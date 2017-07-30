package main

import "os"
import "io"
import "io/ioutil"
import "fmt"
import "encoding/json"
import "strings"
import "path/filepath"
import "runtime"
import "github.com/shiena/ansicolor"

type Emoji struct {
    Emoji, Description, Category, Unicode_version, Ios_version string
    Aliases, Tags []string
}

func load_emojis() []Emoji {

    _, fileName, _, _ := runtime.Caller(1)
    fileName, _ = filepath.Abs(filepath.Dir(fileName))
    fileName = filepath.Join(fileName, "emojis.json")

    stream, err := os.Open(fileName)
    if err != nil {
        panic(err)
    }

    decoder := json.NewDecoder(stream)

    for {
        var emojis []Emoji
        if err := decoder.Decode(&emojis); err == io.EOF {
            break
        } else if err != nil {
            fmt.Println("Got error:", err)
        }
        return emojis
    }

    var emojis []Emoji
    return emojis

}

func hasDataToReadFromStdin()bool {
    fi, err := os.Stdin.Stat()
    return err == nil && (fi.Mode() & os.ModeCharDevice) == 0
}

func getData() string {

    bytes, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        panic(err)
    }
    return string(bytes)
}

func printUsage() {
    fmt.Println("goemoji")
    fmt.Println("-------")
    fmt.Println("Just pipe me!")
    fmt.Println("")
    fmt.Println("$ echo 'hello world :tada:' | goemoji")
    fmt.Println("hello world 🎉")
    fmt.Println("")
    fmt.Println("→ README.md for more info")
}

func main() {

    if hasDataToReadFromStdin() == false {
        printUsage()
        return
    }

    var suffix string = os.Getenv("GOEMOJI_SUFFIX") // by default, it's ""
    if len(os.Args) >= 2 {
        suffix = os.Args[1]
    }

    data := getData()

    emojis := load_emojis()
    for _, emoji := range emojis {
        for _, alias := range emoji.Aliases {
            data = strings.Replace(data, ":" + alias + ":", emoji.Emoji + suffix, -1)
        }
    }

    writer := ansicolor.NewAnsiColorWriter(os.Stdout)
    fmt.Fprint(writer, data)
}
