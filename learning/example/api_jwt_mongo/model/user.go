package model

// import (
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

type User struct {
	// ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ID       string `json:"id,omitempty" bson:"_id,omitempty"` // tag golang
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Role     string `json:"role" bson:"role"`
	Password string `json:"password" bson:"password"`
}
