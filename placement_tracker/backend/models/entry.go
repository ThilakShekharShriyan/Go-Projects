package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Company struct {
	ID          primitive.ObjectID `bson:"id"`
	Companyname *string            `json:"comp_name"`
	Package     *float64           `json:"package"`
	Role        *string            `json:"role"`
}
