// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostAuthorizationDivisionsReader is a Reader for the PostAuthorizationDivisions structure.
type PostAuthorizationDivisionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAuthorizationDivisionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAuthorizationDivisionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAuthorizationDivisionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAuthorizationDivisionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAuthorizationDivisionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAuthorizationDivisionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostAuthorizationDivisionsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostAuthorizationDivisionsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostAuthorizationDivisionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostAuthorizationDivisionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAuthorizationDivisionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAuthorizationDivisionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostAuthorizationDivisionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAuthorizationDivisionsOK creates a PostAuthorizationDivisionsOK with default headers values
func NewPostAuthorizationDivisionsOK() *PostAuthorizationDivisionsOK {
	return &PostAuthorizationDivisionsOK{}
}

/*PostAuthorizationDivisionsOK handles this case with default header values.

successful operation
*/
type PostAuthorizationDivisionsOK struct {
	Payload *models.AuthzDivision
}

func (o *PostAuthorizationDivisionsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/divisions][%d] postAuthorizationDivisionsOK  %+v", 200, o.Payload)
}

func (o *PostAuthorizationDivisionsOK) GetPayload() *models.AuthzDivision {
	return o.Payload
}

func (o *PostAuthorizationDivisionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthzDivision)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationDivisionsBadRequest creates a PostAuthorizationDivisionsBadRequest with default headers values
func NewPostAuthorizationDivisionsBadRequest() *PostAuthorizationDivisionsBadRequest {
	return &PostAuthorizationDivisionsBadRequest{}
}

/*PostAuthorizationDivisionsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostAuthorizationDivisionsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationDivisionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/divisions][%d] postAuthorizationDivisionsBadRequest  %+v", 400, o.Payload)
}

func (o *PostAuthorizationDivisionsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationDivisionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationDivisionsUnauthorized creates a PostAuthorizationDivisionsUnauthorized with default headers values
func NewPostAuthorizationDivisionsUnauthorized() *PostAuthorizationDivisionsUnauthorized {
	return &PostAuthorizationDivisionsUnauthorized{}
}

/*PostAuthorizationDivisionsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostAuthorizationDivisionsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationDivisionsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/divisions][%d] postAuthorizationDivisionsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAuthorizationDivisionsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationDivisionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationDivisionsForbidden creates a PostAuthorizationDivisionsForbidden with default headers values
func NewPostAuthorizationDivisionsForbidden() *PostAuthorizationDivisionsForbidden {
	return &PostAuthorizationDivisionsForbidden{}
}

/*PostAuthorizationDivisionsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostAuthorizationDivisionsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationDivisionsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/divisions][%d] postAuthorizationDivisionsForbidden  %+v", 403, o.Payload)
}

func (o *PostAuthorizationDivisionsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationDivisionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationDivisionsNotFound creates a PostAuthorizationDivisionsNotFound with default headers values
func NewPostAuthorizationDivisionsNotFound() *PostAuthorizationDivisionsNotFound {
	return &PostAuthorizationDivisionsNotFound{}
}

/*PostAuthorizationDivisionsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostAuthorizationDivisionsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationDivisionsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/divisions][%d] postAuthorizationDivisionsNotFound  %+v", 404, o.Payload)
}

func (o *PostAuthorizationDivisionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationDivisionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationDivisionsConflict creates a PostAuthorizationDivisionsConflict with default headers values
func NewPostAuthorizationDivisionsConflict() *PostAuthorizationDivisionsConflict {
	return &PostAuthorizationDivisionsConflict{}
}

/*PostAuthorizationDivisionsConflict handles this case with default header values.

Conflict
*/
type PostAuthorizationDivisionsConflict struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationDivisionsConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/divisions][%d] postAuthorizationDivisionsConflict  %+v", 409, o.Payload)
}

func (o *PostAuthorizationDivisionsConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationDivisionsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationDivisionsRequestEntityTooLarge creates a PostAuthorizationDivisionsRequestEntityTooLarge with default headers values
func NewPostAuthorizationDivisionsRequestEntityTooLarge() *PostAuthorizationDivisionsRequestEntityTooLarge {
	return &PostAuthorizationDivisionsRequestEntityTooLarge{}
}

/*PostAuthorizationDivisionsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostAuthorizationDivisionsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationDivisionsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/divisions][%d] postAuthorizationDivisionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAuthorizationDivisionsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationDivisionsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationDivisionsUnsupportedMediaType creates a PostAuthorizationDivisionsUnsupportedMediaType with default headers values
func NewPostAuthorizationDivisionsUnsupportedMediaType() *PostAuthorizationDivisionsUnsupportedMediaType {
	return &PostAuthorizationDivisionsUnsupportedMediaType{}
}

/*PostAuthorizationDivisionsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostAuthorizationDivisionsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationDivisionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/divisions][%d] postAuthorizationDivisionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAuthorizationDivisionsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationDivisionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationDivisionsTooManyRequests creates a PostAuthorizationDivisionsTooManyRequests with default headers values
func NewPostAuthorizationDivisionsTooManyRequests() *PostAuthorizationDivisionsTooManyRequests {
	return &PostAuthorizationDivisionsTooManyRequests{}
}

/*PostAuthorizationDivisionsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostAuthorizationDivisionsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationDivisionsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/divisions][%d] postAuthorizationDivisionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAuthorizationDivisionsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationDivisionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationDivisionsInternalServerError creates a PostAuthorizationDivisionsInternalServerError with default headers values
func NewPostAuthorizationDivisionsInternalServerError() *PostAuthorizationDivisionsInternalServerError {
	return &PostAuthorizationDivisionsInternalServerError{}
}

/*PostAuthorizationDivisionsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostAuthorizationDivisionsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationDivisionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/divisions][%d] postAuthorizationDivisionsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAuthorizationDivisionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationDivisionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationDivisionsServiceUnavailable creates a PostAuthorizationDivisionsServiceUnavailable with default headers values
func NewPostAuthorizationDivisionsServiceUnavailable() *PostAuthorizationDivisionsServiceUnavailable {
	return &PostAuthorizationDivisionsServiceUnavailable{}
}

/*PostAuthorizationDivisionsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostAuthorizationDivisionsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationDivisionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/divisions][%d] postAuthorizationDivisionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAuthorizationDivisionsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationDivisionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationDivisionsGatewayTimeout creates a PostAuthorizationDivisionsGatewayTimeout with default headers values
func NewPostAuthorizationDivisionsGatewayTimeout() *PostAuthorizationDivisionsGatewayTimeout {
	return &PostAuthorizationDivisionsGatewayTimeout{}
}

/*PostAuthorizationDivisionsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostAuthorizationDivisionsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationDivisionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/divisions][%d] postAuthorizationDivisionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAuthorizationDivisionsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationDivisionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
