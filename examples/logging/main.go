// Command logging 演示 SDK 的 slog 日志接入契约：
//
//  1. 默认静默：Client 未注入 Logger 时不产生任何输出；
//  2. 标准 slog：option.WithLogger 搭配标准库 TextHandler / JSONHandler；
//  3. 自定义 Handler：option.WithLogHandler 注入任意 slog.Handler 实现，
//     演示字段过滤与调用方对日志资源的所有权；
//  4. 日志失败不影响调用：Handler 写出错误只被丢弃，API 结果保持不变。
//
// SDK 只记录筛选后的字段（sdk、operation、duration、error_kind、
// http_status、code），不输出完整报文、URL 查询参数、签名原文或错误文本；
// 日志级别约定：平台业务拒绝为 Warn，成功与主动取消为 Debug，
// 网络、协议与安全失败为 Error。
//
// 示例使用本地 httptest 服务模拟网关，不访问支付宝、不包含真实凭证。
package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/core/option"
	"github.com/tiyee/alipay-sdk-go/services/trade"
)

func main() {
	// 模拟网关：始终返回未签名的 V3 错误响应（非 2xx + {code, message}，
	// 业务拒绝）。SDK 将其按 Verified=false 的 *core.APIError 处理，
	// 日志级别为 Warn。
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"code":"ACQ.TRADE_NOT_EXIST","message":"交易不存在"}`)
	}))
	defer gateway.Close()

	credentials := demoCredentials()
	ctx := context.Background()

	fmt.Println("== 1. 默认静默：不注入 Logger，SDK 不输出任何日志 ==")
	client := newClient(credentials, gateway.URL)
	callOnce(ctx, client)
	fmt.Printf("（以上没有任何 SDK 日志；调用错误仍正常返回）\n\n")

	fmt.Println("== 2. 标准 slog：WithLogger + TextHandler（Info 级别）==")
	client = newClient(credentials, gateway.URL,
		option.WithLogger(slog.New(slog.NewTextHandler(os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelInfo}))))
	callOnce(ctx, client)
	fmt.Printf("（Warn=业务拒绝；成功与取消为 Debug，Info 级别下不显示；网络/协议/安全错误为 Error）\n\n")

	fmt.Println("== 3. 自定义 Handler：WithLogHandler + 字段过滤，写到调用方持有的文件 ==")
	logFile, err := os.CreateTemp("", "alipay-log-*.txt")
	if err != nil {
		fmt.Println("create temp log file:", err)
		os.Exit(1)
	}
	defer func() {
		logFile.Close()
		os.Remove(logFile.Name())
	}()
	client = newClient(credentials, gateway.URL,
		option.WithLogHandler(newFilteredHandler(logFile, "operation", "error_kind")))
	callOnce(ctx, client)
	content, err := os.ReadFile(logFile.Name())
	if err != nil {
		fmt.Println("read temp log file:", err)
		os.Exit(1)
	}
	fmt.Printf("调用方关闭文件前的日志内容：\n%s\n", strings.TrimRight(string(content), "\n"))
	// SDK 不会关闭调用方提供的日志资源；上面 defer 的 logFile.Close()
	// 由本程序负责，覆盖日志轮转、文件生命周期等接入方自己的策略。

	fmt.Println("== 4. Handler 写出失败不影响调用结果 ==")
	client = newClient(credentials, gateway.URL,
		option.WithLogHandler(newFilteredHandler(brokenWriter{}, "operation")))
	err = callOnce(ctx, client)
	var apiErr *core.APIError
	if errors.As(err, &apiErr) {
		fmt.Printf("API 结果不受日志故障影响：code=%s verified=%v\n",
			apiErr.Code, apiErr.Verified)
	}
}

// newClient 构建指向本地模拟网关的 Client；opts 追加日志选项。
func newClient(credentials auth.Credentials, gateway string,
	opts ...core.ClientOption) *core.Client {
	base := []core.ClientOption{
		option.WithAppID("your-app-id"),
		option.WithGateway(gateway),
	}
	client, err := core.NewClient(credentials, append(base, opts...)...)
	if err != nil {
		fmt.Println("build client:", err)
		os.Exit(1)
	}
	return client
}

// callOnce 发起一次查询调用，返回业务错误供第 4 节断言。
func callOnce(ctx context.Context, client *core.Client) error {
	service, err := trade.NewService(client)
	if err != nil {
		fmt.Println("build trade service:", err)
		os.Exit(1)
	}
	_, err = service.Query(ctx, trade.QueryRequest{OutTradeNo: "demo_20261105_001"})
	return err
}

// demoCredentials 生成本地演示用的 RSA2 公钥凭证；
// 测试密钥仅用于本示例与本地模拟网关，不代表任何真实应用。
func demoCredentials() auth.Credentials {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("generate demo key:", err)
		os.Exit(1)
	}
	credentials, err := auth.NewPublicKeyCredentials(auth.PublicKeyConfig{
		AppSigner:       key,
		AlipayPublicKey: &key.PublicKey,
	})
	if err != nil {
		fmt.Println("build credentials:", err)
		os.Exit(1)
	}
	return credentials
}

// filteredHandler 是演示用的最小 slog.Handler：只输出白名单内的属性，
// 其余属性一律丢弃；输出格式为“消息  key=value ...”。
type filteredHandler struct {
	mu     *sync.Mutex
	w      io.Writer
	allow  map[string]struct{}
	preset []slog.Attr
}

// newFilteredHandler 构建 Handler；writer 由调用方创建并关闭。
func newFilteredHandler(w io.Writer, allow ...string) *filteredHandler {
	keys := make(map[string]struct{}, len(allow))
	for _, key := range allow {
		keys[key] = struct{}{}
	}
	return &filteredHandler{mu: &sync.Mutex{}, w: w, allow: keys}
}

// Enabled 实现 slog.Handler：本演示处理所有级别，由白名单决定输出内容。
func (h *filteredHandler) Enabled(context.Context, slog.Level) bool { return true }

// Handle 实现 slog.Handler：写出失败时错误经 slog.Logger 丢弃，
// 不会传播回 API 调用（见第 4 节演示）。
func (h *filteredHandler) Handle(_ context.Context, record slog.Record) error {
	var line strings.Builder
	line.WriteString(record.Message)
	appendAllowed(&line, h.preset, h.allow)
	record.Attrs(func(attr slog.Attr) bool {
		appendAllowed(&line, []slog.Attr{attr}, h.allow)
		return true
	})
	line.WriteByte('\n')
	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := io.WriteString(h.w, line.String())
	return err
}

// WithAttrs 实现 slog.Handler：返回携带预置属性的副本，保持不可变复用。
func (h *filteredHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	clone := *h
	clone.preset = append(append([]slog.Attr(nil), h.preset...), attrs...)
	return &clone
}

// WithGroup 实现 slog.Handler：演示实现不处理分组。
func (h *filteredHandler) WithGroup(string) slog.Handler { return h }

func appendAllowed(line *strings.Builder, attrs []slog.Attr,
	allow map[string]struct{}) {
	for _, attr := range attrs {
		if _, ok := allow[attr.Key]; ok {
			line.WriteString("  ")
			line.WriteString(attr.Key)
			line.WriteString("=")
			line.WriteString(attr.Value.String())
		}
	}
}

// brokenWriter 是始终失败的日志写出目标，用于演示 Handler 错误被丢弃。
type brokenWriter struct{}

// Write 实现 io.Writer。
func (brokenWriter) Write([]byte) (int, error) { return 0, errors.New("log sink down") }
