package users

import (
	"log"

	"feldrise.com/inventory-exercice/graph/model"
	"feldrise.com/inventory-exercice/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	Email        string             `bson:"email"`
	PasswordHash string             `bson:"password_hash"`
}

func (user *User) ToModel() *model.User {
	return &model.User{
		ID:    user.ID.Hex(),
		Email: user.Email,
	}
}

func Create(input model.NewUser) *User {
	hashedPassword, err := HashPassword(input.Password)

	if err != nil {
		log.Fatal(err)
	}

	databaseUser := User{
		ID:           primitive.NewObjectID(),
		Email:        input.Email,
		PasswordHash: hashedPassword,
	}

	_, err = database.CollectionUsers.InsertOne(database.MongoContext, databaseUser)

	if err != nil {
		log.Fatal(err)
	}

	return &databaseUser
}

func GetUserById(id string) (*User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.D{
		primitive.E{
			Key:   "_id",
			Value: objectId,
		},
	}

	users, err := GetFiltered(filter)

	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return &users[0], nil
}

func GetUserByEmail(email string) (*User, error) {
	filter := bson.D{
		primitive.E{
			Key:   "email",
			Value: email,
		},
	}

	users, err := GetFiltered(filter)

	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return &users[0], nil
}

func GetFiltered(filter interface{}) ([]User, error) {
	users := []User{}

	cursor, err := database.CollectionUsers.Find(database.MongoContext, filter)

	if err != nil {
		return users, err
	}

	for cursor.Next(database.MongoContext) {
		var user User

		err := cursor.Decode(&user)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return users, err
	}

	return users, nil
}

func Authenticate(login model.Login) bool {
	user, err := GetUserByEmail(login.Email)

	if user == nil || err != nil {
		return false
	}

	return CheckPasswordHash(login.Password, user.PasswordHash)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
