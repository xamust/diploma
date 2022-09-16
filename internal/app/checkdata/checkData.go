package checkdata

import (
	"fmt"
	"server/internal/app/models"
	"strconv"
	"strings"
)

// CheckData проверка входных данных
type CheckData struct {
}

func (c *CheckData) checkBandwidth(input string) error {
	bandwidth, err := strconv.Atoi(input)
	if err != nil {
		return err
	}
	if bandwidth < 0 || bandwidth > 100 {
		return fmt.Errorf("величина пропускной способности канала %d, не может быть меньше 0 или больше 100", bandwidth)
	}
	return nil
}

func (c *CheckData) checkResponseTime(input string) error {
	if _, err := strconv.Atoi(input); err != nil {
		return err
	}
	return nil
}

func (c *CheckData) checkConnectionStability(input string) (float32, error) {
	float, err := strconv.ParseFloat(input, 32)
	return float32(float), err
}

func (c *CheckData) checkTTFB(input string) (int, error) {
	return strconv.Atoi(input)
}

func (c *CheckData) checkVoicePurity(input string) (int, error) {
	return strconv.Atoi(input)
}

func (c *CheckData) checkMedianOfCallsTime(input string) (int, error) {
	return strconv.Atoi(input)
}

func (c *CheckData) checkDeliveryTime(input string) (int, error) {
	return strconv.Atoi(input)
}

func (c *CheckData) checkData(input []string, lenData int) error {

	//check len...
	if len(input) != lenData {
		return fmt.Errorf("длинна sms.data не соответсвует установленному значению %d", lenData)
	}

	//check Country...
	if err := models.SearchCode(input[0]); err != nil {
		return err
	}

	//check Bandwidth...
	if err := c.checkBandwidth(input[1]); err != nil {
		return err
	}

	//check ResponseTime...
	if err := c.checkResponseTime(input[2]); err != nil {
		return err
	}

	//check Provider...
	if err := models.SearchProvider(input[3]); err != nil {
		return err
	}
	return nil
}

// CheckDataSMS проверка данных SMS
func (c *CheckData) CheckDataSMS(input []string, lenData int) error {
	return c.checkData(input, lenData)
}

// CheckDataMMS проверка данных MMS
func (c *CheckData) CheckDataMMS(input *models.MMSData, lenData int) error {
	//check struct...
	if input.Provider == "" || input.Country == "" || input.Bandwidth == "" || input.ResponseTime == "" {
		return fmt.Errorf("некорректные поля структуры %v", input)
	}
	return c.checkData([]string{input.Country, input.Bandwidth, input.ResponseTime, input.Provider}, lenData)
}

// CheckVoiceCall проверка данных Voice
func (c *CheckData) CheckVoiceCall(input []string, lenData int) (*models.VoiceCallData, error) {

	//check len...
	if len(input) != lenData {
		return nil, fmt.Errorf("длинна voice.data не соответсвует установленному значению %d", lenData)
	}

	//check Country string...
	if err := models.SearchCode(input[0]); err != nil {
		return nil, err
	}

	//check Bandwidth string...
	if err := c.checkBandwidth(input[1]); err != nil {
		return nil, err
	}

	//check ResponseTime string...
	if err := c.checkResponseTime(input[2]); err != nil {
		return nil, err
	}

	//check Provider string...
	if err := models.SearchProviderVoiceCall(input[3]); err != nil {
		return nil, err
	}

	//check ConnectionStability float32...
	dataConStab, err := c.checkConnectionStability(input[4])
	if err != nil {
		return nil, err
	}

	//check TTFB int...
	dataTTFB, err := c.checkTTFB(input[5])
	if err != nil {
		return nil, err
	}

	//check VoicePurity int...
	dataVoicePurity, err := c.checkVoicePurity(input[6])
	if err != nil {
		return nil, err
	}

	//check MedianOfCallsTime int...
	dataMedianOfCallsTime, err := c.checkMedianOfCallsTime(input[7])
	if err != nil {
		return nil, err
	}

	return &models.VoiceCallData{
		Country:             input[0],
		Bandwidth:           input[1],
		ResponseTime:        input[2],
		Provider:            input[3],
		ConnectionStability: dataConStab,
		TTFB:                dataTTFB,
		VoicePurity:         dataVoicePurity,
		MedianOfCallsTime:   dataMedianOfCallsTime,
	}, nil
}

// CheckEmailData проверка данных Email
func (c *CheckData) CheckEmailData(input []string, lenData int) (*models.EmailData, error) {
	//check len...
	if len(input) != lenData {
		return nil, fmt.Errorf("длинна email.data не соответсвует установленному значению %d", lenData)
	}

	//check Country string...
	if err := models.SearchCode(input[0]); err != nil {
		return nil, err
	}

	//check Provider string...
	if err := models.SearchProviderEmail(input[1]); err != nil {
		return nil, err
	}

	//check DeliveryTime int...
	deliveryTime, err := c.checkDeliveryTime(input[2])
	if err != nil {
		return nil, err
	}

	return &models.EmailData{
		Country:      input[0],
		Provider:     input[1],
		DeliveryTime: deliveryTime,
	}, nil
}

// CheckBillingData проверка данных Billing
func (c *CheckData) CheckBillingData(input uint8) *models.BillingData {
	billingData := models.BillingData{
		CreateCustomer: false,
		Purchase:       false,
		Payout:         false,
		Recurring:      false,
		FraudControl:   false,
		CheckoutPage:   false,
	}
	resultString := strings.Split(fmt.Sprintf("%06b", input), "")
	if resultString[5] == "1" {
		billingData.CreateCustomer = true
	}
	if resultString[4] == "1" {
		billingData.Purchase = true
	}
	if resultString[3] == "1" {
		billingData.Payout = true
	}
	if resultString[2] == "1" {
		billingData.Recurring = true
	}
	if resultString[1] == "1" {
		billingData.FraudControl = true
	}
	if resultString[0] == "1" {
		billingData.CheckoutPage = true
	}
	return &billingData
}

// CheckDataSupport проверка данных Support
func (c *CheckData) CheckDataSupport(input *models.SupportData) error {
	//check struct...
	if input.Topic == "" || string(rune(input.ActiveTickets)) == "" {
		return fmt.Errorf("некорректные поля структуры %v", input)
	}
	return nil
}

// CheckDataIncident проверка данных Incident
func (c *CheckData) CheckDataIncident(input *models.IncidentData) error {
	//check struct...
	if input.Topic == "" || input.Status == "" {
		return fmt.Errorf("некорректные поля структуры %v", input)
	}
	if input.Status != "active" && input.Status != "closed" {
		return fmt.Errorf("некорректные данные поля структуры Status %v", input)
	}
	return nil
}
