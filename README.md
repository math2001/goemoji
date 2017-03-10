# GoEmoji

I spent some time looking for something really simple: I wanted to be able to do this:

```bash
$ echo 'Hello world :tada:' | someprogram
Hello world 🎉
```

I didn't find one quickly enough, and as I want to learn go, I thought that doing it myself could be
a good task (and it was :smile: I learned a :shit:-load of new stuff :tada:)

To get it working on your machine, just grab the latest `.exe` from the [tags][], and add an alias
to it or add it in your `PATH`.

```bash
# .bashrc
alias goemoji="path/to/goemoji.exe $*"
```

*Note: the* `$*` *at then end allows to give arguments*

Now, you can do

```
$ echo 'I :heart: code' | goemoji
I ❤ code
```

You can also build it yourself, it's simple:

```
$ go build goemoji.go
```

:tada:

### Add a suffix

I'm on Windows, and I'm using Hyper. At this time, Hyper has a bug, emoji take too much room, which
"eat" the following char. The solution to prevent this is to add a space just after (so that it
eats the space).

So, you can specify a suffix by passing it as an argument:

```
$ git log | goemoji " "
```

Or by setting it as an environment variable:

```
$ export GOEMOJI_SUFFIX=" "
$ git log | goemoji
```

### `--color=always`

Plenty of program add color to their output *only if the destination is a terminal*. So, when you
pipe it, it doesn't use colors (so you get emoji, but no color :disappointed:). It's the case for
`git log`.

Fortunatly, there's often an option to *oblige* colored output. For `git log`, it's `--color=always`

```
$ git log --color="always" --oneline --graph --decorate -10 --all | goemoji
```

:wink:

That's it! Hope you enjoy it (don't forget to :star: this repo if it did :smile:)

[tags]: https://github.com/math2001/goemoji/releases/latest
