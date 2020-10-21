// Code generated by go-swagger; DO NOT EDIT.

package course_memberships

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetLearnAPIPublicV1UsersUserIDCoursesParams creates a new GetLearnAPIPublicV1UsersUserIDCoursesParams object
// with the default values initialized.
func NewGetLearnAPIPublicV1UsersUserIDCoursesParams() *GetLearnAPIPublicV1UsersUserIDCoursesParams {
	var ()
	return &GetLearnAPIPublicV1UsersUserIDCoursesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV1UsersUserIDCoursesParamsWithTimeout creates a new GetLearnAPIPublicV1UsersUserIDCoursesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV1UsersUserIDCoursesParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV1UsersUserIDCoursesParams {
	var ()
	return &GetLearnAPIPublicV1UsersUserIDCoursesParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV1UsersUserIDCoursesParamsWithContext creates a new GetLearnAPIPublicV1UsersUserIDCoursesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV1UsersUserIDCoursesParamsWithContext(ctx context.Context) *GetLearnAPIPublicV1UsersUserIDCoursesParams {
	var ()
	return &GetLearnAPIPublicV1UsersUserIDCoursesParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV1UsersUserIDCoursesParamsWithHTTPClient creates a new GetLearnAPIPublicV1UsersUserIDCoursesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV1UsersUserIDCoursesParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV1UsersUserIDCoursesParams {
	var ()
	return &GetLearnAPIPublicV1UsersUserIDCoursesParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV1UsersUserIDCoursesParams contains all the parameters to send to the API endpoint
for the get learn API public v1 users user ID courses operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV1UsersUserIDCoursesParams struct {

	/*AvailabilityAvailable
	  Search for users with availability.available properties that contain this value.

	**Since**: 3100.0.0


	| Type      | Description
	 | --------- | --------- |
	| Yes |  |
	| No |  |
	| Disabled |   **Since**: 3100.0.0 |


	*/
	AvailabilityAvailable *string
	/*Created
	  Search for memberships with a created date relative to this value.  'createdCompare' may also be sent to control search behavior.

	**Since**: 3100.0.0

	*/
	Created *strfmt.DateTime
	/*CreatedCompare
	  Used alongside the 'created' search parameter.  Supported values:

	- lessThan
	- greaterOrEqual

	Defaults to greaterOrEqual if not specified.

	**Since**: 3100.0.0


	| Type      | Description
	 | --------- | --------- |
	| lessThan |  |
	| greaterOrEqual |  |


	*/
	CreatedCompare *string
	/*DataSourceID
	  Search for memberships with this dataSourceId.  This may optionally be the data source's externalId using the syntax "externalId:math101".

	**Since**: 3100.0.0

	*/
	DataSourceID *string
	/*Expand
	  A comma-delimited list of fields to expand as part of the response. Expanded fields may cause additional load time. Supported fields are:<br><ul><li>course</li></ul>

	*/
	Expand *string
	/*Fields
	  A comma-delimited list of fields to include in the response. If not specified, all fields will be returned.

	*/
	Fields *string
	/*LastAccessed
	  Search for memberships with a last accessed date relative to this value.  'lastAccessedCompare' may also be sent to control search behavior.

	**Since**: 3300.9.0

	*/
	LastAccessed *strfmt.DateTime
	/*LastAccessedCompare
	  Used alongside the 'lastAccessed' search parameter.  Supported values:

	- lessThan
	- greaterOrEqual

	Defaults to greaterOrEqual if not specified.

	**Since**: 3300.9.0


	| Type      | Description
	 | --------- | --------- |
	| lessThan |  |
	| greaterOrEqual |  |


	*/
	LastAccessedCompare *string
	/*Limit
	  The maximum number of results to be returned. There may be less if the query returned less than the maximum.

	*/
	Limit *int32
	/*Modified
	  Search for memberships with a modified date relative to this value. 'modifiedCompare' may also be sent to control search behavior.

	**Since**: 3800.9.0

	*/
	Modified *strfmt.DateTime
	/*ModifiedCompare
	  Used alongside the 'modified' search parameter. Supported values:

	- lessThan
	- greaterOrEqual

	Defaults to greaterOrEqual if not specified.

	**Since**: 3800.9.0


	| Type      | Description
	 | --------- | --------- |
	| lessThan |  |
	| greaterOrEqual |  |


	*/
	ModifiedCompare *string
	/*Offset
	  The number of rows to skip before beginning to return rows. An offset of 0 is the same as omitting the offset parameter.

	*/
	Offset *int32
	/*Role
	  Search for memberships with a course role id that matches this value.

	**Since**: 3500.5.0

	*/
	Role *string
	/*Sort
	  Properties to sort by. This is a comma-delimited list of JSON properties, with an optional "(desc)" suffix to request a descending sort for that property. e.g. "created(desc)" Supported fields are:

	- created
	- lastAccessed (Since 3300.9.0)

	**Since**: 3100.0.0

	*/
	Sort *string
	/*UserID
	 The user ID.  This may be the primary ID, or any of the secondary IDs prefixed with the ID type.

	| ID type    | Example                               |
	|------------|---------------------------------------|
	| primary    | _123_1                                |
	| externalId | externalId:jsmith                     |
	| userName   | userName:jsmith                       |
	| uuid       | uuid:915c7567d76d444abf1eed56aad3beb5 |


	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV1UsersUserIDCoursesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) WithContext(ctx context.Context) *GetLearnAPIPublicV1UsersUserIDCoursesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV1UsersUserIDCoursesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAvailabilityAvailable adds the availabilityAvailable to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) WithAvailabilityAvailable(availabilityAvailable *string) *GetLearnAPIPublicV1UsersUserIDCoursesParams {
	o.SetAvailabilityAvailable(availabilityAvailable)
	return o
}

// SetAvailabilityAvailable adds the availabilityAvailable to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) SetAvailabilityAvailable(availabilityAvailable *string) {
	o.AvailabilityAvailable = availabilityAvailable
}

// WithCreated adds the created to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) WithCreated(created *strfmt.DateTime) *GetLearnAPIPublicV1UsersUserIDCoursesParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) SetCreated(created *strfmt.DateTime) {
	o.Created = created
}

// WithCreatedCompare adds the createdCompare to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) WithCreatedCompare(createdCompare *string) *GetLearnAPIPublicV1UsersUserIDCoursesParams {
	o.SetCreatedCompare(createdCompare)
	return o
}

// SetCreatedCompare adds the createdCompare to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) SetCreatedCompare(createdCompare *string) {
	o.CreatedCompare = createdCompare
}

// WithDataSourceID adds the dataSourceID to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) WithDataSourceID(dataSourceID *string) *GetLearnAPIPublicV1UsersUserIDCoursesParams {
	o.SetDataSourceID(dataSourceID)
	return o
}

// SetDataSourceID adds the dataSourceId to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) SetDataSourceID(dataSourceID *string) {
	o.DataSourceID = dataSourceID
}

// WithExpand adds the expand to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) WithExpand(expand *string) *GetLearnAPIPublicV1UsersUserIDCoursesParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) SetExpand(expand *string) {
	o.Expand = expand
}

// WithFields adds the fields to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) WithFields(fields *string) *GetLearnAPIPublicV1UsersUserIDCoursesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLastAccessed adds the lastAccessed to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) WithLastAccessed(lastAccessed *strfmt.DateTime) *GetLearnAPIPublicV1UsersUserIDCoursesParams {
	o.SetLastAccessed(lastAccessed)
	return o
}

// SetLastAccessed adds the lastAccessed to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) SetLastAccessed(lastAccessed *strfmt.DateTime) {
	o.LastAccessed = lastAccessed
}

// WithLastAccessedCompare adds the lastAccessedCompare to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) WithLastAccessedCompare(lastAccessedCompare *string) *GetLearnAPIPublicV1UsersUserIDCoursesParams {
	o.SetLastAccessedCompare(lastAccessedCompare)
	return o
}

// SetLastAccessedCompare adds the lastAccessedCompare to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) SetLastAccessedCompare(lastAccessedCompare *string) {
	o.LastAccessedCompare = lastAccessedCompare
}

// WithLimit adds the limit to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) WithLimit(limit *int32) *GetLearnAPIPublicV1UsersUserIDCoursesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithModified adds the modified to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) WithModified(modified *strfmt.DateTime) *GetLearnAPIPublicV1UsersUserIDCoursesParams {
	o.SetModified(modified)
	return o
}

// SetModified adds the modified to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) SetModified(modified *strfmt.DateTime) {
	o.Modified = modified
}

// WithModifiedCompare adds the modifiedCompare to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) WithModifiedCompare(modifiedCompare *string) *GetLearnAPIPublicV1UsersUserIDCoursesParams {
	o.SetModifiedCompare(modifiedCompare)
	return o
}

// SetModifiedCompare adds the modifiedCompare to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) SetModifiedCompare(modifiedCompare *string) {
	o.ModifiedCompare = modifiedCompare
}

// WithOffset adds the offset to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) WithOffset(offset *int32) *GetLearnAPIPublicV1UsersUserIDCoursesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithRole adds the role to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) WithRole(role *string) *GetLearnAPIPublicV1UsersUserIDCoursesParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) SetRole(role *string) {
	o.Role = role
}

// WithSort adds the sort to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) WithSort(sort *string) *GetLearnAPIPublicV1UsersUserIDCoursesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithUserID adds the userID to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) WithUserID(userID string) *GetLearnAPIPublicV1UsersUserIDCoursesParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get learn API public v1 users user ID courses params
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV1UsersUserIDCoursesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AvailabilityAvailable != nil {

		// query param availability.available
		var qrAvailabilityAvailable string
		if o.AvailabilityAvailable != nil {
			qrAvailabilityAvailable = *o.AvailabilityAvailable
		}
		qAvailabilityAvailable := qrAvailabilityAvailable
		if qAvailabilityAvailable != "" {
			if err := r.SetQueryParam("availability.available", qAvailabilityAvailable); err != nil {
				return err
			}
		}

	}

	if o.Created != nil {

		// query param created
		var qrCreated strfmt.DateTime
		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated.String()
		if qCreated != "" {
			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}

	}

	if o.CreatedCompare != nil {

		// query param createdCompare
		var qrCreatedCompare string
		if o.CreatedCompare != nil {
			qrCreatedCompare = *o.CreatedCompare
		}
		qCreatedCompare := qrCreatedCompare
		if qCreatedCompare != "" {
			if err := r.SetQueryParam("createdCompare", qCreatedCompare); err != nil {
				return err
			}
		}

	}

	if o.DataSourceID != nil {

		// query param dataSourceId
		var qrDataSourceID string
		if o.DataSourceID != nil {
			qrDataSourceID = *o.DataSourceID
		}
		qDataSourceID := qrDataSourceID
		if qDataSourceID != "" {
			if err := r.SetQueryParam("dataSourceId", qDataSourceID); err != nil {
				return err
			}
		}

	}

	if o.Expand != nil {

		// query param expand
		var qrExpand string
		if o.Expand != nil {
			qrExpand = *o.Expand
		}
		qExpand := qrExpand
		if qExpand != "" {
			if err := r.SetQueryParam("expand", qExpand); err != nil {
				return err
			}
		}

	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	if o.LastAccessed != nil {

		// query param lastAccessed
		var qrLastAccessed strfmt.DateTime
		if o.LastAccessed != nil {
			qrLastAccessed = *o.LastAccessed
		}
		qLastAccessed := qrLastAccessed.String()
		if qLastAccessed != "" {
			if err := r.SetQueryParam("lastAccessed", qLastAccessed); err != nil {
				return err
			}
		}

	}

	if o.LastAccessedCompare != nil {

		// query param lastAccessedCompare
		var qrLastAccessedCompare string
		if o.LastAccessedCompare != nil {
			qrLastAccessedCompare = *o.LastAccessedCompare
		}
		qLastAccessedCompare := qrLastAccessedCompare
		if qLastAccessedCompare != "" {
			if err := r.SetQueryParam("lastAccessedCompare", qLastAccessedCompare); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Modified != nil {

		// query param modified
		var qrModified strfmt.DateTime
		if o.Modified != nil {
			qrModified = *o.Modified
		}
		qModified := qrModified.String()
		if qModified != "" {
			if err := r.SetQueryParam("modified", qModified); err != nil {
				return err
			}
		}

	}

	if o.ModifiedCompare != nil {

		// query param modifiedCompare
		var qrModifiedCompare string
		if o.ModifiedCompare != nil {
			qrModifiedCompare = *o.ModifiedCompare
		}
		qModifiedCompare := qrModifiedCompare
		if qModifiedCompare != "" {
			if err := r.SetQueryParam("modifiedCompare", qModifiedCompare); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Role != nil {

		// query param role
		var qrRole string
		if o.Role != nil {
			qrRole = *o.Role
		}
		qRole := qrRole
		if qRole != "" {
			if err := r.SetQueryParam("role", qRole); err != nil {
				return err
			}
		}

	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}