package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"net/http"
	"os"
)

// https://api.chess.com/pub/player/<player>
type Player struct {
	Avatar     string `json:"avatar"`
	PlayerID   int    `json:"player_id"`
	ID         string `json:"@id"`
	URL        string `json:"url"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Followers  int    `json:"followers"`
	Country    string `json:"country"`
	Location   string `json:"location"`
	LastOnline int    `json:"last_online"`
	Joined     int    `json:"joined"`
	Status     string `json:"status"`
	IsStreamer bool   `json:"is_streamer"`
}

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

type URLStatus struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}



func main() {
	//var monthlyarchives MonthlyArchives
	var mygames PlayerGames
	
	/* Commented temporaily while working on section below
	response, err := http.Get("https://api.chess.com/pub/player/lgbarn/games/archives")
	 if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} 
	
	defer response.Body.Close()
	
	data, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(data))
	json.Unmarshal([]byte(data), &monthlyarchives)

	for _, archive := range monthlyarchives.Archives {
	  fmt.Printf("%s\n", string(archive))
	}
*/
	jsonFile, err := os.Open("games_2010_05.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened games_2010_05.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &mygames)

	for _, game := range mygames.Games {
		fmt.Printf("%s\n\n", string(game.Pgn))
	}
}