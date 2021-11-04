// BWOTSHEWCHb

package handler

import (
	"fmt"
	"strings"
	"net/http"

	"github.com/Manni-MinM/Leprechaun/db"
	"github.com/Manni-MinM/Leprechaun/model"

	"github.com/labstack/echo/v4"
	_"github.com/labstack/echo/v4/middleware"
)

func HomePage(ctx echo.Context) error {
	return ctx.Render(http.StatusOK , "index.html" , nil)
}
func StoreLink(ctx echo.Context) error {
	URL := ctx.FormValue("URL")
	if !strings.HasPrefix(URL , "http://") && !strings.HasPrefix(URL , "https://") {
		URL = "http://" + URL
	}
	newLink := model.GetLink(URL)
	go func() {
		err := db.InsertRecord(newLink)
		if err != nil {
			panic(err)
		}
	}()
	return ctx.String(http.StatusOK , fmt.Sprintf("You Leprechaun URL : %s/link/%s" , ctx.Request().Host , newLink.Hash))
}
func Redirect(ctx echo.Context) error {
	shortLink := ctx.Param("shortLink")
	link , err := db.SelectRecord(shortLink)
	if err != nil {
		return ctx.String(http.StatusBadRequest , "Invalid or Expired Link")
	} else {
		go func() {
			err := db.UpdateRecord(link)
			if err != nil {
				panic(err)
			}
		}()
		return ctx.Redirect(http.StatusTemporaryRedirect , link.URL)
	}
}

