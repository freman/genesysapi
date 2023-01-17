// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostGroupsSearchReader is a Reader for the PostGroupsSearch structure.
type PostGroupsSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostGroupsSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostGroupsSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostGroupsSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostGroupsSearchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostGroupsSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostGroupsSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostGroupsSearchRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostGroupsSearchRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostGroupsSearchUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostGroupsSearchTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostGroupsSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostGroupsSearchServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostGroupsSearchGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostGroupsSearchOK creates a PostGroupsSearchOK with default headers values
func NewPostGroupsSearchOK() *PostGroupsSearchOK {
	return &PostGroupsSearchOK{}
}

/*
PostGroupsSearchOK describes a response with status code 200, with default header values.

successful operation
*/
type PostGroupsSearchOK struct {
	Payload *models.GroupsSearchResponse
}

// IsSuccess returns true when this post groups search o k response has a 2xx status code
func (o *PostGroupsSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post groups search o k response has a 3xx status code
func (o *PostGroupsSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post groups search o k response has a 4xx status code
func (o *PostGroupsSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post groups search o k response has a 5xx status code
func (o *PostGroupsSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post groups search o k response a status code equal to that given
func (o *PostGroupsSearchOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostGroupsSearchOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchOK  %+v", 200, o.Payload)
}

func (o *PostGroupsSearchOK) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchOK  %+v", 200, o.Payload)
}

func (o *PostGroupsSearchOK) GetPayload() *models.GroupsSearchResponse {
	return o.Payload
}

func (o *PostGroupsSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupsSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsSearchBadRequest creates a PostGroupsSearchBadRequest with default headers values
func NewPostGroupsSearchBadRequest() *PostGroupsSearchBadRequest {
	return &PostGroupsSearchBadRequest{}
}

/*
PostGroupsSearchBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostGroupsSearchBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post groups search bad request response has a 2xx status code
func (o *PostGroupsSearchBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post groups search bad request response has a 3xx status code
func (o *PostGroupsSearchBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post groups search bad request response has a 4xx status code
func (o *PostGroupsSearchBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post groups search bad request response has a 5xx status code
func (o *PostGroupsSearchBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post groups search bad request response a status code equal to that given
func (o *PostGroupsSearchBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostGroupsSearchBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchBadRequest  %+v", 400, o.Payload)
}

func (o *PostGroupsSearchBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchBadRequest  %+v", 400, o.Payload)
}

func (o *PostGroupsSearchBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupsSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsSearchUnauthorized creates a PostGroupsSearchUnauthorized with default headers values
func NewPostGroupsSearchUnauthorized() *PostGroupsSearchUnauthorized {
	return &PostGroupsSearchUnauthorized{}
}

/*
PostGroupsSearchUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostGroupsSearchUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post groups search unauthorized response has a 2xx status code
func (o *PostGroupsSearchUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post groups search unauthorized response has a 3xx status code
func (o *PostGroupsSearchUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post groups search unauthorized response has a 4xx status code
func (o *PostGroupsSearchUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post groups search unauthorized response has a 5xx status code
func (o *PostGroupsSearchUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post groups search unauthorized response a status code equal to that given
func (o *PostGroupsSearchUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostGroupsSearchUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *PostGroupsSearchUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *PostGroupsSearchUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupsSearchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsSearchForbidden creates a PostGroupsSearchForbidden with default headers values
func NewPostGroupsSearchForbidden() *PostGroupsSearchForbidden {
	return &PostGroupsSearchForbidden{}
}

/*
PostGroupsSearchForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostGroupsSearchForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post groups search forbidden response has a 2xx status code
func (o *PostGroupsSearchForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post groups search forbidden response has a 3xx status code
func (o *PostGroupsSearchForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post groups search forbidden response has a 4xx status code
func (o *PostGroupsSearchForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post groups search forbidden response has a 5xx status code
func (o *PostGroupsSearchForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post groups search forbidden response a status code equal to that given
func (o *PostGroupsSearchForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostGroupsSearchForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchForbidden  %+v", 403, o.Payload)
}

func (o *PostGroupsSearchForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchForbidden  %+v", 403, o.Payload)
}

func (o *PostGroupsSearchForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupsSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsSearchNotFound creates a PostGroupsSearchNotFound with default headers values
func NewPostGroupsSearchNotFound() *PostGroupsSearchNotFound {
	return &PostGroupsSearchNotFound{}
}

/*
PostGroupsSearchNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostGroupsSearchNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post groups search not found response has a 2xx status code
func (o *PostGroupsSearchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post groups search not found response has a 3xx status code
func (o *PostGroupsSearchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post groups search not found response has a 4xx status code
func (o *PostGroupsSearchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post groups search not found response has a 5xx status code
func (o *PostGroupsSearchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post groups search not found response a status code equal to that given
func (o *PostGroupsSearchNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostGroupsSearchNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchNotFound  %+v", 404, o.Payload)
}

func (o *PostGroupsSearchNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchNotFound  %+v", 404, o.Payload)
}

func (o *PostGroupsSearchNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupsSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsSearchRequestTimeout creates a PostGroupsSearchRequestTimeout with default headers values
func NewPostGroupsSearchRequestTimeout() *PostGroupsSearchRequestTimeout {
	return &PostGroupsSearchRequestTimeout{}
}

/*
PostGroupsSearchRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostGroupsSearchRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post groups search request timeout response has a 2xx status code
func (o *PostGroupsSearchRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post groups search request timeout response has a 3xx status code
func (o *PostGroupsSearchRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post groups search request timeout response has a 4xx status code
func (o *PostGroupsSearchRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post groups search request timeout response has a 5xx status code
func (o *PostGroupsSearchRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post groups search request timeout response a status code equal to that given
func (o *PostGroupsSearchRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostGroupsSearchRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostGroupsSearchRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostGroupsSearchRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupsSearchRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsSearchRequestEntityTooLarge creates a PostGroupsSearchRequestEntityTooLarge with default headers values
func NewPostGroupsSearchRequestEntityTooLarge() *PostGroupsSearchRequestEntityTooLarge {
	return &PostGroupsSearchRequestEntityTooLarge{}
}

/*
PostGroupsSearchRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostGroupsSearchRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post groups search request entity too large response has a 2xx status code
func (o *PostGroupsSearchRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post groups search request entity too large response has a 3xx status code
func (o *PostGroupsSearchRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post groups search request entity too large response has a 4xx status code
func (o *PostGroupsSearchRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post groups search request entity too large response has a 5xx status code
func (o *PostGroupsSearchRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post groups search request entity too large response a status code equal to that given
func (o *PostGroupsSearchRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostGroupsSearchRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostGroupsSearchRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostGroupsSearchRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupsSearchRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsSearchUnsupportedMediaType creates a PostGroupsSearchUnsupportedMediaType with default headers values
func NewPostGroupsSearchUnsupportedMediaType() *PostGroupsSearchUnsupportedMediaType {
	return &PostGroupsSearchUnsupportedMediaType{}
}

/*
PostGroupsSearchUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostGroupsSearchUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post groups search unsupported media type response has a 2xx status code
func (o *PostGroupsSearchUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post groups search unsupported media type response has a 3xx status code
func (o *PostGroupsSearchUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post groups search unsupported media type response has a 4xx status code
func (o *PostGroupsSearchUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post groups search unsupported media type response has a 5xx status code
func (o *PostGroupsSearchUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post groups search unsupported media type response a status code equal to that given
func (o *PostGroupsSearchUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostGroupsSearchUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostGroupsSearchUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostGroupsSearchUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupsSearchUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsSearchTooManyRequests creates a PostGroupsSearchTooManyRequests with default headers values
func NewPostGroupsSearchTooManyRequests() *PostGroupsSearchTooManyRequests {
	return &PostGroupsSearchTooManyRequests{}
}

/*
PostGroupsSearchTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostGroupsSearchTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post groups search too many requests response has a 2xx status code
func (o *PostGroupsSearchTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post groups search too many requests response has a 3xx status code
func (o *PostGroupsSearchTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post groups search too many requests response has a 4xx status code
func (o *PostGroupsSearchTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post groups search too many requests response has a 5xx status code
func (o *PostGroupsSearchTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post groups search too many requests response a status code equal to that given
func (o *PostGroupsSearchTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostGroupsSearchTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostGroupsSearchTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostGroupsSearchTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupsSearchTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsSearchInternalServerError creates a PostGroupsSearchInternalServerError with default headers values
func NewPostGroupsSearchInternalServerError() *PostGroupsSearchInternalServerError {
	return &PostGroupsSearchInternalServerError{}
}

/*
PostGroupsSearchInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostGroupsSearchInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post groups search internal server error response has a 2xx status code
func (o *PostGroupsSearchInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post groups search internal server error response has a 3xx status code
func (o *PostGroupsSearchInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post groups search internal server error response has a 4xx status code
func (o *PostGroupsSearchInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post groups search internal server error response has a 5xx status code
func (o *PostGroupsSearchInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post groups search internal server error response a status code equal to that given
func (o *PostGroupsSearchInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostGroupsSearchInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *PostGroupsSearchInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *PostGroupsSearchInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupsSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsSearchServiceUnavailable creates a PostGroupsSearchServiceUnavailable with default headers values
func NewPostGroupsSearchServiceUnavailable() *PostGroupsSearchServiceUnavailable {
	return &PostGroupsSearchServiceUnavailable{}
}

/*
PostGroupsSearchServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostGroupsSearchServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post groups search service unavailable response has a 2xx status code
func (o *PostGroupsSearchServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post groups search service unavailable response has a 3xx status code
func (o *PostGroupsSearchServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post groups search service unavailable response has a 4xx status code
func (o *PostGroupsSearchServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post groups search service unavailable response has a 5xx status code
func (o *PostGroupsSearchServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post groups search service unavailable response a status code equal to that given
func (o *PostGroupsSearchServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostGroupsSearchServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostGroupsSearchServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostGroupsSearchServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupsSearchServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsSearchGatewayTimeout creates a PostGroupsSearchGatewayTimeout with default headers values
func NewPostGroupsSearchGatewayTimeout() *PostGroupsSearchGatewayTimeout {
	return &PostGroupsSearchGatewayTimeout{}
}

/*
PostGroupsSearchGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostGroupsSearchGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post groups search gateway timeout response has a 2xx status code
func (o *PostGroupsSearchGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post groups search gateway timeout response has a 3xx status code
func (o *PostGroupsSearchGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post groups search gateway timeout response has a 4xx status code
func (o *PostGroupsSearchGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post groups search gateway timeout response has a 5xx status code
func (o *PostGroupsSearchGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post groups search gateway timeout response a status code equal to that given
func (o *PostGroupsSearchGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostGroupsSearchGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostGroupsSearchGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/groups/search][%d] postGroupsSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostGroupsSearchGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGroupsSearchGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
