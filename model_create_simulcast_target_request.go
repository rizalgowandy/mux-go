// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type CreateSimulcastTargetRequest struct {
	// Arbitrary metadata set by you when creating a simulcast target.
	Passthrough string `json:"passthrough,omitempty"`
	// Stream Key represents a stream identifier on the third party live streaming service to send the parent live stream to.
	StreamKey string `json:"stream_key,omitempty"`
	// RTMP hostname including application name for the third party live streaming service. Example: 'rtmp://live.example.com/app'.
	Url string `json:"url"`
}
