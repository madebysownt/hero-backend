package main

type Response struct {
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Error   any    `json:"error,omitempty"`
}
