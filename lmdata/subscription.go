package lmdata

import "time"

// "subscription_" event payload
type SubscriptionPayload struct {
	Meta `json:"meta"`
	Data Subscription `json:"data"`
}

type SubscriptionPause struct {
	Mode      string    `json:"mode"`
	ResumesAt time.Time `json:"resumes_at"`
}

type Subscription struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		StoreID         int               `json:"store_id"`
		CustomerID      int               `json:"customer_id"`
		OrderID         int               `json:"order_id"`
		OrderItemID     int               `json:"order_item_id"`
		ProductID       int               `json:"product_id"`
		VariantID       int               `json:"variant_id"`
		ProductName     string            `json:"product_name"`
		VariantName     string            `json:"variant_name"`
		UserName        string            `json:"user_name"`
		UserEmail       string            `json:"user_email"`
		Status          string            `json:"status"`
		StatusFormatted string            `json:"status_formatted"`
		CardBrand       string            `json:"card_brand"`
		CardLastFour    string            `json:"card_last_four"`
		Pause           SubscriptionPause `json:"pause"`
		Cancelled       bool              `json:"cancelled"`
		TrialEndsAt     time.Time         `json:"trial_ends_at"`
		BillingAnchor   int               `json:"billing_anchor"`
		Urls            struct {
			UpdatePaymentMethod string `json:"update_payment_method"`
		} `json:"urls"`
		RenewsAt  time.Time `json:"renews_at"`
		EndsAt    time.Time `json:"ends_at"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		TestMode  bool      `json:"test_mode"`
	} `json:"attributes"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Relationships struct {
		Order struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"order"`
		Store struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"store"`
		Product struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"product"`
		Variant struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"variant"`
		Customer struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"customer"`
		OrderItem struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"order-item"`
		SubscriptionItems struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"subscription-items"`
		SubscriptionInvoices struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"subscription-invoices"`
	}
}
