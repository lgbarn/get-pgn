package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// https://api.chess.com/pub/player/<player>/games/archives
type MonthlyArchives struct {
	Archives []string `json:"archives"`
}

func main() {
	var monthlyarchives MonthlyArchives
	var CurrPlayer string
	var UseSingleFile bool
	var getLastMonth bool

	flag.StringVar(&CurrPlayer, "p", "", "Player to get pgn games")
    flag.BoolVar(&getLastMonth, "l", false, "Get last month of pgn games")
	flag.BoolVar(&UseSingleFile, "s", false, "Save to a single file")
	flag.Parse()
	if CurrPlayer == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	var currFile = CurrPlayer + ".pgn"
	response, err := http.Get("https://api.chess.com/pub/player/" + CurrPlayer + "/games/archives")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal([]byte(data), &monthlyarchives)

	if getLastMonth {
		newArchive := monthlyarchives.Archives[len(monthlyarchives.Archives)-1]
		monthlyarchives.Archives = nil
		monthlyarchives.Archives = append(monthlyarchives.Archives,newArchive[:])
	}

	for _, archive := range monthlyarchives.Archives {		
		splitArchive := strings.Split(string(archive), "/")
		year := splitArchive[7]
		month := splitArchive[8]
		fmt.Printf("Downloading games from %s/%s for %s\n", month, year, CurrPlayer)

		if !UseSingleFile {
			currFile = (CurrPlayer + "_" + year + "-" + month + ".pgn")
		}
		fmt.Println(currFile)
		f, err := os.OpenFile(currFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		response, err := http.Get(archive + "/pgn")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}
		defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		_, err = f.WriteString(string(data) + "\n")
		if err != nil {
			panic(err)
		}
	}
}
