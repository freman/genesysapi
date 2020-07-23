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

// PostWorkforcemanagementBusinessunitPlanninggroupsReader is a Reader for the PostWorkforcemanagementBusinessunitPlanninggroups structure.
type PostWorkforcemanagementBusinessunitPlanninggroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementBusinessunitPlanninggroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWorkforcemanagementBusinessunitPlanninggroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWorkforcemanagementBusinessunitPlanninggroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementBusinessunitPlanninggroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementBusinessunitPlanninggroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementBusinessunitPlanninggroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostWorkforcemanagementBusinessunitPlanninggroupsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementBusinessunitPlanninggroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWorkforcemanagementBusinessunitPlanninggroupsOK creates a PostWorkforcemanagementBusinessunitPlanninggroupsOK with default headers values
func NewPostWorkforcemanagementBusinessunitPlanninggroupsOK() *PostWorkforcemanagementBusinessunitPlanninggroupsOK {
	return &PostWorkforcemanagementBusinessunitPlanninggroupsOK{}
}

/*PostWorkforcemanagementBusinessunitPlanninggroupsOK handles this case with default header values.

successful operation
*/
type PostWorkforcemanagementBusinessunitPlanninggroupsOK struct {
	Payload *models.PlanningGroup
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] postWorkforcemanagementBusinessunitPlanninggroupsOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsOK) GetPayload() *models.PlanningGroup {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlanningGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitPlanninggroupsBadRequest creates a PostWorkforcemanagementBusinessunitPlanninggroupsBadRequest with default headers values
func NewPostWorkforcemanagementBusinessunitPlanninggroupsBadRequest() *PostWorkforcemanagementBusinessunitPlanninggroupsBadRequest {
	return &PostWorkforcemanagementBusinessunitPlanninggroupsBadRequest{}
}

/*PostWorkforcemanagementBusinessunitPlanninggroupsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementBusinessunitPlanninggroupsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] postWorkforcemanagementBusinessunitPlanninggroupsBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitPlanninggroupsUnauthorized creates a PostWorkforcemanagementBusinessunitPlanninggroupsUnauthorized with default headers values
func NewPostWorkforcemanagementBusinessunitPlanninggroupsUnauthorized() *PostWorkforcemanagementBusinessunitPlanninggroupsUnauthorized {
	return &PostWorkforcemanagementBusinessunitPlanninggroupsUnauthorized{}
}

/*PostWorkforcemanagementBusinessunitPlanninggroupsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementBusinessunitPlanninggroupsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] postWorkforcemanagementBusinessunitPlanninggroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitPlanninggroupsForbidden creates a PostWorkforcemanagementBusinessunitPlanninggroupsForbidden with default headers values
func NewPostWorkforcemanagementBusinessunitPlanninggroupsForbidden() *PostWorkforcemanagementBusinessunitPlanninggroupsForbidden {
	return &PostWorkforcemanagementBusinessunitPlanninggroupsForbidden{}
}

/*PostWorkforcemanagementBusinessunitPlanninggroupsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementBusinessunitPlanninggroupsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] postWorkforcemanagementBusinessunitPlanninggroupsForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitPlanninggroupsNotFound creates a PostWorkforcemanagementBusinessunitPlanninggroupsNotFound with default headers values
func NewPostWorkforcemanagementBusinessunitPlanninggroupsNotFound() *PostWorkforcemanagementBusinessunitPlanninggroupsNotFound {
	return &PostWorkforcemanagementBusinessunitPlanninggroupsNotFound{}
}

/*PostWorkforcemanagementBusinessunitPlanninggroupsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementBusinessunitPlanninggroupsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] postWorkforcemanagementBusinessunitPlanninggroupsNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitPlanninggroupsConflict creates a PostWorkforcemanagementBusinessunitPlanninggroupsConflict with default headers values
func NewPostWorkforcemanagementBusinessunitPlanninggroupsConflict() *PostWorkforcemanagementBusinessunitPlanninggroupsConflict {
	return &PostWorkforcemanagementBusinessunitPlanninggroupsConflict{}
}

/*PostWorkforcemanagementBusinessunitPlanninggroupsConflict handles this case with default header values.

Conflict
*/
type PostWorkforcemanagementBusinessunitPlanninggroupsConflict struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] postWorkforcemanagementBusinessunitPlanninggroupsConflict  %+v", 409, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge creates a PostWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge() *PostWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge {
	return &PostWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge{}
}

/*PostWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] postWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType creates a PostWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType() *PostWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType {
	return &PostWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType{}
}

/*PostWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] postWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests creates a PostWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests with default headers values
func NewPostWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests() *PostWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests {
	return &PostWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests{}
}

/*PostWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] postWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitPlanninggroupsInternalServerError creates a PostWorkforcemanagementBusinessunitPlanninggroupsInternalServerError with default headers values
func NewPostWorkforcemanagementBusinessunitPlanninggroupsInternalServerError() *PostWorkforcemanagementBusinessunitPlanninggroupsInternalServerError {
	return &PostWorkforcemanagementBusinessunitPlanninggroupsInternalServerError{}
}

/*PostWorkforcemanagementBusinessunitPlanninggroupsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementBusinessunitPlanninggroupsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] postWorkforcemanagementBusinessunitPlanninggroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable creates a PostWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable with default headers values
func NewPostWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable() *PostWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable {
	return &PostWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable{}
}

/*PostWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] postWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout creates a PostWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout with default headers values
func NewPostWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout() *PostWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout {
	return &PostWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout{}
}

/*PostWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] postWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
