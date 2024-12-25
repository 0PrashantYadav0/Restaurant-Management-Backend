package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID            primitive.ObjectID `bson:"_id"`
	Order_id      string             `json:"order_id"`
	Order_item_id string             `json:"order_item_id"`
	Food_id       *string            `json:"food_id" validate:"required"`
	Quantity      *string            `json:"quantity" validate:"required,eq=S||eq=M||eq=L"`
	Unit_price    *float64           `json:"total" validate:"required,min=0"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
}
