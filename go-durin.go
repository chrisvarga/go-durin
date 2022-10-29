package durin

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
)

func Raw(message string) (string, error) {
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
	if strings.HasPrefix(data, "(error)") {
		return "", errors.New(data)
	}
	return data, nil
}

func Set(key string, value string) (string, error) {
	conn, err := net.Dial("tcp", "localhost:8045")
	if err != nil {
		return "", errors.New("(error) failed to connect to durin")
	}
	defer conn.Close()
	fmt.Fprintf(conn, "set %s %s\n", key, value)
	reader := bufio.NewReader(conn)
	data, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("(error) connection lost")
	}
	data = strings.TrimSuffix(data, "\n")
	if strings.HasPrefix(data, "(error)") {
		return "", errors.New(data)
	}
	return data, nil
}

func Get(key string) (string, error) {
	conn, err := net.Dial("tcp", "localhost:8045")
	if err != nil {
		return "", errors.New("(error) failed to connect to durin")
	}
	defer conn.Close()
	fmt.Fprintf(conn, "get %s\n", key)
	reader := bufio.NewReader(conn)
	data, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("(error) connection lost")
	}
	data = strings.TrimSuffix(data, "\n")
	if strings.HasPrefix(data, "(error)") {
		return "", errors.New(data)
	}
	return data, nil
}

func Del(key string) (string, error) {
	conn, err := net.Dial("tcp", "localhost:8045")
	if err != nil {
		return "", errors.New("(error) failed to connect to durin")
	}
	defer conn.Close()
	fmt.Fprintf(conn, "del %s\n", key)
	reader := bufio.NewReader(conn)
	data, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("(error) connection lost")
	}
	data = strings.TrimSuffix(data, "\n")
	if strings.HasPrefix(data, "(error)") {
		return "", errors.New(data)
	}
	return data, nil
}

func Keys() (string, error) {
	conn, err := net.Dial("tcp", "localhost:8045")
	if err != nil {
		return "", errors.New("(error) failed to connect to durin")
	}
	defer conn.Close()
	fmt.Fprintf(conn, "keys\n")
	reader := bufio.NewReader(conn)
	data, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("(error) connection lost")
	}
	data = strings.TrimSuffix(data, "\n")
	if strings.HasPrefix(data, "(error)") {
		return "", errors.New(data)
	}
	return data, nil
}

func KeysPrefix(prefix string) (string, error) {
	conn, err := net.Dial("tcp", "localhost:8045")
	if err != nil {
		return "", errors.New("(error) failed to connect to durin")
	}
	defer conn.Close()
	fmt.Fprintf(conn, "keys %s\n", prefix)
	reader := bufio.NewReader(conn)
	data, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("(error) connection lost")
	}
	data = strings.TrimSuffix(data, "\n")
	if strings.HasPrefix(data, "(error)") {
		return "", errors.New(data)
	}
	return data, nil
}

func Json(prefix string) (string, error) {
	conn, err := net.Dial("tcp", "localhost:8045")
	if err != nil {
		return "", errors.New("(error) failed to connect to durin")
	}
	defer conn.Close()
	fmt.Fprintf(conn, "json %s\n", prefix)
	reader := bufio.NewReader(conn)
	data, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("(error) connection lost")
	}
	data = strings.TrimSuffix(data, "\n")
	if strings.HasPrefix(data, "(error)") {
		return "", errors.New(data)
	}
	return data, nil
}
