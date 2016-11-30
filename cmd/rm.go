package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Create the rm command
var cmdRemove = &cobra.Command{
	Use:   "rm WORKFLOW",
	Short: "Remove a workflow",
	Long: `Remove a Swif workflow.

Use "sw cancel WORKFLOW" to cancel a workflow without deleting it.

Usage example:
sw rm ana
`,
	Run: runRemove,
}

func init() {
	cmdSW.AddCommand(cmdRemove)
}

func runRemove(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, `Required workflow argument is missing.

Run "sw rm -h" for usage details.`)
		os.Exit(2)
	}
	run("swif", "cancel", "-workflow", args[0], "-delete")
}
