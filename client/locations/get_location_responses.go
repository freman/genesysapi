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

// GetLocationReader is a Reader for the GetLocation structure.
type GetLocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLocationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLocationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLocationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLocationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLocationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLocationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLocationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLocationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLocationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLocationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLocationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLocationOK creates a GetLocationOK with default headers values
func NewGetLocationOK() *GetLocationOK {
	return &GetLocationOK{}
}

/*GetLocationOK handles this case with default header values.

successful operation
*/
type GetLocationOK struct {
	Payload *models.LocationDefinition
}

func (o *GetLocationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationOK  %+v", 200, o.Payload)
}

func (o *GetLocationOK) GetPayload() *models.LocationDefinition {
	return o.Payload
}

func (o *GetLocationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocationDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationBadRequest creates a GetLocationBadRequest with default headers values
func NewGetLocationBadRequest() *GetLocationBadRequest {
	return &GetLocationBadRequest{}
}

/*GetLocationBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLocationBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetLocationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationBadRequest  %+v", 400, o.Payload)
}

func (o *GetLocationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationUnauthorized creates a GetLocationUnauthorized with default headers values
func NewGetLocationUnauthorized() *GetLocationUnauthorized {
	return &GetLocationUnauthorized{}
}

/*GetLocationUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLocationUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetLocationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLocationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationForbidden creates a GetLocationForbidden with default headers values
func NewGetLocationForbidden() *GetLocationForbidden {
	return &GetLocationForbidden{}
}

/*GetLocationForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetLocationForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetLocationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationForbidden  %+v", 403, o.Payload)
}

func (o *GetLocationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationNotFound creates a GetLocationNotFound with default headers values
func NewGetLocationNotFound() *GetLocationNotFound {
	return &GetLocationNotFound{}
}

/*GetLocationNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetLocationNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetLocationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationNotFound  %+v", 404, o.Payload)
}

func (o *GetLocationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationRequestEntityTooLarge creates a GetLocationRequestEntityTooLarge with default headers values
func NewGetLocationRequestEntityTooLarge() *GetLocationRequestEntityTooLarge {
	return &GetLocationRequestEntityTooLarge{}
}

/*GetLocationRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetLocationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetLocationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLocationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationUnsupportedMediaType creates a GetLocationUnsupportedMediaType with default headers values
func NewGetLocationUnsupportedMediaType() *GetLocationUnsupportedMediaType {
	return &GetLocationUnsupportedMediaType{}
}

/*GetLocationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLocationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetLocationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLocationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationTooManyRequests creates a GetLocationTooManyRequests with default headers values
func NewGetLocationTooManyRequests() *GetLocationTooManyRequests {
	return &GetLocationTooManyRequests{}
}

/*GetLocationTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetLocationTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetLocationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLocationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationInternalServerError creates a GetLocationInternalServerError with default headers values
func NewGetLocationInternalServerError() *GetLocationInternalServerError {
	return &GetLocationInternalServerError{}
}

/*GetLocationInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLocationInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetLocationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLocationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationServiceUnavailable creates a GetLocationServiceUnavailable with default headers values
func NewGetLocationServiceUnavailable() *GetLocationServiceUnavailable {
	return &GetLocationServiceUnavailable{}
}

/*GetLocationServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLocationServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetLocationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLocationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationGatewayTimeout creates a GetLocationGatewayTimeout with default headers values
func NewGetLocationGatewayTimeout() *GetLocationGatewayTimeout {
	return &GetLocationGatewayTimeout{}
}

/*GetLocationGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetLocationGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetLocationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLocationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
