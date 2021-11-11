// BWOTSHEWCHb

package handler

import (
	"fmt"
	"time"
	"testing"
	"net/url"
	"net/http"

	"github.com/Manni-MinM/Leprechaun/db"
	"github.com/Manni-MinM/Leprechaun/model"
	
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func TestServer(t *testing.T) {
	var err error
	err = db.New()
	if err != nil {
		t.Error(err)
	}
	err = db.CreateTable()
	if err != nil {
		t.Error(err)
	}

	server := echo.New()
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	server.GET("/" , HomePage)
	server.POST("/new" , StoreLink)
	server.POST("/usage" , ShowUsage)
	server.GET("/link/:shortlink" , Redirect)

	go func() {
		server.Logger.Fatal(server.Start(":1323"))
	}()
	time.Sleep(time.Millisecond * 100)
}
func TestStoreLink(t *testing.T) {
	link := model.GetLink("http://aut.ac.ir")
	link.SetHash("hell")
	form := url.Values{}
	form.Add("URL" , link.URL)
	form.Add("desired_shortlink" , link.Hash)
	resp , err := http.PostForm("http://localhost:1323/new" , form)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Error("Invalid Status Code")
	}
	linkSel , err := db.SelectRecord(link.Hash)
	fmt.Println(linkSel , err)
	if err != nil || linkSel.URL != link.URL {
		t.Error("Unknown Link Error")
	}
}
func TestShowUsage(t *testing.T) {
	// TODO
}

