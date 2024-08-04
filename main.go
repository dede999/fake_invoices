package main

import (
	"fmt"
	"invoices_print/lib"
	"time"
)

func main() {
	defer lib.RecoverFromPanic()

	parallelFlag, invoiceItems, projects := lib.ParseFlags()
	lib.DisplayOptions(parallelFlag, invoiceItems, projects)
	start := time.Now()
	if parallelFlag {
		lib.AsyncRun(invoiceItems, projects)
	} else {
		lib.SyncRun(invoiceItems, projects)
	}
	fmt.Println("Took:", time.Since(start))
}
