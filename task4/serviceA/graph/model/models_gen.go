// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type PhoneVerificationRequest struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

type PhoneVerificationResponse struct {
	Verified bool `json:"verified"`
}
