// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostWorkforcemanagementManagementunitActivitycodesReader is a Reader for the PostWorkforcemanagementManagementunitActivitycodes structure.
type PostWorkforcemanagementManagementunitActivitycodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementManagementunitActivitycodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewPostWorkforcemanagementManagementunitActivitycodesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementManagementunitActivitycodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementManagementunitActivitycodesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementManagementunitActivitycodesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPostWorkforcemanagementManagementunitActivitycodesGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementManagementunitActivitycodesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementManagementunitActivitycodesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementManagementunitActivitycodesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementManagementunitActivitycodesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPostWorkforcemanagementManagementunitActivitycodesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostWorkforcemanagementManagementunitActivitycodesBadRequest creates a PostWorkforcemanagementManagementunitActivitycodesBadRequest with default headers values
func NewPostWorkforcemanagementManagementunitActivitycodesBadRequest() *PostWorkforcemanagementManagementunitActivitycodesBadRequest {
	return &PostWorkforcemanagementManagementunitActivitycodesBadRequest{}
}

/*PostWorkforcemanagementManagementunitActivitycodesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementManagementunitActivitycodesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitActivitycodesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{muId}/activitycodes][%d] postWorkforcemanagementManagementunitActivitycodesBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitActivitycodesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitActivitycodesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitActivitycodesUnauthorized creates a PostWorkforcemanagementManagementunitActivitycodesUnauthorized with default headers values
func NewPostWorkforcemanagementManagementunitActivitycodesUnauthorized() *PostWorkforcemanagementManagementunitActivitycodesUnauthorized {
	return &PostWorkforcemanagementManagementunitActivitycodesUnauthorized{}
}

/*PostWorkforcemanagementManagementunitActivitycodesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementManagementunitActivitycodesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitActivitycodesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{muId}/activitycodes][%d] postWorkforcemanagementManagementunitActivitycodesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitActivitycodesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitActivitycodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitActivitycodesForbidden creates a PostWorkforcemanagementManagementunitActivitycodesForbidden with default headers values
func NewPostWorkforcemanagementManagementunitActivitycodesForbidden() *PostWorkforcemanagementManagementunitActivitycodesForbidden {
	return &PostWorkforcemanagementManagementunitActivitycodesForbidden{}
}

/*PostWorkforcemanagementManagementunitActivitycodesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementManagementunitActivitycodesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitActivitycodesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{muId}/activitycodes][%d] postWorkforcemanagementManagementunitActivitycodesForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitActivitycodesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitActivitycodesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitActivitycodesNotFound creates a PostWorkforcemanagementManagementunitActivitycodesNotFound with default headers values
func NewPostWorkforcemanagementManagementunitActivitycodesNotFound() *PostWorkforcemanagementManagementunitActivitycodesNotFound {
	return &PostWorkforcemanagementManagementunitActivitycodesNotFound{}
}

/*PostWorkforcemanagementManagementunitActivitycodesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementManagementunitActivitycodesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitActivitycodesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{muId}/activitycodes][%d] postWorkforcemanagementManagementunitActivitycodesNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitActivitycodesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitActivitycodesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitActivitycodesGone creates a PostWorkforcemanagementManagementunitActivitycodesGone with default headers values
func NewPostWorkforcemanagementManagementunitActivitycodesGone() *PostWorkforcemanagementManagementunitActivitycodesGone {
	return &PostWorkforcemanagementManagementunitActivitycodesGone{}
}

/*PostWorkforcemanagementManagementunitActivitycodesGone handles this case with default header values.

Gone
*/
type PostWorkforcemanagementManagementunitActivitycodesGone struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitActivitycodesGone) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{muId}/activitycodes][%d] postWorkforcemanagementManagementunitActivitycodesGone  %+v", 410, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitActivitycodesGone) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitActivitycodesGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge creates a PostWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge() *PostWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge {
	return &PostWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge{}
}

/*PostWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{muId}/activitycodes][%d] postWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType creates a PostWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType() *PostWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType {
	return &PostWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType{}
}

/*PostWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{muId}/activitycodes][%d] postWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitActivitycodesTooManyRequests creates a PostWorkforcemanagementManagementunitActivitycodesTooManyRequests with default headers values
func NewPostWorkforcemanagementManagementunitActivitycodesTooManyRequests() *PostWorkforcemanagementManagementunitActivitycodesTooManyRequests {
	return &PostWorkforcemanagementManagementunitActivitycodesTooManyRequests{}
}

/*PostWorkforcemanagementManagementunitActivitycodesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostWorkforcemanagementManagementunitActivitycodesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitActivitycodesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{muId}/activitycodes][%d] postWorkforcemanagementManagementunitActivitycodesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitActivitycodesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitActivitycodesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitActivitycodesInternalServerError creates a PostWorkforcemanagementManagementunitActivitycodesInternalServerError with default headers values
func NewPostWorkforcemanagementManagementunitActivitycodesInternalServerError() *PostWorkforcemanagementManagementunitActivitycodesInternalServerError {
	return &PostWorkforcemanagementManagementunitActivitycodesInternalServerError{}
}

/*PostWorkforcemanagementManagementunitActivitycodesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementManagementunitActivitycodesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitActivitycodesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{muId}/activitycodes][%d] postWorkforcemanagementManagementunitActivitycodesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitActivitycodesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitActivitycodesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitActivitycodesServiceUnavailable creates a PostWorkforcemanagementManagementunitActivitycodesServiceUnavailable with default headers values
func NewPostWorkforcemanagementManagementunitActivitycodesServiceUnavailable() *PostWorkforcemanagementManagementunitActivitycodesServiceUnavailable {
	return &PostWorkforcemanagementManagementunitActivitycodesServiceUnavailable{}
}

/*PostWorkforcemanagementManagementunitActivitycodesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementManagementunitActivitycodesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitActivitycodesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{muId}/activitycodes][%d] postWorkforcemanagementManagementunitActivitycodesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitActivitycodesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitActivitycodesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitActivitycodesGatewayTimeout creates a PostWorkforcemanagementManagementunitActivitycodesGatewayTimeout with default headers values
func NewPostWorkforcemanagementManagementunitActivitycodesGatewayTimeout() *PostWorkforcemanagementManagementunitActivitycodesGatewayTimeout {
	return &PostWorkforcemanagementManagementunitActivitycodesGatewayTimeout{}
}

/*PostWorkforcemanagementManagementunitActivitycodesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostWorkforcemanagementManagementunitActivitycodesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitActivitycodesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{muId}/activitycodes][%d] postWorkforcemanagementManagementunitActivitycodesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitActivitycodesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitActivitycodesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitActivitycodesDefault creates a PostWorkforcemanagementManagementunitActivitycodesDefault with default headers values
func NewPostWorkforcemanagementManagementunitActivitycodesDefault(code int) *PostWorkforcemanagementManagementunitActivitycodesDefault {
	return &PostWorkforcemanagementManagementunitActivitycodesDefault{
		_statusCode: code,
	}
}

/*PostWorkforcemanagementManagementunitActivitycodesDefault handles this case with default header values.

successful operation
*/
type PostWorkforcemanagementManagementunitActivitycodesDefault struct {
	_statusCode int
}

// Code gets the status code for the post workforcemanagement managementunit activitycodes default response
func (o *PostWorkforcemanagementManagementunitActivitycodesDefault) Code() int {
	return o._statusCode
}

func (o *PostWorkforcemanagementManagementunitActivitycodesDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{muId}/activitycodes][%d] postWorkforcemanagementManagementunitActivitycodes default ", o._statusCode)
}

func (o *PostWorkforcemanagementManagementunitActivitycodesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}