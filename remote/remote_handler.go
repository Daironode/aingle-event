

package remote

import "github.com/Daironode/aingle-event/actor"

func remoteHandler(pid *actor.PID) (actor.Process, bool) {
	ref := newProcess(pid)
	return ref, true
}
