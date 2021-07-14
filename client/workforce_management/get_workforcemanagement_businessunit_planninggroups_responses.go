// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetWorkforcemanagementBusinessunitPlanninggroupsReader is a Reader for the GetWorkforcemanagementBusinessunitPlanninggroups structure.
type GetWorkforcemanagementBusinessunitPlanninggroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsOK creates a GetWorkforcemanagementBusinessunitPlanninggroupsOK with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsOK() *GetWorkforcemanagementBusinessunitPlanninggroupsOK {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsOK{}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupsOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsOK struct {
	Payload *models.PlanningGroupList
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsOK) GetPayload() *models.PlanningGroupList {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlanningGroupList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsBadRequest creates a GetWorkforcemanagementBusinessunitPlanninggroupsBadRequest with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsBadRequest() *GetWorkforcemanagementBusinessunitPlanninggroupsBadRequest {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsBadRequest{}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized creates a GetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized() *GetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized{}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsForbidden creates a GetWorkforcemanagementBusinessunitPlanninggroupsForbidden with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsForbidden() *GetWorkforcemanagementBusinessunitPlanninggroupsForbidden {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsForbidden{}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsNotFound creates a GetWorkforcemanagementBusinessunitPlanninggroupsNotFound with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsNotFound() *GetWorkforcemanagementBusinessunitPlanninggroupsNotFound {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsNotFound{}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout creates a GetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout() *GetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout{}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge creates a GetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge() *GetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge{}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType creates a GetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType() *GetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType{}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests creates a GetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests() *GetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests{}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError creates a GetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError() *GetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError{}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable creates a GetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable() *GetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable{}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout creates a GetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout() *GetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout{}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
