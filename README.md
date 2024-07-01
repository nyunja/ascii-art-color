# ascii-art color

This program allows you to print ASCII art with customizable colors and styles.

## Usage

The program accepts various combinations of arguments to customize the output. Here are the supported formats:

**Format 1: With color, banner file, letters to be colored, and string**

```
go run . --color=<color> <substring-to-be-colored> <string> <banner>
```
Example:
```
go run . --color=red "H" "Hello" standard
```
**Format 2: With color, letters to be colored, and string**
```
go run . --color=<color> <substring-to-be-colored> <string>
```
Example:
```
go run . --color=blue el "Hello"
```

**Format 3: With color, string and banner style**
```
go run . --color=<color> <string> <banner-file> 
```
Example:
```
go run . --color=yellow Hello shadow
```
**Format 4: With color and string**
```
go run . --color=<color> <string>
```
Example:
```
go run . --color=green "Hello"
```

**Format 5: String only**
```
go run . <string>
```

Example:
```
go run . "Hello"
```

## Error Handling

If the program is not provided with the correct number and format of arguments, it will display an error message:

```
Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring-to-be-colored> "something"
```
Example:
```
go run . --color=red Hello
```

## Examples

To print "Hello" in red:
```
go run . --color=red Hello
```
To print "Hello" in blue with a shadow banner:
```
go run . --color=blue "Hello" shadow
```

To print "Hello" in default color with a Thinkertoy banner:
```
go run . "Hello"
```

## Notes

- The `<substring-to-be-colored>` argument should be enclosed in double quotes or single quotes if they contain special characters.
- The `<substring-to-be-colored>` and `"something"` arguments are case-sensitive and will only respect matching cases.
- Available colors include: red, green, blue, and others supported by your terminal.
- Available banner styles include: standard, shadow, and thinkertoy.
- Use `\\standard`, `\\shadow` and `\\thinkertoy` when you want to to use the banner file names as the substring to be colored or the main string.
- To use 2 colors i.e one for the substring and one for the what is left use the syntax `--color=<color>|<color>`, `<substring-to-be-colored>` and `"something"` or `--color=<color>|<color>` and `"something"`. If you use more than 1 color and no substring the first color will be used. The example prints red shadow output.
- You can provide multiple sub-strings by separating your sub-strings with a newline character ```'\n'```. For example ```"sub\nstring"```

```
go run . "--color=red|blue|green" "bus bu" shadow
```
- While the following commands color the substring red and the other characters blue.
```
go run . "--color=red|blue|green" bu "bus bu" shadow

go run . "--color=red|blue" bu "bus bu" shadow
```


## Color Library
Feel free to choose from the following color library and formats

```     
        Color Name       Hex                 RGB                      HSL

		{"red",         "#FF0000",          "rgb(255, 0, 0)",        "hsl(0, 100%, 50%)"},
		{"green",       "#00FF00",          "rgb(0, 255, 0)",        "hsl(120, 100%, 50%)"},
		{"blue",        "#0000FF",          "rgb(0, 0, 255)",        "hsl(240, 100%, 50%)"},
		{"orange",      "#FFA500",          "rgb(255, 165, 0)",      "hsl(39, 100%, 50%)"},
		{"yellow",      "#FFFF00",          "rgb(255, 255, 0)",      "hsl(60, 100%, 50%)"},
		{"black",       "#000000",          "rgb(0, 0, 0)",          "hsl(0, 0%, 0%)"},
		{"white",       "#FFFFFF",          "rgb(255, 255, 255)",    "hsl(0, 0%, 100%)"},
		{"gray",        "#808080",          "rgb(128, 128, 128)",    "hsl(0, 0%, 50%)"},
		{"pink",        "#FF00FF",          "rgb(255, 0, 255)",      "hsl(300, 100%, 50%)"},
		{"teal",        "#008080",          "rgb(0, 128, 128)",      "hsl(180, 100%, 25%)"},
		{"purple",      "#800080",          "rgb(128, 0, 128)",      "hsl(300, 100%, 25%)"},
		{"brown",       "#A52A2A",          "rgb(165, 42, 42)",      "hsl(0, 59%, 41%)", },
		{"beige",       "#F5F5DC",          "rgb(245, 245, 220)",    "hsl(60, 56%, 91%)", },
		{"tan",         "#D2B48C",          "rgb(210, 180, 140)",    "hsl(34, 44%, 69%)", },
		{"peach",       "#FFC0CB",          "rgb(255, 192, 203)",    "hsl(350, 100%, 88%)"},
		{"lime",        "#00FF00",          "rgb(0, 255, 0)",        "hsl(120, 100%, 50%)"},
		{"olive",       "#808000",          "rgb(128, 128, 0)",      "hsl(60, 100%, 25%)"},
		{"turquoise",   "#40E0D0",          "rgb(64, 224, 208)",     "hsl(174, 72%, 56%)"},
		{"navy Blue",   "#000080",          "rgb(0, 0, 128)",        "hsl(240, 100%, 25%)", },
		{"indigo",      "#4B0082",          "rgb(75, 0, 130)",       "hsl(275, 76%, 25%)"},
		{"violet",      "#8A2BE2",          "rgb(138, 43, 226)",     "hsl(271, 76%, 53%)"},
		{"lavender",    "#E6E6FA",          "rgb(230, 230, 250)",    "hsl(240, 100%, 94%)"},
		{"lilac",       "#C8A2C8",          "rgb(200, 162, 200)",    "hsl(270, 25%, 73%)"},
		{"maroon",      "#800000",          "rgb(128, 0, 0)",        "hsl(0, 100%, 25%)"},
		{"crimson",     "#DC143C",          "rgb(220, 20, 60)",      "hsl(348, 83%, 47%)"},
		{"fuchsia",     "#FF00FF",          "rgb(255, 0, 255)",      "hsl(300, 100%, 50%)"},
		{"coral",       "#FF7F50",          "rgb(255, 127, 80)",     "hsl(16, 100%, 66%)"},
		{"saffron",     "#F4C430",          "rgb(244, 196, 48)",     "hsl(45, 87%, 57%)"},
		{"sage",        "#B2AC88",          "rgb(178, 172, 136)",    "hsl(72, 22%, 53%)"},
		{"mint",        "#98FF98",          "rgb(152, 255, 152)",    "hsl(120, 100%, 78%)"},
		{"seafoam",     "#9FE2BF",          "rgb(159, 226, 191)",    "hsl(160, 66%, 75%)"},
		{"emerald",     "#50C878",          "rgb(80, 200, 120)",     "hsl(140, 76%, 55%)"},
		{"sapphire",    "#0F52BA",          "rgb(15, 82, 186)",      "hsl(220, 87%, 39%)"},
		{"ruby",        "#E0115F",          "rgb(224, 17, 95)",      "hsl(340, 85%, 47%)"},
		{"garnet",      "#733635",          "rgb(115, 54, 53)",      "hsl(6, 37%, 33%)"},
		{"ivory",       "#FFFFF0",          "rgb(255, 255, 240)",    "hsl(60, 100%, 97%)"},
		{"cream",       "#FFFDD0",          "rgb(255, 253, 208)",    "hsl(57, 100%, 91%)"},
		{"champagne",   "#F7E7CE",          "rgb(247, 231, 206)",    "hsl(39, 79%, 88%)"},
		{"camel",       "#C19A6B",          "rgb(193, 154, 107)",    "hsl(33, 41%, 59%)"},
		{"khaki",       "#C3B091",          "rgb(195, 176, 145)",    "hsl(37, 34%, 67%)"},
		{"mocha",       "#80655E",          "rgb(128, 101, 94)",     "hsl(13, 16%, 44%)"},
		{"umber",       "#635147",          "rgb(99, 81, 71)",       "hsl(24, 21%, 42%)"},
		{"sienna",      "#882D17",          "rgb(136, 45, 23)",      "hsl(10, 71%, 31%)"},
		{"charcoal",    "#36454F",          "rgb(54, 69, 79)",       "hsl(204, 20%, 26%)"},
		{"slate",       "#708090",          "rgb(112, 128, 144)",    "hsl(210, 13%, 50%)"},
		{"ash",         "#B2BEB5",          "rgb(178, 190, 181)",    "hsl(140, 12%, 72%)"},
```         
## Authors
**John Otieno,**

Apprentice Software Developer, Zone01 Kisumu

**David Odhiambo,**

Apprentice Software Developer, Zone01 Kisumu