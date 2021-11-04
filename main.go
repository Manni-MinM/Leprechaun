// BWOTSHEWCHb

package main

import (
	"io"
	"html/template"

	"github.com/Manni-MinM/Leprechaun/db"
	"github.com/Manni-MinM/Leprechaun/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (template *Template) Render(writer io.Writer , filename string , data interface{} , ctx echo.Context) error {
	return template.templates.ExecuteTemplate(writer , filename , data)
}

func main() {
	// create db and table
	var err error
	err = db.New()
	if err != nil {
		panic(err)
	}
	// TODO : Delete line below
	err = db.DropTable()
	err = db.CreateTable()
	if err != nil {
		panic(err)
	}
	// initialize server and templates
	server := echo.New()
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	templates := &Template {template.Must(template.ParseGlob("templates/*.html"))}
	server.Renderer = templates
	// setup server routing
	server.GET("/" , handler.HomePage)
	server.POST("/new" , handler.StoreLink)
	server.POST("/usage" , handler.ShowUsage)
	server.GET("/link/:shortLink" , handler.Redirect)
	// launch server
	server.Logger.Fatal(server.Start(":1323"))
}

