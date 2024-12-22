package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func main() {
	var willMonitor bool
	defaultTestEnv := "native"

	upload := &cobra.Command{
		Use:   "upload [environment]",
		Short: "Upload firmware to the board",
		Long:  `Upload firmware to the selected board. Defaults to lilygo-t-display-s3 if no environment is specified.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			env := ""
			if len(args) > 0 {
				env = args[0]
			}

			var pioCmd *exec.Cmd
			if willMonitor {
				pioCmd = exec.Command("platformio", "run", "--target", "upload", "--target", "monitor", "--environment", env)
			} else {
				pioCmd = exec.Command("platformio", "run", "--target", "upload", "--environment", env)
			}
			pioCmd.Stdout = os.Stdout
			pioCmd.Stderr = os.Stderr
			pioCmd.Stdin = os.Stdin
			err := pioCmd.Run()
			if err != nil {
				fmt.Println("Error uploading: ", err)
				return
			}
		},
	}

	test := &cobra.Command{
		Use:   "test [environment]",
		Short: "Run test suite",
		Long:  `Run the test suite for the selected environment. Defaults to native if no environment is specified.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			env := defaultTestEnv
			if len(args) > 0 {
				env = args[0]
			}

			pioCmd := exec.Command("platformio", "test", "--environment", env)
			pioCmd.Stdout = os.Stdout
			pioCmd.Stderr = os.Stderr
			err := pioCmd.Run()
			if err != nil {
				fmt.Println("Could not run tests: ", err)
				return
			}
		},
	}

	build := &cobra.Command{
		Use:   "build [environment]",
		Short: "Build the project",
		Long:  `Build the project for the selected environment. Defaults to lilygo-t-display-s3 if no environment is specified.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			env := ""
			if len(args) > 0 {
				env = args[0]
			}

			pioCmd := exec.Command("platformio", "run", "--environment", env)
			pioCmd.Stdout = os.Stdout
			pioCmd.Stderr = os.Stderr
			err := pioCmd.Run()
			if err != nil {
				fmt.Println("Could not build: ", err)
				return
			}
		},
	}

	lsp := &cobra.Command{
		Use:   "lsp [environment]",
		Short: "Configure Clangd for LSP",
		Long:  `Configure Clangd for LSP. Defaults to lilygo-t-display-s3 if no environment is specified.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			env := ""
			if len(args) > 0 {
				env = args[0]
			}

			err := lsp(env)
			if err != nil {
				fmt.Println("Error configuring Clangd for LSP: ", err)
				return
			}
			fmt.Println(".clangd file created successfully.")
		},
	}

	home := &cobra.Command{
		Use:   "home",
		Short: "Open Platformio Home",
		Long:  `Open Platformio Home in the default web browser.`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			pioCmd := exec.Command("platformio", "home")
			pioCmd.Stdout = os.Stdout
			pioCmd.Stderr = os.Stderr
			err := pioCmd.Run()
			if err != nil {
				fmt.Println("Error opening Platformio Home: ", err)
				return
			}
		},
	}

	monitor := &cobra.Command{
		Use:   "monitor",
		Short: "Monitor the serial port",
		Long:  `Monitor the serial port of the connected board.`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			pioCmd := exec.Command("pio", "device", "monitor")
			pioCmd.Stdout = os.Stdout
			pioCmd.Stderr = os.Stderr
			pioCmd.Stdin = os.Stdin
			err := pioCmd.Run()
			if err != nil {
				fmt.Println("Could not monitor the serial port: ", err)
				return
			}
		},
	}

	upload.Flags().BoolVarP(&willMonitor, "monitor", "m", false, "")

	rootCmd := &cobra.Command{Use: "zed-platformio"}
	rootCmd.AddCommand(home, upload, monitor, lsp, test, build)
	rootCmd.Execute()
}
