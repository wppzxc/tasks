package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wppzxc/tasks/src/pkg/types"
	"github.com/wppzxc/tasks/src/pkg/version"
	"os"
)

func NewTasksServerCommand() *cobra.Command {
	cfg :=new(types.Config)
	cmd :=&cobra.Command{
		Use: "tasks",
		Long: "The server of webapp tasks",
		Run: func(cmd *cobra.Command, args []string) {
			getFlagByEnvs(cfg)
			if cfg.Version {
				fmt.Print("tasks : ", version.Get())
				os.Exit(0)
			}
			if err := validateFlags(cfg); err != nil {
				panic(err)
			}
			if err := Run(cfg); err != nil {
				panic(err)
			}
		},
	}
	cmd.Flags().StringVar(&cfg.MysqlUrl,"mysql", "", "The url of mysql")
	cmd.Flags().StringVar(&cfg.BindPort,"bind-ip", "0.0.0.0", "The ip of tasks server")
	cmd.Flags().StringVar(&cfg.BindPort,"bind-port", "8080", "The port of tasks server")
	cmd.Flags().BoolVar(&cfg.Version, "version", false, "Show the version of tasks")
	return cmd
}

func getFlagByEnvs(cfg *types.Config) {
	v := os.Getenv("TASKS_VERSION")
	url := os.Getenv("TASKS_MYSQL_URL")
	ip := os.Getenv("TASKS_BIND_IP")
	port := os.Getenv("TASKS_BIND_PORT")
	if len(v) > 0 {
		fmt.Println("get flag TASKS_VERSION from env : ", v)
		cfg.Version = true
	}
	if len(url) > 0 {
		fmt.Println("get flag TASKS_MYSQL_URL from env : ", url)
		cfg.MysqlUrl = url
	}
	if len(port) > 0 {
		fmt.Println("get flag TASKS_BIND_PORT from env : ", port)
		cfg.BindPort = port
	}
	if len(ip) > 0 {
		fmt.Println("get flag TASKS_BIND_IP from env : ", ip)
		cfg.BindIP = ip
	}
}

func validateFlags(cfg *types.Config) error {
	if len(cfg.MysqlUrl) == 0 {
		return fmt.Errorf("must provide mysql url")
	}
	return nil
}
