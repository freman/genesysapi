// Code generated by go-swagger; DO NOT EDIT.

package locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostLocationsReader is a Reader for the PostLocations structure.
type PostLocationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLocationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLocationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLocationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostLocationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLocationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostLocationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostLocationsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostLocationsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostLocationsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostLocationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostLocationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostLocationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostLocationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLocationsOK creates a PostLocationsOK with default headers values
func NewPostLocationsOK() *PostLocationsOK {
	return &PostLocationsOK{}
}

/*PostLocationsOK handles this case with default header values.

successful operation
*/
type PostLocationsOK struct {
	Payload *models.LocationDefinition
}

func (o *PostLocationsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations][%d] postLocationsOK  %+v", 200, o.Payload)
}

func (o *PostLocationsOK) GetPayload() *models.LocationDefinition {
	return o.Payload
}

func (o *PostLocationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocationDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLocationsBadRequest creates a PostLocationsBadRequest with default headers values
func NewPostLocationsBadRequest() *PostLocationsBadRequest {
	return &PostLocationsBadRequest{}
}

/*PostLocationsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostLocationsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostLocationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations][%d] postLocationsBadRequest  %+v", 400, o.Payload)
}

func (o *PostLocationsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLocationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLocationsUnauthorized creates a PostLocationsUnauthorized with default headers values
func NewPostLocationsUnauthorized() *PostLocationsUnauthorized {
	return &PostLocationsUnauthorized{}
}

/*PostLocationsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostLocationsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostLocationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations][%d] postLocationsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLocationsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLocationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLocationsForbidden creates a PostLocationsForbidden with default headers values
func NewPostLocationsForbidden() *PostLocationsForbidden {
	return &PostLocationsForbidden{}
}

/*PostLocationsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostLocationsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostLocationsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations][%d] postLocationsForbidden  %+v", 403, o.Payload)
}

func (o *PostLocationsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLocationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLocationsNotFound creates a PostLocationsNotFound with default headers values
func NewPostLocationsNotFound() *PostLocationsNotFound {
	return &PostLocationsNotFound{}
}

/*PostLocationsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostLocationsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostLocationsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations][%d] postLocationsNotFound  %+v", 404, o.Payload)
}

func (o *PostLocationsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLocationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLocationsRequestTimeout creates a PostLocationsRequestTimeout with default headers values
func NewPostLocationsRequestTimeout() *PostLocationsRequestTimeout {
	return &PostLocationsRequestTimeout{}
}

/*PostLocationsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostLocationsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostLocationsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations][%d] postLocationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostLocationsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLocationsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLocationsRequestEntityTooLarge creates a PostLocationsRequestEntityTooLarge with default headers values
func NewPostLocationsRequestEntityTooLarge() *PostLocationsRequestEntityTooLarge {
	return &PostLocationsRequestEntityTooLarge{}
}

/*PostLocationsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostLocationsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostLocationsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations][%d] postLocationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLocationsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLocationsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLocationsUnsupportedMediaType creates a PostLocationsUnsupportedMediaType with default headers values
func NewPostLocationsUnsupportedMediaType() *PostLocationsUnsupportedMediaType {
	return &PostLocationsUnsupportedMediaType{}
}

/*PostLocationsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostLocationsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostLocationsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations][%d] postLocationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLocationsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLocationsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLocationsTooManyRequests creates a PostLocationsTooManyRequests with default headers values
func NewPostLocationsTooManyRequests() *PostLocationsTooManyRequests {
	return &PostLocationsTooManyRequests{}
}

/*PostLocationsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostLocationsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostLocationsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations][%d] postLocationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLocationsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLocationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLocationsInternalServerError creates a PostLocationsInternalServerError with default headers values
func NewPostLocationsInternalServerError() *PostLocationsInternalServerError {
	return &PostLocationsInternalServerError{}
}

/*PostLocationsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostLocationsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostLocationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations][%d] postLocationsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLocationsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLocationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLocationsServiceUnavailable creates a PostLocationsServiceUnavailable with default headers values
func NewPostLocationsServiceUnavailable() *PostLocationsServiceUnavailable {
	return &PostLocationsServiceUnavailable{}
}

/*PostLocationsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostLocationsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostLocationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations][%d] postLocationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLocationsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLocationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLocationsGatewayTimeout creates a PostLocationsGatewayTimeout with default headers values
func NewPostLocationsGatewayTimeout() *PostLocationsGatewayTimeout {
	return &PostLocationsGatewayTimeout{}
}

/*PostLocationsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostLocationsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostLocationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations][%d] postLocationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLocationsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLocationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
