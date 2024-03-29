// Code generated by go-swagger; DO NOT EDIT.

package external_contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteExternalcontactsContactReader is a Reader for the DeleteExternalcontactsContact structure.
type DeleteExternalcontactsContactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteExternalcontactsContactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteExternalcontactsContactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteExternalcontactsContactBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteExternalcontactsContactUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteExternalcontactsContactForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteExternalcontactsContactNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteExternalcontactsContactRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteExternalcontactsContactRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteExternalcontactsContactUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteExternalcontactsContactTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteExternalcontactsContactInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteExternalcontactsContactServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteExternalcontactsContactGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteExternalcontactsContactOK creates a DeleteExternalcontactsContactOK with default headers values
func NewDeleteExternalcontactsContactOK() *DeleteExternalcontactsContactOK {
	return &DeleteExternalcontactsContactOK{}
}

/*
DeleteExternalcontactsContactOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteExternalcontactsContactOK struct {
	Payload models.Empty
}

// IsSuccess returns true when this delete externalcontacts contact o k response has a 2xx status code
func (o *DeleteExternalcontactsContactOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete externalcontacts contact o k response has a 3xx status code
func (o *DeleteExternalcontactsContactOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contact o k response has a 4xx status code
func (o *DeleteExternalcontactsContactOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete externalcontacts contact o k response has a 5xx status code
func (o *DeleteExternalcontactsContactOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts contact o k response a status code equal to that given
func (o *DeleteExternalcontactsContactOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteExternalcontactsContactOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactOK  %+v", 200, o.Payload)
}

func (o *DeleteExternalcontactsContactOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactOK  %+v", 200, o.Payload)
}

func (o *DeleteExternalcontactsContactOK) GetPayload() models.Empty {
	return o.Payload
}

func (o *DeleteExternalcontactsContactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactBadRequest creates a DeleteExternalcontactsContactBadRequest with default headers values
func NewDeleteExternalcontactsContactBadRequest() *DeleteExternalcontactsContactBadRequest {
	return &DeleteExternalcontactsContactBadRequest{}
}

/*
DeleteExternalcontactsContactBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteExternalcontactsContactBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts contact bad request response has a 2xx status code
func (o *DeleteExternalcontactsContactBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts contact bad request response has a 3xx status code
func (o *DeleteExternalcontactsContactBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contact bad request response has a 4xx status code
func (o *DeleteExternalcontactsContactBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts contact bad request response has a 5xx status code
func (o *DeleteExternalcontactsContactBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts contact bad request response a status code equal to that given
func (o *DeleteExternalcontactsContactBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteExternalcontactsContactBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteExternalcontactsContactBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteExternalcontactsContactBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactUnauthorized creates a DeleteExternalcontactsContactUnauthorized with default headers values
func NewDeleteExternalcontactsContactUnauthorized() *DeleteExternalcontactsContactUnauthorized {
	return &DeleteExternalcontactsContactUnauthorized{}
}

/*
DeleteExternalcontactsContactUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteExternalcontactsContactUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts contact unauthorized response has a 2xx status code
func (o *DeleteExternalcontactsContactUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts contact unauthorized response has a 3xx status code
func (o *DeleteExternalcontactsContactUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contact unauthorized response has a 4xx status code
func (o *DeleteExternalcontactsContactUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts contact unauthorized response has a 5xx status code
func (o *DeleteExternalcontactsContactUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts contact unauthorized response a status code equal to that given
func (o *DeleteExternalcontactsContactUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteExternalcontactsContactUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteExternalcontactsContactUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteExternalcontactsContactUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactForbidden creates a DeleteExternalcontactsContactForbidden with default headers values
func NewDeleteExternalcontactsContactForbidden() *DeleteExternalcontactsContactForbidden {
	return &DeleteExternalcontactsContactForbidden{}
}

/*
DeleteExternalcontactsContactForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteExternalcontactsContactForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts contact forbidden response has a 2xx status code
func (o *DeleteExternalcontactsContactForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts contact forbidden response has a 3xx status code
func (o *DeleteExternalcontactsContactForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contact forbidden response has a 4xx status code
func (o *DeleteExternalcontactsContactForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts contact forbidden response has a 5xx status code
func (o *DeleteExternalcontactsContactForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts contact forbidden response a status code equal to that given
func (o *DeleteExternalcontactsContactForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteExternalcontactsContactForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactForbidden  %+v", 403, o.Payload)
}

func (o *DeleteExternalcontactsContactForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactForbidden  %+v", 403, o.Payload)
}

func (o *DeleteExternalcontactsContactForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactNotFound creates a DeleteExternalcontactsContactNotFound with default headers values
func NewDeleteExternalcontactsContactNotFound() *DeleteExternalcontactsContactNotFound {
	return &DeleteExternalcontactsContactNotFound{}
}

/*
DeleteExternalcontactsContactNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteExternalcontactsContactNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts contact not found response has a 2xx status code
func (o *DeleteExternalcontactsContactNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts contact not found response has a 3xx status code
func (o *DeleteExternalcontactsContactNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contact not found response has a 4xx status code
func (o *DeleteExternalcontactsContactNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts contact not found response has a 5xx status code
func (o *DeleteExternalcontactsContactNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts contact not found response a status code equal to that given
func (o *DeleteExternalcontactsContactNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteExternalcontactsContactNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactNotFound  %+v", 404, o.Payload)
}

func (o *DeleteExternalcontactsContactNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactNotFound  %+v", 404, o.Payload)
}

func (o *DeleteExternalcontactsContactNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactRequestTimeout creates a DeleteExternalcontactsContactRequestTimeout with default headers values
func NewDeleteExternalcontactsContactRequestTimeout() *DeleteExternalcontactsContactRequestTimeout {
	return &DeleteExternalcontactsContactRequestTimeout{}
}

/*
DeleteExternalcontactsContactRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteExternalcontactsContactRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts contact request timeout response has a 2xx status code
func (o *DeleteExternalcontactsContactRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts contact request timeout response has a 3xx status code
func (o *DeleteExternalcontactsContactRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contact request timeout response has a 4xx status code
func (o *DeleteExternalcontactsContactRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts contact request timeout response has a 5xx status code
func (o *DeleteExternalcontactsContactRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts contact request timeout response a status code equal to that given
func (o *DeleteExternalcontactsContactRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteExternalcontactsContactRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteExternalcontactsContactRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteExternalcontactsContactRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactRequestEntityTooLarge creates a DeleteExternalcontactsContactRequestEntityTooLarge with default headers values
func NewDeleteExternalcontactsContactRequestEntityTooLarge() *DeleteExternalcontactsContactRequestEntityTooLarge {
	return &DeleteExternalcontactsContactRequestEntityTooLarge{}
}

/*
DeleteExternalcontactsContactRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteExternalcontactsContactRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts contact request entity too large response has a 2xx status code
func (o *DeleteExternalcontactsContactRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts contact request entity too large response has a 3xx status code
func (o *DeleteExternalcontactsContactRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contact request entity too large response has a 4xx status code
func (o *DeleteExternalcontactsContactRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts contact request entity too large response has a 5xx status code
func (o *DeleteExternalcontactsContactRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts contact request entity too large response a status code equal to that given
func (o *DeleteExternalcontactsContactRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteExternalcontactsContactRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteExternalcontactsContactRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteExternalcontactsContactRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactUnsupportedMediaType creates a DeleteExternalcontactsContactUnsupportedMediaType with default headers values
func NewDeleteExternalcontactsContactUnsupportedMediaType() *DeleteExternalcontactsContactUnsupportedMediaType {
	return &DeleteExternalcontactsContactUnsupportedMediaType{}
}

/*
DeleteExternalcontactsContactUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteExternalcontactsContactUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts contact unsupported media type response has a 2xx status code
func (o *DeleteExternalcontactsContactUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts contact unsupported media type response has a 3xx status code
func (o *DeleteExternalcontactsContactUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contact unsupported media type response has a 4xx status code
func (o *DeleteExternalcontactsContactUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts contact unsupported media type response has a 5xx status code
func (o *DeleteExternalcontactsContactUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts contact unsupported media type response a status code equal to that given
func (o *DeleteExternalcontactsContactUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteExternalcontactsContactUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteExternalcontactsContactUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteExternalcontactsContactUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactTooManyRequests creates a DeleteExternalcontactsContactTooManyRequests with default headers values
func NewDeleteExternalcontactsContactTooManyRequests() *DeleteExternalcontactsContactTooManyRequests {
	return &DeleteExternalcontactsContactTooManyRequests{}
}

/*
DeleteExternalcontactsContactTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteExternalcontactsContactTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts contact too many requests response has a 2xx status code
func (o *DeleteExternalcontactsContactTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts contact too many requests response has a 3xx status code
func (o *DeleteExternalcontactsContactTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contact too many requests response has a 4xx status code
func (o *DeleteExternalcontactsContactTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts contact too many requests response has a 5xx status code
func (o *DeleteExternalcontactsContactTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts contact too many requests response a status code equal to that given
func (o *DeleteExternalcontactsContactTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteExternalcontactsContactTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteExternalcontactsContactTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteExternalcontactsContactTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactInternalServerError creates a DeleteExternalcontactsContactInternalServerError with default headers values
func NewDeleteExternalcontactsContactInternalServerError() *DeleteExternalcontactsContactInternalServerError {
	return &DeleteExternalcontactsContactInternalServerError{}
}

/*
DeleteExternalcontactsContactInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteExternalcontactsContactInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts contact internal server error response has a 2xx status code
func (o *DeleteExternalcontactsContactInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts contact internal server error response has a 3xx status code
func (o *DeleteExternalcontactsContactInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contact internal server error response has a 4xx status code
func (o *DeleteExternalcontactsContactInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete externalcontacts contact internal server error response has a 5xx status code
func (o *DeleteExternalcontactsContactInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete externalcontacts contact internal server error response a status code equal to that given
func (o *DeleteExternalcontactsContactInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteExternalcontactsContactInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteExternalcontactsContactInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteExternalcontactsContactInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactServiceUnavailable creates a DeleteExternalcontactsContactServiceUnavailable with default headers values
func NewDeleteExternalcontactsContactServiceUnavailable() *DeleteExternalcontactsContactServiceUnavailable {
	return &DeleteExternalcontactsContactServiceUnavailable{}
}

/*
DeleteExternalcontactsContactServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteExternalcontactsContactServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts contact service unavailable response has a 2xx status code
func (o *DeleteExternalcontactsContactServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts contact service unavailable response has a 3xx status code
func (o *DeleteExternalcontactsContactServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contact service unavailable response has a 4xx status code
func (o *DeleteExternalcontactsContactServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete externalcontacts contact service unavailable response has a 5xx status code
func (o *DeleteExternalcontactsContactServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete externalcontacts contact service unavailable response a status code equal to that given
func (o *DeleteExternalcontactsContactServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteExternalcontactsContactServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteExternalcontactsContactServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteExternalcontactsContactServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactGatewayTimeout creates a DeleteExternalcontactsContactGatewayTimeout with default headers values
func NewDeleteExternalcontactsContactGatewayTimeout() *DeleteExternalcontactsContactGatewayTimeout {
	return &DeleteExternalcontactsContactGatewayTimeout{}
}

/*
DeleteExternalcontactsContactGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteExternalcontactsContactGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts contact gateway timeout response has a 2xx status code
func (o *DeleteExternalcontactsContactGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts contact gateway timeout response has a 3xx status code
func (o *DeleteExternalcontactsContactGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contact gateway timeout response has a 4xx status code
func (o *DeleteExternalcontactsContactGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete externalcontacts contact gateway timeout response has a 5xx status code
func (o *DeleteExternalcontactsContactGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete externalcontacts contact gateway timeout response a status code equal to that given
func (o *DeleteExternalcontactsContactGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteExternalcontactsContactGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteExternalcontactsContactGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}][%d] deleteExternalcontactsContactGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteExternalcontactsContactGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
