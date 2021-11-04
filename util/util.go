// BWOTSHEWCHb

package util

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

func ToAbsURL(URL string) string {
	if !strings.HasPrefix(URL , "http://") && !strings.HasPrefix(URL , "https://") {
		URL = "http://" + URL
	}
	return URL
}
func StripURL(ctx echo.Context , URL string) string {
	URL = ToAbsURL(URL)
	key := fmt.Sprintf("http://%s/link/" , ctx.Request().Host)
	if strings.HasPrefix(URL , key) {
		URL = strings.Replace(URL , key , "" , 1)
	}
	return URL
}
func StoreLinkMessage(ctx echo.Context , hash string) string {
	 return fmt.Sprintf("Your Leprechaun URL : %s/link/%s" , ctx.Request().Host , hash)
}
func ShowUsageMessage(usedCount int) string {
	return fmt.Sprintf("This Link has Been Used %d Time(s)" , usedCount)
}
func UnknownURLMessage() string {
	return "Invalid or Expired Link"
}

