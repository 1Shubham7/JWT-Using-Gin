package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time
)"

type User struct{
	ID primitive.ObjectID	`bson:"id"`
	First_name *string	`json:"first_name" validate:"required, min=2, max=100"`
	Last_name *string `json:"last_name" validate:"required, min=2, max=100"`
	Password *string `json:"password" validate:"required, min=6`
	Email	*string `json:"email" validate:"email, required"`  //validate email means it should have an @
	Phone	*string `json:"phone" validate:"required"`
	User_type *string `json:"user_type" validate:"required, eq=ADMIN|eq=USER"`
	Refresh_token *string `json:"refresh_token"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	User_id string `json:"user_id"`
}