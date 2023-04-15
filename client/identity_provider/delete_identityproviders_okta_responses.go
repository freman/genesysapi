// Code generated by go-swagger; DO NOT EDIT.

package identity_provider

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteIdentityprovidersOktaReader is a Reader for the DeleteIdentityprovidersOkta structure.
type DeleteIdentityprovidersOktaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIdentityprovidersOktaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIdentityprovidersOktaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteIdentityprovidersOktaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteIdentityprovidersOktaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteIdentityprovidersOktaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteIdentityprovidersOktaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteIdentityprovidersOktaRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteIdentityprovidersOktaRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteIdentityprovidersOktaUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteIdentityprovidersOktaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteIdentityprovidersOktaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteIdentityprovidersOktaServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteIdentityprovidersOktaGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteIdentityprovidersOktaOK creates a DeleteIdentityprovidersOktaOK with default headers values
func NewDeleteIdentityprovidersOktaOK() *DeleteIdentityprovidersOktaOK {
	return &DeleteIdentityprovidersOktaOK{}
}

/*
DeleteIdentityprovidersOktaOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteIdentityprovidersOktaOK struct {
	Payload models.Empty
}

// IsSuccess returns true when this delete identityproviders okta o k response has a 2xx status code
func (o *DeleteIdentityprovidersOktaOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete identityproviders okta o k response has a 3xx status code
func (o *DeleteIdentityprovidersOktaOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders okta o k response has a 4xx status code
func (o *DeleteIdentityprovidersOktaOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete identityproviders okta o k response has a 5xx status code
func (o *DeleteIdentityprovidersOktaOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders okta o k response a status code equal to that given
func (o *DeleteIdentityprovidersOktaOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteIdentityprovidersOktaOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaOK  %+v", 200, o.Payload)
}

func (o *DeleteIdentityprovidersOktaOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaOK  %+v", 200, o.Payload)
}

func (o *DeleteIdentityprovidersOktaOK) GetPayload() models.Empty {
	return o.Payload
}

func (o *DeleteIdentityprovidersOktaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersOktaBadRequest creates a DeleteIdentityprovidersOktaBadRequest with default headers values
func NewDeleteIdentityprovidersOktaBadRequest() *DeleteIdentityprovidersOktaBadRequest {
	return &DeleteIdentityprovidersOktaBadRequest{}
}

/*
DeleteIdentityprovidersOktaBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteIdentityprovidersOktaBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders okta bad request response has a 2xx status code
func (o *DeleteIdentityprovidersOktaBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders okta bad request response has a 3xx status code
func (o *DeleteIdentityprovidersOktaBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders okta bad request response has a 4xx status code
func (o *DeleteIdentityprovidersOktaBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders okta bad request response has a 5xx status code
func (o *DeleteIdentityprovidersOktaBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders okta bad request response a status code equal to that given
func (o *DeleteIdentityprovidersOktaBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteIdentityprovidersOktaBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteIdentityprovidersOktaBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteIdentityprovidersOktaBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersOktaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersOktaUnauthorized creates a DeleteIdentityprovidersOktaUnauthorized with default headers values
func NewDeleteIdentityprovidersOktaUnauthorized() *DeleteIdentityprovidersOktaUnauthorized {
	return &DeleteIdentityprovidersOktaUnauthorized{}
}

/*
DeleteIdentityprovidersOktaUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteIdentityprovidersOktaUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders okta unauthorized response has a 2xx status code
func (o *DeleteIdentityprovidersOktaUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders okta unauthorized response has a 3xx status code
func (o *DeleteIdentityprovidersOktaUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders okta unauthorized response has a 4xx status code
func (o *DeleteIdentityprovidersOktaUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders okta unauthorized response has a 5xx status code
func (o *DeleteIdentityprovidersOktaUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders okta unauthorized response a status code equal to that given
func (o *DeleteIdentityprovidersOktaUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteIdentityprovidersOktaUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteIdentityprovidersOktaUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteIdentityprovidersOktaUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersOktaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersOktaForbidden creates a DeleteIdentityprovidersOktaForbidden with default headers values
func NewDeleteIdentityprovidersOktaForbidden() *DeleteIdentityprovidersOktaForbidden {
	return &DeleteIdentityprovidersOktaForbidden{}
}

/*
DeleteIdentityprovidersOktaForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteIdentityprovidersOktaForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders okta forbidden response has a 2xx status code
func (o *DeleteIdentityprovidersOktaForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders okta forbidden response has a 3xx status code
func (o *DeleteIdentityprovidersOktaForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders okta forbidden response has a 4xx status code
func (o *DeleteIdentityprovidersOktaForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders okta forbidden response has a 5xx status code
func (o *DeleteIdentityprovidersOktaForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders okta forbidden response a status code equal to that given
func (o *DeleteIdentityprovidersOktaForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteIdentityprovidersOktaForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaForbidden  %+v", 403, o.Payload)
}

func (o *DeleteIdentityprovidersOktaForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaForbidden  %+v", 403, o.Payload)
}

func (o *DeleteIdentityprovidersOktaForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersOktaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersOktaNotFound creates a DeleteIdentityprovidersOktaNotFound with default headers values
func NewDeleteIdentityprovidersOktaNotFound() *DeleteIdentityprovidersOktaNotFound {
	return &DeleteIdentityprovidersOktaNotFound{}
}

/*
DeleteIdentityprovidersOktaNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteIdentityprovidersOktaNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders okta not found response has a 2xx status code
func (o *DeleteIdentityprovidersOktaNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders okta not found response has a 3xx status code
func (o *DeleteIdentityprovidersOktaNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders okta not found response has a 4xx status code
func (o *DeleteIdentityprovidersOktaNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders okta not found response has a 5xx status code
func (o *DeleteIdentityprovidersOktaNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders okta not found response a status code equal to that given
func (o *DeleteIdentityprovidersOktaNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteIdentityprovidersOktaNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIdentityprovidersOktaNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIdentityprovidersOktaNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersOktaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersOktaRequestTimeout creates a DeleteIdentityprovidersOktaRequestTimeout with default headers values
func NewDeleteIdentityprovidersOktaRequestTimeout() *DeleteIdentityprovidersOktaRequestTimeout {
	return &DeleteIdentityprovidersOktaRequestTimeout{}
}

/*
DeleteIdentityprovidersOktaRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteIdentityprovidersOktaRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders okta request timeout response has a 2xx status code
func (o *DeleteIdentityprovidersOktaRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders okta request timeout response has a 3xx status code
func (o *DeleteIdentityprovidersOktaRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders okta request timeout response has a 4xx status code
func (o *DeleteIdentityprovidersOktaRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders okta request timeout response has a 5xx status code
func (o *DeleteIdentityprovidersOktaRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders okta request timeout response a status code equal to that given
func (o *DeleteIdentityprovidersOktaRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteIdentityprovidersOktaRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteIdentityprovidersOktaRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteIdentityprovidersOktaRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersOktaRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersOktaRequestEntityTooLarge creates a DeleteIdentityprovidersOktaRequestEntityTooLarge with default headers values
func NewDeleteIdentityprovidersOktaRequestEntityTooLarge() *DeleteIdentityprovidersOktaRequestEntityTooLarge {
	return &DeleteIdentityprovidersOktaRequestEntityTooLarge{}
}

/*
DeleteIdentityprovidersOktaRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteIdentityprovidersOktaRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders okta request entity too large response has a 2xx status code
func (o *DeleteIdentityprovidersOktaRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders okta request entity too large response has a 3xx status code
func (o *DeleteIdentityprovidersOktaRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders okta request entity too large response has a 4xx status code
func (o *DeleteIdentityprovidersOktaRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders okta request entity too large response has a 5xx status code
func (o *DeleteIdentityprovidersOktaRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders okta request entity too large response a status code equal to that given
func (o *DeleteIdentityprovidersOktaRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteIdentityprovidersOktaRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteIdentityprovidersOktaRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteIdentityprovidersOktaRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersOktaRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersOktaUnsupportedMediaType creates a DeleteIdentityprovidersOktaUnsupportedMediaType with default headers values
func NewDeleteIdentityprovidersOktaUnsupportedMediaType() *DeleteIdentityprovidersOktaUnsupportedMediaType {
	return &DeleteIdentityprovidersOktaUnsupportedMediaType{}
}

/*
DeleteIdentityprovidersOktaUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteIdentityprovidersOktaUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders okta unsupported media type response has a 2xx status code
func (o *DeleteIdentityprovidersOktaUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders okta unsupported media type response has a 3xx status code
func (o *DeleteIdentityprovidersOktaUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders okta unsupported media type response has a 4xx status code
func (o *DeleteIdentityprovidersOktaUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders okta unsupported media type response has a 5xx status code
func (o *DeleteIdentityprovidersOktaUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders okta unsupported media type response a status code equal to that given
func (o *DeleteIdentityprovidersOktaUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteIdentityprovidersOktaUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteIdentityprovidersOktaUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteIdentityprovidersOktaUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersOktaUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersOktaTooManyRequests creates a DeleteIdentityprovidersOktaTooManyRequests with default headers values
func NewDeleteIdentityprovidersOktaTooManyRequests() *DeleteIdentityprovidersOktaTooManyRequests {
	return &DeleteIdentityprovidersOktaTooManyRequests{}
}

/*
DeleteIdentityprovidersOktaTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteIdentityprovidersOktaTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders okta too many requests response has a 2xx status code
func (o *DeleteIdentityprovidersOktaTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders okta too many requests response has a 3xx status code
func (o *DeleteIdentityprovidersOktaTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders okta too many requests response has a 4xx status code
func (o *DeleteIdentityprovidersOktaTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders okta too many requests response has a 5xx status code
func (o *DeleteIdentityprovidersOktaTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders okta too many requests response a status code equal to that given
func (o *DeleteIdentityprovidersOktaTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteIdentityprovidersOktaTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteIdentityprovidersOktaTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteIdentityprovidersOktaTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersOktaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersOktaInternalServerError creates a DeleteIdentityprovidersOktaInternalServerError with default headers values
func NewDeleteIdentityprovidersOktaInternalServerError() *DeleteIdentityprovidersOktaInternalServerError {
	return &DeleteIdentityprovidersOktaInternalServerError{}
}

/*
DeleteIdentityprovidersOktaInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteIdentityprovidersOktaInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders okta internal server error response has a 2xx status code
func (o *DeleteIdentityprovidersOktaInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders okta internal server error response has a 3xx status code
func (o *DeleteIdentityprovidersOktaInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders okta internal server error response has a 4xx status code
func (o *DeleteIdentityprovidersOktaInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete identityproviders okta internal server error response has a 5xx status code
func (o *DeleteIdentityprovidersOktaInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete identityproviders okta internal server error response a status code equal to that given
func (o *DeleteIdentityprovidersOktaInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteIdentityprovidersOktaInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteIdentityprovidersOktaInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteIdentityprovidersOktaInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersOktaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersOktaServiceUnavailable creates a DeleteIdentityprovidersOktaServiceUnavailable with default headers values
func NewDeleteIdentityprovidersOktaServiceUnavailable() *DeleteIdentityprovidersOktaServiceUnavailable {
	return &DeleteIdentityprovidersOktaServiceUnavailable{}
}

/*
DeleteIdentityprovidersOktaServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteIdentityprovidersOktaServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders okta service unavailable response has a 2xx status code
func (o *DeleteIdentityprovidersOktaServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders okta service unavailable response has a 3xx status code
func (o *DeleteIdentityprovidersOktaServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders okta service unavailable response has a 4xx status code
func (o *DeleteIdentityprovidersOktaServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete identityproviders okta service unavailable response has a 5xx status code
func (o *DeleteIdentityprovidersOktaServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete identityproviders okta service unavailable response a status code equal to that given
func (o *DeleteIdentityprovidersOktaServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteIdentityprovidersOktaServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteIdentityprovidersOktaServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteIdentityprovidersOktaServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersOktaServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersOktaGatewayTimeout creates a DeleteIdentityprovidersOktaGatewayTimeout with default headers values
func NewDeleteIdentityprovidersOktaGatewayTimeout() *DeleteIdentityprovidersOktaGatewayTimeout {
	return &DeleteIdentityprovidersOktaGatewayTimeout{}
}

/*
DeleteIdentityprovidersOktaGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteIdentityprovidersOktaGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders okta gateway timeout response has a 2xx status code
func (o *DeleteIdentityprovidersOktaGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders okta gateway timeout response has a 3xx status code
func (o *DeleteIdentityprovidersOktaGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders okta gateway timeout response has a 4xx status code
func (o *DeleteIdentityprovidersOktaGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete identityproviders okta gateway timeout response has a 5xx status code
func (o *DeleteIdentityprovidersOktaGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete identityproviders okta gateway timeout response a status code equal to that given
func (o *DeleteIdentityprovidersOktaGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteIdentityprovidersOktaGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteIdentityprovidersOktaGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/okta][%d] deleteIdentityprovidersOktaGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteIdentityprovidersOktaGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersOktaGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
