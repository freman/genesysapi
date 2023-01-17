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

// GetPresencedefinitionsReader is a Reader for the GetPresencedefinitions structure.
type GetPresencedefinitionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPresencedefinitionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPresencedefinitionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPresencedefinitionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPresencedefinitionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPresencedefinitionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPresencedefinitionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetPresencedefinitionsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetPresencedefinitionsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetPresencedefinitionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPresencedefinitionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPresencedefinitionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetPresencedefinitionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetPresencedefinitionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPresencedefinitionsOK creates a GetPresencedefinitionsOK with default headers values
func NewGetPresencedefinitionsOK() *GetPresencedefinitionsOK {
	return &GetPresencedefinitionsOK{}
}

/*
GetPresencedefinitionsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetPresencedefinitionsOK struct {
	Payload *models.OrganizationPresenceEntityListing
}

// IsSuccess returns true when this get presencedefinitions o k response has a 2xx status code
func (o *GetPresencedefinitionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get presencedefinitions o k response has a 3xx status code
func (o *GetPresencedefinitionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presencedefinitions o k response has a 4xx status code
func (o *GetPresencedefinitionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get presencedefinitions o k response has a 5xx status code
func (o *GetPresencedefinitionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get presencedefinitions o k response a status code equal to that given
func (o *GetPresencedefinitionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPresencedefinitionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsOK  %+v", 200, o.Payload)
}

func (o *GetPresencedefinitionsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsOK  %+v", 200, o.Payload)
}

func (o *GetPresencedefinitionsOK) GetPayload() *models.OrganizationPresenceEntityListing {
	return o.Payload
}

func (o *GetPresencedefinitionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationPresenceEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresencedefinitionsBadRequest creates a GetPresencedefinitionsBadRequest with default headers values
func NewGetPresencedefinitionsBadRequest() *GetPresencedefinitionsBadRequest {
	return &GetPresencedefinitionsBadRequest{}
}

/*
GetPresencedefinitionsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetPresencedefinitionsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get presencedefinitions bad request response has a 2xx status code
func (o *GetPresencedefinitionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presencedefinitions bad request response has a 3xx status code
func (o *GetPresencedefinitionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presencedefinitions bad request response has a 4xx status code
func (o *GetPresencedefinitionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get presencedefinitions bad request response has a 5xx status code
func (o *GetPresencedefinitionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get presencedefinitions bad request response a status code equal to that given
func (o *GetPresencedefinitionsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetPresencedefinitionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetPresencedefinitionsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetPresencedefinitionsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPresencedefinitionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresencedefinitionsUnauthorized creates a GetPresencedefinitionsUnauthorized with default headers values
func NewGetPresencedefinitionsUnauthorized() *GetPresencedefinitionsUnauthorized {
	return &GetPresencedefinitionsUnauthorized{}
}

/*
GetPresencedefinitionsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetPresencedefinitionsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get presencedefinitions unauthorized response has a 2xx status code
func (o *GetPresencedefinitionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presencedefinitions unauthorized response has a 3xx status code
func (o *GetPresencedefinitionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presencedefinitions unauthorized response has a 4xx status code
func (o *GetPresencedefinitionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get presencedefinitions unauthorized response has a 5xx status code
func (o *GetPresencedefinitionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get presencedefinitions unauthorized response a status code equal to that given
func (o *GetPresencedefinitionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetPresencedefinitionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPresencedefinitionsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPresencedefinitionsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPresencedefinitionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresencedefinitionsForbidden creates a GetPresencedefinitionsForbidden with default headers values
func NewGetPresencedefinitionsForbidden() *GetPresencedefinitionsForbidden {
	return &GetPresencedefinitionsForbidden{}
}

/*
GetPresencedefinitionsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetPresencedefinitionsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get presencedefinitions forbidden response has a 2xx status code
func (o *GetPresencedefinitionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presencedefinitions forbidden response has a 3xx status code
func (o *GetPresencedefinitionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presencedefinitions forbidden response has a 4xx status code
func (o *GetPresencedefinitionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get presencedefinitions forbidden response has a 5xx status code
func (o *GetPresencedefinitionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get presencedefinitions forbidden response a status code equal to that given
func (o *GetPresencedefinitionsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetPresencedefinitionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsForbidden  %+v", 403, o.Payload)
}

func (o *GetPresencedefinitionsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsForbidden  %+v", 403, o.Payload)
}

func (o *GetPresencedefinitionsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPresencedefinitionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresencedefinitionsNotFound creates a GetPresencedefinitionsNotFound with default headers values
func NewGetPresencedefinitionsNotFound() *GetPresencedefinitionsNotFound {
	return &GetPresencedefinitionsNotFound{}
}

/*
GetPresencedefinitionsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetPresencedefinitionsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get presencedefinitions not found response has a 2xx status code
func (o *GetPresencedefinitionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presencedefinitions not found response has a 3xx status code
func (o *GetPresencedefinitionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presencedefinitions not found response has a 4xx status code
func (o *GetPresencedefinitionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get presencedefinitions not found response has a 5xx status code
func (o *GetPresencedefinitionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get presencedefinitions not found response a status code equal to that given
func (o *GetPresencedefinitionsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetPresencedefinitionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsNotFound  %+v", 404, o.Payload)
}

func (o *GetPresencedefinitionsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsNotFound  %+v", 404, o.Payload)
}

func (o *GetPresencedefinitionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPresencedefinitionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresencedefinitionsRequestTimeout creates a GetPresencedefinitionsRequestTimeout with default headers values
func NewGetPresencedefinitionsRequestTimeout() *GetPresencedefinitionsRequestTimeout {
	return &GetPresencedefinitionsRequestTimeout{}
}

/*
GetPresencedefinitionsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetPresencedefinitionsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get presencedefinitions request timeout response has a 2xx status code
func (o *GetPresencedefinitionsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presencedefinitions request timeout response has a 3xx status code
func (o *GetPresencedefinitionsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presencedefinitions request timeout response has a 4xx status code
func (o *GetPresencedefinitionsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get presencedefinitions request timeout response has a 5xx status code
func (o *GetPresencedefinitionsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get presencedefinitions request timeout response a status code equal to that given
func (o *GetPresencedefinitionsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetPresencedefinitionsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetPresencedefinitionsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetPresencedefinitionsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPresencedefinitionsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresencedefinitionsRequestEntityTooLarge creates a GetPresencedefinitionsRequestEntityTooLarge with default headers values
func NewGetPresencedefinitionsRequestEntityTooLarge() *GetPresencedefinitionsRequestEntityTooLarge {
	return &GetPresencedefinitionsRequestEntityTooLarge{}
}

/*
GetPresencedefinitionsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetPresencedefinitionsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get presencedefinitions request entity too large response has a 2xx status code
func (o *GetPresencedefinitionsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presencedefinitions request entity too large response has a 3xx status code
func (o *GetPresencedefinitionsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presencedefinitions request entity too large response has a 4xx status code
func (o *GetPresencedefinitionsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get presencedefinitions request entity too large response has a 5xx status code
func (o *GetPresencedefinitionsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get presencedefinitions request entity too large response a status code equal to that given
func (o *GetPresencedefinitionsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetPresencedefinitionsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetPresencedefinitionsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetPresencedefinitionsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPresencedefinitionsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresencedefinitionsUnsupportedMediaType creates a GetPresencedefinitionsUnsupportedMediaType with default headers values
func NewGetPresencedefinitionsUnsupportedMediaType() *GetPresencedefinitionsUnsupportedMediaType {
	return &GetPresencedefinitionsUnsupportedMediaType{}
}

/*
GetPresencedefinitionsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetPresencedefinitionsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get presencedefinitions unsupported media type response has a 2xx status code
func (o *GetPresencedefinitionsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presencedefinitions unsupported media type response has a 3xx status code
func (o *GetPresencedefinitionsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presencedefinitions unsupported media type response has a 4xx status code
func (o *GetPresencedefinitionsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get presencedefinitions unsupported media type response has a 5xx status code
func (o *GetPresencedefinitionsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get presencedefinitions unsupported media type response a status code equal to that given
func (o *GetPresencedefinitionsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetPresencedefinitionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetPresencedefinitionsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetPresencedefinitionsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPresencedefinitionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresencedefinitionsTooManyRequests creates a GetPresencedefinitionsTooManyRequests with default headers values
func NewGetPresencedefinitionsTooManyRequests() *GetPresencedefinitionsTooManyRequests {
	return &GetPresencedefinitionsTooManyRequests{}
}

/*
GetPresencedefinitionsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetPresencedefinitionsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get presencedefinitions too many requests response has a 2xx status code
func (o *GetPresencedefinitionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presencedefinitions too many requests response has a 3xx status code
func (o *GetPresencedefinitionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presencedefinitions too many requests response has a 4xx status code
func (o *GetPresencedefinitionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get presencedefinitions too many requests response has a 5xx status code
func (o *GetPresencedefinitionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get presencedefinitions too many requests response a status code equal to that given
func (o *GetPresencedefinitionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetPresencedefinitionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPresencedefinitionsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPresencedefinitionsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPresencedefinitionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresencedefinitionsInternalServerError creates a GetPresencedefinitionsInternalServerError with default headers values
func NewGetPresencedefinitionsInternalServerError() *GetPresencedefinitionsInternalServerError {
	return &GetPresencedefinitionsInternalServerError{}
}

/*
GetPresencedefinitionsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetPresencedefinitionsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get presencedefinitions internal server error response has a 2xx status code
func (o *GetPresencedefinitionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presencedefinitions internal server error response has a 3xx status code
func (o *GetPresencedefinitionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presencedefinitions internal server error response has a 4xx status code
func (o *GetPresencedefinitionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get presencedefinitions internal server error response has a 5xx status code
func (o *GetPresencedefinitionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get presencedefinitions internal server error response a status code equal to that given
func (o *GetPresencedefinitionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetPresencedefinitionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPresencedefinitionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPresencedefinitionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPresencedefinitionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresencedefinitionsServiceUnavailable creates a GetPresencedefinitionsServiceUnavailable with default headers values
func NewGetPresencedefinitionsServiceUnavailable() *GetPresencedefinitionsServiceUnavailable {
	return &GetPresencedefinitionsServiceUnavailable{}
}

/*
GetPresencedefinitionsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetPresencedefinitionsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get presencedefinitions service unavailable response has a 2xx status code
func (o *GetPresencedefinitionsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presencedefinitions service unavailable response has a 3xx status code
func (o *GetPresencedefinitionsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presencedefinitions service unavailable response has a 4xx status code
func (o *GetPresencedefinitionsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get presencedefinitions service unavailable response has a 5xx status code
func (o *GetPresencedefinitionsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get presencedefinitions service unavailable response a status code equal to that given
func (o *GetPresencedefinitionsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetPresencedefinitionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetPresencedefinitionsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetPresencedefinitionsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPresencedefinitionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresencedefinitionsGatewayTimeout creates a GetPresencedefinitionsGatewayTimeout with default headers values
func NewGetPresencedefinitionsGatewayTimeout() *GetPresencedefinitionsGatewayTimeout {
	return &GetPresencedefinitionsGatewayTimeout{}
}

/*
GetPresencedefinitionsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetPresencedefinitionsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get presencedefinitions gateway timeout response has a 2xx status code
func (o *GetPresencedefinitionsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presencedefinitions gateway timeout response has a 3xx status code
func (o *GetPresencedefinitionsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presencedefinitions gateway timeout response has a 4xx status code
func (o *GetPresencedefinitionsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get presencedefinitions gateway timeout response has a 5xx status code
func (o *GetPresencedefinitionsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get presencedefinitions gateway timeout response a status code equal to that given
func (o *GetPresencedefinitionsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetPresencedefinitionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetPresencedefinitionsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/presencedefinitions][%d] getPresencedefinitionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetPresencedefinitionsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPresencedefinitionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
