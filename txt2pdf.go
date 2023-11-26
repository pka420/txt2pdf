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

    flag.Parse()

    if inputFile == "" {
        fmt.Println("No input file specified")
        return
    }
 
    file, err := os.Open(inputFile) 
    if err != nil {
        fmt.Println("File reading error", err)
    }
    defer file.Close()

    m := maroto.New()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        m.AddRow(10, text.NewCol(12, scanner.Text(), props.Text{
            Size:  10,
            Align: align.Left,
        }))
    }

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
