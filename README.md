# ASCII Art Generator in Go

## Description

This is a simple and efficient ASCII Art generator written in Go. The program takes a string input and converts it into ASCII characters based on a specified font. It dynamically loads banner font files and processes multiple lines while maintaining correct formatting.

## Features

- Supports three fonts: standard, shadow, and thinkertoy.
- Maintains text formatting: Correct alignment, spacing & multi-line handling.
- Optimized for performance: Handles large inputs efficiently.
- Lightweight: Uses only Go’s standard library.
- CLI-based: Simple command-line interface for fast execution.

## Compatibility

This program is designed for Linux/macOS.  
Windows users should use WSL (Windows Subsystem for Linux).

## Installation

### Prerequisites

- [Go](https://go.dev/) installed on your system.

### Clone the repository

```sh
git clone <repository-url>
cd ascii-art-go
```

### Run the program

```sh
go run . "Hello World"
```

### You can also specify a font:

```
go run . "Hello World" "standard"
```

### To check for special characters like newlines in the output, you can use:

```sh
go run . "Hello World" | cat -e
```


## Usage

The program accepts a string input and an optional font argument. \
If the font is not specified, it defaults to `standard.txt`.

### Examples & Output:

```sh
go run . "" | cat -e
           # nothing printed
```


```sh
go run . "\n" | cat -e
$          # a single blank line printed
```


```sh
go run . "Hello" | cat -e
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
```

```
go run . "Hello" "shadow" | cat -e
                                      $
_|    _|            _|  _|            $
_|    _|    _|_|    _|  _|    _|_|    $
_|_|_|_|  _|_|_|_|  _|  _|  _|    _|  $
_|    _|  _|        _|  _|  _|    _|  $
_|    _|    _|_|_|  _|  _|    _|_|    $
                                      $
                                      $
```

## Character Support

This program supports:

Numbers (0-9)

Latin letters (A-Z, a-z)

Spaces

Special characters (e.g., !@#$%^&*()-_=+[]{}|;:'",.<>?/)

Newline character (\n) for multi-line formatting

For characters that have special meaning in the shell (e.g., ", ', \, &, |, *), you must escape them using a backslash (\\\) so they appear as literal characters.

Example:

```sh
go run . "Hello \"World\" \|" | cat -e

 _    _            _    _                    _ _   __          __                    _        _    _ _          __        _   $
| |  | |          | |  | |                  ( | )  \ \        / /                   | |      | |  ( | )         \ \      | |  $
| |__| |    ___   | |  | |    ___            V V    \ \  /\  / /     ___     _ __   | |    __| |   V V           \ \     | |  $
|  __  |   / _ \  | |  | |   / _ \                   \ \/  \/ /     / _ \   | '__|  | |   / _` |                  \ \    | |  $
| |  | |  |  __/  | |  | |  | (_) |                   \  /\  /     | (_) |  | |     | |  | (_| |                   \ \   | |  $
|_|  |_|   \___|  |_|  |_|   \___/                     \/  \/       \___/   |_|     |_|   \__,_|                    \_\  | |  $
                                                                                                                         | |  $
                                                                                                                         |_|  $
```

## Error Handling

- If no input is provided:
  ```sh
  go run .
  Error: No input provided.
  ```
- If the font file is missing:
  ```sh
  go run . "Test" "missingfont"
  Error: Banner font "missingfont" does not exist.
  ```
- If the font file is empty or malformed:
  ```sh
  go run . "Test" "corruptfont"
  Error: Banner file "corruptfont.txt" is empty.
  ```

## Testing

Automated tests are included to validate the ASCII rendering. To run tests:

```sh
go test ./tests
```

## License

This project is open-source and available under the MIT License.

## License

This project is open-source and licensed under the MIT License - see the [LICENSE](LICENSE) file for details.


