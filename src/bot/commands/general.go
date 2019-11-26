package commands

import (
	"context"
	"math/rand"
	"strconv"

	"github.com/skwair/harmony"
)

func ping(args []string, channel *harmony.ChannelResource) {
	channel.SendMessage(context.Background(), "pong")
}

func roll(args []string, channel *harmony.ChannelResource) {
	numDice := 1
	diceSize := 6

	switch len(args) {
	case 1:
		var err error
		numDice, err = strconv.Atoi(args[0])
		if err != nil {
			channel.SendMessage(context.Background(), "Invalid argument: "+args[0])
			return
		}
	case 2:
		var err error
		numDice, err = strconv.Atoi(args[0])
		if err != nil {
			channel.SendMessage(context.Background(), "Invalid argument: "+args[0])
			return
		}

		diceSize, err = strconv.Atoi(args[1])
		if err != nil {
			channel.SendMessage(context.Background(), "Invalid argument: "+args[1])
			return
		}
	}

	roll := 0
	for i := 0; i < numDice; i++ {
		roll += rand.Intn(diceSize)
	}

	channel.SendMessage(context.Background(), strconv.Itoa(roll))
}
