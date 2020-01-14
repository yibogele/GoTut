package main

import (
	"bufio"
	"fmt"
	"net"
)

/**
 * author: will fan
 * created: 2020/1/12 21:03
 * description:
 */
func main() {
	conn, _ := net.Dial("tcp", "baidu.com:80")
	_, _ = fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}