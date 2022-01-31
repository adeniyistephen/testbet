package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/adeniyistephen/testbet/business"
)

func main() {
	go StartPolling()
	business.SaveAllSport()
	odds := business.GetInPlayOddsUk("h2h", "uk")
	fmt.Println("Odds and teams Uk:", odds)

	http.HandleFunc("/testbet", WelcomeToTestbet)
	http.ListenAndServe(":8080", nil)
}

func WelcomeToTestbet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to testbet")
}

func StartPolling() {
	for {
		time.Sleep(1 * time.Hour)
		go business.SaveUpcomingGames()
	}
}
