package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


// https://api.chess.com/pub/player/<player>/games/archives
type MonthlyArchives struct {
	Archives []string `json:"archives"`
}

//https://api.chess.com/pub/player/<player>/games/<year>/<month>
type PlayerGames struct {
	Games []struct {
		URL         string `json:"url"`
		Pgn         string `json:"pgn"`
		TimeControl string `json:"time_control"`
		EndTime     int    `json:"end_time"`
		Rated       bool   `json:"rated"`
		Fen         string `json:"fen"`
		TimeClass   string `json:"time_class"`
		Rules       string `json:"rules"`
		White       struct {
			Rating   int    `json:"rating"`
			Result   string `json:"result"`
			ID       string `json:"@id"`
			Username string `json:"username"`
		} `json:"white"`
		Black struct {
			Rating   int    `json:"rating"`
			Result   string `json:"result"`
			ID       string `json:"@id"`
			Username string `json:"username"`
		} `json:"black"`
	} `json:"games"`
}

func main() {
	var monthlyarchives MonthlyArchives
	var mygames PlayerGames

	response, err := http.Get("https://api.chess.com/pub/player/lgbarn/games/archives")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal([]byte(data), &monthlyarchives)

	for _, archive := range monthlyarchives.Archives {
		response, err := http.Get(archive)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}
		defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(data, &mygames)
		for _, game := range mygames.Games {
			fmt.Printf("%s\n\n", string(game.Pgn))
		}
	}
}
