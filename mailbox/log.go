

package mailbox

import (
	"github.com/Daironode/aingle-event/log"
)

var (
	plog = log.New(log.DebugLevel, "[MAILBOX]")
)

// SetLogLevel sets the log level for the logger.
//
// SetLogLevel is safe to call concurrently
func SetLogLevel(level log.Level) {
	plog.SetLevel(level)
}
