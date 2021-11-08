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

// renders home page on server
func HomePage(ctx echo.Context) error {
	return ctx.Render(http.StatusOK , "index.html" , nil)
}
// receives URL via echo context and adds its hash to the db
func StoreLink(ctx echo.Context) error {
	URL := ctx.FormValue("URL")
	URL = util.ToAbsURL(URL)
	expiryDate := ctx.FormValue("expiry_date")
	newLink := model.GetLink(URL)
	// add record to db
	go func() {
		err := db.InsertRecord(newLink , expiryDate)
		if err != nil {
			panic(err)
		}
	}()
	return ctx.String(http.StatusOK , util.StoreLinkMessage(ctx , newLink.Hash))
}
// get Used Count of short URL from db and renders it on server
func ShowUsage(ctx echo.Context) error {
	shortLink := ctx.FormValue("shortLink")
	hash := util.StripURL(ctx , shortLink)
	// get record from db
	link , err := db.SelectRecord(hash)
	if err != nil {
		return ctx.String(http.StatusBadRequest , util.UnknownURLMessage())
	} else {
		return ctx.String(http.StatusOK , util.ShowUsageMessage(link.UsedCount))
	}
}
// redirects user from short URL to original URL and increments Used Count of link by 1
func Redirect(ctx echo.Context) error {
	shortLink := ctx.Param("shortLink")
	// get record from db
	link , err := db.SelectRecord(shortLink)
	if err != nil {
		return ctx.String(http.StatusBadRequest , util.UnknownURLMessage())
	} else {
		// increment used count of link by 1 in db
		go func() {
			err := db.UpdateRecord(link)
			if err != nil {
				panic(err)
			}
		}()
		return ctx.Redirect(http.StatusTemporaryRedirect , link.URL)
	}
}

