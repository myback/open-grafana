package extensions

import (
	// Upgrade ldapsync from cron to cron.v3 and
	// remove the cron (v1) dependency

	_ "github.com/beevik/etree"
	_ "github.com/crewjam/saml"
	_ "github.com/go-jose/go-jose/v3"
	_ "github.com/gobwas/glob"
	_ "github.com/grpc-ecosystem/go-grpc-middleware"
	_ "github.com/jung-kurt/gofpdf"
	_ "github.com/linkedin/goavro/v2"
	"github.com/myback/open-grafana/pkg/registry"
	"github.com/myback/open-grafana/pkg/services/licensing"
	"github.com/myback/open-grafana/pkg/services/validations"
	_ "github.com/pkg/errors"
	_ "github.com/robfig/cron"
	_ "github.com/robfig/cron/v3"
	_ "github.com/russellhaering/goxmldsig"
	_ "github.com/stretchr/testify/require"
	_ "github.com/timberio/go-datemath"
	_ "golang.org/x/time/rate"
)

func init() {
	registry.RegisterService(&licensing.OSSLicensingService{})
	registry.RegisterService(&validations.OSSPluginRequestValidator{})
}

var IsEnterprise bool = false
