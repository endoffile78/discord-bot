package bot

import (
	"github.com/endoffile78/bot/src/bot/commands"
	"github.com/endoffile78/bot/src/config"
	"github.com/skwair/harmony"
)

type Bot struct {
	Client *harmony.Client
	prefix string
}

func NewBot(client *harmony.Client) *Bot {
	return &Bot{
		Client: client,
		prefix: config.ConfigGet("Bot", "prefix"),
	}
}

func (b Bot) OnNewMessage(msg *harmony.Message) {
	if msg.Author.ID == b.Client.State.CurrentUser().ID {
		return
	}

	cmd, args := commands.CommandParse(b.prefix, msg.Content)
	if cmd == "" {
		return
	}

	channel := b.Client.Channel(msg.ChannelID)
	commands.CommandRun(cmd, args, channel)
}
