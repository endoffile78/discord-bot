package commands

import (
	"strings"

	"github.com/andersfylling/disgord"
)

type CommandFunction func(args []string, msg *disgord.Message, session disgord.Session)

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
	}
)

func CommandParse(prefix string, message string) (string, []string) {
	if !strings.HasPrefix(message, prefix) {
		return "", nil
	}

	cmd := ""
	var args []string = nil

	message = strings.TrimSpace(message)

	spaceIndex := strings.Index(message, " ")
	if spaceIndex > 0 {
		cmd = message[len(prefix):spaceIndex]
		args = strings.Split(message, " ")[1:]
	} else {
		cmd = message[len(prefix):]
	}

	return cmd, args
}

func CommandRun(cmd string, args []string, msg *disgord.Message, session disgord.Session) {
	for _, command := range commands {
		if command.cmd == cmd {
			command.function(args, msg, session)
			break
		}
	}
}
