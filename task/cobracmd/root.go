package cobracmd

import (
	cobra "github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:	"task",
	Short:	"this is a CLI task manager",
}
