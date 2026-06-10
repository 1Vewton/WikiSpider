package construct_ua

import (
	"fmt"

	"github.com/1Vewton/WikiSpider/utils/config"
)

// Construct user agent
func ConstructUA() string {
	// User agent template
	return fmt.Sprintf(
		"%s/%s (%s; %s) %s",
		config.GetProjectName(),
		config.GetVersion(),
		config.GetProjectUrl(),
		config.GetEmail(),
		config.GetRequestPackageName(),
	)
}
