package lmdata

import "time"

// "subscription_payment_" event payload
type SubscriptionInvoicePayload struct {
	Meta `json:"meta"`
	Data SubscriptionInvoice `json:"data"`
}

type SubscriptionInvoice struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		StoreID                int       `json:"store_id"`
		SubscriptionID         int       `json:"subscription_id"`
		BillingReason          string    `json:"billing_reason"`
		CardBrand              string    `json:"card_brand"`
		CardLastFour           string    `json:"card_last_four"`
		Currency               string    `json:"currency"`
		CurrencyRate           string    `json:"currency_rate"`
		Subtotal               int       `json:"subtotal"`
		DiscountTotal          int       `json:"discount_total"`
		Tax                    int       `json:"tax"`
		Total                  int       `json:"total"`
		SubtotalUsd            int       `json:"subtotal_usd"`
		DiscountTotalUsd       int       `json:"discount_total_usd"`
		TaxUsd                 int       `json:"tax_usd"`
		TotalUsd               int       `json:"total_usd"`
		Status                 string    `json:"status"`
		StatusFormatted        string    `json:"status_formatted"`
		Refunded               bool      `json:"refunded"`
		RefundedAt             time.Time `json:"refunded_at"`
		SubtotalFormatted      string    `json:"subtotal_formatted"`
		DiscountTotalFormatted string    `json:"discount_total_formatted"`
		TaxFormatted           string    `json:"tax_formatted"`
		TotalFormatted         string    `json:"total_formatted"`
		UserName               string    `json:"user_name"`
		UserEmail              string    `json:"user_email"`
		CustomerID             int       `json:"customer_id"`
		Urls                   struct {
			InvoiceURL string `json:"invoice_url"`
		} `json:"urls"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		TestMode  bool      `json:"test_mode"`
	} `json:"attributes"`
	Relationships struct {
		Store struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"store"`
		Customer struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"customer"`
		Subscription struct {
			Links struct {
				Self    string `json:"self"`
				Related string `json:"related"`
			} `json:"links"`
		} `json:"subscription"`
	} `json:"relationships"`
}
