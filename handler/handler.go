// BWOTSHEWCHb

package handler

import (
	"net/http"

	"github.com/Manni-MinM/Leprechaun/db"
	"github.com/Manni-MinM/Leprechaun/util"
	"github.com/Manni-MinM/Leprechaun/model"

	"github.com/labstack/echo/v4"
	_"github.com/labstack/echo/v4/middleware"
)

func HomePage(ctx echo.Context) error {
	return ctx.Render(http.StatusOK , "index.html" , nil)
}
func StoreLink(ctx echo.Context) error {
	URL := ctx.FormValue("URL")
	URL = util.ToAbsURL(URL)
	newLink := model.GetLink(URL)
	go func() {
		err := db.InsertRecord(newLink)
		if err != nil {
			panic(err)
		}
	}()
	return ctx.String(http.StatusOK , util.StoreLinkMessage(ctx , newLink.Hash))
}
func ShowUsage(ctx echo.Context) error {
	shortLink := ctx.FormValue("shortLink")
	hash := util.StripURL(ctx , shortLink)
	link , err := db.SelectRecord(hash)
	if err != nil {
		return ctx.String(http.StatusBadRequest , util.UnknownURLMessage())
	} else {
		return ctx.String(http.StatusOK , util.ShowUsageMessage(link.UsedCount))
	}
}
func Redirect(ctx echo.Context) error {
	shortLink := ctx.Param("shortLink")
	link , err := db.SelectRecord(shortLink)
	if err != nil {
		return ctx.String(http.StatusBadRequest , util.UnknownURLMessage())
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

