// BWOTSHEWCHb

package model

import (
	"crypto/md5"
	"encoding/base64"
)

type Link struct {
	Hash string
	URL string
	UsedCount int
}

func getHash(str string) string {
	md5 := md5.Sum([]byte(str))
	hash := base64.StdEncoding.EncodeToString(md5[:])
	return hash[:8]
}

func GetLink(URL string) Link {
	return Link {getHash(URL) , URL , 0}
}

