package handlers

import (
	"encoding/json"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// GameSummary represents the properties for a game
type GameSummary struct {
	GameURL     string `json:"gameUrl,omitempty"`
	Title       string `json:"title,omitempty"`
	Developer   string `json:"developer,omitempty"`
	Genre       string `json:"genre,omitempty"`
	Description string `json:"description,omitempty"`
	ImageURL    string `json:"imageUrl,omitempty"`
}

// storage of all the games from the website
type GameList struct {
	list []*GameSummary
}

var gl = &GameList{}

// Handles requests for the game summary API.
// This adds a cookie when the client first makes a request,
// then responds with a JSON-encoded Game Summary struct containing
// the game summary data.
func GameSummaryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	_, err := r.Cookie("user")
	cookie := &http.Cookie{
		Name:    "user",
		Expires: time.Now().AddDate(0, 0, 1),
		Value:   strconv.FormatInt(time.Now().Unix(), 10),
	}
	http.SetCookie(w, cookie)

	var extract *GameSummary
	if err != nil {
		// get all games for new user
		gl.list = extractGames()

		extract, err = gl.randomGame()
		if err != nil {
			http.Error(w, "error getting game information", http.StatusBadRequest)
			return
		}
	}

	extract, err = gl.randomGame()
	if err != nil {
		http.Error(w, "You viewed all the games! Refresh to start over.", http.StatusBadRequest)
		gl.list = extractGames()
	}
	json.NewEncoder(w).Encode(extract)
}

// Extracts the games available on in the gallery and saves it
// to a slice and returns it.
func extractGames() []*GameSummary {
	var gameSlice []*GameSummary

	doc, err := goquery.NewDocument("https://unity3d.com/showcase/gallery")
	if err != nil {
		log.Fatal(err)
	}

	// scrape HTML for video game information
	doc.Find(".game").Each(func(index int, item *goquery.Selection) {
		game := &GameSummary{}
		game.Title = item.Find(".title").Text()
		game.Developer = item.Find(".developer").Text()
		game.Description = item.Find(".description").Find("p").Text()
		url, _ := item.Find(".description").Find("a").Attr("href")
		game.GameURL = url
		game.Genre = item.Find(".genres").Text()
		image, _ := item.Find(".ic").Attr("src")
		game.ImageURL = image
		gameSlice = append(gameSlice, game)
	})
	return gameSlice
}

// Generates a random game in the game slice.
// Deletes from slice to prevent user from seeing the same game.
func (games *GameList) randomGame() (*GameSummary, error) {
	rand.Seed(time.Now().Unix())
	if len(games.list) == 0 {
		return nil, errors.New("Out of games")
	}
	// random number to choose game
	n := rand.Int() % len(games.list)
	tempGame := games.list[n]

	//delete game from slice by swapping to end and
	//return n-1 elements
	games.list[len(games.list)-1], games.list[n] = games.list[n], games.list[len(games.list)-1]
	games.list = games.list[:len(games.list)-1]

	return tempGame, nil
}
