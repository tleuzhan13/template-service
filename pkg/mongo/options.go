package mongo

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	errNoSuchFile    = errors.New("no such file or directory")
	errPemFile       = errors.New("failed parsing pem file")
	errEmptyFilePath = errors.New("db tls file path is empty")
)

func (m Config) genConnectURL() string {
	var url string
	if m.Username == "" || m.Password == "" {
		url = fmt.Sprintf("mongodb://%s/?tls=%t&retryWrites=false", m.URI, m.TLSEnable)
	} else {
		url = fmt.Sprintf("mongodb://%s:%s@%s/?tls=%t&retryWrites=false", m.Username, m.Password, m.URI, m.TLSEnable)
	}

	return url
}

func (m Config) getTLSConfig(ctx context.Context) (*tls.Config, error) {
	if m.TLSFilePath == "" {
		return nil, errEmptyFilePath
	}

	tlsConfig := new(tls.Config)

	certs, err := os.ReadFile(m.TLSFilePath)
	if err != nil && strings.Contains(err.Error(), errNoSuchFile.Error()) {
		err = downloadKey(ctx, m.TLSFilePath)
		if err != nil {
			return tlsConfig, err
		}

		certs, _ = os.ReadFile(m.TLSFilePath)
	}

	tlsConfig.RootCAs = x509.NewCertPool()
	ok := tlsConfig.RootCAs.AppendCertsFromPEM(certs)

	if !ok {
		return tlsConfig, errPemFile
	}

	tlsConfig.InsecureSkipVerify = true

	return tlsConfig, nil
}

func downloadKey(ctx context.Context, filePath string) error {
	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet, "https://truststore.pki.rds.amazonaws.com/global/global-bundle.pem", nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}
