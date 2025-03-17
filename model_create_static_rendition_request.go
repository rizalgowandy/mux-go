// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type CreateStaticRenditionRequest struct {
	Resolution string `json:"resolution,omitempty"`
	// Arbitrary user-supplied metadata set for the static rendition. Max 255 characters.
	Passthrough string `json:"passthrough,omitempty"`
}
