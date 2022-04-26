package systemsProject

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"math"
	"server/internal/app/models"
	"strconv"
	"strings"
)

type BillingSystem struct {
	check    *CheckData
	logger   *logrus.Logger
	fileName map[string]string
}

func (b *BillingSystem) ReadBillingData() (*models.BillingData, error) {

	// TODO:configs filename....
	data, err := ioutil.ReadFile(b.fileName["billing.data"])
	if err != nil {
		return nil, err
	}
	//TODO:need another way to '\n'...
	byteMask := strings.Split(strings.Split(string(data), "\n")[0], "")
	result, err := b.calcDataBilling(byteMask)
	if err != nil {
		b.logger.Error(err)
		return nil, err
	}
	return b.CheckBillingData(result), nil
}

func (b *BillingSystem) CheckBillingData(input uint8) *models.BillingData {
	return b.check.CheckBillingData(input)
}

func (b *BillingSystem) calcDataBilling(input []string) (uint8, error) {
	//todo:еще глянуть
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
