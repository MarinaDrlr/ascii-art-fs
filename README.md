# ASCII Art FS (Font Selector) in Go

## Description

This is a simple and efficient ASCII Art generator written in Go. The program takes a string input and converts it into ASCII characters based on a specified font. It dynamically loads banner font files and processes multiple lines while maintaining correct formatting.

## Features

- Supports three fonts: standard, shadow, and thinkertoy.
- Maintains text formatting: Correct alignment, spacing & multi-line handling.
- Optimized for performance: It handles large inputs efficiently.
- Lightweight: Uses only Go’s standard library.
- CLI-based: Simple command-line interface for fast execution.

## Compatibility

This program is designed for Linux/macOS.\
Windows users should use WSL (Windows Subsystem for Linux).

## Installation

### Prerequisites

- [Go](https://go.dev/) installed on your system.

### Clone the repository

```sh
git clone https://github.com/MarinaDrlr/ascii-art-fs
cd ascii-art-fs
```

### Run the program

```sh
go run . "Hello World" shadow
```

### To check for special characters like newlines in the output, you can use:

```sh
go run . "Hello World" shadow | cat -e
```

## Usage

The program requires exactly two arguments: a string input and a font name.

### Examples & Output:

```sh
go run . "" | cat -e
Usage: go run . [STRING] [BANNER]
EX: go run . "something" standard
```

```sh
go run . "\n" | cat -e
Usage: go run . [STRING] [BANNER]
EX: go run . "something" standard
```

```sh
go run . "Hello" standard | cat -e
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
go run . "Hello" shadow | cat -e
                                      $
_|    _|            _|  _|            $
_|    _|    _|_|    _|  _|    _|_|    $
_|_|_|_|  _|_|_|_|  _|  _|  _|    _|  $
_|    _|  _|        _|  _|  _|    _|  $
_|    _|    _|_|_|  _|  _|    _|_|    $
                                      $
                                      $
```

```
go run . "Hello" thinkertoy | cat -e
                      $
o  o       o  o       $
|  |       |  |       $
O--O  o-o  |  |  o-o  $
|  |  |-'  |  |  | |  $
o  o  o-o  o  o  o-o  $
                      $
                      $
```

## Character Support

This program supports:

- Numbers (0-9)
- Latin letters (A-Z, a-z)
- Spaces
- Special characters (e.g., !@#\$%^&\*()-\_=+[]{}|;:'",.<>?/)
- Newline character (`\n`) for multi-line formatting

For characters that have special meaning in the shell (e.g., `"`, `'`, `\\`, `&`, `|`, `*`), you must escape them using a backslash (`\\`) so they appear as literal characters.

Example:

```sh
go run . "Hello World" standard | cat -e
 _    _            _    _                   __          __                    _        _   $
| |  | |          | |  | |                  \ \        / /                   | |      | |  $
| |__| |    ___   | |  | |    ___            \ \  /\  / /     ___     _ __   | |    __| |  $
|  __  |   / _ \  | |  | |   / _ \            \ \/  \/ /     / _ \   | '__|  | |   / _` |  $
| |  | |  |  __/  | |  | |  | (_) |            \  /\  /     | (_) |  | |     | |  | (_| |  $
|_|  |_|   \___|  |_|  |_|   \___/              \/  \/       \___/   |_|     |_|   \__,_|  $
                                                                                           $
                                                                                           $
```

## Error Handling

- If the number of arguments is not exactly 2:

  - If **no input is provided**:

    ```sh
    go run .
    Usage: go run . [STRING] [BANNER]
    EX: go run . "something" standard
    ```

  - If **only one argument is provided**:

    ```sh
    go run . "Hello"
    Usage: go run . [STRING] [BANNER]
    EX: go run . "something" standard
    ```

  - If **more than two arguments are provided**:

    ```sh
    go run . "Hello" standard plus
    Usage: go run . [STRING] [BANNER]
    EX: go run . "something" standard
    ```

- If the font file is missing:

  ```sh
  go run . "Test" missingfont
  Error: Banner font missingfont does not exist.
  ```

- If the font file is empty:

  ```sh
  go run . "Test" empty
  Error: Banner file empty.txt is empty.
  ```

  If the font file is malformed:

  ```sh
  go run . "Test" broken
  Error: Banner file broken.txt is corrupted.
  ```

## Testing

Unit tests are included to validate the ASCII rendering. To run tests:

```sh
go test
```

or

```sh
go test -v
```

for verbose output with detailed information about each test case.

### Sample test output

```sh
=== RUN   TestNormalizeInput
--- PASS: TestNormalizeInput (0.00s)
=== RUN   TestNormalizeInputNewlines
--- PASS: TestNormalizeInputNewlines (0.00s)
=== RUN   TestLoadBanner
--- PASS: TestLoadBanner (0.00s)
=== RUN   TestLoadCorruptedBanner
--- PASS: TestLoadCorruptedBanner (0.00s)
=== RUN   TestLoadMissingFontFile
--- PASS: TestLoadMissingFontFile (0.00s)
=== RUN   TestGenerateASCIIArt
--- PASS: TestGenerateASCIIArt (0.00s)
=== RUN   TestGenerateASCIIArtUnsupportedChar
--- PASS: TestGenerateASCIIArtUnsupportedChar (0.00s)
PASS
ok      fs      0.003s
```

## License

This project is open-source and licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

