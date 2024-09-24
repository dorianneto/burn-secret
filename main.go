package main

import (
	"crypto/tls"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dorianneto/burn-secret/cmd/api"
	"github.com/dorianneto/burn-secret/internal"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/acme/autocert"
)

func getSelfSignedOrLetsEncryptCert(certManager *autocert.Manager) func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	return func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		dirCache, ok := certManager.Cache.(autocert.DirCache)
		if !ok {
			dirCache = "certs"
		}

		keyFile := filepath.Join(string(dirCache), hello.ServerName+".key")
		crtFile := filepath.Join(string(dirCache), hello.ServerName+".crt")
		certificate, err := tls.LoadX509KeyPair(crtFile, keyFile)
		if err != nil {
			return certManager.GetCertificate(hello)
		}
		return &certificate, err
	}
}

func main() {
	godotenv.Load(".env." + os.Getenv("APP_ENV"))

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	database, err := internal.NewDatabase(logger)
	if err != nil {
		logger.Error(err.Error())
	}

	app := api.NewApp(logger, database)

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(os.Getenv("APP_DOMAIN")),
		Cache:      autocert.DirCache("certs"),
	}

	tlsConfig := certManager.TLSConfig()
	tlsConfig.GetCertificate = getSelfSignedOrLetsEncryptCert(&certManager)

	server := http.Server{
		Addr:      ":443",
		Handler:   app.Routes(),
		TLSConfig: tlsConfig,
	}

	logger.Info(fmt.Sprintf("server running on port :%s", server.Addr))

	go http.ListenAndServe(":80", certManager.HTTPHandler(nil))

	if err := server.ListenAndServeTLS("", ""); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
