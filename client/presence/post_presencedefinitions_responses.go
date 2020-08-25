// Code generated by go-swagger; DO NOT EDIT.

package presence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostPresencedefinitionsReader is a Reader for the PostPresencedefinitions structure.
type PostPresencedefinitionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPresencedefinitionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostPresencedefinitionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostPresencedefinitionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostPresencedefinitionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostPresencedefinitionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostPresencedefinitionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostPresencedefinitionsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostPresencedefinitionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostPresencedefinitionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostPresencedefinitionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostPresencedefinitionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostPresencedefinitionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostPresencedefinitionsOK creates a PostPresencedefinitionsOK with default headers values
func NewPostPresencedefinitionsOK() *PostPresencedefinitionsOK {
	return &PostPresencedefinitionsOK{}
}

/*PostPresencedefinitionsOK handles this case with default header values.

successful operation
*/
type PostPresencedefinitionsOK struct {
	Payload *models.OrganizationPresence
}

func (o *PostPresencedefinitionsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/presencedefinitions][%d] postPresencedefinitionsOK  %+v", 200, o.Payload)
}

func (o *PostPresencedefinitionsOK) GetPayload() *models.OrganizationPresence {
	return o.Payload
}

func (o *PostPresencedefinitionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationPresence)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPresencedefinitionsBadRequest creates a PostPresencedefinitionsBadRequest with default headers values
func NewPostPresencedefinitionsBadRequest() *PostPresencedefinitionsBadRequest {
	return &PostPresencedefinitionsBadRequest{}
}

/*PostPresencedefinitionsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostPresencedefinitionsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostPresencedefinitionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/presencedefinitions][%d] postPresencedefinitionsBadRequest  %+v", 400, o.Payload)
}

func (o *PostPresencedefinitionsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostPresencedefinitionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPresencedefinitionsUnauthorized creates a PostPresencedefinitionsUnauthorized with default headers values
func NewPostPresencedefinitionsUnauthorized() *PostPresencedefinitionsUnauthorized {
	return &PostPresencedefinitionsUnauthorized{}
}

/*PostPresencedefinitionsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostPresencedefinitionsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostPresencedefinitionsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/presencedefinitions][%d] postPresencedefinitionsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostPresencedefinitionsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostPresencedefinitionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPresencedefinitionsForbidden creates a PostPresencedefinitionsForbidden with default headers values
func NewPostPresencedefinitionsForbidden() *PostPresencedefinitionsForbidden {
	return &PostPresencedefinitionsForbidden{}
}

/*PostPresencedefinitionsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostPresencedefinitionsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostPresencedefinitionsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/presencedefinitions][%d] postPresencedefinitionsForbidden  %+v", 403, o.Payload)
}

func (o *PostPresencedefinitionsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostPresencedefinitionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPresencedefinitionsNotFound creates a PostPresencedefinitionsNotFound with default headers values
func NewPostPresencedefinitionsNotFound() *PostPresencedefinitionsNotFound {
	return &PostPresencedefinitionsNotFound{}
}

/*PostPresencedefinitionsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostPresencedefinitionsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostPresencedefinitionsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/presencedefinitions][%d] postPresencedefinitionsNotFound  %+v", 404, o.Payload)
}

func (o *PostPresencedefinitionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostPresencedefinitionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPresencedefinitionsRequestEntityTooLarge creates a PostPresencedefinitionsRequestEntityTooLarge with default headers values
func NewPostPresencedefinitionsRequestEntityTooLarge() *PostPresencedefinitionsRequestEntityTooLarge {
	return &PostPresencedefinitionsRequestEntityTooLarge{}
}

/*PostPresencedefinitionsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostPresencedefinitionsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostPresencedefinitionsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/presencedefinitions][%d] postPresencedefinitionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostPresencedefinitionsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostPresencedefinitionsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPresencedefinitionsUnsupportedMediaType creates a PostPresencedefinitionsUnsupportedMediaType with default headers values
func NewPostPresencedefinitionsUnsupportedMediaType() *PostPresencedefinitionsUnsupportedMediaType {
	return &PostPresencedefinitionsUnsupportedMediaType{}
}

/*PostPresencedefinitionsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostPresencedefinitionsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostPresencedefinitionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/presencedefinitions][%d] postPresencedefinitionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostPresencedefinitionsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostPresencedefinitionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPresencedefinitionsTooManyRequests creates a PostPresencedefinitionsTooManyRequests with default headers values
func NewPostPresencedefinitionsTooManyRequests() *PostPresencedefinitionsTooManyRequests {
	return &PostPresencedefinitionsTooManyRequests{}
}

/*PostPresencedefinitionsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostPresencedefinitionsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostPresencedefinitionsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/presencedefinitions][%d] postPresencedefinitionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostPresencedefinitionsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostPresencedefinitionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPresencedefinitionsInternalServerError creates a PostPresencedefinitionsInternalServerError with default headers values
func NewPostPresencedefinitionsInternalServerError() *PostPresencedefinitionsInternalServerError {
	return &PostPresencedefinitionsInternalServerError{}
}

/*PostPresencedefinitionsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostPresencedefinitionsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostPresencedefinitionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/presencedefinitions][%d] postPresencedefinitionsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostPresencedefinitionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostPresencedefinitionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPresencedefinitionsServiceUnavailable creates a PostPresencedefinitionsServiceUnavailable with default headers values
func NewPostPresencedefinitionsServiceUnavailable() *PostPresencedefinitionsServiceUnavailable {
	return &PostPresencedefinitionsServiceUnavailable{}
}

/*PostPresencedefinitionsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostPresencedefinitionsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostPresencedefinitionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/presencedefinitions][%d] postPresencedefinitionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostPresencedefinitionsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostPresencedefinitionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPresencedefinitionsGatewayTimeout creates a PostPresencedefinitionsGatewayTimeout with default headers values
func NewPostPresencedefinitionsGatewayTimeout() *PostPresencedefinitionsGatewayTimeout {
	return &PostPresencedefinitionsGatewayTimeout{}
}

/*PostPresencedefinitionsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostPresencedefinitionsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostPresencedefinitionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/presencedefinitions][%d] postPresencedefinitionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostPresencedefinitionsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostPresencedefinitionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}