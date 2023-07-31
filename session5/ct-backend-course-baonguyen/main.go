// You can edit this code!
// Click here and start typing.
package main

import (
	"ct-backend-course-baonguyen/config"
	"ct-backend-course-baonguyen/internal/controller"
	mongostore "ct-backend-course-baonguyen/internal/storage/mongo"
	"ct-backend-course-baonguyen/internal/usecase"
	auth "ct-backend-course-baonguyen/pkg/auth"
	"ct-backend-course-baonguyen/pkg/bucket"
	"ct-backend-course-baonguyen/pkg/validator"
	_ "github.com/labstack/echo-jwt/v4"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	conf := config.Config{
		Port:             "8090",
		MongoURI:         "mongodb+srv://thaovo2962002:KEOF@cluster0.qrmk6cu.mongodb.net/?retryWrites=true&w=majority",
		MongoDB:          "session5",
		MongoCollImage:   "images",
		MongoCollUser:    "user",
		// GoogleCredFile:   "<Input your self>",
		// GoogleBucketName: "<Input your self>",
	}

	userStore := mongostore.NewUserCollection(conf.MongoURI, conf.MongoDB, conf.MongoCollUser)
	imageStore := mongostore.NewImageCollection(conf.MongoURI, conf.MongoDB, conf.MongoCollImage)

	//imgBucket := bucket.MustNewGoogleStorageClient(context.TODO(), conf.GoogleBucketName, conf.GoogleCredFile)
	imgBucket := bucket.NewFake()

	uc := usecase.NewUseCase(imageStore, userStore, imgBucket)
	hdl := controller.NewHandler(uc)
	
	srv := newServer(hdl)
	if err := srv.Start(":8090"); err != nil {
		log.Error(err)
	}
}

func newServer(hdl *controller.Handler) *echo.Echo {
	e := echo.New()
	e.Validator = validator.NewCustomValidator()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	public := e.Group("/api/public")
	private := e.Group("/api/private")
	private.Use(auth.AuthMiddleware(), auth.ExtractUserNameFn)

	public.POST("/register", hdl.Register)
	public.POST("/login", hdl.Login)

	private.GET("/self", hdl.Self)
	private.POST("/upload", hdl.UploadImage)

	return e
}