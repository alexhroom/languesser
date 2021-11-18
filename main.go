package main

import (
	"flag"
	"log"

	src "github.com/alexhroom/languesser/src"
)

func main() {
	path := flag.String("path", "./file.txt", "In learn mode, the directory to learn from. In guess mode, the file to guess the language of.")
	mode := flag.String("mode", "guess", "The mode of the script; whether you are giving it input to learn, or a file to guess from.")
	language := flag.String("lang", "NA", "Only used in learn mode; the language that the file to learn from is in, in all lowercase.")

	flag.Parse()

	if *mode == "learn" {
		if *language == "NA" {
			log.Fatal("You must specify a language in learn mode.")
		}
		src.Learn(*path, *language)
	}
	if *mode == "guess" {
		// not yet implemented
		//src.Guess(*path)
	}
}
