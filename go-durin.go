package durin

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
)

func Init(host string) (*net.Conn, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, errors.New("(error) failed to connect to durin")
	}
	return &conn, nil
}

func Raw(conn *net.Conn, message string) (string, error) {
	fmt.Fprintf(*conn, "%s\n", message)
	reader := bufio.NewReader(*conn)
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

func Set(conn *net.Conn, key string, value string) (string, error) {
	fmt.Fprintf(*conn, "set %s %s\n", key, value)
	reader := bufio.NewReader(*conn)
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

func Get(conn *net.Conn, key string) (string, error) {
	fmt.Fprintf(*conn, "get %s\n", key)
	reader := bufio.NewReader(*conn)
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

func Del(conn *net.Conn, key string) (string, error) {
	fmt.Fprintf(*conn, "del %s\n", key)
	reader := bufio.NewReader(*conn)
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

func Keys(conn *net.Conn) (string, error) {
	fmt.Fprintf(*conn, "keys\n")
	reader := bufio.NewReader(*conn)
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

func KeysPrefix(conn *net.Conn, prefix string) (string, error) {
	fmt.Fprintf(*conn, "keys %s\n", prefix)
	reader := bufio.NewReader(*conn)
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

func Json(conn *net.Conn, prefix string) (string, error) {
	fmt.Fprintf(*conn, "json %s\n", prefix)
	reader := bufio.NewReader(*conn)
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
