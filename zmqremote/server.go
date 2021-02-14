
package zmqremote

import (
	"github.com/Daironode/aingle-event/actor"
	"github.com/Daironode/aingle-event/log"
	zmq "github.com/pebbe/zmq4"
)

var (
	edpReader *endpointReader
	conn      *zmq.Socket
)

func Start(address string) {

	actor.ProcessRegistry.RegisterAddressResolver(remoteHandler)
	actor.ProcessRegistry.Address = address

	spawnActivatorActor()
	startEndpointManager()

	edpReader = &endpointReader{}

	conn, _ = zmq.NewSocket(zmq.ROUTER)
	err := conn.Bind("tcp://" + address)
	if err != nil {
		plog.Error("failed to Bind", log.Error(err))
	}
	plog.Info("Starting Proto.Actor server", log.String("address", address))
	go func() {
		edpReader.Receive(conn)
	}()
}

func Shutdonw() {
	edpReader.suspend(true)
	stopEndpointManager()
	stopActivatorActor()
	conn.Close()
}
