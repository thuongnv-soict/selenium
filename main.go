package main

import (
	"fmt"
	"strconv"
	"strings"
)
import "github.com/speps/go-hashids"

func encodeInvitedCode(salt string, phoneNumber string) string{
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = 8
	h, _ := hashids.NewWithData(hd)

	var number int
	var err error
	if strings.HasPrefix(phoneNumber, "+"){
		number, err = strconv.Atoi(phoneNumber[3:])
		if err != nil{
			panic(err)
		}
	}else{
		number, err = strconv.Atoi(phoneNumber[1:])
		if err != nil{
			panic(err)
		}
	}
	e, _ := h.Encode([]int{number})

	return e
}

func decodeInvitedCode(salt string, code string) []int{
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = 8
	h, _ := hashids.NewWithData(hd)

	d, err := h.DecodeWithError(code)
	if err != nil {
		fmt.Println(err)
	}

	return d
}


func main() {
	phoneNumber := "0367958012"
	encode := encodeInvitedCode("hit.vn", phoneNumber)
	fmt.Println(encode)

	decode := decodeInvitedCode("hit.vn", encode)
	fmt.Println(decode)
}
