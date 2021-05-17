package terraform

import (
	"github.com/hashicorp/terraform/configs"
	"github.com/hashicorp/terraform/internal/addrs"
)

// GraphNodeAttachProviderMetaConfigs is an interface that must be implemented
// by nodes that want provider meta configurations attached.
type GraphNodeAttachProviderMetaConfigs interface {
	GraphNodeConfigResource

	// Sets the configuration
	AttachProviderMetaConfigs(map[addrs.Provider]*configs.ProviderMeta)
}
