// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostIntegrationsCredentialsReader is a Reader for the PostIntegrationsCredentials structure.
type PostIntegrationsCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIntegrationsCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostIntegrationsCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostIntegrationsCredentialsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostIntegrationsCredentialsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostIntegrationsCredentialsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostIntegrationsCredentialsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostIntegrationsCredentialsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostIntegrationsCredentialsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostIntegrationsCredentialsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostIntegrationsCredentialsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostIntegrationsCredentialsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostIntegrationsCredentialsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostIntegrationsCredentialsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostIntegrationsCredentialsOK creates a PostIntegrationsCredentialsOK with default headers values
func NewPostIntegrationsCredentialsOK() *PostIntegrationsCredentialsOK {
	return &PostIntegrationsCredentialsOK{}
}

/*PostIntegrationsCredentialsOK handles this case with default header values.

successful operation
*/
type PostIntegrationsCredentialsOK struct {
	Payload *models.CredentialInfo
}

func (o *PostIntegrationsCredentialsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsOK  %+v", 200, o.Payload)
}

func (o *PostIntegrationsCredentialsOK) GetPayload() *models.CredentialInfo {
	return o.Payload
}

func (o *PostIntegrationsCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsCredentialsBadRequest creates a PostIntegrationsCredentialsBadRequest with default headers values
func NewPostIntegrationsCredentialsBadRequest() *PostIntegrationsCredentialsBadRequest {
	return &PostIntegrationsCredentialsBadRequest{}
}

/*PostIntegrationsCredentialsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostIntegrationsCredentialsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsCredentialsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsBadRequest  %+v", 400, o.Payload)
}

func (o *PostIntegrationsCredentialsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsCredentialsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsCredentialsUnauthorized creates a PostIntegrationsCredentialsUnauthorized with default headers values
func NewPostIntegrationsCredentialsUnauthorized() *PostIntegrationsCredentialsUnauthorized {
	return &PostIntegrationsCredentialsUnauthorized{}
}

/*PostIntegrationsCredentialsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostIntegrationsCredentialsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsCredentialsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostIntegrationsCredentialsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsCredentialsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsCredentialsForbidden creates a PostIntegrationsCredentialsForbidden with default headers values
func NewPostIntegrationsCredentialsForbidden() *PostIntegrationsCredentialsForbidden {
	return &PostIntegrationsCredentialsForbidden{}
}

/*PostIntegrationsCredentialsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostIntegrationsCredentialsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsCredentialsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsForbidden  %+v", 403, o.Payload)
}

func (o *PostIntegrationsCredentialsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsCredentialsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsCredentialsNotFound creates a PostIntegrationsCredentialsNotFound with default headers values
func NewPostIntegrationsCredentialsNotFound() *PostIntegrationsCredentialsNotFound {
	return &PostIntegrationsCredentialsNotFound{}
}

/*PostIntegrationsCredentialsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostIntegrationsCredentialsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsCredentialsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsNotFound  %+v", 404, o.Payload)
}

func (o *PostIntegrationsCredentialsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsCredentialsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsCredentialsRequestTimeout creates a PostIntegrationsCredentialsRequestTimeout with default headers values
func NewPostIntegrationsCredentialsRequestTimeout() *PostIntegrationsCredentialsRequestTimeout {
	return &PostIntegrationsCredentialsRequestTimeout{}
}

/*PostIntegrationsCredentialsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostIntegrationsCredentialsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsCredentialsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostIntegrationsCredentialsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsCredentialsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsCredentialsRequestEntityTooLarge creates a PostIntegrationsCredentialsRequestEntityTooLarge with default headers values
func NewPostIntegrationsCredentialsRequestEntityTooLarge() *PostIntegrationsCredentialsRequestEntityTooLarge {
	return &PostIntegrationsCredentialsRequestEntityTooLarge{}
}

/*PostIntegrationsCredentialsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostIntegrationsCredentialsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsCredentialsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostIntegrationsCredentialsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsCredentialsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsCredentialsUnsupportedMediaType creates a PostIntegrationsCredentialsUnsupportedMediaType with default headers values
func NewPostIntegrationsCredentialsUnsupportedMediaType() *PostIntegrationsCredentialsUnsupportedMediaType {
	return &PostIntegrationsCredentialsUnsupportedMediaType{}
}

/*PostIntegrationsCredentialsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostIntegrationsCredentialsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsCredentialsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostIntegrationsCredentialsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsCredentialsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsCredentialsTooManyRequests creates a PostIntegrationsCredentialsTooManyRequests with default headers values
func NewPostIntegrationsCredentialsTooManyRequests() *PostIntegrationsCredentialsTooManyRequests {
	return &PostIntegrationsCredentialsTooManyRequests{}
}

/*PostIntegrationsCredentialsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostIntegrationsCredentialsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsCredentialsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostIntegrationsCredentialsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsCredentialsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsCredentialsInternalServerError creates a PostIntegrationsCredentialsInternalServerError with default headers values
func NewPostIntegrationsCredentialsInternalServerError() *PostIntegrationsCredentialsInternalServerError {
	return &PostIntegrationsCredentialsInternalServerError{}
}

/*PostIntegrationsCredentialsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostIntegrationsCredentialsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsCredentialsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostIntegrationsCredentialsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsCredentialsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsCredentialsServiceUnavailable creates a PostIntegrationsCredentialsServiceUnavailable with default headers values
func NewPostIntegrationsCredentialsServiceUnavailable() *PostIntegrationsCredentialsServiceUnavailable {
	return &PostIntegrationsCredentialsServiceUnavailable{}
}

/*PostIntegrationsCredentialsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostIntegrationsCredentialsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsCredentialsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostIntegrationsCredentialsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsCredentialsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsCredentialsGatewayTimeout creates a PostIntegrationsCredentialsGatewayTimeout with default headers values
func NewPostIntegrationsCredentialsGatewayTimeout() *PostIntegrationsCredentialsGatewayTimeout {
	return &PostIntegrationsCredentialsGatewayTimeout{}
}

/*PostIntegrationsCredentialsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostIntegrationsCredentialsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsCredentialsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostIntegrationsCredentialsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsCredentialsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
