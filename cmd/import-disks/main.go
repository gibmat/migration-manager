package main

import (
	"context"
	"fmt"
	"os"

	"github.com/lxc/incus/v6/shared/util"
	"github.com/spf13/cobra"

	"github.com/FuturFusion/migration-manager/cmd/common"
	"github.com/FuturFusion/migration-manager/util/vmware"
)

type appFlags struct {
	common.CmdGlobalFlags
	common.CmdVMwareFlags

	cutover bool
	vmName  string
}

func main() {
	appCmd := appFlags{}
	app := appCmd.Command()
	app.SilenceUsage = true
	app.CompletionOptions = cobra.CompletionOptions{DisableDefaultCmd: true}

	// Workaround for main command
	app.Args = cobra.ArbitraryArgs

	// Version handling
	app.SetVersionTemplate("{{.Version}}\n")
	app.Version = common.Version

	// Run the main command and handle errors
	err := app.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func (c *appFlags) Command() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "import-disks"
	cmd.Short = "Import VMware VM disks into Incus"
	cmd.Long = `Description:
  Import VMware VM disks into Incus

  This tool imports VM disks from VMware into Incus. Prior to importing any
  disks, you must first run the `+"`import-vm-metadata`"+` utility to
  populate VM metadata in Incus.

  By default, the state of VMs running in VMware will be unchanged. A
  snapshot of the disk(s) will be created, then transferred via this tool.
  If possible, incremental copies are used to speed up the process.

  When you are ready to cut over to Incus, pass the `+"`--cutover`"+` flag
  which will shutdown the VMware VMs, do a final data transfer and then
  start the VMs in Incus.
`

	cmd.RunE = c.Run

	c.CmdGlobalFlags.AddFlags(cmd)
	c.CmdVMwareFlags.AddFlags(cmd)
	cmd.Flags().BoolVar(&c.cutover, "cutover", false, "Shutdown VMware VMs, perform a final data transfer and start Incus VMs")
	cmd.Flags().StringVar(&c.vmName, "vm-name", "", "VM name from which to import disk (required)")
	cmd.MarkFlagRequired("vm-name")

	return cmd
}

func (c *appFlags) Run(cmd *cobra.Command, args []string) error {
	if os.Geteuid() != 0 {
		return fmt.Errorf("This tool must be run as root")
	}

	if !util.PathExists("/dev/virtio-ports/org.linuxcontainers.incus") {
		return fmt.Errorf("This tool is designed to be run within an Incus VM")
	}

	ctx := context.TODO()

	// Connect to VMware endpoint.
	vmwareClient, err := vmware.NewVMwareClient(ctx, c.VmwareEndpoint, c.VmwareInsecure, c.VmwareUsername, c.VmwarePassword)
	if err != nil {
		return err
	}

	// Get a list of all VMs.
	vmwareVms, err := vmwareClient.GetVMs()
	if err != nil {
		return err
	}

	if c.cutover {
		fmt.Printf("cut over not currently supported!\n")
		return nil
	}

	for _, vm := range vmwareVms {
		if vm.Name() != c.vmName {
			continue
		}

		fmt.Printf("Exporting disks attached to VM %q...\n", vm.Name())

		err := vmwareClient.DeleteVMSnapshot(vm, "incusMigration")
		if err != nil {
			return err
		}

		err = vmwareClient.ExportDisks(vm)
		if err != nil {
			return err
		}
	}

	return nil
}
