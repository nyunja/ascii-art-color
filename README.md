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

## Authors
**John Otieno,**

Apprentice Software Developer, Zone01 Kisumu

**David Odhiambo,**

Apprentice Software Developer, Zone01 Kisumu