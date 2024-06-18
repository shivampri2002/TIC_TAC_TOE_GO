package main

import (
	"bufio"
	// "fmt"
	"net"
	// "strings"
	"sync"
)

//when game is over close the connection

// TCPClient manages the connection to the game server
type TCPClient struct {
	conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
	mu     sync.Mutex
}

func (c *TCPClient) Connect(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	c.conn = conn
	c.reader = bufio.NewReader(conn)
	c.writer = bufio.NewWriter(conn)
	return nil
}

func (c *TCPClient) Send(message string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err := c.writer.WriteString(message + "\n")
	if err != nil {
		return err
	}
	return c.writer.Flush()
}

func (c *TCPClient) Read() (string, error) {
	return c.reader.ReadString('\n')
}

func (c *TCPClient) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.conn != nil {
		err := c.conn.Close()
		c.conn = nil
		return err
	}
	return nil
}
