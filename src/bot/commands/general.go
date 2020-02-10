package commands

import (
	"context"

	"github.com/andersfylling/disgord"
)

func ping(args []string, msg *disgord.Message, session disgord.Session) {
	msg.Reply(context.Background(), session, "pong")
}
