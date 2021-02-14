
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/Daironode/aingle-event/actor"
	"github.com/Daironode/aingle-event/mailbox"
	"github.com/Daironode/aingle-event/zmqremote"
)

type localActor struct {
	count        int
	wgStop       *sync.WaitGroup
	messageCount int
}

func (state *localActor) Receive(context actor.Context) {
	switch context.Message().(type) {
	case *zmqremote.MsgData:
		state.count++
		if state.count%50 == 0 {
			fmt.Println(state.count)
		}
		if state.count == state.messageCount {
			state.wgStop.Done()
		}
	}
}

func newLocalActor(stop *sync.WaitGroup, messageCount int) actor.Producer {
	return func() actor.Actor {
		return &localActor{
			wgStop:       stop,
			messageCount: messageCount,
		}
	}
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU() * 1)
	runtime.GC()

	var wg sync.WaitGroup

	messageCount := 500

	zmqremote.Start("127.0.0.1:8081")

	props := actor.
		FromProducer(newLocalActor(&wg, messageCount)).
		WithMailbox(mailbox.Bounded(1000000))

	pid := actor.Spawn(props)
	fmt.Println(pid)

	remotePid := actor.NewPID("127.0.0.1:8080", "remote")
	wg.Add(1)

	start := time.Now()
	fmt.Println("Starting to send")

	message := &zmqremote.MsgData{MsgType: 1, Data: []byte("123")}
	for i := 0; i < messageCount; i++ {
		remotePid.Request(message, pid)
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Elapsed %s", elapsed)

	x := int(float32(messageCount*2) / (float32(elapsed) / float32(time.Second)))
	fmt.Printf("Msg per sec %v", x)
}
