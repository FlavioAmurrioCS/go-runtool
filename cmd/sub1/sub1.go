package sub1

import (
	"fmt"
	"os"

	"github.com/FlavioAmurrioCS/go-runtool/cmd/sub1/sub2"
	"github.com/spf13/cobra"
)

var Sub1Cmd = &cobra.Command{
	Use:   "sub1",
	Short: "A brief description of your command",
}

func init() {
	Sub1Cmd.AddCommand(sub2.Sub2Cmd)
}

func Execute() {
	if err := Sub1Cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
