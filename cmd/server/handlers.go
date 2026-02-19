package main

import (
	"fmt"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/gamelogic"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
)

func handlerGameLogs() func(log routing.GameLog) pubsub.Acktype {
	return func(log routing.GameLog) pubsub.Acktype {
		defer fmt.Print("> ")

		err := gamelogic.WriteLog(log)
		if err != nil {
			fmt.Printf("could not write log: %v\n", err)
			return pubsub.NackRequeue
		}
		return pubsub.Ack
	}
}
