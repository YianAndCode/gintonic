package lmdata

type Meta struct {
	TestMode   bool           `json:"test_mode"`
	EventName  string         `json:"event_name"`
	CustomData map[string]any `json:"custom_data,omitempty"`
}
