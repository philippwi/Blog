//Matrikelnummern: 3229403, 9964427

package main

import (
	"Blog/server"
	"flag"
	"fmt"
	"Blog/config"
	"strconv"
)

//Programmstart
//startet Server mit Konfigurationsmöglichkeit über Flags
func main() {
	sessionExp := flag.Int("exp", config.DefaultCookieAge, "Cookie expiration (minutes)")
	port := flag.String("port", config.DefaultPort, "Server port")

	flag.Parse()

	fmt.Println(
		"Port: " + *port +
			"\nSession: " + strconv.Itoa(*sessionExp) + " minutes" +
			"\nServer running: https://localhost:" + *port)

	fmt.Println()

	server.StartServer(*sessionExp, *port)
}
