package bot

import (
	"github.com/endoffile78/bot/src/bot/commands"
	"github.com/skwair/harmony"
)

type Bot struct {
	Client *harmony.Client
}

func (b Bot) OnNewMessage(msg *harmony.Message) {
	if msg.Author.ID == b.Client.State.CurrentUser().ID {
		return
	}

	cmd, args := commands.CommandParse(msg.Content)
	if cmd == "" {
		return
	}

	channel := b.Client.Channel(msg.ChannelID)
	commands.CommandRun(cmd, args, channel)
}
