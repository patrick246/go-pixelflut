package producers

import (
	"image"
	"log"
	"os"

	"github.com/patrick246/go-pixelflut/pkg/client"
)

type PrerenderedImageProducer struct {
	buf []client.WriteCommand
}

func NewPrerenderedImageProducer(filepath string) (*PrerenderedImageProducer, error) {
	fileReader, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(fileReader)
	if err != nil {
		log.Fatal(err)
	}

	prod := &PrerenderedImageProducer{
		buf: make([]client.WriteCommand, 0, img.Bounds().Size().Y*img.Bounds().Size().X),
	}

	for y := range img.Bounds().Size().Y {
		for x := range img.Bounds().Size().X {
			r, g, b, a := img.At(x, y).RGBA()
			if a == 0 {
				continue
			}

			prod.buf = append(prod.buf, client.WriteCommand{
				X: x,
				Y: y,
				Color: client.Color{
					R: uint8(r >> 8),
					G: uint8(g >> 8),
					B: uint8(b >> 8),
					A: uint8(a >> 8),
				},
			})
		}
	}

	return prod, nil
}

func (p *PrerenderedImageProducer) Produce(commands chan<- client.WriteCommand) {
	for {
		for i := range p.buf {
			commands <- p.buf[i]
		}
	}
}
