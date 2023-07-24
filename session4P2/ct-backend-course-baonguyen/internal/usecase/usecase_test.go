package usecase

import (
	"testing"
	"os"
	"context"
	inmemory "ct-backend-course-baonguyen/internal/storage/in-memory"
	bucket "ct-backend-course-baonguyen/pkg/bucket"
	"ct-backend-course-baonguyen/internal/entity"
)

func TestUcImplement_UploadImage(t *testing.T) {
	// TODO implement
	userStore := inmemory.NewUserStore()
	imgBucket := bucket.NewFake()
	uc := NewUseCase(userStore, imgBucket)

	file, err := os.OpenFile("imageTest.png", os.O_RDWR, 0644)
	if err != nil {
		return
	}
	imageInfo, err := uc.UploadImage(context.TODO(), &entity.UploadImageRequest{Image: *file})
	if err != nil {
		return
	}
	target := "/static/images/imageTest.png"

	if imageInfo.ImageUrl != target {
		t.Errorf("got %q, wanted %q", imageInfo.ImageUrl, target)
	}
}
