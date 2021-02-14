
package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Daironode/aingle-event/actor"
	msg "github.com/Daironode/aingle-event/example/services/messages"
	"github.com/Daironode/aingle-event/example/services/servicea"
	"github.com/Daironode/aingle-event/example/services/serviceb"
)

func main() {
	sva := actor.FromProducer(func() actor.Actor { return &servicea.ServiceA{} })
	svb := actor.FromProducer(func() actor.Actor { return &serviceb.ServiceB{} })

	pipA, _ := actor.SpawnNamed(sva, "serviceA")
	pipB, _ := actor.SpawnNamed(svb, "serviceB")

	pipA.Request(&msg.ServiceARequest{"TESTA"}, pipB)

	pipB.Request(&msg.ServiceBRequest{"TESTB"}, pipA)
	time.Sleep(2 * time.Second)

	f := pipA.RequestFuture(1, 50*time.Microsecond)
	result, err := f.Result()
	if err != nil {
		fmt.Println("errors:", err.Error())
	}
	fmt.Println("get sync call result :" + strconv.Itoa(result.(int)))

}
