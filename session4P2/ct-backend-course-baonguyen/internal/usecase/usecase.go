package usecase

import (
	"context"
	"ct-backend-course-baonguyen/internal/entity"
	"ct-backend-course-baonguyen/pkg/auth"
	"errors"
	"fmt"
	"io"
	_ "net/http"
	"time"
)

type UserStore interface {
	Save(info entity.UserInfo) error
	Get(username string) (entity.UserInfo, error)
}

type ImageBucket interface {
	SaveImage(ctx context.Context, name string, r io.Reader) (string, error)
}

func NewUseCase(userStore UserStore, imageBucket ImageBucket) *ucImplement {
	return &ucImplement{
		userStore: userStore,
		imgBucket: imageBucket,
	}
}

type ucImplement struct {
	userStore UserStore
	imgBucket ImageBucket
}

func (uc *ucImplement) Register(ctx context.Context, req *entity.RegisterRequest) (*entity.RegisterResponse, error) {
	if err := uc.userStore.Save(entity.UserInfo{
		Username: req.Username,
		Password: req.Password,
		FullName: req.FullName,
		Address:  req.Address,
	}); err != nil {
		return nil, err
	}

	return &entity.RegisterResponse{UserId: req.Username}, nil
}

func (uc *ucImplement) Login(ctx context.Context, req *entity.LoginRequest) (*entity.LoginResponse, error) {
	// panic("TODO implement me")
	user, err := uc.userStore.Get(req.Username)
	if err != nil {
		return nil, err
	}
	if user.Password != req.Password {
		return nil, ErrPasswordMisMatch
	}
	token, err := auth.GenerateToken(user.Username, 24*time.Minute)
	if err != nil {
        return nil, err		
	}

	resp := entity.LoginResponse{Token: token}
	return &resp, nil
}

func (uc *ucImplement) Self(ctx context.Context, req *entity.SelfRequest) (*entity.SelfResponse, error) {
	user, err := uc.userStore.Get(req.Username)
	if err != nil {
		fmt.Print("error self uc")
		return nil , err
	}
	selfResp := entity.SelfResponse{
		Username: user.Username,
		Password: user.Password,
		FullName: user.FullName,
		Address:  user.Address,
	}
	return &selfResp, nil
}

func (uc *ucImplement) UploadImage(ctx context.Context, req *entity.UploadImageRequest) (*entity.UploadImageResponse, error) {
	var imagePath = req.Image
	imageName, err := uc.imgBucket.SaveImage(ctx, imagePath.Name(), &imagePath)
	if err != nil {
		return nil, err
	}
	return &entity.UploadImageResponse{ImageUrl: imageName}, nil
}

var ErrPasswordMisMatch = errors.New("Password mismatch")