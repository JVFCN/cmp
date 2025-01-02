package cmd

import (
	"CLMusicPlayer/MusicList"
	"CLMusicPlayer/Play"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "cmp",
	Short: "CLMusicPlayer让您在终端中也能享受音乐",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version: 1.0.1")
		fmt.Println(cmd.Help())
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(Play.Play)
	Play.Play.Flags().Int16P("list", "l", -1, "Choose a playlist to play")
	Play.Play.Flags().StringP("path", "p", "", "Choose a file to play")
	Play.Play.Args = cobra.NoArgs

	rootCmd.AddCommand(MusicList.Add)
	rootCmd.AddCommand(MusicList.List)
	MusicList.List.Flags().Int16P("number", "n", -1, "Display information about a playlist")
	MusicList.List.Args = cobra.NoArgs
}
