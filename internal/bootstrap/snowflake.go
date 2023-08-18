// Package bootstrap
package bootstrap

import (
	"fmt"

	"gitlab.com/willysihombing/task-c3/internal/helper"
	"gitlab.com/willysihombing/task-c3/pkg/generator"
	"gitlab.com/willysihombing/task-c3/pkg/logger"
)

// RegistrySnowflake setup snowflake generator
func RegistrySnowflake() {
	hs := helper.GetHostname()
	nodeID := uint64(helper.GetNodeID(hs))

	lf := logger.NewFields(
		logger.EventName("SetupSnowflake"),
		logger.Any("node_id", nodeID),
		logger.Any("hostname", hs),
	)

	logger.Info(fmt.Sprintf(`generate node id for snowflake is %d`, nodeID), lf...)
	generator.Setup(nodeID)
}
