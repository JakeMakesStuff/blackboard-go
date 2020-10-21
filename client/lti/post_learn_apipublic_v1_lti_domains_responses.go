// Code generated by go-swagger; DO NOT EDIT.

package lti

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/jakemakesstuff/blackboard-go/models"
)

// PostLearnApipublicV1LtiDomainsReader is a Reader for the PostLearnApipublicV1LtiDomains structure.
type PostLearnApipublicV1LtiDomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLearnApipublicV1LtiDomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostLearnApipublicV1LtiDomainsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLearnApipublicV1LtiDomainsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLearnApipublicV1LtiDomainsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLearnApipublicV1LtiDomainsCreated creates a PostLearnApipublicV1LtiDomainsCreated with default headers values
func NewPostLearnApipublicV1LtiDomainsCreated() *PostLearnApipublicV1LtiDomainsCreated {
	return &PostLearnApipublicV1LtiDomainsCreated{}
}

/*PostLearnApipublicV1LtiDomainsCreated handles this case with default header values.

Created
*/
type PostLearnApipublicV1LtiDomainsCreated struct {
	Payload *models.LTIDomainConfig
}

func (o *PostLearnApipublicV1LtiDomainsCreated) Error() string {
	return fmt.Sprintf("[POST /learn/apipublic/v1/lti/domains][%d] postLearnApipublicV1LtiDomainsCreated  %+v", 201, o.Payload)
}

func (o *PostLearnApipublicV1LtiDomainsCreated) GetPayload() *models.LTIDomainConfig {
	return o.Payload
}

func (o *PostLearnApipublicV1LtiDomainsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LTIDomainConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnApipublicV1LtiDomainsBadRequest creates a PostLearnApipublicV1LtiDomainsBadRequest with default headers values
func NewPostLearnApipublicV1LtiDomainsBadRequest() *PostLearnApipublicV1LtiDomainsBadRequest {
	return &PostLearnApipublicV1LtiDomainsBadRequest{}
}

/*PostLearnApipublicV1LtiDomainsBadRequest handles this case with default header values.

Bad Request
*/
type PostLearnApipublicV1LtiDomainsBadRequest struct {
	Payload *models.RestException
}

func (o *PostLearnApipublicV1LtiDomainsBadRequest) Error() string {
	return fmt.Sprintf("[POST /learn/apipublic/v1/lti/domains][%d] postLearnApipublicV1LtiDomainsBadRequest  %+v", 400, o.Payload)
}

func (o *PostLearnApipublicV1LtiDomainsBadRequest) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnApipublicV1LtiDomainsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearnApipublicV1LtiDomainsForbidden creates a PostLearnApipublicV1LtiDomainsForbidden with default headers values
func NewPostLearnApipublicV1LtiDomainsForbidden() *PostLearnApipublicV1LtiDomainsForbidden {
	return &PostLearnApipublicV1LtiDomainsForbidden{}
}

/*PostLearnApipublicV1LtiDomainsForbidden handles this case with default header values.

Forbidden
*/
type PostLearnApipublicV1LtiDomainsForbidden struct {
	Payload *models.RestException
}

func (o *PostLearnApipublicV1LtiDomainsForbidden) Error() string {
	return fmt.Sprintf("[POST /learn/apipublic/v1/lti/domains][%d] postLearnApipublicV1LtiDomainsForbidden  %+v", 403, o.Payload)
}

func (o *PostLearnApipublicV1LtiDomainsForbidden) GetPayload() *models.RestException {
	return o.Payload
}

func (o *PostLearnApipublicV1LtiDomainsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestException)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostLearnApipublicV1LtiDomainsBody post learn apipublic v1 lti domains body
swagger:model PostLearnApipublicV1LtiDomainsBody
*/
type PostLearnApipublicV1LtiDomainsBody struct {

	// Whether the Tool is allowed to call the LTI Assignment and Grades service and manage line items and grades.
	//
	// **Since**: 3600.0.0
	AllowGradesService bool `json:"allowGradesService,omitempty"`

	// Whether the Tool is allowed to call the LTI Names and Roles service and get the course memberships.
	AllowMembershipService bool `json:"allowMembershipService,omitempty"`

	// The client id associated with this configuration. Only applicable for LTI versions 1.3+, excluding 2.0
	//
	// **Since**: 3600.0.0
	ClientID string `json:"clientId,omitempty"`

	// The custom parameters for the given domain.
	CustomParameters map[string]string `json:"customParameters,omitempty"`

	// The JWKS URL of the tool, if specified. It is optional and can be null
	//
	// **Since**: 3800.17.0
	JwksURL string `json:"jwksUrl,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// The primary domain name associated with this configuration.
	PrimaryDomain string `json:"primaryDomain,omitempty"`

	// The public key of the tool, if specified. It is optional and can be null
	//
	// **Since**: 3800.17.0
	PublicKey string `json:"publicKey,omitempty"`

	// secret
	Secret string `json:"secret,omitempty"`

	// Whether the user's email address can be sent to the LTI tool provider.
	SendEmail bool `json:"sendEmail,omitempty"`

	// Whether the user's name can be sent to the LTI tool provider.
	SendName bool `json:"sendName,omitempty"`

	// Whether the user's role can be sent to the LTI tool provider.
	SendRole bool `json:"sendRole,omitempty"`

	// Enum indicating when user data can be sent to the LTI tool provider.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Never |  |
	// | Sslonly |  |
	// | Always |   **Since**: 3300.9.0 |
	//
	// Enum: [Never Sslonly Always]
	SendUserDataType string `json:"sendUserDataType,omitempty"`

	// Enum that indicates if the set of domains associated with this config can or cannot be linked to.
	//
	//
	// | Type      | Description
	//  | --------- | --------- |
	// | Approved |  |
	// | Excluded |  |
	// | NeedsApproval |   **Since**: 3300.9.0 |
	//
	// Enum: [Approved Excluded NeedsApproval]
	Status string `json:"status,omitempty"`

	// Whether a splash screen is shown before launching the LTI link.
	//
	// Cannot be set to true if allowMembershipService is true.
	UseSplashScreen bool `json:"useSplashScreen,omitempty"`
}

// Validate validates this post learn apipublic v1 lti domains body
func (o *PostLearnApipublicV1LtiDomainsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSendUserDataType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postLearnApipublicV1LtiDomainsBodyTypeSendUserDataTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Never","Sslonly","Always"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postLearnApipublicV1LtiDomainsBodyTypeSendUserDataTypePropEnum = append(postLearnApipublicV1LtiDomainsBodyTypeSendUserDataTypePropEnum, v)
	}
}

const (

	// PostLearnApipublicV1LtiDomainsBodySendUserDataTypeNever captures enum value "Never"
	PostLearnApipublicV1LtiDomainsBodySendUserDataTypeNever string = "Never"

	// PostLearnApipublicV1LtiDomainsBodySendUserDataTypeSslonly captures enum value "Sslonly"
	PostLearnApipublicV1LtiDomainsBodySendUserDataTypeSslonly string = "Sslonly"

	// PostLearnApipublicV1LtiDomainsBodySendUserDataTypeAlways captures enum value "Always"
	PostLearnApipublicV1LtiDomainsBodySendUserDataTypeAlways string = "Always"
)

// prop value enum
func (o *PostLearnApipublicV1LtiDomainsBody) validateSendUserDataTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postLearnApipublicV1LtiDomainsBodyTypeSendUserDataTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostLearnApipublicV1LtiDomainsBody) validateSendUserDataType(formats strfmt.Registry) error {

	if swag.IsZero(o.SendUserDataType) { // not required
		return nil
	}

	// value enum
	if err := o.validateSendUserDataTypeEnum("input"+"."+"sendUserDataType", "body", o.SendUserDataType); err != nil {
		return err
	}

	return nil
}

var postLearnApipublicV1LtiDomainsBodyTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Approved","Excluded","NeedsApproval"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postLearnApipublicV1LtiDomainsBodyTypeStatusPropEnum = append(postLearnApipublicV1LtiDomainsBodyTypeStatusPropEnum, v)
	}
}

const (

	// PostLearnApipublicV1LtiDomainsBodyStatusApproved captures enum value "Approved"
	PostLearnApipublicV1LtiDomainsBodyStatusApproved string = "Approved"

	// PostLearnApipublicV1LtiDomainsBodyStatusExcluded captures enum value "Excluded"
	PostLearnApipublicV1LtiDomainsBodyStatusExcluded string = "Excluded"

	// PostLearnApipublicV1LtiDomainsBodyStatusNeedsApproval captures enum value "NeedsApproval"
	PostLearnApipublicV1LtiDomainsBodyStatusNeedsApproval string = "NeedsApproval"
)

// prop value enum
func (o *PostLearnApipublicV1LtiDomainsBody) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postLearnApipublicV1LtiDomainsBodyTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostLearnApipublicV1LtiDomainsBody) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("input"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostLearnApipublicV1LtiDomainsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLearnApipublicV1LtiDomainsBody) UnmarshalBinary(b []byte) error {
	var res PostLearnApipublicV1LtiDomainsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}