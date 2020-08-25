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

// GetWorkforcemanagementManagementunitActivitycodesReader is a Reader for the GetWorkforcemanagementManagementunitActivitycodes structure.
type GetWorkforcemanagementManagementunitActivitycodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementManagementunitActivitycodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementManagementunitActivitycodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementManagementunitActivitycodesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementManagementunitActivitycodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementManagementunitActivitycodesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementManagementunitActivitycodesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementManagementunitActivitycodesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementManagementunitActivitycodesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementManagementunitActivitycodesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementManagementunitActivitycodesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementManagementunitActivitycodesOK creates a GetWorkforcemanagementManagementunitActivitycodesOK with default headers values
func NewGetWorkforcemanagementManagementunitActivitycodesOK() *GetWorkforcemanagementManagementunitActivitycodesOK {
	return &GetWorkforcemanagementManagementunitActivitycodesOK{}
}

/*GetWorkforcemanagementManagementunitActivitycodesOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementManagementunitActivitycodesOK struct {
	Payload *models.ActivityCodeContainer
}

func (o *GetWorkforcemanagementManagementunitActivitycodesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes][%d] getWorkforcemanagementManagementunitActivitycodesOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitActivitycodesOK) GetPayload() *models.ActivityCodeContainer {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitActivitycodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActivityCodeContainer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitActivitycodesBadRequest creates a GetWorkforcemanagementManagementunitActivitycodesBadRequest with default headers values
func NewGetWorkforcemanagementManagementunitActivitycodesBadRequest() *GetWorkforcemanagementManagementunitActivitycodesBadRequest {
	return &GetWorkforcemanagementManagementunitActivitycodesBadRequest{}
}

/*GetWorkforcemanagementManagementunitActivitycodesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementManagementunitActivitycodesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitActivitycodesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes][%d] getWorkforcemanagementManagementunitActivitycodesBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitActivitycodesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitActivitycodesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitActivitycodesUnauthorized creates a GetWorkforcemanagementManagementunitActivitycodesUnauthorized with default headers values
func NewGetWorkforcemanagementManagementunitActivitycodesUnauthorized() *GetWorkforcemanagementManagementunitActivitycodesUnauthorized {
	return &GetWorkforcemanagementManagementunitActivitycodesUnauthorized{}
}

/*GetWorkforcemanagementManagementunitActivitycodesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementManagementunitActivitycodesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitActivitycodesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes][%d] getWorkforcemanagementManagementunitActivitycodesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitActivitycodesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitActivitycodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitActivitycodesForbidden creates a GetWorkforcemanagementManagementunitActivitycodesForbidden with default headers values
func NewGetWorkforcemanagementManagementunitActivitycodesForbidden() *GetWorkforcemanagementManagementunitActivitycodesForbidden {
	return &GetWorkforcemanagementManagementunitActivitycodesForbidden{}
}

/*GetWorkforcemanagementManagementunitActivitycodesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementManagementunitActivitycodesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitActivitycodesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes][%d] getWorkforcemanagementManagementunitActivitycodesForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitActivitycodesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitActivitycodesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitActivitycodesNotFound creates a GetWorkforcemanagementManagementunitActivitycodesNotFound with default headers values
func NewGetWorkforcemanagementManagementunitActivitycodesNotFound() *GetWorkforcemanagementManagementunitActivitycodesNotFound {
	return &GetWorkforcemanagementManagementunitActivitycodesNotFound{}
}

/*GetWorkforcemanagementManagementunitActivitycodesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementManagementunitActivitycodesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitActivitycodesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes][%d] getWorkforcemanagementManagementunitActivitycodesNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitActivitycodesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitActivitycodesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge creates a GetWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge() *GetWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge {
	return &GetWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge{}
}

/*GetWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes][%d] getWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitActivitycodesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType creates a GetWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType() *GetWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType {
	return &GetWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType{}
}

/*GetWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes][%d] getWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitActivitycodesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitActivitycodesTooManyRequests creates a GetWorkforcemanagementManagementunitActivitycodesTooManyRequests with default headers values
func NewGetWorkforcemanagementManagementunitActivitycodesTooManyRequests() *GetWorkforcemanagementManagementunitActivitycodesTooManyRequests {
	return &GetWorkforcemanagementManagementunitActivitycodesTooManyRequests{}
}

/*GetWorkforcemanagementManagementunitActivitycodesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetWorkforcemanagementManagementunitActivitycodesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitActivitycodesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes][%d] getWorkforcemanagementManagementunitActivitycodesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitActivitycodesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitActivitycodesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitActivitycodesInternalServerError creates a GetWorkforcemanagementManagementunitActivitycodesInternalServerError with default headers values
func NewGetWorkforcemanagementManagementunitActivitycodesInternalServerError() *GetWorkforcemanagementManagementunitActivitycodesInternalServerError {
	return &GetWorkforcemanagementManagementunitActivitycodesInternalServerError{}
}

/*GetWorkforcemanagementManagementunitActivitycodesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementManagementunitActivitycodesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitActivitycodesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes][%d] getWorkforcemanagementManagementunitActivitycodesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitActivitycodesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitActivitycodesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitActivitycodesServiceUnavailable creates a GetWorkforcemanagementManagementunitActivitycodesServiceUnavailable with default headers values
func NewGetWorkforcemanagementManagementunitActivitycodesServiceUnavailable() *GetWorkforcemanagementManagementunitActivitycodesServiceUnavailable {
	return &GetWorkforcemanagementManagementunitActivitycodesServiceUnavailable{}
}

/*GetWorkforcemanagementManagementunitActivitycodesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementManagementunitActivitycodesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitActivitycodesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes][%d] getWorkforcemanagementManagementunitActivitycodesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitActivitycodesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitActivitycodesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitActivitycodesGatewayTimeout creates a GetWorkforcemanagementManagementunitActivitycodesGatewayTimeout with default headers values
func NewGetWorkforcemanagementManagementunitActivitycodesGatewayTimeout() *GetWorkforcemanagementManagementunitActivitycodesGatewayTimeout {
	return &GetWorkforcemanagementManagementunitActivitycodesGatewayTimeout{}
}

/*GetWorkforcemanagementManagementunitActivitycodesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementManagementunitActivitycodesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitActivitycodesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes][%d] getWorkforcemanagementManagementunitActivitycodesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitActivitycodesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitActivitycodesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}