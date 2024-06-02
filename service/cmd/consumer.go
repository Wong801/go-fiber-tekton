/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// consumerCmd represents the consumer command
var consumerCmd = &cobra.Command{
	Use:   "consumer NAME",
	Short: "Start consumer.",
	Long: `Start consumer with name of the consumer as argument.`,
	Args: cobra.ExactArgs(1),
	
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("consumer called")
	},
}

func init() {
	serveCmd.AddCommand(consumerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// consumerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// consumerCmd.Flags().StringP("name", "n", "", "Specify the name of the consumer")
}
