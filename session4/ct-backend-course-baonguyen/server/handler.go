package main

import (
	"net/http"
	"time"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4/middleware"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	
)

type RegisterRequest struct {
	Address string `json:"address"`
	Username string `json:"username" validate:"required,gt=0"`
	Password string `json:"password" validate:"required,gt=7"`
	Fullname string `json:"full_name"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,gt=0"`
	Password string `json:"password" validate:"required,gt=7"`
}

type CustomValidator struct {
    validator *validator.Validate
}

type LoginResponse struct {
	Token string
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}
	return nil
}

func Register(c echo.Context) error {
    var req RegisterRequest

    if err := c.Bind(&req); err != nil {
        return err
    }
	if err := c.Validate(req); err != nil {
		return err
	}

	_, exist := userStore.data[req.Username]
	if (exist) {
		return echo.NewHTTPError(http.StatusBadRequest,"user existed")
	}
	
	userInfo := UserInfo{
		Fullname: req.Fullname,
		Address: req.Address,
		Username: req.Username,
		Password: req.Password,
	}

	if err := userStore.Save(userInfo); err != nil {
		return err
	}
    return c.JSON(http.StatusCreated, "register successfully")
}

func Login(c echo.Context) error {
    var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := userStore.Get(req.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}

    if user.Password != req.Password {
        return echo.NewHTTPError(http.StatusUnauthorized, "Password not match")
    }
	
	token, err := GenerateToken(user.Username, 24*time.Minute)
	if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())		
	}

	resp := LoginResponse{Token: token}
    return c.JSON(http.StatusOK, resp)
}

func Self(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	extractUserNameFn := func(authenticationHeader string) (string, error) {

		var name string
		token, err := jwt.Parse(authenticationHeader, func(token *jwt.Token) (interface{}, error) {
			// check token signing method etc
			return []byte("ct-secret-key"), nil
		})
		if err != nil {
			return "", err
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			name = fmt.Sprint(claims["sub"])
		}

		if name == "" {
			return "", fmt.Errorf("invalid token payload")
		}
		return name, nil
	}

	username, err := extractUserNameFn(authHeader)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())		
	}

	user, _ := userStore.Get(username)
    return c.JSON(http.StatusOK, user)

}