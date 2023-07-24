package entity

import (
	"os"
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
	Image os.File 
}

type UploadImageResponse struct {
	ImageUrl string `json:"imageURL"`
}

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
}

type ImageInfo struct {
	ImagePath string `json:"image"`
}