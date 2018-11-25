package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	var CurrPlayer string

	flag.StringVar(&CurrPlayer, "p", "", "Player to get pgn games")
	flag.Parse()
	if CurrPlayer == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	response, err := http.Get("https://lichess.org/api/games/user/" + CurrPlayer + "")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	response.Header.Set("Authorization", "Bearer xYAsDILHMwOiAJfT")
	defer response.Body.Close()
	reader := bufio.NewReader(response.Body)
	var currFile = CurrPlayer + ".pgn"
	fmt.Println(currFile)
	f, err := os.OpenFile(currFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for {
		line, _ := reader.ReadString('\n')
		_, err = f.WriteString(string(line) + "")
		if err != nil {
			panic(err)
		}
	}
}
