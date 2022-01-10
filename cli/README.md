# CLI

## How to run it

```
go run main.go
```

## Features

* teach words: collect all the words. Randomly show a word and ask the user to type the translation (do it both directions)
* teach phrases: collect all the phrases. Randomly show a phrase and ask the user to type the translation (do it both directions)
* Randomly select one of the challenges
* There is a weight system for each type of challange. The current type of challenge is selected by a random number.


## TODO


* Allow the user to adjust the weight system to pick which challenge
* Introduce a command line option that will start the CLI interface
* Add GUI, try it on various OS-es including mobile phones.

* Shall we teach the words skill-by-skill?
* Read and display the md files if they exist
* Accept words case-insenstive way
* Desrager punctuation when checking the responses.

* [Spaced repetition](https://en.wikipedia.org/wiki/Spaced_repetition)


## How it was created


```
go mod init github.com/szabgab/gogo/cli
cobra init
go get gopkg.in/yaml.v3


go get fyne.io/fyne/v2/cmd/fyne
go get fyne.io/fyne/v2
```

Error

```
/usr/bin/ld: cannot find -lXxf86vm
```

Sollution

```
sudo apt-get install libxxf86vm-dev
```

Explanations

* [CLI](https://dev.to/aurelievache/learning-go-by-examples-part-3-create-a-cli-app-in-go-1h43)
* [GUI](https://dev.to/aurelievache/learning-go-by-examples-part-7-create-a-cross-platform-gui-desktop-app-in-go-44j1)
* [Fyne](https://fyne.io/)
