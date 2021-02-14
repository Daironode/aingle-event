
package main

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/Daironode/aingle-event/actor"
	"github.com/Daironode/aingle-event/example/remotebenchmark/messages"
	"github.com/Daironode/aingle-event/mailbox"
	"github.com/Daironode/aingle-event/remote"
)

type localActor struct {
	count        int
	wgStop       *sync.WaitGroup
	messageCount int
}

func (state *localActor) Receive(context actor.Context) {
	switch context.Message().(type) {
	case *messages.Pong:
		state.count++
		if state.count%50000 == 0 {
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

	messageCount := 50000
	remote.Start("127.0.0.1:8081")

	props := actor.
		FromProducer(newLocalActor(&wg, messageCount)).
		WithMailbox(mailbox.Bounded(1000000))

	pid := actor.Spawn(props)

	remotePid := actor.NewPID("127.0.0.1:8080", "remote")
	if err := remotePid.
		RequestFuture(&messages.StartRemote{
			Sender: pid,
		}, 5*time.Second).
		Wait(); err != nil {
		fmt.Println("remote request failed:", err)
		return
	}

	wg.Add(1)

	start := time.Now()
	fmt.Println("Starting to send")

	bb := bytes.NewBuffer([]byte(""))
	for i := 0; i < 2000; i++ {
		bb.WriteString("1234567890")
	}
	message := &messages.Ping{Data: bb.Bytes()}
	for i := 0; i < messageCount; i++ {
		remotePid.Tell(message)
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Elapsed %s \n", elapsed)

	x := int(float32(messageCount*2) / (float32(elapsed) / float32(time.Second)))
	fmt.Printf("Msg per sec %v \n", x)
}
