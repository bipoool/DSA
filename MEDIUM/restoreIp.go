package medium

import "fmt"

var resips []string

func resipstoreIpAddresipsses(s string) []string {

	resips = make([]string, 0)
	dfsip(s, "", 0)
	return resips
}

func dfsip(s string, currString string, currNum int) {

	fmt.Println(currString)
	if currNum > 255 {
		return
	}
	if len(s) == len(currString) {
		resips = append(resips, currString)
		return
	}
	dfsip(s[1:], currString+string(s[0]), currNum*10+int(s[0]-'0'))
}
