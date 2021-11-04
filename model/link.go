// BWOTSHEWCHb

package model

import (
	"fmt"
	"time"
	"math/rand"
	"crypto/md5"
	"encoding/base64"
)

type Link struct {
	Hash string
	URL string
	UsedCount int
}

func getRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	str := make([]byte , length)
	rand.Read(str)
	return fmt.Sprintf("%s" , str)[:length]
}
func getHash(str string) string {
	str += getRandomString(8)
	md5 := md5.Sum([]byte(str))
	hash := base64.StdEncoding.EncodeToString(md5[:])
	return hash[:8]
}

func GetLink(URL string) Link {
	return Link {getHash(URL) , URL , 0}
}

