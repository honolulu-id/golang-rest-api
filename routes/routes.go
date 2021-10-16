package routes

import (
	"crypto/subtle"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang-rest-api/controllers"
	"net/http"
)

func Init(port string) *echo.Echo {
	e := echo.New()
	// basic auth token api
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("GolangSampling")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("1234")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.GET("/pariwisata", controllers.pariwisata)

	if port == "" {
		port = "80"
	}

	// address := fmt.Sprintf("%s:%s", "0.0.0.0", port)
	address := fmt.Sprintf("%s:%s", "localhost", port)
	fmt.Println(address)
	e.Start(address)

	return e
}
