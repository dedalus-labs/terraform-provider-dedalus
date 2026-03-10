// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workspace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*WorkspaceResource)(nil)

func (r *WorkspaceResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
