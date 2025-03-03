package sub2

import (
	"fmt"

	"github.com/spf13/cobra"
)

var leadACmd = &cobra.Command{
	Use:   "leafA",
	Short: "A brief description of your command",
	Run:   func(cmd *cobra.Command, args []string) { fmt.Println("sub2 leafA called") },
}
