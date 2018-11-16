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

// monthlyArchives https://api.chess.com/pub/player/<player>/games/archives
type monthlyArchives struct {
	Archives []string `json:"archives"`
}

// getArchiveList returns a list of achive links
func (MonthlyArchives *monthlyArchives) getArchiveList() []string {
	var archiveList []string
	for _, archive := range MonthlyArchives.Archives {
		archiveList = append(archiveList, archive)
	}
	return archiveList
}

// reverseArchiveList returns archive links in reverse order
func (MonthlyArchives *monthlyArchives) reverseArchiveList() []string {
	var archiveList []string
	for _, archive := range MonthlyArchives.Archives {
		archiveList = append(archiveList, archive)
	}
	for i := len(archiveList)/2 - 1; i >= 0; i-- {
		opp := len(archiveList) - 1 - i
		archiveList[i], archiveList[opp] = archiveList[opp], archiveList[i]
	}
	return archiveList
}

// getLastXof get X number months
func getLastXof(number int, inList []string) []string {
	var outList []string
	for i := len(inList)/2 - 1; i >= 0; i-- {
		opp := len(inList) - 1 - i
		inList[i], inList[opp] = inList[opp], inList[i]
	}
	for i := 0; i < number; i++ {
		outList = append(outList, inList[i])
	}
	for i := len(outList)/2 - 1; i >= 0; i-- {
		opp := len(outList) - 1 - i
		outList[i], outList[opp] = outList[opp], outList[i]
	}
	return outList
}

func main() {
	var monthlyarchives monthlyArchives
	var CurrPlayer string
	var UseSingleFile bool
	var getLastMonth int

	flag.StringVar(&CurrPlayer, "p", "", "Player to get pgn games")
	flag.IntVar(&getLastMonth, "l", 0, "Get last month of pgn games")
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

	if getLastMonth > 0 {
		if getLastMonth > len(monthlyarchives.Archives) {
			getLastMonth = len(monthlyarchives.Archives)
			fmt.Printf("Only %d Number of months are available, returning %d number of months instead\n", getLastMonth, getLastMonth)
		}
		monthlyarchives.Archives = getLastXof(getLastMonth, monthlyarchives.getArchiveList())
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
