// Copyright 2019-present Vic Shóstak. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/create-go-app/cli/pkg/cgapp"
	"github.com/create-go-app/cli/pkg/registry"
	"github.com/spf13/cobra"
)

// deployCmd represents the `deploy` command.
var deployCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"push"},
	Short:   "Deploy your project to the remote server via Ansible",
	Long:    "\nDeploy your project to the remote server by Ansible playbooks and roles.",
	Run:     runDeployCmd,
}

// runDeployCmd represents runner for the `deploy` command.
var runDeployCmd = func(cmd *cobra.Command, args []string) {
	// Start message.
	cgapp.ShowMessage(
		"",
		fmt.Sprintf("Deploying project via Create Go App CLI v`%v`...", registry.CLIVersion),
		true, true,
	)

	// Start timer.
	startTimer := time.Now()

	// Set Ansible playbook and inventory files.
	if askBecomePass {
		// With entering password.
		options = []string{"playbook.yml", "-i", "hosts.ini", "-K"}
	} else {
		// Without entering password.
		options = []string{"playbook.yml", "-i", "hosts.ini"}
	}

	// Create config files for your project.
	cgapp.ShowMessage("info", "Run Ansible playbook for deploy your project...", true, true)

	// Run execution for Ansible playbook.
	if err := cgapp.ExecCommand("ansible-playbook", options, false); err != nil {
		log.Fatal(cgapp.ShowError(err.Error()))
	}

	// Stop timer.
	stopTimer := fmt.Sprintf("%.0f", time.Since(startTimer).Seconds())

	// End messages.
	cgapp.ShowMessage(
		"info",
		fmt.Sprintf("Completed in %v seconds!", stopTimer),
		true, true,
	)
	cgapp.ShowMessage(
		"",
		"Have a great project launch! :)",
		false, true,
	)
}

func init() {
	rootCmd.AddCommand(deployCmd)
	deployCmd.PersistentFlags().BoolVarP(
		&askBecomePass,
		"", "K", false,
		"prompt you to provide the remote user sudo password (standard Ansible `--ask-become-pass` option)",
	)
}
