package invoice_items

import (
	"fmt"
	"strconv"
	"syreclabs.com/go/faker"
)

type Project struct {
	Name  string
	price float64
}

type Invoice struct {
	Project       *Project
	TotalValue    float64
	TotalDiscount float64
	Discount      float64
}

type InvoiceItems struct {
	Invoice    *Invoice
	taskNumber int64
	hours      float64
	value      float64
}

func (invoiceItem InvoiceItems) String() string {
	return fmt.Sprintf("%d;%s;%.2f;%.2f",
		invoiceItem.taskNumber, invoiceItem.Invoice.Project.Name,
		invoiceItem.hours, invoiceItem.value,
	)
}

func PanicOnError(err error, message string) {
	if err != nil {
		panic(fmt.Sprintf("%s: %s", message, err.Error()))
	}
}

func NewProject() *Project {
	price, err := strconv.ParseFloat(
		faker.Number().Decimal(5, 2),
		64,
	)
	PanicOnError(err, "while creating project")

	return &Project{
		Name:  faker.Company().Name(),
		price: price,
	}
}

func (project *Project) NewInvoice() *Invoice {
	discount, err := strconv.ParseFloat(faker.Number().Between(0, 20), 64)
	PanicOnError(err, "while creating invoice")

	return &Invoice{
		Project:  project,
		Discount: discount / 100.0,
	}
}

func (invoice *Invoice) NewInvoiceItem() *InvoiceItems {
	hours, err := strconv.ParseFloat(faker.Number().Decimal(3, 2), 64)
	PanicOnError(err, "while creating invoice item")

	value := hours * invoice.Project.price
	invoice.TotalValue += value * (1 - invoice.Discount)
	invoice.TotalDiscount += value * invoice.Discount

	return &InvoiceItems{
		Invoice:    invoice,
		hours:      hours,
		value:      value,
		taskNumber: faker.Number().NumberInt64(4),
	}
}

func (invoice *Invoice) CollectInvoiceItemsSync(invoiceCount int) []*InvoiceItems {
	var items []*InvoiceItems
	for i := 0; i < invoiceCount; i++ {
		items = append(items, invoice.NewInvoiceItem())
	}
	return items
}

func (invoice *Invoice) CollectInvoiceItemsAsync(invoiceCount int, chanel chan *InvoiceItems) {
	for i := 0; i < invoiceCount; i++ {
		item := invoice.NewInvoiceItem()
		chanel <- item
	}
}
