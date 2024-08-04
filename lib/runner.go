package lib

import (
	"flag"
	"fmt"
	"invoices_print/invoice_items"
)

func RecoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Printf("An error occurred: %v\n", r)
		fmt.Println("Application terminated gracefully.")
	} else {
		fmt.Println("Application executed successfully.")
	}
}

func DisplayOptions(parallelFlag bool, invoiceCount, projects int) {
	underlines := "___________________"
	fmt.Println("In this run:")
	fmt.Println(underlines)
	fmt.Print("in parallel?\t")
	if parallelFlag {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	fmt.Println("invoice items:\t", invoiceCount)
	fmt.Println("projects:\t", projects)
	fmt.Println(underlines)
}

func SyncRun(invoiceCount, projects int) {
	fmt.Println("taskNumber;projectName;hours;value")
	for i := 0; i < projects; i++ {
		proj := invoice_items.NewProject()
		invoice := proj.NewInvoice()
		for _, item := range invoice.CollectInvoiceItemsSync(invoiceCount) {
			fmt.Println(item)
		}
		fmt.Printf("The billed value is $ %.2f\n", invoice.TotalValue)
		fmt.Printf(
			"The discount was $ %.2f (%.2f pct)\n",
			invoice.TotalDiscount, invoice.Discount*100,
		)
	}
}

func PrintProjectInvoicesTotal(
	items map[*invoice_items.Invoice][]*invoice_items.InvoiceItems,
) {
	var total float64
	var totalDiscount float64
	var discount float64
	fmt.Println("Project Name\tTotal Value\tTotalDiscount\tDiscount")
	for invoice := range items {
		total += invoice.TotalValue
		totalDiscount += invoice.TotalDiscount
		fmt.Println(invoice.Project.Name, "\t", invoice.TotalValue, "\t", invoice.TotalDiscount, "\t", invoice.Discount)
	}
	discount = total / totalDiscount
	fmt.Println("-----\t", total, "\t", totalDiscount, "\t", discount)
}

func AsyncRun(invoiceCount, projects int) {
	ch := make(chan *invoice_items.InvoiceItems, 3)
	var items = make(map[*invoice_items.Invoice][]*invoice_items.InvoiceItems)

	receivedItems := 0
	for i := 0; i < projects; i++ {
		proj := invoice_items.NewProject()
		items[proj.NewInvoice()] = []*invoice_items.InvoiceItems{}
	}

	for invoice := range items {
		go invoice.CollectInvoiceItemsAsync(invoiceCount, ch)
	}

	for receivedItems != projects*invoiceCount {
		item, more := <-ch
		if !more {
			break
		}
		receivedItems++
		invoiceArray := items[item.Invoice]
		fmt.Println(receivedItems, "->", item)
		invoiceArray = append(invoiceArray, item)
	}
	PrintProjectInvoicesTotal(items)
}

func ParseFlags() (bool, int, int) {
	parallelFlag := flag.Bool("parallel", false, "Will it be run in parallel mode? (default: FALSE)")
	invoiceItems := flag.Int("items", 20, "The number of invoices to generate")
	projects := flag.Int("projects", 1, "The number of projects to generate")
	flag.Parse()

	return *parallelFlag, *invoiceItems, *projects
}
