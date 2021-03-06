// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OAuth2Error o auth2 error
//
// swagger:model OAuth2Error
type OAuth2Error struct {

	// Error code indicating high level source of error
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | invalid_request | The request is missing a required parameter, includes an unsupported parameter value (other than grant type), repeats a parameter, includes multiple credentials, utilizes more than one mechanism for authenticating the client, or is otherwise malformed. |
	// | invalid_client | Client authentication failed (e.g., unknown client, no client authentication included, or unsupported authentication method).  The authorization server MAY return an HTTP 401 (Unauthorized) status code to indicate which HTTP authentication schemes are supported.  If the client attempted to authenticate via the 'Authorization' request header field, the authorization server MUST respond with an HTTP 401 (Unauthorized) status code and include the 'WWW-Authenticate' response header field matching the authentication scheme used by the client. |
	// | invalid_grant | The provided authorization grant (e.g., authorization code, resource owner credentials) or refresh token is invalid, expired, revoked, does not match the redirection URI used in the authorization request, or was issued to another client. |
	// | unauthorized_client | The authenticated client is not authorized to use this authorization grant type. |
	// | unsupported_grant_type | The authorization grant type is not supported by the authorization server. |
	// | invalid_scope | The requested scope is invalid, unknown, malformed, or exceeds the scope granted by the resource owner. |
	// | unsupported_response_type | The authorization server does not support obtaining an authorization code using this method. |
	// | server_error | The authorization server encountered an unexpected condition that prevented it from fulfilling the request. (This error code is needed because a 500 Internal Server Error HTTP status code cannot be returned to the client via a HTTP redirect.) |
	//
	// Enum: [invalid_request invalid_client invalid_grant unauthorized_client unsupported_grant_type invalid_scope unsupported_response_type server_error]
	Error string `json:"error,omitempty"`

	// Optional text providing additional information about the error condition.
	ErrorDescription string `json:"error_description,omitempty"`
}

// Validate validates this o auth2 error
func (m *OAuth2Error) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var oAuth2ErrorTypeErrorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["invalid_request","invalid_client","invalid_grant","unauthorized_client","unsupported_grant_type","invalid_scope","unsupported_response_type","server_error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oAuth2ErrorTypeErrorPropEnum = append(oAuth2ErrorTypeErrorPropEnum, v)
	}
}

const (

	// OAuth2ErrorErrorInvalidRequest captures enum value "invalid_request"
	OAuth2ErrorErrorInvalidRequest string = "invalid_request"

	// OAuth2ErrorErrorInvalidClient captures enum value "invalid_client"
	OAuth2ErrorErrorInvalidClient string = "invalid_client"

	// OAuth2ErrorErrorInvalidGrant captures enum value "invalid_grant"
	OAuth2ErrorErrorInvalidGrant string = "invalid_grant"

	// OAuth2ErrorErrorUnauthorizedClient captures enum value "unauthorized_client"
	OAuth2ErrorErrorUnauthorizedClient string = "unauthorized_client"

	// OAuth2ErrorErrorUnsupportedGrantType captures enum value "unsupported_grant_type"
	OAuth2ErrorErrorUnsupportedGrantType string = "unsupported_grant_type"

	// OAuth2ErrorErrorInvalidScope captures enum value "invalid_scope"
	OAuth2ErrorErrorInvalidScope string = "invalid_scope"

	// OAuth2ErrorErrorUnsupportedResponseType captures enum value "unsupported_response_type"
	OAuth2ErrorErrorUnsupportedResponseType string = "unsupported_response_type"

	// OAuth2ErrorErrorServerError captures enum value "server_error"
	OAuth2ErrorErrorServerError string = "server_error"
)

// prop value enum
func (m *OAuth2Error) validateErrorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, oAuth2ErrorTypeErrorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OAuth2Error) validateError(formats strfmt.Registry) error {

	if swag.IsZero(m.Error) { // not required
		return nil
	}

	// value enum
	if err := m.validateErrorEnum("error", "body", m.Error); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OAuth2Error) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OAuth2Error) UnmarshalBinary(b []byte) error {
	var res OAuth2Error
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
