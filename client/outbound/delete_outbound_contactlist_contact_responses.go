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

// DeleteOutboundContactlistContactReader is a Reader for the DeleteOutboundContactlistContact structure.
type DeleteOutboundContactlistContactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOutboundContactlistContactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOutboundContactlistContactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteOutboundContactlistContactBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteOutboundContactlistContactUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteOutboundContactlistContactForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOutboundContactlistContactNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteOutboundContactlistContactRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteOutboundContactlistContactRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteOutboundContactlistContactUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteOutboundContactlistContactTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOutboundContactlistContactInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteOutboundContactlistContactServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteOutboundContactlistContactGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOutboundContactlistContactOK creates a DeleteOutboundContactlistContactOK with default headers values
func NewDeleteOutboundContactlistContactOK() *DeleteOutboundContactlistContactOK {
	return &DeleteOutboundContactlistContactOK{}
}

/*
DeleteOutboundContactlistContactOK describes a response with status code 200, with default header values.

Operation was successful.
*/
type DeleteOutboundContactlistContactOK struct {
}

// IsSuccess returns true when this delete outbound contactlist contact o k response has a 2xx status code
func (o *DeleteOutboundContactlistContactOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete outbound contactlist contact o k response has a 3xx status code
func (o *DeleteOutboundContactlistContactOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound contactlist contact o k response has a 4xx status code
func (o *DeleteOutboundContactlistContactOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete outbound contactlist contact o k response has a 5xx status code
func (o *DeleteOutboundContactlistContactOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound contactlist contact o k response a status code equal to that given
func (o *DeleteOutboundContactlistContactOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteOutboundContactlistContactOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactOK ", 200)
}

func (o *DeleteOutboundContactlistContactOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactOK ", 200)
}

func (o *DeleteOutboundContactlistContactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOutboundContactlistContactBadRequest creates a DeleteOutboundContactlistContactBadRequest with default headers values
func NewDeleteOutboundContactlistContactBadRequest() *DeleteOutboundContactlistContactBadRequest {
	return &DeleteOutboundContactlistContactBadRequest{}
}

/*
DeleteOutboundContactlistContactBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteOutboundContactlistContactBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound contactlist contact bad request response has a 2xx status code
func (o *DeleteOutboundContactlistContactBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound contactlist contact bad request response has a 3xx status code
func (o *DeleteOutboundContactlistContactBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound contactlist contact bad request response has a 4xx status code
func (o *DeleteOutboundContactlistContactBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound contactlist contact bad request response has a 5xx status code
func (o *DeleteOutboundContactlistContactBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound contactlist contact bad request response a status code equal to that given
func (o *DeleteOutboundContactlistContactBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteOutboundContactlistContactBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOutboundContactlistContactBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOutboundContactlistContactBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistContactBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistContactUnauthorized creates a DeleteOutboundContactlistContactUnauthorized with default headers values
func NewDeleteOutboundContactlistContactUnauthorized() *DeleteOutboundContactlistContactUnauthorized {
	return &DeleteOutboundContactlistContactUnauthorized{}
}

/*
DeleteOutboundContactlistContactUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteOutboundContactlistContactUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound contactlist contact unauthorized response has a 2xx status code
func (o *DeleteOutboundContactlistContactUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound contactlist contact unauthorized response has a 3xx status code
func (o *DeleteOutboundContactlistContactUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound contactlist contact unauthorized response has a 4xx status code
func (o *DeleteOutboundContactlistContactUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound contactlist contact unauthorized response has a 5xx status code
func (o *DeleteOutboundContactlistContactUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound contactlist contact unauthorized response a status code equal to that given
func (o *DeleteOutboundContactlistContactUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteOutboundContactlistContactUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOutboundContactlistContactUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOutboundContactlistContactUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistContactUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistContactForbidden creates a DeleteOutboundContactlistContactForbidden with default headers values
func NewDeleteOutboundContactlistContactForbidden() *DeleteOutboundContactlistContactForbidden {
	return &DeleteOutboundContactlistContactForbidden{}
}

/*
DeleteOutboundContactlistContactForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteOutboundContactlistContactForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound contactlist contact forbidden response has a 2xx status code
func (o *DeleteOutboundContactlistContactForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound contactlist contact forbidden response has a 3xx status code
func (o *DeleteOutboundContactlistContactForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound contactlist contact forbidden response has a 4xx status code
func (o *DeleteOutboundContactlistContactForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound contactlist contact forbidden response has a 5xx status code
func (o *DeleteOutboundContactlistContactForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound contactlist contact forbidden response a status code equal to that given
func (o *DeleteOutboundContactlistContactForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteOutboundContactlistContactForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOutboundContactlistContactForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOutboundContactlistContactForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistContactForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistContactNotFound creates a DeleteOutboundContactlistContactNotFound with default headers values
func NewDeleteOutboundContactlistContactNotFound() *DeleteOutboundContactlistContactNotFound {
	return &DeleteOutboundContactlistContactNotFound{}
}

/*
DeleteOutboundContactlistContactNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteOutboundContactlistContactNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound contactlist contact not found response has a 2xx status code
func (o *DeleteOutboundContactlistContactNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound contactlist contact not found response has a 3xx status code
func (o *DeleteOutboundContactlistContactNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound contactlist contact not found response has a 4xx status code
func (o *DeleteOutboundContactlistContactNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound contactlist contact not found response has a 5xx status code
func (o *DeleteOutboundContactlistContactNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound contactlist contact not found response a status code equal to that given
func (o *DeleteOutboundContactlistContactNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteOutboundContactlistContactNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOutboundContactlistContactNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOutboundContactlistContactNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistContactNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistContactRequestTimeout creates a DeleteOutboundContactlistContactRequestTimeout with default headers values
func NewDeleteOutboundContactlistContactRequestTimeout() *DeleteOutboundContactlistContactRequestTimeout {
	return &DeleteOutboundContactlistContactRequestTimeout{}
}

/*
DeleteOutboundContactlistContactRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteOutboundContactlistContactRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound contactlist contact request timeout response has a 2xx status code
func (o *DeleteOutboundContactlistContactRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound contactlist contact request timeout response has a 3xx status code
func (o *DeleteOutboundContactlistContactRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound contactlist contact request timeout response has a 4xx status code
func (o *DeleteOutboundContactlistContactRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound contactlist contact request timeout response has a 5xx status code
func (o *DeleteOutboundContactlistContactRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound contactlist contact request timeout response a status code equal to that given
func (o *DeleteOutboundContactlistContactRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteOutboundContactlistContactRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteOutboundContactlistContactRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteOutboundContactlistContactRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistContactRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistContactRequestEntityTooLarge creates a DeleteOutboundContactlistContactRequestEntityTooLarge with default headers values
func NewDeleteOutboundContactlistContactRequestEntityTooLarge() *DeleteOutboundContactlistContactRequestEntityTooLarge {
	return &DeleteOutboundContactlistContactRequestEntityTooLarge{}
}

/*
DeleteOutboundContactlistContactRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteOutboundContactlistContactRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound contactlist contact request entity too large response has a 2xx status code
func (o *DeleteOutboundContactlistContactRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound contactlist contact request entity too large response has a 3xx status code
func (o *DeleteOutboundContactlistContactRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound contactlist contact request entity too large response has a 4xx status code
func (o *DeleteOutboundContactlistContactRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound contactlist contact request entity too large response has a 5xx status code
func (o *DeleteOutboundContactlistContactRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound contactlist contact request entity too large response a status code equal to that given
func (o *DeleteOutboundContactlistContactRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteOutboundContactlistContactRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteOutboundContactlistContactRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteOutboundContactlistContactRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistContactRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistContactUnsupportedMediaType creates a DeleteOutboundContactlistContactUnsupportedMediaType with default headers values
func NewDeleteOutboundContactlistContactUnsupportedMediaType() *DeleteOutboundContactlistContactUnsupportedMediaType {
	return &DeleteOutboundContactlistContactUnsupportedMediaType{}
}

/*
DeleteOutboundContactlistContactUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteOutboundContactlistContactUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound contactlist contact unsupported media type response has a 2xx status code
func (o *DeleteOutboundContactlistContactUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound contactlist contact unsupported media type response has a 3xx status code
func (o *DeleteOutboundContactlistContactUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound contactlist contact unsupported media type response has a 4xx status code
func (o *DeleteOutboundContactlistContactUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound contactlist contact unsupported media type response has a 5xx status code
func (o *DeleteOutboundContactlistContactUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound contactlist contact unsupported media type response a status code equal to that given
func (o *DeleteOutboundContactlistContactUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteOutboundContactlistContactUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteOutboundContactlistContactUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteOutboundContactlistContactUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistContactUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistContactTooManyRequests creates a DeleteOutboundContactlistContactTooManyRequests with default headers values
func NewDeleteOutboundContactlistContactTooManyRequests() *DeleteOutboundContactlistContactTooManyRequests {
	return &DeleteOutboundContactlistContactTooManyRequests{}
}

/*
DeleteOutboundContactlistContactTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteOutboundContactlistContactTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound contactlist contact too many requests response has a 2xx status code
func (o *DeleteOutboundContactlistContactTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound contactlist contact too many requests response has a 3xx status code
func (o *DeleteOutboundContactlistContactTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound contactlist contact too many requests response has a 4xx status code
func (o *DeleteOutboundContactlistContactTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound contactlist contact too many requests response has a 5xx status code
func (o *DeleteOutboundContactlistContactTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound contactlist contact too many requests response a status code equal to that given
func (o *DeleteOutboundContactlistContactTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteOutboundContactlistContactTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteOutboundContactlistContactTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteOutboundContactlistContactTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistContactTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistContactInternalServerError creates a DeleteOutboundContactlistContactInternalServerError with default headers values
func NewDeleteOutboundContactlistContactInternalServerError() *DeleteOutboundContactlistContactInternalServerError {
	return &DeleteOutboundContactlistContactInternalServerError{}
}

/*
DeleteOutboundContactlistContactInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteOutboundContactlistContactInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound contactlist contact internal server error response has a 2xx status code
func (o *DeleteOutboundContactlistContactInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound contactlist contact internal server error response has a 3xx status code
func (o *DeleteOutboundContactlistContactInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound contactlist contact internal server error response has a 4xx status code
func (o *DeleteOutboundContactlistContactInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete outbound contactlist contact internal server error response has a 5xx status code
func (o *DeleteOutboundContactlistContactInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete outbound contactlist contact internal server error response a status code equal to that given
func (o *DeleteOutboundContactlistContactInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteOutboundContactlistContactInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOutboundContactlistContactInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOutboundContactlistContactInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistContactInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistContactServiceUnavailable creates a DeleteOutboundContactlistContactServiceUnavailable with default headers values
func NewDeleteOutboundContactlistContactServiceUnavailable() *DeleteOutboundContactlistContactServiceUnavailable {
	return &DeleteOutboundContactlistContactServiceUnavailable{}
}

/*
DeleteOutboundContactlistContactServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteOutboundContactlistContactServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound contactlist contact service unavailable response has a 2xx status code
func (o *DeleteOutboundContactlistContactServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound contactlist contact service unavailable response has a 3xx status code
func (o *DeleteOutboundContactlistContactServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound contactlist contact service unavailable response has a 4xx status code
func (o *DeleteOutboundContactlistContactServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete outbound contactlist contact service unavailable response has a 5xx status code
func (o *DeleteOutboundContactlistContactServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete outbound contactlist contact service unavailable response a status code equal to that given
func (o *DeleteOutboundContactlistContactServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteOutboundContactlistContactServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteOutboundContactlistContactServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteOutboundContactlistContactServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistContactServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistContactGatewayTimeout creates a DeleteOutboundContactlistContactGatewayTimeout with default headers values
func NewDeleteOutboundContactlistContactGatewayTimeout() *DeleteOutboundContactlistContactGatewayTimeout {
	return &DeleteOutboundContactlistContactGatewayTimeout{}
}

/*
DeleteOutboundContactlistContactGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteOutboundContactlistContactGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound contactlist contact gateway timeout response has a 2xx status code
func (o *DeleteOutboundContactlistContactGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound contactlist contact gateway timeout response has a 3xx status code
func (o *DeleteOutboundContactlistContactGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound contactlist contact gateway timeout response has a 4xx status code
func (o *DeleteOutboundContactlistContactGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete outbound contactlist contact gateway timeout response has a 5xx status code
func (o *DeleteOutboundContactlistContactGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete outbound contactlist contact gateway timeout response a status code equal to that given
func (o *DeleteOutboundContactlistContactGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteOutboundContactlistContactGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteOutboundContactlistContactGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] deleteOutboundContactlistContactGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteOutboundContactlistContactGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistContactGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
