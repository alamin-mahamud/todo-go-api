package cmd

import (
	"fmt"
	"github.com/alaminmahamud/todo-app/app"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "Todo App",
	Short: "Todo App now only supports Task List",
	Long: `A fast and flexible static task list built with 
	- Cobra
	- Echo Framework
	- Clean Architecture
	`,
	Run: func(cmd *cobra.Command, args []string) {
		app.Run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}