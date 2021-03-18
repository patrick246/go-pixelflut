package main

import (
	"flag"
	"github.com/patrick246/go-pixelflut/pkg/client"
	"github.com/patrick246/go-pixelflut/pkg/producers"
	"log"
)

var flagAddr = flag.String("addr", "localhost:1234", "Pixelflut server address, hostname:port")

func main() {
	flag.Parse()

	var clients []*client.PixelflutClient
	for i := 0; i < 8; i++ {
		c, err := client.NewPixelflutClient(*flagAddr)
		if err != nil {
			log.Fatal(err)
		}
		clients = append(clients, c)
	}

	commands := make(chan client.WriteCommand)

	for _, c := range clients {
		go c.WriteFromChannel(commands)
	}

	var prod []producers.Producer
	for i := 0; i < 4; i++ {
		prod = append(prod, &producers.ImageProducer{
			Filepath: "/home/patrick/Bilder/logov1_small.png",
			OffsetY:  200,
			OffsetX:  150,
		})
	}

	for _, p := range prod {
		go p.Produce(commands)
	}

	select {}
}
