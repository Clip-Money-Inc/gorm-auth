package gormaws

import (
	"context"
	"crypto/tls"

	awscerts "github.com/Clip-Money-Inc/gorm-auth/aws/certs"
	"github.com/Invicton-Labs/go-stackerr"
)

// GetTlsConfig will get a *tls.Config that trusts the AWS Root CAs
// for the given host.
func GetTlsConfig(ctx context.Context, host string) (*tls.Config, stackerr.Error) {
	rootCaPool, err := awscerts.GetGlobalRootCertPool(nil)
	if err != nil {
		return nil, stackerr.Wrap(err)
	}
	return &tls.Config{
		RootCAs:    rootCaPool,
		ServerName: host,
	}, nil
}
