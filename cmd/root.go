package cmd

import (
	"os"

	"github.com/dfuentes/torrent-grabber/config"
	"github.com/dfuentes/torrent-grabber/grabber"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "torrent-grabber",
	Short: "Simple tool for downloading torrent files from an rss feed",
	RunE:  runRoot,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")

}

func runRoot(cmd *cobra.Command, args []string) error {
	if cfgFile == "" {
		return errors.New("must specify config file path")
	}

	config, err := config.Load(cfgFile)
	if err != nil {
		return errors.Wrap(err, "could not load config: ")
	}

	grabber.Grab(config)

	return nil
}
