package products

import (
	"fmt"
	"strings"
	"time"
)

type Product struct {
	CustomerID  string    `json:"CustomerID"`
	Name        string    `json:"Name"`
	ProductType string    `json:"ProductType"`
	StartDate   time.Time `json:"StartDate"`
	Duration    int       `json:"Duration"`

	DomainRegistration *Product `json:"DomainRegistration,omitempty"`
}

func (p *Product) Expires() time.Time {
	return p.StartDate.AddDate(0, p.Duration, 0)
}

func (p *Product) IsActive() bool {
	return p.Expires().After(time.Now())
}

func (p *Product) IsValid() (bool, error) {
	switch p.ProductType {
	case "domain":
		return validateDomainProduct(p)
	case "hosting":
		return validateHostingProduct(p)
	case "email":
		return validateEmailProduct(p)
	case "pdomain":
		return validateProtectedDoaminProduct(p)
	case "edomain":
		return validateEduDomainProduct(p)
	default:
		return false, fmt.Errorf("Unknown product type: %s", p.ProductType)
	}
}

func (p *Product) EmailDates() []time.Time {
	var resp []time.Time

	switch p.ProductType {
	case "hosting":
		resp = []time.Time{
			p.StartDate.AddDate(0, 0, 1),  // One day after activation
			p.Expires().AddDate(0, 0, -3), // Three days before expiration
		}

	case "email":
		resp = []time.Time{
			p.Expires().AddDate(0, 0, -1), // One day before expiration
		}

	case "pdomain", "edomain", "domain":
		resp = []time.Time{
			p.Expires().AddDate(0, 0, -2), // Two days before expiration
		}

	}

	return resp
}

// TODO: Validators

func validateDomainProduct(p *Product) (bool, error) {
	if !strings.HasSuffix(p.Name, "org") && !strings.HasSuffix(p.Name, "com") {
		return false, fmt.Errorf("invalid tld for domain")
	}

	return true, nil
}

func validateHostingProduct(p *Product) (bool, error) {
	return true, nil
}

func validateEmailProduct(p *Product) (bool, error) {
	return true, nil
}

func validateProtectedDoaminProduct(p *Product) (bool, error) {
	return true, nil
}

func validateEduDomainProduct(p *Product) (bool, error) {
	if !strings.HasSuffix(p.Name, "edu") {
		return false, fmt.Errorf("invalid tld for domain")
	}

	return true, nil
}
