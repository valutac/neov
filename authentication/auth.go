package authentication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Valutac/neov"
	"github.com/Valutac/neov/cfg"
)

type service struct {
	cfg *cfg.Configuration
}

func NewService(cfg *cfg.Configuration) *service {
	return &service{cfg}
}

func (srv *service) Login() (string, error) {
	authURL := fmt.Sprintf("%s/auth/tokens?nocatalog", srv.cfg.Credential.Host)
	// fmt.Printf("authenticating to: %s\n", authURL)
	creds := data{}
	creds.Auth.Identity.Methods = []string{"password"}
	creds.Auth.Identity.Password.User.Domain = &domain{neov.UserDomainName}
	creds.Auth.Identity.Password.User.Name = srv.cfg.Credential.Username
	creds.Auth.Identity.Password.User.Password = srv.cfg.Credential.Password
	creds.Auth.Scope = &scope{
		Project: project{ID: srv.cfg.Credential.ProjectID},
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(creds)
	res, _ := http.Post(authURL, "application/json; charset=utf-8", b)

	token := res.Header.Get("X-Subject-Token")
	if token == "" {
		return token, ErrAuthenticationFailed
	}

	return token, nil
}
