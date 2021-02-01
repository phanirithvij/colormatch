package photo

import (
	"testing"

	"github.com/phanirithvij/colormatch/colormatch/color"
)

// TestSaturatedColor returns true on a saturated colour
func TestSaturatedColor(t *testing.T) {
	c := color.NewHex("#004b96")
	e := NewHistEntry(c)
	valid := ValidHue(*e)
	if !valid {
		t.Fatalf(`ValidHue("#004b96") = %v, want true`, valid)
	}
}

// TestBrightSaturatedColor returns true on a bright, saturated colour
func TestBrightSaturatedColor(t *testing.T) {
	c := color.NewHex("#2aff00")
	e := NewHistEntry(c)
	valid := ValidHue(*e)
	if !valid {
		t.Fatalf(`ValidHue("#2aff00") = %v, want true`, valid)
	}
}

// TestSoftValidColor returns true on a soft but valid colour
func TestSoftValidColor(t *testing.T) {
	c := color.NewHex("#41573d")
	e := NewHistEntry(c)
	valid := ValidHue(*e)
	if !valid {
		t.Fatalf(`ValidHue("#41573d") = %v, want true`, valid)
	}
}

// TestWhiteColor returns false on white
func TestWhiteColor(t *testing.T) {
	c := color.NewHex("#FFFFFF")
	e := NewHistEntry(c)
	valid := ValidHue(*e)
	if valid {
		t.Fatalf(`ValidHue("#FFFFFF") = %v, want false`, valid)
	}
}

// TestBlackColor returns false on black
func TestBlackColor(t *testing.T) {
	c := color.NewHex("#000")
	e := NewHistEntry(c)
	valid := ValidHue(*e)
	if valid {
		t.Fatalf(`ValidHue("#000") = %v, want false`, valid)
	}
}

// TestGreyColor returns false on grey
func TestGreyColor(t *testing.T) {
	c := color.NewHex("#777")
	e := NewHistEntry(c)
	valid := ValidHue(*e)
	if valid {
		t.Fatalf(`ValidHue("#777") = %v, want false`, valid)
	}
}

// TestNavyBlueColor returns false on a navy blue
func TestNavyBlueColor(t *testing.T) {
	c := color.NewHex("#00091a")
	e := NewHistEntry(c)
	valid := ValidHue(*e)
	if valid {
		t.Fatalf(`ValidHue("#00091a") = %v, want false`, valid)
	}
}
