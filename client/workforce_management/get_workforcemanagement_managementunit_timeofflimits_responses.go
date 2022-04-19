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

// GetWorkforcemanagementManagementunitTimeofflimitsReader is a Reader for the GetWorkforcemanagementManagementunitTimeofflimits structure.
type GetWorkforcemanagementManagementunitTimeofflimitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementManagementunitTimeofflimitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsOK creates a GetWorkforcemanagementManagementunitTimeofflimitsOK with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsOK() *GetWorkforcemanagementManagementunitTimeofflimitsOK {
	return &GetWorkforcemanagementManagementunitTimeofflimitsOK{}
}

/*GetWorkforcemanagementManagementunitTimeofflimitsOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementManagementunitTimeofflimitsOK struct {
	Payload *models.TimeOffLimitListing
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsOK) GetPayload() *models.TimeOffLimitListing {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TimeOffLimitListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsBadRequest creates a GetWorkforcemanagementManagementunitTimeofflimitsBadRequest with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsBadRequest() *GetWorkforcemanagementManagementunitTimeofflimitsBadRequest {
	return &GetWorkforcemanagementManagementunitTimeofflimitsBadRequest{}
}

/*GetWorkforcemanagementManagementunitTimeofflimitsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementManagementunitTimeofflimitsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsUnauthorized creates a GetWorkforcemanagementManagementunitTimeofflimitsUnauthorized with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsUnauthorized() *GetWorkforcemanagementManagementunitTimeofflimitsUnauthorized {
	return &GetWorkforcemanagementManagementunitTimeofflimitsUnauthorized{}
}

/*GetWorkforcemanagementManagementunitTimeofflimitsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementManagementunitTimeofflimitsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsForbidden creates a GetWorkforcemanagementManagementunitTimeofflimitsForbidden with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsForbidden() *GetWorkforcemanagementManagementunitTimeofflimitsForbidden {
	return &GetWorkforcemanagementManagementunitTimeofflimitsForbidden{}
}

/*GetWorkforcemanagementManagementunitTimeofflimitsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementManagementunitTimeofflimitsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsNotFound creates a GetWorkforcemanagementManagementunitTimeofflimitsNotFound with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsNotFound() *GetWorkforcemanagementManagementunitTimeofflimitsNotFound {
	return &GetWorkforcemanagementManagementunitTimeofflimitsNotFound{}
}

/*GetWorkforcemanagementManagementunitTimeofflimitsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementManagementunitTimeofflimitsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout creates a GetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout() *GetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout {
	return &GetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout{}
}

/*GetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge creates a GetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge() *GetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge {
	return &GetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge{}
}

/*GetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType creates a GetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType() *GetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType {
	return &GetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType{}
}

/*GetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests creates a GetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests() *GetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests {
	return &GetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests{}
}

/*GetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsInternalServerError creates a GetWorkforcemanagementManagementunitTimeofflimitsInternalServerError with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsInternalServerError() *GetWorkforcemanagementManagementunitTimeofflimitsInternalServerError {
	return &GetWorkforcemanagementManagementunitTimeofflimitsInternalServerError{}
}

/*GetWorkforcemanagementManagementunitTimeofflimitsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementManagementunitTimeofflimitsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable creates a GetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable() *GetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable {
	return &GetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable{}
}

/*GetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout creates a GetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout() *GetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout {
	return &GetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout{}
}

/*GetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
