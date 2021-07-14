// Code generated by go-swagger; DO NOT EDIT.

package locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetLocationSublocationsReader is a Reader for the GetLocationSublocations structure.
type GetLocationSublocationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLocationSublocationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLocationSublocationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLocationSublocationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLocationSublocationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLocationSublocationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLocationSublocationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetLocationSublocationsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLocationSublocationsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLocationSublocationsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLocationSublocationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLocationSublocationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLocationSublocationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLocationSublocationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLocationSublocationsOK creates a GetLocationSublocationsOK with default headers values
func NewGetLocationSublocationsOK() *GetLocationSublocationsOK {
	return &GetLocationSublocationsOK{}
}

/*GetLocationSublocationsOK handles this case with default header values.

successful operation
*/
type GetLocationSublocationsOK struct {
	Payload *models.LocationEntityListing
}

func (o *GetLocationSublocationsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}/sublocations][%d] getLocationSublocationsOK  %+v", 200, o.Payload)
}

func (o *GetLocationSublocationsOK) GetPayload() *models.LocationEntityListing {
	return o.Payload
}

func (o *GetLocationSublocationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocationEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationSublocationsBadRequest creates a GetLocationSublocationsBadRequest with default headers values
func NewGetLocationSublocationsBadRequest() *GetLocationSublocationsBadRequest {
	return &GetLocationSublocationsBadRequest{}
}

/*GetLocationSublocationsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLocationSublocationsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetLocationSublocationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}/sublocations][%d] getLocationSublocationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetLocationSublocationsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationSublocationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationSublocationsUnauthorized creates a GetLocationSublocationsUnauthorized with default headers values
func NewGetLocationSublocationsUnauthorized() *GetLocationSublocationsUnauthorized {
	return &GetLocationSublocationsUnauthorized{}
}

/*GetLocationSublocationsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLocationSublocationsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetLocationSublocationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}/sublocations][%d] getLocationSublocationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLocationSublocationsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationSublocationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationSublocationsForbidden creates a GetLocationSublocationsForbidden with default headers values
func NewGetLocationSublocationsForbidden() *GetLocationSublocationsForbidden {
	return &GetLocationSublocationsForbidden{}
}

/*GetLocationSublocationsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetLocationSublocationsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetLocationSublocationsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}/sublocations][%d] getLocationSublocationsForbidden  %+v", 403, o.Payload)
}

func (o *GetLocationSublocationsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationSublocationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationSublocationsNotFound creates a GetLocationSublocationsNotFound with default headers values
func NewGetLocationSublocationsNotFound() *GetLocationSublocationsNotFound {
	return &GetLocationSublocationsNotFound{}
}

/*GetLocationSublocationsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetLocationSublocationsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetLocationSublocationsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}/sublocations][%d] getLocationSublocationsNotFound  %+v", 404, o.Payload)
}

func (o *GetLocationSublocationsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationSublocationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationSublocationsRequestTimeout creates a GetLocationSublocationsRequestTimeout with default headers values
func NewGetLocationSublocationsRequestTimeout() *GetLocationSublocationsRequestTimeout {
	return &GetLocationSublocationsRequestTimeout{}
}

/*GetLocationSublocationsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetLocationSublocationsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetLocationSublocationsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}/sublocations][%d] getLocationSublocationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLocationSublocationsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationSublocationsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationSublocationsRequestEntityTooLarge creates a GetLocationSublocationsRequestEntityTooLarge with default headers values
func NewGetLocationSublocationsRequestEntityTooLarge() *GetLocationSublocationsRequestEntityTooLarge {
	return &GetLocationSublocationsRequestEntityTooLarge{}
}

/*GetLocationSublocationsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetLocationSublocationsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetLocationSublocationsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}/sublocations][%d] getLocationSublocationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLocationSublocationsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationSublocationsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationSublocationsUnsupportedMediaType creates a GetLocationSublocationsUnsupportedMediaType with default headers values
func NewGetLocationSublocationsUnsupportedMediaType() *GetLocationSublocationsUnsupportedMediaType {
	return &GetLocationSublocationsUnsupportedMediaType{}
}

/*GetLocationSublocationsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLocationSublocationsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetLocationSublocationsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}/sublocations][%d] getLocationSublocationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLocationSublocationsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationSublocationsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationSublocationsTooManyRequests creates a GetLocationSublocationsTooManyRequests with default headers values
func NewGetLocationSublocationsTooManyRequests() *GetLocationSublocationsTooManyRequests {
	return &GetLocationSublocationsTooManyRequests{}
}

/*GetLocationSublocationsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetLocationSublocationsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetLocationSublocationsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}/sublocations][%d] getLocationSublocationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLocationSublocationsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationSublocationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationSublocationsInternalServerError creates a GetLocationSublocationsInternalServerError with default headers values
func NewGetLocationSublocationsInternalServerError() *GetLocationSublocationsInternalServerError {
	return &GetLocationSublocationsInternalServerError{}
}

/*GetLocationSublocationsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLocationSublocationsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetLocationSublocationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}/sublocations][%d] getLocationSublocationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLocationSublocationsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationSublocationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationSublocationsServiceUnavailable creates a GetLocationSublocationsServiceUnavailable with default headers values
func NewGetLocationSublocationsServiceUnavailable() *GetLocationSublocationsServiceUnavailable {
	return &GetLocationSublocationsServiceUnavailable{}
}

/*GetLocationSublocationsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLocationSublocationsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetLocationSublocationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}/sublocations][%d] getLocationSublocationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLocationSublocationsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationSublocationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationSublocationsGatewayTimeout creates a GetLocationSublocationsGatewayTimeout with default headers values
func NewGetLocationSublocationsGatewayTimeout() *GetLocationSublocationsGatewayTimeout {
	return &GetLocationSublocationsGatewayTimeout{}
}

/*GetLocationSublocationsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetLocationSublocationsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetLocationSublocationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}/sublocations][%d] getLocationSublocationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLocationSublocationsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationSublocationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
