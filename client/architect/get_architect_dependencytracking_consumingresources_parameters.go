// Code generated by go-swagger; DO NOT EDIT.

package architect

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

// NewGetArchitectDependencytrackingConsumingresourcesParams creates a new GetArchitectDependencytrackingConsumingresourcesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetArchitectDependencytrackingConsumingresourcesParams() *GetArchitectDependencytrackingConsumingresourcesParams {
	return &GetArchitectDependencytrackingConsumingresourcesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetArchitectDependencytrackingConsumingresourcesParamsWithTimeout creates a new GetArchitectDependencytrackingConsumingresourcesParams object
// with the ability to set a timeout on a request.
func NewGetArchitectDependencytrackingConsumingresourcesParamsWithTimeout(timeout time.Duration) *GetArchitectDependencytrackingConsumingresourcesParams {
	return &GetArchitectDependencytrackingConsumingresourcesParams{
		timeout: timeout,
	}
}

// NewGetArchitectDependencytrackingConsumingresourcesParamsWithContext creates a new GetArchitectDependencytrackingConsumingresourcesParams object
// with the ability to set a context for a request.
func NewGetArchitectDependencytrackingConsumingresourcesParamsWithContext(ctx context.Context) *GetArchitectDependencytrackingConsumingresourcesParams {
	return &GetArchitectDependencytrackingConsumingresourcesParams{
		Context: ctx,
	}
}

// NewGetArchitectDependencytrackingConsumingresourcesParamsWithHTTPClient creates a new GetArchitectDependencytrackingConsumingresourcesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetArchitectDependencytrackingConsumingresourcesParamsWithHTTPClient(client *http.Client) *GetArchitectDependencytrackingConsumingresourcesParams {
	return &GetArchitectDependencytrackingConsumingresourcesParams{
		HTTPClient: client,
	}
}

/*
GetArchitectDependencytrackingConsumingresourcesParams contains all the parameters to send to the API endpoint

	for the get architect dependencytracking consumingresources operation.

	Typically these are written to a http.Request.
*/
type GetArchitectDependencytrackingConsumingresourcesParams struct {

	/* FlowFilter.

	   Show only checkedIn or published flows
	*/
	FlowFilter *string

	/* ID.

	   Consumed object ID
	*/
	ID string

	/* ObjectType.

	   Consumed object type
	*/
	ObjectType string

	/* PageNumber.

	   Page number

	   Format: int32
	   Default: 1
	*/
	PageNumber *int32

	/* PageSize.

	   Page size

	   Format: int32
	   Default: 25
	*/
	PageSize *int32

	/* ResourceType.

	   Types of consuming resources to show.  Only versioned types are allowed here.
	*/
	ResourceType []string

	/* Version.

	   Object version
	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get architect dependencytracking consumingresources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetArchitectDependencytrackingConsumingresourcesParams) WithDefaults() *GetArchitectDependencytrackingConsumingresourcesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get architect dependencytracking consumingresources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetArchitectDependencytrackingConsumingresourcesParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)
	)

	val := GetArchitectDependencytrackingConsumingresourcesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get architect dependencytracking consumingresources params
func (o *GetArchitectDependencytrackingConsumingresourcesParams) WithTimeout(timeout time.Duration) *GetArchitectDependencytrackingConsumingresourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get architect dependencytracking consumingresources params
func (o *GetArchitectDependencytrackingConsumingresourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get architect dependencytracking consumingresources params
func (o *GetArchitectDependencytrackingConsumingresourcesParams) WithContext(ctx context.Context) *GetArchitectDependencytrackingConsumingresourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get architect dependencytracking consumingresources params
func (o *GetArchitectDependencytrackingConsumingresourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get architect dependencytracking consumingresources params
func (o *GetArchitectDependencytrackingConsumingresourcesParams) WithHTTPClient(client *http.Client) *GetArchitectDependencytrackingConsumingresourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get architect dependencytracking consumingresources params
func (o *GetArchitectDependencytrackingConsumingresourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFlowFilter adds the flowFilter to the get architect dependencytracking consumingresources params
func (o *GetArchitectDependencytrackingConsumingresourcesParams) WithFlowFilter(flowFilter *string) *GetArchitectDependencytrackingConsumingresourcesParams {
	o.SetFlowFilter(flowFilter)
	return o
}

// SetFlowFilter adds the flowFilter to the get architect dependencytracking consumingresources params
func (o *GetArchitectDependencytrackingConsumingresourcesParams) SetFlowFilter(flowFilter *string) {
	o.FlowFilter = flowFilter
}

// WithID adds the id to the get architect dependencytracking consumingresources params
func (o *GetArchitectDependencytrackingConsumingresourcesParams) WithID(id string) *GetArchitectDependencytrackingConsumingresourcesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get architect dependencytracking consumingresources params
func (o *GetArchitectDependencytrackingConsumingresourcesParams) SetID(id string) {
	o.ID = id
}

// WithObjectType adds the objectType to the get architect dependencytracking consumingresources params
func (o *GetArchitectDependencytrackingConsumingresourcesParams) WithObjectType(objectType string) *GetArchitectDependencytrackingConsumingresourcesParams {
	o.SetObjectType(objectType)
	return o
}

// SetObjectType adds the objectType to the get architect dependencytracking consumingresources params
func (o *GetArchitectDependencytrackingConsumingresourcesParams) SetObjectType(objectType string) {
	o.ObjectType = objectType
}

// WithPageNumber adds the pageNumber to the get architect dependencytracking consumingresources params
func (o *GetArchitectDependencytrackingConsumingresourcesParams) WithPageNumber(pageNumber *int32) *GetArchitectDependencytrackingConsumingresourcesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get architect dependencytracking consumingresources params
func (o *GetArchitectDependencytrackingConsumingresourcesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get architect dependencytracking consumingresources params
func (o *GetArchitectDependencytrackingConsumingresourcesParams) WithPageSize(pageSize *int32) *GetArchitectDependencytrackingConsumingresourcesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get architect dependencytracking consumingresources params
func (o *GetArchitectDependencytrackingConsumingresourcesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithResourceType adds the resourceType to the get architect dependencytracking consumingresources params
func (o *GetArchitectDependencytrackingConsumingresourcesParams) WithResourceType(resourceType []string) *GetArchitectDependencytrackingConsumingresourcesParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the get architect dependencytracking consumingresources params
func (o *GetArchitectDependencytrackingConsumingresourcesParams) SetResourceType(resourceType []string) {
	o.ResourceType = resourceType
}

// WithVersion adds the version to the get architect dependencytracking consumingresources params
func (o *GetArchitectDependencytrackingConsumingresourcesParams) WithVersion(version *string) *GetArchitectDependencytrackingConsumingresourcesParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get architect dependencytracking consumingresources params
func (o *GetArchitectDependencytrackingConsumingresourcesParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetArchitectDependencytrackingConsumingresourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FlowFilter != nil {

		// query param flowFilter
		var qrFlowFilter string

		if o.FlowFilter != nil {
			qrFlowFilter = *o.FlowFilter
		}
		qFlowFilter := qrFlowFilter
		if qFlowFilter != "" {

			if err := r.SetQueryParam("flowFilter", qFlowFilter); err != nil {
				return err
			}
		}
	}

	// query param id
	qrID := o.ID
	qID := qrID
	if qID != "" {

		if err := r.SetQueryParam("id", qID); err != nil {
			return err
		}
	}

	// query param objectType
	qrObjectType := o.ObjectType
	qObjectType := qrObjectType
	if qObjectType != "" {

		if err := r.SetQueryParam("objectType", qObjectType); err != nil {
			return err
		}
	}

	if o.PageNumber != nil {

		// query param pageNumber
		var qrPageNumber int32

		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt32(qrPageNumber)
		if qPageNumber != "" {

			if err := r.SetQueryParam("pageNumber", qPageNumber); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.ResourceType != nil {

		// binding items for resourceType
		joinedResourceType := o.bindParamResourceType(reg)

		// query array param resourceType
		if err := r.SetQueryParam("resourceType", joinedResourceType...); err != nil {
			return err
		}
	}

	if o.Version != nil {

		// query param version
		var qrVersion string

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetArchitectDependencytrackingConsumingresources binds the parameter resourceType
func (o *GetArchitectDependencytrackingConsumingresourcesParams) bindParamResourceType(formats strfmt.Registry) []string {
	resourceTypeIR := o.ResourceType

	var resourceTypeIC []string
	for _, resourceTypeIIR := range resourceTypeIR { // explode []string

		resourceTypeIIV := resourceTypeIIR // string as string
		resourceTypeIC = append(resourceTypeIC, resourceTypeIIV)
	}

	// items.CollectionFormat: "multi"
	resourceTypeIS := swag.JoinByFormat(resourceTypeIC, "multi")

	return resourceTypeIS
}
