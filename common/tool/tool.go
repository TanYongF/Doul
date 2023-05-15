package tool

import (
	"crypto/md5"
	"fmt"
	"github.com/rs/xid"
)

var (
	defaultSecrete = "douyinxiangmu" //salt for MD5 encryption
)

func Md5(pwd string) string {
	return Md5WithSalt(pwd, defaultSecrete)
}

func Md5WithSalt(pwd string, salt string) string {
	str := fmt.Sprintf("%c%c%c%s%c%c%c", salt[9], salt[0], salt[4],
		pwd, salt[3], salt[12], salt[6]) //combine salt and password.
	data := []byte(str)               //convert string to byte slice.
	hash := md5.Sum(data)             //MD5 encryption
	md5str := fmt.Sprintf("%x", hash) //convert []byte to hexadecimal.
	return md5str
}

// TokenGenerator 生成用户token
func TokenGenerator() string {
	return xid.New().String()
}
