package business

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/adeniyistephen/testbet/database"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var SportCollection *mongo.Collection = database.OpenCollection(database.Client, "sports")
var UpcomingCollection *mongo.Collection = database.OpenCollection(database.Client, "upcominggames")

func SaveAllSport() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// getting env variables API_KEY
	apiKey := os.ExpandEnv("https://api.the-odds-api.com/v4/sports/?apiKey=$API_KEY&all=true")
	response, err := http.Get(apiKey)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var sport []database.Sport
	var sp database.Sport
	json.Unmarshal(responseData, &sport)

	for i := 0; i < len(sport); i++ {
		sp.Active = sport[i].Active
		sp.Key = sport[i].Key
		sp.Title = sport[i].Title
		sp.Description = sport[i].Description
		sp.Group = sport[i].Group
		sp.Has_Outrights = sport[i].Has_Outrights

		insertResult, err := SportCollection.InsertOne(context.Background(), sp)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted Document:", sp, "InsertID: ", insertResult.InsertedID)
	}
}

func SaveUpcomingGames() database.In_Play {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//getting env variable API_KEY
	apiKey := os.ExpandEnv("https://api.the-odds-api.com/v4/sports/upcoming/odds/?regions=us,uk,eu,au&apiKey=$API_KEY")
	response, err := http.Get(apiKey)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var upcomingGames []database.UpcomingGames
	var ug database.UpcomingGames
	json.Unmarshal(responseData, &upcomingGames)

	for i := 0; i < len(upcomingGames); i++ {
		ug.Id = upcomingGames[i].Id
		ug.Sport_Key = upcomingGames[i].Sport_Key
		ug.Sport_Title = upcomingGames[i].Sport_Title
		ug.Commence_Time = upcomingGames[i].Commence_Time
		ug.Away_Team = upcomingGames[i].Away_Team
		ug.Home_Team = upcomingGames[i].Home_Team
		ug.Bookmakers = upcomingGames[i].Bookmakers
		insertResult, err := UpcomingCollection.InsertOne(context.Background(), ug)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted Document:", ug, "InsertID: ", insertResult.InsertedID)
	}

	var ip database.In_Play
	ip.InPlay = ug
	return ip
}

func GetInPlayOddsUk(markets, region string) []database.Outcome {
	//getting odds h2h odds for the UK market
	h2h_Uk := os.ExpandEnv("https://api.the-odds-api.com/v4/sports/upcoming/odds/?regions=" + region + "&markets=" + markets + "&apiKey=$API_KEY")
	response, err := http.Get(h2h_Uk)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var inPlayGames []database.UpcomingGames
	var odds []database.Outcome
	json.Unmarshal(responseData, &inPlayGames)

	for i := 0; i < len(inPlayGames); i++ {
		book := inPlayGames[i].Bookmakers
		for i := 0; i < len(book); i++ {
			market := book[i].Markets
			for i := 0; i < len(market); i++ {
				outcome := market[i].Outcomes
				odds = append(odds, outcome...)
			}
		}
	}
	return odds
}
