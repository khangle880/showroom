package utils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/khangle880/showroom/models"
	"github.com/manifoldco/promptui"
)

func PromptGetSelect(rooms *[]models.LiveRoom) models.LiveRoom {
	var items []string
	for _, room := range *rooms {
		label := fmt.Sprintf("%s (ID: %d) (Viewer: %d)", room.MainName, room.RoomID, room.ViewNum)
		items = append(items, label)
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . | cyan }}?",
		Active:   "\U0001F449 {{ . | green | underline }}",
		Inactive: "  {{ . | cyan }}",
		Selected: "\U0001F449 {{ . | green | underline }}",
	}

	prompt := promptui.Select{
		Label: "Select Room",
		Items: items,
		// Size:      10,
		Templates: templates,
	}

	i, _, err := prompt.Run()
	if err != nil {
		LogError(err)
		os.Exit(1)
	}

	selectedRoom := (*rooms)[i]

	return selectedRoom
}
func PromptSelectQuality(room *models.LiveRoom) string {
	urlList := room.StreamingURLList
	template := &promptui.SelectTemplates{
		Label:    "{{ . | cyan }}?",
		Active:   "\U0001F449 {{ . | green | underline }}",
		Inactive: "  {{ . | cyan }}",
	}

	var labels []string
	for _, label := range urlList {
		labels = append(labels, label.Label)
	}

	prompt := promptui.Select{
		Label:     "Select Quality",
		Items:     append(labels, "Get more"),
		Templates: template,
	}

	i, _, err := prompt.Run()
	if err != nil {
		LogError(err)
		os.Exit(1)
	}

	if i == len(labels) {
		room.StreamingURLList = *GetStreamingUrlByRoomId(room.RoomID)
		return PromptSelectQuality(room)
	}

	return (urlList)[i].URL
}

func StartVLCService(url string) {
	cmd := exec.Command("/Applications/VLC.app/Contents/MacOS/VLC", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		LogError(err)
		os.Exit(1)
	}
}
