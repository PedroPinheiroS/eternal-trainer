package tibiadata

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PedroPinheiroS/eternal-trainer/env"
)

type GetWorld struct {
	Worlds struct {
		World struct {
			Name                string    `json:"name"`
			Status              string    `json:"status"`
			PlayersOnline       int       `json:"players_online"`
			RecordPlayers       int       `json:"record_players"`
			RecordDate          time.Time `json:"record_date"`
			CreationDate        string    `json:"creation_date"`
			Location            string    `json:"location"`
			PvpType             string    `json:"pvp_type"`
			PremiumOnly         bool      `json:"premium_only"`
			TransferType        string    `json:"transfer_type"`
			WorldQuestTitles    []string  `json:"world_quest_titles"`
			BattleyeProtected   bool      `json:"battleye_protected"`
			BattleyeDate        string    `json:"battleye_date"`
			GameWorldType       string    `json:"game_world_type"`
			TournamentWorldType string    `json:"tournament_world_type"`
			OnlinePlayers       []struct {
				Name     string `json:"name"`
				Level    int    `json:"level"`
				Vocation string `json:"vocation"`
			} `json:"online_players"`
		} `json:"world"`
	} `json:"worlds"`
	Information struct {
		APIVersion int       `json:"api_version"`
		Timestamp  time.Time `json:"timestamp"`
	} `json:"information"`
}

func GetOnline() (bool, error) {
	requestURL := fmt.Sprintf("https://api.tibiadata.com/v3/world/%s", env.Mundo)

	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	defer res.Body.Close()

	var getWorld GetWorld

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err)
	}

	err = json.Unmarshal(bodyBytes, &getWorld)
	if err != nil {
		fmt.Print(err)
	}

	online := false
	for _, player := range getWorld.Worlds.World.OnlinePlayers {
		if player.Name == env.Personagem {
			online = true
			log.Println("Online")
			break
		}
	}

	return online, nil
}
