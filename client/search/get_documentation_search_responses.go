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

// GetDocumentationSearchReader is a Reader for the GetDocumentationSearch structure.
type GetDocumentationSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDocumentationSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDocumentationSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDocumentationSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDocumentationSearchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDocumentationSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDocumentationSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetDocumentationSearchRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetDocumentationSearchRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetDocumentationSearchUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetDocumentationSearchTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDocumentationSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetDocumentationSearchServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetDocumentationSearchGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDocumentationSearchOK creates a GetDocumentationSearchOK with default headers values
func NewGetDocumentationSearchOK() *GetDocumentationSearchOK {
	return &GetDocumentationSearchOK{}
}

/*
GetDocumentationSearchOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDocumentationSearchOK struct {
	Payload *models.DocumentationSearchResponse
}

// IsSuccess returns true when this get documentation search o k response has a 2xx status code
func (o *GetDocumentationSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get documentation search o k response has a 3xx status code
func (o *GetDocumentationSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation search o k response has a 4xx status code
func (o *GetDocumentationSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get documentation search o k response has a 5xx status code
func (o *GetDocumentationSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get documentation search o k response a status code equal to that given
func (o *GetDocumentationSearchOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDocumentationSearchOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchOK  %+v", 200, o.Payload)
}

func (o *GetDocumentationSearchOK) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchOK  %+v", 200, o.Payload)
}

func (o *GetDocumentationSearchOK) GetPayload() *models.DocumentationSearchResponse {
	return o.Payload
}

func (o *GetDocumentationSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DocumentationSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationSearchBadRequest creates a GetDocumentationSearchBadRequest with default headers values
func NewGetDocumentationSearchBadRequest() *GetDocumentationSearchBadRequest {
	return &GetDocumentationSearchBadRequest{}
}

/*
GetDocumentationSearchBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetDocumentationSearchBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get documentation search bad request response has a 2xx status code
func (o *GetDocumentationSearchBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get documentation search bad request response has a 3xx status code
func (o *GetDocumentationSearchBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation search bad request response has a 4xx status code
func (o *GetDocumentationSearchBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get documentation search bad request response has a 5xx status code
func (o *GetDocumentationSearchBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get documentation search bad request response a status code equal to that given
func (o *GetDocumentationSearchBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetDocumentationSearchBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchBadRequest  %+v", 400, o.Payload)
}

func (o *GetDocumentationSearchBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchBadRequest  %+v", 400, o.Payload)
}

func (o *GetDocumentationSearchBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationSearchUnauthorized creates a GetDocumentationSearchUnauthorized with default headers values
func NewGetDocumentationSearchUnauthorized() *GetDocumentationSearchUnauthorized {
	return &GetDocumentationSearchUnauthorized{}
}

/*
GetDocumentationSearchUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetDocumentationSearchUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get documentation search unauthorized response has a 2xx status code
func (o *GetDocumentationSearchUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get documentation search unauthorized response has a 3xx status code
func (o *GetDocumentationSearchUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation search unauthorized response has a 4xx status code
func (o *GetDocumentationSearchUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get documentation search unauthorized response has a 5xx status code
func (o *GetDocumentationSearchUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get documentation search unauthorized response a status code equal to that given
func (o *GetDocumentationSearchUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetDocumentationSearchUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDocumentationSearchUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDocumentationSearchUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationSearchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationSearchForbidden creates a GetDocumentationSearchForbidden with default headers values
func NewGetDocumentationSearchForbidden() *GetDocumentationSearchForbidden {
	return &GetDocumentationSearchForbidden{}
}

/*
GetDocumentationSearchForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetDocumentationSearchForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get documentation search forbidden response has a 2xx status code
func (o *GetDocumentationSearchForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get documentation search forbidden response has a 3xx status code
func (o *GetDocumentationSearchForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation search forbidden response has a 4xx status code
func (o *GetDocumentationSearchForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get documentation search forbidden response has a 5xx status code
func (o *GetDocumentationSearchForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get documentation search forbidden response a status code equal to that given
func (o *GetDocumentationSearchForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetDocumentationSearchForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchForbidden  %+v", 403, o.Payload)
}

func (o *GetDocumentationSearchForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchForbidden  %+v", 403, o.Payload)
}

func (o *GetDocumentationSearchForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationSearchNotFound creates a GetDocumentationSearchNotFound with default headers values
func NewGetDocumentationSearchNotFound() *GetDocumentationSearchNotFound {
	return &GetDocumentationSearchNotFound{}
}

/*
GetDocumentationSearchNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetDocumentationSearchNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get documentation search not found response has a 2xx status code
func (o *GetDocumentationSearchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get documentation search not found response has a 3xx status code
func (o *GetDocumentationSearchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation search not found response has a 4xx status code
func (o *GetDocumentationSearchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get documentation search not found response has a 5xx status code
func (o *GetDocumentationSearchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get documentation search not found response a status code equal to that given
func (o *GetDocumentationSearchNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetDocumentationSearchNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchNotFound  %+v", 404, o.Payload)
}

func (o *GetDocumentationSearchNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchNotFound  %+v", 404, o.Payload)
}

func (o *GetDocumentationSearchNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationSearchRequestTimeout creates a GetDocumentationSearchRequestTimeout with default headers values
func NewGetDocumentationSearchRequestTimeout() *GetDocumentationSearchRequestTimeout {
	return &GetDocumentationSearchRequestTimeout{}
}

/*
GetDocumentationSearchRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetDocumentationSearchRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get documentation search request timeout response has a 2xx status code
func (o *GetDocumentationSearchRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get documentation search request timeout response has a 3xx status code
func (o *GetDocumentationSearchRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation search request timeout response has a 4xx status code
func (o *GetDocumentationSearchRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get documentation search request timeout response has a 5xx status code
func (o *GetDocumentationSearchRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get documentation search request timeout response a status code equal to that given
func (o *GetDocumentationSearchRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetDocumentationSearchRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetDocumentationSearchRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetDocumentationSearchRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationSearchRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationSearchRequestEntityTooLarge creates a GetDocumentationSearchRequestEntityTooLarge with default headers values
func NewGetDocumentationSearchRequestEntityTooLarge() *GetDocumentationSearchRequestEntityTooLarge {
	return &GetDocumentationSearchRequestEntityTooLarge{}
}

/*
GetDocumentationSearchRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetDocumentationSearchRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get documentation search request entity too large response has a 2xx status code
func (o *GetDocumentationSearchRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get documentation search request entity too large response has a 3xx status code
func (o *GetDocumentationSearchRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation search request entity too large response has a 4xx status code
func (o *GetDocumentationSearchRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get documentation search request entity too large response has a 5xx status code
func (o *GetDocumentationSearchRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get documentation search request entity too large response a status code equal to that given
func (o *GetDocumentationSearchRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetDocumentationSearchRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetDocumentationSearchRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetDocumentationSearchRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationSearchRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationSearchUnsupportedMediaType creates a GetDocumentationSearchUnsupportedMediaType with default headers values
func NewGetDocumentationSearchUnsupportedMediaType() *GetDocumentationSearchUnsupportedMediaType {
	return &GetDocumentationSearchUnsupportedMediaType{}
}

/*
GetDocumentationSearchUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetDocumentationSearchUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get documentation search unsupported media type response has a 2xx status code
func (o *GetDocumentationSearchUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get documentation search unsupported media type response has a 3xx status code
func (o *GetDocumentationSearchUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation search unsupported media type response has a 4xx status code
func (o *GetDocumentationSearchUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get documentation search unsupported media type response has a 5xx status code
func (o *GetDocumentationSearchUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get documentation search unsupported media type response a status code equal to that given
func (o *GetDocumentationSearchUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetDocumentationSearchUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetDocumentationSearchUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetDocumentationSearchUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationSearchUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationSearchTooManyRequests creates a GetDocumentationSearchTooManyRequests with default headers values
func NewGetDocumentationSearchTooManyRequests() *GetDocumentationSearchTooManyRequests {
	return &GetDocumentationSearchTooManyRequests{}
}

/*
GetDocumentationSearchTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetDocumentationSearchTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get documentation search too many requests response has a 2xx status code
func (o *GetDocumentationSearchTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get documentation search too many requests response has a 3xx status code
func (o *GetDocumentationSearchTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation search too many requests response has a 4xx status code
func (o *GetDocumentationSearchTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get documentation search too many requests response has a 5xx status code
func (o *GetDocumentationSearchTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get documentation search too many requests response a status code equal to that given
func (o *GetDocumentationSearchTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetDocumentationSearchTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDocumentationSearchTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDocumentationSearchTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationSearchTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationSearchInternalServerError creates a GetDocumentationSearchInternalServerError with default headers values
func NewGetDocumentationSearchInternalServerError() *GetDocumentationSearchInternalServerError {
	return &GetDocumentationSearchInternalServerError{}
}

/*
GetDocumentationSearchInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetDocumentationSearchInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get documentation search internal server error response has a 2xx status code
func (o *GetDocumentationSearchInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get documentation search internal server error response has a 3xx status code
func (o *GetDocumentationSearchInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation search internal server error response has a 4xx status code
func (o *GetDocumentationSearchInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get documentation search internal server error response has a 5xx status code
func (o *GetDocumentationSearchInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get documentation search internal server error response a status code equal to that given
func (o *GetDocumentationSearchInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetDocumentationSearchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDocumentationSearchInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDocumentationSearchInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationSearchServiceUnavailable creates a GetDocumentationSearchServiceUnavailable with default headers values
func NewGetDocumentationSearchServiceUnavailable() *GetDocumentationSearchServiceUnavailable {
	return &GetDocumentationSearchServiceUnavailable{}
}

/*
GetDocumentationSearchServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetDocumentationSearchServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get documentation search service unavailable response has a 2xx status code
func (o *GetDocumentationSearchServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get documentation search service unavailable response has a 3xx status code
func (o *GetDocumentationSearchServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation search service unavailable response has a 4xx status code
func (o *GetDocumentationSearchServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get documentation search service unavailable response has a 5xx status code
func (o *GetDocumentationSearchServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get documentation search service unavailable response a status code equal to that given
func (o *GetDocumentationSearchServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetDocumentationSearchServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetDocumentationSearchServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetDocumentationSearchServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationSearchServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationSearchGatewayTimeout creates a GetDocumentationSearchGatewayTimeout with default headers values
func NewGetDocumentationSearchGatewayTimeout() *GetDocumentationSearchGatewayTimeout {
	return &GetDocumentationSearchGatewayTimeout{}
}

/*
GetDocumentationSearchGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetDocumentationSearchGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get documentation search gateway timeout response has a 2xx status code
func (o *GetDocumentationSearchGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get documentation search gateway timeout response has a 3xx status code
func (o *GetDocumentationSearchGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation search gateway timeout response has a 4xx status code
func (o *GetDocumentationSearchGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get documentation search gateway timeout response has a 5xx status code
func (o *GetDocumentationSearchGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get documentation search gateway timeout response a status code equal to that given
func (o *GetDocumentationSearchGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetDocumentationSearchGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetDocumentationSearchGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetDocumentationSearchGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationSearchGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
