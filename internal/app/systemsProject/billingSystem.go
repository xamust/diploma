package systemsProject

import (
	"log"
	"math"
	"os"
	"server/internal/app/checkdata"
	"server/internal/app/models"
	"strconv"
	"strings"
)

type BillingSystem struct {
	check    *checkdata.CheckData
	fileName map[string]string
}

const dBilling = "billing.data"

func NewBillingSystem(fileName map[string]string) *BillingSystem {
	return &BillingSystem{
		check:    &checkdata.CheckData{},
		fileName: fileName,
	}
}

func (b *BillingSystem) readBilling() (*models.BillingData, error) {

	data, err := os.ReadFile(b.fileName[dBilling])
	if err != nil {
		return nil, err
	}
	byteMask := strings.Split(strings.Split(string(data), "\n")[0], "")
	result, err := b.calcDataBilling(byteMask)
	if err != nil {
		return nil, err
	}
	return b.check.CheckBillingData(result), nil
}

func (b *BillingSystem) calcDataBilling(input []string) (uint8, error) {
	intRes := 0
	for i := len(input) - 1; i >= 0; i-- {
		tempInt, err := strconv.Atoi(input[i])
		if err != nil {
			log.Fatalln(err.Error())
			return 0, err
		}
		intRes += tempInt * int(math.Pow(2, float64(len(input)-i-1)))
	}
	return uint8(intRes), nil
}

func (b *BillingSystem) GetBillingData() (*models.BillingData, error) {
	type Result struct {
		Payload *models.BillingData
		Error   error
	}
	in := make(chan Result)
	go func() {
		billingData, err := b.readBilling()
		if err != nil {
			in <- Result{
				Payload: nil,
				Error:   err,
			}
		}
		in <- Result{
			Payload: billingData,
			Error:   nil,
		}
	}()
	result := <-in
	return result.Payload, result.Error
}
