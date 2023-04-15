// Code generated by go-swagger; DO NOT EDIT.

package response_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostResponsemanagementResponseassetsSearchReader is a Reader for the PostResponsemanagementResponseassetsSearch structure.
type PostResponsemanagementResponseassetsSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostResponsemanagementResponseassetsSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostResponsemanagementResponseassetsSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostResponsemanagementResponseassetsSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostResponsemanagementResponseassetsSearchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostResponsemanagementResponseassetsSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostResponsemanagementResponseassetsSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostResponsemanagementResponseassetsSearchRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostResponsemanagementResponseassetsSearchRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostResponsemanagementResponseassetsSearchUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostResponsemanagementResponseassetsSearchTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostResponsemanagementResponseassetsSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostResponsemanagementResponseassetsSearchServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostResponsemanagementResponseassetsSearchGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostResponsemanagementResponseassetsSearchOK creates a PostResponsemanagementResponseassetsSearchOK with default headers values
func NewPostResponsemanagementResponseassetsSearchOK() *PostResponsemanagementResponseassetsSearchOK {
	return &PostResponsemanagementResponseassetsSearchOK{}
}

/*
PostResponsemanagementResponseassetsSearchOK describes a response with status code 200, with default header values.

successful operation
*/
type PostResponsemanagementResponseassetsSearchOK struct {
	Payload *models.ResponseAssetSearchResults
}

// IsSuccess returns true when this post responsemanagement responseassets search o k response has a 2xx status code
func (o *PostResponsemanagementResponseassetsSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post responsemanagement responseassets search o k response has a 3xx status code
func (o *PostResponsemanagementResponseassetsSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responseassets search o k response has a 4xx status code
func (o *PostResponsemanagementResponseassetsSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post responsemanagement responseassets search o k response has a 5xx status code
func (o *PostResponsemanagementResponseassetsSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post responsemanagement responseassets search o k response a status code equal to that given
func (o *PostResponsemanagementResponseassetsSearchOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostResponsemanagementResponseassetsSearchOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchOK  %+v", 200, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchOK) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchOK  %+v", 200, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchOK) GetPayload() *models.ResponseAssetSearchResults {
	return o.Payload
}

func (o *PostResponsemanagementResponseassetsSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseAssetSearchResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementResponseassetsSearchBadRequest creates a PostResponsemanagementResponseassetsSearchBadRequest with default headers values
func NewPostResponsemanagementResponseassetsSearchBadRequest() *PostResponsemanagementResponseassetsSearchBadRequest {
	return &PostResponsemanagementResponseassetsSearchBadRequest{}
}

/*
PostResponsemanagementResponseassetsSearchBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostResponsemanagementResponseassetsSearchBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post responsemanagement responseassets search bad request response has a 2xx status code
func (o *PostResponsemanagementResponseassetsSearchBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post responsemanagement responseassets search bad request response has a 3xx status code
func (o *PostResponsemanagementResponseassetsSearchBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responseassets search bad request response has a 4xx status code
func (o *PostResponsemanagementResponseassetsSearchBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post responsemanagement responseassets search bad request response has a 5xx status code
func (o *PostResponsemanagementResponseassetsSearchBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post responsemanagement responseassets search bad request response a status code equal to that given
func (o *PostResponsemanagementResponseassetsSearchBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostResponsemanagementResponseassetsSearchBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchBadRequest  %+v", 400, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchBadRequest  %+v", 400, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementResponseassetsSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementResponseassetsSearchUnauthorized creates a PostResponsemanagementResponseassetsSearchUnauthorized with default headers values
func NewPostResponsemanagementResponseassetsSearchUnauthorized() *PostResponsemanagementResponseassetsSearchUnauthorized {
	return &PostResponsemanagementResponseassetsSearchUnauthorized{}
}

/*
PostResponsemanagementResponseassetsSearchUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostResponsemanagementResponseassetsSearchUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post responsemanagement responseassets search unauthorized response has a 2xx status code
func (o *PostResponsemanagementResponseassetsSearchUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post responsemanagement responseassets search unauthorized response has a 3xx status code
func (o *PostResponsemanagementResponseassetsSearchUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responseassets search unauthorized response has a 4xx status code
func (o *PostResponsemanagementResponseassetsSearchUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post responsemanagement responseassets search unauthorized response has a 5xx status code
func (o *PostResponsemanagementResponseassetsSearchUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post responsemanagement responseassets search unauthorized response a status code equal to that given
func (o *PostResponsemanagementResponseassetsSearchUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostResponsemanagementResponseassetsSearchUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementResponseassetsSearchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementResponseassetsSearchForbidden creates a PostResponsemanagementResponseassetsSearchForbidden with default headers values
func NewPostResponsemanagementResponseassetsSearchForbidden() *PostResponsemanagementResponseassetsSearchForbidden {
	return &PostResponsemanagementResponseassetsSearchForbidden{}
}

/*
PostResponsemanagementResponseassetsSearchForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostResponsemanagementResponseassetsSearchForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post responsemanagement responseassets search forbidden response has a 2xx status code
func (o *PostResponsemanagementResponseassetsSearchForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post responsemanagement responseassets search forbidden response has a 3xx status code
func (o *PostResponsemanagementResponseassetsSearchForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responseassets search forbidden response has a 4xx status code
func (o *PostResponsemanagementResponseassetsSearchForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post responsemanagement responseassets search forbidden response has a 5xx status code
func (o *PostResponsemanagementResponseassetsSearchForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post responsemanagement responseassets search forbidden response a status code equal to that given
func (o *PostResponsemanagementResponseassetsSearchForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostResponsemanagementResponseassetsSearchForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchForbidden  %+v", 403, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchForbidden  %+v", 403, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementResponseassetsSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementResponseassetsSearchNotFound creates a PostResponsemanagementResponseassetsSearchNotFound with default headers values
func NewPostResponsemanagementResponseassetsSearchNotFound() *PostResponsemanagementResponseassetsSearchNotFound {
	return &PostResponsemanagementResponseassetsSearchNotFound{}
}

/*
PostResponsemanagementResponseassetsSearchNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostResponsemanagementResponseassetsSearchNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post responsemanagement responseassets search not found response has a 2xx status code
func (o *PostResponsemanagementResponseassetsSearchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post responsemanagement responseassets search not found response has a 3xx status code
func (o *PostResponsemanagementResponseassetsSearchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responseassets search not found response has a 4xx status code
func (o *PostResponsemanagementResponseassetsSearchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post responsemanagement responseassets search not found response has a 5xx status code
func (o *PostResponsemanagementResponseassetsSearchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post responsemanagement responseassets search not found response a status code equal to that given
func (o *PostResponsemanagementResponseassetsSearchNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostResponsemanagementResponseassetsSearchNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchNotFound  %+v", 404, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchNotFound  %+v", 404, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementResponseassetsSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementResponseassetsSearchRequestTimeout creates a PostResponsemanagementResponseassetsSearchRequestTimeout with default headers values
func NewPostResponsemanagementResponseassetsSearchRequestTimeout() *PostResponsemanagementResponseassetsSearchRequestTimeout {
	return &PostResponsemanagementResponseassetsSearchRequestTimeout{}
}

/*
PostResponsemanagementResponseassetsSearchRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostResponsemanagementResponseassetsSearchRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post responsemanagement responseassets search request timeout response has a 2xx status code
func (o *PostResponsemanagementResponseassetsSearchRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post responsemanagement responseassets search request timeout response has a 3xx status code
func (o *PostResponsemanagementResponseassetsSearchRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responseassets search request timeout response has a 4xx status code
func (o *PostResponsemanagementResponseassetsSearchRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post responsemanagement responseassets search request timeout response has a 5xx status code
func (o *PostResponsemanagementResponseassetsSearchRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post responsemanagement responseassets search request timeout response a status code equal to that given
func (o *PostResponsemanagementResponseassetsSearchRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostResponsemanagementResponseassetsSearchRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementResponseassetsSearchRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementResponseassetsSearchRequestEntityTooLarge creates a PostResponsemanagementResponseassetsSearchRequestEntityTooLarge with default headers values
func NewPostResponsemanagementResponseassetsSearchRequestEntityTooLarge() *PostResponsemanagementResponseassetsSearchRequestEntityTooLarge {
	return &PostResponsemanagementResponseassetsSearchRequestEntityTooLarge{}
}

/*
PostResponsemanagementResponseassetsSearchRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostResponsemanagementResponseassetsSearchRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post responsemanagement responseassets search request entity too large response has a 2xx status code
func (o *PostResponsemanagementResponseassetsSearchRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post responsemanagement responseassets search request entity too large response has a 3xx status code
func (o *PostResponsemanagementResponseassetsSearchRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responseassets search request entity too large response has a 4xx status code
func (o *PostResponsemanagementResponseassetsSearchRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post responsemanagement responseassets search request entity too large response has a 5xx status code
func (o *PostResponsemanagementResponseassetsSearchRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post responsemanagement responseassets search request entity too large response a status code equal to that given
func (o *PostResponsemanagementResponseassetsSearchRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostResponsemanagementResponseassetsSearchRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementResponseassetsSearchRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementResponseassetsSearchUnsupportedMediaType creates a PostResponsemanagementResponseassetsSearchUnsupportedMediaType with default headers values
func NewPostResponsemanagementResponseassetsSearchUnsupportedMediaType() *PostResponsemanagementResponseassetsSearchUnsupportedMediaType {
	return &PostResponsemanagementResponseassetsSearchUnsupportedMediaType{}
}

/*
PostResponsemanagementResponseassetsSearchUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostResponsemanagementResponseassetsSearchUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post responsemanagement responseassets search unsupported media type response has a 2xx status code
func (o *PostResponsemanagementResponseassetsSearchUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post responsemanagement responseassets search unsupported media type response has a 3xx status code
func (o *PostResponsemanagementResponseassetsSearchUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responseassets search unsupported media type response has a 4xx status code
func (o *PostResponsemanagementResponseassetsSearchUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post responsemanagement responseassets search unsupported media type response has a 5xx status code
func (o *PostResponsemanagementResponseassetsSearchUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post responsemanagement responseassets search unsupported media type response a status code equal to that given
func (o *PostResponsemanagementResponseassetsSearchUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostResponsemanagementResponseassetsSearchUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementResponseassetsSearchUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementResponseassetsSearchTooManyRequests creates a PostResponsemanagementResponseassetsSearchTooManyRequests with default headers values
func NewPostResponsemanagementResponseassetsSearchTooManyRequests() *PostResponsemanagementResponseassetsSearchTooManyRequests {
	return &PostResponsemanagementResponseassetsSearchTooManyRequests{}
}

/*
PostResponsemanagementResponseassetsSearchTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostResponsemanagementResponseassetsSearchTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post responsemanagement responseassets search too many requests response has a 2xx status code
func (o *PostResponsemanagementResponseassetsSearchTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post responsemanagement responseassets search too many requests response has a 3xx status code
func (o *PostResponsemanagementResponseassetsSearchTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responseassets search too many requests response has a 4xx status code
func (o *PostResponsemanagementResponseassetsSearchTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post responsemanagement responseassets search too many requests response has a 5xx status code
func (o *PostResponsemanagementResponseassetsSearchTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post responsemanagement responseassets search too many requests response a status code equal to that given
func (o *PostResponsemanagementResponseassetsSearchTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostResponsemanagementResponseassetsSearchTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementResponseassetsSearchTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementResponseassetsSearchInternalServerError creates a PostResponsemanagementResponseassetsSearchInternalServerError with default headers values
func NewPostResponsemanagementResponseassetsSearchInternalServerError() *PostResponsemanagementResponseassetsSearchInternalServerError {
	return &PostResponsemanagementResponseassetsSearchInternalServerError{}
}

/*
PostResponsemanagementResponseassetsSearchInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostResponsemanagementResponseassetsSearchInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post responsemanagement responseassets search internal server error response has a 2xx status code
func (o *PostResponsemanagementResponseassetsSearchInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post responsemanagement responseassets search internal server error response has a 3xx status code
func (o *PostResponsemanagementResponseassetsSearchInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responseassets search internal server error response has a 4xx status code
func (o *PostResponsemanagementResponseassetsSearchInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post responsemanagement responseassets search internal server error response has a 5xx status code
func (o *PostResponsemanagementResponseassetsSearchInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post responsemanagement responseassets search internal server error response a status code equal to that given
func (o *PostResponsemanagementResponseassetsSearchInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostResponsemanagementResponseassetsSearchInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementResponseassetsSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementResponseassetsSearchServiceUnavailable creates a PostResponsemanagementResponseassetsSearchServiceUnavailable with default headers values
func NewPostResponsemanagementResponseassetsSearchServiceUnavailable() *PostResponsemanagementResponseassetsSearchServiceUnavailable {
	return &PostResponsemanagementResponseassetsSearchServiceUnavailable{}
}

/*
PostResponsemanagementResponseassetsSearchServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostResponsemanagementResponseassetsSearchServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post responsemanagement responseassets search service unavailable response has a 2xx status code
func (o *PostResponsemanagementResponseassetsSearchServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post responsemanagement responseassets search service unavailable response has a 3xx status code
func (o *PostResponsemanagementResponseassetsSearchServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responseassets search service unavailable response has a 4xx status code
func (o *PostResponsemanagementResponseassetsSearchServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post responsemanagement responseassets search service unavailable response has a 5xx status code
func (o *PostResponsemanagementResponseassetsSearchServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post responsemanagement responseassets search service unavailable response a status code equal to that given
func (o *PostResponsemanagementResponseassetsSearchServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostResponsemanagementResponseassetsSearchServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementResponseassetsSearchServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementResponseassetsSearchGatewayTimeout creates a PostResponsemanagementResponseassetsSearchGatewayTimeout with default headers values
func NewPostResponsemanagementResponseassetsSearchGatewayTimeout() *PostResponsemanagementResponseassetsSearchGatewayTimeout {
	return &PostResponsemanagementResponseassetsSearchGatewayTimeout{}
}

/*
PostResponsemanagementResponseassetsSearchGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostResponsemanagementResponseassetsSearchGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post responsemanagement responseassets search gateway timeout response has a 2xx status code
func (o *PostResponsemanagementResponseassetsSearchGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post responsemanagement responseassets search gateway timeout response has a 3xx status code
func (o *PostResponsemanagementResponseassetsSearchGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responseassets search gateway timeout response has a 4xx status code
func (o *PostResponsemanagementResponseassetsSearchGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post responsemanagement responseassets search gateway timeout response has a 5xx status code
func (o *PostResponsemanagementResponseassetsSearchGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post responsemanagement responseassets search gateway timeout response a status code equal to that given
func (o *PostResponsemanagementResponseassetsSearchGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostResponsemanagementResponseassetsSearchGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responseassets/search][%d] postResponsemanagementResponseassetsSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostResponsemanagementResponseassetsSearchGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementResponseassetsSearchGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
