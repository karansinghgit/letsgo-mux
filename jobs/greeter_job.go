package jobs

import (
	letslog "github.com/letsgo-framework/letsgo-mux/log"
)

func Greet() {
	letslog.Debug("Hello Jobs")
}
