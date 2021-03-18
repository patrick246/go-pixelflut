package producers

import "github.com/patrick246/go-pixelflut/pkg/client"

type Producer interface {
	Produce(chan<- client.WriteCommand)
}
