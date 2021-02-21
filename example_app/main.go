package main

import (
	"fmt"
	"net/http"
	"flag"

	"github.com/labstack/echo/v4"
	"google.golang.org/appengine"
)

var flaglocalrun bool
func init() {
	flag.BoolVar(&flaglocalrun, "local", false, "run locally")
	flag.Parse()
}

func main() {
	e := createMux()
	if flaglocalrun{
		e.Logger.Fatal(e.Start(":8080"))
		}else{
		// the appengine package provides a convenient method to handle the health-check requests
		// and also run the app on the correct port. We just need to add Echo to the default handler
		http.Handle("/", e)
		appengine.Main()
	}
}

func createMux() *echo.Echo {
	// note: we don't need to provide the middleware or static handlers, that's taken care of by the platform
	// app engine has it's own "main" wrapper - we just need to hook echo into the default handler
	e := echo.New()
  	e.GET("/", hello)
	return e
}


func hello(c echo.Context) error {
	msg := fmt.Sprintf("Hello, world!")
	return c.String(http.StatusOK, msg)
}