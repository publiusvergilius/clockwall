package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	// Brazil/East=localhost:8080
	args := os.Args 
	
	v := getArgs(args)
	addr, tz := v[0], v[1]
	
	listener, err := net.Listen("tcp", addr)	
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Running on ", addr)

	for {
		conn, err := listener.Accept()	
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn, tz)
	}
}

func handleConn(c net.Conn, tz string){
	defer c.Close()	

	for {
		timeNow := getTimeLocation(tz)
		response := fmt.Sprintf("%s-%s", tz,timeNow.Format("15:04:05\n"))
		_, err := io.WriteString(c, response) 
		if err != nil {
			return 
		}
		 time.Sleep(1 * time.Second)
	}
}

func getTimeLocation(tz string) time.Time {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		panic(err)	
	}

	now := time.Now()

	return now.In(loc)
}

func getArgs (args []string) [2]string {
	if len(args) < 2 {
		log.Fatal("You should provide one arguments like: `Brazil/East=localhost:8080`")
	}

	if len(args) > 2 {
		log.Fatal("Too many arguments")
	}

	if !strings.ContainsRune(args[1], '=') {
		log.Fatal("Wrong format")
	}


	parts := strings.Split(args[1], "=")
	addr, tz := parts[1], parts[0]

	return [2]string{addr, tz}
}
