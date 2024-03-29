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

// DeleteIdentityprovidersCicReader is a Reader for the DeleteIdentityprovidersCic structure.
type DeleteIdentityprovidersCicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIdentityprovidersCicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIdentityprovidersCicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteIdentityprovidersCicBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteIdentityprovidersCicUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteIdentityprovidersCicForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteIdentityprovidersCicNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteIdentityprovidersCicRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteIdentityprovidersCicRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteIdentityprovidersCicUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteIdentityprovidersCicTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteIdentityprovidersCicInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteIdentityprovidersCicServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteIdentityprovidersCicGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteIdentityprovidersCicOK creates a DeleteIdentityprovidersCicOK with default headers values
func NewDeleteIdentityprovidersCicOK() *DeleteIdentityprovidersCicOK {
	return &DeleteIdentityprovidersCicOK{}
}

/*
DeleteIdentityprovidersCicOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteIdentityprovidersCicOK struct {
	Payload models.Empty
}

// IsSuccess returns true when this delete identityproviders cic o k response has a 2xx status code
func (o *DeleteIdentityprovidersCicOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete identityproviders cic o k response has a 3xx status code
func (o *DeleteIdentityprovidersCicOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders cic o k response has a 4xx status code
func (o *DeleteIdentityprovidersCicOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete identityproviders cic o k response has a 5xx status code
func (o *DeleteIdentityprovidersCicOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders cic o k response a status code equal to that given
func (o *DeleteIdentityprovidersCicOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteIdentityprovidersCicOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicOK  %+v", 200, o.Payload)
}

func (o *DeleteIdentityprovidersCicOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicOK  %+v", 200, o.Payload)
}

func (o *DeleteIdentityprovidersCicOK) GetPayload() models.Empty {
	return o.Payload
}

func (o *DeleteIdentityprovidersCicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersCicBadRequest creates a DeleteIdentityprovidersCicBadRequest with default headers values
func NewDeleteIdentityprovidersCicBadRequest() *DeleteIdentityprovidersCicBadRequest {
	return &DeleteIdentityprovidersCicBadRequest{}
}

/*
DeleteIdentityprovidersCicBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteIdentityprovidersCicBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders cic bad request response has a 2xx status code
func (o *DeleteIdentityprovidersCicBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders cic bad request response has a 3xx status code
func (o *DeleteIdentityprovidersCicBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders cic bad request response has a 4xx status code
func (o *DeleteIdentityprovidersCicBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders cic bad request response has a 5xx status code
func (o *DeleteIdentityprovidersCicBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders cic bad request response a status code equal to that given
func (o *DeleteIdentityprovidersCicBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteIdentityprovidersCicBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteIdentityprovidersCicBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteIdentityprovidersCicBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersCicBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersCicUnauthorized creates a DeleteIdentityprovidersCicUnauthorized with default headers values
func NewDeleteIdentityprovidersCicUnauthorized() *DeleteIdentityprovidersCicUnauthorized {
	return &DeleteIdentityprovidersCicUnauthorized{}
}

/*
DeleteIdentityprovidersCicUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteIdentityprovidersCicUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders cic unauthorized response has a 2xx status code
func (o *DeleteIdentityprovidersCicUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders cic unauthorized response has a 3xx status code
func (o *DeleteIdentityprovidersCicUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders cic unauthorized response has a 4xx status code
func (o *DeleteIdentityprovidersCicUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders cic unauthorized response has a 5xx status code
func (o *DeleteIdentityprovidersCicUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders cic unauthorized response a status code equal to that given
func (o *DeleteIdentityprovidersCicUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteIdentityprovidersCicUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteIdentityprovidersCicUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteIdentityprovidersCicUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersCicUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersCicForbidden creates a DeleteIdentityprovidersCicForbidden with default headers values
func NewDeleteIdentityprovidersCicForbidden() *DeleteIdentityprovidersCicForbidden {
	return &DeleteIdentityprovidersCicForbidden{}
}

/*
DeleteIdentityprovidersCicForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteIdentityprovidersCicForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders cic forbidden response has a 2xx status code
func (o *DeleteIdentityprovidersCicForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders cic forbidden response has a 3xx status code
func (o *DeleteIdentityprovidersCicForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders cic forbidden response has a 4xx status code
func (o *DeleteIdentityprovidersCicForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders cic forbidden response has a 5xx status code
func (o *DeleteIdentityprovidersCicForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders cic forbidden response a status code equal to that given
func (o *DeleteIdentityprovidersCicForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteIdentityprovidersCicForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicForbidden  %+v", 403, o.Payload)
}

func (o *DeleteIdentityprovidersCicForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicForbidden  %+v", 403, o.Payload)
}

func (o *DeleteIdentityprovidersCicForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersCicForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersCicNotFound creates a DeleteIdentityprovidersCicNotFound with default headers values
func NewDeleteIdentityprovidersCicNotFound() *DeleteIdentityprovidersCicNotFound {
	return &DeleteIdentityprovidersCicNotFound{}
}

/*
DeleteIdentityprovidersCicNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteIdentityprovidersCicNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders cic not found response has a 2xx status code
func (o *DeleteIdentityprovidersCicNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders cic not found response has a 3xx status code
func (o *DeleteIdentityprovidersCicNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders cic not found response has a 4xx status code
func (o *DeleteIdentityprovidersCicNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders cic not found response has a 5xx status code
func (o *DeleteIdentityprovidersCicNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders cic not found response a status code equal to that given
func (o *DeleteIdentityprovidersCicNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteIdentityprovidersCicNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIdentityprovidersCicNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIdentityprovidersCicNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersCicNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersCicRequestTimeout creates a DeleteIdentityprovidersCicRequestTimeout with default headers values
func NewDeleteIdentityprovidersCicRequestTimeout() *DeleteIdentityprovidersCicRequestTimeout {
	return &DeleteIdentityprovidersCicRequestTimeout{}
}

/*
DeleteIdentityprovidersCicRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteIdentityprovidersCicRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders cic request timeout response has a 2xx status code
func (o *DeleteIdentityprovidersCicRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders cic request timeout response has a 3xx status code
func (o *DeleteIdentityprovidersCicRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders cic request timeout response has a 4xx status code
func (o *DeleteIdentityprovidersCicRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders cic request timeout response has a 5xx status code
func (o *DeleteIdentityprovidersCicRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders cic request timeout response a status code equal to that given
func (o *DeleteIdentityprovidersCicRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteIdentityprovidersCicRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteIdentityprovidersCicRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteIdentityprovidersCicRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersCicRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersCicRequestEntityTooLarge creates a DeleteIdentityprovidersCicRequestEntityTooLarge with default headers values
func NewDeleteIdentityprovidersCicRequestEntityTooLarge() *DeleteIdentityprovidersCicRequestEntityTooLarge {
	return &DeleteIdentityprovidersCicRequestEntityTooLarge{}
}

/*
DeleteIdentityprovidersCicRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteIdentityprovidersCicRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders cic request entity too large response has a 2xx status code
func (o *DeleteIdentityprovidersCicRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders cic request entity too large response has a 3xx status code
func (o *DeleteIdentityprovidersCicRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders cic request entity too large response has a 4xx status code
func (o *DeleteIdentityprovidersCicRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders cic request entity too large response has a 5xx status code
func (o *DeleteIdentityprovidersCicRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders cic request entity too large response a status code equal to that given
func (o *DeleteIdentityprovidersCicRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteIdentityprovidersCicRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteIdentityprovidersCicRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteIdentityprovidersCicRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersCicRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersCicUnsupportedMediaType creates a DeleteIdentityprovidersCicUnsupportedMediaType with default headers values
func NewDeleteIdentityprovidersCicUnsupportedMediaType() *DeleteIdentityprovidersCicUnsupportedMediaType {
	return &DeleteIdentityprovidersCicUnsupportedMediaType{}
}

/*
DeleteIdentityprovidersCicUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteIdentityprovidersCicUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders cic unsupported media type response has a 2xx status code
func (o *DeleteIdentityprovidersCicUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders cic unsupported media type response has a 3xx status code
func (o *DeleteIdentityprovidersCicUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders cic unsupported media type response has a 4xx status code
func (o *DeleteIdentityprovidersCicUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders cic unsupported media type response has a 5xx status code
func (o *DeleteIdentityprovidersCicUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders cic unsupported media type response a status code equal to that given
func (o *DeleteIdentityprovidersCicUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteIdentityprovidersCicUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteIdentityprovidersCicUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteIdentityprovidersCicUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersCicUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersCicTooManyRequests creates a DeleteIdentityprovidersCicTooManyRequests with default headers values
func NewDeleteIdentityprovidersCicTooManyRequests() *DeleteIdentityprovidersCicTooManyRequests {
	return &DeleteIdentityprovidersCicTooManyRequests{}
}

/*
DeleteIdentityprovidersCicTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteIdentityprovidersCicTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders cic too many requests response has a 2xx status code
func (o *DeleteIdentityprovidersCicTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders cic too many requests response has a 3xx status code
func (o *DeleteIdentityprovidersCicTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders cic too many requests response has a 4xx status code
func (o *DeleteIdentityprovidersCicTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders cic too many requests response has a 5xx status code
func (o *DeleteIdentityprovidersCicTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders cic too many requests response a status code equal to that given
func (o *DeleteIdentityprovidersCicTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteIdentityprovidersCicTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteIdentityprovidersCicTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteIdentityprovidersCicTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersCicTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersCicInternalServerError creates a DeleteIdentityprovidersCicInternalServerError with default headers values
func NewDeleteIdentityprovidersCicInternalServerError() *DeleteIdentityprovidersCicInternalServerError {
	return &DeleteIdentityprovidersCicInternalServerError{}
}

/*
DeleteIdentityprovidersCicInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteIdentityprovidersCicInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders cic internal server error response has a 2xx status code
func (o *DeleteIdentityprovidersCicInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders cic internal server error response has a 3xx status code
func (o *DeleteIdentityprovidersCicInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders cic internal server error response has a 4xx status code
func (o *DeleteIdentityprovidersCicInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete identityproviders cic internal server error response has a 5xx status code
func (o *DeleteIdentityprovidersCicInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete identityproviders cic internal server error response a status code equal to that given
func (o *DeleteIdentityprovidersCicInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteIdentityprovidersCicInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteIdentityprovidersCicInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteIdentityprovidersCicInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersCicInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersCicServiceUnavailable creates a DeleteIdentityprovidersCicServiceUnavailable with default headers values
func NewDeleteIdentityprovidersCicServiceUnavailable() *DeleteIdentityprovidersCicServiceUnavailable {
	return &DeleteIdentityprovidersCicServiceUnavailable{}
}

/*
DeleteIdentityprovidersCicServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteIdentityprovidersCicServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders cic service unavailable response has a 2xx status code
func (o *DeleteIdentityprovidersCicServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders cic service unavailable response has a 3xx status code
func (o *DeleteIdentityprovidersCicServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders cic service unavailable response has a 4xx status code
func (o *DeleteIdentityprovidersCicServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete identityproviders cic service unavailable response has a 5xx status code
func (o *DeleteIdentityprovidersCicServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete identityproviders cic service unavailable response a status code equal to that given
func (o *DeleteIdentityprovidersCicServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteIdentityprovidersCicServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteIdentityprovidersCicServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteIdentityprovidersCicServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersCicServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersCicGatewayTimeout creates a DeleteIdentityprovidersCicGatewayTimeout with default headers values
func NewDeleteIdentityprovidersCicGatewayTimeout() *DeleteIdentityprovidersCicGatewayTimeout {
	return &DeleteIdentityprovidersCicGatewayTimeout{}
}

/*
DeleteIdentityprovidersCicGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteIdentityprovidersCicGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders cic gateway timeout response has a 2xx status code
func (o *DeleteIdentityprovidersCicGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders cic gateway timeout response has a 3xx status code
func (o *DeleteIdentityprovidersCicGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders cic gateway timeout response has a 4xx status code
func (o *DeleteIdentityprovidersCicGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete identityproviders cic gateway timeout response has a 5xx status code
func (o *DeleteIdentityprovidersCicGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete identityproviders cic gateway timeout response a status code equal to that given
func (o *DeleteIdentityprovidersCicGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteIdentityprovidersCicGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteIdentityprovidersCicGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/cic][%d] deleteIdentityprovidersCicGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteIdentityprovidersCicGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersCicGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
