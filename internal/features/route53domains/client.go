//Package route53domains provides gucumber integration tests suppport.
package route53domains

import (
	"github.com/gexue/galaxy-ks3-sdk-go/internal/features/shared"
	"github.com/gexue/galaxy-ks3-sdk-go/service/route53domains"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@route53domains", func() {
		World["client"] = route53domains.New(nil)
	})
}
