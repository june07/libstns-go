package main

import (
	"fmt"

	"github.com/STNS/libstns-go/libstns"
	"github.com/k0kubun/pp"
)

func main() {
	stns, err := libstns.NewSTNS("https://stns.lolipop.io/v1/", nil)
	if err != nil {
		panic(err)
	}

	user, err := stns.GetUserByName("pyama")
	if err != nil {
		panic(err)
	}
	pp.Println(user)

	signature, err := stns.Signature([]byte("secret message"))
	if err != nil {
		panic(err)
	}

	// it is ok
	fmt.Println(stns.VerifyWithUser("pyama", signature))

	// verify error
	fmt.Println(stns.VerifyWithUser("pyama", []byte("dummy")))
}
