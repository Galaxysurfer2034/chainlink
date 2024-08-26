package chainlink

import (
	"math/big"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink/v2/core/config"
	"github.com/smartcontractkit/chainlink/v2/core/config/toml"
)

var _ config.Capabilities = (*capabilitiesConfig)(nil)

type capabilitiesConfig struct {
	c toml.Capabilities
}

func (c *capabilitiesConfig) Peering() config.P2P {
	return &p2p{c: c.c.Peering}
}

func (c *capabilitiesConfig) ExternalRegistry() config.CapabilitiesExternalRegistry {
	return &capabilitiesExternalRegistry{
		c: c.c.ExternalRegistry,
	}
}

func (c *capabilitiesConfig) WorkflowConnectorConfig() config.WorkflowConnectorConfig {
	// I don't understand how the mixing of the toml.WorkflowConnectorConfig can be forced to pick up the config.GatewayConnectorConfig.
	// when it's coming from the toml.  It seems to be intermixing the toml.WorkflowConnectorConfig with the config.WorkflowConnectorConfig
	// And I don't understand the use of a reference access with the use of the c: c.c.* syntax
	// error described below.
	// cannot use &workflowConnectorConfig{…} (value of type *workflowConnectorConfig) as
	// "github.com/smartcontractkit/chainlink/v2/core/config".WorkflowConnectorConfig value in return statement:
	//  *workflowConnectorConfig does not implement "github.com/smartcontractkit/chainlink/v2/core/config".WorkflowConnectorConfig (wrong type for method GatewayConnectorConfig)
	//	have GatewayConnectorConfig() "github.com/smartcontractkit/chainlink/v2/core/config/toml".GatewayConnectorConfig
	//	want GatewayConnectorConfig() "github.com/smartcontractkit/chainlink/v2/core/config".GatewayConnectorConfigcompilerInvalidIfaceAssign
	return &workflowConnectorConfig{
		c: c.c.WorkflowConnectorConfig,
	}
}

type capabilitiesExternalRegistry struct {
	c toml.ExternalRegistry
}

func (c *capabilitiesExternalRegistry) RelayID() types.RelayID {
	return types.NewRelayID(c.NetworkID(), c.ChainID())
}

func (c *capabilitiesExternalRegistry) NetworkID() string {
	return *c.c.NetworkID
}

func (c *capabilitiesExternalRegistry) ChainID() string {
	return *c.c.ChainID
}

func (c *capabilitiesExternalRegistry) Address() string {
	return *c.c.Address
}

type workflowConnectorConfig struct {
	c toml.WorkflowConnectorConfig
}

func (c *workflowConnectorConfig) ChainIDForNodeKey() big.Int {
	return c.c.ChainIDForNodeKey
}

func (c *workflowConnectorConfig) GatewayConnectorConfig() toml.GatewayConnectorConfig {
	// invalid operation: cannot indirect
	// c.c.GatewayConnectorConfig (variable of type "github.com/smartcontractkit/chainlink/v2/core/config/toml".ConnectorConfig)
	// compilerInvalidIndirection
	return c.c.GatewayConnectorConfig
}
