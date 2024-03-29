// Code generated by go-swagger; DO NOT EDIT.

package license

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostLicenseToggleReader is a Reader for the PostLicenseToggle structure.
type PostLicenseToggleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLicenseToggleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLicenseToggleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLicenseToggleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostLicenseToggleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLicenseToggleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostLicenseToggleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostLicenseToggleRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostLicenseToggleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostLicenseToggleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostLicenseToggleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostLicenseToggleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostLicenseToggleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostLicenseToggleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLicenseToggleOK creates a PostLicenseToggleOK with default headers values
func NewPostLicenseToggleOK() *PostLicenseToggleOK {
	return &PostLicenseToggleOK{}
}

/*
PostLicenseToggleOK describes a response with status code 200, with default header values.

successful operation
*/
type PostLicenseToggleOK struct {
	Payload *models.LicenseOrgToggle
}

// IsSuccess returns true when this post license toggle o k response has a 2xx status code
func (o *PostLicenseToggleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post license toggle o k response has a 3xx status code
func (o *PostLicenseToggleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license toggle o k response has a 4xx status code
func (o *PostLicenseToggleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post license toggle o k response has a 5xx status code
func (o *PostLicenseToggleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post license toggle o k response a status code equal to that given
func (o *PostLicenseToggleOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostLicenseToggleOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleOK  %+v", 200, o.Payload)
}

func (o *PostLicenseToggleOK) String() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleOK  %+v", 200, o.Payload)
}

func (o *PostLicenseToggleOK) GetPayload() *models.LicenseOrgToggle {
	return o.Payload
}

func (o *PostLicenseToggleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LicenseOrgToggle)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseToggleBadRequest creates a PostLicenseToggleBadRequest with default headers values
func NewPostLicenseToggleBadRequest() *PostLicenseToggleBadRequest {
	return &PostLicenseToggleBadRequest{}
}

/*
PostLicenseToggleBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostLicenseToggleBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post license toggle bad request response has a 2xx status code
func (o *PostLicenseToggleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post license toggle bad request response has a 3xx status code
func (o *PostLicenseToggleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license toggle bad request response has a 4xx status code
func (o *PostLicenseToggleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post license toggle bad request response has a 5xx status code
func (o *PostLicenseToggleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post license toggle bad request response a status code equal to that given
func (o *PostLicenseToggleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostLicenseToggleBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleBadRequest  %+v", 400, o.Payload)
}

func (o *PostLicenseToggleBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleBadRequest  %+v", 400, o.Payload)
}

func (o *PostLicenseToggleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseToggleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseToggleUnauthorized creates a PostLicenseToggleUnauthorized with default headers values
func NewPostLicenseToggleUnauthorized() *PostLicenseToggleUnauthorized {
	return &PostLicenseToggleUnauthorized{}
}

/*
PostLicenseToggleUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostLicenseToggleUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post license toggle unauthorized response has a 2xx status code
func (o *PostLicenseToggleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post license toggle unauthorized response has a 3xx status code
func (o *PostLicenseToggleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license toggle unauthorized response has a 4xx status code
func (o *PostLicenseToggleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post license toggle unauthorized response has a 5xx status code
func (o *PostLicenseToggleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post license toggle unauthorized response a status code equal to that given
func (o *PostLicenseToggleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostLicenseToggleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLicenseToggleUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLicenseToggleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseToggleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseToggleForbidden creates a PostLicenseToggleForbidden with default headers values
func NewPostLicenseToggleForbidden() *PostLicenseToggleForbidden {
	return &PostLicenseToggleForbidden{}
}

/*
PostLicenseToggleForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostLicenseToggleForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post license toggle forbidden response has a 2xx status code
func (o *PostLicenseToggleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post license toggle forbidden response has a 3xx status code
func (o *PostLicenseToggleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license toggle forbidden response has a 4xx status code
func (o *PostLicenseToggleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post license toggle forbidden response has a 5xx status code
func (o *PostLicenseToggleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post license toggle forbidden response a status code equal to that given
func (o *PostLicenseToggleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostLicenseToggleForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleForbidden  %+v", 403, o.Payload)
}

func (o *PostLicenseToggleForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleForbidden  %+v", 403, o.Payload)
}

func (o *PostLicenseToggleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseToggleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseToggleNotFound creates a PostLicenseToggleNotFound with default headers values
func NewPostLicenseToggleNotFound() *PostLicenseToggleNotFound {
	return &PostLicenseToggleNotFound{}
}

/*
PostLicenseToggleNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostLicenseToggleNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post license toggle not found response has a 2xx status code
func (o *PostLicenseToggleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post license toggle not found response has a 3xx status code
func (o *PostLicenseToggleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license toggle not found response has a 4xx status code
func (o *PostLicenseToggleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post license toggle not found response has a 5xx status code
func (o *PostLicenseToggleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post license toggle not found response a status code equal to that given
func (o *PostLicenseToggleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostLicenseToggleNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleNotFound  %+v", 404, o.Payload)
}

func (o *PostLicenseToggleNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleNotFound  %+v", 404, o.Payload)
}

func (o *PostLicenseToggleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseToggleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseToggleRequestTimeout creates a PostLicenseToggleRequestTimeout with default headers values
func NewPostLicenseToggleRequestTimeout() *PostLicenseToggleRequestTimeout {
	return &PostLicenseToggleRequestTimeout{}
}

/*
PostLicenseToggleRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostLicenseToggleRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post license toggle request timeout response has a 2xx status code
func (o *PostLicenseToggleRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post license toggle request timeout response has a 3xx status code
func (o *PostLicenseToggleRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license toggle request timeout response has a 4xx status code
func (o *PostLicenseToggleRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post license toggle request timeout response has a 5xx status code
func (o *PostLicenseToggleRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post license toggle request timeout response a status code equal to that given
func (o *PostLicenseToggleRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostLicenseToggleRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostLicenseToggleRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostLicenseToggleRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseToggleRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseToggleRequestEntityTooLarge creates a PostLicenseToggleRequestEntityTooLarge with default headers values
func NewPostLicenseToggleRequestEntityTooLarge() *PostLicenseToggleRequestEntityTooLarge {
	return &PostLicenseToggleRequestEntityTooLarge{}
}

/*
PostLicenseToggleRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostLicenseToggleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post license toggle request entity too large response has a 2xx status code
func (o *PostLicenseToggleRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post license toggle request entity too large response has a 3xx status code
func (o *PostLicenseToggleRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license toggle request entity too large response has a 4xx status code
func (o *PostLicenseToggleRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post license toggle request entity too large response has a 5xx status code
func (o *PostLicenseToggleRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post license toggle request entity too large response a status code equal to that given
func (o *PostLicenseToggleRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostLicenseToggleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLicenseToggleRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLicenseToggleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseToggleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseToggleUnsupportedMediaType creates a PostLicenseToggleUnsupportedMediaType with default headers values
func NewPostLicenseToggleUnsupportedMediaType() *PostLicenseToggleUnsupportedMediaType {
	return &PostLicenseToggleUnsupportedMediaType{}
}

/*
PostLicenseToggleUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostLicenseToggleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post license toggle unsupported media type response has a 2xx status code
func (o *PostLicenseToggleUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post license toggle unsupported media type response has a 3xx status code
func (o *PostLicenseToggleUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license toggle unsupported media type response has a 4xx status code
func (o *PostLicenseToggleUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post license toggle unsupported media type response has a 5xx status code
func (o *PostLicenseToggleUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post license toggle unsupported media type response a status code equal to that given
func (o *PostLicenseToggleUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostLicenseToggleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLicenseToggleUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLicenseToggleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseToggleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseToggleTooManyRequests creates a PostLicenseToggleTooManyRequests with default headers values
func NewPostLicenseToggleTooManyRequests() *PostLicenseToggleTooManyRequests {
	return &PostLicenseToggleTooManyRequests{}
}

/*
PostLicenseToggleTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostLicenseToggleTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post license toggle too many requests response has a 2xx status code
func (o *PostLicenseToggleTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post license toggle too many requests response has a 3xx status code
func (o *PostLicenseToggleTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license toggle too many requests response has a 4xx status code
func (o *PostLicenseToggleTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post license toggle too many requests response has a 5xx status code
func (o *PostLicenseToggleTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post license toggle too many requests response a status code equal to that given
func (o *PostLicenseToggleTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostLicenseToggleTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLicenseToggleTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLicenseToggleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseToggleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseToggleInternalServerError creates a PostLicenseToggleInternalServerError with default headers values
func NewPostLicenseToggleInternalServerError() *PostLicenseToggleInternalServerError {
	return &PostLicenseToggleInternalServerError{}
}

/*
PostLicenseToggleInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostLicenseToggleInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post license toggle internal server error response has a 2xx status code
func (o *PostLicenseToggleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post license toggle internal server error response has a 3xx status code
func (o *PostLicenseToggleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license toggle internal server error response has a 4xx status code
func (o *PostLicenseToggleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post license toggle internal server error response has a 5xx status code
func (o *PostLicenseToggleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post license toggle internal server error response a status code equal to that given
func (o *PostLicenseToggleInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostLicenseToggleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLicenseToggleInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLicenseToggleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseToggleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseToggleServiceUnavailable creates a PostLicenseToggleServiceUnavailable with default headers values
func NewPostLicenseToggleServiceUnavailable() *PostLicenseToggleServiceUnavailable {
	return &PostLicenseToggleServiceUnavailable{}
}

/*
PostLicenseToggleServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostLicenseToggleServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post license toggle service unavailable response has a 2xx status code
func (o *PostLicenseToggleServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post license toggle service unavailable response has a 3xx status code
func (o *PostLicenseToggleServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license toggle service unavailable response has a 4xx status code
func (o *PostLicenseToggleServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post license toggle service unavailable response has a 5xx status code
func (o *PostLicenseToggleServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post license toggle service unavailable response a status code equal to that given
func (o *PostLicenseToggleServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostLicenseToggleServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLicenseToggleServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLicenseToggleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseToggleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseToggleGatewayTimeout creates a PostLicenseToggleGatewayTimeout with default headers values
func NewPostLicenseToggleGatewayTimeout() *PostLicenseToggleGatewayTimeout {
	return &PostLicenseToggleGatewayTimeout{}
}

/*
PostLicenseToggleGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostLicenseToggleGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post license toggle gateway timeout response has a 2xx status code
func (o *PostLicenseToggleGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post license toggle gateway timeout response has a 3xx status code
func (o *PostLicenseToggleGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license toggle gateway timeout response has a 4xx status code
func (o *PostLicenseToggleGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post license toggle gateway timeout response has a 5xx status code
func (o *PostLicenseToggleGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post license toggle gateway timeout response a status code equal to that given
func (o *PostLicenseToggleGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostLicenseToggleGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLicenseToggleGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/license/toggles/{featureName}][%d] postLicenseToggleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLicenseToggleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseToggleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
