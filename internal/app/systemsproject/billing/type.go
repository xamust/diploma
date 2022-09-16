package billing

import "server/internal/app/models"

const (
	dBilling = "billing.data"
)

type Billing interface {
	readBilling() (*models.BillingData, error)
	GetBillingData() (*models.BillingData, error)
	calcDataBilling(input []string) (uint8, error)
}
