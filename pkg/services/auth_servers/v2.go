package authservers

import (
	"encoding/json"
	"fmt"
	"github.com/onelogin/onelogin-go-sdk/pkg/services"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/olhttp"
)

const errAuthServersV2Context = "auth_servers v2 service"

// V2Service holds the information needed to interface with a repository
type V2Service struct {
	Endpoint, ErrorContext string
	Repository             services.Repository
}

// New creates the new svc service v2.
func New(repo services.Repository, host string) *V2Service {
	return &V2Service{
		Endpoint:     fmt.Sprintf("%s/api/2/api_authorizations", host),
		Repository:   repo,
		ErrorContext: errAuthServersV2Context,
	}
}

// Query retrieves all the auth_servers from the repository that meet the query criteria passed in the
// request payload. If an empty payload is given, it will retrieve all auth_servers
func (svc *V2Service) Query(query *AuthServerQuery) ([]AuthServer, error) {
	resp, err := svc.Repository.Read(olhttp.OLHTTPRequest{
		URL:        svc.Endpoint,
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    query,
	})
	if err != nil {
		return nil, err
	}

	var auth_servers []AuthServer
	json.Unmarshal(resp, &auth_servers)
	return auth_servers, nil
}

// GetOne retrieves the user by id and returns it
func (svc *V2Service) GetOne(id int32) (*AuthServer, error) {
	resp, err := svc.Repository.Read(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
	})
	if err != nil {
		return nil, err
	}
	var user AuthServer
	json.Unmarshal(resp, &user)
	return &user, nil
}

// Create takes a user without an id and attempts to use the parameters to create it
// in the API. Modifies the user in place, or returns an error if one occurs
func (svc *V2Service) Create(user *AuthServer) error {
	resp, err := svc.Repository.Create(olhttp.OLHTTPRequest{
		URL:        svc.Endpoint,
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    user,
	})
	if err != nil {
		return err
	}
	json.Unmarshal(resp, user)
	return nil
}

// Update takes a user and an id and attempts to use the parameters to update it
// in the API. Modifies the user in place, or returns an error if one occurs
func (svc *V2Service) Update(id int32, user *AuthServer) error {
	resp, err := svc.Repository.Update(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    user,
	})
	if err != nil {
		return err
	}
	json.Unmarshal(resp, user)
	return nil
}

// Destroy deletes the user with the given id, and if successful, it returns nil
func (svc *V2Service) Destroy(id int32) error {
	if _, err := svc.Repository.Destroy(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
	}); err != nil {
		return err
	}
	return nil
}