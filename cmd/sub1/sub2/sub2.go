package sub2

import (
	"github.com/spf13/cobra"
)

var Sub2Cmd = &cobra.Command{
	Use:   "sub2",
	Short: "A brief description of your command",
}

func init() {
	Sub2Cmd.AddCommand(leadACmd)
	Sub2Cmd.AddCommand(leadBCmd)
}
