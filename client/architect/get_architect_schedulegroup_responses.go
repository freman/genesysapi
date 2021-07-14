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

// GetArchitectSchedulegroupReader is a Reader for the GetArchitectSchedulegroup structure.
type GetArchitectSchedulegroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArchitectSchedulegroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArchitectSchedulegroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetArchitectSchedulegroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetArchitectSchedulegroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetArchitectSchedulegroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetArchitectSchedulegroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetArchitectSchedulegroupRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetArchitectSchedulegroupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetArchitectSchedulegroupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetArchitectSchedulegroupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetArchitectSchedulegroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetArchitectSchedulegroupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetArchitectSchedulegroupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetArchitectSchedulegroupOK creates a GetArchitectSchedulegroupOK with default headers values
func NewGetArchitectSchedulegroupOK() *GetArchitectSchedulegroupOK {
	return &GetArchitectSchedulegroupOK{}
}

/*GetArchitectSchedulegroupOK handles this case with default header values.

successful operation
*/
type GetArchitectSchedulegroupOK struct {
	Payload *models.ScheduleGroup
}

func (o *GetArchitectSchedulegroupOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] getArchitectSchedulegroupOK  %+v", 200, o.Payload)
}

func (o *GetArchitectSchedulegroupOK) GetPayload() *models.ScheduleGroup {
	return o.Payload
}

func (o *GetArchitectSchedulegroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScheduleGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSchedulegroupBadRequest creates a GetArchitectSchedulegroupBadRequest with default headers values
func NewGetArchitectSchedulegroupBadRequest() *GetArchitectSchedulegroupBadRequest {
	return &GetArchitectSchedulegroupBadRequest{}
}

/*GetArchitectSchedulegroupBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetArchitectSchedulegroupBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSchedulegroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] getArchitectSchedulegroupBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectSchedulegroupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSchedulegroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSchedulegroupUnauthorized creates a GetArchitectSchedulegroupUnauthorized with default headers values
func NewGetArchitectSchedulegroupUnauthorized() *GetArchitectSchedulegroupUnauthorized {
	return &GetArchitectSchedulegroupUnauthorized{}
}

/*GetArchitectSchedulegroupUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetArchitectSchedulegroupUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSchedulegroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] getArchitectSchedulegroupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectSchedulegroupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSchedulegroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSchedulegroupForbidden creates a GetArchitectSchedulegroupForbidden with default headers values
func NewGetArchitectSchedulegroupForbidden() *GetArchitectSchedulegroupForbidden {
	return &GetArchitectSchedulegroupForbidden{}
}

/*GetArchitectSchedulegroupForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetArchitectSchedulegroupForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSchedulegroupForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] getArchitectSchedulegroupForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectSchedulegroupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSchedulegroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSchedulegroupNotFound creates a GetArchitectSchedulegroupNotFound with default headers values
func NewGetArchitectSchedulegroupNotFound() *GetArchitectSchedulegroupNotFound {
	return &GetArchitectSchedulegroupNotFound{}
}

/*GetArchitectSchedulegroupNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetArchitectSchedulegroupNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSchedulegroupNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] getArchitectSchedulegroupNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectSchedulegroupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSchedulegroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSchedulegroupRequestTimeout creates a GetArchitectSchedulegroupRequestTimeout with default headers values
func NewGetArchitectSchedulegroupRequestTimeout() *GetArchitectSchedulegroupRequestTimeout {
	return &GetArchitectSchedulegroupRequestTimeout{}
}

/*GetArchitectSchedulegroupRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetArchitectSchedulegroupRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSchedulegroupRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] getArchitectSchedulegroupRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetArchitectSchedulegroupRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSchedulegroupRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSchedulegroupRequestEntityTooLarge creates a GetArchitectSchedulegroupRequestEntityTooLarge with default headers values
func NewGetArchitectSchedulegroupRequestEntityTooLarge() *GetArchitectSchedulegroupRequestEntityTooLarge {
	return &GetArchitectSchedulegroupRequestEntityTooLarge{}
}

/*GetArchitectSchedulegroupRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetArchitectSchedulegroupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSchedulegroupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] getArchitectSchedulegroupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectSchedulegroupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSchedulegroupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSchedulegroupUnsupportedMediaType creates a GetArchitectSchedulegroupUnsupportedMediaType with default headers values
func NewGetArchitectSchedulegroupUnsupportedMediaType() *GetArchitectSchedulegroupUnsupportedMediaType {
	return &GetArchitectSchedulegroupUnsupportedMediaType{}
}

/*GetArchitectSchedulegroupUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetArchitectSchedulegroupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSchedulegroupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] getArchitectSchedulegroupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectSchedulegroupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSchedulegroupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSchedulegroupTooManyRequests creates a GetArchitectSchedulegroupTooManyRequests with default headers values
func NewGetArchitectSchedulegroupTooManyRequests() *GetArchitectSchedulegroupTooManyRequests {
	return &GetArchitectSchedulegroupTooManyRequests{}
}

/*GetArchitectSchedulegroupTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetArchitectSchedulegroupTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSchedulegroupTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] getArchitectSchedulegroupTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectSchedulegroupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSchedulegroupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSchedulegroupInternalServerError creates a GetArchitectSchedulegroupInternalServerError with default headers values
func NewGetArchitectSchedulegroupInternalServerError() *GetArchitectSchedulegroupInternalServerError {
	return &GetArchitectSchedulegroupInternalServerError{}
}

/*GetArchitectSchedulegroupInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetArchitectSchedulegroupInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSchedulegroupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] getArchitectSchedulegroupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectSchedulegroupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSchedulegroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSchedulegroupServiceUnavailable creates a GetArchitectSchedulegroupServiceUnavailable with default headers values
func NewGetArchitectSchedulegroupServiceUnavailable() *GetArchitectSchedulegroupServiceUnavailable {
	return &GetArchitectSchedulegroupServiceUnavailable{}
}

/*GetArchitectSchedulegroupServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetArchitectSchedulegroupServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSchedulegroupServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] getArchitectSchedulegroupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectSchedulegroupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSchedulegroupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSchedulegroupGatewayTimeout creates a GetArchitectSchedulegroupGatewayTimeout with default headers values
func NewGetArchitectSchedulegroupGatewayTimeout() *GetArchitectSchedulegroupGatewayTimeout {
	return &GetArchitectSchedulegroupGatewayTimeout{}
}

/*GetArchitectSchedulegroupGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetArchitectSchedulegroupGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSchedulegroupGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/schedulegroups/{scheduleGroupId}][%d] getArchitectSchedulegroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectSchedulegroupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSchedulegroupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
