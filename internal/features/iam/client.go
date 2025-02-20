//Package iam provides gucumber integration tests suppport.
package iam

import (
	"github.com/gexue/galaxy-ks3-sdk-go/internal/features/shared"
	"github.com/gexue/galaxy-ks3-sdk-go/service/iam"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@iam", func() {
		World["client"] = iam.New(nil)
	})
}
