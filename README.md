# Go Utils

A comprehensive Go library providing a wide range of utility functions for text manipulation, file handling, terminal output with colors, and data manipulation. Useful for various common tasks in Go projects. Primarily because I copy+paste all these functions in my other Go projects, thought I should make it its own package.

## Features

- **Text and Terminal Utilities**:
  - Colorize text output in terminal using CSS color codes.
  - Capture input from stdio with optional colored prompts.
  - Remove quotes from input strings using regular expressions.
- **File Handling**:
  - Globbing function to find files matching specific patterns.
  - Read from standard input.
  - Write content directly to a file.
- **Data Manipulation**:
  - Check presence of a string in a slice.
  - Deduplicate entries in a string slice.
  - Extract keys from a JSON object.
  - Break out maps into separate slices of keys and values.
  - Reverse slices of any type.
  - Convert a slice of strings to a slice of empty interfaces.
- **Stream Handling**:
  - Display ongoing completion stream results from the OpenAI API.

## Getting Started

### Prerequisites

This project is built using Go 1.22.0. Ensure you have Go installed on your system. You can download it from the [official Go website](https://golang.org/dl/).

### Installation

```bash
go get github.com/chand1012/go-utils@latest
```

### Usage

All functions are contained within the `utils` package, and thus you should import this package to access the functionalities:

```go
import (
    "github.com/chand1012/go-utils/utils"
)
```

#### Examples

1. **Printing colored text**: 

   ```go
   utils.PrintColoredText("Hello, world!", "#FF5733")
   utils.PrintColoredTextLn("This will be on a new line with green color.", "#00FF00")
   ```

2. **File globbing**:

   ```go
   files, err := utils.GlobAll([]string{"*.txt", "*.docx"})
   if err != nil {
       fmt.Println("Error:", err)
   }
   fmt.Println("Matched files:", files)
   ```

3. **Reading input**:

   ```go
   input, err := utils.Input("Enter your name: ")
   if err != nil {
       fmt.Println("Error:", err)
   }
   fmt.Println("You entered:", input)

   coloredInput, err := utils.InputWithColor("Enter your age: ", "#4CAF50")
   if err != nil {
       fmt.Println("Error:", err)
   }
   fmt.Println("You entered:", coloredInput)
   ```

4. **Working with JSON**:

   ```go
   json := map[string]interface{}{"name": "John", "age": 30}
   keys, err := utils.GetJSONKeys(json)
   if err != nil {
       fmt.Println("Error:", err)
   }
   fmt.Println("JSON Keys:", keys)
   ```

### Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to help improve the project.
