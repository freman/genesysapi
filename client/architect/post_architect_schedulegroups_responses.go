// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostArchitectSchedulegroupsReader is a Reader for the PostArchitectSchedulegroups structure.
type PostArchitectSchedulegroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostArchitectSchedulegroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostArchitectSchedulegroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostArchitectSchedulegroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostArchitectSchedulegroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostArchitectSchedulegroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostArchitectSchedulegroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostArchitectSchedulegroupsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostArchitectSchedulegroupsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostArchitectSchedulegroupsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostArchitectSchedulegroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostArchitectSchedulegroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostArchitectSchedulegroupsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostArchitectSchedulegroupsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostArchitectSchedulegroupsOK creates a PostArchitectSchedulegroupsOK with default headers values
func NewPostArchitectSchedulegroupsOK() *PostArchitectSchedulegroupsOK {
	return &PostArchitectSchedulegroupsOK{}
}

/*PostArchitectSchedulegroupsOK handles this case with default header values.

successful operation
*/
type PostArchitectSchedulegroupsOK struct {
	Payload *models.ScheduleGroup
}

func (o *PostArchitectSchedulegroupsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsOK  %+v", 200, o.Payload)
}

func (o *PostArchitectSchedulegroupsOK) GetPayload() *models.ScheduleGroup {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScheduleGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulegroupsBadRequest creates a PostArchitectSchedulegroupsBadRequest with default headers values
func NewPostArchitectSchedulegroupsBadRequest() *PostArchitectSchedulegroupsBadRequest {
	return &PostArchitectSchedulegroupsBadRequest{}
}

/*PostArchitectSchedulegroupsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostArchitectSchedulegroupsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectSchedulegroupsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsBadRequest  %+v", 400, o.Payload)
}

func (o *PostArchitectSchedulegroupsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulegroupsUnauthorized creates a PostArchitectSchedulegroupsUnauthorized with default headers values
func NewPostArchitectSchedulegroupsUnauthorized() *PostArchitectSchedulegroupsUnauthorized {
	return &PostArchitectSchedulegroupsUnauthorized{}
}

/*PostArchitectSchedulegroupsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostArchitectSchedulegroupsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectSchedulegroupsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostArchitectSchedulegroupsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulegroupsForbidden creates a PostArchitectSchedulegroupsForbidden with default headers values
func NewPostArchitectSchedulegroupsForbidden() *PostArchitectSchedulegroupsForbidden {
	return &PostArchitectSchedulegroupsForbidden{}
}

/*PostArchitectSchedulegroupsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostArchitectSchedulegroupsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectSchedulegroupsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsForbidden  %+v", 403, o.Payload)
}

func (o *PostArchitectSchedulegroupsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulegroupsNotFound creates a PostArchitectSchedulegroupsNotFound with default headers values
func NewPostArchitectSchedulegroupsNotFound() *PostArchitectSchedulegroupsNotFound {
	return &PostArchitectSchedulegroupsNotFound{}
}

/*PostArchitectSchedulegroupsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostArchitectSchedulegroupsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectSchedulegroupsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsNotFound  %+v", 404, o.Payload)
}

func (o *PostArchitectSchedulegroupsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulegroupsRequestTimeout creates a PostArchitectSchedulegroupsRequestTimeout with default headers values
func NewPostArchitectSchedulegroupsRequestTimeout() *PostArchitectSchedulegroupsRequestTimeout {
	return &PostArchitectSchedulegroupsRequestTimeout{}
}

/*PostArchitectSchedulegroupsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostArchitectSchedulegroupsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectSchedulegroupsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostArchitectSchedulegroupsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulegroupsRequestEntityTooLarge creates a PostArchitectSchedulegroupsRequestEntityTooLarge with default headers values
func NewPostArchitectSchedulegroupsRequestEntityTooLarge() *PostArchitectSchedulegroupsRequestEntityTooLarge {
	return &PostArchitectSchedulegroupsRequestEntityTooLarge{}
}

/*PostArchitectSchedulegroupsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostArchitectSchedulegroupsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectSchedulegroupsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostArchitectSchedulegroupsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulegroupsUnsupportedMediaType creates a PostArchitectSchedulegroupsUnsupportedMediaType with default headers values
func NewPostArchitectSchedulegroupsUnsupportedMediaType() *PostArchitectSchedulegroupsUnsupportedMediaType {
	return &PostArchitectSchedulegroupsUnsupportedMediaType{}
}

/*PostArchitectSchedulegroupsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostArchitectSchedulegroupsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectSchedulegroupsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostArchitectSchedulegroupsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulegroupsTooManyRequests creates a PostArchitectSchedulegroupsTooManyRequests with default headers values
func NewPostArchitectSchedulegroupsTooManyRequests() *PostArchitectSchedulegroupsTooManyRequests {
	return &PostArchitectSchedulegroupsTooManyRequests{}
}

/*PostArchitectSchedulegroupsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostArchitectSchedulegroupsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectSchedulegroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostArchitectSchedulegroupsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulegroupsInternalServerError creates a PostArchitectSchedulegroupsInternalServerError with default headers values
func NewPostArchitectSchedulegroupsInternalServerError() *PostArchitectSchedulegroupsInternalServerError {
	return &PostArchitectSchedulegroupsInternalServerError{}
}

/*PostArchitectSchedulegroupsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostArchitectSchedulegroupsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectSchedulegroupsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostArchitectSchedulegroupsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulegroupsServiceUnavailable creates a PostArchitectSchedulegroupsServiceUnavailable with default headers values
func NewPostArchitectSchedulegroupsServiceUnavailable() *PostArchitectSchedulegroupsServiceUnavailable {
	return &PostArchitectSchedulegroupsServiceUnavailable{}
}

/*PostArchitectSchedulegroupsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostArchitectSchedulegroupsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectSchedulegroupsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostArchitectSchedulegroupsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSchedulegroupsGatewayTimeout creates a PostArchitectSchedulegroupsGatewayTimeout with default headers values
func NewPostArchitectSchedulegroupsGatewayTimeout() *PostArchitectSchedulegroupsGatewayTimeout {
	return &PostArchitectSchedulegroupsGatewayTimeout{}
}

/*PostArchitectSchedulegroupsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostArchitectSchedulegroupsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectSchedulegroupsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/schedulegroups][%d] postArchitectSchedulegroupsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostArchitectSchedulegroupsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSchedulegroupsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
