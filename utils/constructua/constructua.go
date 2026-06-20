// Constructing user agent
package constructua

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
		config.GetProjectURL(),
		config.GetEmail(),
		config.GetRequestPackageName(),
	)
}
