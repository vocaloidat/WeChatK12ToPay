package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	b1 := `{"out_user_id":"qq123","organization_id":"ww123"}`
	type User struct {
		Scene         string `json:"scene"`
		Web_init_data string `json:"web_init_data"`
	}
	u1 := User{
		Scene:         "WEBSESSION",
		Web_init_data: b1,
	}
	json_str, err := json.Marshal(u1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(json_str))
}
