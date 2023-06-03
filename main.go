package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
)

func main() {
	c := NewPwnClient("", true)
	if err := c.Connect(); err != nil {
		panic(err)
	}
	defer c.Close()
	c.StartInteractive()

	c.RecvLine()
	c.SendLine("aiue")
	c.StartInteractive()
}

type PwnClient struct {
	addr   string
	conn   net.Conn
	stdout bool
}

func NewPwnClient(addr string, stdout bool) *PwnClient {
	return &PwnClient{addr: addr, stdout: stdout}
}

func (c *PwnClient) Connect() error {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *PwnClient) Close() {
	c.conn.Close()
}

func (c *PwnClient) RecvBytes(length int) []byte {
	buf := make([]byte, length)
	n, err := c.conn.Read(buf)
	if err != nil {
		return []byte{}
	}
	return buf[:n]
}

func (c *PwnClient) Recv(length int) string {
	buf := c.Recv(length)
	if len(buf) > 0 {
		return string(buf)
	}
	return ""
}

func (c *PwnClient) RecvLine() string {
	buf, err := c.RecvLineBytes()
	if err != nil {
		panic(err)
	}
	if len(buf) <= 0 {
		return string(buf)
	}
	str := string(buf)
	fmt.Println(str)
	return str
}

func (c *PwnClient) RecvLineBytes() ([]byte, error) {
	reader := bufio.NewReader(c.conn)
	line, _, err := reader.ReadLine()
	if err != nil {
		return []byte{}, err
	}
	return line, nil
}

func (c *PwnClient) SendLine(str string) int {
	writer := bufio.NewWriter(c.conn)
	n, err := writer.WriteString(str)
	if err != nil {
		panic(err)
	}
	writer.Flush()
	return n
}

func (c *PwnClient) StartInteractive() {
	scanner := bufio.NewScanner(os.Stdin)
	ctx := context.Background()
	go c.interactiveReceive(ctx)
	for {
		ok := scanner.Scan()
		if ok {
			c.SendLine(fmt.Sprintf("%s\n", scanner.Text()))
		}
	}
}

func (c *PwnClient) interactiveReceive(ctx context.Context) {
	for true {
		c.RecvLine()
	}
}
