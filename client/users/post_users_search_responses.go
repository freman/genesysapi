// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostUsersSearchReader is a Reader for the PostUsersSearch structure.
type PostUsersSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUsersSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUsersSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUsersSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUsersSearchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUsersSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostUsersSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostUsersSearchRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostUsersSearchRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostUsersSearchUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostUsersSearchTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUsersSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostUsersSearchServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostUsersSearchGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostUsersSearchOK creates a PostUsersSearchOK with default headers values
func NewPostUsersSearchOK() *PostUsersSearchOK {
	return &PostUsersSearchOK{}
}

/*
PostUsersSearchOK describes a response with status code 200, with default header values.

successful operation
*/
type PostUsersSearchOK struct {
	Payload *models.UsersSearchResponse
}

// IsSuccess returns true when this post users search o k response has a 2xx status code
func (o *PostUsersSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post users search o k response has a 3xx status code
func (o *PostUsersSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search o k response has a 4xx status code
func (o *PostUsersSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post users search o k response has a 5xx status code
func (o *PostUsersSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post users search o k response a status code equal to that given
func (o *PostUsersSearchOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostUsersSearchOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchOK  %+v", 200, o.Payload)
}

func (o *PostUsersSearchOK) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchOK  %+v", 200, o.Payload)
}

func (o *PostUsersSearchOK) GetPayload() *models.UsersSearchResponse {
	return o.Payload
}

func (o *PostUsersSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UsersSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersSearchBadRequest creates a PostUsersSearchBadRequest with default headers values
func NewPostUsersSearchBadRequest() *PostUsersSearchBadRequest {
	return &PostUsersSearchBadRequest{}
}

/*
PostUsersSearchBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostUsersSearchBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users search bad request response has a 2xx status code
func (o *PostUsersSearchBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users search bad request response has a 3xx status code
func (o *PostUsersSearchBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search bad request response has a 4xx status code
func (o *PostUsersSearchBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users search bad request response has a 5xx status code
func (o *PostUsersSearchBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post users search bad request response a status code equal to that given
func (o *PostUsersSearchBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostUsersSearchBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchBadRequest  %+v", 400, o.Payload)
}

func (o *PostUsersSearchBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchBadRequest  %+v", 400, o.Payload)
}

func (o *PostUsersSearchBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersSearchUnauthorized creates a PostUsersSearchUnauthorized with default headers values
func NewPostUsersSearchUnauthorized() *PostUsersSearchUnauthorized {
	return &PostUsersSearchUnauthorized{}
}

/*
PostUsersSearchUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostUsersSearchUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users search unauthorized response has a 2xx status code
func (o *PostUsersSearchUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users search unauthorized response has a 3xx status code
func (o *PostUsersSearchUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search unauthorized response has a 4xx status code
func (o *PostUsersSearchUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users search unauthorized response has a 5xx status code
func (o *PostUsersSearchUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post users search unauthorized response a status code equal to that given
func (o *PostUsersSearchUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostUsersSearchUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUsersSearchUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUsersSearchUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersSearchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersSearchForbidden creates a PostUsersSearchForbidden with default headers values
func NewPostUsersSearchForbidden() *PostUsersSearchForbidden {
	return &PostUsersSearchForbidden{}
}

/*
PostUsersSearchForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostUsersSearchForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users search forbidden response has a 2xx status code
func (o *PostUsersSearchForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users search forbidden response has a 3xx status code
func (o *PostUsersSearchForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search forbidden response has a 4xx status code
func (o *PostUsersSearchForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users search forbidden response has a 5xx status code
func (o *PostUsersSearchForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post users search forbidden response a status code equal to that given
func (o *PostUsersSearchForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostUsersSearchForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchForbidden  %+v", 403, o.Payload)
}

func (o *PostUsersSearchForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchForbidden  %+v", 403, o.Payload)
}

func (o *PostUsersSearchForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersSearchNotFound creates a PostUsersSearchNotFound with default headers values
func NewPostUsersSearchNotFound() *PostUsersSearchNotFound {
	return &PostUsersSearchNotFound{}
}

/*
PostUsersSearchNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostUsersSearchNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users search not found response has a 2xx status code
func (o *PostUsersSearchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users search not found response has a 3xx status code
func (o *PostUsersSearchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search not found response has a 4xx status code
func (o *PostUsersSearchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users search not found response has a 5xx status code
func (o *PostUsersSearchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post users search not found response a status code equal to that given
func (o *PostUsersSearchNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostUsersSearchNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchNotFound  %+v", 404, o.Payload)
}

func (o *PostUsersSearchNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchNotFound  %+v", 404, o.Payload)
}

func (o *PostUsersSearchNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersSearchRequestTimeout creates a PostUsersSearchRequestTimeout with default headers values
func NewPostUsersSearchRequestTimeout() *PostUsersSearchRequestTimeout {
	return &PostUsersSearchRequestTimeout{}
}

/*
PostUsersSearchRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostUsersSearchRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users search request timeout response has a 2xx status code
func (o *PostUsersSearchRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users search request timeout response has a 3xx status code
func (o *PostUsersSearchRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search request timeout response has a 4xx status code
func (o *PostUsersSearchRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users search request timeout response has a 5xx status code
func (o *PostUsersSearchRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post users search request timeout response a status code equal to that given
func (o *PostUsersSearchRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostUsersSearchRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostUsersSearchRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostUsersSearchRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersSearchRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersSearchRequestEntityTooLarge creates a PostUsersSearchRequestEntityTooLarge with default headers values
func NewPostUsersSearchRequestEntityTooLarge() *PostUsersSearchRequestEntityTooLarge {
	return &PostUsersSearchRequestEntityTooLarge{}
}

/*
PostUsersSearchRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostUsersSearchRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users search request entity too large response has a 2xx status code
func (o *PostUsersSearchRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users search request entity too large response has a 3xx status code
func (o *PostUsersSearchRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search request entity too large response has a 4xx status code
func (o *PostUsersSearchRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users search request entity too large response has a 5xx status code
func (o *PostUsersSearchRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post users search request entity too large response a status code equal to that given
func (o *PostUsersSearchRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostUsersSearchRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostUsersSearchRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostUsersSearchRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersSearchRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersSearchUnsupportedMediaType creates a PostUsersSearchUnsupportedMediaType with default headers values
func NewPostUsersSearchUnsupportedMediaType() *PostUsersSearchUnsupportedMediaType {
	return &PostUsersSearchUnsupportedMediaType{}
}

/*
PostUsersSearchUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostUsersSearchUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users search unsupported media type response has a 2xx status code
func (o *PostUsersSearchUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users search unsupported media type response has a 3xx status code
func (o *PostUsersSearchUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search unsupported media type response has a 4xx status code
func (o *PostUsersSearchUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users search unsupported media type response has a 5xx status code
func (o *PostUsersSearchUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post users search unsupported media type response a status code equal to that given
func (o *PostUsersSearchUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostUsersSearchUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostUsersSearchUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostUsersSearchUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersSearchUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersSearchTooManyRequests creates a PostUsersSearchTooManyRequests with default headers values
func NewPostUsersSearchTooManyRequests() *PostUsersSearchTooManyRequests {
	return &PostUsersSearchTooManyRequests{}
}

/*
PostUsersSearchTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostUsersSearchTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users search too many requests response has a 2xx status code
func (o *PostUsersSearchTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users search too many requests response has a 3xx status code
func (o *PostUsersSearchTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search too many requests response has a 4xx status code
func (o *PostUsersSearchTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users search too many requests response has a 5xx status code
func (o *PostUsersSearchTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post users search too many requests response a status code equal to that given
func (o *PostUsersSearchTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostUsersSearchTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostUsersSearchTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostUsersSearchTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersSearchTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersSearchInternalServerError creates a PostUsersSearchInternalServerError with default headers values
func NewPostUsersSearchInternalServerError() *PostUsersSearchInternalServerError {
	return &PostUsersSearchInternalServerError{}
}

/*
PostUsersSearchInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostUsersSearchInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users search internal server error response has a 2xx status code
func (o *PostUsersSearchInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users search internal server error response has a 3xx status code
func (o *PostUsersSearchInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search internal server error response has a 4xx status code
func (o *PostUsersSearchInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post users search internal server error response has a 5xx status code
func (o *PostUsersSearchInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post users search internal server error response a status code equal to that given
func (o *PostUsersSearchInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostUsersSearchInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUsersSearchInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUsersSearchInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersSearchServiceUnavailable creates a PostUsersSearchServiceUnavailable with default headers values
func NewPostUsersSearchServiceUnavailable() *PostUsersSearchServiceUnavailable {
	return &PostUsersSearchServiceUnavailable{}
}

/*
PostUsersSearchServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostUsersSearchServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users search service unavailable response has a 2xx status code
func (o *PostUsersSearchServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users search service unavailable response has a 3xx status code
func (o *PostUsersSearchServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search service unavailable response has a 4xx status code
func (o *PostUsersSearchServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post users search service unavailable response has a 5xx status code
func (o *PostUsersSearchServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post users search service unavailable response a status code equal to that given
func (o *PostUsersSearchServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostUsersSearchServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUsersSearchServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUsersSearchServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersSearchServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersSearchGatewayTimeout creates a PostUsersSearchGatewayTimeout with default headers values
func NewPostUsersSearchGatewayTimeout() *PostUsersSearchGatewayTimeout {
	return &PostUsersSearchGatewayTimeout{}
}

/*
PostUsersSearchGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostUsersSearchGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users search gateway timeout response has a 2xx status code
func (o *PostUsersSearchGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users search gateway timeout response has a 3xx status code
func (o *PostUsersSearchGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search gateway timeout response has a 4xx status code
func (o *PostUsersSearchGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post users search gateway timeout response has a 5xx status code
func (o *PostUsersSearchGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post users search gateway timeout response a status code equal to that given
func (o *PostUsersSearchGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostUsersSearchGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUsersSearchGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search][%d] postUsersSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUsersSearchGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersSearchGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
