package main

import "os"
import "io"
import "fmt"
import "encoding/json"
import "strings"

// bytes, err := ioutil.ReadAll(os.Stdin)
// if err != nil {
//     fmt.Println("Error when reading stdin", err)
//     return
// }

type Emoji struct {
    Emoji, Description, Category, Unicode_version, Ios_version string
    Aliases, Tags []string
}

func load_emojis() []Emoji {
    stream, err := os.Open("emoji.json")
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

func getData() string{
    return "hello :+1: world :fr:"
}

func main() {

    data := getData()

    emojis := load_emojis()
    for _, emoji := range emojis {
        for _, alias := range emoji.Aliases {
            data = strings.Replace(data, ":" + alias + ":", emoji.Emoji, -1)
        }
    }

    print(data)

}
