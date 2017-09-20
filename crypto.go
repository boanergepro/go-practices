package main 

import (

	"fmt"
	"crypto/md5"
    "encoding/hex"

)

func main (){
	fmt.Println(GetMD5Hash("marcianita"))
}

func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}