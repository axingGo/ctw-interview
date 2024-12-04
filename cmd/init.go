package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "airbnb-cli",
	Short: "",
	Long:  "",
}

func CmdExec() error {
	return RootCmd.Execute()
}
