package bot

import (
	"context"
	"fmt"

	"github.com/andersfylling/disgord"
	"github.com/endoffile78/bot/bot/commands"
	"github.com/endoffile78/bot/config"
)

type Bot struct {
	client *disgord.Client
	prefix string
}

func NewBot(client *disgord.Client) *Bot {
	return &Bot{
		client: client,
		prefix: config.ConfigGet("Bot", "prefix"),
	}
}

func (b Bot) onReady() {
	fmt.Println("Bot ready!")
}

func (b Bot) Run() {
	defer b.client.StayConnectedUntilInterrupted(context.Background())

	invite, err := b.client.InviteURL(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Invite: ", invite)

	b.client.On(disgord.EvtReady, b.onReady)
	b.client.On(disgord.EvtMessageCreate, b.onNewMessage)
	b.client.On(disgord.EvtGuildMemberAdd, b.onJoin)
}

func (b Bot) onNewMessage(session disgord.Session, evt *disgord.MessageCreate) {
	msg := evt.Message

	user, err := session.GetCurrentUser(context.Background())
	if err != nil {
		fmt.Println("Unable to get user")
		return
	}

	if msg.Author.ID == user.ID {
		return
	}

	cmd, args := commands.CommandParse(b.prefix, msg.Content)
	if cmd == "" {
		return
	}

	commands.CommandRun(cmd, args, msg, session)
}

func (b Bot) onJoin(session disgord.Session, evt *disgord.GuildMemberAdd) {

}
