package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/willdot/uGit/pkg/cli"
)

func main() {
	if err := execute(); err != nil {
		log.Fatal(err)
	}
}

type ugitCli struct {
	rootCmd *cobra.Command
}

func execute() error {
	rootCmd := &cobra.Command{
		Use:   "uGit",
		Short: "uGit is an application to run common git commands quicker ",
	}

	rootCmd.AddCommand(cli.CheckoutCommand())
	rootCmd.AddCommand(cli.CommitCommand())
	rootCmd.AddCommand(cli.MergeCommand())
	rootCmd.AddCommand(cli.DeleteCommand())

	cli := ugitCli{
		rootCmd: rootCmd,
	}

	if err := cli.rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}
