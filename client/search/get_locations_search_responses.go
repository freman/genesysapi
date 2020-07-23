// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetLocationsSearchReader is a Reader for the GetLocationsSearch structure.
type GetLocationsSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLocationsSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLocationsSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLocationsSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLocationsSearchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLocationsSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLocationsSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLocationsSearchRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLocationsSearchUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLocationsSearchTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLocationsSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLocationsSearchServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLocationsSearchGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLocationsSearchOK creates a GetLocationsSearchOK with default headers values
func NewGetLocationsSearchOK() *GetLocationsSearchOK {
	return &GetLocationsSearchOK{}
}

/*GetLocationsSearchOK handles this case with default header values.

successful operation
*/
type GetLocationsSearchOK struct {
	Payload *models.LocationsSearchResponse
}

func (o *GetLocationsSearchOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/search][%d] getLocationsSearchOK  %+v", 200, o.Payload)
}

func (o *GetLocationsSearchOK) GetPayload() *models.LocationsSearchResponse {
	return o.Payload
}

func (o *GetLocationsSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocationsSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationsSearchBadRequest creates a GetLocationsSearchBadRequest with default headers values
func NewGetLocationsSearchBadRequest() *GetLocationsSearchBadRequest {
	return &GetLocationsSearchBadRequest{}
}

/*GetLocationsSearchBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLocationsSearchBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetLocationsSearchBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/search][%d] getLocationsSearchBadRequest  %+v", 400, o.Payload)
}

func (o *GetLocationsSearchBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationsSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationsSearchUnauthorized creates a GetLocationsSearchUnauthorized with default headers values
func NewGetLocationsSearchUnauthorized() *GetLocationsSearchUnauthorized {
	return &GetLocationsSearchUnauthorized{}
}

/*GetLocationsSearchUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLocationsSearchUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetLocationsSearchUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/search][%d] getLocationsSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLocationsSearchUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationsSearchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationsSearchForbidden creates a GetLocationsSearchForbidden with default headers values
func NewGetLocationsSearchForbidden() *GetLocationsSearchForbidden {
	return &GetLocationsSearchForbidden{}
}

/*GetLocationsSearchForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetLocationsSearchForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetLocationsSearchForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/search][%d] getLocationsSearchForbidden  %+v", 403, o.Payload)
}

func (o *GetLocationsSearchForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationsSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationsSearchNotFound creates a GetLocationsSearchNotFound with default headers values
func NewGetLocationsSearchNotFound() *GetLocationsSearchNotFound {
	return &GetLocationsSearchNotFound{}
}

/*GetLocationsSearchNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetLocationsSearchNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetLocationsSearchNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/search][%d] getLocationsSearchNotFound  %+v", 404, o.Payload)
}

func (o *GetLocationsSearchNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationsSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationsSearchRequestEntityTooLarge creates a GetLocationsSearchRequestEntityTooLarge with default headers values
func NewGetLocationsSearchRequestEntityTooLarge() *GetLocationsSearchRequestEntityTooLarge {
	return &GetLocationsSearchRequestEntityTooLarge{}
}

/*GetLocationsSearchRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetLocationsSearchRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetLocationsSearchRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/search][%d] getLocationsSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLocationsSearchRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationsSearchRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationsSearchUnsupportedMediaType creates a GetLocationsSearchUnsupportedMediaType with default headers values
func NewGetLocationsSearchUnsupportedMediaType() *GetLocationsSearchUnsupportedMediaType {
	return &GetLocationsSearchUnsupportedMediaType{}
}

/*GetLocationsSearchUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLocationsSearchUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetLocationsSearchUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/search][%d] getLocationsSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLocationsSearchUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationsSearchUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationsSearchTooManyRequests creates a GetLocationsSearchTooManyRequests with default headers values
func NewGetLocationsSearchTooManyRequests() *GetLocationsSearchTooManyRequests {
	return &GetLocationsSearchTooManyRequests{}
}

/*GetLocationsSearchTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetLocationsSearchTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetLocationsSearchTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/search][%d] getLocationsSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLocationsSearchTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationsSearchTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationsSearchInternalServerError creates a GetLocationsSearchInternalServerError with default headers values
func NewGetLocationsSearchInternalServerError() *GetLocationsSearchInternalServerError {
	return &GetLocationsSearchInternalServerError{}
}

/*GetLocationsSearchInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLocationsSearchInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetLocationsSearchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/search][%d] getLocationsSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLocationsSearchInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationsSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationsSearchServiceUnavailable creates a GetLocationsSearchServiceUnavailable with default headers values
func NewGetLocationsSearchServiceUnavailable() *GetLocationsSearchServiceUnavailable {
	return &GetLocationsSearchServiceUnavailable{}
}

/*GetLocationsSearchServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLocationsSearchServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetLocationsSearchServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/search][%d] getLocationsSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLocationsSearchServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationsSearchServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationsSearchGatewayTimeout creates a GetLocationsSearchGatewayTimeout with default headers values
func NewGetLocationsSearchGatewayTimeout() *GetLocationsSearchGatewayTimeout {
	return &GetLocationsSearchGatewayTimeout{}
}

/*GetLocationsSearchGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetLocationsSearchGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetLocationsSearchGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/search][%d] getLocationsSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLocationsSearchGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationsSearchGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
