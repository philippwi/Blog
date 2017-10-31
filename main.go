//Matrikelnummern: 3229403, 9964427

package main

import (
	"Blog/server"
	"flag"
	"fmt"
	"Blog/config"
	"strconv"
)

func main() {
	sessionExp := flag.Int("exp", config.DefaultCookieAge, "Cookie expiration (minutes)")
	port := flag.String("port", config.DefaultPort, "Server port")

	flag.Parse()

	fmt.Println(
		"Port: " + *port +
			"Session: " + strconv.Itoa(*sessionExp) + " minutes")

	server.StartServer(*sessionExp, *port)
}
