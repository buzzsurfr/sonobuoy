package cmd

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/buzzsurfr/sonobuoy/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thediveo/enumflag"
)

type ServerType enumflag.Flag

const (
	Tcp ServerType = iota
	Http
	Grpc
)

var ServerTypeIds = map[ServerType][]string{
	Tcp:  {"tcp"},
	Http: {"http"},
	Grpc: {"grpc"},
}

var (
	cfgFile  string
	port     string
	protocol ServerType

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "sonobuoy",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			lis, err := net.Listen("tcp", port)
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			var s server.EchoServer
			switch protocol {
			case Tcp:
				s = &server.TcpServer{}
			case Http:
				s = &server.HttpServer{}
			case Grpc:
				s = &server.GrpcServer{}
			}
			if err := s.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
			defer lis.Close()
		},
	}
)

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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sonobuoy.yaml)")
	rootCmd.PersistentFlags().VarP(
		enumflag.New(&protocol, "protocol", ServerTypeIds, enumflag.EnumCaseInsensitive),
		"protocol", "P", "Protocol: tcp | http | grpc")
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", ":2869", "Port")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".sonobuoy" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".sonobuoy")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
