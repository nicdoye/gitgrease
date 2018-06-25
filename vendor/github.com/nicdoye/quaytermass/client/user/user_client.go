// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new user API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateStar Star a repository.
*/
func (a *Client) CreateStar(params *CreateStarParams, authInfo runtime.ClientAuthInfoWriter) (*CreateStarCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateStarParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createStar",
		Method:             "POST",
		PathPattern:        "/api/v1/user/starred",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateStarReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateStarCreated), nil

}

/*
DeleteStar Removes a star from a repository.
*/
func (a *Client) DeleteStar(params *DeleteStarParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteStarNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteStarParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteStar",
		Method:             "DELETE",
		PathPattern:        "/api/v1/user/starred/{repository}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteStarReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteStarNoContent), nil

}

/*
GetLoggedInUser Get user information for the authenticated user.
*/
func (a *Client) GetLoggedInUser(params *GetLoggedInUserParams, authInfo runtime.ClientAuthInfoWriter) (*GetLoggedInUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoggedInUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLoggedInUser",
		Method:             "GET",
		PathPattern:        "/api/v1/user/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLoggedInUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLoggedInUserOK), nil

}

/*
GetUserInformation Get user information for the specified user.
*/
func (a *Client) GetUserInformation(params *GetUserInformationParams) (*GetUserInformationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserInformationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserInformation",
		Method:             "GET",
		PathPattern:        "/api/v1/users/{username}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserInformationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserInformationOK), nil

}

/*
ListStarredRepos List all starred repositories.
*/
func (a *Client) ListStarredRepos(params *ListStarredReposParams, authInfo runtime.ClientAuthInfoWriter) (*ListStarredReposOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListStarredReposParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listStarredRepos",
		Method:             "GET",
		PathPattern:        "/api/v1/user/starred",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListStarredReposReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListStarredReposOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}