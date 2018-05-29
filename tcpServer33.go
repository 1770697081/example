package main

import( 
    "net"
	"fmt" 


)


func HandleConn(conn net.Conn) {
    defer conn.Close()

    for {
        // read from the connection
        // ... ...
        // write to the connection
        //... ...
        fmt.Println("-------server get a connect----------")
        break
    }
}

func main() {
    listen, err := net.Listen("tcp", ":8888")
    if err != nil {
        fmt.Println("listen error: ", err)
        return
    }

    for {
        conn, err := listen.Accept()
        if err != nil {
            fmt.Println("accept error: ", err)
            break
        }

        // start a new goroutine to handle the new connection
        go HandleConn(conn)
    }
}