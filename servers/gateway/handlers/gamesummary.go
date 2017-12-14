package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	//"golang.org/x/net/html"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	//"strconv"
	"log"
	"strings"
)

const unityURL = "https://unity3d.com/showcase/gallery"

var Games []*GameSummary

type GameSummary struct {
	GameURL     string `json:"gameUrl,omitempty"`
	Title       string `json:"title,omitempty"`
	Developer   string `json:"developer,omitempty"`
	Genre       string `json:"genre,omitempty"`
	Description string `json:"description,omitempty"`
	ImageURL    string `json:"imageUrl,omitempty"`
}

func GameSummaryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(headerContentType, contentTypeJSON)
	w.Header().Add(headerAccessControlAllowOrigin, "*")

	htmlStream, err := fetchHTML(unityURL)
	if err != nil {
		http.Error(w, "error fetching html", http.StatusBadRequest)
		return
	}
	extract, err := extractGames()
	// extract, err := extractGames(unityURL, htmlStream)
	if err != nil {
		http.Error(w, "error extracting game summary", http.StatusBadRequest)
		return
	}

	defer htmlStream.Close()

	json.NewEncoder(w).Encode(extract)
}

func fetchHTML(pageURL string) (io.ReadCloser, error) {
	resp, err := http.Get(pageURL)
	if err != nil {
		return nil, errors.New("error fecting URL")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("response status code >= 400")
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/html") {
		return nil, errors.New("response content type was not text/html")
	}
	return resp.Body, nil
}
func extractGames() ([]*GameSummary, error) {
	var gameSlice []*GameSummary

	doc, err := goquery.NewDocument("https://unity3d.com/showcase/gallery")
	if err != nil {
		log.Fatal(err)
	}
	var count = 0
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
		count++
		gameSlice = append(gameSlice, game)
	})
	fmt.Println(len(gameSlice))
	return gameSlice, nil
}
