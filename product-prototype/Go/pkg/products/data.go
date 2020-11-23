package products

import (
	"encoding/json"
	"fmt"
	"time"
)

type DataIndex map[string][]*Product

type CurrentData struct {
	CustomerData    DataIndex
	DomainNameIndex DataIndex
}

func (c *CurrentData) ListProductsByCustomer() []*Product {
	resp := []*Product{}
	for k := range c.CustomerData {
		resp = append(resp, c.CustomerData[k]...)
	}

	return resp
}

func (c *CurrentData) LookupDomainRegistration(domainName string) (*Product, error) {
	for _, product := range c.DomainNameIndex[domainName] {
		if product.ProductType == "domain" ||
			product.ProductType == "pdomain" ||
			product.ProductType == "edomain" {
			return product, nil
		}
	}

	return nil, fmt.Errorf("unable to find domain registation for %s", domainName)
}

func (c *CurrentData) AddRecord(newRecord Product) error {
	if newRecord.ProductType == "email" || newRecord.ProductType == "hosting" {
		domainRegistration, err := c.LookupDomainRegistration(newRecord.Name)
		// TODO: allow adding old records
		if err != nil || !domainRegistration.IsActive() {
			return fmt.Errorf("Active domain registration required")
		}

		newRecord.DomainRegistration = domainRegistration
	}

	c.loadRecord(&newRecord)
	return nil
}

func (c *CurrentData) loadRecord(p *Product) {
	// TODO: dup detection
	_, ok := c.CustomerData[p.CustomerID]
	if !ok {
		c.CustomerData[p.CustomerID] = []*Product{p}
	} else {
		c.CustomerData[p.CustomerID] = append(c.CustomerData[p.CustomerID], p)
	}

	_, ok = c.DomainNameIndex[p.Name]
	if !ok {
		c.DomainNameIndex[p.Name] = []*Product{p}
	} else {
		c.DomainNameIndex[p.Name] = append(c.DomainNameIndex[p.Name], p)
	}
}

func LoadData() (*CurrentData, error) {
	rawLoadData := make([]map[string]interface{}, 0)

	if err := json.Unmarshal([]byte(rawData), &rawLoadData); err != nil {
		return nil, fmt.Errorf("Unable to unmarshal initial data: %w\n", err)
	}

	data := &CurrentData{
		CustomerData:    make(map[string][]*Product),
		DomainNameIndex: make(map[string][]*Product),
	}

	for _, rawProductData := range rawLoadData {
		timestamp, err := time.Parse("2006-2-1", rawProductData["Created"].(string))
		if err != nil {
			return nil, fmt.Errorf("Unable to create Product struct: %w", err)
		}

		data.loadRecord(&Product{
			CustomerID:  rawProductData["CustomerID"].(string),
			Name:        rawProductData["DomainName"].(string),
			ProductType: rawProductData["ProductType"].(string),
			Duration:    int(rawProductData["Duration"].(float64)),
			StartDate:   timestamp,
		})
	}

	// TODO: un-nest
	for _, plist := range data.CustomerData {
		for _, product := range plist {
			if product.ProductType == "hosting" || product.ProductType == "email" {
				registration, err := data.LookupDomainRegistration(product.Name)
				if err != nil {
					return nil, fmt.Errorf("bad data; unable to find domain registration for %s", product.Name)
				}

				product.DomainRegistration = registration
			}
		}
	}

	return data, nil
}

const (
	rawData = `[
		{
			"CustomerID": "Cust123",
			"ProductType": "domain",
			"DomainName": "xyzzy.com",
			"Created": "2020-1-1",
			"Duration": 12
		},
		{
			"CustomerID": "Cust123",
			"ProductType": "hosting",
			"DomainName": "xyzzy.com",
			"Created": "2020-1-1",
			"Duration": 6
		},
		{
			"CustomerID": "Cust234",
			"ProductType": "domain",
			"DomainName": "plugh.org",
			"Created": "2020-2-1",
			"Duration": 24
		},
		{
			"CustomerID": "Cust123",
			"ProductType": "domain",
			"DomainName": "mydomain.com",
			"Created": "2020-3-1",
			"Duration": 12
		},
		{
			"CustomerID": "Cust123",
			"ProductType": "email",
			"DomainName": "mydomain.com",
			"Created": "2020-3-1",
			"Duration": 12
		},
		{
			"CustomerID": "Cust345",
			"ProductType": "pdomain",
			"DomainName": "protected.org",
			"Created": "2020-3-1",
			"Duration": 36
		},
		{
			"CustomerID": "Cust456",
			"ProductType": "edomain",
			"DomainName": "school.edu",
			"Created": "2020-4-1",
			"Duration": 12
		},
		{
			"CustomerID": "Cust345",
			"ProductType": "hosting",
			"DomainName": "protected.org",
			"Created": "2020-4-1",
			"Duration": 11
		}
	]`
)
