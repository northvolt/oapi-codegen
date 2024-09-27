// Package xgojsonignore provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/northvolt/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package xgojsonignore

// Client defines model for Client.
type Client struct {
	ComplexField *struct {
		AccountName *string `json:"accountName,omitempty"`
		Name        *string `json:"name,omitempty"`
	} `json:"complexField,omitempty"`
	Name string `json:"name"`
}

// ClientWithExtension defines model for ClientWithExtension.
type ClientWithExtension struct {
	ComplexField *struct {
		AccountName *string `json:"accountName,omitempty"`
		Name        *string `json:"name,omitempty"`
	} `json:"-"`
	Name string `json:"name"`
}
