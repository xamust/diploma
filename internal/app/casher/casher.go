package casher

import (
	"server/internal/app/models"
	"server/internal/app/systemsproject"
	"sync"
	"time"
)

type Casher struct {
	systems   *systemsproject.SystemsProject
	CashModel *models.ResultT
	mu        *sync.Mutex
	config    *Config
}

func NewCasher(systems *systemsproject.SystemsProject, config *Config) *Casher {
	return &Casher{
		systems:   systems,
		CashModel: &models.ResultT{},
		mu:        &sync.Mutex{},
		config:    config,
	}
}

func (c *Casher) setCash() {
	data, err := c.systems.GetResultData()
	if err != nil {
		c.CashModel.Error = err.Error()
		c.CashModel.Status = false
		c.CashModel.Data = models.ResultSetT{}
	} else {
		c.CashModel.Error = ""
		c.CashModel.Status = true
		c.CashModel.Data = *data
	}
}

func (c *Casher) ResultSet() {
	c.setCash()
	go func() {
		for range time.Tick(time.Second * time.Duration(c.config.StorageTimeout)) {
			c.mu.Lock()
			c.setCash()
			c.mu.Unlock()
		}
	}()
}

func (c *Casher) ToHandler() *models.ResultT {
	return c.CashModel
}
