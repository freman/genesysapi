// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteRoutingSmsPhonenumberReader is a Reader for the DeleteRoutingSmsPhonenumber structure.
type DeleteRoutingSmsPhonenumberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoutingSmsPhonenumberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteRoutingSmsPhonenumberAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRoutingSmsPhonenumberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRoutingSmsPhonenumberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRoutingSmsPhonenumberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRoutingSmsPhonenumberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteRoutingSmsPhonenumberRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteRoutingSmsPhonenumberConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteRoutingSmsPhonenumberRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteRoutingSmsPhonenumberUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteRoutingSmsPhonenumberTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRoutingSmsPhonenumberInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteRoutingSmsPhonenumberServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteRoutingSmsPhonenumberGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRoutingSmsPhonenumberAccepted creates a DeleteRoutingSmsPhonenumberAccepted with default headers values
func NewDeleteRoutingSmsPhonenumberAccepted() *DeleteRoutingSmsPhonenumberAccepted {
	return &DeleteRoutingSmsPhonenumberAccepted{}
}

/*
DeleteRoutingSmsPhonenumberAccepted describes a response with status code 202, with default header values.

Accepted -The phone number delete is in progress.
*/
type DeleteRoutingSmsPhonenumberAccepted struct {
}

// IsSuccess returns true when this delete routing sms phonenumber accepted response has a 2xx status code
func (o *DeleteRoutingSmsPhonenumberAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete routing sms phonenumber accepted response has a 3xx status code
func (o *DeleteRoutingSmsPhonenumberAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing sms phonenumber accepted response has a 4xx status code
func (o *DeleteRoutingSmsPhonenumberAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing sms phonenumber accepted response has a 5xx status code
func (o *DeleteRoutingSmsPhonenumberAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing sms phonenumber accepted response a status code equal to that given
func (o *DeleteRoutingSmsPhonenumberAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *DeleteRoutingSmsPhonenumberAccepted) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberAccepted ", 202)
}

func (o *DeleteRoutingSmsPhonenumberAccepted) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberAccepted ", 202)
}

func (o *DeleteRoutingSmsPhonenumberAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRoutingSmsPhonenumberBadRequest creates a DeleteRoutingSmsPhonenumberBadRequest with default headers values
func NewDeleteRoutingSmsPhonenumberBadRequest() *DeleteRoutingSmsPhonenumberBadRequest {
	return &DeleteRoutingSmsPhonenumberBadRequest{}
}

/*
DeleteRoutingSmsPhonenumberBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteRoutingSmsPhonenumberBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing sms phonenumber bad request response has a 2xx status code
func (o *DeleteRoutingSmsPhonenumberBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing sms phonenumber bad request response has a 3xx status code
func (o *DeleteRoutingSmsPhonenumberBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing sms phonenumber bad request response has a 4xx status code
func (o *DeleteRoutingSmsPhonenumberBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing sms phonenumber bad request response has a 5xx status code
func (o *DeleteRoutingSmsPhonenumberBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing sms phonenumber bad request response a status code equal to that given
func (o *DeleteRoutingSmsPhonenumberBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteRoutingSmsPhonenumberBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSmsPhonenumberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSmsPhonenumberUnauthorized creates a DeleteRoutingSmsPhonenumberUnauthorized with default headers values
func NewDeleteRoutingSmsPhonenumberUnauthorized() *DeleteRoutingSmsPhonenumberUnauthorized {
	return &DeleteRoutingSmsPhonenumberUnauthorized{}
}

/*
DeleteRoutingSmsPhonenumberUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteRoutingSmsPhonenumberUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing sms phonenumber unauthorized response has a 2xx status code
func (o *DeleteRoutingSmsPhonenumberUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing sms phonenumber unauthorized response has a 3xx status code
func (o *DeleteRoutingSmsPhonenumberUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing sms phonenumber unauthorized response has a 4xx status code
func (o *DeleteRoutingSmsPhonenumberUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing sms phonenumber unauthorized response has a 5xx status code
func (o *DeleteRoutingSmsPhonenumberUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing sms phonenumber unauthorized response a status code equal to that given
func (o *DeleteRoutingSmsPhonenumberUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteRoutingSmsPhonenumberUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSmsPhonenumberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSmsPhonenumberForbidden creates a DeleteRoutingSmsPhonenumberForbidden with default headers values
func NewDeleteRoutingSmsPhonenumberForbidden() *DeleteRoutingSmsPhonenumberForbidden {
	return &DeleteRoutingSmsPhonenumberForbidden{}
}

/*
DeleteRoutingSmsPhonenumberForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteRoutingSmsPhonenumberForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing sms phonenumber forbidden response has a 2xx status code
func (o *DeleteRoutingSmsPhonenumberForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing sms phonenumber forbidden response has a 3xx status code
func (o *DeleteRoutingSmsPhonenumberForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing sms phonenumber forbidden response has a 4xx status code
func (o *DeleteRoutingSmsPhonenumberForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing sms phonenumber forbidden response has a 5xx status code
func (o *DeleteRoutingSmsPhonenumberForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing sms phonenumber forbidden response a status code equal to that given
func (o *DeleteRoutingSmsPhonenumberForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteRoutingSmsPhonenumberForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSmsPhonenumberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSmsPhonenumberNotFound creates a DeleteRoutingSmsPhonenumberNotFound with default headers values
func NewDeleteRoutingSmsPhonenumberNotFound() *DeleteRoutingSmsPhonenumberNotFound {
	return &DeleteRoutingSmsPhonenumberNotFound{}
}

/*
DeleteRoutingSmsPhonenumberNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteRoutingSmsPhonenumberNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing sms phonenumber not found response has a 2xx status code
func (o *DeleteRoutingSmsPhonenumberNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing sms phonenumber not found response has a 3xx status code
func (o *DeleteRoutingSmsPhonenumberNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing sms phonenumber not found response has a 4xx status code
func (o *DeleteRoutingSmsPhonenumberNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing sms phonenumber not found response has a 5xx status code
func (o *DeleteRoutingSmsPhonenumberNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing sms phonenumber not found response a status code equal to that given
func (o *DeleteRoutingSmsPhonenumberNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteRoutingSmsPhonenumberNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSmsPhonenumberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSmsPhonenumberRequestTimeout creates a DeleteRoutingSmsPhonenumberRequestTimeout with default headers values
func NewDeleteRoutingSmsPhonenumberRequestTimeout() *DeleteRoutingSmsPhonenumberRequestTimeout {
	return &DeleteRoutingSmsPhonenumberRequestTimeout{}
}

/*
DeleteRoutingSmsPhonenumberRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteRoutingSmsPhonenumberRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing sms phonenumber request timeout response has a 2xx status code
func (o *DeleteRoutingSmsPhonenumberRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing sms phonenumber request timeout response has a 3xx status code
func (o *DeleteRoutingSmsPhonenumberRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing sms phonenumber request timeout response has a 4xx status code
func (o *DeleteRoutingSmsPhonenumberRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing sms phonenumber request timeout response has a 5xx status code
func (o *DeleteRoutingSmsPhonenumberRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing sms phonenumber request timeout response a status code equal to that given
func (o *DeleteRoutingSmsPhonenumberRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteRoutingSmsPhonenumberRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSmsPhonenumberRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSmsPhonenumberConflict creates a DeleteRoutingSmsPhonenumberConflict with default headers values
func NewDeleteRoutingSmsPhonenumberConflict() *DeleteRoutingSmsPhonenumberConflict {
	return &DeleteRoutingSmsPhonenumberConflict{}
}

/*
DeleteRoutingSmsPhonenumberConflict describes a response with status code 409, with default header values.

Conflict
*/
type DeleteRoutingSmsPhonenumberConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing sms phonenumber conflict response has a 2xx status code
func (o *DeleteRoutingSmsPhonenumberConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing sms phonenumber conflict response has a 3xx status code
func (o *DeleteRoutingSmsPhonenumberConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing sms phonenumber conflict response has a 4xx status code
func (o *DeleteRoutingSmsPhonenumberConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing sms phonenumber conflict response has a 5xx status code
func (o *DeleteRoutingSmsPhonenumberConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing sms phonenumber conflict response a status code equal to that given
func (o *DeleteRoutingSmsPhonenumberConflict) IsCode(code int) bool {
	return code == 409
}

func (o *DeleteRoutingSmsPhonenumberConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberConflict  %+v", 409, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberConflict) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberConflict  %+v", 409, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSmsPhonenumberConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSmsPhonenumberRequestEntityTooLarge creates a DeleteRoutingSmsPhonenumberRequestEntityTooLarge with default headers values
func NewDeleteRoutingSmsPhonenumberRequestEntityTooLarge() *DeleteRoutingSmsPhonenumberRequestEntityTooLarge {
	return &DeleteRoutingSmsPhonenumberRequestEntityTooLarge{}
}

/*
DeleteRoutingSmsPhonenumberRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteRoutingSmsPhonenumberRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing sms phonenumber request entity too large response has a 2xx status code
func (o *DeleteRoutingSmsPhonenumberRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing sms phonenumber request entity too large response has a 3xx status code
func (o *DeleteRoutingSmsPhonenumberRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing sms phonenumber request entity too large response has a 4xx status code
func (o *DeleteRoutingSmsPhonenumberRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing sms phonenumber request entity too large response has a 5xx status code
func (o *DeleteRoutingSmsPhonenumberRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing sms phonenumber request entity too large response a status code equal to that given
func (o *DeleteRoutingSmsPhonenumberRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteRoutingSmsPhonenumberRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSmsPhonenumberRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSmsPhonenumberUnsupportedMediaType creates a DeleteRoutingSmsPhonenumberUnsupportedMediaType with default headers values
func NewDeleteRoutingSmsPhonenumberUnsupportedMediaType() *DeleteRoutingSmsPhonenumberUnsupportedMediaType {
	return &DeleteRoutingSmsPhonenumberUnsupportedMediaType{}
}

/*
DeleteRoutingSmsPhonenumberUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteRoutingSmsPhonenumberUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing sms phonenumber unsupported media type response has a 2xx status code
func (o *DeleteRoutingSmsPhonenumberUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing sms phonenumber unsupported media type response has a 3xx status code
func (o *DeleteRoutingSmsPhonenumberUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing sms phonenumber unsupported media type response has a 4xx status code
func (o *DeleteRoutingSmsPhonenumberUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing sms phonenumber unsupported media type response has a 5xx status code
func (o *DeleteRoutingSmsPhonenumberUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing sms phonenumber unsupported media type response a status code equal to that given
func (o *DeleteRoutingSmsPhonenumberUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteRoutingSmsPhonenumberUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSmsPhonenumberUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSmsPhonenumberTooManyRequests creates a DeleteRoutingSmsPhonenumberTooManyRequests with default headers values
func NewDeleteRoutingSmsPhonenumberTooManyRequests() *DeleteRoutingSmsPhonenumberTooManyRequests {
	return &DeleteRoutingSmsPhonenumberTooManyRequests{}
}

/*
DeleteRoutingSmsPhonenumberTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteRoutingSmsPhonenumberTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing sms phonenumber too many requests response has a 2xx status code
func (o *DeleteRoutingSmsPhonenumberTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing sms phonenumber too many requests response has a 3xx status code
func (o *DeleteRoutingSmsPhonenumberTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing sms phonenumber too many requests response has a 4xx status code
func (o *DeleteRoutingSmsPhonenumberTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing sms phonenumber too many requests response has a 5xx status code
func (o *DeleteRoutingSmsPhonenumberTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing sms phonenumber too many requests response a status code equal to that given
func (o *DeleteRoutingSmsPhonenumberTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteRoutingSmsPhonenumberTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSmsPhonenumberTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSmsPhonenumberInternalServerError creates a DeleteRoutingSmsPhonenumberInternalServerError with default headers values
func NewDeleteRoutingSmsPhonenumberInternalServerError() *DeleteRoutingSmsPhonenumberInternalServerError {
	return &DeleteRoutingSmsPhonenumberInternalServerError{}
}

/*
DeleteRoutingSmsPhonenumberInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteRoutingSmsPhonenumberInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing sms phonenumber internal server error response has a 2xx status code
func (o *DeleteRoutingSmsPhonenumberInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing sms phonenumber internal server error response has a 3xx status code
func (o *DeleteRoutingSmsPhonenumberInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing sms phonenumber internal server error response has a 4xx status code
func (o *DeleteRoutingSmsPhonenumberInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing sms phonenumber internal server error response has a 5xx status code
func (o *DeleteRoutingSmsPhonenumberInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete routing sms phonenumber internal server error response a status code equal to that given
func (o *DeleteRoutingSmsPhonenumberInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteRoutingSmsPhonenumberInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSmsPhonenumberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSmsPhonenumberServiceUnavailable creates a DeleteRoutingSmsPhonenumberServiceUnavailable with default headers values
func NewDeleteRoutingSmsPhonenumberServiceUnavailable() *DeleteRoutingSmsPhonenumberServiceUnavailable {
	return &DeleteRoutingSmsPhonenumberServiceUnavailable{}
}

/*
DeleteRoutingSmsPhonenumberServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteRoutingSmsPhonenumberServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing sms phonenumber service unavailable response has a 2xx status code
func (o *DeleteRoutingSmsPhonenumberServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing sms phonenumber service unavailable response has a 3xx status code
func (o *DeleteRoutingSmsPhonenumberServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing sms phonenumber service unavailable response has a 4xx status code
func (o *DeleteRoutingSmsPhonenumberServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing sms phonenumber service unavailable response has a 5xx status code
func (o *DeleteRoutingSmsPhonenumberServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete routing sms phonenumber service unavailable response a status code equal to that given
func (o *DeleteRoutingSmsPhonenumberServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteRoutingSmsPhonenumberServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSmsPhonenumberServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSmsPhonenumberGatewayTimeout creates a DeleteRoutingSmsPhonenumberGatewayTimeout with default headers values
func NewDeleteRoutingSmsPhonenumberGatewayTimeout() *DeleteRoutingSmsPhonenumberGatewayTimeout {
	return &DeleteRoutingSmsPhonenumberGatewayTimeout{}
}

/*
DeleteRoutingSmsPhonenumberGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteRoutingSmsPhonenumberGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing sms phonenumber gateway timeout response has a 2xx status code
func (o *DeleteRoutingSmsPhonenumberGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing sms phonenumber gateway timeout response has a 3xx status code
func (o *DeleteRoutingSmsPhonenumberGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing sms phonenumber gateway timeout response has a 4xx status code
func (o *DeleteRoutingSmsPhonenumberGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing sms phonenumber gateway timeout response has a 5xx status code
func (o *DeleteRoutingSmsPhonenumberGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete routing sms phonenumber gateway timeout response a status code equal to that given
func (o *DeleteRoutingSmsPhonenumberGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteRoutingSmsPhonenumberGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/sms/phonenumbers/{addressId}][%d] deleteRoutingSmsPhonenumberGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteRoutingSmsPhonenumberGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSmsPhonenumberGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
