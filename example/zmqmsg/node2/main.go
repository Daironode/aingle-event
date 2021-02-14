
package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/Daironode/aingle-event/actor"
	"github.com/Daironode/aingle-event/mailbox"
	"github.com/Daironode/aingle-event/zmqremote"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU() * 1)
	runtime.GC()

	zmqremote.Start("127.0.0.1:8080")

	props := actor.
		FromFunc(
			func(context actor.Context) {
				switch context.Message().(type) {
				case *zmqremote.MsgData:
					context.Sender().Tell(&zmqremote.MsgData{MsgType: 1, Data: []byte("123")})
				}
			}).
		WithMailbox(mailbox.Bounded(1000000))

	pid, _ := actor.SpawnNamed(props, "remote")
	fmt.Println(pid)

	for {
		time.Sleep(1 * time.Second)
	}
}
