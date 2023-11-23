package cmd

import (
	"fmt"
	"os"

	"github.com/khangle880/showroom/models"
	"github.com/khangle880/showroom/utils"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var watchLive = &cobra.Command{
	Use:   "watch",
	Short: "Watch onlive streams",
	Run: func(cmd *cobra.Command, args []string) {
		watchLiveRun(cmd)
	},
}

func watchLiveRun(cmd *cobra.Command) {
	roomsChan := make(chan *[]models.LiveRoom)

	utils.LogInfo("💫 Getting all live status...")
	fmt.Fprint(cmd.OutOrStderr(), "\n")
	progressbar := progressbar.NewOptions(88, progressbar.OptionSetWidth(35), progressbar.OptionOnCompletion(func() {
		fmt.Println()
	}))

	go utils.GetAllOnliveRooms(progressbar, roomsChan)
	rooms := <-roomsChan

	selectedRoom := utils.PromptGetSelect(rooms)
	if len(selectedRoom.StreamingURLList) != 0 {
		streamUrl := utils.PromptSelectQuality(&selectedRoom)
		utils.StartVLCService(streamUrl)
	} else {
		utils.LogError(fmt.Errorf("⚠️ Error: Not found streaming url"))
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(watchLive)
}
