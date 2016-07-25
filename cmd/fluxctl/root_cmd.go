package main

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/weaveworks/fluxy"
)

type rootOpts struct {
	URL   string
	Fluxd flux.Service
}

func newRoot() *rootOpts {
	return &rootOpts{}
}

var rootLongHelp = strings.TrimSpace(`
fluxctl helps you deploy your code.

Workflow:
  fluxctl service list                             # Which services are running?
  fluxctl service images -s helloworld             # Which images are available?
  fluxctl config update --TODO                     # Update local files, use new image.
  fluxctl service release -s helloworld -f rc.yaml # Release new version.
`)

func (opts *rootOpts) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "fluxctl",
		Long:              rootLongHelp,
		SilenceUsage:      true,
		PersistentPreRunE: opts.PersistentPreRunE,
	}
	cmd.PersistentFlags().StringVarP(&opts.URL, "url", "u", "http://localhost:3030/v0", "base URL of the fluxd API server")
	return cmd
}

func (opts *rootOpts) PersistentPreRunE(*cobra.Command, []string) error {
	var err error
	opts.Fluxd, err = flux.NewClient(opts.URL)
	return err
}
