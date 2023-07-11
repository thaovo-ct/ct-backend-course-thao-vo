package main

import (
    "github.com/go-playground/validator"
    "github.com/labstack/echo/v4"
)

func main() {
	// http.HandleFunc("/api/public/register", register)
	// http.HandleFunc("/api/public/login", login)
	// http.HandleFunc("/api/private/self", self)

	// // http.HandleFunc("/api/public/log/register", LogWrapper(register))
	// // http.HandleFunc("/api/public/log/login", LogWrapper(login))
	// // http.HandleFunc("/api/private/log/self", LogWrapper(self))

	// http.ListenAndServe(":8090", nil)
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
    e.POST("/api/public/register", Register)
	e.POST("/api/public/login", Login)
	e.POST("/api/private/self", Self)
    e.Logger.Fatal(e.Start(":1323"))
}


/*
		TODO #2:
		- implement the logic to register a new user (username, password, full_name, address)
	  	- Validate username (not empty and unique)
	  	- Validate password (length should at least 8)
// */
/*
TODO: extra wrapper
Print some logs to console
  - Path
  - Http Status code
  - Time start, Duration
*/
// func LogWrapper(handler http.HandlerFunc) http.HandlerFunc {
// 	panic("TODO implement me")
// }

/*
	TODO #1: implement in-memory user store
	TODO #2: implement register handler
	TODO #3: implement login handler
	TODO #4: implement self handler

	Extra: implement log handler
*/