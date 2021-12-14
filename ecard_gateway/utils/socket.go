package utils

import (
    "fmt"
	"net"
)

func RecvAll(conn net.Conn, b []byte) error {
	var receiver []byte
	for len(receiver) < len(b) {
		temp := make([]byte, len(b))
		n, err := conn.Read(temp)
		if err != nil {
			return err
		}
		receiver = append(receiver, temp[:n]...)
	}
	copy(b, receiver)
	return nil
}

func SendAll(conn net.Conn, b []byte) error {
	count := 0
	for count < len(b) {
		n, err := conn.Write(b[count:])
		if err != nil {
			return fmt.Errorf("error writing: " + err.Error())
		}
		if n == 0 {
			return fmt.Errorf("error writing: no enough data to write")
		}
		count += n
	}
	return nil
}