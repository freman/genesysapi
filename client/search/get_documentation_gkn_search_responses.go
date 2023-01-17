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

// GetDocumentationGknSearchReader is a Reader for the GetDocumentationGknSearch structure.
type GetDocumentationGknSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDocumentationGknSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDocumentationGknSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDocumentationGknSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDocumentationGknSearchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDocumentationGknSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDocumentationGknSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetDocumentationGknSearchRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetDocumentationGknSearchRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetDocumentationGknSearchUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetDocumentationGknSearchTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDocumentationGknSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetDocumentationGknSearchServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetDocumentationGknSearchGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDocumentationGknSearchOK creates a GetDocumentationGknSearchOK with default headers values
func NewGetDocumentationGknSearchOK() *GetDocumentationGknSearchOK {
	return &GetDocumentationGknSearchOK{}
}

/*
GetDocumentationGknSearchOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDocumentationGknSearchOK struct {
	Payload *models.GKNDocumentationSearchResponse
}

// IsSuccess returns true when this get documentation gkn search o k response has a 2xx status code
func (o *GetDocumentationGknSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get documentation gkn search o k response has a 3xx status code
func (o *GetDocumentationGknSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation gkn search o k response has a 4xx status code
func (o *GetDocumentationGknSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get documentation gkn search o k response has a 5xx status code
func (o *GetDocumentationGknSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get documentation gkn search o k response a status code equal to that given
func (o *GetDocumentationGknSearchOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDocumentationGknSearchOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchOK  %+v", 200, o.Payload)
}

func (o *GetDocumentationGknSearchOK) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchOK  %+v", 200, o.Payload)
}

func (o *GetDocumentationGknSearchOK) GetPayload() *models.GKNDocumentationSearchResponse {
	return o.Payload
}

func (o *GetDocumentationGknSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GKNDocumentationSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationGknSearchBadRequest creates a GetDocumentationGknSearchBadRequest with default headers values
func NewGetDocumentationGknSearchBadRequest() *GetDocumentationGknSearchBadRequest {
	return &GetDocumentationGknSearchBadRequest{}
}

/*
GetDocumentationGknSearchBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetDocumentationGknSearchBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get documentation gkn search bad request response has a 2xx status code
func (o *GetDocumentationGknSearchBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get documentation gkn search bad request response has a 3xx status code
func (o *GetDocumentationGknSearchBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation gkn search bad request response has a 4xx status code
func (o *GetDocumentationGknSearchBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get documentation gkn search bad request response has a 5xx status code
func (o *GetDocumentationGknSearchBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get documentation gkn search bad request response a status code equal to that given
func (o *GetDocumentationGknSearchBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetDocumentationGknSearchBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchBadRequest  %+v", 400, o.Payload)
}

func (o *GetDocumentationGknSearchBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchBadRequest  %+v", 400, o.Payload)
}

func (o *GetDocumentationGknSearchBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationGknSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationGknSearchUnauthorized creates a GetDocumentationGknSearchUnauthorized with default headers values
func NewGetDocumentationGknSearchUnauthorized() *GetDocumentationGknSearchUnauthorized {
	return &GetDocumentationGknSearchUnauthorized{}
}

/*
GetDocumentationGknSearchUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetDocumentationGknSearchUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get documentation gkn search unauthorized response has a 2xx status code
func (o *GetDocumentationGknSearchUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get documentation gkn search unauthorized response has a 3xx status code
func (o *GetDocumentationGknSearchUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation gkn search unauthorized response has a 4xx status code
func (o *GetDocumentationGknSearchUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get documentation gkn search unauthorized response has a 5xx status code
func (o *GetDocumentationGknSearchUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get documentation gkn search unauthorized response a status code equal to that given
func (o *GetDocumentationGknSearchUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetDocumentationGknSearchUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDocumentationGknSearchUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDocumentationGknSearchUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationGknSearchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationGknSearchForbidden creates a GetDocumentationGknSearchForbidden with default headers values
func NewGetDocumentationGknSearchForbidden() *GetDocumentationGknSearchForbidden {
	return &GetDocumentationGknSearchForbidden{}
}

/*
GetDocumentationGknSearchForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetDocumentationGknSearchForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get documentation gkn search forbidden response has a 2xx status code
func (o *GetDocumentationGknSearchForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get documentation gkn search forbidden response has a 3xx status code
func (o *GetDocumentationGknSearchForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation gkn search forbidden response has a 4xx status code
func (o *GetDocumentationGknSearchForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get documentation gkn search forbidden response has a 5xx status code
func (o *GetDocumentationGknSearchForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get documentation gkn search forbidden response a status code equal to that given
func (o *GetDocumentationGknSearchForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetDocumentationGknSearchForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchForbidden  %+v", 403, o.Payload)
}

func (o *GetDocumentationGknSearchForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchForbidden  %+v", 403, o.Payload)
}

func (o *GetDocumentationGknSearchForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationGknSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationGknSearchNotFound creates a GetDocumentationGknSearchNotFound with default headers values
func NewGetDocumentationGknSearchNotFound() *GetDocumentationGknSearchNotFound {
	return &GetDocumentationGknSearchNotFound{}
}

/*
GetDocumentationGknSearchNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetDocumentationGknSearchNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get documentation gkn search not found response has a 2xx status code
func (o *GetDocumentationGknSearchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get documentation gkn search not found response has a 3xx status code
func (o *GetDocumentationGknSearchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation gkn search not found response has a 4xx status code
func (o *GetDocumentationGknSearchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get documentation gkn search not found response has a 5xx status code
func (o *GetDocumentationGknSearchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get documentation gkn search not found response a status code equal to that given
func (o *GetDocumentationGknSearchNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetDocumentationGknSearchNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchNotFound  %+v", 404, o.Payload)
}

func (o *GetDocumentationGknSearchNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchNotFound  %+v", 404, o.Payload)
}

func (o *GetDocumentationGknSearchNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationGknSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationGknSearchRequestTimeout creates a GetDocumentationGknSearchRequestTimeout with default headers values
func NewGetDocumentationGknSearchRequestTimeout() *GetDocumentationGknSearchRequestTimeout {
	return &GetDocumentationGknSearchRequestTimeout{}
}

/*
GetDocumentationGknSearchRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetDocumentationGknSearchRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get documentation gkn search request timeout response has a 2xx status code
func (o *GetDocumentationGknSearchRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get documentation gkn search request timeout response has a 3xx status code
func (o *GetDocumentationGknSearchRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation gkn search request timeout response has a 4xx status code
func (o *GetDocumentationGknSearchRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get documentation gkn search request timeout response has a 5xx status code
func (o *GetDocumentationGknSearchRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get documentation gkn search request timeout response a status code equal to that given
func (o *GetDocumentationGknSearchRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetDocumentationGknSearchRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetDocumentationGknSearchRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetDocumentationGknSearchRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationGknSearchRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationGknSearchRequestEntityTooLarge creates a GetDocumentationGknSearchRequestEntityTooLarge with default headers values
func NewGetDocumentationGknSearchRequestEntityTooLarge() *GetDocumentationGknSearchRequestEntityTooLarge {
	return &GetDocumentationGknSearchRequestEntityTooLarge{}
}

/*
GetDocumentationGknSearchRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetDocumentationGknSearchRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get documentation gkn search request entity too large response has a 2xx status code
func (o *GetDocumentationGknSearchRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get documentation gkn search request entity too large response has a 3xx status code
func (o *GetDocumentationGknSearchRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation gkn search request entity too large response has a 4xx status code
func (o *GetDocumentationGknSearchRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get documentation gkn search request entity too large response has a 5xx status code
func (o *GetDocumentationGknSearchRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get documentation gkn search request entity too large response a status code equal to that given
func (o *GetDocumentationGknSearchRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetDocumentationGknSearchRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetDocumentationGknSearchRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetDocumentationGknSearchRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationGknSearchRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationGknSearchUnsupportedMediaType creates a GetDocumentationGknSearchUnsupportedMediaType with default headers values
func NewGetDocumentationGknSearchUnsupportedMediaType() *GetDocumentationGknSearchUnsupportedMediaType {
	return &GetDocumentationGknSearchUnsupportedMediaType{}
}

/*
GetDocumentationGknSearchUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetDocumentationGknSearchUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get documentation gkn search unsupported media type response has a 2xx status code
func (o *GetDocumentationGknSearchUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get documentation gkn search unsupported media type response has a 3xx status code
func (o *GetDocumentationGknSearchUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation gkn search unsupported media type response has a 4xx status code
func (o *GetDocumentationGknSearchUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get documentation gkn search unsupported media type response has a 5xx status code
func (o *GetDocumentationGknSearchUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get documentation gkn search unsupported media type response a status code equal to that given
func (o *GetDocumentationGknSearchUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetDocumentationGknSearchUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetDocumentationGknSearchUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetDocumentationGknSearchUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationGknSearchUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationGknSearchTooManyRequests creates a GetDocumentationGknSearchTooManyRequests with default headers values
func NewGetDocumentationGknSearchTooManyRequests() *GetDocumentationGknSearchTooManyRequests {
	return &GetDocumentationGknSearchTooManyRequests{}
}

/*
GetDocumentationGknSearchTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetDocumentationGknSearchTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get documentation gkn search too many requests response has a 2xx status code
func (o *GetDocumentationGknSearchTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get documentation gkn search too many requests response has a 3xx status code
func (o *GetDocumentationGknSearchTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation gkn search too many requests response has a 4xx status code
func (o *GetDocumentationGknSearchTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get documentation gkn search too many requests response has a 5xx status code
func (o *GetDocumentationGknSearchTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get documentation gkn search too many requests response a status code equal to that given
func (o *GetDocumentationGknSearchTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetDocumentationGknSearchTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDocumentationGknSearchTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDocumentationGknSearchTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationGknSearchTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationGknSearchInternalServerError creates a GetDocumentationGknSearchInternalServerError with default headers values
func NewGetDocumentationGknSearchInternalServerError() *GetDocumentationGknSearchInternalServerError {
	return &GetDocumentationGknSearchInternalServerError{}
}

/*
GetDocumentationGknSearchInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetDocumentationGknSearchInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get documentation gkn search internal server error response has a 2xx status code
func (o *GetDocumentationGknSearchInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get documentation gkn search internal server error response has a 3xx status code
func (o *GetDocumentationGknSearchInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation gkn search internal server error response has a 4xx status code
func (o *GetDocumentationGknSearchInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get documentation gkn search internal server error response has a 5xx status code
func (o *GetDocumentationGknSearchInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get documentation gkn search internal server error response a status code equal to that given
func (o *GetDocumentationGknSearchInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetDocumentationGknSearchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDocumentationGknSearchInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDocumentationGknSearchInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationGknSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationGknSearchServiceUnavailable creates a GetDocumentationGknSearchServiceUnavailable with default headers values
func NewGetDocumentationGknSearchServiceUnavailable() *GetDocumentationGknSearchServiceUnavailable {
	return &GetDocumentationGknSearchServiceUnavailable{}
}

/*
GetDocumentationGknSearchServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetDocumentationGknSearchServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get documentation gkn search service unavailable response has a 2xx status code
func (o *GetDocumentationGknSearchServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get documentation gkn search service unavailable response has a 3xx status code
func (o *GetDocumentationGknSearchServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation gkn search service unavailable response has a 4xx status code
func (o *GetDocumentationGknSearchServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get documentation gkn search service unavailable response has a 5xx status code
func (o *GetDocumentationGknSearchServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get documentation gkn search service unavailable response a status code equal to that given
func (o *GetDocumentationGknSearchServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetDocumentationGknSearchServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetDocumentationGknSearchServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetDocumentationGknSearchServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationGknSearchServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationGknSearchGatewayTimeout creates a GetDocumentationGknSearchGatewayTimeout with default headers values
func NewGetDocumentationGknSearchGatewayTimeout() *GetDocumentationGknSearchGatewayTimeout {
	return &GetDocumentationGknSearchGatewayTimeout{}
}

/*
GetDocumentationGknSearchGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetDocumentationGknSearchGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get documentation gkn search gateway timeout response has a 2xx status code
func (o *GetDocumentationGknSearchGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get documentation gkn search gateway timeout response has a 3xx status code
func (o *GetDocumentationGknSearchGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get documentation gkn search gateway timeout response has a 4xx status code
func (o *GetDocumentationGknSearchGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get documentation gkn search gateway timeout response has a 5xx status code
func (o *GetDocumentationGknSearchGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get documentation gkn search gateway timeout response a status code equal to that given
func (o *GetDocumentationGknSearchGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetDocumentationGknSearchGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetDocumentationGknSearchGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/documentation/gkn/search][%d] getDocumentationGknSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetDocumentationGknSearchGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationGknSearchGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
