package producers

import (
	"github.com/patrick246/go-pixelflut/pkg/client"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math/rand"
	"os"
)

type ImageProducer struct {
	Filepath         string
	OffsetX, OffsetY int
}

func (i *ImageProducer) Produce(commands chan<- client.WriteCommand) {
	fileReader, err := os.Open(i.Filepath)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(fileReader)
	if err != nil {
		log.Fatal(err)
	}

	for {
		y := rand.Intn(img.Bounds().Size().Y)
		x := rand.Intn(img.Bounds().Size().X)

		r, g, b, a := img.At(x, y).RGBA()
		if a == 0 {
			continue
		}
		commands <- client.WriteCommand{
			X: i.OffsetX + x,
			Y: i.OffsetY + y,
			Color: client.Color{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			},
		}
	}

}
