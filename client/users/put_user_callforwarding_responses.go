// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutUserCallforwardingReader is a Reader for the PutUserCallforwarding structure.
type PutUserCallforwardingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutUserCallforwardingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutUserCallforwardingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutUserCallforwardingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutUserCallforwardingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutUserCallforwardingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutUserCallforwardingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutUserCallforwardingRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutUserCallforwardingConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutUserCallforwardingRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutUserCallforwardingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutUserCallforwardingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutUserCallforwardingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutUserCallforwardingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutUserCallforwardingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutUserCallforwardingOK creates a PutUserCallforwardingOK with default headers values
func NewPutUserCallforwardingOK() *PutUserCallforwardingOK {
	return &PutUserCallforwardingOK{}
}

/*PutUserCallforwardingOK handles this case with default header values.

successful operation
*/
type PutUserCallforwardingOK struct {
	Payload *models.CallForwarding
}

func (o *PutUserCallforwardingOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/callforwarding][%d] putUserCallforwardingOK  %+v", 200, o.Payload)
}

func (o *PutUserCallforwardingOK) GetPayload() *models.CallForwarding {
	return o.Payload
}

func (o *PutUserCallforwardingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CallForwarding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserCallforwardingBadRequest creates a PutUserCallforwardingBadRequest with default headers values
func NewPutUserCallforwardingBadRequest() *PutUserCallforwardingBadRequest {
	return &PutUserCallforwardingBadRequest{}
}

/*PutUserCallforwardingBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutUserCallforwardingBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutUserCallforwardingBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/callforwarding][%d] putUserCallforwardingBadRequest  %+v", 400, o.Payload)
}

func (o *PutUserCallforwardingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserCallforwardingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserCallforwardingUnauthorized creates a PutUserCallforwardingUnauthorized with default headers values
func NewPutUserCallforwardingUnauthorized() *PutUserCallforwardingUnauthorized {
	return &PutUserCallforwardingUnauthorized{}
}

/*PutUserCallforwardingUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutUserCallforwardingUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutUserCallforwardingUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/callforwarding][%d] putUserCallforwardingUnauthorized  %+v", 401, o.Payload)
}

func (o *PutUserCallforwardingUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserCallforwardingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserCallforwardingForbidden creates a PutUserCallforwardingForbidden with default headers values
func NewPutUserCallforwardingForbidden() *PutUserCallforwardingForbidden {
	return &PutUserCallforwardingForbidden{}
}

/*PutUserCallforwardingForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutUserCallforwardingForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutUserCallforwardingForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/callforwarding][%d] putUserCallforwardingForbidden  %+v", 403, o.Payload)
}

func (o *PutUserCallforwardingForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserCallforwardingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserCallforwardingNotFound creates a PutUserCallforwardingNotFound with default headers values
func NewPutUserCallforwardingNotFound() *PutUserCallforwardingNotFound {
	return &PutUserCallforwardingNotFound{}
}

/*PutUserCallforwardingNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutUserCallforwardingNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutUserCallforwardingNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/callforwarding][%d] putUserCallforwardingNotFound  %+v", 404, o.Payload)
}

func (o *PutUserCallforwardingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserCallforwardingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserCallforwardingRequestTimeout creates a PutUserCallforwardingRequestTimeout with default headers values
func NewPutUserCallforwardingRequestTimeout() *PutUserCallforwardingRequestTimeout {
	return &PutUserCallforwardingRequestTimeout{}
}

/*PutUserCallforwardingRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutUserCallforwardingRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutUserCallforwardingRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/callforwarding][%d] putUserCallforwardingRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutUserCallforwardingRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserCallforwardingRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserCallforwardingConflict creates a PutUserCallforwardingConflict with default headers values
func NewPutUserCallforwardingConflict() *PutUserCallforwardingConflict {
	return &PutUserCallforwardingConflict{}
}

/*PutUserCallforwardingConflict handles this case with default header values.

Conflict
*/
type PutUserCallforwardingConflict struct {
	Payload *models.ErrorBody
}

func (o *PutUserCallforwardingConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/callforwarding][%d] putUserCallforwardingConflict  %+v", 409, o.Payload)
}

func (o *PutUserCallforwardingConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserCallforwardingConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserCallforwardingRequestEntityTooLarge creates a PutUserCallforwardingRequestEntityTooLarge with default headers values
func NewPutUserCallforwardingRequestEntityTooLarge() *PutUserCallforwardingRequestEntityTooLarge {
	return &PutUserCallforwardingRequestEntityTooLarge{}
}

/*PutUserCallforwardingRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutUserCallforwardingRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutUserCallforwardingRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/callforwarding][%d] putUserCallforwardingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutUserCallforwardingRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserCallforwardingRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserCallforwardingUnsupportedMediaType creates a PutUserCallforwardingUnsupportedMediaType with default headers values
func NewPutUserCallforwardingUnsupportedMediaType() *PutUserCallforwardingUnsupportedMediaType {
	return &PutUserCallforwardingUnsupportedMediaType{}
}

/*PutUserCallforwardingUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutUserCallforwardingUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutUserCallforwardingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/callforwarding][%d] putUserCallforwardingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutUserCallforwardingUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserCallforwardingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserCallforwardingTooManyRequests creates a PutUserCallforwardingTooManyRequests with default headers values
func NewPutUserCallforwardingTooManyRequests() *PutUserCallforwardingTooManyRequests {
	return &PutUserCallforwardingTooManyRequests{}
}

/*PutUserCallforwardingTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutUserCallforwardingTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutUserCallforwardingTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/callforwarding][%d] putUserCallforwardingTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutUserCallforwardingTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserCallforwardingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserCallforwardingInternalServerError creates a PutUserCallforwardingInternalServerError with default headers values
func NewPutUserCallforwardingInternalServerError() *PutUserCallforwardingInternalServerError {
	return &PutUserCallforwardingInternalServerError{}
}

/*PutUserCallforwardingInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutUserCallforwardingInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutUserCallforwardingInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/callforwarding][%d] putUserCallforwardingInternalServerError  %+v", 500, o.Payload)
}

func (o *PutUserCallforwardingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserCallforwardingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserCallforwardingServiceUnavailable creates a PutUserCallforwardingServiceUnavailable with default headers values
func NewPutUserCallforwardingServiceUnavailable() *PutUserCallforwardingServiceUnavailable {
	return &PutUserCallforwardingServiceUnavailable{}
}

/*PutUserCallforwardingServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutUserCallforwardingServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutUserCallforwardingServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/callforwarding][%d] putUserCallforwardingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutUserCallforwardingServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserCallforwardingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserCallforwardingGatewayTimeout creates a PutUserCallforwardingGatewayTimeout with default headers values
func NewPutUserCallforwardingGatewayTimeout() *PutUserCallforwardingGatewayTimeout {
	return &PutUserCallforwardingGatewayTimeout{}
}

/*PutUserCallforwardingGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutUserCallforwardingGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutUserCallforwardingGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/callforwarding][%d] putUserCallforwardingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutUserCallforwardingGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserCallforwardingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
