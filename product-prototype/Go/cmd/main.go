package main

import (
	"fmt"
	"time"

	"github.com/rwhelan/coding-challenge/product-prototype/Go/pkg/products"
)

func main() {
	data, err := products.LoadData()
	if err != nil {
		panic(err)
	}

	err = data.AddRecord(products.Product{
		CustomerID:  "Cust234",
		ProductType: "hosting",
		Name:        "plugh.org",
		Duration:    6,
		StartDate:   time.Now(),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(" -- Existing Products by Customer -- ")
	for _, p := range data.ListProductsByCustomer() {
		fmt.Printf("%s %s\t%s\t%s\t%d\n",
			p.CustomerID, p.ProductType, p.Name, p.StartDate.Format("2006-1-2"), p.Duration)
	}

	fmt.Println("\n -- Email Notifications -- ")
	for _, p := range data.EmailList() {
		fmt.Printf("%s %s\t%s\t%s\n",
			p.CustomerID, p.ProductType, p.Name, p.Date.Format("2006-1-2"))
	}

}
