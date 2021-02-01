package main

import (
	"os"

	"github.com/phanirithvij/colormatch/colormatch/color"
	clog "github.com/phanirithvij/colormatch/colormatch/log"
	"github.com/phanirithvij/colormatch/colormatch/photo"
)

func main() {
	logger := clog.ColoredLogger()
	defer logger.Sync()
	log := logger.Sugar()
	hex := color.NewHex(os.Args[1])
	log.Info("rgb ", hex.RGB())
	log.Info("hex ", hex)
	log.Info("csshex ", hex.RGBHex())
	data, err := photo.GetHistogramData(os.Args[2], 22, false, 0)
	if err != nil {
		log.Fatal(err)
	}
	log.Infow("main", "hist", data)
}
