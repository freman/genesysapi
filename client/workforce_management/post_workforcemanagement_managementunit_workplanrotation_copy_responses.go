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

// PostWorkforcemanagementManagementunitWorkplanrotationCopyReader is a Reader for the PostWorkforcemanagementManagementunitWorkplanrotationCopy structure.
type PostWorkforcemanagementManagementunitWorkplanrotationCopyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWorkforcemanagementManagementunitWorkplanrotationCopyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWorkforcemanagementManagementunitWorkplanrotationCopyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementManagementunitWorkplanrotationCopyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementManagementunitWorkplanrotationCopyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementManagementunitWorkplanrotationCopyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostWorkforcemanagementManagementunitWorkplanrotationCopyRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementManagementunitWorkplanrotationCopyRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementManagementunitWorkplanrotationCopyUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementManagementunitWorkplanrotationCopyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementManagementunitWorkplanrotationCopyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementManagementunitWorkplanrotationCopyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementManagementunitWorkplanrotationCopyGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWorkforcemanagementManagementunitWorkplanrotationCopyOK creates a PostWorkforcemanagementManagementunitWorkplanrotationCopyOK with default headers values
func NewPostWorkforcemanagementManagementunitWorkplanrotationCopyOK() *PostWorkforcemanagementManagementunitWorkplanrotationCopyOK {
	return &PostWorkforcemanagementManagementunitWorkplanrotationCopyOK{}
}

/*PostWorkforcemanagementManagementunitWorkplanrotationCopyOK handles this case with default header values.

successful operation
*/
type PostWorkforcemanagementManagementunitWorkplanrotationCopyOK struct {
	Payload *models.WorkPlanRotationResponse
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}/copy][%d] postWorkforcemanagementManagementunitWorkplanrotationCopyOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyOK) GetPayload() *models.WorkPlanRotationResponse {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkPlanRotationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplanrotationCopyBadRequest creates a PostWorkforcemanagementManagementunitWorkplanrotationCopyBadRequest with default headers values
func NewPostWorkforcemanagementManagementunitWorkplanrotationCopyBadRequest() *PostWorkforcemanagementManagementunitWorkplanrotationCopyBadRequest {
	return &PostWorkforcemanagementManagementunitWorkplanrotationCopyBadRequest{}
}

/*PostWorkforcemanagementManagementunitWorkplanrotationCopyBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementManagementunitWorkplanrotationCopyBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}/copy][%d] postWorkforcemanagementManagementunitWorkplanrotationCopyBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplanrotationCopyUnauthorized creates a PostWorkforcemanagementManagementunitWorkplanrotationCopyUnauthorized with default headers values
func NewPostWorkforcemanagementManagementunitWorkplanrotationCopyUnauthorized() *PostWorkforcemanagementManagementunitWorkplanrotationCopyUnauthorized {
	return &PostWorkforcemanagementManagementunitWorkplanrotationCopyUnauthorized{}
}

/*PostWorkforcemanagementManagementunitWorkplanrotationCopyUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementManagementunitWorkplanrotationCopyUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}/copy][%d] postWorkforcemanagementManagementunitWorkplanrotationCopyUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplanrotationCopyForbidden creates a PostWorkforcemanagementManagementunitWorkplanrotationCopyForbidden with default headers values
func NewPostWorkforcemanagementManagementunitWorkplanrotationCopyForbidden() *PostWorkforcemanagementManagementunitWorkplanrotationCopyForbidden {
	return &PostWorkforcemanagementManagementunitWorkplanrotationCopyForbidden{}
}

/*PostWorkforcemanagementManagementunitWorkplanrotationCopyForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementManagementunitWorkplanrotationCopyForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}/copy][%d] postWorkforcemanagementManagementunitWorkplanrotationCopyForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplanrotationCopyNotFound creates a PostWorkforcemanagementManagementunitWorkplanrotationCopyNotFound with default headers values
func NewPostWorkforcemanagementManagementunitWorkplanrotationCopyNotFound() *PostWorkforcemanagementManagementunitWorkplanrotationCopyNotFound {
	return &PostWorkforcemanagementManagementunitWorkplanrotationCopyNotFound{}
}

/*PostWorkforcemanagementManagementunitWorkplanrotationCopyNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementManagementunitWorkplanrotationCopyNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}/copy][%d] postWorkforcemanagementManagementunitWorkplanrotationCopyNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplanrotationCopyRequestTimeout creates a PostWorkforcemanagementManagementunitWorkplanrotationCopyRequestTimeout with default headers values
func NewPostWorkforcemanagementManagementunitWorkplanrotationCopyRequestTimeout() *PostWorkforcemanagementManagementunitWorkplanrotationCopyRequestTimeout {
	return &PostWorkforcemanagementManagementunitWorkplanrotationCopyRequestTimeout{}
}

/*PostWorkforcemanagementManagementunitWorkplanrotationCopyRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostWorkforcemanagementManagementunitWorkplanrotationCopyRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}/copy][%d] postWorkforcemanagementManagementunitWorkplanrotationCopyRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplanrotationCopyRequestEntityTooLarge creates a PostWorkforcemanagementManagementunitWorkplanrotationCopyRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementManagementunitWorkplanrotationCopyRequestEntityTooLarge() *PostWorkforcemanagementManagementunitWorkplanrotationCopyRequestEntityTooLarge {
	return &PostWorkforcemanagementManagementunitWorkplanrotationCopyRequestEntityTooLarge{}
}

/*PostWorkforcemanagementManagementunitWorkplanrotationCopyRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostWorkforcemanagementManagementunitWorkplanrotationCopyRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}/copy][%d] postWorkforcemanagementManagementunitWorkplanrotationCopyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplanrotationCopyUnsupportedMediaType creates a PostWorkforcemanagementManagementunitWorkplanrotationCopyUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementManagementunitWorkplanrotationCopyUnsupportedMediaType() *PostWorkforcemanagementManagementunitWorkplanrotationCopyUnsupportedMediaType {
	return &PostWorkforcemanagementManagementunitWorkplanrotationCopyUnsupportedMediaType{}
}

/*PostWorkforcemanagementManagementunitWorkplanrotationCopyUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementManagementunitWorkplanrotationCopyUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}/copy][%d] postWorkforcemanagementManagementunitWorkplanrotationCopyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplanrotationCopyTooManyRequests creates a PostWorkforcemanagementManagementunitWorkplanrotationCopyTooManyRequests with default headers values
func NewPostWorkforcemanagementManagementunitWorkplanrotationCopyTooManyRequests() *PostWorkforcemanagementManagementunitWorkplanrotationCopyTooManyRequests {
	return &PostWorkforcemanagementManagementunitWorkplanrotationCopyTooManyRequests{}
}

/*PostWorkforcemanagementManagementunitWorkplanrotationCopyTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostWorkforcemanagementManagementunitWorkplanrotationCopyTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}/copy][%d] postWorkforcemanagementManagementunitWorkplanrotationCopyTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplanrotationCopyInternalServerError creates a PostWorkforcemanagementManagementunitWorkplanrotationCopyInternalServerError with default headers values
func NewPostWorkforcemanagementManagementunitWorkplanrotationCopyInternalServerError() *PostWorkforcemanagementManagementunitWorkplanrotationCopyInternalServerError {
	return &PostWorkforcemanagementManagementunitWorkplanrotationCopyInternalServerError{}
}

/*PostWorkforcemanagementManagementunitWorkplanrotationCopyInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementManagementunitWorkplanrotationCopyInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}/copy][%d] postWorkforcemanagementManagementunitWorkplanrotationCopyInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplanrotationCopyServiceUnavailable creates a PostWorkforcemanagementManagementunitWorkplanrotationCopyServiceUnavailable with default headers values
func NewPostWorkforcemanagementManagementunitWorkplanrotationCopyServiceUnavailable() *PostWorkforcemanagementManagementunitWorkplanrotationCopyServiceUnavailable {
	return &PostWorkforcemanagementManagementunitWorkplanrotationCopyServiceUnavailable{}
}

/*PostWorkforcemanagementManagementunitWorkplanrotationCopyServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementManagementunitWorkplanrotationCopyServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}/copy][%d] postWorkforcemanagementManagementunitWorkplanrotationCopyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementManagementunitWorkplanrotationCopyGatewayTimeout creates a PostWorkforcemanagementManagementunitWorkplanrotationCopyGatewayTimeout with default headers values
func NewPostWorkforcemanagementManagementunitWorkplanrotationCopyGatewayTimeout() *PostWorkforcemanagementManagementunitWorkplanrotationCopyGatewayTimeout {
	return &PostWorkforcemanagementManagementunitWorkplanrotationCopyGatewayTimeout{}
}

/*PostWorkforcemanagementManagementunitWorkplanrotationCopyGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostWorkforcemanagementManagementunitWorkplanrotationCopyGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}/copy][%d] postWorkforcemanagementManagementunitWorkplanrotationCopyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementManagementunitWorkplanrotationCopyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
