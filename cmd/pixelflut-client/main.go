package main

import (
	"flag"
	"fmt"
	"github.com/patrick246/go-pixelflut/pkg/client"
	"strconv"
)

var (
	flagAddr = flag.String("addr", "localhost:1234", "Pixelflut server addr, in the host:port format")
	flagOp   = flag.String("op", "help", "Pixelflut command, like help, size, px, ...")
)

func main() {
	flag.Parse()

	c, err := client.NewPixelflutClient(*flagAddr)
	if err != nil {
		fmt.Printf("pixelflut connect error: %v\n", err)
		return
	}

	switch *flagOp {
	case "help":
		fmt.Println("not implemented yet")
	case "size":
		x, y, err := c.Size()
		if err != nil {
			fmt.Printf("pixelflut error: %v\n", err)
			return
		}
		fmt.Printf("%dx%d", x, y)
	case "px":
		remainingArgs := flag.Args()
		if len(remainingArgs) != 3 {
			fmt.Printf("need 3 args, x, y and hex color, got %d", len(remainingArgs))
			return
		}
		x, err := strconv.Atoi(remainingArgs[0])
		if err != nil {
			fmt.Printf("error reading x pos: %v", err)
			return
		}

		y, err := strconv.Atoi(remainingArgs[1])
		if err != nil {
			fmt.Printf("error reading y pos: %v", err)
			return
		}
		err = c.SetPixelString(x, y, remainingArgs[2])
		if err != nil {
			fmt.Printf("error sending command: %v", err)
		}
	}
}
