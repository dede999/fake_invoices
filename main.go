package main

import (
	"flag"
	"fmt"
	"invoices_print/lib"
	"time"
)

func main() {
	defer lib.RecoverFromPanic()

	parallelFlag := flag.Bool("parallel", false, "Will it be run in parallel mode? (default: FALSE)")
	invoiceItems := flag.Int("items", 20, "The number of invoices to generate")
	projects := flag.Int("projects", 1, "The number of projects to generate")
	flag.Parse()
	lib.DisplayOptions(parallelFlag, invoiceItems, projects)
	start := time.Now()
	if *parallelFlag {
		lib.AsyncRun(invoiceItems, projects)
	} else {
		lib.SyncRun(invoiceItems, projects)
	}
	fmt.Println("Took:", time.Since(start))
}
