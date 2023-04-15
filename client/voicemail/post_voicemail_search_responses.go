// Code generated by go-swagger; DO NOT EDIT.

package voicemail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostVoicemailSearchReader is a Reader for the PostVoicemailSearch structure.
type PostVoicemailSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVoicemailSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostVoicemailSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostVoicemailSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostVoicemailSearchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostVoicemailSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostVoicemailSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostVoicemailSearchRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostVoicemailSearchRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostVoicemailSearchUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostVoicemailSearchTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostVoicemailSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostVoicemailSearchServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostVoicemailSearchGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostVoicemailSearchOK creates a PostVoicemailSearchOK with default headers values
func NewPostVoicemailSearchOK() *PostVoicemailSearchOK {
	return &PostVoicemailSearchOK{}
}

/*
PostVoicemailSearchOK describes a response with status code 200, with default header values.

successful operation
*/
type PostVoicemailSearchOK struct {
	Payload *models.VoicemailsSearchResponse
}

// IsSuccess returns true when this post voicemail search o k response has a 2xx status code
func (o *PostVoicemailSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post voicemail search o k response has a 3xx status code
func (o *PostVoicemailSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post voicemail search o k response has a 4xx status code
func (o *PostVoicemailSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post voicemail search o k response has a 5xx status code
func (o *PostVoicemailSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post voicemail search o k response a status code equal to that given
func (o *PostVoicemailSearchOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostVoicemailSearchOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchOK  %+v", 200, o.Payload)
}

func (o *PostVoicemailSearchOK) String() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchOK  %+v", 200, o.Payload)
}

func (o *PostVoicemailSearchOK) GetPayload() *models.VoicemailsSearchResponse {
	return o.Payload
}

func (o *PostVoicemailSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VoicemailsSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVoicemailSearchBadRequest creates a PostVoicemailSearchBadRequest with default headers values
func NewPostVoicemailSearchBadRequest() *PostVoicemailSearchBadRequest {
	return &PostVoicemailSearchBadRequest{}
}

/*
PostVoicemailSearchBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostVoicemailSearchBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post voicemail search bad request response has a 2xx status code
func (o *PostVoicemailSearchBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post voicemail search bad request response has a 3xx status code
func (o *PostVoicemailSearchBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post voicemail search bad request response has a 4xx status code
func (o *PostVoicemailSearchBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post voicemail search bad request response has a 5xx status code
func (o *PostVoicemailSearchBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post voicemail search bad request response a status code equal to that given
func (o *PostVoicemailSearchBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostVoicemailSearchBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchBadRequest  %+v", 400, o.Payload)
}

func (o *PostVoicemailSearchBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchBadRequest  %+v", 400, o.Payload)
}

func (o *PostVoicemailSearchBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostVoicemailSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVoicemailSearchUnauthorized creates a PostVoicemailSearchUnauthorized with default headers values
func NewPostVoicemailSearchUnauthorized() *PostVoicemailSearchUnauthorized {
	return &PostVoicemailSearchUnauthorized{}
}

/*
PostVoicemailSearchUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostVoicemailSearchUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post voicemail search unauthorized response has a 2xx status code
func (o *PostVoicemailSearchUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post voicemail search unauthorized response has a 3xx status code
func (o *PostVoicemailSearchUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post voicemail search unauthorized response has a 4xx status code
func (o *PostVoicemailSearchUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post voicemail search unauthorized response has a 5xx status code
func (o *PostVoicemailSearchUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post voicemail search unauthorized response a status code equal to that given
func (o *PostVoicemailSearchUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostVoicemailSearchUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *PostVoicemailSearchUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *PostVoicemailSearchUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostVoicemailSearchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVoicemailSearchForbidden creates a PostVoicemailSearchForbidden with default headers values
func NewPostVoicemailSearchForbidden() *PostVoicemailSearchForbidden {
	return &PostVoicemailSearchForbidden{}
}

/*
PostVoicemailSearchForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostVoicemailSearchForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post voicemail search forbidden response has a 2xx status code
func (o *PostVoicemailSearchForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post voicemail search forbidden response has a 3xx status code
func (o *PostVoicemailSearchForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post voicemail search forbidden response has a 4xx status code
func (o *PostVoicemailSearchForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post voicemail search forbidden response has a 5xx status code
func (o *PostVoicemailSearchForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post voicemail search forbidden response a status code equal to that given
func (o *PostVoicemailSearchForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostVoicemailSearchForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchForbidden  %+v", 403, o.Payload)
}

func (o *PostVoicemailSearchForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchForbidden  %+v", 403, o.Payload)
}

func (o *PostVoicemailSearchForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostVoicemailSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVoicemailSearchNotFound creates a PostVoicemailSearchNotFound with default headers values
func NewPostVoicemailSearchNotFound() *PostVoicemailSearchNotFound {
	return &PostVoicemailSearchNotFound{}
}

/*
PostVoicemailSearchNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostVoicemailSearchNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post voicemail search not found response has a 2xx status code
func (o *PostVoicemailSearchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post voicemail search not found response has a 3xx status code
func (o *PostVoicemailSearchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post voicemail search not found response has a 4xx status code
func (o *PostVoicemailSearchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post voicemail search not found response has a 5xx status code
func (o *PostVoicemailSearchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post voicemail search not found response a status code equal to that given
func (o *PostVoicemailSearchNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostVoicemailSearchNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchNotFound  %+v", 404, o.Payload)
}

func (o *PostVoicemailSearchNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchNotFound  %+v", 404, o.Payload)
}

func (o *PostVoicemailSearchNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostVoicemailSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVoicemailSearchRequestTimeout creates a PostVoicemailSearchRequestTimeout with default headers values
func NewPostVoicemailSearchRequestTimeout() *PostVoicemailSearchRequestTimeout {
	return &PostVoicemailSearchRequestTimeout{}
}

/*
PostVoicemailSearchRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostVoicemailSearchRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post voicemail search request timeout response has a 2xx status code
func (o *PostVoicemailSearchRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post voicemail search request timeout response has a 3xx status code
func (o *PostVoicemailSearchRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post voicemail search request timeout response has a 4xx status code
func (o *PostVoicemailSearchRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post voicemail search request timeout response has a 5xx status code
func (o *PostVoicemailSearchRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post voicemail search request timeout response a status code equal to that given
func (o *PostVoicemailSearchRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostVoicemailSearchRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostVoicemailSearchRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostVoicemailSearchRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostVoicemailSearchRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVoicemailSearchRequestEntityTooLarge creates a PostVoicemailSearchRequestEntityTooLarge with default headers values
func NewPostVoicemailSearchRequestEntityTooLarge() *PostVoicemailSearchRequestEntityTooLarge {
	return &PostVoicemailSearchRequestEntityTooLarge{}
}

/*
PostVoicemailSearchRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostVoicemailSearchRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post voicemail search request entity too large response has a 2xx status code
func (o *PostVoicemailSearchRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post voicemail search request entity too large response has a 3xx status code
func (o *PostVoicemailSearchRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post voicemail search request entity too large response has a 4xx status code
func (o *PostVoicemailSearchRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post voicemail search request entity too large response has a 5xx status code
func (o *PostVoicemailSearchRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post voicemail search request entity too large response a status code equal to that given
func (o *PostVoicemailSearchRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostVoicemailSearchRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostVoicemailSearchRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostVoicemailSearchRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostVoicemailSearchRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVoicemailSearchUnsupportedMediaType creates a PostVoicemailSearchUnsupportedMediaType with default headers values
func NewPostVoicemailSearchUnsupportedMediaType() *PostVoicemailSearchUnsupportedMediaType {
	return &PostVoicemailSearchUnsupportedMediaType{}
}

/*
PostVoicemailSearchUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostVoicemailSearchUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post voicemail search unsupported media type response has a 2xx status code
func (o *PostVoicemailSearchUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post voicemail search unsupported media type response has a 3xx status code
func (o *PostVoicemailSearchUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post voicemail search unsupported media type response has a 4xx status code
func (o *PostVoicemailSearchUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post voicemail search unsupported media type response has a 5xx status code
func (o *PostVoicemailSearchUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post voicemail search unsupported media type response a status code equal to that given
func (o *PostVoicemailSearchUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostVoicemailSearchUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostVoicemailSearchUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostVoicemailSearchUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostVoicemailSearchUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVoicemailSearchTooManyRequests creates a PostVoicemailSearchTooManyRequests with default headers values
func NewPostVoicemailSearchTooManyRequests() *PostVoicemailSearchTooManyRequests {
	return &PostVoicemailSearchTooManyRequests{}
}

/*
PostVoicemailSearchTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostVoicemailSearchTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post voicemail search too many requests response has a 2xx status code
func (o *PostVoicemailSearchTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post voicemail search too many requests response has a 3xx status code
func (o *PostVoicemailSearchTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post voicemail search too many requests response has a 4xx status code
func (o *PostVoicemailSearchTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post voicemail search too many requests response has a 5xx status code
func (o *PostVoicemailSearchTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post voicemail search too many requests response a status code equal to that given
func (o *PostVoicemailSearchTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostVoicemailSearchTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostVoicemailSearchTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostVoicemailSearchTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostVoicemailSearchTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVoicemailSearchInternalServerError creates a PostVoicemailSearchInternalServerError with default headers values
func NewPostVoicemailSearchInternalServerError() *PostVoicemailSearchInternalServerError {
	return &PostVoicemailSearchInternalServerError{}
}

/*
PostVoicemailSearchInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostVoicemailSearchInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post voicemail search internal server error response has a 2xx status code
func (o *PostVoicemailSearchInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post voicemail search internal server error response has a 3xx status code
func (o *PostVoicemailSearchInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post voicemail search internal server error response has a 4xx status code
func (o *PostVoicemailSearchInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post voicemail search internal server error response has a 5xx status code
func (o *PostVoicemailSearchInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post voicemail search internal server error response a status code equal to that given
func (o *PostVoicemailSearchInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostVoicemailSearchInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *PostVoicemailSearchInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *PostVoicemailSearchInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostVoicemailSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVoicemailSearchServiceUnavailable creates a PostVoicemailSearchServiceUnavailable with default headers values
func NewPostVoicemailSearchServiceUnavailable() *PostVoicemailSearchServiceUnavailable {
	return &PostVoicemailSearchServiceUnavailable{}
}

/*
PostVoicemailSearchServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostVoicemailSearchServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post voicemail search service unavailable response has a 2xx status code
func (o *PostVoicemailSearchServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post voicemail search service unavailable response has a 3xx status code
func (o *PostVoicemailSearchServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post voicemail search service unavailable response has a 4xx status code
func (o *PostVoicemailSearchServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post voicemail search service unavailable response has a 5xx status code
func (o *PostVoicemailSearchServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post voicemail search service unavailable response a status code equal to that given
func (o *PostVoicemailSearchServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostVoicemailSearchServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostVoicemailSearchServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostVoicemailSearchServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostVoicemailSearchServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVoicemailSearchGatewayTimeout creates a PostVoicemailSearchGatewayTimeout with default headers values
func NewPostVoicemailSearchGatewayTimeout() *PostVoicemailSearchGatewayTimeout {
	return &PostVoicemailSearchGatewayTimeout{}
}

/*
PostVoicemailSearchGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostVoicemailSearchGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post voicemail search gateway timeout response has a 2xx status code
func (o *PostVoicemailSearchGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post voicemail search gateway timeout response has a 3xx status code
func (o *PostVoicemailSearchGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post voicemail search gateway timeout response has a 4xx status code
func (o *PostVoicemailSearchGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post voicemail search gateway timeout response has a 5xx status code
func (o *PostVoicemailSearchGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post voicemail search gateway timeout response a status code equal to that given
func (o *PostVoicemailSearchGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostVoicemailSearchGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostVoicemailSearchGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/search][%d] postVoicemailSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostVoicemailSearchGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostVoicemailSearchGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
