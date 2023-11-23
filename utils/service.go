package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/khangle880/showroom/models"
	"github.com/schollz/progressbar/v3"
)

func GetAllOnliveRooms(bar *progressbar.ProgressBar, roomsChan chan<- *[]models.LiveRoom) {
	var rooms []models.LiveRoom

	res, err := http.Get(OnlivesApiURL)
	if err != nil {
		LogError(err)
		os.Exit(1)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		LogError(fmt.Errorf("⚠️ Error: %s", err))
		os.Exit(1)
	}

	type ApiResponse struct {
		Onlives []models.Onlives `json:"onlives"`
	}
	var decodedData ApiResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&decodedData); err != nil {
		LogError(err)
	}

	for _, data := range decodedData.Onlives {
		rooms = append(rooms, data.Lives...)
	}

	bar.Finish()

	roomsChan <- &rooms
}

func GetJKT48Rooms(bar *progressbar.ProgressBar, result chan<- *[]models.Room) {
	var jkt48 []models.Room

	res, err := http.Get(AKB48RoomURL)
	if err != nil {
		LogError(err)
		os.Exit(1)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		LogError(fmt.Errorf("⚠️ Error: %s", res.Status))
		os.Exit(1)
	}

	var decodedData []models.Room
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&decodedData); err != nil {
		LogError(err)
	}

	for _, room := range decodedData {
		if strings.Contains(room.Name, "JKT48") {
			jkt48 = append(jkt48, room)
		}
	}

	bar.ChangeMax(bar.GetMax() + len(jkt48))
	bar.Add(len(jkt48))

	for _, data := range AddedRooms {
		res, err := http.Get(fmt.Sprintf("%s/profile?room_id=%d", RoomApiURL, data.RoomId))
		if err != nil {
			LogError(err)
		}
		defer res.Body.Close()
		var decodedData models.LiveRoom
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&decodedData); err != nil {
			LogError(err)
		}

		jkt48 = append(jkt48, models.Room{
			Id:          decodedData.RoomID,
			Name:        decodedData.MainName,
			URLKey:      decodedData.RoomURLKey,
			ImageURL:    decodedData.ImageSquare,
			FollowerNum: decodedData.FollowerNum,
			IsLive:      decodedData.IsLive,
		})
		bar.Add(1)
	}

	result <- &jkt48
}

func GetStreamingUrlByRoomId(roomId int) *[]models.StreamingURL {
	res, err := http.Get(fmt.Sprintf("%s/streaming_url?room_id=%d", LiveApiURL, roomId))
	if err != nil {
		LogError(err)
		os.Exit(1)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		LogError(fmt.Errorf("⚠️ Error: %s", res.Status))
		os.Exit(1)
	}

	var decodedData struct {
		StreamingUrlList []models.StreamingURL `json:"streaming_url_list"`
	}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&decodedData); err != nil {
		LogError(err)
	}

	return &decodedData.StreamingUrlList
}
