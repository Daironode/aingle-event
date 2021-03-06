
package router

import "github.com/Daironode/aingle-event/actor"

type AddRoutee struct {
	PID *actor.PID
}

type RemoveRoutee struct {
	PID *actor.PID
}

type AdjustPoolSize struct {
	Change int32
}
type GetRoutees struct {
}

type Routees struct {
	PIDs []*actor.PID
}
