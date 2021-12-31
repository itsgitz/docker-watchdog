/*
Copyright © 2022 Anggit M Ginanjar <anggit@isi.co.id> from PT Infinys System Indonesia

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

	"docker-watchdog/watchdog"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultCfgPath = "/opt"
	defaultCfgFile = ".docker-watchdog"
)

var cfgFile string
var parseConfig watchdog.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "docker-watchdog",
	Short: "Watch and detect all stopped docker containers",
	Long: `Docker Watchdog is Go program that used for detects
all stopped containers that have exited, paused, or dead state/status.
It uses list containers endpoints from Docker Engine API.

The watcher will gather all docker containers information
repeatedly every 3 seconds using go Ticker function.
If there are stopped containers detected, Docker Watchdog will
send an email alert to Developer or System Administrator.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		"config file (default is /opt/.docker-watchdog.yaml)",
	)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(defaultCfgPath)
		viper.SetConfigType("yaml")
		viper.SetConfigName(defaultCfgFile)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		watchdog.InformationText.Fprintln(os.Stderr, "[*] Using config file:", viper.ConfigFileUsed())
	}

	//Unmarshalling the config value
	if err := viper.Unmarshal(&parseConfig); err != nil {
		cobra.CheckErr(err)
	}
}
