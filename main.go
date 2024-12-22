package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func main() {
	var willMonitor bool
	environment := "lilygo-t-display-s3"
	test_env := "native"

	home := &cobra.Command{
		Use:   "home",
		Short: "Open Platformio Home",
		Long:  `Open Platformio Home in the default web browser.`,
		Args:  cobra.MinimumNArgs(0),
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

	upload := &cobra.Command{
		Use:   "upload",
		Short: "Upload firmware to the board",
		Long:  `Upload firmware to the selected board. Configure the board at zed-platformio.json.`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			var pioCmd *exec.Cmd
			if willMonitor {
				pioCmd = exec.Command("platformio", "run", "--target", "upload", "--target", "monitor", "--environment", environment)
			} else {
				pioCmd = exec.Command("platformio", "run", "--target", "upload", "--environment", environment)
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

	monitor := &cobra.Command{
		Use:   "monitor",
		Short: "Monitor the serial port",
		Long:  `Monitor the serial port of the connected board.`,
		Args:  cobra.MinimumNArgs(0),
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

	test := &cobra.Command{
		Use:   "test",
		Short: "Run test suite",
		Long:  `Run the test suite for the selected environment.`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			pioCmd := exec.Command("platformio", "test", "--environment", test_env)
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
		Use:   "build",
		Short: "Build the project",
		Long:  `Build the project for the selected environment.`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			pioCmd := exec.Command("platformio", "run", "--environment", environment)
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
		Use:   "lsp",
		Short: "Configure Clangd for LSP",
		Long:  `Configure Clangd for LSP.`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			err := lsp(environment)
			if err != nil {
				fmt.Println("Error configuring Clangd for LSP: ", err)
				return
			}
			fmt.Println(".clangd file created successfully.")
		},
	}

	upload.Flags().BoolVarP(&willMonitor, "monitor", "m", false, "")

	rootCmd := &cobra.Command{Use: "zed-platformio"}
	rootCmd.AddCommand(home, upload, monitor, lsp, test, build)
	rootCmd.Execute()
}
