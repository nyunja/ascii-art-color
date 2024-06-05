# ascii-art color

This program allows you to print ASCII art with customizable colors and styles.

## Usage

The program accepts various combinations of arguments to customize the output. Here are the supported formats:

**Format 1: With color, banner file, letters to be colored, and string**

```
go run . --color=<color> -<banner-file> <letters-to-be-colored> <string>
```
Example:
```
go run . --color=red -standard "H" "Hello"
```
**Format 2: With color, letters to be colored, and string**
```
go run . --color=<color> <letters-to-be-colored> <string>
```
Example:
```
go run . --color=blue "Hello"
```

**Format 3: With banner style, letters to be colored, and string**
```
go run . -<banner-file> <letters-to-be-colored> <string>
```
Example:
```
go run . -shadow "ello" "Hello"
```
**Format 4: With color and string**
```
go run . --color=<color> <string>
```
Example:
```
go run . --color=green "Hello"
```

**Format 5: With banner style and string**
```
go run . -<banner-file> <string>
```

Example:
```
go run . -thinkertoy "Hello"
```

**Format 6: String only**
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
go run . --color=blue -shadow "Hello"
```

To print "Hello" in default color with a Thinkertoy banner:
```
go run . -thinkertoy "Hello"
```
To print "Hello" in default color with the defalt -standard banner:
```
go run . "Hello"
```

## Notes

- The `<letters-to-be-colored>` argument should be enclosed in double quotes or single quotes if they contain special characters.
- Available colors include: red, green, blue, and others supported by your terminal.
- Available banner styles include: standard, shadow, and thinkertoy.
## Authors
**John Otieno,**

Apprentice Software Developer, Zone01 Kisumu

**David Odhiambo,**

Apprentice Software Developer, Zone01 Kisumu