package main 

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

/** 
* !TODO: Change the the display function for support a single connection
* !TODO: Display function will show only the time not the timezone
*/

type Server struct {
	conn net.Conn
	tz string
}

func main () {
	conn1 := connect(8080)
	conn2 := connect(8090)

	servers := []io.Reader{ conn1, conn2, }

	displayWall(os.Stdout, servers)
}


func displayWall(dst io.Writer, servers []io.Reader) {
	buf := make([]byte, 1024)
	c := 0

	for true {
		/*
		_, err := servers[1].Read(buf)
		if err != nil {
			log.Println("Read error: ", err)
		}

		v := strings.Split(string(buf), "-")
		local, timeNow := v[0], v[1] 
			

		if c == 0 {
			dst.Write([]byte(local + "\n"))
		}
		dst.Write([]byte(timeNow))
		*/

		_buf := make([]byte, 1024)
		log.Println(string(_buf))

		/** 
		* This approach needs the use of channels.
		* For this to work each function needs it's own 
		* concurrent displayToWall function.
		*/
		for i := range servers{

			_, err := servers[i].Read(buf)
			_buf = append(_buf, buf...)

			_, err = servers[i].Read(buf)

			if err != nil {
				log.Println("Read error: ", err)
			}

			v := strings.Split(string(buf), "-")
			local, timeNow := v[0], v[1] 
				
			if c == 0 {
				dst.Write([]byte(local + " "))
				if i == len(servers) - 1 {
					dst.Write([]byte("\n"))
				}
			} else {
				dst.Write([]byte(timeNow))
			}
		}
		c++
	}

	// dst.Write([]byte("\n"))
}

func connect(port int) net.Conn {
	addr := fmt.Sprintf("localhost:%s", strconv.Itoa(port))
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	// defer conn.Close()

	return conn
}
