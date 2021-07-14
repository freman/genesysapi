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

// GetOrganizationsLimitsChangerequestsReader is a Reader for the GetOrganizationsLimitsChangerequests structure.
type GetOrganizationsLimitsChangerequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationsLimitsChangerequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationsLimitsChangerequestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrganizationsLimitsChangerequestsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrganizationsLimitsChangerequestsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrganizationsLimitsChangerequestsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrganizationsLimitsChangerequestsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOrganizationsLimitsChangerequestsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOrganizationsLimitsChangerequestsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOrganizationsLimitsChangerequestsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrganizationsLimitsChangerequestsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrganizationsLimitsChangerequestsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOrganizationsLimitsChangerequestsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOrganizationsLimitsChangerequestsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationsLimitsChangerequestsOK creates a GetOrganizationsLimitsChangerequestsOK with default headers values
func NewGetOrganizationsLimitsChangerequestsOK() *GetOrganizationsLimitsChangerequestsOK {
	return &GetOrganizationsLimitsChangerequestsOK{}
}

/*GetOrganizationsLimitsChangerequestsOK handles this case with default header values.

successful operation
*/
type GetOrganizationsLimitsChangerequestsOK struct {
	Payload *models.LimitChangeRequestsEntityListing
}

func (o *GetOrganizationsLimitsChangerequestsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests][%d] getOrganizationsLimitsChangerequestsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestsOK) GetPayload() *models.LimitChangeRequestsEntityListing {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LimitChangeRequestsEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsChangerequestsBadRequest creates a GetOrganizationsLimitsChangerequestsBadRequest with default headers values
func NewGetOrganizationsLimitsChangerequestsBadRequest() *GetOrganizationsLimitsChangerequestsBadRequest {
	return &GetOrganizationsLimitsChangerequestsBadRequest{}
}

/*GetOrganizationsLimitsChangerequestsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOrganizationsLimitsChangerequestsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsLimitsChangerequestsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests][%d] getOrganizationsLimitsChangerequestsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsChangerequestsUnauthorized creates a GetOrganizationsLimitsChangerequestsUnauthorized with default headers values
func NewGetOrganizationsLimitsChangerequestsUnauthorized() *GetOrganizationsLimitsChangerequestsUnauthorized {
	return &GetOrganizationsLimitsChangerequestsUnauthorized{}
}

/*GetOrganizationsLimitsChangerequestsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOrganizationsLimitsChangerequestsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsLimitsChangerequestsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests][%d] getOrganizationsLimitsChangerequestsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsChangerequestsForbidden creates a GetOrganizationsLimitsChangerequestsForbidden with default headers values
func NewGetOrganizationsLimitsChangerequestsForbidden() *GetOrganizationsLimitsChangerequestsForbidden {
	return &GetOrganizationsLimitsChangerequestsForbidden{}
}

/*GetOrganizationsLimitsChangerequestsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOrganizationsLimitsChangerequestsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsLimitsChangerequestsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests][%d] getOrganizationsLimitsChangerequestsForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsChangerequestsNotFound creates a GetOrganizationsLimitsChangerequestsNotFound with default headers values
func NewGetOrganizationsLimitsChangerequestsNotFound() *GetOrganizationsLimitsChangerequestsNotFound {
	return &GetOrganizationsLimitsChangerequestsNotFound{}
}

/*GetOrganizationsLimitsChangerequestsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOrganizationsLimitsChangerequestsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsLimitsChangerequestsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests][%d] getOrganizationsLimitsChangerequestsNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsChangerequestsRequestTimeout creates a GetOrganizationsLimitsChangerequestsRequestTimeout with default headers values
func NewGetOrganizationsLimitsChangerequestsRequestTimeout() *GetOrganizationsLimitsChangerequestsRequestTimeout {
	return &GetOrganizationsLimitsChangerequestsRequestTimeout{}
}

/*GetOrganizationsLimitsChangerequestsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOrganizationsLimitsChangerequestsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsLimitsChangerequestsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests][%d] getOrganizationsLimitsChangerequestsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsChangerequestsRequestEntityTooLarge creates a GetOrganizationsLimitsChangerequestsRequestEntityTooLarge with default headers values
func NewGetOrganizationsLimitsChangerequestsRequestEntityTooLarge() *GetOrganizationsLimitsChangerequestsRequestEntityTooLarge {
	return &GetOrganizationsLimitsChangerequestsRequestEntityTooLarge{}
}

/*GetOrganizationsLimitsChangerequestsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOrganizationsLimitsChangerequestsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsLimitsChangerequestsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests][%d] getOrganizationsLimitsChangerequestsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsChangerequestsUnsupportedMediaType creates a GetOrganizationsLimitsChangerequestsUnsupportedMediaType with default headers values
func NewGetOrganizationsLimitsChangerequestsUnsupportedMediaType() *GetOrganizationsLimitsChangerequestsUnsupportedMediaType {
	return &GetOrganizationsLimitsChangerequestsUnsupportedMediaType{}
}

/*GetOrganizationsLimitsChangerequestsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOrganizationsLimitsChangerequestsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsLimitsChangerequestsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests][%d] getOrganizationsLimitsChangerequestsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsChangerequestsTooManyRequests creates a GetOrganizationsLimitsChangerequestsTooManyRequests with default headers values
func NewGetOrganizationsLimitsChangerequestsTooManyRequests() *GetOrganizationsLimitsChangerequestsTooManyRequests {
	return &GetOrganizationsLimitsChangerequestsTooManyRequests{}
}

/*GetOrganizationsLimitsChangerequestsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOrganizationsLimitsChangerequestsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsLimitsChangerequestsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests][%d] getOrganizationsLimitsChangerequestsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsChangerequestsInternalServerError creates a GetOrganizationsLimitsChangerequestsInternalServerError with default headers values
func NewGetOrganizationsLimitsChangerequestsInternalServerError() *GetOrganizationsLimitsChangerequestsInternalServerError {
	return &GetOrganizationsLimitsChangerequestsInternalServerError{}
}

/*GetOrganizationsLimitsChangerequestsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOrganizationsLimitsChangerequestsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsLimitsChangerequestsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests][%d] getOrganizationsLimitsChangerequestsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsChangerequestsServiceUnavailable creates a GetOrganizationsLimitsChangerequestsServiceUnavailable with default headers values
func NewGetOrganizationsLimitsChangerequestsServiceUnavailable() *GetOrganizationsLimitsChangerequestsServiceUnavailable {
	return &GetOrganizationsLimitsChangerequestsServiceUnavailable{}
}

/*GetOrganizationsLimitsChangerequestsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOrganizationsLimitsChangerequestsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsLimitsChangerequestsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests][%d] getOrganizationsLimitsChangerequestsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsLimitsChangerequestsGatewayTimeout creates a GetOrganizationsLimitsChangerequestsGatewayTimeout with default headers values
func NewGetOrganizationsLimitsChangerequestsGatewayTimeout() *GetOrganizationsLimitsChangerequestsGatewayTimeout {
	return &GetOrganizationsLimitsChangerequestsGatewayTimeout{}
}

/*GetOrganizationsLimitsChangerequestsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOrganizationsLimitsChangerequestsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOrganizationsLimitsChangerequestsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/organizations/limits/changerequests][%d] getOrganizationsLimitsChangerequestsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOrganizationsLimitsChangerequestsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrganizationsLimitsChangerequestsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}