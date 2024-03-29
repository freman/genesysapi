// Code generated by go-swagger; DO NOT EDIT.

package language_understanding

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetLanguageunderstandingDomainReader is a Reader for the GetLanguageunderstandingDomain structure.
type GetLanguageunderstandingDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLanguageunderstandingDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLanguageunderstandingDomainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLanguageunderstandingDomainBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLanguageunderstandingDomainUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLanguageunderstandingDomainForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLanguageunderstandingDomainNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetLanguageunderstandingDomainRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLanguageunderstandingDomainRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLanguageunderstandingDomainUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLanguageunderstandingDomainTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLanguageunderstandingDomainInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLanguageunderstandingDomainServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLanguageunderstandingDomainGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLanguageunderstandingDomainOK creates a GetLanguageunderstandingDomainOK with default headers values
func NewGetLanguageunderstandingDomainOK() *GetLanguageunderstandingDomainOK {
	return &GetLanguageunderstandingDomainOK{}
}

/*
GetLanguageunderstandingDomainOK describes a response with status code 200, with default header values.

successful operation
*/
type GetLanguageunderstandingDomainOK struct {
	Payload *models.NluDomain
}

// IsSuccess returns true when this get languageunderstanding domain o k response has a 2xx status code
func (o *GetLanguageunderstandingDomainOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get languageunderstanding domain o k response has a 3xx status code
func (o *GetLanguageunderstandingDomainOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding domain o k response has a 4xx status code
func (o *GetLanguageunderstandingDomainOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get languageunderstanding domain o k response has a 5xx status code
func (o *GetLanguageunderstandingDomainOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding domain o k response a status code equal to that given
func (o *GetLanguageunderstandingDomainOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetLanguageunderstandingDomainOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainOK  %+v", 200, o.Payload)
}

func (o *GetLanguageunderstandingDomainOK) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainOK  %+v", 200, o.Payload)
}

func (o *GetLanguageunderstandingDomainOK) GetPayload() *models.NluDomain {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NluDomain)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainBadRequest creates a GetLanguageunderstandingDomainBadRequest with default headers values
func NewGetLanguageunderstandingDomainBadRequest() *GetLanguageunderstandingDomainBadRequest {
	return &GetLanguageunderstandingDomainBadRequest{}
}

/*
GetLanguageunderstandingDomainBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLanguageunderstandingDomainBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding domain bad request response has a 2xx status code
func (o *GetLanguageunderstandingDomainBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding domain bad request response has a 3xx status code
func (o *GetLanguageunderstandingDomainBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding domain bad request response has a 4xx status code
func (o *GetLanguageunderstandingDomainBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding domain bad request response has a 5xx status code
func (o *GetLanguageunderstandingDomainBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding domain bad request response a status code equal to that given
func (o *GetLanguageunderstandingDomainBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetLanguageunderstandingDomainBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainBadRequest  %+v", 400, o.Payload)
}

func (o *GetLanguageunderstandingDomainBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainBadRequest  %+v", 400, o.Payload)
}

func (o *GetLanguageunderstandingDomainBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainUnauthorized creates a GetLanguageunderstandingDomainUnauthorized with default headers values
func NewGetLanguageunderstandingDomainUnauthorized() *GetLanguageunderstandingDomainUnauthorized {
	return &GetLanguageunderstandingDomainUnauthorized{}
}

/*
GetLanguageunderstandingDomainUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLanguageunderstandingDomainUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding domain unauthorized response has a 2xx status code
func (o *GetLanguageunderstandingDomainUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding domain unauthorized response has a 3xx status code
func (o *GetLanguageunderstandingDomainUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding domain unauthorized response has a 4xx status code
func (o *GetLanguageunderstandingDomainUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding domain unauthorized response has a 5xx status code
func (o *GetLanguageunderstandingDomainUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding domain unauthorized response a status code equal to that given
func (o *GetLanguageunderstandingDomainUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetLanguageunderstandingDomainUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLanguageunderstandingDomainUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLanguageunderstandingDomainUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainForbidden creates a GetLanguageunderstandingDomainForbidden with default headers values
func NewGetLanguageunderstandingDomainForbidden() *GetLanguageunderstandingDomainForbidden {
	return &GetLanguageunderstandingDomainForbidden{}
}

/*
GetLanguageunderstandingDomainForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetLanguageunderstandingDomainForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding domain forbidden response has a 2xx status code
func (o *GetLanguageunderstandingDomainForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding domain forbidden response has a 3xx status code
func (o *GetLanguageunderstandingDomainForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding domain forbidden response has a 4xx status code
func (o *GetLanguageunderstandingDomainForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding domain forbidden response has a 5xx status code
func (o *GetLanguageunderstandingDomainForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding domain forbidden response a status code equal to that given
func (o *GetLanguageunderstandingDomainForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetLanguageunderstandingDomainForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainForbidden  %+v", 403, o.Payload)
}

func (o *GetLanguageunderstandingDomainForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainForbidden  %+v", 403, o.Payload)
}

func (o *GetLanguageunderstandingDomainForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainNotFound creates a GetLanguageunderstandingDomainNotFound with default headers values
func NewGetLanguageunderstandingDomainNotFound() *GetLanguageunderstandingDomainNotFound {
	return &GetLanguageunderstandingDomainNotFound{}
}

/*
GetLanguageunderstandingDomainNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetLanguageunderstandingDomainNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding domain not found response has a 2xx status code
func (o *GetLanguageunderstandingDomainNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding domain not found response has a 3xx status code
func (o *GetLanguageunderstandingDomainNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding domain not found response has a 4xx status code
func (o *GetLanguageunderstandingDomainNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding domain not found response has a 5xx status code
func (o *GetLanguageunderstandingDomainNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding domain not found response a status code equal to that given
func (o *GetLanguageunderstandingDomainNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetLanguageunderstandingDomainNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainNotFound  %+v", 404, o.Payload)
}

func (o *GetLanguageunderstandingDomainNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainNotFound  %+v", 404, o.Payload)
}

func (o *GetLanguageunderstandingDomainNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainRequestTimeout creates a GetLanguageunderstandingDomainRequestTimeout with default headers values
func NewGetLanguageunderstandingDomainRequestTimeout() *GetLanguageunderstandingDomainRequestTimeout {
	return &GetLanguageunderstandingDomainRequestTimeout{}
}

/*
GetLanguageunderstandingDomainRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetLanguageunderstandingDomainRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding domain request timeout response has a 2xx status code
func (o *GetLanguageunderstandingDomainRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding domain request timeout response has a 3xx status code
func (o *GetLanguageunderstandingDomainRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding domain request timeout response has a 4xx status code
func (o *GetLanguageunderstandingDomainRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding domain request timeout response has a 5xx status code
func (o *GetLanguageunderstandingDomainRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding domain request timeout response a status code equal to that given
func (o *GetLanguageunderstandingDomainRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetLanguageunderstandingDomainRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLanguageunderstandingDomainRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLanguageunderstandingDomainRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainRequestEntityTooLarge creates a GetLanguageunderstandingDomainRequestEntityTooLarge with default headers values
func NewGetLanguageunderstandingDomainRequestEntityTooLarge() *GetLanguageunderstandingDomainRequestEntityTooLarge {
	return &GetLanguageunderstandingDomainRequestEntityTooLarge{}
}

/*
GetLanguageunderstandingDomainRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetLanguageunderstandingDomainRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding domain request entity too large response has a 2xx status code
func (o *GetLanguageunderstandingDomainRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding domain request entity too large response has a 3xx status code
func (o *GetLanguageunderstandingDomainRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding domain request entity too large response has a 4xx status code
func (o *GetLanguageunderstandingDomainRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding domain request entity too large response has a 5xx status code
func (o *GetLanguageunderstandingDomainRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding domain request entity too large response a status code equal to that given
func (o *GetLanguageunderstandingDomainRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetLanguageunderstandingDomainRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLanguageunderstandingDomainRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLanguageunderstandingDomainRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainUnsupportedMediaType creates a GetLanguageunderstandingDomainUnsupportedMediaType with default headers values
func NewGetLanguageunderstandingDomainUnsupportedMediaType() *GetLanguageunderstandingDomainUnsupportedMediaType {
	return &GetLanguageunderstandingDomainUnsupportedMediaType{}
}

/*
GetLanguageunderstandingDomainUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLanguageunderstandingDomainUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding domain unsupported media type response has a 2xx status code
func (o *GetLanguageunderstandingDomainUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding domain unsupported media type response has a 3xx status code
func (o *GetLanguageunderstandingDomainUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding domain unsupported media type response has a 4xx status code
func (o *GetLanguageunderstandingDomainUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding domain unsupported media type response has a 5xx status code
func (o *GetLanguageunderstandingDomainUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding domain unsupported media type response a status code equal to that given
func (o *GetLanguageunderstandingDomainUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetLanguageunderstandingDomainUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLanguageunderstandingDomainUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLanguageunderstandingDomainUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainTooManyRequests creates a GetLanguageunderstandingDomainTooManyRequests with default headers values
func NewGetLanguageunderstandingDomainTooManyRequests() *GetLanguageunderstandingDomainTooManyRequests {
	return &GetLanguageunderstandingDomainTooManyRequests{}
}

/*
GetLanguageunderstandingDomainTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetLanguageunderstandingDomainTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding domain too many requests response has a 2xx status code
func (o *GetLanguageunderstandingDomainTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding domain too many requests response has a 3xx status code
func (o *GetLanguageunderstandingDomainTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding domain too many requests response has a 4xx status code
func (o *GetLanguageunderstandingDomainTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languageunderstanding domain too many requests response has a 5xx status code
func (o *GetLanguageunderstandingDomainTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get languageunderstanding domain too many requests response a status code equal to that given
func (o *GetLanguageunderstandingDomainTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetLanguageunderstandingDomainTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLanguageunderstandingDomainTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLanguageunderstandingDomainTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainInternalServerError creates a GetLanguageunderstandingDomainInternalServerError with default headers values
func NewGetLanguageunderstandingDomainInternalServerError() *GetLanguageunderstandingDomainInternalServerError {
	return &GetLanguageunderstandingDomainInternalServerError{}
}

/*
GetLanguageunderstandingDomainInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLanguageunderstandingDomainInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding domain internal server error response has a 2xx status code
func (o *GetLanguageunderstandingDomainInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding domain internal server error response has a 3xx status code
func (o *GetLanguageunderstandingDomainInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding domain internal server error response has a 4xx status code
func (o *GetLanguageunderstandingDomainInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get languageunderstanding domain internal server error response has a 5xx status code
func (o *GetLanguageunderstandingDomainInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get languageunderstanding domain internal server error response a status code equal to that given
func (o *GetLanguageunderstandingDomainInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetLanguageunderstandingDomainInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLanguageunderstandingDomainInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLanguageunderstandingDomainInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainServiceUnavailable creates a GetLanguageunderstandingDomainServiceUnavailable with default headers values
func NewGetLanguageunderstandingDomainServiceUnavailable() *GetLanguageunderstandingDomainServiceUnavailable {
	return &GetLanguageunderstandingDomainServiceUnavailable{}
}

/*
GetLanguageunderstandingDomainServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLanguageunderstandingDomainServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding domain service unavailable response has a 2xx status code
func (o *GetLanguageunderstandingDomainServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding domain service unavailable response has a 3xx status code
func (o *GetLanguageunderstandingDomainServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding domain service unavailable response has a 4xx status code
func (o *GetLanguageunderstandingDomainServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get languageunderstanding domain service unavailable response has a 5xx status code
func (o *GetLanguageunderstandingDomainServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get languageunderstanding domain service unavailable response a status code equal to that given
func (o *GetLanguageunderstandingDomainServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetLanguageunderstandingDomainServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLanguageunderstandingDomainServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLanguageunderstandingDomainServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainGatewayTimeout creates a GetLanguageunderstandingDomainGatewayTimeout with default headers values
func NewGetLanguageunderstandingDomainGatewayTimeout() *GetLanguageunderstandingDomainGatewayTimeout {
	return &GetLanguageunderstandingDomainGatewayTimeout{}
}

/*
GetLanguageunderstandingDomainGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetLanguageunderstandingDomainGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languageunderstanding domain gateway timeout response has a 2xx status code
func (o *GetLanguageunderstandingDomainGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languageunderstanding domain gateway timeout response has a 3xx status code
func (o *GetLanguageunderstandingDomainGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languageunderstanding domain gateway timeout response has a 4xx status code
func (o *GetLanguageunderstandingDomainGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get languageunderstanding domain gateway timeout response has a 5xx status code
func (o *GetLanguageunderstandingDomainGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get languageunderstanding domain gateway timeout response a status code equal to that given
func (o *GetLanguageunderstandingDomainGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetLanguageunderstandingDomainGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLanguageunderstandingDomainGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}][%d] getLanguageunderstandingDomainGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLanguageunderstandingDomainGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
