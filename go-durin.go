package durin

import (
	"bufio"
	"fmt"
	"net"
)

func Durin(message string) string {
	conn, err := net.Dial("tcp", "localhost:8045")
	if err != nil {
		return "(error) failed to connect to durin"
	}
	defer conn.Close()
	fmt.Fprintf(conn, "%s\n", message)
	reader := bufio.NewReader(conn)
	data, err := reader.ReadString('\n')
	if err != nil {
		return "(error) connection lost"
	}
	return data
}
