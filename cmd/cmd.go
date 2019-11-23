package cmd

import (
	"github.com/spf13/cobra"
)

var dowMp3Cmd = &cobra.Command{
	Use:   "down-song",
	Short: "Download music song with specify url",
	Long:  `Download music song with specify url`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var dowMp4Cmd = &cobra.Command{
	Use:   "down-video",
	Short: "Download music video with specify url",
	Long:  `Download music video with specify url`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	RootCmd.AddCommand(dowMp3Cmd)
	RootCmd.AddCommand(dowMp4Cmd)
}
