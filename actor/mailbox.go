
package actor

import "github.com/Daironode/aingle-event/mailbox"

var (
	defaultDispatcher = mailbox.NewDefaultDispatcher(300)
)

var defaultMailboxProducer = mailbox.Unbounded()
