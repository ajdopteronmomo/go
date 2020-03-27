package main

import (
	"log"

	"github.com/zeromq/goczmq"
)

func main() {
	log.Println("dealer created and connected")

	router, err := goczmq.NewRouter("tcp://*.5555")
	if err != nil {
		log.Fatal(err)
	}
	defer router.Destroy()

	// log.Println("route createed and bound")

	// dealer, err := goczmq.NewDealer("tcp://127.0.0.1:5555")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer dealer.Destroy()
	// log.Println("dealer created and connected")

	// err = dealer.SendFrame([]byte("hello"), goczmq.FlagNone)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("dealer send hello")

	// request, err := router.RecvMessage()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("router received '%s' from '%v'", request[1], request[0])

	// err = router.SendFrame(request[0], goczmq.FlagNone)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("route sent 'world'")

	// err = router.SendFrame([]byte("world"), goczmq.FlagNone)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// reply, err := dealer.RecvMessage()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("dealer received '%s'", string(reply[0]))
}
