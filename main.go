package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// https://api.chess.com/pub/player/<player>/games/archives
type MonthlyArchives struct {
	Archives []string `json:"archives"`
}

func main() {
	var monthlyarchives MonthlyArchives
	var CurrPlayer string

	flag.StringVar(&CurrPlayer, "player", "", "Player to get pgn games")
	flag.Parse()
	if CurrPlayer == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}
	response, err := http.Get("https://api.chess.com/pub/player/" + CurrPlayer + "/games/archives")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal([]byte(data), &monthlyarchives)
	f, err := os.OpenFile(CurrPlayer+".pgn", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, archive := range monthlyarchives.Archives {
		response, err := http.Get(archive + "/pgn")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}
		defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		if _, err = f.WriteString(string(data) + "\n"); err != nil {
			panic(err)
		}
	}
}
