// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostRoutingLanguagesReader is a Reader for the PostRoutingLanguages structure.
type PostRoutingLanguagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRoutingLanguagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostRoutingLanguagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRoutingLanguagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostRoutingLanguagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRoutingLanguagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRoutingLanguagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostRoutingLanguagesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostRoutingLanguagesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostRoutingLanguagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostRoutingLanguagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostRoutingLanguagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostRoutingLanguagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostRoutingLanguagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRoutingLanguagesOK creates a PostRoutingLanguagesOK with default headers values
func NewPostRoutingLanguagesOK() *PostRoutingLanguagesOK {
	return &PostRoutingLanguagesOK{}
}

/*PostRoutingLanguagesOK handles this case with default header values.

successful operation
*/
type PostRoutingLanguagesOK struct {
	Payload *models.Language
}

func (o *PostRoutingLanguagesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesOK  %+v", 200, o.Payload)
}

func (o *PostRoutingLanguagesOK) GetPayload() *models.Language {
	return o.Payload
}

func (o *PostRoutingLanguagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Language)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingLanguagesBadRequest creates a PostRoutingLanguagesBadRequest with default headers values
func NewPostRoutingLanguagesBadRequest() *PostRoutingLanguagesBadRequest {
	return &PostRoutingLanguagesBadRequest{}
}

/*PostRoutingLanguagesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostRoutingLanguagesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingLanguagesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesBadRequest  %+v", 400, o.Payload)
}

func (o *PostRoutingLanguagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingLanguagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingLanguagesUnauthorized creates a PostRoutingLanguagesUnauthorized with default headers values
func NewPostRoutingLanguagesUnauthorized() *PostRoutingLanguagesUnauthorized {
	return &PostRoutingLanguagesUnauthorized{}
}

/*PostRoutingLanguagesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostRoutingLanguagesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingLanguagesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRoutingLanguagesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingLanguagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingLanguagesForbidden creates a PostRoutingLanguagesForbidden with default headers values
func NewPostRoutingLanguagesForbidden() *PostRoutingLanguagesForbidden {
	return &PostRoutingLanguagesForbidden{}
}

/*PostRoutingLanguagesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostRoutingLanguagesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingLanguagesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesForbidden  %+v", 403, o.Payload)
}

func (o *PostRoutingLanguagesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingLanguagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingLanguagesNotFound creates a PostRoutingLanguagesNotFound with default headers values
func NewPostRoutingLanguagesNotFound() *PostRoutingLanguagesNotFound {
	return &PostRoutingLanguagesNotFound{}
}

/*PostRoutingLanguagesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostRoutingLanguagesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingLanguagesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesNotFound  %+v", 404, o.Payload)
}

func (o *PostRoutingLanguagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingLanguagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingLanguagesRequestTimeout creates a PostRoutingLanguagesRequestTimeout with default headers values
func NewPostRoutingLanguagesRequestTimeout() *PostRoutingLanguagesRequestTimeout {
	return &PostRoutingLanguagesRequestTimeout{}
}

/*PostRoutingLanguagesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostRoutingLanguagesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingLanguagesRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostRoutingLanguagesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingLanguagesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingLanguagesRequestEntityTooLarge creates a PostRoutingLanguagesRequestEntityTooLarge with default headers values
func NewPostRoutingLanguagesRequestEntityTooLarge() *PostRoutingLanguagesRequestEntityTooLarge {
	return &PostRoutingLanguagesRequestEntityTooLarge{}
}

/*PostRoutingLanguagesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostRoutingLanguagesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingLanguagesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRoutingLanguagesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingLanguagesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingLanguagesUnsupportedMediaType creates a PostRoutingLanguagesUnsupportedMediaType with default headers values
func NewPostRoutingLanguagesUnsupportedMediaType() *PostRoutingLanguagesUnsupportedMediaType {
	return &PostRoutingLanguagesUnsupportedMediaType{}
}

/*PostRoutingLanguagesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostRoutingLanguagesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingLanguagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRoutingLanguagesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingLanguagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingLanguagesTooManyRequests creates a PostRoutingLanguagesTooManyRequests with default headers values
func NewPostRoutingLanguagesTooManyRequests() *PostRoutingLanguagesTooManyRequests {
	return &PostRoutingLanguagesTooManyRequests{}
}

/*PostRoutingLanguagesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostRoutingLanguagesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingLanguagesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRoutingLanguagesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingLanguagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingLanguagesInternalServerError creates a PostRoutingLanguagesInternalServerError with default headers values
func NewPostRoutingLanguagesInternalServerError() *PostRoutingLanguagesInternalServerError {
	return &PostRoutingLanguagesInternalServerError{}
}

/*PostRoutingLanguagesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostRoutingLanguagesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingLanguagesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRoutingLanguagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingLanguagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingLanguagesServiceUnavailable creates a PostRoutingLanguagesServiceUnavailable with default headers values
func NewPostRoutingLanguagesServiceUnavailable() *PostRoutingLanguagesServiceUnavailable {
	return &PostRoutingLanguagesServiceUnavailable{}
}

/*PostRoutingLanguagesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostRoutingLanguagesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingLanguagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRoutingLanguagesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingLanguagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingLanguagesGatewayTimeout creates a PostRoutingLanguagesGatewayTimeout with default headers values
func NewPostRoutingLanguagesGatewayTimeout() *PostRoutingLanguagesGatewayTimeout {
	return &PostRoutingLanguagesGatewayTimeout{}
}

/*PostRoutingLanguagesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostRoutingLanguagesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingLanguagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRoutingLanguagesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingLanguagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
