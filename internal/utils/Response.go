package utils

type Response[T any] struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func Format[T any](status int, message string, data T) *Response[T] {
	return &Response[T]{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
