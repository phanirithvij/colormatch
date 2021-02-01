package scripts

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/phanirithvij/colormatch/colormatch/log"
	"github.com/phanirithvij/colormatch/colormatch/models"
	"go.uber.org/zap"
)

// GenerateColormap generates a colormap
func GenerateColormap(colors []models.ColorModel) {
	logger := log.ColoredLogger()
	dateStr := strings.ReplaceAll(time.Now().Format(time.UnixDate), ":", "-")
	filepath := fmt.Sprintf("lib/assets/images/colormap_%s.png", dateStr)
	// alright, this is going to be a very long terminal command. Hopefully there aren't limits.
	queryStr := fmt.Sprintf("convert -size %dx1 xc:white ", len(colors))
	for i, c := range colors {
		queryStr += fmt.Sprintf("-fill '%s' -draw 'point %d,0' ", c.Hex, i)
	}
	queryStr += filepath

	cmd := exec.Command(queryStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	logger.Error("error", zap.Error(cmd.Run()))
}
