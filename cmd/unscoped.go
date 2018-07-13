package cmd

import (
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var unscopedCmd = &cobra.Command{
	Use:   "unscoped",
	Short: "Unscoped Products",
	Run: func(cmd *cobra.Command, args []string) {

		products, err := apiService.UnscopedProducts()
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Version", "Arch"})

		for _, product := range products {
			table.Append([]string{product.Name, product.Version, product.Arch})
		}

		table.Render()
	},
}
