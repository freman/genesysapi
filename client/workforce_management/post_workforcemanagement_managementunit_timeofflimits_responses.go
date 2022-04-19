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

// PostWorkforcemanagementManagementunitTimeofflimitsReader is a Reader for the PostWorkforcemanagementManagementunitTimeofflimits structure.
type PostWorkforcemanagementManagementunitTimeofflimitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementManagementunitTimeofflimitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWorkforcemanagementManagementunitTimeofflimitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostWorkforcemanagementManagementunitTimeofflimitsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWorkforcemanagementManagementunitTimeofflimitsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementManagementunitTimeofflimitsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementManagementunitTimeofflimitsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementManagementunitTimeofflimitsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostWorkforcemanagementManagementunitTimeofflimitsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementManagementunitTimeofflimitsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementManagementunitTimeofflimitsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWorkforcemanagementManagementunitTimeofflimitsOK creates a PostWorkforcemanagementManagementunitTimeofflimitsOK with default headers values
func NewPostWorkforcemanagementManagementunitTimeofflimitsOK() *PostWorkforcemanagementManagementunitTimeofflimitsOK {
	return &PostWorkforcemanagementManagementunitTimeofflimitsOK{}
}

/*PostWorkforcemanagementManagementunitTimeofflimitsOK handles this case with default header values.

successful operation
*/
type PostWorkforcemanagementManagementunitTimeofflimitsOK struct {
	Payload *models.TimeOffLimit
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] postWorkforcemanagementManagementunitTimeofflimitsOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsOK) GetPayload() *models.TimeOffLimit {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TimeOffLimit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitTimeofflimitsCreated creates a PostWorkforcemanagementManagementunitTimeofflimitsCreated with default headers values
func NewPostWorkforcemanagementManagementunitTimeofflimitsCreated() *PostWorkforcemanagementManagementunitTimeofflimitsCreated {
	return &PostWorkforcemanagementManagementunitTimeofflimitsCreated{}
}

/*PostWorkforcemanagementManagementunitTimeofflimitsCreated handles this case with default header values.

The time off limit was successfully created
*/
type PostWorkforcemanagementManagementunitTimeofflimitsCreated struct {
	Payload *models.TimeOffLimit
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] postWorkforcemanagementManagementunitTimeofflimitsCreated  %+v", 201, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsCreated) GetPayload() *models.TimeOffLimit {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TimeOffLimit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitTimeofflimitsBadRequest creates a PostWorkforcemanagementManagementunitTimeofflimitsBadRequest with default headers values
func NewPostWorkforcemanagementManagementunitTimeofflimitsBadRequest() *PostWorkforcemanagementManagementunitTimeofflimitsBadRequest {
	return &PostWorkforcemanagementManagementunitTimeofflimitsBadRequest{}
}

/*PostWorkforcemanagementManagementunitTimeofflimitsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementManagementunitTimeofflimitsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] postWorkforcemanagementManagementunitTimeofflimitsBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitTimeofflimitsUnauthorized creates a PostWorkforcemanagementManagementunitTimeofflimitsUnauthorized with default headers values
func NewPostWorkforcemanagementManagementunitTimeofflimitsUnauthorized() *PostWorkforcemanagementManagementunitTimeofflimitsUnauthorized {
	return &PostWorkforcemanagementManagementunitTimeofflimitsUnauthorized{}
}

/*PostWorkforcemanagementManagementunitTimeofflimitsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementManagementunitTimeofflimitsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] postWorkforcemanagementManagementunitTimeofflimitsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitTimeofflimitsForbidden creates a PostWorkforcemanagementManagementunitTimeofflimitsForbidden with default headers values
func NewPostWorkforcemanagementManagementunitTimeofflimitsForbidden() *PostWorkforcemanagementManagementunitTimeofflimitsForbidden {
	return &PostWorkforcemanagementManagementunitTimeofflimitsForbidden{}
}

/*PostWorkforcemanagementManagementunitTimeofflimitsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementManagementunitTimeofflimitsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] postWorkforcemanagementManagementunitTimeofflimitsForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitTimeofflimitsNotFound creates a PostWorkforcemanagementManagementunitTimeofflimitsNotFound with default headers values
func NewPostWorkforcemanagementManagementunitTimeofflimitsNotFound() *PostWorkforcemanagementManagementunitTimeofflimitsNotFound {
	return &PostWorkforcemanagementManagementunitTimeofflimitsNotFound{}
}

/*PostWorkforcemanagementManagementunitTimeofflimitsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementManagementunitTimeofflimitsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] postWorkforcemanagementManagementunitTimeofflimitsNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitTimeofflimitsRequestTimeout creates a PostWorkforcemanagementManagementunitTimeofflimitsRequestTimeout with default headers values
func NewPostWorkforcemanagementManagementunitTimeofflimitsRequestTimeout() *PostWorkforcemanagementManagementunitTimeofflimitsRequestTimeout {
	return &PostWorkforcemanagementManagementunitTimeofflimitsRequestTimeout{}
}

/*PostWorkforcemanagementManagementunitTimeofflimitsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostWorkforcemanagementManagementunitTimeofflimitsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] postWorkforcemanagementManagementunitTimeofflimitsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge creates a PostWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge() *PostWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge {
	return &PostWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge{}
}

/*PostWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] postWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType creates a PostWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType() *PostWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType {
	return &PostWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType{}
}

/*PostWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] postWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitTimeofflimitsTooManyRequests creates a PostWorkforcemanagementManagementunitTimeofflimitsTooManyRequests with default headers values
func NewPostWorkforcemanagementManagementunitTimeofflimitsTooManyRequests() *PostWorkforcemanagementManagementunitTimeofflimitsTooManyRequests {
	return &PostWorkforcemanagementManagementunitTimeofflimitsTooManyRequests{}
}

/*PostWorkforcemanagementManagementunitTimeofflimitsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostWorkforcemanagementManagementunitTimeofflimitsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] postWorkforcemanagementManagementunitTimeofflimitsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitTimeofflimitsInternalServerError creates a PostWorkforcemanagementManagementunitTimeofflimitsInternalServerError with default headers values
func NewPostWorkforcemanagementManagementunitTimeofflimitsInternalServerError() *PostWorkforcemanagementManagementunitTimeofflimitsInternalServerError {
	return &PostWorkforcemanagementManagementunitTimeofflimitsInternalServerError{}
}

/*PostWorkforcemanagementManagementunitTimeofflimitsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementManagementunitTimeofflimitsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] postWorkforcemanagementManagementunitTimeofflimitsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable creates a PostWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable with default headers values
func NewPostWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable() *PostWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable {
	return &PostWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable{}
}

/*PostWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] postWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout creates a PostWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout with default headers values
func NewPostWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout() *PostWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout {
	return &PostWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout{}
}

/*PostWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] postWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
