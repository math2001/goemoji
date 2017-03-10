package main

import "os"
import "io"
import "io/ioutil"
import "fmt"
import "encoding/json"
import "strings"
import "path/filepath"
import "runtime"


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
        fmt.Println("Error when reading stdin", err)
        return ""

    }
    return string(bytes)
}

func printUsage() {
    fmt.Println("goEmoji")
    fmt.Println("-------")
    fmt.Println("Just pip me!")
    fmt.Println("")
    fmt.Println("$ echo 'hello world :tada:' | goemoji")
    fmt.Println("")
    fmt.Println("→ README.md for more info")
}

func main() {

    if hasDataToReadFromStdin() == false {
        printUsage()
        return
    }
    var suffix string
    if len(os.Args) >= 2 {
        suffix = os.Args[1]
    } else {
        suffix = ""
    }

    data := getData()

    emojis := load_emojis()
    for _, emoji := range emojis {
        for _, alias := range emoji.Aliases {
            data = strings.Replace(data, ":" + alias + ":", emoji.Emoji + suffix, -1)
        }
    }

    print(data)

}
