package gintonic

import (
	"encoding/json"
	"errors"

	"github.com/YianAndCode/gintonic/lmdata"
)

const (
	Event_OrderCreated                 string = "order_created"
	Event_OrderRefunded                string = "order_refunded"
	Event_SubscriptionCreated          string = "subscription_created"
	Event_SubscriptionUpdated          string = "subscription_updated"
	Event_SubscriptionCancelled        string = "subscription_cancelled"
	Event_SubscriptionResumed          string = "subscription_resumed"
	Event_SubscriptionExpired          string = "subscription_expired"
	Event_SubscriptionPaused           string = "subscription_paused"
	Event_SubscriptionUnpaused         string = "subscription_unpaused"
	Event_SubscriptionPaymentSuccess   string = "subscription_payment_success"
	Event_SubscriptionPaymentFailed    string = "subscription_payment_failed"
	Event_SubscriptionPaymentRecovered string = "subscription_payment_recovered"
	Event_LicenseKeyCreated            string = "license_key_created"
	Event_LicenseKeyUpdated            string = "license_key_updated"
)

func ParseEventPayload(eventName string, payload []byte) (any, error) {
	switch eventName {
	case Event_OrderCreated:
		fallthrough
	case Event_OrderRefunded:
		payloadObj := lmdata.OrderPayload{}
		err := json.Unmarshal(payload, &payloadObj)
		return payloadObj, err
	case Event_SubscriptionCreated:
		fallthrough
	case Event_SubscriptionUpdated:
		fallthrough
	case Event_SubscriptionCancelled:
		fallthrough
	case Event_SubscriptionResumed:
		fallthrough
	case Event_SubscriptionExpired:
		fallthrough
	case Event_SubscriptionPaused:
		fallthrough
	case Event_SubscriptionUnpaused:
		payloadObj := lmdata.SubscriptionPayload{}
		err := json.Unmarshal(payload, &payloadObj)
		return payloadObj, err
	case Event_SubscriptionPaymentSuccess:
		fallthrough
	case Event_SubscriptionPaymentFailed:
		fallthrough
	case Event_SubscriptionPaymentRecovered:
		payloadObj := lmdata.SubscriptionInvoicePayload{}
		err := json.Unmarshal(payload, &payloadObj)
		return payloadObj, err
	case Event_LicenseKeyCreated:
		fallthrough
	case Event_LicenseKeyUpdated:
		payloadObj := lmdata.LicenseKeyPayload{}
		err := json.Unmarshal(payload, &payloadObj)
		return payloadObj, err
	}
	return nil, errors.New("invalid event:" + string(eventName))
}
