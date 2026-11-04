// Command tradenotify 是普通支付通知的最小 net/http 接入示例。
//
// 链路：支付宝 POST 表单 → 有界读取并解码一次 → RSA2 验签 →
// 领域映射 → 官方要求的业务校验（APPID、收款方、订单、金额）→
// 幂等处理 → 应答 success / fail。
//
// 配置全部来自环境变量，示例不包含任何真实凭证：
//
//	ALIPAY_APP_ID          本应用 APPID（必填）
//	ALIPAY_PUBLIC_KEY      支付宝公钥 PEM 或单行 base64（必填）
//	ALIPAY_SELLER_ID       卖家支付宝用户号，多收款账户时必校验
//	ALIPAY_LISTEN_ADDR     监听地址，默认 :8080
//
// 本示例不处理订单持久化；memoryOrders 仅演示幂等判断的接入位置，
// 商户应使用数据库唯一约束或状态机保证幂等。
package main

import (
	"log/slog"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/tiyee/go-alipay/core/auth/verifiers"
	"github.com/tiyee/go-alipay/core/notify"
	"github.com/tiyee/go-alipay/services/trade"
)

func main() {
	appID := os.Getenv("ALIPAY_APP_ID")
	publicKey := os.Getenv("ALIPAY_PUBLIC_KEY")
	sellerID := os.Getenv("ALIPAY_SELLER_ID")
	addr := os.Getenv("ALIPAY_LISTEN_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	if appID == "" || publicKey == "" {
		slog.Error("ALIPAY_APP_ID and ALIPAY_PUBLIC_KEY are required")
		os.Exit(2)
	}
	verifier, err := verifiers.NewRSA2(publicKey)
	if err != nil {
		slog.Error("invalid ALIPAY_PUBLIC_KEY")
		os.Exit(2)
	}

	orders := newMemoryOrders()
	handler := &notifyHandler{
		verifier:  verifier,
		appID:     appID,
		sellerID:  sellerID,
		orders:    orders,
		logger:    slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		notifySvc: trade.NotificationDecoder(),
	}
	// 生产环境应为该端点配置 HTTPS 与访问日志，并保证公网可直达
	// （通知地址不能重定向）。
	server := &http.Server{
		Addr:              addr,
		Handler:           http.HandlerFunc(handler.serveNotify),
		ReadHeaderTimeout: 5 * time.Second,
	}
	slog.Info("listening", "addr", addr)
	if err := server.ListenAndServe(); err != nil {
		slog.Error("server exited", "err_kind", "listen")
		os.Exit(1)
	}
}

type notifyHandler struct {
	verifier  *verifiers.RSA2
	appID     string
	sellerID  string
	orders    *memoryOrders
	logger    *slog.Logger
	notifySvc notify.Decoder[trade.Notification]
}

func (h *notifyHandler) serveNotify(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// ParseRequest 完成有界读取、单次解码、验签与领域映射；
	// 日志注入复用 slog 契约，默认输出仅包含分类字段。
	notification, err := notify.ParseRequest(ctx, r,
		notify.PolicyTradeStatusSync, h.verifier, h.notifySvc,
		notify.WithLogger(h.logger))
	if err != nil {
		// 未知来源的请求：应答 fail，支付宝会按策略重发。
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("fail"))
		return
	}
	data := notification.Data

	// —— 业务校验接入位置：验签通过后仍必须逐项核对官方要求的四项信息 ——
	switch {
	case data.AppID != h.appID:
		h.reject(w, "unknown app_id")
		return
	case h.sellerID != "" && (data.SellerID == nil || *data.SellerID != h.sellerID):
		// 多收款账户的商户必须校验收款方；单账户商户也应核对。
		h.reject(w, "unknown seller")
		return
	case !h.orders.exists(data.OutTradeNo):
		h.reject(w, "unknown order")
		return
	case h.orders.amount(data.OutTradeNo) != data.TotalAmount:
		h.reject(w, "amount mismatch")
		return
	}

	// —— 幂等接入位置：同一通知（NotifyID 不变）重发时只处理一次 ——
	if !h.orders.markProcessed(data.NotifyID) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("success"))
		return
	}

	// 只有 TRADE_SUCCESS / TRADE_FINISHED 表示买家付款成功；
	// 其他状态（TRADE_CLOSED、未知新状态）按业务语义另行处理。
	if data.Paid() {
		h.orders.markPaid(data.OutTradeNo, data.TradeNo)
	}
	// 处理完成后原样应答 success（小写、无空白与 HTML），
	// 否则支付宝会按 4m、10m、10m、1h、2h、6h、15h 重发。
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("success"))
}

func (h *notifyHandler) reject(w http.ResponseWriter, reason string) {
	h.logger.Warn("notify rejected", "reason_kind", reason)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("fail"))
}

// memoryOrders 用内存数据演示订单存在性、金额一致性校验与幂等标记。
type memoryOrders struct {
	mu         sync.Mutex
	amounts    map[string]string
	processed  map[string]bool
	paidOrders map[string]string
}

func newMemoryOrders() *memoryOrders {
	return &memoryOrders{
		amounts:    map[string]string{"ORDER-20261105-001": "88.88"},
		processed:  map[string]bool{},
		paidOrders: map[string]string{},
	}
}

func (m *memoryOrders) exists(outTradeNo string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.amounts[outTradeNo]
	return ok
}

func (m *memoryOrders) amount(outTradeNo string) string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.amounts[outTradeNo]
}

func (m *memoryOrders) markProcessed(notifyID string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.processed[notifyID] {
		return false
	}
	m.processed[notifyID] = true
	return true
}

func (m *memoryOrders) markPaid(outTradeNo, tradeNo string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.paidOrders[outTradeNo] = tradeNo
}
