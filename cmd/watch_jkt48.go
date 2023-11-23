package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/khangle880/showroom/models"
	"github.com/khangle880/showroom/utils"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var watchJkt48Live = &cobra.Command{
	Use:   "jkt48",
	Short: "Watch JKT48 member live streams",
	Run: func(cmd *cobra.Command, args []string) {
		watchJkt48(cmd)
	},
}

func watchJkt48(cmd *cobra.Command) {
	roomChan := make(chan *[]models.Room)

	utils.LogInfo("💫 step 1 of 2 | Getting all JKT48 members live status...")
	fmt.Println()
	progressbar1 := progressbar.NewOptions(len(utils.AddedRooms), progressbar.OptionSetWidth(35), progressbar.OptionOnCompletion(func() {
		fmt.Fprint(cmd.OutOrStdout(), "\n")
	}))

	go utils.GetJKT48Rooms(progressbar1, roomChan)
	rooms := <-roomChan

	fmt.Println()
	utils.LogInfo("💫 step 2 of 2 | Seaching their streaming url's")
	fmt.Println()
	progressbar2 := progressbar.NewOptions(len(*rooms), progressbar.OptionSetWidth(35), progressbar.OptionOnCompletion(func() {
		fmt.Fprint(cmd.OutOrStdout(), "\n")
	}))

	go func() {
		for i := 0; i < progressbar2.GetMax(); i++ {
			if i == progressbar2.GetMax()-1 && !progressbar2.IsFinished() {
				progressbar2.ChangeMax(progressbar2.GetMax() + 10)
			}
			progressbar2.Add(1)
			time.Sleep(1 * time.Second)
		}
	}()

	var liveRooms []models.LiveRoom
	for _, data := range *rooms {
		if data.IsLive {
			urlList := utils.GetStreamingUrlByRoomId(data.Id)
			liveRooms = append(liveRooms, models.LiveRoom{
				RoomID:           data.Id,
				MainName:         data.Name,
				StreamingURLList: *urlList,
			})
		}
		progressbar2.Finish()
	}

	if len(liveRooms) == 0 {
		fmt.Println()
		utils.LogInfo("🤭 Oops, seems like there's no active rooms right now")
		os.Exit(0)
	}

	selectedRoom := utils.PromptGetSelect(&liveRooms)
	if len(selectedRoom.StreamingURLList) != 0 {
		streamUrl := utils.PromptSelectQuality(&selectedRoom)
		utils.StartVLCService(streamUrl)
	} else {
		utils.LogError(fmt.Errorf("⚠️ Error: Not found streaming url"))
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(watchJkt48Live)
}
