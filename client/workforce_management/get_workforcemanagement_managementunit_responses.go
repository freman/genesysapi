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

// GetWorkforcemanagementManagementunitReader is a Reader for the GetWorkforcemanagementManagementunit structure.
type GetWorkforcemanagementManagementunitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementManagementunitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementManagementunitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementManagementunitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementManagementunitUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementManagementunitForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementManagementunitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementManagementunitRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementManagementunitUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementManagementunitTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementManagementunitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementManagementunitServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementManagementunitGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementManagementunitOK creates a GetWorkforcemanagementManagementunitOK with default headers values
func NewGetWorkforcemanagementManagementunitOK() *GetWorkforcemanagementManagementunitOK {
	return &GetWorkforcemanagementManagementunitOK{}
}

/*GetWorkforcemanagementManagementunitOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementManagementunitOK struct {
	Payload *models.ManagementUnit
}

func (o *GetWorkforcemanagementManagementunitOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}][%d] getWorkforcemanagementManagementunitOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitOK) GetPayload() *models.ManagementUnit {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManagementUnit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitBadRequest creates a GetWorkforcemanagementManagementunitBadRequest with default headers values
func NewGetWorkforcemanagementManagementunitBadRequest() *GetWorkforcemanagementManagementunitBadRequest {
	return &GetWorkforcemanagementManagementunitBadRequest{}
}

/*GetWorkforcemanagementManagementunitBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementManagementunitBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}][%d] getWorkforcemanagementManagementunitBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUnauthorized creates a GetWorkforcemanagementManagementunitUnauthorized with default headers values
func NewGetWorkforcemanagementManagementunitUnauthorized() *GetWorkforcemanagementManagementunitUnauthorized {
	return &GetWorkforcemanagementManagementunitUnauthorized{}
}

/*GetWorkforcemanagementManagementunitUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementManagementunitUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}][%d] getWorkforcemanagementManagementunitUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitForbidden creates a GetWorkforcemanagementManagementunitForbidden with default headers values
func NewGetWorkforcemanagementManagementunitForbidden() *GetWorkforcemanagementManagementunitForbidden {
	return &GetWorkforcemanagementManagementunitForbidden{}
}

/*GetWorkforcemanagementManagementunitForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementManagementunitForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}][%d] getWorkforcemanagementManagementunitForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitNotFound creates a GetWorkforcemanagementManagementunitNotFound with default headers values
func NewGetWorkforcemanagementManagementunitNotFound() *GetWorkforcemanagementManagementunitNotFound {
	return &GetWorkforcemanagementManagementunitNotFound{}
}

/*GetWorkforcemanagementManagementunitNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementManagementunitNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}][%d] getWorkforcemanagementManagementunitNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitRequestEntityTooLarge creates a GetWorkforcemanagementManagementunitRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementManagementunitRequestEntityTooLarge() *GetWorkforcemanagementManagementunitRequestEntityTooLarge {
	return &GetWorkforcemanagementManagementunitRequestEntityTooLarge{}
}

/*GetWorkforcemanagementManagementunitRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetWorkforcemanagementManagementunitRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}][%d] getWorkforcemanagementManagementunitRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUnsupportedMediaType creates a GetWorkforcemanagementManagementunitUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementManagementunitUnsupportedMediaType() *GetWorkforcemanagementManagementunitUnsupportedMediaType {
	return &GetWorkforcemanagementManagementunitUnsupportedMediaType{}
}

/*GetWorkforcemanagementManagementunitUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementManagementunitUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}][%d] getWorkforcemanagementManagementunitUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTooManyRequests creates a GetWorkforcemanagementManagementunitTooManyRequests with default headers values
func NewGetWorkforcemanagementManagementunitTooManyRequests() *GetWorkforcemanagementManagementunitTooManyRequests {
	return &GetWorkforcemanagementManagementunitTooManyRequests{}
}

/*GetWorkforcemanagementManagementunitTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetWorkforcemanagementManagementunitTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}][%d] getWorkforcemanagementManagementunitTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitInternalServerError creates a GetWorkforcemanagementManagementunitInternalServerError with default headers values
func NewGetWorkforcemanagementManagementunitInternalServerError() *GetWorkforcemanagementManagementunitInternalServerError {
	return &GetWorkforcemanagementManagementunitInternalServerError{}
}

/*GetWorkforcemanagementManagementunitInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementManagementunitInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}][%d] getWorkforcemanagementManagementunitInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitServiceUnavailable creates a GetWorkforcemanagementManagementunitServiceUnavailable with default headers values
func NewGetWorkforcemanagementManagementunitServiceUnavailable() *GetWorkforcemanagementManagementunitServiceUnavailable {
	return &GetWorkforcemanagementManagementunitServiceUnavailable{}
}

/*GetWorkforcemanagementManagementunitServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementManagementunitServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}][%d] getWorkforcemanagementManagementunitServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitGatewayTimeout creates a GetWorkforcemanagementManagementunitGatewayTimeout with default headers values
func NewGetWorkforcemanagementManagementunitGatewayTimeout() *GetWorkforcemanagementManagementunitGatewayTimeout {
	return &GetWorkforcemanagementManagementunitGatewayTimeout{}
}

/*GetWorkforcemanagementManagementunitGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementManagementunitGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}][%d] getWorkforcemanagementManagementunitGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}