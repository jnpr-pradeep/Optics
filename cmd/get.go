/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	optics "optics/pkg/utils"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		sku, _ := cmd.Flags().GetString("sku")
		speed, _ := cmd.Flags().GetString("speed")
		cable, _ := cmd.Flags().GetString("cable")
		fmt.Println("==========================================================")
		fmt.Printf("Optics for sku=%v, speed=%v, cable=%v\n", sku, speed, cable)
		fmt.Println("==========================================================")

		// Get the new optics processor
		o := optics.NewOpticsProcessor()
		// Set the parameters
		if sku != "" {
			o.SetSKU(sku)
		}
		if speed != "" {
			o.SetSpeed(speed)
		}
		if cable != "" {
			o.SetCableType(cable)
		}

		// Get the Optics
		o.GetOpticsWithCableTypeAndSpeed()

		fmt.Println(o.GetSelectedOptics())
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringP("sku", "s", "", "Interested Device Model")
	getCmd.Flags().StringP("speed", "v", "", "Speed of the Port")
	getCmd.Flags().StringP("cable", "c", "", "Cable Type")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
