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

// PostWorkforcemanagementManagementunitsReader is a Reader for the PostWorkforcemanagementManagementunits structure.
type PostWorkforcemanagementManagementunitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementManagementunitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWorkforcemanagementManagementunitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostWorkforcemanagementManagementunitsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWorkforcemanagementManagementunitsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementManagementunitsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementManagementunitsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementManagementunitsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostWorkforcemanagementManagementunitsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementManagementunitsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementManagementunitsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementManagementunitsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementManagementunitsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementManagementunitsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementManagementunitsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWorkforcemanagementManagementunitsOK creates a PostWorkforcemanagementManagementunitsOK with default headers values
func NewPostWorkforcemanagementManagementunitsOK() *PostWorkforcemanagementManagementunitsOK {
	return &PostWorkforcemanagementManagementunitsOK{}
}

/*PostWorkforcemanagementManagementunitsOK handles this case with default header values.

successful operation
*/
type PostWorkforcemanagementManagementunitsOK struct {
	Payload *models.ManagementUnit
}

func (o *PostWorkforcemanagementManagementunitsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsOK) GetPayload() *models.ManagementUnit {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManagementUnit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsCreated creates a PostWorkforcemanagementManagementunitsCreated with default headers values
func NewPostWorkforcemanagementManagementunitsCreated() *PostWorkforcemanagementManagementunitsCreated {
	return &PostWorkforcemanagementManagementunitsCreated{}
}

/*PostWorkforcemanagementManagementunitsCreated handles this case with default header values.

The management unit was successfully created
*/
type PostWorkforcemanagementManagementunitsCreated struct {
	Payload *models.ManagementUnit
}

func (o *PostWorkforcemanagementManagementunitsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsCreated  %+v", 201, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsCreated) GetPayload() *models.ManagementUnit {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManagementUnit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsBadRequest creates a PostWorkforcemanagementManagementunitsBadRequest with default headers values
func NewPostWorkforcemanagementManagementunitsBadRequest() *PostWorkforcemanagementManagementunitsBadRequest {
	return &PostWorkforcemanagementManagementunitsBadRequest{}
}

/*PostWorkforcemanagementManagementunitsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementManagementunitsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsUnauthorized creates a PostWorkforcemanagementManagementunitsUnauthorized with default headers values
func NewPostWorkforcemanagementManagementunitsUnauthorized() *PostWorkforcemanagementManagementunitsUnauthorized {
	return &PostWorkforcemanagementManagementunitsUnauthorized{}
}

/*PostWorkforcemanagementManagementunitsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementManagementunitsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsForbidden creates a PostWorkforcemanagementManagementunitsForbidden with default headers values
func NewPostWorkforcemanagementManagementunitsForbidden() *PostWorkforcemanagementManagementunitsForbidden {
	return &PostWorkforcemanagementManagementunitsForbidden{}
}

/*PostWorkforcemanagementManagementunitsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementManagementunitsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsNotFound creates a PostWorkforcemanagementManagementunitsNotFound with default headers values
func NewPostWorkforcemanagementManagementunitsNotFound() *PostWorkforcemanagementManagementunitsNotFound {
	return &PostWorkforcemanagementManagementunitsNotFound{}
}

/*PostWorkforcemanagementManagementunitsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementManagementunitsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsRequestTimeout creates a PostWorkforcemanagementManagementunitsRequestTimeout with default headers values
func NewPostWorkforcemanagementManagementunitsRequestTimeout() *PostWorkforcemanagementManagementunitsRequestTimeout {
	return &PostWorkforcemanagementManagementunitsRequestTimeout{}
}

/*PostWorkforcemanagementManagementunitsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostWorkforcemanagementManagementunitsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsRequestEntityTooLarge creates a PostWorkforcemanagementManagementunitsRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementManagementunitsRequestEntityTooLarge() *PostWorkforcemanagementManagementunitsRequestEntityTooLarge {
	return &PostWorkforcemanagementManagementunitsRequestEntityTooLarge{}
}

/*PostWorkforcemanagementManagementunitsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostWorkforcemanagementManagementunitsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsUnsupportedMediaType creates a PostWorkforcemanagementManagementunitsUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementManagementunitsUnsupportedMediaType() *PostWorkforcemanagementManagementunitsUnsupportedMediaType {
	return &PostWorkforcemanagementManagementunitsUnsupportedMediaType{}
}

/*PostWorkforcemanagementManagementunitsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementManagementunitsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsTooManyRequests creates a PostWorkforcemanagementManagementunitsTooManyRequests with default headers values
func NewPostWorkforcemanagementManagementunitsTooManyRequests() *PostWorkforcemanagementManagementunitsTooManyRequests {
	return &PostWorkforcemanagementManagementunitsTooManyRequests{}
}

/*PostWorkforcemanagementManagementunitsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostWorkforcemanagementManagementunitsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsInternalServerError creates a PostWorkforcemanagementManagementunitsInternalServerError with default headers values
func NewPostWorkforcemanagementManagementunitsInternalServerError() *PostWorkforcemanagementManagementunitsInternalServerError {
	return &PostWorkforcemanagementManagementunitsInternalServerError{}
}

/*PostWorkforcemanagementManagementunitsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementManagementunitsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsServiceUnavailable creates a PostWorkforcemanagementManagementunitsServiceUnavailable with default headers values
func NewPostWorkforcemanagementManagementunitsServiceUnavailable() *PostWorkforcemanagementManagementunitsServiceUnavailable {
	return &PostWorkforcemanagementManagementunitsServiceUnavailable{}
}

/*PostWorkforcemanagementManagementunitsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementManagementunitsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitsGatewayTimeout creates a PostWorkforcemanagementManagementunitsGatewayTimeout with default headers values
func NewPostWorkforcemanagementManagementunitsGatewayTimeout() *PostWorkforcemanagementManagementunitsGatewayTimeout {
	return &PostWorkforcemanagementManagementunitsGatewayTimeout{}
}

/*PostWorkforcemanagementManagementunitsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostWorkforcemanagementManagementunitsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits][%d] postWorkforcemanagementManagementunitsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
