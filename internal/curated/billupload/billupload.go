// Package billupload provides the higher-level bill-upload workflows that are
// implemented by unified-go-client rather than the OpenAPI operation set.
package billupload

import (
	"fmt"

	"github.com/spf13/cobra"

	clipkg "github.com/flexera-public/flexera-cli/internal/cli"
	flexera "github.com/flexera-public/unified-go-client"
)

// NewPushCmd builds the bill-upload push workflow command.
func NewPushCmd() *cobra.Command {
	var (
		billConnectID string
		billingPeriod string
		files         []string
	)
	c := &cobra.Command{
		Use:   "push",
		Short: "Create, upload, and commit a bill upload",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			deps := clipkg.DepsFrom(cmd.Context())
			if err := deps.Config.RequireOrgID(); err != nil {
				return err
			}
			client, err := deps.APIClient()
			if err != nil {
				return err
			}
			input := flexera.BillUploadPushInput{
				OrgID:         deps.Config.OrgID,
				BillConnectID: billConnectID,
				BillingPeriod: billingPeriod,
				Files:         files,
			}
			output, err := flexera.NewBillUploadPushTool(client.ClientInterface).Invoke(cmd.Context(), input)
			if err != nil {
				return err
			}
			return deps.Printer.Render(deps.Stdout, deps.Config.Output, output)
		},
	}
	c.Flags().StringVar(&billConnectID, "bill-connect-id", "", "bill connect ID (required)")
	c.Flags().StringVar(&billingPeriod, "billing-period", "", "billing period in yyyy-mm format (required)")
	c.Flags().StringSliceVar(&files, "file", nil, "local bill file to upload (repeatable, required)")
	return c
}

// Attach adds the curated push workflow below the generated bill-upload
// command, preserving the generated direct endpoint commands.
func Attach(root *cobra.Command) error {
	for _, command := range root.Commands() {
		if command.Name() == "bill-upload" {
			command.AddCommand(NewPushCmd())
			return nil
		}
	}
	return fmt.Errorf("bill-upload generated command is not registered")
}
