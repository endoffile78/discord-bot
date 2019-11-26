package commands

import (
	"context"

	"github.com/skwair/harmony"
)

func ping(args []string, channel *harmony.ChannelResource) {
	channel.SendMessage(context.Background(), "pong")
}
