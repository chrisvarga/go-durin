package durin

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
)

func Durin(message string) (string, error) {
	conn, err := net.Dial("tcp", "localhost:8045")
	if err != nil {
		return "", errors.New("(error) failed to connect to durin")
	}
	defer conn.Close()
	fmt.Fprintf(conn, "%s\n", message)
	reader := bufio.NewReader(conn)
	data, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("(error) connection lost")
	}
	data = strings.TrimSuffix(data, "\n")
	return data, nil
}
