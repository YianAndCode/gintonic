package lmdata

import "time"

// "license_key_" event payload
type LicenseKeyPayload struct {
	Meta `json:"meta"`
	Data LicenseKey `json:"data"`
}

type LicenseKey struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		StoreID         int       `json:"store_id"`
		CustomerID      int       `json:"customer_id"`
		OrderID         int       `json:"order_id"`
		OrderItemID     int       `json:"order_item_id"`
		ProductID       int       `json:"product_id"`
		UserName        string    `json:"user_name"`
		UserEmail       string    `json:"user_email"`
		Key             string    `json:"key"`
		KeyShort        string    `json:"key_short"`
		ActivationLimit int       `json:"activation_limit"`
		InstancesCount  int       `json:"instances_count"`
		Disabled        int       `json:"disabled"`
		Status          string    `json:"status"`
		StatusFormatted string    `json:"status_formatted"`
		ExpiresAt       time.Time `json:"expires_at"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
	} `json:"attributes"`
}
