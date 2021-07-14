// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetOrganizationsLimitsChangerequestReader is a Reader for the GetOrganizationsLimitsChangerequest structure.
type GetOrganizationsLimitsChangerequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationsLimitsChangerequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationsLimitsChangerequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrganizationsLimitsChangerequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrganizationsLimitsChangerequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrganizationsLimitsChangerequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrganizationsLimitsChangerequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOrganizationsLimitsChangerequestRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOrganizationsLimitsChangerequestRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOrganizationsLimitsChangerequestUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrganizationsLimitsChangerequestTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrganizationsLimitsChangerequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOrganizationsLimitsChangerequestServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOrganizationsLimitsChangerequestGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationsLimitsChangerequestOK creates a GetOrganizationsLimitsChangerequestOK with default headers values
func NewGetOrganizationsLimitsChangerequestOK() *GetOrganizationsLimitsChangerequestOK {
	return &GetOrganizationsLimitsChangerequestOK{}
}

/*GetOrganizationsLimitsChangerequestOK handles this case with default header values.

successful operation
*/
type GetOrganizationsLimitsChangerequestOK struct {
	Payload *models.LimitChangeRequestDetails
}

func (o *GetOrganizationsLimitsChangerequestOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests/{requestId}][%d] getOrganizationsLimitsChangerequestOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestOK) GetPayload() *models.LimitChangeRequestDetails {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LimitChangeRequestDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsChangerequestBadRequest creates a GetOrganizationsLimitsChangerequestBadRequest with default headers values
func NewGetOrganizationsLimitsChangerequestBadRequest() *GetOrganizationsLimitsChangerequestBadRequest {
	return &GetOrganizationsLimitsChangerequestBadRequest{}
}

/*GetOrganizationsLimitsChangerequestBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOrganizationsLimitsChangerequestBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsLimitsChangerequestBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests/{requestId}][%d] getOrganizationsLimitsChangerequestBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsChangerequestUnauthorized creates a GetOrganizationsLimitsChangerequestUnauthorized with default headers values
func NewGetOrganizationsLimitsChangerequestUnauthorized() *GetOrganizationsLimitsChangerequestUnauthorized {
	return &GetOrganizationsLimitsChangerequestUnauthorized{}
}

/*GetOrganizationsLimitsChangerequestUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOrganizationsLimitsChangerequestUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsLimitsChangerequestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests/{requestId}][%d] getOrganizationsLimitsChangerequestUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsChangerequestForbidden creates a GetOrganizationsLimitsChangerequestForbidden with default headers values
func NewGetOrganizationsLimitsChangerequestForbidden() *GetOrganizationsLimitsChangerequestForbidden {
	return &GetOrganizationsLimitsChangerequestForbidden{}
}

/*GetOrganizationsLimitsChangerequestForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOrganizationsLimitsChangerequestForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsLimitsChangerequestForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests/{requestId}][%d] getOrganizationsLimitsChangerequestForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsChangerequestNotFound creates a GetOrganizationsLimitsChangerequestNotFound with default headers values
func NewGetOrganizationsLimitsChangerequestNotFound() *GetOrganizationsLimitsChangerequestNotFound {
	return &GetOrganizationsLimitsChangerequestNotFound{}
}

/*GetOrganizationsLimitsChangerequestNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOrganizationsLimitsChangerequestNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsLimitsChangerequestNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests/{requestId}][%d] getOrganizationsLimitsChangerequestNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsChangerequestRequestTimeout creates a GetOrganizationsLimitsChangerequestRequestTimeout with default headers values
func NewGetOrganizationsLimitsChangerequestRequestTimeout() *GetOrganizationsLimitsChangerequestRequestTimeout {
	return &GetOrganizationsLimitsChangerequestRequestTimeout{}
}

/*GetOrganizationsLimitsChangerequestRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOrganizationsLimitsChangerequestRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsLimitsChangerequestRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests/{requestId}][%d] getOrganizationsLimitsChangerequestRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsChangerequestRequestEntityTooLarge creates a GetOrganizationsLimitsChangerequestRequestEntityTooLarge with default headers values
func NewGetOrganizationsLimitsChangerequestRequestEntityTooLarge() *GetOrganizationsLimitsChangerequestRequestEntityTooLarge {
	return &GetOrganizationsLimitsChangerequestRequestEntityTooLarge{}
}

/*GetOrganizationsLimitsChangerequestRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOrganizationsLimitsChangerequestRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsLimitsChangerequestRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests/{requestId}][%d] getOrganizationsLimitsChangerequestRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsChangerequestUnsupportedMediaType creates a GetOrganizationsLimitsChangerequestUnsupportedMediaType with default headers values
func NewGetOrganizationsLimitsChangerequestUnsupportedMediaType() *GetOrganizationsLimitsChangerequestUnsupportedMediaType {
	return &GetOrganizationsLimitsChangerequestUnsupportedMediaType{}
}

/*GetOrganizationsLimitsChangerequestUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOrganizationsLimitsChangerequestUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsLimitsChangerequestUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests/{requestId}][%d] getOrganizationsLimitsChangerequestUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsChangerequestTooManyRequests creates a GetOrganizationsLimitsChangerequestTooManyRequests with default headers values
func NewGetOrganizationsLimitsChangerequestTooManyRequests() *GetOrganizationsLimitsChangerequestTooManyRequests {
	return &GetOrganizationsLimitsChangerequestTooManyRequests{}
}

/*GetOrganizationsLimitsChangerequestTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOrganizationsLimitsChangerequestTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsLimitsChangerequestTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests/{requestId}][%d] getOrganizationsLimitsChangerequestTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsChangerequestInternalServerError creates a GetOrganizationsLimitsChangerequestInternalServerError with default headers values
func NewGetOrganizationsLimitsChangerequestInternalServerError() *GetOrganizationsLimitsChangerequestInternalServerError {
	return &GetOrganizationsLimitsChangerequestInternalServerError{}
}

/*GetOrganizationsLimitsChangerequestInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOrganizationsLimitsChangerequestInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsLimitsChangerequestInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests/{requestId}][%d] getOrganizationsLimitsChangerequestInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsChangerequestServiceUnavailable creates a GetOrganizationsLimitsChangerequestServiceUnavailable with default headers values
func NewGetOrganizationsLimitsChangerequestServiceUnavailable() *GetOrganizationsLimitsChangerequestServiceUnavailable {
	return &GetOrganizationsLimitsChangerequestServiceUnavailable{}
}

/*GetOrganizationsLimitsChangerequestServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOrganizationsLimitsChangerequestServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsLimitsChangerequestServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests/{requestId}][%d] getOrganizationsLimitsChangerequestServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsChangerequestGatewayTimeout creates a GetOrganizationsLimitsChangerequestGatewayTimeout with default headers values
func NewGetOrganizationsLimitsChangerequestGatewayTimeout() *GetOrganizationsLimitsChangerequestGatewayTimeout {
	return &GetOrganizationsLimitsChangerequestGatewayTimeout{}
}

/*GetOrganizationsLimitsChangerequestGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOrganizationsLimitsChangerequestGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsLimitsChangerequestGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests/{requestId}][%d] getOrganizationsLimitsChangerequestGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}