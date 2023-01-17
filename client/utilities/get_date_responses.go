// Code generated by go-swagger; DO NOT EDIT.

package utilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetDateReader is a Reader for the GetDate structure.
type GetDateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetDateRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetDateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetDateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetDateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetDateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetDateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDateOK creates a GetDateOK with default headers values
func NewGetDateOK() *GetDateOK {
	return &GetDateOK{}
}

/*
GetDateOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDateOK struct {
	Payload *models.ServerDate
}

// IsSuccess returns true when this get date o k response has a 2xx status code
func (o *GetDateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get date o k response has a 3xx status code
func (o *GetDateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get date o k response has a 4xx status code
func (o *GetDateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get date o k response has a 5xx status code
func (o *GetDateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get date o k response a status code equal to that given
func (o *GetDateOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDateOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateOK  %+v", 200, o.Payload)
}

func (o *GetDateOK) String() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateOK  %+v", 200, o.Payload)
}

func (o *GetDateOK) GetPayload() *models.ServerDate {
	return o.Payload
}

func (o *GetDateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerDate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDateBadRequest creates a GetDateBadRequest with default headers values
func NewGetDateBadRequest() *GetDateBadRequest {
	return &GetDateBadRequest{}
}

/*
GetDateBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetDateBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get date bad request response has a 2xx status code
func (o *GetDateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get date bad request response has a 3xx status code
func (o *GetDateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get date bad request response has a 4xx status code
func (o *GetDateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get date bad request response has a 5xx status code
func (o *GetDateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get date bad request response a status code equal to that given
func (o *GetDateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetDateBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateBadRequest  %+v", 400, o.Payload)
}

func (o *GetDateBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateBadRequest  %+v", 400, o.Payload)
}

func (o *GetDateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDateUnauthorized creates a GetDateUnauthorized with default headers values
func NewGetDateUnauthorized() *GetDateUnauthorized {
	return &GetDateUnauthorized{}
}

/*
GetDateUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetDateUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get date unauthorized response has a 2xx status code
func (o *GetDateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get date unauthorized response has a 3xx status code
func (o *GetDateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get date unauthorized response has a 4xx status code
func (o *GetDateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get date unauthorized response has a 5xx status code
func (o *GetDateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get date unauthorized response a status code equal to that given
func (o *GetDateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetDateUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDateUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDateForbidden creates a GetDateForbidden with default headers values
func NewGetDateForbidden() *GetDateForbidden {
	return &GetDateForbidden{}
}

/*
GetDateForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetDateForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get date forbidden response has a 2xx status code
func (o *GetDateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get date forbidden response has a 3xx status code
func (o *GetDateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get date forbidden response has a 4xx status code
func (o *GetDateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get date forbidden response has a 5xx status code
func (o *GetDateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get date forbidden response a status code equal to that given
func (o *GetDateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetDateForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateForbidden  %+v", 403, o.Payload)
}

func (o *GetDateForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateForbidden  %+v", 403, o.Payload)
}

func (o *GetDateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDateNotFound creates a GetDateNotFound with default headers values
func NewGetDateNotFound() *GetDateNotFound {
	return &GetDateNotFound{}
}

/*
GetDateNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetDateNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get date not found response has a 2xx status code
func (o *GetDateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get date not found response has a 3xx status code
func (o *GetDateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get date not found response has a 4xx status code
func (o *GetDateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get date not found response has a 5xx status code
func (o *GetDateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get date not found response a status code equal to that given
func (o *GetDateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetDateNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateNotFound  %+v", 404, o.Payload)
}

func (o *GetDateNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateNotFound  %+v", 404, o.Payload)
}

func (o *GetDateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDateRequestTimeout creates a GetDateRequestTimeout with default headers values
func NewGetDateRequestTimeout() *GetDateRequestTimeout {
	return &GetDateRequestTimeout{}
}

/*
GetDateRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetDateRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get date request timeout response has a 2xx status code
func (o *GetDateRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get date request timeout response has a 3xx status code
func (o *GetDateRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get date request timeout response has a 4xx status code
func (o *GetDateRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get date request timeout response has a 5xx status code
func (o *GetDateRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get date request timeout response a status code equal to that given
func (o *GetDateRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetDateRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetDateRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetDateRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDateRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDateRequestEntityTooLarge creates a GetDateRequestEntityTooLarge with default headers values
func NewGetDateRequestEntityTooLarge() *GetDateRequestEntityTooLarge {
	return &GetDateRequestEntityTooLarge{}
}

/*
GetDateRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetDateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get date request entity too large response has a 2xx status code
func (o *GetDateRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get date request entity too large response has a 3xx status code
func (o *GetDateRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get date request entity too large response has a 4xx status code
func (o *GetDateRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get date request entity too large response has a 5xx status code
func (o *GetDateRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get date request entity too large response a status code equal to that given
func (o *GetDateRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetDateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetDateRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetDateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDateUnsupportedMediaType creates a GetDateUnsupportedMediaType with default headers values
func NewGetDateUnsupportedMediaType() *GetDateUnsupportedMediaType {
	return &GetDateUnsupportedMediaType{}
}

/*
GetDateUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetDateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get date unsupported media type response has a 2xx status code
func (o *GetDateUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get date unsupported media type response has a 3xx status code
func (o *GetDateUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get date unsupported media type response has a 4xx status code
func (o *GetDateUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get date unsupported media type response has a 5xx status code
func (o *GetDateUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get date unsupported media type response a status code equal to that given
func (o *GetDateUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetDateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetDateUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetDateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDateTooManyRequests creates a GetDateTooManyRequests with default headers values
func NewGetDateTooManyRequests() *GetDateTooManyRequests {
	return &GetDateTooManyRequests{}
}

/*
GetDateTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetDateTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get date too many requests response has a 2xx status code
func (o *GetDateTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get date too many requests response has a 3xx status code
func (o *GetDateTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get date too many requests response has a 4xx status code
func (o *GetDateTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get date too many requests response has a 5xx status code
func (o *GetDateTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get date too many requests response a status code equal to that given
func (o *GetDateTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetDateTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDateTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDateInternalServerError creates a GetDateInternalServerError with default headers values
func NewGetDateInternalServerError() *GetDateInternalServerError {
	return &GetDateInternalServerError{}
}

/*
GetDateInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetDateInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get date internal server error response has a 2xx status code
func (o *GetDateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get date internal server error response has a 3xx status code
func (o *GetDateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get date internal server error response has a 4xx status code
func (o *GetDateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get date internal server error response has a 5xx status code
func (o *GetDateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get date internal server error response a status code equal to that given
func (o *GetDateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetDateInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDateInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDateServiceUnavailable creates a GetDateServiceUnavailable with default headers values
func NewGetDateServiceUnavailable() *GetDateServiceUnavailable {
	return &GetDateServiceUnavailable{}
}

/*
GetDateServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetDateServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get date service unavailable response has a 2xx status code
func (o *GetDateServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get date service unavailable response has a 3xx status code
func (o *GetDateServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get date service unavailable response has a 4xx status code
func (o *GetDateServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get date service unavailable response has a 5xx status code
func (o *GetDateServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get date service unavailable response a status code equal to that given
func (o *GetDateServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetDateServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetDateServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetDateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDateGatewayTimeout creates a GetDateGatewayTimeout with default headers values
func NewGetDateGatewayTimeout() *GetDateGatewayTimeout {
	return &GetDateGatewayTimeout{}
}

/*
GetDateGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetDateGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get date gateway timeout response has a 2xx status code
func (o *GetDateGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get date gateway timeout response has a 3xx status code
func (o *GetDateGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get date gateway timeout response has a 4xx status code
func (o *GetDateGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get date gateway timeout response has a 5xx status code
func (o *GetDateGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get date gateway timeout response a status code equal to that given
func (o *GetDateGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetDateGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetDateGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/date][%d] getDateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetDateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
