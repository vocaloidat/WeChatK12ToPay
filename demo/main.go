package main

import "fmt"

func main() {
	Wechatpay_Serial := "200"
	PlatformSerialNo1 := "100"
	PlatformSerialNo2 := "200"

	if Wechatpay_Serial != PlatformSerialNo1 && Wechatpay_Serial != PlatformSerialNo2 {
		fmt.Println("解码失败")
	} else {
		fmt.Println("解码成功")
	}
}
