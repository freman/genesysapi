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

// GetWorkforcemanagementBusinessunitsReader is a Reader for the GetWorkforcemanagementBusinessunits structure.
type GetWorkforcemanagementBusinessunitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementBusinessunitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementBusinessunitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementBusinessunitsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementBusinessunitsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementBusinessunitsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementBusinessunitsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementBusinessunitsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementBusinessunitsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementBusinessunitsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementBusinessunitsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementBusinessunitsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementBusinessunitsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementBusinessunitsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementBusinessunitsOK creates a GetWorkforcemanagementBusinessunitsOK with default headers values
func NewGetWorkforcemanagementBusinessunitsOK() *GetWorkforcemanagementBusinessunitsOK {
	return &GetWorkforcemanagementBusinessunitsOK{}
}

/*GetWorkforcemanagementBusinessunitsOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementBusinessunitsOK struct {
	Payload *models.BusinessUnitListing
}

func (o *GetWorkforcemanagementBusinessunitsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits][%d] getWorkforcemanagementBusinessunitsOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsOK) GetPayload() *models.BusinessUnitListing {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BusinessUnitListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitsBadRequest creates a GetWorkforcemanagementBusinessunitsBadRequest with default headers values
func NewGetWorkforcemanagementBusinessunitsBadRequest() *GetWorkforcemanagementBusinessunitsBadRequest {
	return &GetWorkforcemanagementBusinessunitsBadRequest{}
}

/*GetWorkforcemanagementBusinessunitsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementBusinessunitsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits][%d] getWorkforcemanagementBusinessunitsBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitsUnauthorized creates a GetWorkforcemanagementBusinessunitsUnauthorized with default headers values
func NewGetWorkforcemanagementBusinessunitsUnauthorized() *GetWorkforcemanagementBusinessunitsUnauthorized {
	return &GetWorkforcemanagementBusinessunitsUnauthorized{}
}

/*GetWorkforcemanagementBusinessunitsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementBusinessunitsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits][%d] getWorkforcemanagementBusinessunitsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitsForbidden creates a GetWorkforcemanagementBusinessunitsForbidden with default headers values
func NewGetWorkforcemanagementBusinessunitsForbidden() *GetWorkforcemanagementBusinessunitsForbidden {
	return &GetWorkforcemanagementBusinessunitsForbidden{}
}

/*GetWorkforcemanagementBusinessunitsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementBusinessunitsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits][%d] getWorkforcemanagementBusinessunitsForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitsNotFound creates a GetWorkforcemanagementBusinessunitsNotFound with default headers values
func NewGetWorkforcemanagementBusinessunitsNotFound() *GetWorkforcemanagementBusinessunitsNotFound {
	return &GetWorkforcemanagementBusinessunitsNotFound{}
}

/*GetWorkforcemanagementBusinessunitsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementBusinessunitsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits][%d] getWorkforcemanagementBusinessunitsNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitsRequestTimeout creates a GetWorkforcemanagementBusinessunitsRequestTimeout with default headers values
func NewGetWorkforcemanagementBusinessunitsRequestTimeout() *GetWorkforcemanagementBusinessunitsRequestTimeout {
	return &GetWorkforcemanagementBusinessunitsRequestTimeout{}
}

/*GetWorkforcemanagementBusinessunitsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementBusinessunitsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits][%d] getWorkforcemanagementBusinessunitsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitsRequestEntityTooLarge creates a GetWorkforcemanagementBusinessunitsRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementBusinessunitsRequestEntityTooLarge() *GetWorkforcemanagementBusinessunitsRequestEntityTooLarge {
	return &GetWorkforcemanagementBusinessunitsRequestEntityTooLarge{}
}

/*GetWorkforcemanagementBusinessunitsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetWorkforcemanagementBusinessunitsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits][%d] getWorkforcemanagementBusinessunitsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitsUnsupportedMediaType creates a GetWorkforcemanagementBusinessunitsUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementBusinessunitsUnsupportedMediaType() *GetWorkforcemanagementBusinessunitsUnsupportedMediaType {
	return &GetWorkforcemanagementBusinessunitsUnsupportedMediaType{}
}

/*GetWorkforcemanagementBusinessunitsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementBusinessunitsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits][%d] getWorkforcemanagementBusinessunitsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitsTooManyRequests creates a GetWorkforcemanagementBusinessunitsTooManyRequests with default headers values
func NewGetWorkforcemanagementBusinessunitsTooManyRequests() *GetWorkforcemanagementBusinessunitsTooManyRequests {
	return &GetWorkforcemanagementBusinessunitsTooManyRequests{}
}

/*GetWorkforcemanagementBusinessunitsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementBusinessunitsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits][%d] getWorkforcemanagementBusinessunitsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitsInternalServerError creates a GetWorkforcemanagementBusinessunitsInternalServerError with default headers values
func NewGetWorkforcemanagementBusinessunitsInternalServerError() *GetWorkforcemanagementBusinessunitsInternalServerError {
	return &GetWorkforcemanagementBusinessunitsInternalServerError{}
}

/*GetWorkforcemanagementBusinessunitsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementBusinessunitsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits][%d] getWorkforcemanagementBusinessunitsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitsServiceUnavailable creates a GetWorkforcemanagementBusinessunitsServiceUnavailable with default headers values
func NewGetWorkforcemanagementBusinessunitsServiceUnavailable() *GetWorkforcemanagementBusinessunitsServiceUnavailable {
	return &GetWorkforcemanagementBusinessunitsServiceUnavailable{}
}

/*GetWorkforcemanagementBusinessunitsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementBusinessunitsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits][%d] getWorkforcemanagementBusinessunitsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitsGatewayTimeout creates a GetWorkforcemanagementBusinessunitsGatewayTimeout with default headers values
func NewGetWorkforcemanagementBusinessunitsGatewayTimeout() *GetWorkforcemanagementBusinessunitsGatewayTimeout {
	return &GetWorkforcemanagementBusinessunitsGatewayTimeout{}
}

/*GetWorkforcemanagementBusinessunitsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementBusinessunitsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits][%d] getWorkforcemanagementBusinessunitsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
