// BWOTSHEWCHb

package model

import (
	"fmt"
	"time"
	"math/rand"
	"crypto/md5"
	"encoding/hex"
)

// Link struct stores URL hash , original URL and number of redirections
type Link struct {
	Hash string
	URL string
	UsedCount int
}

// returns random string with specified length
func getRandomString(length int) string {
	// set random seed by time
	rand.Seed(time.Now().UnixNano())
	// create random string
	str := make([]byte , length)
	rand.Read(str)
	return fmt.Sprintf("%s" , str)[:length]
}
// returns hash value of original URL
func getHash(str string) string {
	// append random string of length 8 to URL
	str += getRandomString(8)
	// calculate md5 hash of string
	md5 := md5.Sum([]byte(str))
	// encode hash to hex format (only letters and numbers)
	hash := hex.EncodeToString(md5[:])
	return hash[:8]
}

// returns new Link instance of specified URL
func GetLink(URL string) Link {
	return Link {getHash(URL) , URL , 0}
}

func (link *Link) SetHash(hash string) {
	link.Hash = hash
	return
}

