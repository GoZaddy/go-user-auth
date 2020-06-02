package database

import (
	"context"
	"errors"
	"net/http"

	"github.com/gozaddy/AuthProject/models"
	"go.mongodb.org/mongo-driver/bson"
)

//StoreUser in DB
func StoreUser(user models.User) error {
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}

//GetUser from DB
func GetUser(email string) (models.User, error) {
	var user models.User
	filter := bson.M{
		"email": email,
	}
	collection.FindOne(context.TODO(), filter).Decode(&user)
	if user.Email != "" {
		return user, nil
	}
	return models.User{}, errors.New("user does not exist")
}

//AlreadyLoggedIn check if user is logged in
func AlreadyLoggedIn(w http.ResponseWriter, r *http.Request) (bool, string) {
	c, err := r.Cookie("session-id")
	if err != nil {
		return false, ""
	}
	response, err := Cache.Do("GET", c.Value)

	emailInBytes, _ := response.([]byte)
	if err != nil || emailInBytes == nil {
		return false, ""
	}
	return true, string(emailInBytes)
}
