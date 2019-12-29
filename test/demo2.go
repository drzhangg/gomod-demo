package main

import (
	"encoding/hex"
	"fmt"
	"github.com/axgle/mahonia"
)

func main() {

	str := mahonia.NewEncoder("utf8").ConvertString("周口")
	fmt.Println(str)


	str1 := mahonia.NewDecoder("gbk").ConvertString("周口")
	_,result,_ := mahonia.NewDecoder("utf-8").Translate([]byte(str1),true)
	fmt.Println(string(result))



	str2 := "乱码的字符串变量"
	str2  = ConvertToString(str2, "gbk", "utf-8")
	fmt.Println(str2)


	fmt.Println(hex.EncodeToString([]byte("周口")))


}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

