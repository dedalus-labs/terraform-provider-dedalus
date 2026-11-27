// File generated from our OpenAPI spec by Dedalus. See CONTRIBUTING.md for details.

package machine_port

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*MachinePortResource)(nil)

func (r *MachinePortResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
