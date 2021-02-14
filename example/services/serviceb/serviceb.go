
package serviceb

import (
	"fmt"

	"github.com/Daironode/aingle-event/actor"
	message "github.com/Daironode/aingle-event/example/services/messages"
)

type ServiceB struct {
}

func (this *ServiceB) Receive(context actor.Context) {
	switch msg := context.Message().(type) {

	case *message.ServiceBRequest:
		fmt.Println("Receive ServiceBRequest:", msg.Message)
		context.Sender().Request(&message.ServiceBResponse{"response from serviceB"}, context.Self())

	case *message.ServiceAResponse:
		fmt.Println("Receive ServiceAResonse:", msg.Message)

	default:
		//fmt.Println("unknown message")
	}
}
