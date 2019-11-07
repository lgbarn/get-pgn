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

// define ArchiveReader interface
type ArchiveReader interface {
	getArchiveList() []string
}

// define ArchiveConstructor interface
type ArchiveConstructer interface {
	constructArchive(data []uint8)
}

// define ArchiveConstructorReader interface
type ArchiveConstructorReader interface {
	ArchiveConstructer
	ArchiveReader
}

// monthlyArchives https://api.chess.com/pub/player/<player>/games/archives
type monthlyArchives struct {
	Archives []string `json:"archives"`
}

// getArchiveList returns a list of archive links
func (ma *monthlyArchives) getArchiveList() []string {
	return ma.Archives
}

func (ma *monthlyArchives) constructArchive(data []uint8) {
	newData := []byte(data)
	err := json.Unmarshal(newData, ma)
	if err != nil {
		panic(err)
	}
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

func checkMonthsAvailable(getLastMonth int, acr ArchiveConstructorReader) int {
	if getLastMonth > 0 {
		if getLastMonth > len(acr.getArchiveList()) {
			getLastMonth = len(acr.getArchiveList())
			fmt.Printf("Downloading %d months of archived games\n", getLastMonth)
		}
	}
	return getLastMonth
}

func writePGNFiles(getLastMonth int, acr ArchiveConstructorReader, currFile string, CurrPlayer string, UseSingleFile bool) {
	for _, archive := range getLastXof(getLastMonth, acr.getArchiveList()) {
		currFile = getFileName(archive, CurrPlayer, UseSingleFile, currFile)
		fmt.Println(currFile)
		f, err := os.OpenFile(currFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		if f != nil {
			defer f.Close()
		}
		response, err := http.Get(archive + "/pgn")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}
		if response != nil {
			defer response.Body.Close()
		}
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		_, err = f.WriteString(string(data) + "\n")
		if err != nil {
			panic(err)
		}
	}
}

func getFileName(archive string, CurrPlayer string, UseSingleFile bool, currFile string) string {
	splitArchive := strings.Split(archive, "/")
	year := splitArchive[7]
	month := splitArchive[8]
	fmt.Printf("Downloading games from %s/%s for %s\n", month, year, CurrPlayer)
	if !UseSingleFile {
		currFile = CurrPlayer + "_" + year + "-" + month + ".pgn"
	}
	return currFile
}

func getPlayerData(CurrPlayer string) []byte {
	response, err := http.Get("https://api.chess.com/pub/player/" + CurrPlayer + "/games/archives")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return data
}

func main() {
	var CurrPlayer string
	var UseSingleFile bool
	var getLastMonth int

	var acr ArchiveConstructorReader = &monthlyArchives{}

	flag.StringVar(&CurrPlayer, "p", "", "Player to get pgn games")
	flag.IntVar(&getLastMonth, "l", 9999, "Get last month of pgn games")
	flag.BoolVar(&UseSingleFile, "s", false, "Save to a single file")
	flag.Parse()
	if CurrPlayer == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	var currFile = CurrPlayer + ".pgn"
	playerData := getPlayerData(CurrPlayer)
	acr.constructArchive(playerData)
	getLastMonth = checkMonthsAvailable(getLastMonth, acr)
	writePGNFiles(getLastMonth, acr, currFile, CurrPlayer, UseSingleFile)
}
