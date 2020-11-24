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
	client, err := snap7.ConnentTo2(host, 0x0101, 0x0101, 0)

	fmt.Println(client)
	fmt.Println(err)

	result, _ := client.GetPlcDateTime()
	fmt.Println(result)

}
