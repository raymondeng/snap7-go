package main

import (
	"fmt"

	"github.com/raymondeng/snap7-go"
)

// func main() {
// 	snap7_client, err := snap7.ConnentTo("127.0.0.1", 0, 0)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(snap7_client)
// }

func main() {
	host := "192.168.2.3"
	// client, err := snap7.ConnentTo2(host, 0x0101, 0x0101, 1)
	client, err := snap7.ConnentTo(host, 1, 1, 1)

	fmt.Println(client)
	fmt.Println(err)

	resultPlcTime, _ := client.GetPlcDateTime()
	fmt.Println(resultPlcTime)

	resultDr3000, _ := client.DBRead(3000, 0, 4)
	fmt.Println(resultDr3000)

}
