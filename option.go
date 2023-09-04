package gintonic

type ginTonicOption interface {
	apply(*GinTonic)
}

type handlerOption struct {
	handlerType string
	handlerFunc any
}

func (o handlerOption) apply(g *GinTonic) {
	g.handlers[o.handlerType] = o.handlerFunc
}

func WithDefaultHandler(handler DefaultHandler) ginTonicOption {
	return handlerOption{
		handlerType: "default",
		handlerFunc: handler,
	}
}

func WithOrderCreatedHandler(handler OrderCreatedHandler) ginTonicOption {
	return handlerOption{
		handlerType: "order_created",
		handlerFunc: handler,
	}
}

func WithOrderRefundedHandler(handler OrderRefundedHandler) ginTonicOption {
	return handlerOption{
		handlerType: "order_refunded",
		handlerFunc: handler,
	}
}

func WithSubscriptionCreatedHandler(handler SubscriptionCreatedHandler) ginTonicOption {
	return handlerOption{
		handlerType: "subscription_created",
		handlerFunc: handler,
	}
}

func WithSubscriptionUpdatedHandler(handler SubscriptionUpdatedHandler) ginTonicOption {
	return handlerOption{
		handlerType: "subscription_updated",
		handlerFunc: handler,
	}
}
func WithSubscriptionCancelledHandler(handler SubscriptionCancelledHandler) ginTonicOption {
	return handlerOption{
		handlerType: "subscription_cancelled",
		handlerFunc: handler,
	}
}
func WithSubscriptionResumedHandler(handler SubscriptionResumedHandler) ginTonicOption {
	return handlerOption{
		handlerType: "subscription_resumed",
		handlerFunc: handler,
	}
}
func WithSubscriptionExpiredHandler(handler SubscriptionExpiredHandler) ginTonicOption {
	return handlerOption{
		handlerType: "subscription_expired",
		handlerFunc: handler,
	}
}
func WithSubscriptionPausedHandler(handler SubscriptionPausedHandler) ginTonicOption {
	return handlerOption{
		handlerType: "subscription_paused",
		handlerFunc: handler,
	}
}
func WithSubscriptionUnpausedHandler(handler SubscriptionUnpausedHandler) ginTonicOption {
	return handlerOption{
		handlerType: "subscription_unpaused",
		handlerFunc: handler,
	}
}

func WithSubscriptionPaymentSuccessHandler(handler SubscriptionPaymentSuccessHandler) ginTonicOption {
	return handlerOption{
		handlerType: "subscription_payment_success",
		handlerFunc: handler,
	}
}
func WithSubscriptionPaymentFailedHandler(handler SubscriptionPaymentFailedHandler) ginTonicOption {
	return handlerOption{
		handlerType: "subscription_payment_failed",
		handlerFunc: handler,
	}
}
func WithSubscriptionPaymentRecoveredHandler(handler SubscriptionPaymentRecoveredHandler) ginTonicOption {
	return handlerOption{
		handlerType: "subscription_payment_recovered",
		handlerFunc: handler,
	}
}

func WithLicenseKeyCreatedHandler(handler LicenseKeyCreatedHandler) ginTonicOption {
	return handlerOption{
		handlerType: "license_key_created",
		handlerFunc: handler,
	}
}
func WithLicenseKeyUpdatedHandler(handler LicenseKeyUpdatedHandler) ginTonicOption {
	return handlerOption{
		handlerType: "license_key_updated",
		handlerFunc: handler,
	}
}
