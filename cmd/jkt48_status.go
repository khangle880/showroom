package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/khangle880/showroom/models"
	"github.com/khangle880/showroom/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var jkt48StatusCmd = &cobra.Command{
	Use:   "jkt48status",
	Short: "Show all JKT48 members live status",
	Run: func(cmd *cobra.Command, args []string) {
		roomChan := make(chan *[]models.Room)

		utils.LogInfo("💫 Getting all JKT48 members live status...")
		fmt.Println()
		progressBar := progressbar.NewOptions(len(utils.AddedRooms), progressbar.OptionSetWidth(35), progressbar.OptionOnCompletion(func() {
			fmt.Fprint(cmd.OutOrStdout(), "\n")
		}))

		go utils.GetJKT48Rooms(progressBar, roomChan)

		rooms := <-roomChan
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Status"})
		table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_CENTER})

		for _, data := range *rooms {
			if data.IsLive {
				table.Append([]string{data.Name, color.GreenString("ONLINE")})
			} else {
				table.Append([]string{data.Name, color.RedString("OFFLINE")})
			}
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(jkt48StatusCmd)
}
