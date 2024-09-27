// Package anyofallofoneof provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/northvolt/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package anyofallofoneof

import (
	"encoding/json"

	"github.com/oapi-codegen/runtime"
)

// Client defines model for Client.
type Client struct {
	Name string `json:"name"`
}

// ClientAndMaybeIdentity defines model for ClientAndMaybeIdentity.
type ClientAndMaybeIdentity struct {
	union json.RawMessage
}

// ClientOrIdentity defines model for ClientOrIdentity.
type ClientOrIdentity struct {
	union json.RawMessage
}

// ClientWithId defines model for ClientWithId.
type ClientWithId struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// Identity defines model for Identity.
type Identity struct {
	Issuer string `json:"issuer"`
}

// IdentityWithDuplicateField defines model for IdentityWithDuplicateField.
type IdentityWithDuplicateField struct {
	Issuer struct {
		Name string `json:"name"`
	} `json:"issuer"`
}

// AsClient returns the union data inside the ClientAndMaybeIdentity as a Client
func (t ClientAndMaybeIdentity) AsClient() (Client, error) {
	var body Client
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromClient overwrites any union data inside the ClientAndMaybeIdentity as the provided Client
func (t *ClientAndMaybeIdentity) FromClient(v Client) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeClient performs a merge with any union data inside the ClientAndMaybeIdentity, using the provided Client
func (t *ClientAndMaybeIdentity) MergeClient(v Client) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsIdentity returns the union data inside the ClientAndMaybeIdentity as a Identity
func (t ClientAndMaybeIdentity) AsIdentity() (Identity, error) {
	var body Identity
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromIdentity overwrites any union data inside the ClientAndMaybeIdentity as the provided Identity
func (t *ClientAndMaybeIdentity) FromIdentity(v Identity) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeIdentity performs a merge with any union data inside the ClientAndMaybeIdentity, using the provided Identity
func (t *ClientAndMaybeIdentity) MergeIdentity(v Identity) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t ClientAndMaybeIdentity) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ClientAndMaybeIdentity) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsClient returns the union data inside the ClientOrIdentity as a Client
func (t ClientOrIdentity) AsClient() (Client, error) {
	var body Client
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromClient overwrites any union data inside the ClientOrIdentity as the provided Client
func (t *ClientOrIdentity) FromClient(v Client) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeClient performs a merge with any union data inside the ClientOrIdentity, using the provided Client
func (t *ClientOrIdentity) MergeClient(v Client) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsIdentity returns the union data inside the ClientOrIdentity as a Identity
func (t ClientOrIdentity) AsIdentity() (Identity, error) {
	var body Identity
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromIdentity overwrites any union data inside the ClientOrIdentity as the provided Identity
func (t *ClientOrIdentity) FromIdentity(v Identity) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeIdentity performs a merge with any union data inside the ClientOrIdentity, using the provided Identity
func (t *ClientOrIdentity) MergeIdentity(v Identity) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t ClientOrIdentity) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ClientOrIdentity) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}
