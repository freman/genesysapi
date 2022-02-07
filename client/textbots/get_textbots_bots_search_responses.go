// Code generated by go-swagger; DO NOT EDIT.

package textbots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetTextbotsBotsSearchReader is a Reader for the GetTextbotsBotsSearch structure.
type GetTextbotsBotsSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTextbotsBotsSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTextbotsBotsSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTextbotsBotsSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTextbotsBotsSearchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTextbotsBotsSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTextbotsBotsSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTextbotsBotsSearchRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTextbotsBotsSearchRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTextbotsBotsSearchUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTextbotsBotsSearchTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTextbotsBotsSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTextbotsBotsSearchServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTextbotsBotsSearchGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTextbotsBotsSearchOK creates a GetTextbotsBotsSearchOK with default headers values
func NewGetTextbotsBotsSearchOK() *GetTextbotsBotsSearchOK {
	return &GetTextbotsBotsSearchOK{}
}

/*GetTextbotsBotsSearchOK handles this case with default header values.

successful operation
*/
type GetTextbotsBotsSearchOK struct {
	Payload *models.BotSearchResponseEntityListing
}

func (o *GetTextbotsBotsSearchOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchOK  %+v", 200, o.Payload)
}

func (o *GetTextbotsBotsSearchOK) GetPayload() *models.BotSearchResponseEntityListing {
	return o.Payload
}

func (o *GetTextbotsBotsSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BotSearchResponseEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTextbotsBotsSearchBadRequest creates a GetTextbotsBotsSearchBadRequest with default headers values
func NewGetTextbotsBotsSearchBadRequest() *GetTextbotsBotsSearchBadRequest {
	return &GetTextbotsBotsSearchBadRequest{}
}

/*GetTextbotsBotsSearchBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTextbotsBotsSearchBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTextbotsBotsSearchBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchBadRequest  %+v", 400, o.Payload)
}

func (o *GetTextbotsBotsSearchBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTextbotsBotsSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTextbotsBotsSearchUnauthorized creates a GetTextbotsBotsSearchUnauthorized with default headers values
func NewGetTextbotsBotsSearchUnauthorized() *GetTextbotsBotsSearchUnauthorized {
	return &GetTextbotsBotsSearchUnauthorized{}
}

/*GetTextbotsBotsSearchUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTextbotsBotsSearchUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTextbotsBotsSearchUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTextbotsBotsSearchUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTextbotsBotsSearchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTextbotsBotsSearchForbidden creates a GetTextbotsBotsSearchForbidden with default headers values
func NewGetTextbotsBotsSearchForbidden() *GetTextbotsBotsSearchForbidden {
	return &GetTextbotsBotsSearchForbidden{}
}

/*GetTextbotsBotsSearchForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTextbotsBotsSearchForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTextbotsBotsSearchForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchForbidden  %+v", 403, o.Payload)
}

func (o *GetTextbotsBotsSearchForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTextbotsBotsSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTextbotsBotsSearchNotFound creates a GetTextbotsBotsSearchNotFound with default headers values
func NewGetTextbotsBotsSearchNotFound() *GetTextbotsBotsSearchNotFound {
	return &GetTextbotsBotsSearchNotFound{}
}

/*GetTextbotsBotsSearchNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTextbotsBotsSearchNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTextbotsBotsSearchNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchNotFound  %+v", 404, o.Payload)
}

func (o *GetTextbotsBotsSearchNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTextbotsBotsSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTextbotsBotsSearchRequestTimeout creates a GetTextbotsBotsSearchRequestTimeout with default headers values
func NewGetTextbotsBotsSearchRequestTimeout() *GetTextbotsBotsSearchRequestTimeout {
	return &GetTextbotsBotsSearchRequestTimeout{}
}

/*GetTextbotsBotsSearchRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTextbotsBotsSearchRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTextbotsBotsSearchRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTextbotsBotsSearchRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTextbotsBotsSearchRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTextbotsBotsSearchRequestEntityTooLarge creates a GetTextbotsBotsSearchRequestEntityTooLarge with default headers values
func NewGetTextbotsBotsSearchRequestEntityTooLarge() *GetTextbotsBotsSearchRequestEntityTooLarge {
	return &GetTextbotsBotsSearchRequestEntityTooLarge{}
}

/*GetTextbotsBotsSearchRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTextbotsBotsSearchRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTextbotsBotsSearchRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTextbotsBotsSearchRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTextbotsBotsSearchRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTextbotsBotsSearchUnsupportedMediaType creates a GetTextbotsBotsSearchUnsupportedMediaType with default headers values
func NewGetTextbotsBotsSearchUnsupportedMediaType() *GetTextbotsBotsSearchUnsupportedMediaType {
	return &GetTextbotsBotsSearchUnsupportedMediaType{}
}

/*GetTextbotsBotsSearchUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTextbotsBotsSearchUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTextbotsBotsSearchUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTextbotsBotsSearchUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTextbotsBotsSearchUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTextbotsBotsSearchTooManyRequests creates a GetTextbotsBotsSearchTooManyRequests with default headers values
func NewGetTextbotsBotsSearchTooManyRequests() *GetTextbotsBotsSearchTooManyRequests {
	return &GetTextbotsBotsSearchTooManyRequests{}
}

/*GetTextbotsBotsSearchTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTextbotsBotsSearchTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTextbotsBotsSearchTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTextbotsBotsSearchTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTextbotsBotsSearchTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTextbotsBotsSearchInternalServerError creates a GetTextbotsBotsSearchInternalServerError with default headers values
func NewGetTextbotsBotsSearchInternalServerError() *GetTextbotsBotsSearchInternalServerError {
	return &GetTextbotsBotsSearchInternalServerError{}
}

/*GetTextbotsBotsSearchInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTextbotsBotsSearchInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTextbotsBotsSearchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTextbotsBotsSearchInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTextbotsBotsSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTextbotsBotsSearchServiceUnavailable creates a GetTextbotsBotsSearchServiceUnavailable with default headers values
func NewGetTextbotsBotsSearchServiceUnavailable() *GetTextbotsBotsSearchServiceUnavailable {
	return &GetTextbotsBotsSearchServiceUnavailable{}
}

/*GetTextbotsBotsSearchServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTextbotsBotsSearchServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTextbotsBotsSearchServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTextbotsBotsSearchServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTextbotsBotsSearchServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTextbotsBotsSearchGatewayTimeout creates a GetTextbotsBotsSearchGatewayTimeout with default headers values
func NewGetTextbotsBotsSearchGatewayTimeout() *GetTextbotsBotsSearchGatewayTimeout {
	return &GetTextbotsBotsSearchGatewayTimeout{}
}

/*GetTextbotsBotsSearchGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTextbotsBotsSearchGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTextbotsBotsSearchGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTextbotsBotsSearchGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTextbotsBotsSearchGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}