package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
	response, err := http.Get("https://api.chess.com/pub/player/lgbarn")
	if err != nil {
		fmt.Print("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}