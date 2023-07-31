package entity

import (
	"io"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
}

type RegisterResponse struct {
	UserId string `json:"user_id"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,gt=0"`
	Password string `json:"password" validate:"required,gt=7"`
}

type LoginResponse struct {
	Token string
}

type SelfRequest struct {
	Username string
}

type SelfResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
}

type UploadImageRequest struct {
	Username string
	FileName string
	Content  io.Reader
}

type UploadImageResponse struct {
	ImageUrl string `json:"imageURL"`
}

type UserInfo struct {
	Id primitive.ObjectID 		`bson:"_id"`
	CreatedAt time.Time         `bson:"createdAt"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	FullName string `json:"full_name" bson:"full_name"`
	Address  string `json:"address" bson:"address"`
}

type ImageInfo struct {
	UserName  string `json:"username"`
	ImagePath string `json:"image_path"`
	FileName  string `json:"file_name"`
}