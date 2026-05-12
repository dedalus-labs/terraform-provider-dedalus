// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package machine_ssh_session

import (
	"context"

	"github.com/dedalus-labs/dedalus-go"
	"github.com/dedalus-labs/terraform-provider-dedalus/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/datasource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MachineSSHSessionDataSourceModel struct {
	ID                          types.String                                                                          `tfsdk:"id" path:"session_id,computed"`
	SessionID                   types.String                                                                          `tfsdk:"session_id" path:"session_id,required"`
	MachineID                   types.String                                                                          `tfsdk:"machine_id" path:"machine_id,required"`
	CreatedAt                   timetypes.RFC3339                                                                     `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ErrorCode                   types.String                                                                          `tfsdk:"error_code" json:"error_code,computed"`
	ErrorMessage                types.String                                                                          `tfsdk:"error_message" json:"error_message,computed"`
	ExpiresAt                   timetypes.RFC3339                                                                     `tfsdk:"expires_at" json:"expires_at,computed" format:"date-time"`
	ReadyAt                     timetypes.RFC3339                                                                     `tfsdk:"ready_at" json:"ready_at,computed" format:"date-time"`
	RetryAfterMs                types.Int64                                                                           `tfsdk:"retry_after_ms" json:"retry_after_ms,computed"`
	Status                      types.String                                                                          `tfsdk:"status" json:"status,computed"`
	MachineSSHSessionConnection customfield.NestedObject[MachineSSHSessionMachineSSHSessionConnectionDataSourceModel] `tfsdk:"machine_ssh_session_connection" json:"connection,computed"`
	Timeouts                    timeouts.Value                                                                        `tfsdk:"timeouts"`
}

func (m *MachineSSHSessionDataSourceModel) toReadParams(_ context.Context) (params dedalus.MachineSSHGetParams, diags diag.Diagnostics) {
	params = dedalus.MachineSSHGetParams{
		MachineID: m.MachineID.ValueString(),
		SessionID: m.SessionID.ValueString(),
	}

	return
}

type MachineSSHSessionMachineSSHSessionConnectionDataSourceModel struct {
	Endpoint        types.String                                                                                   `tfsdk:"endpoint" json:"endpoint,computed"`
	Port            types.Int64                                                                                    `tfsdk:"port" json:"port,computed"`
	SSHUsername     types.String                                                                                   `tfsdk:"ssh_username" json:"ssh_username,computed"`
	HostTrust       customfield.NestedObject[MachineSSHSessionMachineSSHSessionConnectionHostTrustDataSourceModel] `tfsdk:"host_trust" json:"host_trust,computed"`
	UserCertificate types.String                                                                                   `tfsdk:"user_certificate" json:"user_certificate,computed"`
}

type MachineSSHSessionMachineSSHSessionConnectionHostTrustDataSourceModel struct {
	HostPattern types.String `tfsdk:"host_pattern" json:"host_pattern,computed"`
	Kind        types.String `tfsdk:"kind" json:"kind,computed"`
	PublicKey   types.String `tfsdk:"public_key" json:"public_key,computed"`
}
