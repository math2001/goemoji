package main

import "os"
import "io"
import "io/ioutil"
import "fmt"
import "encoding/json"
import "strings"


type Emoji struct {
    Emoji, Description, Category, Unicode_version, Ios_version string
    Aliases, Tags []string
}

func load_emojis() []Emoji {
    stream, err := os.Open("emojis.json")
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
    stat, _ := os.Stdin.Stat()
    if (stat.Mode() & os.ModeCharDevice) == 0 {
        return true
    }
    return false
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
    print("goEmoji")
    print("-------")
    print("Just pip me!")
    print("")
    print("echo 'hello world :tada:' | emo")
}

func main() {

    if hasDataToReadFromStdin() == false {
        printUsage()
        return
    }

    data := getData()

    emojis := load_emojis()
    for _, emoji := range emojis {
        for _, alias := range emoji.Aliases {
            data = strings.Replace(data, ":" + alias + ":", emoji.Emoji, -1)
        }
    }

    print(data)

}
