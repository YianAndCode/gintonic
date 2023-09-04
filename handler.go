package gintonic

import (
	"context"
	"gintonic/lmdata"
)

type DefaultHandler func(ctx context.Context, eventName string, meta lmdata.Meta, data interface{}) error

type OrderCreatedHandler func(ctx context.Context, meta lmdata.Meta, data lmdata.Order) error
type OrderRefundedHandler func(ctx context.Context, meta lmdata.Meta, data lmdata.Order) error

type SubscriptionCreatedHandler func(ctx context.Context, meta lmdata.Meta, data lmdata.Subscription) error
type SubscriptionUpdatedHandler func(ctx context.Context, meta lmdata.Meta, data lmdata.Subscription) error
type SubscriptionCancelledHandler func(ctx context.Context, meta lmdata.Meta, data lmdata.Subscription) error
type SubscriptionResumedHandler func(ctx context.Context, meta lmdata.Meta, data lmdata.Subscription) error
type SubscriptionExpiredHandler func(ctx context.Context, meta lmdata.Meta, data lmdata.Subscription) error
type SubscriptionPausedHandler func(ctx context.Context, meta lmdata.Meta, data lmdata.Subscription) error
type SubscriptionUnpausedHandler func(ctx context.Context, meta lmdata.Meta, data lmdata.Subscription) error

type SubscriptionPaymentSuccessHandler func(ctx context.Context, meta lmdata.Meta, data lmdata.SubscriptionInvoice) error
type SubscriptionPaymentFailedHandler func(ctx context.Context, meta lmdata.Meta, data lmdata.SubscriptionInvoice) error
type SubscriptionPaymentRecoveredHandler func(ctx context.Context, meta lmdata.Meta, data lmdata.SubscriptionInvoice) error

type LicenseKeyCreatedHandler func(ctx context.Context, meta lmdata.Meta, data lmdata.LicenseKey) error
type LicenseKeyUpdatedHandler func(ctx context.Context, meta lmdata.Meta, data lmdata.LicenseKey) error
