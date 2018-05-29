package main

import( 
	"net"
	"fmt" 

)

//一个tcp的服务器测试程序;
func HandleConn(conn net.Conn) {
	defer conn.Close()
	
    for {
		data := make([]byte, 128)  
		conn.Read(data)
        fmt.Println("--------HandleConn read buff-------",string(data))
        
        recv := string(data[:3])
        switch recv {
        case "bye":
            fmt.Println("one")
            conn.Write([]byte("one"))
        case "hel":
            fmt.Println("two")
            conn.Write([]byte("two"))
        case "yes":
            fmt.Println("three")
            conn.Write([]byte("three"))
        default:
            fmt.Println("default")
            conn.Write([]byte("default"))
        }


        //conn.Write([]byte{'f', 'i', 'n', 'i', 's', 'h'})
        //conn.Write([]byte("finisssss"))  

        
       
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

        go HandleConn(conn)
    }
}