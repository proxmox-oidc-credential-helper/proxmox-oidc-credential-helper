package proxmox

import (
	"fmt"
	"github.com/camaeel/proxmox-oidc-credential-helper/pkg/config"
)

func getCallbackUrl(cfg config.Config) string {
	return fmt.Sprintf("http://localhost:%d%s", cfg.CallbackPort, cfg.CallbackPath)
}
