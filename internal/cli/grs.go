package cli

import (
	"strings"

	flexera "github.com/flexera-public/unified-go-client"
)

// EnvGRSBaseURL overrides the GRS base URL.
const EnvGRSBaseURL = "FLEXERA_CLI_GRS_BASE_URL"

// GRSBaseURL resolves the GRS (Governance / Resource Service) host base URL.
//
// GRS is a legacy service outside the unified gateway; project enumeration has
// no IAM equivalent, so the CLI points a client at this host. Precedence:
// override flag > FLEXERA_CLI_GRS_BASE_URL env > --api-base-url > the zone's
// default grs-front host (flexera.GrsBaseURLForZone).
func (d *Deps) GRSBaseURL(override string) string {
	if v := strings.TrimSpace(override); v != "" {
		return v
	}
	if d.Getenv != nil {
		if v := strings.TrimSpace(d.Getenv(EnvGRSBaseURL)); v != "" {
			return v
		}
	}
	if v := strings.TrimSpace(d.Config.APIBaseURL); v != "" {
		return v
	}
	return flexera.GrsBaseURLForZone(d.Config.Zone)
}
