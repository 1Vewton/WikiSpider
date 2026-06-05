package construct_ua

import (
	"fmt"

	"github.com/1Vewton/WikiSpider/utils/config"
)

func ConstructUA() string {
	return fmt.Sprintf(
		"%s/%s (%s; %s) %s",
		config.GetProjectName(),
		config.GetVersion(),
		config.GetProjectUrl(),
		config.GetEmail(),
		config.GetRequestPackageName(),
	)
}
