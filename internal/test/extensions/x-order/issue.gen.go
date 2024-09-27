// Package xorder provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/northvolt/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package xorder

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// DateInterval defines model for DateInterval.
type DateInterval struct {
	Start *openapi_types.Date `json:"start,omitempty"`
	End   *openapi_types.Date `json:"end,omitempty"`
}

// Port defines model for Port.
type Port = int

// PortInterval defines model for PortInterval.
type PortInterval struct {
	Start   Port             `json:"start"`
	End     Port             `json:"end"`
	VeryEnd *LowPriorityPort `json:"very_end,omitempty"`
}

// LowPriorityPort defines model for LowPriorityPort.
type LowPriorityPort = int
