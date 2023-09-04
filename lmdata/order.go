package lmdata

import "time"

// "order_" event payload
type OrderPayload struct {
	Meta `json:"meta"`
	Data Order `json:"data"`
}

const (
	OrderStatus_Pending  string = "pending"
	OrderStatus_Failed   string = "failed"
	OrderStatus_Paid     string = "paid"
	OrderStatus_Refunded string = "refunded"
)

type Order struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		StoreID                int       `json:"store_id"`
		CustomerID             int       `json:"customer_id"`
		Identifier             string    `json:"identifier"`
		OrderNumber            int       `json:"order_number"`
		UserName               string    `json:"user_name"`
		UserEmail              string    `json:"user_email"`
		Currency               string    `json:"currency"`
		CurrencyRate           string    `json:"currency_rate"`
		TaxName                string    `json:"tax_name"`
		TaxRate                string    `json:"tax_rate"`
		Status                 string    `json:"status"`
		StatusFormatted        string    `json:"status_formatted"`
		Refunded               bool      `json:"refunded"`
		RefundedAt             time.Time `json:"refunded_at"`
		Subtotal               int       `json:"subtotal"`
		DiscountTotal          int       `json:"discount_total"`
		Tax                    int       `json:"tax"`
		Total                  int       `json:"total"`
		SubtotalUsd            int       `json:"subtotal_usd"`
		DiscountTotalUsd       int       `json:"discount_total_usd"`
		TaxUsd                 int       `json:"tax_usd"`
		TotalUsd               int       `json:"total_usd"`
		SubtotalFormatted      string    `json:"subtotal_formatted"`
		DiscountTotalFormatted string    `json:"discount_total_formatted"`
		TaxFormatted           string    `json:"tax_formatted"`
		TotalFormatted         string    `json:"total_formatted"`
		FirstOrderItem         struct {
			ID          int       `json:"id"`
			OrderID     int       `json:"order_id"`
			ProductID   int       `json:"product_id"`
			VariantID   int       `json:"variant_id"`
			ProductName string    `json:"product_name"`
			VariantName string    `json:"variant_name"`
			Price       int       `json:"price"`
			CreatedAt   time.Time `json:"created_at"`
			UpdatedAt   time.Time `json:"updated_at"`
			TestMode    bool      `json:"test_mode"`
		} `json:"first_order_item"`
		Urls struct {
			Receipt string `json:"receipt"`
		} `json:"urls"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		TestMode  bool      `json:"test_mode"`
	} `json:"attributes"`
	Relationships struct {
		Store struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"store"`
		Customer struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"customer"`
		OrderItems struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"order-items"`
		Subscriptions struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"subscriptions"`
		LicenseKeys struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"license-keys"`
		DiscountRedemptions struct {
			Links struct {
				Related string `json:"related"`
				Self    string `json:"self"`
			} `json:"links"`
		} `json:"discount-redemptions"`
	} `json:"relationships"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}
