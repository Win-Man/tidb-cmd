package command

import (
	"github.com/spf13/cobra"
)

func newTestCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "test <name>",
		Short: "test toolkit",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return cmd.Help()
			}
			switch args[0] {
			case "log":
				logTest()
			default:
				return cmd.Help()
			}
			return nil
		},
	}
	return cmd
}

func logTest() {

}
