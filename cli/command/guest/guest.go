package guest

import (
	"github.com/br4v3st4rr/proxmox-api-go/cli"
	"github.com/spf13/cobra"
)

var GuestCmd = &cobra.Command{
	Use:   "guest",
	Short: "Commands to interact with the qemu and LXC guests on Proxmox",
}

func init() {
	cli.RootCmd.AddCommand(GuestCmd)
}
