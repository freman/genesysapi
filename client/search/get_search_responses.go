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

// GetSearchReader is a Reader for the GetSearch structure.
type GetSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSearchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetSearchRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetSearchUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSearchTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetSearchServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetSearchGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSearchOK creates a GetSearchOK with default headers values
func NewGetSearchOK() *GetSearchOK {
	return &GetSearchOK{}
}

/*GetSearchOK handles this case with default header values.

successful operation
*/
type GetSearchOK struct {
	Payload *models.JSONNodeSearchResponse
}

func (o *GetSearchOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchOK  %+v", 200, o.Payload)
}

func (o *GetSearchOK) GetPayload() *models.JSONNodeSearchResponse {
	return o.Payload
}

func (o *GetSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONNodeSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchBadRequest creates a GetSearchBadRequest with default headers values
func NewGetSearchBadRequest() *GetSearchBadRequest {
	return &GetSearchBadRequest{}
}

/*GetSearchBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetSearchBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetSearchBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchBadRequest  %+v", 400, o.Payload)
}

func (o *GetSearchBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchUnauthorized creates a GetSearchUnauthorized with default headers values
func NewGetSearchUnauthorized() *GetSearchUnauthorized {
	return &GetSearchUnauthorized{}
}

/*GetSearchUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetSearchUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetSearchUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSearchUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSearchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchForbidden creates a GetSearchForbidden with default headers values
func NewGetSearchForbidden() *GetSearchForbidden {
	return &GetSearchForbidden{}
}

/*GetSearchForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetSearchForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetSearchForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchForbidden  %+v", 403, o.Payload)
}

func (o *GetSearchForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchNotFound creates a GetSearchNotFound with default headers values
func NewGetSearchNotFound() *GetSearchNotFound {
	return &GetSearchNotFound{}
}

/*GetSearchNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetSearchNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetSearchNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchNotFound  %+v", 404, o.Payload)
}

func (o *GetSearchNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchRequestEntityTooLarge creates a GetSearchRequestEntityTooLarge with default headers values
func NewGetSearchRequestEntityTooLarge() *GetSearchRequestEntityTooLarge {
	return &GetSearchRequestEntityTooLarge{}
}

/*GetSearchRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetSearchRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetSearchRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetSearchRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSearchRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchUnsupportedMediaType creates a GetSearchUnsupportedMediaType with default headers values
func NewGetSearchUnsupportedMediaType() *GetSearchUnsupportedMediaType {
	return &GetSearchUnsupportedMediaType{}
}

/*GetSearchUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetSearchUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetSearchUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetSearchUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSearchUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchTooManyRequests creates a GetSearchTooManyRequests with default headers values
func NewGetSearchTooManyRequests() *GetSearchTooManyRequests {
	return &GetSearchTooManyRequests{}
}

/*GetSearchTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetSearchTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetSearchTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSearchTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSearchTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchInternalServerError creates a GetSearchInternalServerError with default headers values
func NewGetSearchInternalServerError() *GetSearchInternalServerError {
	return &GetSearchInternalServerError{}
}

/*GetSearchInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetSearchInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetSearchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSearchInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchServiceUnavailable creates a GetSearchServiceUnavailable with default headers values
func NewGetSearchServiceUnavailable() *GetSearchServiceUnavailable {
	return &GetSearchServiceUnavailable{}
}

/*GetSearchServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetSearchServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetSearchServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSearchServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSearchServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchGatewayTimeout creates a GetSearchGatewayTimeout with default headers values
func NewGetSearchGatewayTimeout() *GetSearchGatewayTimeout {
	return &GetSearchGatewayTimeout{}
}

/*GetSearchGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetSearchGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetSearchGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSearchGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSearchGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
