package producers

import "github.com/patrick246/go-pixelflut/pkg/client"

type RectangleProducer struct {
	OffsetX, OffsetY, SizeX, SizeY int
	color                          client.Color
}

func (rp *RectangleProducer) Produce(commandChannel chan<- client.WriteCommand) {
	for {
		for i := 0; i < rp.SizeY; i++ {
			for j := 0; j < rp.SizeX; j++ {
				commandChannel <- client.WriteCommand{
					X: rp.OffsetX + j,
					Y: rp.OffsetY + i,
					Color: client.Color{
						R: 0xff,
						G: 0xa5,
						B: 0,
						A: 0xff,
					},
				}
			}
		}
	}
}
