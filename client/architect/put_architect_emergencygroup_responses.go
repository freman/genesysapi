// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutArchitectEmergencygroupReader is a Reader for the PutArchitectEmergencygroup structure.
type PutArchitectEmergencygroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutArchitectEmergencygroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutArchitectEmergencygroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutArchitectEmergencygroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutArchitectEmergencygroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutArchitectEmergencygroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutArchitectEmergencygroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutArchitectEmergencygroupRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutArchitectEmergencygroupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutArchitectEmergencygroupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutArchitectEmergencygroupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutArchitectEmergencygroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutArchitectEmergencygroupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutArchitectEmergencygroupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutArchitectEmergencygroupOK creates a PutArchitectEmergencygroupOK with default headers values
func NewPutArchitectEmergencygroupOK() *PutArchitectEmergencygroupOK {
	return &PutArchitectEmergencygroupOK{}
}

/*PutArchitectEmergencygroupOK handles this case with default header values.

successful operation
*/
type PutArchitectEmergencygroupOK struct {
	Payload *models.EmergencyGroup
}

func (o *PutArchitectEmergencygroupOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/emergencygroups/{emergencyGroupId}][%d] putArchitectEmergencygroupOK  %+v", 200, o.Payload)
}

func (o *PutArchitectEmergencygroupOK) GetPayload() *models.EmergencyGroup {
	return o.Payload
}

func (o *PutArchitectEmergencygroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmergencyGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectEmergencygroupBadRequest creates a PutArchitectEmergencygroupBadRequest with default headers values
func NewPutArchitectEmergencygroupBadRequest() *PutArchitectEmergencygroupBadRequest {
	return &PutArchitectEmergencygroupBadRequest{}
}

/*PutArchitectEmergencygroupBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutArchitectEmergencygroupBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutArchitectEmergencygroupBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/emergencygroups/{emergencyGroupId}][%d] putArchitectEmergencygroupBadRequest  %+v", 400, o.Payload)
}

func (o *PutArchitectEmergencygroupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectEmergencygroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectEmergencygroupUnauthorized creates a PutArchitectEmergencygroupUnauthorized with default headers values
func NewPutArchitectEmergencygroupUnauthorized() *PutArchitectEmergencygroupUnauthorized {
	return &PutArchitectEmergencygroupUnauthorized{}
}

/*PutArchitectEmergencygroupUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutArchitectEmergencygroupUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutArchitectEmergencygroupUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/emergencygroups/{emergencyGroupId}][%d] putArchitectEmergencygroupUnauthorized  %+v", 401, o.Payload)
}

func (o *PutArchitectEmergencygroupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectEmergencygroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectEmergencygroupForbidden creates a PutArchitectEmergencygroupForbidden with default headers values
func NewPutArchitectEmergencygroupForbidden() *PutArchitectEmergencygroupForbidden {
	return &PutArchitectEmergencygroupForbidden{}
}

/*PutArchitectEmergencygroupForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutArchitectEmergencygroupForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutArchitectEmergencygroupForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/emergencygroups/{emergencyGroupId}][%d] putArchitectEmergencygroupForbidden  %+v", 403, o.Payload)
}

func (o *PutArchitectEmergencygroupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectEmergencygroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectEmergencygroupNotFound creates a PutArchitectEmergencygroupNotFound with default headers values
func NewPutArchitectEmergencygroupNotFound() *PutArchitectEmergencygroupNotFound {
	return &PutArchitectEmergencygroupNotFound{}
}

/*PutArchitectEmergencygroupNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutArchitectEmergencygroupNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutArchitectEmergencygroupNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/emergencygroups/{emergencyGroupId}][%d] putArchitectEmergencygroupNotFound  %+v", 404, o.Payload)
}

func (o *PutArchitectEmergencygroupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectEmergencygroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectEmergencygroupRequestTimeout creates a PutArchitectEmergencygroupRequestTimeout with default headers values
func NewPutArchitectEmergencygroupRequestTimeout() *PutArchitectEmergencygroupRequestTimeout {
	return &PutArchitectEmergencygroupRequestTimeout{}
}

/*PutArchitectEmergencygroupRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutArchitectEmergencygroupRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutArchitectEmergencygroupRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/emergencygroups/{emergencyGroupId}][%d] putArchitectEmergencygroupRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutArchitectEmergencygroupRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectEmergencygroupRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectEmergencygroupRequestEntityTooLarge creates a PutArchitectEmergencygroupRequestEntityTooLarge with default headers values
func NewPutArchitectEmergencygroupRequestEntityTooLarge() *PutArchitectEmergencygroupRequestEntityTooLarge {
	return &PutArchitectEmergencygroupRequestEntityTooLarge{}
}

/*PutArchitectEmergencygroupRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutArchitectEmergencygroupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutArchitectEmergencygroupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/emergencygroups/{emergencyGroupId}][%d] putArchitectEmergencygroupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutArchitectEmergencygroupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectEmergencygroupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectEmergencygroupUnsupportedMediaType creates a PutArchitectEmergencygroupUnsupportedMediaType with default headers values
func NewPutArchitectEmergencygroupUnsupportedMediaType() *PutArchitectEmergencygroupUnsupportedMediaType {
	return &PutArchitectEmergencygroupUnsupportedMediaType{}
}

/*PutArchitectEmergencygroupUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutArchitectEmergencygroupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutArchitectEmergencygroupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/emergencygroups/{emergencyGroupId}][%d] putArchitectEmergencygroupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutArchitectEmergencygroupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectEmergencygroupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectEmergencygroupTooManyRequests creates a PutArchitectEmergencygroupTooManyRequests with default headers values
func NewPutArchitectEmergencygroupTooManyRequests() *PutArchitectEmergencygroupTooManyRequests {
	return &PutArchitectEmergencygroupTooManyRequests{}
}

/*PutArchitectEmergencygroupTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutArchitectEmergencygroupTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutArchitectEmergencygroupTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/emergencygroups/{emergencyGroupId}][%d] putArchitectEmergencygroupTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutArchitectEmergencygroupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectEmergencygroupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectEmergencygroupInternalServerError creates a PutArchitectEmergencygroupInternalServerError with default headers values
func NewPutArchitectEmergencygroupInternalServerError() *PutArchitectEmergencygroupInternalServerError {
	return &PutArchitectEmergencygroupInternalServerError{}
}

/*PutArchitectEmergencygroupInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutArchitectEmergencygroupInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutArchitectEmergencygroupInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/emergencygroups/{emergencyGroupId}][%d] putArchitectEmergencygroupInternalServerError  %+v", 500, o.Payload)
}

func (o *PutArchitectEmergencygroupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectEmergencygroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectEmergencygroupServiceUnavailable creates a PutArchitectEmergencygroupServiceUnavailable with default headers values
func NewPutArchitectEmergencygroupServiceUnavailable() *PutArchitectEmergencygroupServiceUnavailable {
	return &PutArchitectEmergencygroupServiceUnavailable{}
}

/*PutArchitectEmergencygroupServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutArchitectEmergencygroupServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutArchitectEmergencygroupServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/emergencygroups/{emergencyGroupId}][%d] putArchitectEmergencygroupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutArchitectEmergencygroupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectEmergencygroupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutArchitectEmergencygroupGatewayTimeout creates a PutArchitectEmergencygroupGatewayTimeout with default headers values
func NewPutArchitectEmergencygroupGatewayTimeout() *PutArchitectEmergencygroupGatewayTimeout {
	return &PutArchitectEmergencygroupGatewayTimeout{}
}

/*PutArchitectEmergencygroupGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutArchitectEmergencygroupGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutArchitectEmergencygroupGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/architect/emergencygroups/{emergencyGroupId}][%d] putArchitectEmergencygroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutArchitectEmergencygroupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutArchitectEmergencygroupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
