// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type Broadcast struct {
	// Unique identifier for the broadcast. Max 255 characters.
	Id string `json:"id"`
	// Arbitrary user-supplied metadata that will be included in the broadcast details and related webhooks. Max: 255 characters.
	Passthrough string `json:"passthrough,omitempty"`
	// The ID of the live stream that the broadcast will be sent to.
	LiveStreamId string          `json:"live_stream_id"`
	Status       BroadcastStatus `json:"status"`
	Layout       BroadcastLayout `json:"layout"`
	// URL of an image to display as the background of the broadcast. Its dimensions should match the provided resolution.
	Background string              `json:"background,omitempty"`
	Resolution BroadcastResolution `json:"resolution"`
}
