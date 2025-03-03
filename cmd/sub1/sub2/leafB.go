package sub2

import (
	"fmt"

	"github.com/spf13/cobra"
)

var leadBCmd = &cobra.Command{
	Use:   "leafB",
	Short: "A brief description of your command",
	Run:   func(cmd *cobra.Command, args []string) { fmt.Println("sub2 leafB called") },
}
