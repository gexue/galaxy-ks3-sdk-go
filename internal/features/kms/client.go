//Package kms provides gucumber integration tests suppport.
package kms

import (
	"github.com/gexue/galaxy-ks3-sdk-go/internal/features/shared"
	"github.com/gexue/galaxy-ks3-sdk-go/service/kms"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@kms", func() {
		World["client"] = kms.New(nil)
	})
}
