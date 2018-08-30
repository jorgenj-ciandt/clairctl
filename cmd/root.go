package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/coreos/pkg/capnslog"
	"github.com/spf13/cobra"
)

var errInternalError = errors.New("client quit unexpectedly")

var log = capnslog.NewPackageLogger("github.com/jorgenj-ciandt/clairctl", "cmd")

var logLevel string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "clair-report",
	Short: "Generate HTML report from Clair vulnerabilities JSON",
	Long:  ``,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
}
