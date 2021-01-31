package color

import (
	"image/color"
	"regexp"
	"strconv"
	"strings"

	"github.com/phanirithvij/colormatch/colormatch/log"
	"go.uber.org/zap"
)

// Hex hex string color format
type Hex struct {
	h string
}

// NewHex exported way to create hex
func NewHex(str string) Hex {
	logger := log.ColoredLogger()
	var re *regexp.Regexp
	// #rgb #rrggbb #rrggbbaa
	re = regexp.MustCompile(`(#|0x|0X){1}(([a-fA-F0-9]{3}){1,2}|([a-fA-F0-9]{2}){4})$`)
	logger.Debug("hex matches", zap.String("match", re.FindString(str)))
	if !re.MatchString(str) {
		logger.Warn("Invalid hex", zap.String("hex", str))
		return Hex{}
	}
	str = strings.Replace(str, "#", "", 1)
	str = strings.Replace(str, "0x", "", 1)
	str = strings.Replace(str, "0X", "", 1)
	if len(str) == 3 {
		re = regexp.MustCompile(`(\w)`)
		str = re.ReplaceAllString(str, "$1$1")
	}
	return Hex{h: "#" + str}
}

// RGBHex will return #rrggbb without alpha
func (hex Hex) RGBHex() string {
	return hex.h[:7]
}

// Hex noop conversion use RGBHex to get #rrggbb
func (hex Hex) Hex() Hex {
	return hex
}

// HSB convert to hsb
func (hex Hex) HSB() HSB {
	return hex.RGB().HSB()
}

// Lab convert to Lab
func (hex Hex) Lab() Lab {
	return hex.RGB().Lab()
}

// RGB convert to rgba
func (hex Hex) RGB() RGBA {
	logger := log.ColoredLogger()
	// flushes buffer, if any
	defer logger.Sync()
	log := logger.Sugar()
	str := hex.h
	str = strings.Replace(str, "#", "", 1)
	var re *regexp.Regexp
	reStr := `(..)(..)(..)`
	if len(str) == 8 {
		reStr = `(..)(..)(..)(..)`
	}
	re = regexp.MustCompile(reStr)
	res := re.FindStringSubmatch(str)
	if res == nil {
		log.Warnw(
			"hex string is invalid regex failed",
			"string", str,
			"expected", "#rrggbb or #rrggbbaa",
		)
		return RGBA{}
	}
	log.Debug(res)
	var r uint8 = str2hex(res[1])
	var g uint8 = str2hex(res[2])
	var b uint8 = str2hex(res[3])
	var a uint8 = 255
	if len(res) > 4 {
		a = str2hex(res[4])
	}
	return RGBA{
		NRGBA: color.NRGBA{R: r, G: g, B: b, A: a},
	}
}

// RGBA image/color.Color compat
func (hex Hex) RGBA() (uint32, uint32, uint32, uint32) {
	return hex.RGB().RGBA()
}

// HexModel is the color.Model for the Hex type.
var HexModel = color.ModelFunc(func(c color.Color) color.Color {
	if _, ok := c.(Hex); ok {
		return c
	}
	nrgba := color.NRGBAModel.Convert(c).(color.NRGBA)
	return RGBA{NRGBA: nrgba}.Hex()
})

// str2hex
func str2hex(str string) uint8 {
	n, err := strconv.ParseUint(str, 16, 8)
	if err != nil {
		return 0
	}
	return uint8(n)
}

// String to string
func (hex Hex) String() string {
	return hex.h
}
