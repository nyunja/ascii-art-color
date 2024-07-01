package Lib

import "strings"

// colorPicker picks the ANSI color code and reset code based on the provided colorFlag.
// It returns the ANSI color code and reset code as strings.
func colorPicker(colorFlag string) (string, string, string) {
	// Initialize utility variables
	// Set default color1 if colorFlag
	color1 := ""
	color2 := ""
	reset := "\033[0m"

	// Define hex, RGB, hsl, and ANSI color codes
	colors := []struct {
		Name string
		Hex  string
		RGB  string
		HSL  string
		ANSI string
	}{
		{"red", "#FF0000", "rgb(255, 0, 0)", "hsl(0, 100%, 50%)", "\033[31m"},
		{"green", "#00FF00", "rgb(0, 255, 0)", "hsl(120, 100%, 50%)", "\033[32m"},
		{"blue", "#0000FF", "rgb(0, 0, 255)", "hsl(240, 100%, 50%)", "\033[34m"},
		{"orange", "#FFA500", "rgb(255, 165, 0)", "hsl(39, 100%, 50%)", "\033[38;5;208m"},
		{"yellow", "#FFFF00", "rgb(255, 255, 0)", "hsl(60, 100%, 50%)", "\033[33m"},
		{"black", "#000000", "rgb(0, 0, 0)", "hsl(0, 0%, 0%)", "\033[30m"},
		{"white", "#FFFFFF", "rgb(255, 255, 255)", "hsl(0, 0%, 100%)", ""},
		{"gray", "#808080", "rgb(128, 128, 128)", "hsl(0, 0%, 50%)", "\033[90m"},
		{"pink", "#FF00FF", "rgb(255, 0, 255)", "hsl(300, 100%, 50%)", "\033[95m"},
		{"teal", "#008080", "rgb(0, 128, 128)", "hsl(180, 100%, 25%)", "\033[36m"},
		{"purple", "#800080", "rgb(128, 0, 128)", "hsl(300, 100%, 25%)", "\033[35m"},
		{"brown", "#A52A2A", "rgb(165, 42, 42)", "hsl(0, 59%, 41%)", "\033[33;2m"},
		{"beige", "#F5F5DC", "rgb(245, 245, 220)", "hsl(60, 56%, 91%)", "\033[33;2m"},
		{"tan", "#D2B48C", "rgb(210, 180, 140)", "hsl(34, 44%, 69%)", "\033[33;2m"},
		{"peach", "#FFC0CB", "rgb(255, 192, 203)", "hsl(350, 100%, 88%)", "\033[95m"},
		{"lime", "#00FF00", "rgb(0, 255, 0)", "hsl(120, 100%, 50%)", "\033[92m"},
		{"olive", "#808000", "rgb(128, 128, 0)", "hsl(60, 100%, 25%)", "\033[92m"},
		{"turquoise", "#40E0D0", "rgb(64, 224, 208)", "hsl(174, 72%, 56%)", "\033[96m"},
		{"navy Blue", "#000080", "rgb(0, 0, 128)", "hsl(240, 100%, 25%)", "\033[34;1m"},
		{"indigo", "#4B0082", "rgb(75, 0, 130)", "hsl(275, 76%, 25%)", "\033[94m"},
		{"violet", "#8A2BE2", "rgb(138, 43, 226)", "hsl(271, 76%, 53%)", "\033[94m"},
		{"lavender", "#E6E6FA", "rgb(230, 230, 250)", "hsl(240, 100%, 94%)", "\033[94m"},
		{"lilac", "#C8A2C8", "rgb(200, 162, 200)", "hsl(270, 25%, 73%)", "\033[94m"},
		{"maroon", "#800000", "rgb(128, 0, 0)", "hsl(0, 100%, 25%)", "\033[31;2m"},
		{"crimson", "#DC143C", "rgb(220, 20, 60)", "hsl(348, 83%, 47%)", "\033[31;2m"},
		{"fuchsia", "#FF00FF", "rgb(255, 0, 255)", "hsl(300, 100%, 50%)", "\033[95m"},
		{"coral", "#FF7F50", "rgb(255, 127, 80)", "hsl(16, 100%, 66%)", "\033[95m"},
		{"saffron", "#F4C430", "rgb(244, 196, 48)", "hsl(45, 87%, 57%)", "\033[93m"},
		{"sage", "#B2AC88", "rgb(178, 172, 136)", "hsl(72, 22%, 53%)", "\033[92m"},
		{"mint", "#98FF98", "rgb(152, 255, 152)", "hsl(120, 100%, 78%)", "\033[92m"},
		{"seafoam", "#9FE2BF", "rgb(159, 226, 191)", "hsl(160, 66%, 75%)", "\033[96m"},
		{"emerald", "#50C878", "rgb(80, 200, 120)", "hsl(140, 76%, 55%)", "\033[92m"},
		{"sapphire", "#0F52BA", "rgb(15, 82, 186)", "hsl(220, 87%, 39%)", "\033[94m"},
		{"ruby", "#E0115F", "rgb(224, 17, 95)", "hsl(340, 85%, 47%)", "\033[91m"},
		{"garnet", "#733635", "rgb(115, 54, 53)", "hsl(6, 37%, 33%)", "\033[91m"},
		{"ivory", "#FFFFF0", "rgb(255, 255, 240)", "hsl(60, 100%, 97%)", "\033[97m"},
		{"cream", "#FFFDD0", "rgb(255, 253, 208)", "hsl(57, 100%, 91%)", "\033[97m"},
		{"champagne", "#F7E7CE", "rgb(247, 231, 206)", "hsl(39, 79%, 88%)", "\033[97m"},
		{"camel", "#C19A6B", "rgb(193, 154, 107)", "hsl(33, 41%, 59%)", "\033[33;2m"},
		{"khaki", "#C3B091", "rgb(195, 176, 145)", "hsl(37, 34%, 67%)", "\033[33;2m"},
		{"mocha", "#80655E", "rgb(128, 101, 94)", "hsl(13, 16%, 44%)", "\033[33;2m"},
		{"umber", "#635147", "rgb(99, 81, 71)", "hsl(24, 21%, 42%)", "\033[33;2m"},
		{"sienna", "#882D17", "rgb(136, 45, 23)", "hsl(10, 71%, 31%)", "\033[33;2m"},
		{"charcoal", "#36454F", "rgb(54, 69, 79)", "hsl(204, 20%, 26%)", "\033[90m"},
		{"slate", "#708090", "rgb(112, 128, 144)", "hsl(210, 13%, 50%)", "\033[90m"},
		{"ash", "#B2BEB5", "rgb(178, 190, 181)", "hsl(140, 12%, 72%)", "\033[90m"},
	}

	// Handle cases where colorFlag is provided with valid format
	if len(colorFlag) > 0 {
		if hasPrefix(colorFlag, "--color=") {
			key := trimSpace(colorFlag[8:])
			if key == "" {
				PrintError()
			} else {
				colorInput := strings.Split(key, "|")
				// Get ANSI color code based on the provided key
				var match bool
				for _, item := range colors {
					if item.Name == toLower(colorInput[0]) || item.Hex == colorInput[0] || item.HSL == toLower(colorInput[0]) || item.RGB == toLower(colorInput[0]) {
						color1 = item.ANSI
						match = true
					}
					if len(colorInput) > 1 && (item.Name == toLower(colorInput[1]) || item.Hex == colorInput[1] || item.HSL == toLower(colorInput[1]) || item.RGB == toLower(colorInput[1])) {
						color2 = item.ANSI
						match = true
					}
				}

				// Print error and end program if color is not contained in the color librabry
				if !match {
					PrintError()
				}
			}
		} else {
			PrintError()
		}
	}

	// Return the ANSI color code and reset code
	return color1, color2, reset
}
