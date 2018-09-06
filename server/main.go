package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jzelinskie/geddit"

	"github.com/gorilla/mux"
)

var o *geddit.OAuthSession

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}

	mux := mux.NewRouter()
	mux.HandleFunc("/test", TestHandler)
	mux.HandleFunc("/login", RedditLoginHandler)
	log.Printf("server is listening at http://%s...", addr)
	log.Fatal(http.ListenAndServe(":4000", mux))
}

// TestHandler is a handler
func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(HeaderContentType, ContentTypeText)
	w.Header().Add(HeaderAccessControlAllowHeaders, "*")
	w.Write([]byte("Hello Test WOW!"))
}

//RedditLoginHandler is a handler
func RedditLoginHandler(w http.ResponseWriter, r *http.Request) {
	id := reqEnv("CLIENT_ID")
	secret := reqEnv("CLIENT_SECRET")
	redditID := reqEnv("REDDIT_ID")
	redditPass := reqEnv("REDDIT_PASS")
	o, _ = geddit.NewOAuthSession(
		id,
		secret,
		"Testing oauth bot by u/BirdUpBrotendoo v0.1 see source https://github.com/ask710/reddit-essentials",
		"http://127.0.0.1:4000/login",
	)

	// Create new auth token for confidential clients (personal scripts/apps).
	err := o.LoginAuth(redditID, redditPass)
	if err != nil {
		log.Printf("Error getting new auth: %v", err)
	}
	subOpts := geddit.ListingOptions{
		Limit: 10,
	}
	// submissions, _ := o.Frontpage(geddit.DefaultPopularity, subOpts)
	// for _, s := range submissions {
	// 	w.Write([]byte(fmt.Sprintf("Title: %s\nAuthor: %s\n", s.Title, s.Author)))
	// }
	hhh, _ := o.SubredditSubmissions("nba", geddit.DefaultPopularity, subOpts)
	for _, s := range hhh {

		// if strings.Contains(s.Title, "[FRESH]") {
		w.Write([]byte(fmt.Sprintf("Title: %s\nAuthor: %s\nDate: %v\nClicked: %v\n", s.Title, s.Author, s.DateCreated, s.WasClicked)))

		// }
	}
}
func reqEnv(name string) string {
	val := os.Getenv(name)
	if len(val) == 0 {
		//Fatal?
		log.Fatalf("Please set %s variable", name)
		os.Exit(1)
	}
	return val
}
