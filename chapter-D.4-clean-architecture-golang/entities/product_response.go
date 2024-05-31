package entities

// Struct request and response product
type ProductResponseJSON struct {
	Data    []Products `json:"data"`
	Count   int64      `json:"count"`
	Success bool       `json:"success"`
	Message string     `json:"message"`
}
