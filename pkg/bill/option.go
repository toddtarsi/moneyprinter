package bill

import (
	"github.com/sneakynet/moneyprinter/pkg/db"
)

// WithDatabase sets up the database reference for the bill processor.
func WithDatabase(d *db.DB) ProcessorOption { return func(p *Processor) { p.db = d } }
