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

// GetWorkforcemanagementManagementunitUserTimeoffrequestsReader is a Reader for the GetWorkforcemanagementManagementunitUserTimeoffrequests structure.
type GetWorkforcemanagementManagementunitUserTimeoffrequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestsOK creates a GetWorkforcemanagementManagementunitUserTimeoffrequestsOK with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestsOK() *GetWorkforcemanagementManagementunitUserTimeoffrequestsOK {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestsOK{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestsOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestsOK struct {
	Payload *models.TimeOffRequestList
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests][%d] getWorkforcemanagementManagementunitUserTimeoffrequestsOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsOK) GetPayload() *models.TimeOffRequestList {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TimeOffRequestList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestsBadRequest creates a GetWorkforcemanagementManagementunitUserTimeoffrequestsBadRequest with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestsBadRequest() *GetWorkforcemanagementManagementunitUserTimeoffrequestsBadRequest {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestsBadRequest{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests][%d] getWorkforcemanagementManagementunitUserTimeoffrequestsBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestsUnauthorized creates a GetWorkforcemanagementManagementunitUserTimeoffrequestsUnauthorized with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestsUnauthorized() *GetWorkforcemanagementManagementunitUserTimeoffrequestsUnauthorized {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestsUnauthorized{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests][%d] getWorkforcemanagementManagementunitUserTimeoffrequestsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestsForbidden creates a GetWorkforcemanagementManagementunitUserTimeoffrequestsForbidden with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestsForbidden() *GetWorkforcemanagementManagementunitUserTimeoffrequestsForbidden {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestsForbidden{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests][%d] getWorkforcemanagementManagementunitUserTimeoffrequestsForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestsNotFound creates a GetWorkforcemanagementManagementunitUserTimeoffrequestsNotFound with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestsNotFound() *GetWorkforcemanagementManagementunitUserTimeoffrequestsNotFound {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestsNotFound{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests][%d] getWorkforcemanagementManagementunitUserTimeoffrequestsNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestsRequestTimeout creates a GetWorkforcemanagementManagementunitUserTimeoffrequestsRequestTimeout with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestsRequestTimeout() *GetWorkforcemanagementManagementunitUserTimeoffrequestsRequestTimeout {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestsRequestTimeout{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests][%d] getWorkforcemanagementManagementunitUserTimeoffrequestsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestsRequestEntityTooLarge creates a GetWorkforcemanagementManagementunitUserTimeoffrequestsRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestsRequestEntityTooLarge() *GetWorkforcemanagementManagementunitUserTimeoffrequestsRequestEntityTooLarge {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestsRequestEntityTooLarge{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests][%d] getWorkforcemanagementManagementunitUserTimeoffrequestsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestsUnsupportedMediaType creates a GetWorkforcemanagementManagementunitUserTimeoffrequestsUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestsUnsupportedMediaType() *GetWorkforcemanagementManagementunitUserTimeoffrequestsUnsupportedMediaType {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestsUnsupportedMediaType{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests][%d] getWorkforcemanagementManagementunitUserTimeoffrequestsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestsTooManyRequests creates a GetWorkforcemanagementManagementunitUserTimeoffrequestsTooManyRequests with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestsTooManyRequests() *GetWorkforcemanagementManagementunitUserTimeoffrequestsTooManyRequests {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestsTooManyRequests{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests][%d] getWorkforcemanagementManagementunitUserTimeoffrequestsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestsInternalServerError creates a GetWorkforcemanagementManagementunitUserTimeoffrequestsInternalServerError with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestsInternalServerError() *GetWorkforcemanagementManagementunitUserTimeoffrequestsInternalServerError {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestsInternalServerError{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests][%d] getWorkforcemanagementManagementunitUserTimeoffrequestsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestsServiceUnavailable creates a GetWorkforcemanagementManagementunitUserTimeoffrequestsServiceUnavailable with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestsServiceUnavailable() *GetWorkforcemanagementManagementunitUserTimeoffrequestsServiceUnavailable {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestsServiceUnavailable{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests][%d] getWorkforcemanagementManagementunitUserTimeoffrequestsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestsGatewayTimeout creates a GetWorkforcemanagementManagementunitUserTimeoffrequestsGatewayTimeout with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestsGatewayTimeout() *GetWorkforcemanagementManagementunitUserTimeoffrequestsGatewayTimeout {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestsGatewayTimeout{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests][%d] getWorkforcemanagementManagementunitUserTimeoffrequestsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
