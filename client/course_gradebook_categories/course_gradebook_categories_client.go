// Code generated by go-swagger; DO NOT EDIT.

package course_gradebook_categories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new course gradebook categories API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for course gradebook categories API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetLearnAPIPublicV1CoursesCourseIDGradebookCategories(params *GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesParams, authInfo runtime.ClientAuthInfoWriter) (*GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesOK, error)

	GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryID(params *GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetLearnAPIPublicV1CoursesCourseIDGradebookCategories gets gradebook categories

  Returns a list of gradebook categories in a particular course.

Users with entitlements 'course.gradebook.MODIFY' or 'course.user.grades.VIEW', or users enrolled in the specified course can retrieve the list of gradebook categories.

**Since**: 3400.2.0
*/
func (a *Client) GetLearnAPIPublicV1CoursesCourseIDGradebookCategories(params *GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesParams, authInfo runtime.ClientAuthInfoWriter) (*GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLearnAPIPublicV1CoursesCourseIDGradebookCategories",
		Method:             "GET",
		PathPattern:        "/learn/api/public/v1/courses/{courseId}/gradebook/categories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetLearnAPIPublicV1CoursesCourseIDGradebookCategories: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryID gets gradebook category

  Returns the details of a gradebook category

Users with entitlements 'course.gradebook.MODIFY' or 'course.user.grades.VIEW', or users enrolled in the specified course can retrieve gradebook category details.

**Since**: 3400.2.0
*/
func (a *Client) GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryID(params *GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryID",
		Method:             "GET",
		PathPattern:        "/learn/api/public/v1/courses/{courseId}/gradebook/categories/{categoryId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetLearnAPIPublicV1CoursesCourseIDGradebookCategoriesCategoryID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
