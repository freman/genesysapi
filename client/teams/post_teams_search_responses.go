// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostTeamsSearchReader is a Reader for the PostTeamsSearch structure.
type PostTeamsSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTeamsSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTeamsSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTeamsSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTeamsSearchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTeamsSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTeamsSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostTeamsSearchRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTeamsSearchRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTeamsSearchUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTeamsSearchTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTeamsSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTeamsSearchServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTeamsSearchGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTeamsSearchOK creates a PostTeamsSearchOK with default headers values
func NewPostTeamsSearchOK() *PostTeamsSearchOK {
	return &PostTeamsSearchOK{}
}

/*
PostTeamsSearchOK describes a response with status code 200, with default header values.

successful operation
*/
type PostTeamsSearchOK struct {
	Payload *models.TeamsSearchResponse
}

// IsSuccess returns true when this post teams search o k response has a 2xx status code
func (o *PostTeamsSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post teams search o k response has a 3xx status code
func (o *PostTeamsSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post teams search o k response has a 4xx status code
func (o *PostTeamsSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post teams search o k response has a 5xx status code
func (o *PostTeamsSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post teams search o k response a status code equal to that given
func (o *PostTeamsSearchOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostTeamsSearchOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchOK  %+v", 200, o.Payload)
}

func (o *PostTeamsSearchOK) String() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchOK  %+v", 200, o.Payload)
}

func (o *PostTeamsSearchOK) GetPayload() *models.TeamsSearchResponse {
	return o.Payload
}

func (o *PostTeamsSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TeamsSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsSearchBadRequest creates a PostTeamsSearchBadRequest with default headers values
func NewPostTeamsSearchBadRequest() *PostTeamsSearchBadRequest {
	return &PostTeamsSearchBadRequest{}
}

/*
PostTeamsSearchBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTeamsSearchBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post teams search bad request response has a 2xx status code
func (o *PostTeamsSearchBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post teams search bad request response has a 3xx status code
func (o *PostTeamsSearchBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post teams search bad request response has a 4xx status code
func (o *PostTeamsSearchBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post teams search bad request response has a 5xx status code
func (o *PostTeamsSearchBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post teams search bad request response a status code equal to that given
func (o *PostTeamsSearchBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostTeamsSearchBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchBadRequest  %+v", 400, o.Payload)
}

func (o *PostTeamsSearchBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchBadRequest  %+v", 400, o.Payload)
}

func (o *PostTeamsSearchBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTeamsSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsSearchUnauthorized creates a PostTeamsSearchUnauthorized with default headers values
func NewPostTeamsSearchUnauthorized() *PostTeamsSearchUnauthorized {
	return &PostTeamsSearchUnauthorized{}
}

/*
PostTeamsSearchUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTeamsSearchUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post teams search unauthorized response has a 2xx status code
func (o *PostTeamsSearchUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post teams search unauthorized response has a 3xx status code
func (o *PostTeamsSearchUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post teams search unauthorized response has a 4xx status code
func (o *PostTeamsSearchUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post teams search unauthorized response has a 5xx status code
func (o *PostTeamsSearchUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post teams search unauthorized response a status code equal to that given
func (o *PostTeamsSearchUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostTeamsSearchUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTeamsSearchUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTeamsSearchUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTeamsSearchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsSearchForbidden creates a PostTeamsSearchForbidden with default headers values
func NewPostTeamsSearchForbidden() *PostTeamsSearchForbidden {
	return &PostTeamsSearchForbidden{}
}

/*
PostTeamsSearchForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostTeamsSearchForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post teams search forbidden response has a 2xx status code
func (o *PostTeamsSearchForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post teams search forbidden response has a 3xx status code
func (o *PostTeamsSearchForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post teams search forbidden response has a 4xx status code
func (o *PostTeamsSearchForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post teams search forbidden response has a 5xx status code
func (o *PostTeamsSearchForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post teams search forbidden response a status code equal to that given
func (o *PostTeamsSearchForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostTeamsSearchForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchForbidden  %+v", 403, o.Payload)
}

func (o *PostTeamsSearchForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchForbidden  %+v", 403, o.Payload)
}

func (o *PostTeamsSearchForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTeamsSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsSearchNotFound creates a PostTeamsSearchNotFound with default headers values
func NewPostTeamsSearchNotFound() *PostTeamsSearchNotFound {
	return &PostTeamsSearchNotFound{}
}

/*
PostTeamsSearchNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostTeamsSearchNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post teams search not found response has a 2xx status code
func (o *PostTeamsSearchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post teams search not found response has a 3xx status code
func (o *PostTeamsSearchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post teams search not found response has a 4xx status code
func (o *PostTeamsSearchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post teams search not found response has a 5xx status code
func (o *PostTeamsSearchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post teams search not found response a status code equal to that given
func (o *PostTeamsSearchNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostTeamsSearchNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchNotFound  %+v", 404, o.Payload)
}

func (o *PostTeamsSearchNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchNotFound  %+v", 404, o.Payload)
}

func (o *PostTeamsSearchNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTeamsSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsSearchRequestTimeout creates a PostTeamsSearchRequestTimeout with default headers values
func NewPostTeamsSearchRequestTimeout() *PostTeamsSearchRequestTimeout {
	return &PostTeamsSearchRequestTimeout{}
}

/*
PostTeamsSearchRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostTeamsSearchRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post teams search request timeout response has a 2xx status code
func (o *PostTeamsSearchRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post teams search request timeout response has a 3xx status code
func (o *PostTeamsSearchRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post teams search request timeout response has a 4xx status code
func (o *PostTeamsSearchRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post teams search request timeout response has a 5xx status code
func (o *PostTeamsSearchRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post teams search request timeout response a status code equal to that given
func (o *PostTeamsSearchRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostTeamsSearchRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostTeamsSearchRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostTeamsSearchRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTeamsSearchRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsSearchRequestEntityTooLarge creates a PostTeamsSearchRequestEntityTooLarge with default headers values
func NewPostTeamsSearchRequestEntityTooLarge() *PostTeamsSearchRequestEntityTooLarge {
	return &PostTeamsSearchRequestEntityTooLarge{}
}

/*
PostTeamsSearchRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostTeamsSearchRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post teams search request entity too large response has a 2xx status code
func (o *PostTeamsSearchRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post teams search request entity too large response has a 3xx status code
func (o *PostTeamsSearchRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post teams search request entity too large response has a 4xx status code
func (o *PostTeamsSearchRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post teams search request entity too large response has a 5xx status code
func (o *PostTeamsSearchRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post teams search request entity too large response a status code equal to that given
func (o *PostTeamsSearchRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostTeamsSearchRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTeamsSearchRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTeamsSearchRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTeamsSearchRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsSearchUnsupportedMediaType creates a PostTeamsSearchUnsupportedMediaType with default headers values
func NewPostTeamsSearchUnsupportedMediaType() *PostTeamsSearchUnsupportedMediaType {
	return &PostTeamsSearchUnsupportedMediaType{}
}

/*
PostTeamsSearchUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTeamsSearchUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post teams search unsupported media type response has a 2xx status code
func (o *PostTeamsSearchUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post teams search unsupported media type response has a 3xx status code
func (o *PostTeamsSearchUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post teams search unsupported media type response has a 4xx status code
func (o *PostTeamsSearchUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post teams search unsupported media type response has a 5xx status code
func (o *PostTeamsSearchUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post teams search unsupported media type response a status code equal to that given
func (o *PostTeamsSearchUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostTeamsSearchUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTeamsSearchUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTeamsSearchUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTeamsSearchUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsSearchTooManyRequests creates a PostTeamsSearchTooManyRequests with default headers values
func NewPostTeamsSearchTooManyRequests() *PostTeamsSearchTooManyRequests {
	return &PostTeamsSearchTooManyRequests{}
}

/*
PostTeamsSearchTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostTeamsSearchTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post teams search too many requests response has a 2xx status code
func (o *PostTeamsSearchTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post teams search too many requests response has a 3xx status code
func (o *PostTeamsSearchTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post teams search too many requests response has a 4xx status code
func (o *PostTeamsSearchTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post teams search too many requests response has a 5xx status code
func (o *PostTeamsSearchTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post teams search too many requests response a status code equal to that given
func (o *PostTeamsSearchTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostTeamsSearchTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTeamsSearchTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTeamsSearchTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTeamsSearchTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsSearchInternalServerError creates a PostTeamsSearchInternalServerError with default headers values
func NewPostTeamsSearchInternalServerError() *PostTeamsSearchInternalServerError {
	return &PostTeamsSearchInternalServerError{}
}

/*
PostTeamsSearchInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTeamsSearchInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post teams search internal server error response has a 2xx status code
func (o *PostTeamsSearchInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post teams search internal server error response has a 3xx status code
func (o *PostTeamsSearchInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post teams search internal server error response has a 4xx status code
func (o *PostTeamsSearchInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post teams search internal server error response has a 5xx status code
func (o *PostTeamsSearchInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post teams search internal server error response a status code equal to that given
func (o *PostTeamsSearchInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostTeamsSearchInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTeamsSearchInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTeamsSearchInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTeamsSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsSearchServiceUnavailable creates a PostTeamsSearchServiceUnavailable with default headers values
func NewPostTeamsSearchServiceUnavailable() *PostTeamsSearchServiceUnavailable {
	return &PostTeamsSearchServiceUnavailable{}
}

/*
PostTeamsSearchServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTeamsSearchServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post teams search service unavailable response has a 2xx status code
func (o *PostTeamsSearchServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post teams search service unavailable response has a 3xx status code
func (o *PostTeamsSearchServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post teams search service unavailable response has a 4xx status code
func (o *PostTeamsSearchServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post teams search service unavailable response has a 5xx status code
func (o *PostTeamsSearchServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post teams search service unavailable response a status code equal to that given
func (o *PostTeamsSearchServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostTeamsSearchServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTeamsSearchServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTeamsSearchServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTeamsSearchServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTeamsSearchGatewayTimeout creates a PostTeamsSearchGatewayTimeout with default headers values
func NewPostTeamsSearchGatewayTimeout() *PostTeamsSearchGatewayTimeout {
	return &PostTeamsSearchGatewayTimeout{}
}

/*
PostTeamsSearchGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostTeamsSearchGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post teams search gateway timeout response has a 2xx status code
func (o *PostTeamsSearchGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post teams search gateway timeout response has a 3xx status code
func (o *PostTeamsSearchGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post teams search gateway timeout response has a 4xx status code
func (o *PostTeamsSearchGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post teams search gateway timeout response has a 5xx status code
func (o *PostTeamsSearchGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post teams search gateway timeout response a status code equal to that given
func (o *PostTeamsSearchGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostTeamsSearchGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTeamsSearchGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/teams/search][%d] postTeamsSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTeamsSearchGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTeamsSearchGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
