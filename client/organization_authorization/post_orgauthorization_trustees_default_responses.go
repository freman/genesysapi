// Code generated by go-swagger; DO NOT EDIT.

package organization_authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostOrgauthorizationTrusteesDefaultReader is a Reader for the PostOrgauthorizationTrusteesDefault structure.
type PostOrgauthorizationTrusteesDefaultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOrgauthorizationTrusteesDefaultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOrgauthorizationTrusteesDefaultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOrgauthorizationTrusteesDefaultBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostOrgauthorizationTrusteesDefaultUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostOrgauthorizationTrusteesDefaultForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostOrgauthorizationTrusteesDefaultNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostOrgauthorizationTrusteesDefaultRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostOrgauthorizationTrusteesDefaultRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostOrgauthorizationTrusteesDefaultUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostOrgauthorizationTrusteesDefaultTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOrgauthorizationTrusteesDefaultInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostOrgauthorizationTrusteesDefaultServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostOrgauthorizationTrusteesDefaultGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOrgauthorizationTrusteesDefaultOK creates a PostOrgauthorizationTrusteesDefaultOK with default headers values
func NewPostOrgauthorizationTrusteesDefaultOK() *PostOrgauthorizationTrusteesDefaultOK {
	return &PostOrgauthorizationTrusteesDefaultOK{}
}

/*PostOrgauthorizationTrusteesDefaultOK handles this case with default header values.

successful operation
*/
type PostOrgauthorizationTrusteesDefaultOK struct {
	Payload *models.Trustee
}

func (o *PostOrgauthorizationTrusteesDefaultOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/default][%d] postOrgauthorizationTrusteesDefaultOK  %+v", 200, o.Payload)
}

func (o *PostOrgauthorizationTrusteesDefaultOK) GetPayload() *models.Trustee {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteesDefaultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Trustee)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteesDefaultBadRequest creates a PostOrgauthorizationTrusteesDefaultBadRequest with default headers values
func NewPostOrgauthorizationTrusteesDefaultBadRequest() *PostOrgauthorizationTrusteesDefaultBadRequest {
	return &PostOrgauthorizationTrusteesDefaultBadRequest{}
}

/*PostOrgauthorizationTrusteesDefaultBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostOrgauthorizationTrusteesDefaultBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrusteesDefaultBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/default][%d] postOrgauthorizationTrusteesDefaultBadRequest  %+v", 400, o.Payload)
}

func (o *PostOrgauthorizationTrusteesDefaultBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteesDefaultBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteesDefaultUnauthorized creates a PostOrgauthorizationTrusteesDefaultUnauthorized with default headers values
func NewPostOrgauthorizationTrusteesDefaultUnauthorized() *PostOrgauthorizationTrusteesDefaultUnauthorized {
	return &PostOrgauthorizationTrusteesDefaultUnauthorized{}
}

/*PostOrgauthorizationTrusteesDefaultUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostOrgauthorizationTrusteesDefaultUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrusteesDefaultUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/default][%d] postOrgauthorizationTrusteesDefaultUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOrgauthorizationTrusteesDefaultUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteesDefaultUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteesDefaultForbidden creates a PostOrgauthorizationTrusteesDefaultForbidden with default headers values
func NewPostOrgauthorizationTrusteesDefaultForbidden() *PostOrgauthorizationTrusteesDefaultForbidden {
	return &PostOrgauthorizationTrusteesDefaultForbidden{}
}

/*PostOrgauthorizationTrusteesDefaultForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostOrgauthorizationTrusteesDefaultForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrusteesDefaultForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/default][%d] postOrgauthorizationTrusteesDefaultForbidden  %+v", 403, o.Payload)
}

func (o *PostOrgauthorizationTrusteesDefaultForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteesDefaultForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteesDefaultNotFound creates a PostOrgauthorizationTrusteesDefaultNotFound with default headers values
func NewPostOrgauthorizationTrusteesDefaultNotFound() *PostOrgauthorizationTrusteesDefaultNotFound {
	return &PostOrgauthorizationTrusteesDefaultNotFound{}
}

/*PostOrgauthorizationTrusteesDefaultNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostOrgauthorizationTrusteesDefaultNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrusteesDefaultNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/default][%d] postOrgauthorizationTrusteesDefaultNotFound  %+v", 404, o.Payload)
}

func (o *PostOrgauthorizationTrusteesDefaultNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteesDefaultNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteesDefaultRequestTimeout creates a PostOrgauthorizationTrusteesDefaultRequestTimeout with default headers values
func NewPostOrgauthorizationTrusteesDefaultRequestTimeout() *PostOrgauthorizationTrusteesDefaultRequestTimeout {
	return &PostOrgauthorizationTrusteesDefaultRequestTimeout{}
}

/*PostOrgauthorizationTrusteesDefaultRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostOrgauthorizationTrusteesDefaultRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrusteesDefaultRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/default][%d] postOrgauthorizationTrusteesDefaultRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOrgauthorizationTrusteesDefaultRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteesDefaultRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteesDefaultRequestEntityTooLarge creates a PostOrgauthorizationTrusteesDefaultRequestEntityTooLarge with default headers values
func NewPostOrgauthorizationTrusteesDefaultRequestEntityTooLarge() *PostOrgauthorizationTrusteesDefaultRequestEntityTooLarge {
	return &PostOrgauthorizationTrusteesDefaultRequestEntityTooLarge{}
}

/*PostOrgauthorizationTrusteesDefaultRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostOrgauthorizationTrusteesDefaultRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrusteesDefaultRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/default][%d] postOrgauthorizationTrusteesDefaultRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOrgauthorizationTrusteesDefaultRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteesDefaultRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteesDefaultUnsupportedMediaType creates a PostOrgauthorizationTrusteesDefaultUnsupportedMediaType with default headers values
func NewPostOrgauthorizationTrusteesDefaultUnsupportedMediaType() *PostOrgauthorizationTrusteesDefaultUnsupportedMediaType {
	return &PostOrgauthorizationTrusteesDefaultUnsupportedMediaType{}
}

/*PostOrgauthorizationTrusteesDefaultUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostOrgauthorizationTrusteesDefaultUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrusteesDefaultUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/default][%d] postOrgauthorizationTrusteesDefaultUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOrgauthorizationTrusteesDefaultUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteesDefaultUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteesDefaultTooManyRequests creates a PostOrgauthorizationTrusteesDefaultTooManyRequests with default headers values
func NewPostOrgauthorizationTrusteesDefaultTooManyRequests() *PostOrgauthorizationTrusteesDefaultTooManyRequests {
	return &PostOrgauthorizationTrusteesDefaultTooManyRequests{}
}

/*PostOrgauthorizationTrusteesDefaultTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostOrgauthorizationTrusteesDefaultTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrusteesDefaultTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/default][%d] postOrgauthorizationTrusteesDefaultTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOrgauthorizationTrusteesDefaultTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteesDefaultTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteesDefaultInternalServerError creates a PostOrgauthorizationTrusteesDefaultInternalServerError with default headers values
func NewPostOrgauthorizationTrusteesDefaultInternalServerError() *PostOrgauthorizationTrusteesDefaultInternalServerError {
	return &PostOrgauthorizationTrusteesDefaultInternalServerError{}
}

/*PostOrgauthorizationTrusteesDefaultInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostOrgauthorizationTrusteesDefaultInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrusteesDefaultInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/default][%d] postOrgauthorizationTrusteesDefaultInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOrgauthorizationTrusteesDefaultInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteesDefaultInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteesDefaultServiceUnavailable creates a PostOrgauthorizationTrusteesDefaultServiceUnavailable with default headers values
func NewPostOrgauthorizationTrusteesDefaultServiceUnavailable() *PostOrgauthorizationTrusteesDefaultServiceUnavailable {
	return &PostOrgauthorizationTrusteesDefaultServiceUnavailable{}
}

/*PostOrgauthorizationTrusteesDefaultServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostOrgauthorizationTrusteesDefaultServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrusteesDefaultServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/default][%d] postOrgauthorizationTrusteesDefaultServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOrgauthorizationTrusteesDefaultServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteesDefaultServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteesDefaultGatewayTimeout creates a PostOrgauthorizationTrusteesDefaultGatewayTimeout with default headers values
func NewPostOrgauthorizationTrusteesDefaultGatewayTimeout() *PostOrgauthorizationTrusteesDefaultGatewayTimeout {
	return &PostOrgauthorizationTrusteesDefaultGatewayTimeout{}
}

/*PostOrgauthorizationTrusteesDefaultGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostOrgauthorizationTrusteesDefaultGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrusteesDefaultGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/default][%d] postOrgauthorizationTrusteesDefaultGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOrgauthorizationTrusteesDefaultGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteesDefaultGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
