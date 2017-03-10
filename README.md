# GoEmoji

I spent some time looking for something really simple: I wanted to be able to do this:

```bash
echo 'Hello world :tada:' | someprogram
```

I didn't find one quickly enough, and as I want to learn go, I thought that doing it myself could be a good task (and it was :smile: I learned a :shit:-load of new stuff :tada:)

To get it working on your machine, just grab the latest `.exe` from the [tags][], and add an alias to it or add it in your `PATH`.

```bash
# .bashrc
alias goemoji="path/to/goemoji"
```

Now, you can do

```
$ echo 'I :heart: code' | goemoji
```

You can also build it yourself, it's simple:

```
$ go build goemoji.go
```

:tada:

That's it! Hope you enjoy it (don't forget to :star: this repo if it did :smile:)
