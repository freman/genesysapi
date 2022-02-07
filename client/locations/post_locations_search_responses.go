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

// PostLocationsSearchReader is a Reader for the PostLocationsSearch structure.
type PostLocationsSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLocationsSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLocationsSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLocationsSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostLocationsSearchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLocationsSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostLocationsSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostLocationsSearchRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostLocationsSearchRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostLocationsSearchUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostLocationsSearchTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostLocationsSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostLocationsSearchServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostLocationsSearchGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLocationsSearchOK creates a PostLocationsSearchOK with default headers values
func NewPostLocationsSearchOK() *PostLocationsSearchOK {
	return &PostLocationsSearchOK{}
}

/*PostLocationsSearchOK handles this case with default header values.

successful operation
*/
type PostLocationsSearchOK struct {
	Payload *models.LocationsSearchResponse
}

func (o *PostLocationsSearchOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations/search][%d] postLocationsSearchOK  %+v", 200, o.Payload)
}

func (o *PostLocationsSearchOK) GetPayload() *models.LocationsSearchResponse {
	return o.Payload
}

func (o *PostLocationsSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocationsSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLocationsSearchBadRequest creates a PostLocationsSearchBadRequest with default headers values
func NewPostLocationsSearchBadRequest() *PostLocationsSearchBadRequest {
	return &PostLocationsSearchBadRequest{}
}

/*PostLocationsSearchBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostLocationsSearchBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostLocationsSearchBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations/search][%d] postLocationsSearchBadRequest  %+v", 400, o.Payload)
}

func (o *PostLocationsSearchBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLocationsSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLocationsSearchUnauthorized creates a PostLocationsSearchUnauthorized with default headers values
func NewPostLocationsSearchUnauthorized() *PostLocationsSearchUnauthorized {
	return &PostLocationsSearchUnauthorized{}
}

/*PostLocationsSearchUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostLocationsSearchUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostLocationsSearchUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations/search][%d] postLocationsSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLocationsSearchUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLocationsSearchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLocationsSearchForbidden creates a PostLocationsSearchForbidden with default headers values
func NewPostLocationsSearchForbidden() *PostLocationsSearchForbidden {
	return &PostLocationsSearchForbidden{}
}

/*PostLocationsSearchForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostLocationsSearchForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostLocationsSearchForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations/search][%d] postLocationsSearchForbidden  %+v", 403, o.Payload)
}

func (o *PostLocationsSearchForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLocationsSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLocationsSearchNotFound creates a PostLocationsSearchNotFound with default headers values
func NewPostLocationsSearchNotFound() *PostLocationsSearchNotFound {
	return &PostLocationsSearchNotFound{}
}

/*PostLocationsSearchNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostLocationsSearchNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostLocationsSearchNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations/search][%d] postLocationsSearchNotFound  %+v", 404, o.Payload)
}

func (o *PostLocationsSearchNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLocationsSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLocationsSearchRequestTimeout creates a PostLocationsSearchRequestTimeout with default headers values
func NewPostLocationsSearchRequestTimeout() *PostLocationsSearchRequestTimeout {
	return &PostLocationsSearchRequestTimeout{}
}

/*PostLocationsSearchRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostLocationsSearchRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostLocationsSearchRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations/search][%d] postLocationsSearchRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostLocationsSearchRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLocationsSearchRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLocationsSearchRequestEntityTooLarge creates a PostLocationsSearchRequestEntityTooLarge with default headers values
func NewPostLocationsSearchRequestEntityTooLarge() *PostLocationsSearchRequestEntityTooLarge {
	return &PostLocationsSearchRequestEntityTooLarge{}
}

/*PostLocationsSearchRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostLocationsSearchRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostLocationsSearchRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations/search][%d] postLocationsSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLocationsSearchRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLocationsSearchRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLocationsSearchUnsupportedMediaType creates a PostLocationsSearchUnsupportedMediaType with default headers values
func NewPostLocationsSearchUnsupportedMediaType() *PostLocationsSearchUnsupportedMediaType {
	return &PostLocationsSearchUnsupportedMediaType{}
}

/*PostLocationsSearchUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostLocationsSearchUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostLocationsSearchUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations/search][%d] postLocationsSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLocationsSearchUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLocationsSearchUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLocationsSearchTooManyRequests creates a PostLocationsSearchTooManyRequests with default headers values
func NewPostLocationsSearchTooManyRequests() *PostLocationsSearchTooManyRequests {
	return &PostLocationsSearchTooManyRequests{}
}

/*PostLocationsSearchTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostLocationsSearchTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostLocationsSearchTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations/search][%d] postLocationsSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLocationsSearchTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLocationsSearchTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLocationsSearchInternalServerError creates a PostLocationsSearchInternalServerError with default headers values
func NewPostLocationsSearchInternalServerError() *PostLocationsSearchInternalServerError {
	return &PostLocationsSearchInternalServerError{}
}

/*PostLocationsSearchInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostLocationsSearchInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostLocationsSearchInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations/search][%d] postLocationsSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLocationsSearchInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLocationsSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLocationsSearchServiceUnavailable creates a PostLocationsSearchServiceUnavailable with default headers values
func NewPostLocationsSearchServiceUnavailable() *PostLocationsSearchServiceUnavailable {
	return &PostLocationsSearchServiceUnavailable{}
}

/*PostLocationsSearchServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostLocationsSearchServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostLocationsSearchServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations/search][%d] postLocationsSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLocationsSearchServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLocationsSearchServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLocationsSearchGatewayTimeout creates a PostLocationsSearchGatewayTimeout with default headers values
func NewPostLocationsSearchGatewayTimeout() *PostLocationsSearchGatewayTimeout {
	return &PostLocationsSearchGatewayTimeout{}
}

/*PostLocationsSearchGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostLocationsSearchGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostLocationsSearchGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/locations/search][%d] postLocationsSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLocationsSearchGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLocationsSearchGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}