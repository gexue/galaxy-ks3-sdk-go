//Package elb provides gucumber integration tests suppport.
package elb

import (
	"github.com/gexue/galaxy-ks3-sdk-go/internal/features/shared"
	"github.com/gexue/galaxy-ks3-sdk-go/service/elb"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@elb", func() {
		World["client"] = elb.New(nil)
	})
}
