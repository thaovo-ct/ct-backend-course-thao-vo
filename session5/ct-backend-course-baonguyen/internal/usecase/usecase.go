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
	"go.mongodb.org/mongo-driver/bson/primitive"

)

type UserStore interface {
	Save(info entity.UserInfo) error
	Get(username string) (entity.UserInfo, error)
}

type ImageStore interface {
	Save(info entity.ImageInfo) error
}

type ImageBucket interface {
	SaveImage(ctx context.Context, name string, r io.Reader) (string, error)
}

func NewUseCase(imageStore ImageStore, userStore UserStore, imageBucket ImageBucket) *ucImplement {
	return &ucImplement{
		userStore:  userStore,
		imageStore: imageStore,
		imgBucket:  imageBucket,
	}
}

type ucImplement struct {
	userStore  UserStore
	imageStore ImageStore
	imgBucket  ImageBucket
}

func (uc *ucImplement) Register(ctx context.Context, req *entity.RegisterRequest) (*entity.RegisterResponse, error) {
	if err := uc.userStore.Save(entity.UserInfo{
		Id:    primitive.NewObjectID(),
        CreatedAt: time.Now().UTC(),
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
		return nil, err
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
	imageName, err := uc.imgBucket.SaveImage(ctx, req.FileName, req.Content)
	if err != nil {
		return nil, err
	}
	user, err := uc.userStore.Get(req.Username)
	if err != nil {
		fmt.Print("error self uc")
		return nil, err
	}
	// TODO: save image info to mongoDB image collection
	imgInfo := entity.ImageInfo{
		UserName: user.Username,
		ImagePath: imageName,
		FileName: req.FileName,
	}
	err = uc.imageStore.Save(imgInfo)
	if err != nil {
		return nil, err
	}
	return &entity.UploadImageResponse{ImageUrl: imageName}, nil
}

var ErrPasswordMisMatch = errors.New("Password mismatch")