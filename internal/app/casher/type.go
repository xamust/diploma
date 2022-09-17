package casher

import "server/internal/app/models"

type Cash interface {
	setCash()
	ResultSet()
	ToHandler() *models.ResultT
}
