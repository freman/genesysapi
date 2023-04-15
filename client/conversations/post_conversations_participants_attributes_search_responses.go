// Code generated by go-swagger; DO NOT EDIT.

package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostConversationsParticipantsAttributesSearchReader is a Reader for the PostConversationsParticipantsAttributesSearch structure.
type PostConversationsParticipantsAttributesSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationsParticipantsAttributesSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostConversationsParticipantsAttributesSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationsParticipantsAttributesSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationsParticipantsAttributesSearchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationsParticipantsAttributesSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationsParticipantsAttributesSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostConversationsParticipantsAttributesSearchRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationsParticipantsAttributesSearchRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationsParticipantsAttributesSearchUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationsParticipantsAttributesSearchTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationsParticipantsAttributesSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationsParticipantsAttributesSearchServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationsParticipantsAttributesSearchGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationsParticipantsAttributesSearchOK creates a PostConversationsParticipantsAttributesSearchOK with default headers values
func NewPostConversationsParticipantsAttributesSearchOK() *PostConversationsParticipantsAttributesSearchOK {
	return &PostConversationsParticipantsAttributesSearchOK{}
}

/*
PostConversationsParticipantsAttributesSearchOK describes a response with status code 200, with default header values.

successful operation
*/
type PostConversationsParticipantsAttributesSearchOK struct {
	Payload *models.JSONCursorSearchResponse
}

// IsSuccess returns true when this post conversations participants attributes search o k response has a 2xx status code
func (o *PostConversationsParticipantsAttributesSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post conversations participants attributes search o k response has a 3xx status code
func (o *PostConversationsParticipantsAttributesSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations participants attributes search o k response has a 4xx status code
func (o *PostConversationsParticipantsAttributesSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations participants attributes search o k response has a 5xx status code
func (o *PostConversationsParticipantsAttributesSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations participants attributes search o k response a status code equal to that given
func (o *PostConversationsParticipantsAttributesSearchOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostConversationsParticipantsAttributesSearchOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchOK  %+v", 200, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchOK) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchOK  %+v", 200, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchOK) GetPayload() *models.JSONCursorSearchResponse {
	return o.Payload
}

func (o *PostConversationsParticipantsAttributesSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONCursorSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsParticipantsAttributesSearchBadRequest creates a PostConversationsParticipantsAttributesSearchBadRequest with default headers values
func NewPostConversationsParticipantsAttributesSearchBadRequest() *PostConversationsParticipantsAttributesSearchBadRequest {
	return &PostConversationsParticipantsAttributesSearchBadRequest{}
}

/*
PostConversationsParticipantsAttributesSearchBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsParticipantsAttributesSearchBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations participants attributes search bad request response has a 2xx status code
func (o *PostConversationsParticipantsAttributesSearchBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations participants attributes search bad request response has a 3xx status code
func (o *PostConversationsParticipantsAttributesSearchBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations participants attributes search bad request response has a 4xx status code
func (o *PostConversationsParticipantsAttributesSearchBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations participants attributes search bad request response has a 5xx status code
func (o *PostConversationsParticipantsAttributesSearchBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations participants attributes search bad request response a status code equal to that given
func (o *PostConversationsParticipantsAttributesSearchBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostConversationsParticipantsAttributesSearchBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsParticipantsAttributesSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsParticipantsAttributesSearchUnauthorized creates a PostConversationsParticipantsAttributesSearchUnauthorized with default headers values
func NewPostConversationsParticipantsAttributesSearchUnauthorized() *PostConversationsParticipantsAttributesSearchUnauthorized {
	return &PostConversationsParticipantsAttributesSearchUnauthorized{}
}

/*
PostConversationsParticipantsAttributesSearchUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsParticipantsAttributesSearchUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations participants attributes search unauthorized response has a 2xx status code
func (o *PostConversationsParticipantsAttributesSearchUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations participants attributes search unauthorized response has a 3xx status code
func (o *PostConversationsParticipantsAttributesSearchUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations participants attributes search unauthorized response has a 4xx status code
func (o *PostConversationsParticipantsAttributesSearchUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations participants attributes search unauthorized response has a 5xx status code
func (o *PostConversationsParticipantsAttributesSearchUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations participants attributes search unauthorized response a status code equal to that given
func (o *PostConversationsParticipantsAttributesSearchUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostConversationsParticipantsAttributesSearchUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsParticipantsAttributesSearchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsParticipantsAttributesSearchForbidden creates a PostConversationsParticipantsAttributesSearchForbidden with default headers values
func NewPostConversationsParticipantsAttributesSearchForbidden() *PostConversationsParticipantsAttributesSearchForbidden {
	return &PostConversationsParticipantsAttributesSearchForbidden{}
}

/*
PostConversationsParticipantsAttributesSearchForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsParticipantsAttributesSearchForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations participants attributes search forbidden response has a 2xx status code
func (o *PostConversationsParticipantsAttributesSearchForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations participants attributes search forbidden response has a 3xx status code
func (o *PostConversationsParticipantsAttributesSearchForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations participants attributes search forbidden response has a 4xx status code
func (o *PostConversationsParticipantsAttributesSearchForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations participants attributes search forbidden response has a 5xx status code
func (o *PostConversationsParticipantsAttributesSearchForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations participants attributes search forbidden response a status code equal to that given
func (o *PostConversationsParticipantsAttributesSearchForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostConversationsParticipantsAttributesSearchForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsParticipantsAttributesSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsParticipantsAttributesSearchNotFound creates a PostConversationsParticipantsAttributesSearchNotFound with default headers values
func NewPostConversationsParticipantsAttributesSearchNotFound() *PostConversationsParticipantsAttributesSearchNotFound {
	return &PostConversationsParticipantsAttributesSearchNotFound{}
}

/*
PostConversationsParticipantsAttributesSearchNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostConversationsParticipantsAttributesSearchNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations participants attributes search not found response has a 2xx status code
func (o *PostConversationsParticipantsAttributesSearchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations participants attributes search not found response has a 3xx status code
func (o *PostConversationsParticipantsAttributesSearchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations participants attributes search not found response has a 4xx status code
func (o *PostConversationsParticipantsAttributesSearchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations participants attributes search not found response has a 5xx status code
func (o *PostConversationsParticipantsAttributesSearchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations participants attributes search not found response a status code equal to that given
func (o *PostConversationsParticipantsAttributesSearchNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostConversationsParticipantsAttributesSearchNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsParticipantsAttributesSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsParticipantsAttributesSearchRequestTimeout creates a PostConversationsParticipantsAttributesSearchRequestTimeout with default headers values
func NewPostConversationsParticipantsAttributesSearchRequestTimeout() *PostConversationsParticipantsAttributesSearchRequestTimeout {
	return &PostConversationsParticipantsAttributesSearchRequestTimeout{}
}

/*
PostConversationsParticipantsAttributesSearchRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostConversationsParticipantsAttributesSearchRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations participants attributes search request timeout response has a 2xx status code
func (o *PostConversationsParticipantsAttributesSearchRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations participants attributes search request timeout response has a 3xx status code
func (o *PostConversationsParticipantsAttributesSearchRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations participants attributes search request timeout response has a 4xx status code
func (o *PostConversationsParticipantsAttributesSearchRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations participants attributes search request timeout response has a 5xx status code
func (o *PostConversationsParticipantsAttributesSearchRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations participants attributes search request timeout response a status code equal to that given
func (o *PostConversationsParticipantsAttributesSearchRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostConversationsParticipantsAttributesSearchRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsParticipantsAttributesSearchRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsParticipantsAttributesSearchRequestEntityTooLarge creates a PostConversationsParticipantsAttributesSearchRequestEntityTooLarge with default headers values
func NewPostConversationsParticipantsAttributesSearchRequestEntityTooLarge() *PostConversationsParticipantsAttributesSearchRequestEntityTooLarge {
	return &PostConversationsParticipantsAttributesSearchRequestEntityTooLarge{}
}

/*
PostConversationsParticipantsAttributesSearchRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostConversationsParticipantsAttributesSearchRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations participants attributes search request entity too large response has a 2xx status code
func (o *PostConversationsParticipantsAttributesSearchRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations participants attributes search request entity too large response has a 3xx status code
func (o *PostConversationsParticipantsAttributesSearchRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations participants attributes search request entity too large response has a 4xx status code
func (o *PostConversationsParticipantsAttributesSearchRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations participants attributes search request entity too large response has a 5xx status code
func (o *PostConversationsParticipantsAttributesSearchRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations participants attributes search request entity too large response a status code equal to that given
func (o *PostConversationsParticipantsAttributesSearchRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostConversationsParticipantsAttributesSearchRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsParticipantsAttributesSearchRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsParticipantsAttributesSearchUnsupportedMediaType creates a PostConversationsParticipantsAttributesSearchUnsupportedMediaType with default headers values
func NewPostConversationsParticipantsAttributesSearchUnsupportedMediaType() *PostConversationsParticipantsAttributesSearchUnsupportedMediaType {
	return &PostConversationsParticipantsAttributesSearchUnsupportedMediaType{}
}

/*
PostConversationsParticipantsAttributesSearchUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsParticipantsAttributesSearchUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations participants attributes search unsupported media type response has a 2xx status code
func (o *PostConversationsParticipantsAttributesSearchUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations participants attributes search unsupported media type response has a 3xx status code
func (o *PostConversationsParticipantsAttributesSearchUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations participants attributes search unsupported media type response has a 4xx status code
func (o *PostConversationsParticipantsAttributesSearchUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations participants attributes search unsupported media type response has a 5xx status code
func (o *PostConversationsParticipantsAttributesSearchUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations participants attributes search unsupported media type response a status code equal to that given
func (o *PostConversationsParticipantsAttributesSearchUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostConversationsParticipantsAttributesSearchUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsParticipantsAttributesSearchUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsParticipantsAttributesSearchTooManyRequests creates a PostConversationsParticipantsAttributesSearchTooManyRequests with default headers values
func NewPostConversationsParticipantsAttributesSearchTooManyRequests() *PostConversationsParticipantsAttributesSearchTooManyRequests {
	return &PostConversationsParticipantsAttributesSearchTooManyRequests{}
}

/*
PostConversationsParticipantsAttributesSearchTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostConversationsParticipantsAttributesSearchTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations participants attributes search too many requests response has a 2xx status code
func (o *PostConversationsParticipantsAttributesSearchTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations participants attributes search too many requests response has a 3xx status code
func (o *PostConversationsParticipantsAttributesSearchTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations participants attributes search too many requests response has a 4xx status code
func (o *PostConversationsParticipantsAttributesSearchTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations participants attributes search too many requests response has a 5xx status code
func (o *PostConversationsParticipantsAttributesSearchTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations participants attributes search too many requests response a status code equal to that given
func (o *PostConversationsParticipantsAttributesSearchTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostConversationsParticipantsAttributesSearchTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsParticipantsAttributesSearchTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsParticipantsAttributesSearchInternalServerError creates a PostConversationsParticipantsAttributesSearchInternalServerError with default headers values
func NewPostConversationsParticipantsAttributesSearchInternalServerError() *PostConversationsParticipantsAttributesSearchInternalServerError {
	return &PostConversationsParticipantsAttributesSearchInternalServerError{}
}

/*
PostConversationsParticipantsAttributesSearchInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsParticipantsAttributesSearchInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations participants attributes search internal server error response has a 2xx status code
func (o *PostConversationsParticipantsAttributesSearchInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations participants attributes search internal server error response has a 3xx status code
func (o *PostConversationsParticipantsAttributesSearchInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations participants attributes search internal server error response has a 4xx status code
func (o *PostConversationsParticipantsAttributesSearchInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations participants attributes search internal server error response has a 5xx status code
func (o *PostConversationsParticipantsAttributesSearchInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversations participants attributes search internal server error response a status code equal to that given
func (o *PostConversationsParticipantsAttributesSearchInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostConversationsParticipantsAttributesSearchInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsParticipantsAttributesSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsParticipantsAttributesSearchServiceUnavailable creates a PostConversationsParticipantsAttributesSearchServiceUnavailable with default headers values
func NewPostConversationsParticipantsAttributesSearchServiceUnavailable() *PostConversationsParticipantsAttributesSearchServiceUnavailable {
	return &PostConversationsParticipantsAttributesSearchServiceUnavailable{}
}

/*
PostConversationsParticipantsAttributesSearchServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsParticipantsAttributesSearchServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations participants attributes search service unavailable response has a 2xx status code
func (o *PostConversationsParticipantsAttributesSearchServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations participants attributes search service unavailable response has a 3xx status code
func (o *PostConversationsParticipantsAttributesSearchServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations participants attributes search service unavailable response has a 4xx status code
func (o *PostConversationsParticipantsAttributesSearchServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations participants attributes search service unavailable response has a 5xx status code
func (o *PostConversationsParticipantsAttributesSearchServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversations participants attributes search service unavailable response a status code equal to that given
func (o *PostConversationsParticipantsAttributesSearchServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostConversationsParticipantsAttributesSearchServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsParticipantsAttributesSearchServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsParticipantsAttributesSearchGatewayTimeout creates a PostConversationsParticipantsAttributesSearchGatewayTimeout with default headers values
func NewPostConversationsParticipantsAttributesSearchGatewayTimeout() *PostConversationsParticipantsAttributesSearchGatewayTimeout {
	return &PostConversationsParticipantsAttributesSearchGatewayTimeout{}
}

/*
PostConversationsParticipantsAttributesSearchGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostConversationsParticipantsAttributesSearchGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations participants attributes search gateway timeout response has a 2xx status code
func (o *PostConversationsParticipantsAttributesSearchGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations participants attributes search gateway timeout response has a 3xx status code
func (o *PostConversationsParticipantsAttributesSearchGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations participants attributes search gateway timeout response has a 4xx status code
func (o *PostConversationsParticipantsAttributesSearchGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations participants attributes search gateway timeout response has a 5xx status code
func (o *PostConversationsParticipantsAttributesSearchGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversations participants attributes search gateway timeout response a status code equal to that given
func (o *PostConversationsParticipantsAttributesSearchGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostConversationsParticipantsAttributesSearchGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/participants/attributes/search][%d] postConversationsParticipantsAttributesSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsParticipantsAttributesSearchGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsParticipantsAttributesSearchGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
