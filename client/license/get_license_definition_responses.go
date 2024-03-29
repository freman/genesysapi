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

// GetLicenseDefinitionReader is a Reader for the GetLicenseDefinition structure.
type GetLicenseDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLicenseDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLicenseDefinitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLicenseDefinitionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLicenseDefinitionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLicenseDefinitionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLicenseDefinitionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetLicenseDefinitionRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLicenseDefinitionRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLicenseDefinitionUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLicenseDefinitionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLicenseDefinitionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLicenseDefinitionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLicenseDefinitionGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLicenseDefinitionOK creates a GetLicenseDefinitionOK with default headers values
func NewGetLicenseDefinitionOK() *GetLicenseDefinitionOK {
	return &GetLicenseDefinitionOK{}
}

/*
GetLicenseDefinitionOK describes a response with status code 200, with default header values.

successful operation
*/
type GetLicenseDefinitionOK struct {
	Payload *models.LicenseDefinition
}

// IsSuccess returns true when this get license definition o k response has a 2xx status code
func (o *GetLicenseDefinitionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get license definition o k response has a 3xx status code
func (o *GetLicenseDefinitionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license definition o k response has a 4xx status code
func (o *GetLicenseDefinitionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get license definition o k response has a 5xx status code
func (o *GetLicenseDefinitionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get license definition o k response a status code equal to that given
func (o *GetLicenseDefinitionOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetLicenseDefinitionOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionOK  %+v", 200, o.Payload)
}

func (o *GetLicenseDefinitionOK) String() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionOK  %+v", 200, o.Payload)
}

func (o *GetLicenseDefinitionOK) GetPayload() *models.LicenseDefinition {
	return o.Payload
}

func (o *GetLicenseDefinitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LicenseDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionBadRequest creates a GetLicenseDefinitionBadRequest with default headers values
func NewGetLicenseDefinitionBadRequest() *GetLicenseDefinitionBadRequest {
	return &GetLicenseDefinitionBadRequest{}
}

/*
GetLicenseDefinitionBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLicenseDefinitionBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get license definition bad request response has a 2xx status code
func (o *GetLicenseDefinitionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license definition bad request response has a 3xx status code
func (o *GetLicenseDefinitionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license definition bad request response has a 4xx status code
func (o *GetLicenseDefinitionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get license definition bad request response has a 5xx status code
func (o *GetLicenseDefinitionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get license definition bad request response a status code equal to that given
func (o *GetLicenseDefinitionBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetLicenseDefinitionBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionBadRequest  %+v", 400, o.Payload)
}

func (o *GetLicenseDefinitionBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionBadRequest  %+v", 400, o.Payload)
}

func (o *GetLicenseDefinitionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionUnauthorized creates a GetLicenseDefinitionUnauthorized with default headers values
func NewGetLicenseDefinitionUnauthorized() *GetLicenseDefinitionUnauthorized {
	return &GetLicenseDefinitionUnauthorized{}
}

/*
GetLicenseDefinitionUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLicenseDefinitionUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get license definition unauthorized response has a 2xx status code
func (o *GetLicenseDefinitionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license definition unauthorized response has a 3xx status code
func (o *GetLicenseDefinitionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license definition unauthorized response has a 4xx status code
func (o *GetLicenseDefinitionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get license definition unauthorized response has a 5xx status code
func (o *GetLicenseDefinitionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get license definition unauthorized response a status code equal to that given
func (o *GetLicenseDefinitionUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetLicenseDefinitionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLicenseDefinitionUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLicenseDefinitionUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionForbidden creates a GetLicenseDefinitionForbidden with default headers values
func NewGetLicenseDefinitionForbidden() *GetLicenseDefinitionForbidden {
	return &GetLicenseDefinitionForbidden{}
}

/*
GetLicenseDefinitionForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetLicenseDefinitionForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get license definition forbidden response has a 2xx status code
func (o *GetLicenseDefinitionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license definition forbidden response has a 3xx status code
func (o *GetLicenseDefinitionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license definition forbidden response has a 4xx status code
func (o *GetLicenseDefinitionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get license definition forbidden response has a 5xx status code
func (o *GetLicenseDefinitionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get license definition forbidden response a status code equal to that given
func (o *GetLicenseDefinitionForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetLicenseDefinitionForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionForbidden  %+v", 403, o.Payload)
}

func (o *GetLicenseDefinitionForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionForbidden  %+v", 403, o.Payload)
}

func (o *GetLicenseDefinitionForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionNotFound creates a GetLicenseDefinitionNotFound with default headers values
func NewGetLicenseDefinitionNotFound() *GetLicenseDefinitionNotFound {
	return &GetLicenseDefinitionNotFound{}
}

/*
GetLicenseDefinitionNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetLicenseDefinitionNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get license definition not found response has a 2xx status code
func (o *GetLicenseDefinitionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license definition not found response has a 3xx status code
func (o *GetLicenseDefinitionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license definition not found response has a 4xx status code
func (o *GetLicenseDefinitionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get license definition not found response has a 5xx status code
func (o *GetLicenseDefinitionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get license definition not found response a status code equal to that given
func (o *GetLicenseDefinitionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetLicenseDefinitionNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionNotFound  %+v", 404, o.Payload)
}

func (o *GetLicenseDefinitionNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionNotFound  %+v", 404, o.Payload)
}

func (o *GetLicenseDefinitionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionRequestTimeout creates a GetLicenseDefinitionRequestTimeout with default headers values
func NewGetLicenseDefinitionRequestTimeout() *GetLicenseDefinitionRequestTimeout {
	return &GetLicenseDefinitionRequestTimeout{}
}

/*
GetLicenseDefinitionRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetLicenseDefinitionRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get license definition request timeout response has a 2xx status code
func (o *GetLicenseDefinitionRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license definition request timeout response has a 3xx status code
func (o *GetLicenseDefinitionRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license definition request timeout response has a 4xx status code
func (o *GetLicenseDefinitionRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get license definition request timeout response has a 5xx status code
func (o *GetLicenseDefinitionRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get license definition request timeout response a status code equal to that given
func (o *GetLicenseDefinitionRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetLicenseDefinitionRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLicenseDefinitionRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLicenseDefinitionRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionRequestEntityTooLarge creates a GetLicenseDefinitionRequestEntityTooLarge with default headers values
func NewGetLicenseDefinitionRequestEntityTooLarge() *GetLicenseDefinitionRequestEntityTooLarge {
	return &GetLicenseDefinitionRequestEntityTooLarge{}
}

/*
GetLicenseDefinitionRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetLicenseDefinitionRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get license definition request entity too large response has a 2xx status code
func (o *GetLicenseDefinitionRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license definition request entity too large response has a 3xx status code
func (o *GetLicenseDefinitionRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license definition request entity too large response has a 4xx status code
func (o *GetLicenseDefinitionRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get license definition request entity too large response has a 5xx status code
func (o *GetLicenseDefinitionRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get license definition request entity too large response a status code equal to that given
func (o *GetLicenseDefinitionRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetLicenseDefinitionRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLicenseDefinitionRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLicenseDefinitionRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionUnsupportedMediaType creates a GetLicenseDefinitionUnsupportedMediaType with default headers values
func NewGetLicenseDefinitionUnsupportedMediaType() *GetLicenseDefinitionUnsupportedMediaType {
	return &GetLicenseDefinitionUnsupportedMediaType{}
}

/*
GetLicenseDefinitionUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLicenseDefinitionUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get license definition unsupported media type response has a 2xx status code
func (o *GetLicenseDefinitionUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license definition unsupported media type response has a 3xx status code
func (o *GetLicenseDefinitionUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license definition unsupported media type response has a 4xx status code
func (o *GetLicenseDefinitionUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get license definition unsupported media type response has a 5xx status code
func (o *GetLicenseDefinitionUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get license definition unsupported media type response a status code equal to that given
func (o *GetLicenseDefinitionUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetLicenseDefinitionUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLicenseDefinitionUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLicenseDefinitionUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionTooManyRequests creates a GetLicenseDefinitionTooManyRequests with default headers values
func NewGetLicenseDefinitionTooManyRequests() *GetLicenseDefinitionTooManyRequests {
	return &GetLicenseDefinitionTooManyRequests{}
}

/*
GetLicenseDefinitionTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetLicenseDefinitionTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get license definition too many requests response has a 2xx status code
func (o *GetLicenseDefinitionTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license definition too many requests response has a 3xx status code
func (o *GetLicenseDefinitionTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license definition too many requests response has a 4xx status code
func (o *GetLicenseDefinitionTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get license definition too many requests response has a 5xx status code
func (o *GetLicenseDefinitionTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get license definition too many requests response a status code equal to that given
func (o *GetLicenseDefinitionTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetLicenseDefinitionTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLicenseDefinitionTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLicenseDefinitionTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionInternalServerError creates a GetLicenseDefinitionInternalServerError with default headers values
func NewGetLicenseDefinitionInternalServerError() *GetLicenseDefinitionInternalServerError {
	return &GetLicenseDefinitionInternalServerError{}
}

/*
GetLicenseDefinitionInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLicenseDefinitionInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get license definition internal server error response has a 2xx status code
func (o *GetLicenseDefinitionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license definition internal server error response has a 3xx status code
func (o *GetLicenseDefinitionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license definition internal server error response has a 4xx status code
func (o *GetLicenseDefinitionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get license definition internal server error response has a 5xx status code
func (o *GetLicenseDefinitionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get license definition internal server error response a status code equal to that given
func (o *GetLicenseDefinitionInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetLicenseDefinitionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLicenseDefinitionInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLicenseDefinitionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionServiceUnavailable creates a GetLicenseDefinitionServiceUnavailable with default headers values
func NewGetLicenseDefinitionServiceUnavailable() *GetLicenseDefinitionServiceUnavailable {
	return &GetLicenseDefinitionServiceUnavailable{}
}

/*
GetLicenseDefinitionServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLicenseDefinitionServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get license definition service unavailable response has a 2xx status code
func (o *GetLicenseDefinitionServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license definition service unavailable response has a 3xx status code
func (o *GetLicenseDefinitionServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license definition service unavailable response has a 4xx status code
func (o *GetLicenseDefinitionServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get license definition service unavailable response has a 5xx status code
func (o *GetLicenseDefinitionServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get license definition service unavailable response a status code equal to that given
func (o *GetLicenseDefinitionServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetLicenseDefinitionServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLicenseDefinitionServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLicenseDefinitionServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseDefinitionGatewayTimeout creates a GetLicenseDefinitionGatewayTimeout with default headers values
func NewGetLicenseDefinitionGatewayTimeout() *GetLicenseDefinitionGatewayTimeout {
	return &GetLicenseDefinitionGatewayTimeout{}
}

/*
GetLicenseDefinitionGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetLicenseDefinitionGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get license definition gateway timeout response has a 2xx status code
func (o *GetLicenseDefinitionGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license definition gateway timeout response has a 3xx status code
func (o *GetLicenseDefinitionGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license definition gateway timeout response has a 4xx status code
func (o *GetLicenseDefinitionGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get license definition gateway timeout response has a 5xx status code
func (o *GetLicenseDefinitionGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get license definition gateway timeout response a status code equal to that given
func (o *GetLicenseDefinitionGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetLicenseDefinitionGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLicenseDefinitionGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/license/definitions/{licenseId}][%d] getLicenseDefinitionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLicenseDefinitionGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseDefinitionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
