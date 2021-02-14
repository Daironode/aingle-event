
package servicea

import (
	"fmt"

	"github.com/Daironode/aingle-event/actor"
	message "github.com/Daironode/aingle-event/example/services/messages"
)

type ServiceA struct {
}

func (this *ServiceA) Receive(context actor.Context) {
	switch msg := context.Message().(type) {

	case *message.ServiceARequest:
		fmt.Println("Receive ServiceARequest:", msg.Message)
		context.Sender().Tell(&message.ServiceAResponse{"I got your message"})

	case *message.ServiceBResponse:
		fmt.Println("Receive ServiceBResponse:", msg.Message)

	case int:
		context.Sender().Tell(msg + 1)

	default:
		fmt.Printf("unknown message:%v\n", msg)
	}
}
