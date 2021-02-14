
package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/Daironode/aingle-event/actor"
	"github.com/Daironode/aingle-event/example/remotebenchmark/messages"
	"github.com/Daironode/aingle-event/mailbox"
	"github.com/Daironode/aingle-event/remote"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 1)
	runtime.GC()

	remote.Start("127.0.0.1:8080")
	var sender *actor.PID
	props := actor.
		FromFunc(
			func(context actor.Context) {
				switch msg := context.Message().(type) {
				case *messages.StartRemote:
					fmt.Println("Starting")
					sender = msg.Sender
					context.Respond(&messages.Start{})
				case *messages.Ping:
					sender.Tell(&messages.Pong{})
				}
			}).
		WithMailbox(mailbox.Bounded(1000000))
	actor.SpawnNamed(props, "remote")
	for {
		time.Sleep(1 * time.Second)
	}
}
