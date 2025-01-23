package client

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func (c *Color) ToString() string {
	if c.A != 0xFF {
		return fmt.Sprintf("%.2X%.2X%.2X%.2X", c.R, c.G, c.B, c.A)
	}
	return fmt.Sprintf("%.2X%.2X%.2X", c.R, c.G, c.B)
}

type WriteCommand struct {
	X, Y  int
	Color Color
}

type PixelflutClient struct {
	addr string
	conn net.Conn
}

func NewPixelflutClient(addr string) (*PixelflutClient, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &PixelflutClient{
		addr: addr,
		conn: conn,
	}, nil
}

func (p *PixelflutClient) Size() (int, int, error) {
	_, err := p.conn.Write([]byte("SIZE\n"))
	if err != nil {
		return 0, 0, err
	}

	readBuffer := make([]byte, 512)
	n, err := p.conn.Read(readBuffer)
	if err != nil {
		return 0, 0, err
	}

	sizeResponse := string(readBuffer[:n])
	endOfResponse := strings.Index(sizeResponse, "\n")
	if endOfResponse == -1 {
		return 0, 0, errors.New("malformed response, not correctly terminated")
	}

	sizeParts := strings.SplitN(sizeResponse[:endOfResponse], " ", 3)
	if len(sizeParts) != 3 {
		return 0, 0, errors.New("malformed response, not three parts: " + sizeResponse)
	}

	sizeX, err := strconv.Atoi(sizeParts[1])
	if err != nil {
		return 0, 0, errors.New("malformed response, X value not an int: " + sizeParts[1] + ", error " + err.Error())
	}

	sizeY, err := strconv.Atoi(sizeParts[2])
	if err != nil {
		return 0, 0, errors.New("malformed response, Y value not an int: " + sizeParts[2] + ", error " + err.Error())
	}

	return sizeX, sizeY, nil
}

func (p *PixelflutClient) SetPixel(x, y int, color Color) error {
	return p.SetPixelString(x, y, color.ToString())
}

func (p *PixelflutClient) SetPixelString(x, y int, color string) error {
	command := fmt.Sprintf("PX %d %d %s\n", x, y, color)
	_, err := p.conn.Write([]byte(command))
	return err
}

func (p *PixelflutClient) WriteFromChannel(commandChannel <-chan WriteCommand) {
	writer := bufio.NewWriterSize(p.conn, 64*1024)
	for cmd := range commandChannel {
		_, err := fmt.Fprintf(writer, "PX %d %d %s\n", cmd.X, cmd.Y, cmd.Color.ToString())
		if err != nil {
			fmt.Printf("error sending command: %v\n", err)
		}
	}
}
