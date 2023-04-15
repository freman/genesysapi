// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetOutboundContactlistContactReader is a Reader for the GetOutboundContactlistContact structure.
type GetOutboundContactlistContactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundContactlistContactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundContactlistContactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundContactlistContactBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundContactlistContactUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundContactlistContactForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundContactlistContactNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundContactlistContactRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundContactlistContactRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundContactlistContactUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundContactlistContactTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundContactlistContactInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundContactlistContactServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundContactlistContactGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundContactlistContactOK creates a GetOutboundContactlistContactOK with default headers values
func NewGetOutboundContactlistContactOK() *GetOutboundContactlistContactOK {
	return &GetOutboundContactlistContactOK{}
}

/*
GetOutboundContactlistContactOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOutboundContactlistContactOK struct {
	Payload *models.DialerContact
}

// IsSuccess returns true when this get outbound contactlist contact o k response has a 2xx status code
func (o *GetOutboundContactlistContactOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get outbound contactlist contact o k response has a 3xx status code
func (o *GetOutboundContactlistContactOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist contact o k response has a 4xx status code
func (o *GetOutboundContactlistContactOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound contactlist contact o k response has a 5xx status code
func (o *GetOutboundContactlistContactOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound contactlist contact o k response a status code equal to that given
func (o *GetOutboundContactlistContactOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOutboundContactlistContactOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactOK  %+v", 200, o.Payload)
}

func (o *GetOutboundContactlistContactOK) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactOK  %+v", 200, o.Payload)
}

func (o *GetOutboundContactlistContactOK) GetPayload() *models.DialerContact {
	return o.Payload
}

func (o *GetOutboundContactlistContactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DialerContact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistContactBadRequest creates a GetOutboundContactlistContactBadRequest with default headers values
func NewGetOutboundContactlistContactBadRequest() *GetOutboundContactlistContactBadRequest {
	return &GetOutboundContactlistContactBadRequest{}
}

/*
GetOutboundContactlistContactBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundContactlistContactBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound contactlist contact bad request response has a 2xx status code
func (o *GetOutboundContactlistContactBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound contactlist contact bad request response has a 3xx status code
func (o *GetOutboundContactlistContactBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist contact bad request response has a 4xx status code
func (o *GetOutboundContactlistContactBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound contactlist contact bad request response has a 5xx status code
func (o *GetOutboundContactlistContactBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound contactlist contact bad request response a status code equal to that given
func (o *GetOutboundContactlistContactBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOutboundContactlistContactBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundContactlistContactBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundContactlistContactBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistContactBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistContactUnauthorized creates a GetOutboundContactlistContactUnauthorized with default headers values
func NewGetOutboundContactlistContactUnauthorized() *GetOutboundContactlistContactUnauthorized {
	return &GetOutboundContactlistContactUnauthorized{}
}

/*
GetOutboundContactlistContactUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundContactlistContactUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound contactlist contact unauthorized response has a 2xx status code
func (o *GetOutboundContactlistContactUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound contactlist contact unauthorized response has a 3xx status code
func (o *GetOutboundContactlistContactUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist contact unauthorized response has a 4xx status code
func (o *GetOutboundContactlistContactUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound contactlist contact unauthorized response has a 5xx status code
func (o *GetOutboundContactlistContactUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound contactlist contact unauthorized response a status code equal to that given
func (o *GetOutboundContactlistContactUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOutboundContactlistContactUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundContactlistContactUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundContactlistContactUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistContactUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistContactForbidden creates a GetOutboundContactlistContactForbidden with default headers values
func NewGetOutboundContactlistContactForbidden() *GetOutboundContactlistContactForbidden {
	return &GetOutboundContactlistContactForbidden{}
}

/*
GetOutboundContactlistContactForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundContactlistContactForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound contactlist contact forbidden response has a 2xx status code
func (o *GetOutboundContactlistContactForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound contactlist contact forbidden response has a 3xx status code
func (o *GetOutboundContactlistContactForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist contact forbidden response has a 4xx status code
func (o *GetOutboundContactlistContactForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound contactlist contact forbidden response has a 5xx status code
func (o *GetOutboundContactlistContactForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound contactlist contact forbidden response a status code equal to that given
func (o *GetOutboundContactlistContactForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOutboundContactlistContactForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundContactlistContactForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundContactlistContactForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistContactForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistContactNotFound creates a GetOutboundContactlistContactNotFound with default headers values
func NewGetOutboundContactlistContactNotFound() *GetOutboundContactlistContactNotFound {
	return &GetOutboundContactlistContactNotFound{}
}

/*
GetOutboundContactlistContactNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOutboundContactlistContactNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound contactlist contact not found response has a 2xx status code
func (o *GetOutboundContactlistContactNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound contactlist contact not found response has a 3xx status code
func (o *GetOutboundContactlistContactNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist contact not found response has a 4xx status code
func (o *GetOutboundContactlistContactNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound contactlist contact not found response has a 5xx status code
func (o *GetOutboundContactlistContactNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound contactlist contact not found response a status code equal to that given
func (o *GetOutboundContactlistContactNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOutboundContactlistContactNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundContactlistContactNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundContactlistContactNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistContactNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistContactRequestTimeout creates a GetOutboundContactlistContactRequestTimeout with default headers values
func NewGetOutboundContactlistContactRequestTimeout() *GetOutboundContactlistContactRequestTimeout {
	return &GetOutboundContactlistContactRequestTimeout{}
}

/*
GetOutboundContactlistContactRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundContactlistContactRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound contactlist contact request timeout response has a 2xx status code
func (o *GetOutboundContactlistContactRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound contactlist contact request timeout response has a 3xx status code
func (o *GetOutboundContactlistContactRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist contact request timeout response has a 4xx status code
func (o *GetOutboundContactlistContactRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound contactlist contact request timeout response has a 5xx status code
func (o *GetOutboundContactlistContactRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound contactlist contact request timeout response a status code equal to that given
func (o *GetOutboundContactlistContactRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOutboundContactlistContactRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundContactlistContactRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundContactlistContactRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistContactRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistContactRequestEntityTooLarge creates a GetOutboundContactlistContactRequestEntityTooLarge with default headers values
func NewGetOutboundContactlistContactRequestEntityTooLarge() *GetOutboundContactlistContactRequestEntityTooLarge {
	return &GetOutboundContactlistContactRequestEntityTooLarge{}
}

/*
GetOutboundContactlistContactRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetOutboundContactlistContactRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound contactlist contact request entity too large response has a 2xx status code
func (o *GetOutboundContactlistContactRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound contactlist contact request entity too large response has a 3xx status code
func (o *GetOutboundContactlistContactRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist contact request entity too large response has a 4xx status code
func (o *GetOutboundContactlistContactRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound contactlist contact request entity too large response has a 5xx status code
func (o *GetOutboundContactlistContactRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound contactlist contact request entity too large response a status code equal to that given
func (o *GetOutboundContactlistContactRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOutboundContactlistContactRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundContactlistContactRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundContactlistContactRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistContactRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistContactUnsupportedMediaType creates a GetOutboundContactlistContactUnsupportedMediaType with default headers values
func NewGetOutboundContactlistContactUnsupportedMediaType() *GetOutboundContactlistContactUnsupportedMediaType {
	return &GetOutboundContactlistContactUnsupportedMediaType{}
}

/*
GetOutboundContactlistContactUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundContactlistContactUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound contactlist contact unsupported media type response has a 2xx status code
func (o *GetOutboundContactlistContactUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound contactlist contact unsupported media type response has a 3xx status code
func (o *GetOutboundContactlistContactUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist contact unsupported media type response has a 4xx status code
func (o *GetOutboundContactlistContactUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound contactlist contact unsupported media type response has a 5xx status code
func (o *GetOutboundContactlistContactUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound contactlist contact unsupported media type response a status code equal to that given
func (o *GetOutboundContactlistContactUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOutboundContactlistContactUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundContactlistContactUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundContactlistContactUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistContactUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistContactTooManyRequests creates a GetOutboundContactlistContactTooManyRequests with default headers values
func NewGetOutboundContactlistContactTooManyRequests() *GetOutboundContactlistContactTooManyRequests {
	return &GetOutboundContactlistContactTooManyRequests{}
}

/*
GetOutboundContactlistContactTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundContactlistContactTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound contactlist contact too many requests response has a 2xx status code
func (o *GetOutboundContactlistContactTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound contactlist contact too many requests response has a 3xx status code
func (o *GetOutboundContactlistContactTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist contact too many requests response has a 4xx status code
func (o *GetOutboundContactlistContactTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound contactlist contact too many requests response has a 5xx status code
func (o *GetOutboundContactlistContactTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound contactlist contact too many requests response a status code equal to that given
func (o *GetOutboundContactlistContactTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOutboundContactlistContactTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundContactlistContactTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundContactlistContactTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistContactTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistContactInternalServerError creates a GetOutboundContactlistContactInternalServerError with default headers values
func NewGetOutboundContactlistContactInternalServerError() *GetOutboundContactlistContactInternalServerError {
	return &GetOutboundContactlistContactInternalServerError{}
}

/*
GetOutboundContactlistContactInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundContactlistContactInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound contactlist contact internal server error response has a 2xx status code
func (o *GetOutboundContactlistContactInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound contactlist contact internal server error response has a 3xx status code
func (o *GetOutboundContactlistContactInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist contact internal server error response has a 4xx status code
func (o *GetOutboundContactlistContactInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound contactlist contact internal server error response has a 5xx status code
func (o *GetOutboundContactlistContactInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound contactlist contact internal server error response a status code equal to that given
func (o *GetOutboundContactlistContactInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOutboundContactlistContactInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundContactlistContactInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundContactlistContactInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistContactInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistContactServiceUnavailable creates a GetOutboundContactlistContactServiceUnavailable with default headers values
func NewGetOutboundContactlistContactServiceUnavailable() *GetOutboundContactlistContactServiceUnavailable {
	return &GetOutboundContactlistContactServiceUnavailable{}
}

/*
GetOutboundContactlistContactServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundContactlistContactServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound contactlist contact service unavailable response has a 2xx status code
func (o *GetOutboundContactlistContactServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound contactlist contact service unavailable response has a 3xx status code
func (o *GetOutboundContactlistContactServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist contact service unavailable response has a 4xx status code
func (o *GetOutboundContactlistContactServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound contactlist contact service unavailable response has a 5xx status code
func (o *GetOutboundContactlistContactServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound contactlist contact service unavailable response a status code equal to that given
func (o *GetOutboundContactlistContactServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOutboundContactlistContactServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundContactlistContactServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundContactlistContactServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistContactServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistContactGatewayTimeout creates a GetOutboundContactlistContactGatewayTimeout with default headers values
func NewGetOutboundContactlistContactGatewayTimeout() *GetOutboundContactlistContactGatewayTimeout {
	return &GetOutboundContactlistContactGatewayTimeout{}
}

/*
GetOutboundContactlistContactGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOutboundContactlistContactGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound contactlist contact gateway timeout response has a 2xx status code
func (o *GetOutboundContactlistContactGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound contactlist contact gateway timeout response has a 3xx status code
func (o *GetOutboundContactlistContactGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound contactlist contact gateway timeout response has a 4xx status code
func (o *GetOutboundContactlistContactGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound contactlist contact gateway timeout response has a 5xx status code
func (o *GetOutboundContactlistContactGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound contactlist contact gateway timeout response a status code equal to that given
func (o *GetOutboundContactlistContactGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOutboundContactlistContactGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundContactlistContactGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundContactlistContactGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistContactGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
