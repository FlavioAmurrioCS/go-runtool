package cmd

import (
	"fmt"

	"github.com/FlavioAmurrioCS/lazystream"
	"github.com/spf13/cobra"
)

var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "Filter links based on OS and Arch",
	Run: func(cmd *cobra.Command, args []string) {
		lazystream.
			FromStdin().
			Append(args...).
			ForEach(func(line string) {
				fmt.Println(line)
			})
	},
}
