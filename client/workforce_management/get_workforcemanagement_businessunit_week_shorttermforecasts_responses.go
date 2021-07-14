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

// GetWorkforcemanagementBusinessunitWeekShorttermforecastsReader is a Reader for the GetWorkforcemanagementBusinessunitWeekShorttermforecasts structure.
type GetWorkforcemanagementBusinessunitWeekShorttermforecastsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsOK creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastsOK with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsOK() *GetWorkforcemanagementBusinessunitWeekShorttermforecastsOK {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastsOK{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastsOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastsOK struct {
	Payload *models.BuShortTermForecastListing
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastsOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsOK) GetPayload() *models.BuShortTermForecastListing {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuShortTermForecastListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsBadRequest creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastsBadRequest with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsBadRequest() *GetWorkforcemanagementBusinessunitWeekShorttermforecastsBadRequest {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastsBadRequest{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastsBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsUnauthorized creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastsUnauthorized with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsUnauthorized() *GetWorkforcemanagementBusinessunitWeekShorttermforecastsUnauthorized {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastsUnauthorized{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsForbidden creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastsForbidden with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsForbidden() *GetWorkforcemanagementBusinessunitWeekShorttermforecastsForbidden {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastsForbidden{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastsForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsNotFound creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastsNotFound with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsNotFound() *GetWorkforcemanagementBusinessunitWeekShorttermforecastsNotFound {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastsNotFound{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastsNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsRequestTimeout creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastsRequestTimeout with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsRequestTimeout() *GetWorkforcemanagementBusinessunitWeekShorttermforecastsRequestTimeout {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastsRequestTimeout{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsRequestEntityTooLarge creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastsRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsRequestEntityTooLarge() *GetWorkforcemanagementBusinessunitWeekShorttermforecastsRequestEntityTooLarge {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastsRequestEntityTooLarge{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsUnsupportedMediaType creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastsUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsUnsupportedMediaType() *GetWorkforcemanagementBusinessunitWeekShorttermforecastsUnsupportedMediaType {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastsUnsupportedMediaType{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsTooManyRequests creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastsTooManyRequests with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsTooManyRequests() *GetWorkforcemanagementBusinessunitWeekShorttermforecastsTooManyRequests {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastsTooManyRequests{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsInternalServerError creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastsInternalServerError with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsInternalServerError() *GetWorkforcemanagementBusinessunitWeekShorttermforecastsInternalServerError {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastsInternalServerError{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsServiceUnavailable creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastsServiceUnavailable with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsServiceUnavailable() *GetWorkforcemanagementBusinessunitWeekShorttermforecastsServiceUnavailable {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastsServiceUnavailable{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsGatewayTimeout creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastsGatewayTimeout with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastsGatewayTimeout() *GetWorkforcemanagementBusinessunitWeekShorttermforecastsGatewayTimeout {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastsGatewayTimeout{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
