// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new system API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetLearnAPIPublicV1SystemPoliciesPrivacy(params *GetLearnAPIPublicV1SystemPoliciesPrivacyParams, authInfo runtime.ClientAuthInfoWriter) (*GetLearnAPIPublicV1SystemPoliciesPrivacyOK, error)

	GetLearnAPIPublicV1SystemTasksTaskID(params *GetLearnAPIPublicV1SystemTasksTaskIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetLearnAPIPublicV1SystemTasksTaskIDOK, error)

	GetLearnAPIPublicV1SystemVersion(params *GetLearnAPIPublicV1SystemVersionParams, authInfo runtime.ClientAuthInfoWriter) (*GetLearnAPIPublicV1SystemVersionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetLearnAPIPublicV1SystemPoliciesPrivacy gets policies

  Returns the links to the Blackboard and Institution privacy policies

**Since**: 3400.5.0
*/
func (a *Client) GetLearnAPIPublicV1SystemPoliciesPrivacy(params *GetLearnAPIPublicV1SystemPoliciesPrivacyParams, authInfo runtime.ClientAuthInfoWriter) (*GetLearnAPIPublicV1SystemPoliciesPrivacyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLearnAPIPublicV1SystemPoliciesPrivacyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLearnAPIPublicV1SystemPoliciesPrivacy",
		Method:             "GET",
		PathPattern:        "/learn/api/public/v1/system/policies/privacy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLearnAPIPublicV1SystemPoliciesPrivacyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLearnAPIPublicV1SystemPoliciesPrivacyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetLearnAPIPublicV1SystemPoliciesPrivacy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLearnAPIPublicV1SystemTasksTaskID gets system task

  Get the background task by the given task Id.

**Since**: 3800.1.0
*/
func (a *Client) GetLearnAPIPublicV1SystemTasksTaskID(params *GetLearnAPIPublicV1SystemTasksTaskIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetLearnAPIPublicV1SystemTasksTaskIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLearnAPIPublicV1SystemTasksTaskIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLearnAPIPublicV1SystemTasksTaskID",
		Method:             "GET",
		PathPattern:        "/learn/api/public/v1/system/tasks/{taskId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLearnAPIPublicV1SystemTasksTaskIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLearnAPIPublicV1SystemTasksTaskIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetLearnAPIPublicV1SystemTasksTaskID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLearnAPIPublicV1SystemVersion gets version

  Gets the current Learn server version.

**Since**: 3000.3.0
*/
func (a *Client) GetLearnAPIPublicV1SystemVersion(params *GetLearnAPIPublicV1SystemVersionParams, authInfo runtime.ClientAuthInfoWriter) (*GetLearnAPIPublicV1SystemVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLearnAPIPublicV1SystemVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLearnAPIPublicV1SystemVersion",
		Method:             "GET",
		PathPattern:        "/learn/api/public/v1/system/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLearnAPIPublicV1SystemVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLearnAPIPublicV1SystemVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetLearnAPIPublicV1SystemVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
