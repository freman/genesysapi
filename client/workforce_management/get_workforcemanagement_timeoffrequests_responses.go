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

// GetWorkforcemanagementTimeoffrequestsReader is a Reader for the GetWorkforcemanagementTimeoffrequests structure.
type GetWorkforcemanagementTimeoffrequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementTimeoffrequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementTimeoffrequestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementTimeoffrequestsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementTimeoffrequestsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementTimeoffrequestsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementTimeoffrequestsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementTimeoffrequestsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementTimeoffrequestsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementTimeoffrequestsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementTimeoffrequestsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementTimeoffrequestsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementTimeoffrequestsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementTimeoffrequestsOK creates a GetWorkforcemanagementTimeoffrequestsOK with default headers values
func NewGetWorkforcemanagementTimeoffrequestsOK() *GetWorkforcemanagementTimeoffrequestsOK {
	return &GetWorkforcemanagementTimeoffrequestsOK{}
}

/*GetWorkforcemanagementTimeoffrequestsOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementTimeoffrequestsOK struct {
	Payload *models.TimeOffRequestList
}

func (o *GetWorkforcemanagementTimeoffrequestsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/timeoffrequests][%d] getWorkforcemanagementTimeoffrequestsOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementTimeoffrequestsOK) GetPayload() *models.TimeOffRequestList {
	return o.Payload
}

func (o *GetWorkforcemanagementTimeoffrequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TimeOffRequestList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementTimeoffrequestsBadRequest creates a GetWorkforcemanagementTimeoffrequestsBadRequest with default headers values
func NewGetWorkforcemanagementTimeoffrequestsBadRequest() *GetWorkforcemanagementTimeoffrequestsBadRequest {
	return &GetWorkforcemanagementTimeoffrequestsBadRequest{}
}

/*GetWorkforcemanagementTimeoffrequestsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementTimeoffrequestsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementTimeoffrequestsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/timeoffrequests][%d] getWorkforcemanagementTimeoffrequestsBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementTimeoffrequestsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementTimeoffrequestsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementTimeoffrequestsUnauthorized creates a GetWorkforcemanagementTimeoffrequestsUnauthorized with default headers values
func NewGetWorkforcemanagementTimeoffrequestsUnauthorized() *GetWorkforcemanagementTimeoffrequestsUnauthorized {
	return &GetWorkforcemanagementTimeoffrequestsUnauthorized{}
}

/*GetWorkforcemanagementTimeoffrequestsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementTimeoffrequestsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementTimeoffrequestsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/timeoffrequests][%d] getWorkforcemanagementTimeoffrequestsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementTimeoffrequestsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementTimeoffrequestsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementTimeoffrequestsForbidden creates a GetWorkforcemanagementTimeoffrequestsForbidden with default headers values
func NewGetWorkforcemanagementTimeoffrequestsForbidden() *GetWorkforcemanagementTimeoffrequestsForbidden {
	return &GetWorkforcemanagementTimeoffrequestsForbidden{}
}

/*GetWorkforcemanagementTimeoffrequestsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementTimeoffrequestsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementTimeoffrequestsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/timeoffrequests][%d] getWorkforcemanagementTimeoffrequestsForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementTimeoffrequestsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementTimeoffrequestsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementTimeoffrequestsNotFound creates a GetWorkforcemanagementTimeoffrequestsNotFound with default headers values
func NewGetWorkforcemanagementTimeoffrequestsNotFound() *GetWorkforcemanagementTimeoffrequestsNotFound {
	return &GetWorkforcemanagementTimeoffrequestsNotFound{}
}

/*GetWorkforcemanagementTimeoffrequestsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementTimeoffrequestsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementTimeoffrequestsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/timeoffrequests][%d] getWorkforcemanagementTimeoffrequestsNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementTimeoffrequestsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementTimeoffrequestsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementTimeoffrequestsRequestEntityTooLarge creates a GetWorkforcemanagementTimeoffrequestsRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementTimeoffrequestsRequestEntityTooLarge() *GetWorkforcemanagementTimeoffrequestsRequestEntityTooLarge {
	return &GetWorkforcemanagementTimeoffrequestsRequestEntityTooLarge{}
}

/*GetWorkforcemanagementTimeoffrequestsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetWorkforcemanagementTimeoffrequestsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementTimeoffrequestsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/timeoffrequests][%d] getWorkforcemanagementTimeoffrequestsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementTimeoffrequestsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementTimeoffrequestsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementTimeoffrequestsUnsupportedMediaType creates a GetWorkforcemanagementTimeoffrequestsUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementTimeoffrequestsUnsupportedMediaType() *GetWorkforcemanagementTimeoffrequestsUnsupportedMediaType {
	return &GetWorkforcemanagementTimeoffrequestsUnsupportedMediaType{}
}

/*GetWorkforcemanagementTimeoffrequestsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementTimeoffrequestsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementTimeoffrequestsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/timeoffrequests][%d] getWorkforcemanagementTimeoffrequestsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementTimeoffrequestsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementTimeoffrequestsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementTimeoffrequestsTooManyRequests creates a GetWorkforcemanagementTimeoffrequestsTooManyRequests with default headers values
func NewGetWorkforcemanagementTimeoffrequestsTooManyRequests() *GetWorkforcemanagementTimeoffrequestsTooManyRequests {
	return &GetWorkforcemanagementTimeoffrequestsTooManyRequests{}
}

/*GetWorkforcemanagementTimeoffrequestsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetWorkforcemanagementTimeoffrequestsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementTimeoffrequestsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/timeoffrequests][%d] getWorkforcemanagementTimeoffrequestsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementTimeoffrequestsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementTimeoffrequestsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementTimeoffrequestsInternalServerError creates a GetWorkforcemanagementTimeoffrequestsInternalServerError with default headers values
func NewGetWorkforcemanagementTimeoffrequestsInternalServerError() *GetWorkforcemanagementTimeoffrequestsInternalServerError {
	return &GetWorkforcemanagementTimeoffrequestsInternalServerError{}
}

/*GetWorkforcemanagementTimeoffrequestsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementTimeoffrequestsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementTimeoffrequestsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/timeoffrequests][%d] getWorkforcemanagementTimeoffrequestsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementTimeoffrequestsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementTimeoffrequestsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementTimeoffrequestsServiceUnavailable creates a GetWorkforcemanagementTimeoffrequestsServiceUnavailable with default headers values
func NewGetWorkforcemanagementTimeoffrequestsServiceUnavailable() *GetWorkforcemanagementTimeoffrequestsServiceUnavailable {
	return &GetWorkforcemanagementTimeoffrequestsServiceUnavailable{}
}

/*GetWorkforcemanagementTimeoffrequestsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementTimeoffrequestsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementTimeoffrequestsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/timeoffrequests][%d] getWorkforcemanagementTimeoffrequestsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementTimeoffrequestsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementTimeoffrequestsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementTimeoffrequestsGatewayTimeout creates a GetWorkforcemanagementTimeoffrequestsGatewayTimeout with default headers values
func NewGetWorkforcemanagementTimeoffrequestsGatewayTimeout() *GetWorkforcemanagementTimeoffrequestsGatewayTimeout {
	return &GetWorkforcemanagementTimeoffrequestsGatewayTimeout{}
}

/*GetWorkforcemanagementTimeoffrequestsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementTimeoffrequestsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementTimeoffrequestsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/timeoffrequests][%d] getWorkforcemanagementTimeoffrequestsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementTimeoffrequestsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementTimeoffrequestsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
