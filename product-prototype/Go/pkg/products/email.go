package products

import (
	"sort"
	"time"
)

type EmailNotification struct {
	CustomerID  string
	ProductType string
	Name        string
	Date        time.Time
}
type emaillist struct {
	slice []time.Time
}

func (el emaillist) Len() int {
	return len(el.slice)
}

func (el emaillist) Swap(i, j int) {
	el.slice[i], el.slice[j] = el.slice[j], el.slice[i]
}

func (el emaillist) Less(i, j int) bool {
	return el.slice[i].Before(el.slice[j])
}

func (c *CurrentData) EmailList() []EmailNotification {
	// TODO: rewrite this method

	emailList := make(map[time.Time][]*Product)
	emailIndex := emaillist{
		slice: make([]time.Time, 0),
	}

	for _, domains := range c.DomainNameIndex {
		for _, product := range domains {
			for _, date := range product.EmailDates() {
				if _, ok := emailList[date]; !ok {
					emailList[date] = []*Product{product}
					emailIndex.slice = append(emailIndex.slice, date)
				} else {
					emailList[date] = append(emailList[date], product)
				}
			}
		}
	}

	sort.Sort(emailIndex)

	resp := make([]EmailNotification, 0)
	for _, idx := range emailIndex.slice {
		for _, p := range emailList[idx] {
			resp = append(resp,
				EmailNotification{
					CustomerID:  p.CustomerID,
					ProductType: p.ProductType,
					Name:        p.Name,
					Date:        idx,
				},
			)
		}
	}

	return resp
}
