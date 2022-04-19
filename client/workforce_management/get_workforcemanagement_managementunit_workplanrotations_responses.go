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

// GetWorkforcemanagementManagementunitWorkplanrotationsReader is a Reader for the GetWorkforcemanagementManagementunitWorkplanrotations structure.
type GetWorkforcemanagementManagementunitWorkplanrotationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementManagementunitWorkplanrotationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementManagementunitWorkplanrotationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementManagementunitWorkplanrotationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementManagementunitWorkplanrotationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementManagementunitWorkplanrotationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementManagementunitWorkplanrotationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementManagementunitWorkplanrotationsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementManagementunitWorkplanrotationsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementManagementunitWorkplanrotationsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementManagementunitWorkplanrotationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementManagementunitWorkplanrotationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementManagementunitWorkplanrotationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementManagementunitWorkplanrotationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementManagementunitWorkplanrotationsOK creates a GetWorkforcemanagementManagementunitWorkplanrotationsOK with default headers values
func NewGetWorkforcemanagementManagementunitWorkplanrotationsOK() *GetWorkforcemanagementManagementunitWorkplanrotationsOK {
	return &GetWorkforcemanagementManagementunitWorkplanrotationsOK{}
}

/*GetWorkforcemanagementManagementunitWorkplanrotationsOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementManagementunitWorkplanrotationsOK struct {
	Payload *models.WorkPlanRotationListResponse
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations][%d] getWorkforcemanagementManagementunitWorkplanrotationsOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsOK) GetPayload() *models.WorkPlanRotationListResponse {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkPlanRotationListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitWorkplanrotationsBadRequest creates a GetWorkforcemanagementManagementunitWorkplanrotationsBadRequest with default headers values
func NewGetWorkforcemanagementManagementunitWorkplanrotationsBadRequest() *GetWorkforcemanagementManagementunitWorkplanrotationsBadRequest {
	return &GetWorkforcemanagementManagementunitWorkplanrotationsBadRequest{}
}

/*GetWorkforcemanagementManagementunitWorkplanrotationsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementManagementunitWorkplanrotationsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations][%d] getWorkforcemanagementManagementunitWorkplanrotationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitWorkplanrotationsUnauthorized creates a GetWorkforcemanagementManagementunitWorkplanrotationsUnauthorized with default headers values
func NewGetWorkforcemanagementManagementunitWorkplanrotationsUnauthorized() *GetWorkforcemanagementManagementunitWorkplanrotationsUnauthorized {
	return &GetWorkforcemanagementManagementunitWorkplanrotationsUnauthorized{}
}

/*GetWorkforcemanagementManagementunitWorkplanrotationsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementManagementunitWorkplanrotationsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations][%d] getWorkforcemanagementManagementunitWorkplanrotationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitWorkplanrotationsForbidden creates a GetWorkforcemanagementManagementunitWorkplanrotationsForbidden with default headers values
func NewGetWorkforcemanagementManagementunitWorkplanrotationsForbidden() *GetWorkforcemanagementManagementunitWorkplanrotationsForbidden {
	return &GetWorkforcemanagementManagementunitWorkplanrotationsForbidden{}
}

/*GetWorkforcemanagementManagementunitWorkplanrotationsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementManagementunitWorkplanrotationsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations][%d] getWorkforcemanagementManagementunitWorkplanrotationsForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitWorkplanrotationsNotFound creates a GetWorkforcemanagementManagementunitWorkplanrotationsNotFound with default headers values
func NewGetWorkforcemanagementManagementunitWorkplanrotationsNotFound() *GetWorkforcemanagementManagementunitWorkplanrotationsNotFound {
	return &GetWorkforcemanagementManagementunitWorkplanrotationsNotFound{}
}

/*GetWorkforcemanagementManagementunitWorkplanrotationsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementManagementunitWorkplanrotationsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations][%d] getWorkforcemanagementManagementunitWorkplanrotationsNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitWorkplanrotationsRequestTimeout creates a GetWorkforcemanagementManagementunitWorkplanrotationsRequestTimeout with default headers values
func NewGetWorkforcemanagementManagementunitWorkplanrotationsRequestTimeout() *GetWorkforcemanagementManagementunitWorkplanrotationsRequestTimeout {
	return &GetWorkforcemanagementManagementunitWorkplanrotationsRequestTimeout{}
}

/*GetWorkforcemanagementManagementunitWorkplanrotationsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementManagementunitWorkplanrotationsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations][%d] getWorkforcemanagementManagementunitWorkplanrotationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitWorkplanrotationsRequestEntityTooLarge creates a GetWorkforcemanagementManagementunitWorkplanrotationsRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementManagementunitWorkplanrotationsRequestEntityTooLarge() *GetWorkforcemanagementManagementunitWorkplanrotationsRequestEntityTooLarge {
	return &GetWorkforcemanagementManagementunitWorkplanrotationsRequestEntityTooLarge{}
}

/*GetWorkforcemanagementManagementunitWorkplanrotationsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetWorkforcemanagementManagementunitWorkplanrotationsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations][%d] getWorkforcemanagementManagementunitWorkplanrotationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitWorkplanrotationsUnsupportedMediaType creates a GetWorkforcemanagementManagementunitWorkplanrotationsUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementManagementunitWorkplanrotationsUnsupportedMediaType() *GetWorkforcemanagementManagementunitWorkplanrotationsUnsupportedMediaType {
	return &GetWorkforcemanagementManagementunitWorkplanrotationsUnsupportedMediaType{}
}

/*GetWorkforcemanagementManagementunitWorkplanrotationsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementManagementunitWorkplanrotationsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations][%d] getWorkforcemanagementManagementunitWorkplanrotationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitWorkplanrotationsTooManyRequests creates a GetWorkforcemanagementManagementunitWorkplanrotationsTooManyRequests with default headers values
func NewGetWorkforcemanagementManagementunitWorkplanrotationsTooManyRequests() *GetWorkforcemanagementManagementunitWorkplanrotationsTooManyRequests {
	return &GetWorkforcemanagementManagementunitWorkplanrotationsTooManyRequests{}
}

/*GetWorkforcemanagementManagementunitWorkplanrotationsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementManagementunitWorkplanrotationsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations][%d] getWorkforcemanagementManagementunitWorkplanrotationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitWorkplanrotationsInternalServerError creates a GetWorkforcemanagementManagementunitWorkplanrotationsInternalServerError with default headers values
func NewGetWorkforcemanagementManagementunitWorkplanrotationsInternalServerError() *GetWorkforcemanagementManagementunitWorkplanrotationsInternalServerError {
	return &GetWorkforcemanagementManagementunitWorkplanrotationsInternalServerError{}
}

/*GetWorkforcemanagementManagementunitWorkplanrotationsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementManagementunitWorkplanrotationsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations][%d] getWorkforcemanagementManagementunitWorkplanrotationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitWorkplanrotationsServiceUnavailable creates a GetWorkforcemanagementManagementunitWorkplanrotationsServiceUnavailable with default headers values
func NewGetWorkforcemanagementManagementunitWorkplanrotationsServiceUnavailable() *GetWorkforcemanagementManagementunitWorkplanrotationsServiceUnavailable {
	return &GetWorkforcemanagementManagementunitWorkplanrotationsServiceUnavailable{}
}

/*GetWorkforcemanagementManagementunitWorkplanrotationsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementManagementunitWorkplanrotationsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations][%d] getWorkforcemanagementManagementunitWorkplanrotationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitWorkplanrotationsGatewayTimeout creates a GetWorkforcemanagementManagementunitWorkplanrotationsGatewayTimeout with default headers values
func NewGetWorkforcemanagementManagementunitWorkplanrotationsGatewayTimeout() *GetWorkforcemanagementManagementunitWorkplanrotationsGatewayTimeout {
	return &GetWorkforcemanagementManagementunitWorkplanrotationsGatewayTimeout{}
}

/*GetWorkforcemanagementManagementunitWorkplanrotationsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementManagementunitWorkplanrotationsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations][%d] getWorkforcemanagementManagementunitWorkplanrotationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitWorkplanrotationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
