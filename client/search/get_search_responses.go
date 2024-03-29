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
	case 408:
		result := NewGetSearchRequestTimeout()
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

/*
GetSearchOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSearchOK struct {
	Payload *models.JSONNodeSearchResponse
}

// IsSuccess returns true when this get search o k response has a 2xx status code
func (o *GetSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get search o k response has a 3xx status code
func (o *GetSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search o k response has a 4xx status code
func (o *GetSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get search o k response has a 5xx status code
func (o *GetSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get search o k response a status code equal to that given
func (o *GetSearchOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSearchOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchOK  %+v", 200, o.Payload)
}

func (o *GetSearchOK) String() string {
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

/*
GetSearchBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetSearchBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get search bad request response has a 2xx status code
func (o *GetSearchBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search bad request response has a 3xx status code
func (o *GetSearchBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search bad request response has a 4xx status code
func (o *GetSearchBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search bad request response has a 5xx status code
func (o *GetSearchBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get search bad request response a status code equal to that given
func (o *GetSearchBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetSearchBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchBadRequest  %+v", 400, o.Payload)
}

func (o *GetSearchBadRequest) String() string {
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

/*
GetSearchUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetSearchUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get search unauthorized response has a 2xx status code
func (o *GetSearchUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search unauthorized response has a 3xx status code
func (o *GetSearchUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search unauthorized response has a 4xx status code
func (o *GetSearchUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search unauthorized response has a 5xx status code
func (o *GetSearchUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get search unauthorized response a status code equal to that given
func (o *GetSearchUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetSearchUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSearchUnauthorized) String() string {
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

/*
GetSearchForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetSearchForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get search forbidden response has a 2xx status code
func (o *GetSearchForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search forbidden response has a 3xx status code
func (o *GetSearchForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search forbidden response has a 4xx status code
func (o *GetSearchForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search forbidden response has a 5xx status code
func (o *GetSearchForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get search forbidden response a status code equal to that given
func (o *GetSearchForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetSearchForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchForbidden  %+v", 403, o.Payload)
}

func (o *GetSearchForbidden) String() string {
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

/*
GetSearchNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetSearchNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get search not found response has a 2xx status code
func (o *GetSearchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search not found response has a 3xx status code
func (o *GetSearchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search not found response has a 4xx status code
func (o *GetSearchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search not found response has a 5xx status code
func (o *GetSearchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get search not found response a status code equal to that given
func (o *GetSearchNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetSearchNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchNotFound  %+v", 404, o.Payload)
}

func (o *GetSearchNotFound) String() string {
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

// NewGetSearchRequestTimeout creates a GetSearchRequestTimeout with default headers values
func NewGetSearchRequestTimeout() *GetSearchRequestTimeout {
	return &GetSearchRequestTimeout{}
}

/*
GetSearchRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetSearchRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get search request timeout response has a 2xx status code
func (o *GetSearchRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search request timeout response has a 3xx status code
func (o *GetSearchRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search request timeout response has a 4xx status code
func (o *GetSearchRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search request timeout response has a 5xx status code
func (o *GetSearchRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get search request timeout response a status code equal to that given
func (o *GetSearchRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetSearchRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetSearchRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetSearchRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSearchRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
GetSearchRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetSearchRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get search request entity too large response has a 2xx status code
func (o *GetSearchRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search request entity too large response has a 3xx status code
func (o *GetSearchRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search request entity too large response has a 4xx status code
func (o *GetSearchRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search request entity too large response has a 5xx status code
func (o *GetSearchRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get search request entity too large response a status code equal to that given
func (o *GetSearchRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetSearchRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetSearchRequestEntityTooLarge) String() string {
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

/*
GetSearchUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetSearchUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get search unsupported media type response has a 2xx status code
func (o *GetSearchUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search unsupported media type response has a 3xx status code
func (o *GetSearchUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search unsupported media type response has a 4xx status code
func (o *GetSearchUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search unsupported media type response has a 5xx status code
func (o *GetSearchUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get search unsupported media type response a status code equal to that given
func (o *GetSearchUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetSearchUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetSearchUnsupportedMediaType) String() string {
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

/*
GetSearchTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetSearchTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get search too many requests response has a 2xx status code
func (o *GetSearchTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search too many requests response has a 3xx status code
func (o *GetSearchTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search too many requests response has a 4xx status code
func (o *GetSearchTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search too many requests response has a 5xx status code
func (o *GetSearchTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get search too many requests response a status code equal to that given
func (o *GetSearchTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetSearchTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSearchTooManyRequests) String() string {
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

/*
GetSearchInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetSearchInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get search internal server error response has a 2xx status code
func (o *GetSearchInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search internal server error response has a 3xx status code
func (o *GetSearchInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search internal server error response has a 4xx status code
func (o *GetSearchInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get search internal server error response has a 5xx status code
func (o *GetSearchInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get search internal server error response a status code equal to that given
func (o *GetSearchInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetSearchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSearchInternalServerError) String() string {
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

/*
GetSearchServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetSearchServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get search service unavailable response has a 2xx status code
func (o *GetSearchServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search service unavailable response has a 3xx status code
func (o *GetSearchServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search service unavailable response has a 4xx status code
func (o *GetSearchServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get search service unavailable response has a 5xx status code
func (o *GetSearchServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get search service unavailable response a status code equal to that given
func (o *GetSearchServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetSearchServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSearchServiceUnavailable) String() string {
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

/*
GetSearchGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetSearchGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get search gateway timeout response has a 2xx status code
func (o *GetSearchGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search gateway timeout response has a 3xx status code
func (o *GetSearchGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search gateway timeout response has a 4xx status code
func (o *GetSearchGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get search gateway timeout response has a 5xx status code
func (o *GetSearchGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get search gateway timeout response a status code equal to that given
func (o *GetSearchGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetSearchGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/search][%d] getSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSearchGatewayTimeout) String() string {
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
