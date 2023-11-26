package main

import (
    "github.com/johnfercher/maroto/v2"
    "github.com/johnfercher/maroto/v2/pkg/components/text"
    "github.com/johnfercher/maroto/v2/pkg/props"
    "github.com/johnfercher/maroto/v2/pkg/consts/align"
    "os"
    "fmt"
    "flag"
    "bufio"
    "strings"
)

func init() {
    flag.Usage = func() {
        h := []string{
            "text to pdf utility written in Go",
            "",
            "Options:",
            "  -i, --input <path>         path to input text file",
            "  -o, --output <path>       path to output pdf file",
            "  -f, --font <font>         font to use, default is Arial",
            "",
        }

        fmt.Fprintf(os.Stderr, strings.Join(h, "\n"))
    }
}


func main() {
    var inputFile string
    flag.StringVar(&inputFile, "input", "", "")
    flag.StringVar(&inputFile, "i", "", "")


    var outputFile string
    flag.StringVar(&outputFile, "output", "", "")
    flag.StringVar(&outputFile, "o", "", "")

    // var font string
    // var font_size int

    flag.Parse()

    if inputFile == "" {
        fmt.Println("No input file specified")
        return
    }
 
    //read input file and add to pdf
    file, err := os.Open(inputFile) // replace yourfile.txt with your file name
    if err != nil {
        fmt.Println("File reading error", err)
    }
    defer file.Close() // Make sure to close the file when you're done

    m := maroto.New()

    scanner := bufio.NewScanner(file)

    // Use Scan() to read the next line.
    for scanner.Scan() {
        m.AddRow(10, text.NewCol(12, scanner.Text(), props.Text{
            Size:  10,
            Align: align.Left,
        }))
    }

    // Check for errors during Scan. End of file is expected and not reported by Scan as an error.
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file", err)
    }

    document, err := m.Generate()
    if err != nil {
        fmt.Println("Error generating pdf file", err)
    }
    err = document.Save(outputFile)
    if err != nil {
        fmt.Println("Error saving pdf file", err)
    }
    return;
}
