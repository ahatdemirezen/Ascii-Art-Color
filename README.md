# Colorful Texts With ASCII Art
ASCII Art Generator is a program written in Go. This program takes text input and converts it into a colorful ascii art representation using ASCII characters based on the selected colors. 

## Features

- ***Colorful ASCII Art Creation:*** Create colorful ASCII art by displaying specific letters in the desired color.
- ***Various Color Options:*** Choose from different colors and display your texts in the color of your choice.
- ***Easy to Use:*** Utilizes a simple command-line interface for easy usage.
- ***File Handling:*** Read ASCII art templates from text files to display various texts in color.
- ***Error Handling:*** Provides user-friendly error messages for handling missing or incorrect inputs.

## How it Works

This program utilizes a command-line interface to create colorful ASCII art by altering the color of specific letters or words in the input text. Here's a breakdown of how it works:

1. **Parsing Command-Line Arguments:** The program starts by parsing the command-line arguments provided by the user. It looks for the "--color" flag to determine the color to be used and identifies the letters or words specified for coloring.

2. **Text Processing:** After parsing the arguments, the program processes the input text to identify the letters or words that need to be colored. It prepares the text for colorization based on the user's specifications.

3. **Colorizing Text:** Using the specified color and target letters/words, the program applies the chosen color to the corresponding characters in the text. It ensures that only the specified elements are colorized while maintaining the rest of the text's appearance.

4. **Output Generation:** Once the colorization process is complete, the program generates the final output, which is the colorful ASCII art representing the input text with the specified color alterations.

### Command Line Arguments

- **--color=color_name:** Specifies the color to be used for coloring the text. Available color options include red, green, blue, yellow, and others.
- **letters_to_be_colored:** Specifies the letters or words in the input text that should be colored.
- **"TEXT":** Represents the input text to be processed.
-  If you do not enter the letters or words to be colored, it will turn the entire text into the color you want

### Usage 

1. To color specific letters or words in the input text:
```bash
$ go run . [--color=color_name] [letters_to_be_colored] [TEXT]
```
Example 1:
```bash
$ go run . --color=blue "som" "something"
```
2. To color all text:
```bash
 $ go run . [--color=color_name] [TEXT]
```
Example 2:
```bash
$ go run . --color=blue "something"
```
---

#### Author
[Ahat Demirezen](https://github.com/ahatdemirezen)

#### LinkedIn
[Ahat Demirezen](https://www.linkedin.com/in/ahat-demirezen-8bb5262a9/)
 ## License
 [MIT](https://choosealicense.com/licenses/mit/)
