package photo

import (
	"bytes"
	"errors"
	"fmt"
	"image/color"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"

	col "github.com/phanirithvij/colormatch/colormatch/color"
	"github.com/phanirithvij/colormatch/colormatch/log"
	"go.uber.org/zap"
)

// file: app/services/photo/get_histogram_data.rb

// HistData histogram data
type HistData struct {
	Width  int
	Height int
	Pixels int
	Colors []HistEntry
}

// HistEntry histogram entry
type HistEntry struct {
	Count int
	RGB   col.RGBA
	Hex   col.Hex
	HSB   col.HSB
	Lab   col.Lab
}

// NewHistEntry build hist entry from color
func NewHistEntry(color col.Color) (e *HistEntry) {
	e = &HistEntry{
		Count: 0,
		RGB:   color.RGB(),
		Hex:   color.Hex(),
		HSB:   color.HSB(),
		Lab:   color.Lab(),
	}
	return e
}

// GetHistogramData get the histogram data
// TODO functional arguments
func GetHistogramData(
	path string,
	colors int,
	resize bool,
	resizeDimension int,
) (HistData, error) {
	path = stripVersion(path)
	resizeStr := resizeStr(resize, resizeDimension)
	colorsx, err := makeHistogram(path, resizeStr, colors)
	if err != nil {
		log.ColoredLogger().Fatal("error", zap.Error(err))
	}
	log.ColoredLogger().Sugar().Debug("num colors in hist ", len(colorsx))
	dimens, err := getDimensions(path, resizeStr)
	if err != nil {
		log.ColoredLogger().Fatal("error", zap.Error(err))
	}
	log.ColoredLogger().Debug("dimens", zap.Any("w, h", dimens))
	return HistData{
		Colors: colorsx,
		Width:  dimens[0],
		Height: dimens[1],
		Pixels: dimens[0] * dimens[1],
	}, nil
}

func stripVersion(path string) string {
	re := regexp.MustCompile(`\?v=[\d]+`)
	return re.ReplaceAllString(path, "")
}

func resizeStr(resize bool, dimension int) string {
	if resize {
		return fmt.Sprintf("-resize %dx%d", dimension, dimension)
	}
	return ""
}

func makeHistogram(path string, resize string, colors int) ([]HistEntry, error) {
	args := []string{
		path, "-format", "%c",
	}
	if resize != "" {
		args = append(args, resize)
	}
	args = append(args,
		"-colors", strconv.Itoa(colors),
		"histogram:info:-",
	)
	out, err := basicExec("convert", args...)
	logger := log.ColoredLogger()
	if err != nil {
		logger.Error("an error occurred in convert", zap.Error(err))
		return nil, err
	}
	// TODO improve this hacky regex
	// original /(?<occurances>[\d]{1,8}):\s\(\s*(?<c1>[\d]{1,3}),\s*(?<c2>[\d]{1,3}),\s*(?<c3>[\d]{1,3})/
	re := regexp.MustCompile(`[ ]*([0-9]{1,8}):.*srgb\(([0-9]+)\..*,([0-9]+)\..*,([0-9]+)\..*\)`)
	match := re.FindAllStringSubmatch(out, -1)
	cols := make([]HistEntry, len(match))
	for i, name := range match {
		count, err := strconv.Atoi(name[1])
		if err != nil {
			return nil, err
		}
		r, err := strconv.Atoi(name[2])
		if err != nil {
			return nil, err
		}
		g, err := strconv.Atoi(name[3])
		if err != nil {
			return nil, err
		}
		b, err := strconv.Atoi(name[4])
		if err != nil {
			return nil, err
		}
		rgb := col.RGBA{NRGBA: color.NRGBA{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
			A: uint8(255),
		}}
		cols[i] = HistEntry{
			Count: count,
			RGB:   rgb,
			Hex:   rgb.Hex(),
			HSB:   rgb.HSB(),
			Lab:   rgb.Lab(),
		}
	}
	// sort -n -r
	sort.Slice(cols, func(i, j int) bool {
		return cols[i].Count < cols[j].Count
	})

	return cols, nil
}

// returns an array [w, h]
func getDimensions(path string, resize string) ([]int, error) {
	cmd := ""
	if resize == "" {
		cmd = fmt.Sprintf("convert %s -ping -format %%[fx:w],%%[fx:h] info:", path)
	} else {
		cmd = fmt.Sprintf("convert %s %s -ping -format %%[fx:w],%%[fx:h] info:", path, resize)
	}
	cmdargs := strings.Split(cmd, " ")
	out, err := basicExec(cmdargs[0], cmdargs[1:]...)
	if err != nil {
		return nil, err
	}
	wh := strings.Split(out, ",")
	w, err := strconv.Atoi(wh[0])
	h, err := strconv.Atoi(wh[1])
	if err != nil {
		return nil, err
	}
	return []int{w, h}, nil
}

func basicExec(name string, arg ...string) (string, error) {
	logger := log.ColoredLogger()
	logger.Debug("basicExec", zap.String("cmd", name), zap.Any("args", arg))
	defer logger.Sync()
	cmd := exec.Command(name, arg...)

	var out bytes.Buffer
	var errx bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errx

	err := cmd.Run()

	if err != nil {
		err = errors.New(errx.String())
		logger.Error("an error occurred", zap.Error(err))
		return "", err
	}

	stdout := out.String()
	if len(stdout) < 200 {
		logger.Debug("output", zap.String("output", stdout))
	}
	return stdout, nil
}
