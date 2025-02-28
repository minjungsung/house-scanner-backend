package repositories

import (
	"context"
	"house-scanner-backend/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(email string) error
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{collection: db.Collection("users")}
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	filter := bson.M{"email": email}
	err := r.collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(user *models.User) error {
	_, err := r.collection.InsertOne(context.TODO(), user)
	return err
}

func (r *userRepository) UpdateUser(user *models.User) error {
	filter := bson.M{"email": user.Email}
	update := bson.M{"$set": user}
	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *userRepository) DeleteUser(email string) error {
	filter := bson.M{"email": email}
	_, err := r.collection.DeleteOne(context.TODO(), filter)
	return err
}
