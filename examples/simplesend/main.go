package main

import (
	"github.com/tomek7667/goimail/icloud"
)

func main() {
	c, err := icloud.New(".....@icloud.com", "....@custom-domain.com", "xxxx-xxxx-xxxx-xxxx")
	if err != nil {
		panic(err)
	}
	err = c.SendMail("This is test of the client", "<h1>hiii</h1>", nil, "receiver@example.com")
	if err != nil {
		panic(err)
	}
}
