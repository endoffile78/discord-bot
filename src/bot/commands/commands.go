package commands

import (
	"strings"

	"github.com/endoffile78/bot/src/config"
	"github.com/skwair/harmony"
)

type CommandFunction func(args []string, channel *harmony.ChannelResource)

type Command struct {
	cmd      string
	function CommandFunction
	admin    bool
}

var (
	commands = []Command{
		Command{
			cmd:      "ping",
			function: ping,
			admin:    false,
		},
		Command{
			cmd:      "roll",
			function: roll,
			admin:    false,
		},
	}
)

func CommandParse(message string) (string, []string) {
	prefix := config.ConfigGet("Bot", "prefix")
	if !strings.HasPrefix(message, prefix) {
		return "", nil
	}

	cmd := ""
	var args []string = nil

	spaceIndex := strings.Index(message, " ")
	if spaceIndex > 0 {
		cmd = message[len(prefix):spaceIndex]
		args = strings.Split(message, " ")[1:]
	} else {
		cmd = message[len(prefix):]
	}

	return cmd, args
}

func CommandRun(cmd string, args []string, channel *harmony.ChannelResource) {
	for _, command := range commands {
		if command.cmd == cmd {
			command.function(args, channel)
			break
		}
	}
}
