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

/*
GetTextbotsBotsSearchOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTextbotsBotsSearchOK struct {
	Payload *models.BotSearchResponseEntityListing
}

// IsSuccess returns true when this get textbots bots search o k response has a 2xx status code
func (o *GetTextbotsBotsSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get textbots bots search o k response has a 3xx status code
func (o *GetTextbotsBotsSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get textbots bots search o k response has a 4xx status code
func (o *GetTextbotsBotsSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get textbots bots search o k response has a 5xx status code
func (o *GetTextbotsBotsSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get textbots bots search o k response a status code equal to that given
func (o *GetTextbotsBotsSearchOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTextbotsBotsSearchOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchOK  %+v", 200, o.Payload)
}

func (o *GetTextbotsBotsSearchOK) String() string {
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

/*
GetTextbotsBotsSearchBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTextbotsBotsSearchBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get textbots bots search bad request response has a 2xx status code
func (o *GetTextbotsBotsSearchBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get textbots bots search bad request response has a 3xx status code
func (o *GetTextbotsBotsSearchBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get textbots bots search bad request response has a 4xx status code
func (o *GetTextbotsBotsSearchBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get textbots bots search bad request response has a 5xx status code
func (o *GetTextbotsBotsSearchBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get textbots bots search bad request response a status code equal to that given
func (o *GetTextbotsBotsSearchBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetTextbotsBotsSearchBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchBadRequest  %+v", 400, o.Payload)
}

func (o *GetTextbotsBotsSearchBadRequest) String() string {
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

/*
GetTextbotsBotsSearchUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTextbotsBotsSearchUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get textbots bots search unauthorized response has a 2xx status code
func (o *GetTextbotsBotsSearchUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get textbots bots search unauthorized response has a 3xx status code
func (o *GetTextbotsBotsSearchUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get textbots bots search unauthorized response has a 4xx status code
func (o *GetTextbotsBotsSearchUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get textbots bots search unauthorized response has a 5xx status code
func (o *GetTextbotsBotsSearchUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get textbots bots search unauthorized response a status code equal to that given
func (o *GetTextbotsBotsSearchUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetTextbotsBotsSearchUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTextbotsBotsSearchUnauthorized) String() string {
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

/*
GetTextbotsBotsSearchForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetTextbotsBotsSearchForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get textbots bots search forbidden response has a 2xx status code
func (o *GetTextbotsBotsSearchForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get textbots bots search forbidden response has a 3xx status code
func (o *GetTextbotsBotsSearchForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get textbots bots search forbidden response has a 4xx status code
func (o *GetTextbotsBotsSearchForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get textbots bots search forbidden response has a 5xx status code
func (o *GetTextbotsBotsSearchForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get textbots bots search forbidden response a status code equal to that given
func (o *GetTextbotsBotsSearchForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetTextbotsBotsSearchForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchForbidden  %+v", 403, o.Payload)
}

func (o *GetTextbotsBotsSearchForbidden) String() string {
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

/*
GetTextbotsBotsSearchNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetTextbotsBotsSearchNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get textbots bots search not found response has a 2xx status code
func (o *GetTextbotsBotsSearchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get textbots bots search not found response has a 3xx status code
func (o *GetTextbotsBotsSearchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get textbots bots search not found response has a 4xx status code
func (o *GetTextbotsBotsSearchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get textbots bots search not found response has a 5xx status code
func (o *GetTextbotsBotsSearchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get textbots bots search not found response a status code equal to that given
func (o *GetTextbotsBotsSearchNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetTextbotsBotsSearchNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchNotFound  %+v", 404, o.Payload)
}

func (o *GetTextbotsBotsSearchNotFound) String() string {
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

/*
GetTextbotsBotsSearchRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTextbotsBotsSearchRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get textbots bots search request timeout response has a 2xx status code
func (o *GetTextbotsBotsSearchRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get textbots bots search request timeout response has a 3xx status code
func (o *GetTextbotsBotsSearchRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get textbots bots search request timeout response has a 4xx status code
func (o *GetTextbotsBotsSearchRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get textbots bots search request timeout response has a 5xx status code
func (o *GetTextbotsBotsSearchRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get textbots bots search request timeout response a status code equal to that given
func (o *GetTextbotsBotsSearchRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetTextbotsBotsSearchRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTextbotsBotsSearchRequestTimeout) String() string {
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

/*
GetTextbotsBotsSearchRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetTextbotsBotsSearchRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get textbots bots search request entity too large response has a 2xx status code
func (o *GetTextbotsBotsSearchRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get textbots bots search request entity too large response has a 3xx status code
func (o *GetTextbotsBotsSearchRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get textbots bots search request entity too large response has a 4xx status code
func (o *GetTextbotsBotsSearchRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get textbots bots search request entity too large response has a 5xx status code
func (o *GetTextbotsBotsSearchRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get textbots bots search request entity too large response a status code equal to that given
func (o *GetTextbotsBotsSearchRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetTextbotsBotsSearchRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTextbotsBotsSearchRequestEntityTooLarge) String() string {
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

/*
GetTextbotsBotsSearchUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTextbotsBotsSearchUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get textbots bots search unsupported media type response has a 2xx status code
func (o *GetTextbotsBotsSearchUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get textbots bots search unsupported media type response has a 3xx status code
func (o *GetTextbotsBotsSearchUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get textbots bots search unsupported media type response has a 4xx status code
func (o *GetTextbotsBotsSearchUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get textbots bots search unsupported media type response has a 5xx status code
func (o *GetTextbotsBotsSearchUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get textbots bots search unsupported media type response a status code equal to that given
func (o *GetTextbotsBotsSearchUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetTextbotsBotsSearchUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTextbotsBotsSearchUnsupportedMediaType) String() string {
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

/*
GetTextbotsBotsSearchTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTextbotsBotsSearchTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get textbots bots search too many requests response has a 2xx status code
func (o *GetTextbotsBotsSearchTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get textbots bots search too many requests response has a 3xx status code
func (o *GetTextbotsBotsSearchTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get textbots bots search too many requests response has a 4xx status code
func (o *GetTextbotsBotsSearchTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get textbots bots search too many requests response has a 5xx status code
func (o *GetTextbotsBotsSearchTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get textbots bots search too many requests response a status code equal to that given
func (o *GetTextbotsBotsSearchTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetTextbotsBotsSearchTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTextbotsBotsSearchTooManyRequests) String() string {
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

/*
GetTextbotsBotsSearchInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTextbotsBotsSearchInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get textbots bots search internal server error response has a 2xx status code
func (o *GetTextbotsBotsSearchInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get textbots bots search internal server error response has a 3xx status code
func (o *GetTextbotsBotsSearchInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get textbots bots search internal server error response has a 4xx status code
func (o *GetTextbotsBotsSearchInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get textbots bots search internal server error response has a 5xx status code
func (o *GetTextbotsBotsSearchInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get textbots bots search internal server error response a status code equal to that given
func (o *GetTextbotsBotsSearchInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTextbotsBotsSearchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTextbotsBotsSearchInternalServerError) String() string {
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

/*
GetTextbotsBotsSearchServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTextbotsBotsSearchServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get textbots bots search service unavailable response has a 2xx status code
func (o *GetTextbotsBotsSearchServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get textbots bots search service unavailable response has a 3xx status code
func (o *GetTextbotsBotsSearchServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get textbots bots search service unavailable response has a 4xx status code
func (o *GetTextbotsBotsSearchServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get textbots bots search service unavailable response has a 5xx status code
func (o *GetTextbotsBotsSearchServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get textbots bots search service unavailable response a status code equal to that given
func (o *GetTextbotsBotsSearchServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetTextbotsBotsSearchServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTextbotsBotsSearchServiceUnavailable) String() string {
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

/*
GetTextbotsBotsSearchGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetTextbotsBotsSearchGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get textbots bots search gateway timeout response has a 2xx status code
func (o *GetTextbotsBotsSearchGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get textbots bots search gateway timeout response has a 3xx status code
func (o *GetTextbotsBotsSearchGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get textbots bots search gateway timeout response has a 4xx status code
func (o *GetTextbotsBotsSearchGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get textbots bots search gateway timeout response has a 5xx status code
func (o *GetTextbotsBotsSearchGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get textbots bots search gateway timeout response a status code equal to that given
func (o *GetTextbotsBotsSearchGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetTextbotsBotsSearchGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/textbots/bots/search][%d] getTextbotsBotsSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTextbotsBotsSearchGatewayTimeout) String() string {
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
