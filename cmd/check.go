package cmd

import (
	"fmt"
	"strings"

	"github.com/go-frame/go-base-framework/g"
	"github.com/spf13/cobra"
)

var Check = &cobra.Command{
	Use:   "check [Module ...]",
	Short: "Check the status of Bp-Test modules",
	Long: `
Check if the specified Bp-Test modules are running.
Modules:
  ` + "all " + strings.Join(g.AllModulesInOrder, " "),
	RunE: check,
}

func check(c *cobra.Command, args []string) error {
	args = g.RmDup(args)

	if len(args) == 0 {
		args = g.AllModulesInOrder
	}

	for _, moduleName := range args {
		if !g.HasModule(moduleName) {
			return fmt.Errorf("%s doesn't exist", moduleName)
		}

		if g.IsRunning(moduleName) {
			fmt.Printf("%20s %10s %15s \n", g.ModuleApps[moduleName], "UP", g.Pid(moduleName))
		} else {
			fmt.Printf("%20s %10s %15s \n", g.ModuleApps[moduleName], "DOWN", "-")
		}
	}

	return nil
}
