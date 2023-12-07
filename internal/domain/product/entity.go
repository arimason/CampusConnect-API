package product

import "campusconnect-api/pkg/utils"

type Entity struct {
	ID    utils.Identity `json:"id"`
	Name  string         `json:"name"`
	Price float64        `json:"price"`
}
