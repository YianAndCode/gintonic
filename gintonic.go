package gintonic

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/http"

	"github.com/YianAndCode/gintonic/lmdata"
	"github.com/gin-gonic/gin"
)

type GinTonic struct {
	secret string

	handlers map[string]any
}

func New(secret string, opts ...ginTonicOption) *GinTonic {
	g := GinTonic{
		secret:   secret,
		handlers: make(map[string]any),
	}

	for _, opt := range opts {
		opt.apply(&g)
	}

	return &g
}

func (g *GinTonic) LemonSqueezyWebhook(c *gin.Context) {
	payload, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	eventName := c.GetHeader("X-Event-Name")
	signature := c.GetHeader("X-Signature")

	if !checkSign(g.secret, signature, payload) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid signature."})
		return
	}

	payloadObj, err := ParseEventPayload(eventName, payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	err = g.handle(c, eventName, payloadObj)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func checkSign(secret, signature string, payload []byte) bool {
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write(payload)
	h := hash.Sum(nil)
	signatureBytes, _ := hex.DecodeString(signature)
	return hmac.Equal(h, signatureBytes)
}

func (g *GinTonic) handle(ctx *gin.Context, eventName string, payload any) (err error) {
	handler, exist := g.handlers[eventName]
	if exist {
		switch eventName {
		case Event_OrderCreated:
			h := handler.(OrderCreatedHandler)
			d := payload.(lmdata.OrderPayload)
			err = h(ctx, d.Meta, d.Data)

		case Event_OrderRefunded:
			h := handler.(OrderRefundedHandler)
			d := payload.(lmdata.OrderPayload)
			err = h(ctx, d.Meta, d.Data)

		case Event_SubscriptionCreated:
			h := handler.(SubscriptionCreatedHandler)
			d := payload.(lmdata.SubscriptionPayload)
			err = h(ctx, d.Meta, d.Data)

		case Event_SubscriptionUpdated:
			h := handler.(SubscriptionUpdatedHandler)
			d := payload.(lmdata.SubscriptionPayload)
			err = h(ctx, d.Meta, d.Data)

		case Event_SubscriptionCancelled:
			h := handler.(SubscriptionCancelledHandler)
			d := payload.(lmdata.SubscriptionPayload)
			err = h(ctx, d.Meta, d.Data)

		case Event_SubscriptionResumed:
			h := handler.(SubscriptionResumedHandler)
			d := payload.(lmdata.SubscriptionPayload)
			err = h(ctx, d.Meta, d.Data)

		case Event_SubscriptionExpired:
			h := handler.(SubscriptionExpiredHandler)
			d := payload.(lmdata.SubscriptionPayload)
			err = h(ctx, d.Meta, d.Data)

		case Event_SubscriptionPaused:
			h := handler.(SubscriptionPausedHandler)
			d := payload.(lmdata.SubscriptionPayload)
			err = h(ctx, d.Meta, d.Data)

		case Event_SubscriptionUnpaused:
			h := handler.(SubscriptionUnpausedHandler)
			d := payload.(lmdata.SubscriptionPayload)
			err = h(ctx, d.Meta, d.Data)

		case Event_SubscriptionPaymentSuccess:
			h := handler.(SubscriptionPaymentSuccessHandler)
			d := payload.(lmdata.SubscriptionInvoicePayload)
			err = h(ctx, d.Meta, d.Data)

		case Event_SubscriptionPaymentFailed:
			h := handler.(SubscriptionPaymentFailedHandler)
			d := payload.(lmdata.SubscriptionInvoicePayload)
			err = h(ctx, d.Meta, d.Data)

		case Event_SubscriptionPaymentRecovered:
			h := handler.(SubscriptionPaymentRecoveredHandler)
			d := payload.(lmdata.SubscriptionInvoicePayload)
			err = h(ctx, d.Meta, d.Data)

		case Event_LicenseKeyCreated:
			h := handler.(LicenseKeyCreatedHandler)
			d := payload.(lmdata.LicenseKeyPayload)
			err = h(ctx, d.Meta, d.Data)

		case Event_LicenseKeyUpdated:
			h := handler.(LicenseKeyUpdatedHandler)
			d := payload.(lmdata.LicenseKeyPayload)
			err = h(ctx, d.Meta, d.Data)
		}
	} else {
		defaultHandler, exist := g.handlers["default"]
		if !exist {
			err = errors.New("handler not found")
			return
		}
		h := defaultHandler.(DefaultHandler)
		meta, data := g.getPayloadDatas(payload)
		err = h(ctx, eventName, meta, data)
	}

	return
}

func (g *GinTonic) getPayloadDatas(payload any) (lmdata.Meta, any) {
	if p, ok := payload.(lmdata.OrderPayload); ok {
		return p.Meta, p.Data
	}

	if p, ok := payload.(lmdata.SubscriptionPayload); ok {
		return p.Meta, p.Data
	}

	if p, ok := payload.(lmdata.SubscriptionInvoicePayload); ok {
		return p.Meta, p.Data
	}

	if p, ok := payload.(lmdata.LicenseKeyPayload); ok {
		return p.Meta, p.Data
	}

	return lmdata.Meta{}, nil
}
