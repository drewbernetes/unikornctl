package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/drewbernetes/unikornctl/cmd/create"
	"github.com/drewbernetes/unikornctl/cmd/delete"
	"github.com/drewbernetes/unikornctl/cmd/get"
	"github.com/drewbernetes/unikornctl/cmd/update"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfgFile                          string
	url, username, password, project string
	insecure                         bool
)

func Execute() {
	cmd := NewRootCommand()
	commands := []*cobra.Command{
		create.NewCreateCommand(),
		delete.NewDeleteCommand(),
		get.NewGetCommand(),
		update.NewUpdateCommand(),
	}
	cmd.AddCommand(commands...)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "unikornctl",
		Short: "A CLI for working with the Unikorn Kubernetes Service",
		Long:  "unikornctl - manage Kubernetes clusters via the Unikorn Kubernetes Service",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initializeConfig(cmd)
		},
	}

	rootCmd.PersistentFlags().StringVar(&url, "url", "https://clusters.unikorn-cloud.org", "URL to Unikorn API")
	rootCmd.PersistentFlags().StringVar(&username, "username", "", "Username")
	rootCmd.PersistentFlags().StringVar(&password, "password", "", "Password")
	rootCmd.PersistentFlags().StringVar(&project, "project", "", "Project ID")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Path to configuration file")
	rootCmd.PersistentFlags().BoolVar(&insecure, "insecure", false, "Disable server certificate validation, for use when testing")

	return rootCmd
}

func initializeConfig(cmd *cobra.Command) error {
	v := viper.New()

	if cfgFile != "" {
		v.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".unikornctl" (without extension).
		v.AddConfigPath(home)
		v.SetConfigType("yaml")
		v.SetConfigName(".unikornctl")
	}

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	bindFlags(cmd, v)

	return nil
}

func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		configName := f.Name

		if !f.Changed && v.IsSet(configName) {
			val := v.Get(configName)
			err := cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
			if err != nil {
				log.Fatalln(err)
			}
		}
	})
}
