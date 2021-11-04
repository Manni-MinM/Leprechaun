// BWOTSHEWCHb

package util

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

// puts http at the beginning of URL if it already isnt
func ToAbsURL(URL string) string {
	if !strings.HasPrefix(URL , "http://") && !strings.HasPrefix(URL , "https://") {
		URL = "http://" + URL
	}
	return URL
}
// returns hash value part of URL
func StripURL(ctx echo.Context , URL string) string {
	URL = ToAbsURL(URL)
	// replaces first part of URL with empty string
	key := fmt.Sprintf("http://%s/link/" , ctx.Request().Host)
	if strings.HasPrefix(URL , key) {
		URL = strings.Replace(URL , key , "" , 1)
	}
	return URL
}
// message displayed when URL is stored successfully
func StoreLinkMessage(ctx echo.Context , hash string) string {
	 return fmt.Sprintf("Your Leprechaun URL : %s/link/%s" , ctx.Request().Host , hash)
}
// message displayed when client requests Used Count of short link
func ShowUsageMessage(usedCount int) string {
	return fmt.Sprintf("This Link has Been Used %d Time(s)" , usedCount)
}
// message displayed when URL is unknown
func UnknownURLMessage() string {
	return "Invalid or Expired Link"
}

