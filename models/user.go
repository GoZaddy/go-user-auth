package models

//User model
type User struct {
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Password []byte `json:"password,omitempty" bson:"password,omitempty"`
}
