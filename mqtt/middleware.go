package mqtt

import (
	"context"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// Handler is a function that processes a Pub/Sub message and returns an error
type Handler func(ctx context.Context, msg MQTT.Message) error

// Middleware is a Pub/Sub middleware
type Middleware func(Handler) Handler

// nopMiddleWare is the default middleware, and does nothing.
func nopMiddleWare(h Handler) Handler {
	return h
}
