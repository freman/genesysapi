// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostOutboundContactlistClearReader is a Reader for the PostOutboundContactlistClear structure.
type PostOutboundContactlistClearReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOutboundContactlistClearReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostOutboundContactlistClearNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOutboundContactlistClearBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostOutboundContactlistClearUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostOutboundContactlistClearForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostOutboundContactlistClearNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostOutboundContactlistClearRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostOutboundContactlistClearRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostOutboundContactlistClearUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostOutboundContactlistClearTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOutboundContactlistClearInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostOutboundContactlistClearServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostOutboundContactlistClearGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOutboundContactlistClearNoContent creates a PostOutboundContactlistClearNoContent with default headers values
func NewPostOutboundContactlistClearNoContent() *PostOutboundContactlistClearNoContent {
	return &PostOutboundContactlistClearNoContent{}
}

/*
PostOutboundContactlistClearNoContent describes a response with status code 204, with default header values.

Contacts will be deleted.
*/
type PostOutboundContactlistClearNoContent struct {
}

// IsSuccess returns true when this post outbound contactlist clear no content response has a 2xx status code
func (o *PostOutboundContactlistClearNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post outbound contactlist clear no content response has a 3xx status code
func (o *PostOutboundContactlistClearNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist clear no content response has a 4xx status code
func (o *PostOutboundContactlistClearNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound contactlist clear no content response has a 5xx status code
func (o *PostOutboundContactlistClearNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound contactlist clear no content response a status code equal to that given
func (o *PostOutboundContactlistClearNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *PostOutboundContactlistClearNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearNoContent ", 204)
}

func (o *PostOutboundContactlistClearNoContent) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearNoContent ", 204)
}

func (o *PostOutboundContactlistClearNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostOutboundContactlistClearBadRequest creates a PostOutboundContactlistClearBadRequest with default headers values
func NewPostOutboundContactlistClearBadRequest() *PostOutboundContactlistClearBadRequest {
	return &PostOutboundContactlistClearBadRequest{}
}

/*
PostOutboundContactlistClearBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostOutboundContactlistClearBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound contactlist clear bad request response has a 2xx status code
func (o *PostOutboundContactlistClearBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound contactlist clear bad request response has a 3xx status code
func (o *PostOutboundContactlistClearBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist clear bad request response has a 4xx status code
func (o *PostOutboundContactlistClearBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound contactlist clear bad request response has a 5xx status code
func (o *PostOutboundContactlistClearBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound contactlist clear bad request response a status code equal to that given
func (o *PostOutboundContactlistClearBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostOutboundContactlistClearBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearBadRequest  %+v", 400, o.Payload)
}

func (o *PostOutboundContactlistClearBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearBadRequest  %+v", 400, o.Payload)
}

func (o *PostOutboundContactlistClearBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistClearBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistClearUnauthorized creates a PostOutboundContactlistClearUnauthorized with default headers values
func NewPostOutboundContactlistClearUnauthorized() *PostOutboundContactlistClearUnauthorized {
	return &PostOutboundContactlistClearUnauthorized{}
}

/*
PostOutboundContactlistClearUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostOutboundContactlistClearUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound contactlist clear unauthorized response has a 2xx status code
func (o *PostOutboundContactlistClearUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound contactlist clear unauthorized response has a 3xx status code
func (o *PostOutboundContactlistClearUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist clear unauthorized response has a 4xx status code
func (o *PostOutboundContactlistClearUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound contactlist clear unauthorized response has a 5xx status code
func (o *PostOutboundContactlistClearUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound contactlist clear unauthorized response a status code equal to that given
func (o *PostOutboundContactlistClearUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostOutboundContactlistClearUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOutboundContactlistClearUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOutboundContactlistClearUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistClearUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistClearForbidden creates a PostOutboundContactlistClearForbidden with default headers values
func NewPostOutboundContactlistClearForbidden() *PostOutboundContactlistClearForbidden {
	return &PostOutboundContactlistClearForbidden{}
}

/*
PostOutboundContactlistClearForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostOutboundContactlistClearForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound contactlist clear forbidden response has a 2xx status code
func (o *PostOutboundContactlistClearForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound contactlist clear forbidden response has a 3xx status code
func (o *PostOutboundContactlistClearForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist clear forbidden response has a 4xx status code
func (o *PostOutboundContactlistClearForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound contactlist clear forbidden response has a 5xx status code
func (o *PostOutboundContactlistClearForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound contactlist clear forbidden response a status code equal to that given
func (o *PostOutboundContactlistClearForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostOutboundContactlistClearForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearForbidden  %+v", 403, o.Payload)
}

func (o *PostOutboundContactlistClearForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearForbidden  %+v", 403, o.Payload)
}

func (o *PostOutboundContactlistClearForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistClearForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistClearNotFound creates a PostOutboundContactlistClearNotFound with default headers values
func NewPostOutboundContactlistClearNotFound() *PostOutboundContactlistClearNotFound {
	return &PostOutboundContactlistClearNotFound{}
}

/*
PostOutboundContactlistClearNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostOutboundContactlistClearNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound contactlist clear not found response has a 2xx status code
func (o *PostOutboundContactlistClearNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound contactlist clear not found response has a 3xx status code
func (o *PostOutboundContactlistClearNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist clear not found response has a 4xx status code
func (o *PostOutboundContactlistClearNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound contactlist clear not found response has a 5xx status code
func (o *PostOutboundContactlistClearNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound contactlist clear not found response a status code equal to that given
func (o *PostOutboundContactlistClearNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostOutboundContactlistClearNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearNotFound  %+v", 404, o.Payload)
}

func (o *PostOutboundContactlistClearNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearNotFound  %+v", 404, o.Payload)
}

func (o *PostOutboundContactlistClearNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistClearNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistClearRequestTimeout creates a PostOutboundContactlistClearRequestTimeout with default headers values
func NewPostOutboundContactlistClearRequestTimeout() *PostOutboundContactlistClearRequestTimeout {
	return &PostOutboundContactlistClearRequestTimeout{}
}

/*
PostOutboundContactlistClearRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostOutboundContactlistClearRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound contactlist clear request timeout response has a 2xx status code
func (o *PostOutboundContactlistClearRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound contactlist clear request timeout response has a 3xx status code
func (o *PostOutboundContactlistClearRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist clear request timeout response has a 4xx status code
func (o *PostOutboundContactlistClearRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound contactlist clear request timeout response has a 5xx status code
func (o *PostOutboundContactlistClearRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound contactlist clear request timeout response a status code equal to that given
func (o *PostOutboundContactlistClearRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostOutboundContactlistClearRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOutboundContactlistClearRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOutboundContactlistClearRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistClearRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistClearRequestEntityTooLarge creates a PostOutboundContactlistClearRequestEntityTooLarge with default headers values
func NewPostOutboundContactlistClearRequestEntityTooLarge() *PostOutboundContactlistClearRequestEntityTooLarge {
	return &PostOutboundContactlistClearRequestEntityTooLarge{}
}

/*
PostOutboundContactlistClearRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostOutboundContactlistClearRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound contactlist clear request entity too large response has a 2xx status code
func (o *PostOutboundContactlistClearRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound contactlist clear request entity too large response has a 3xx status code
func (o *PostOutboundContactlistClearRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist clear request entity too large response has a 4xx status code
func (o *PostOutboundContactlistClearRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound contactlist clear request entity too large response has a 5xx status code
func (o *PostOutboundContactlistClearRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound contactlist clear request entity too large response a status code equal to that given
func (o *PostOutboundContactlistClearRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostOutboundContactlistClearRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOutboundContactlistClearRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOutboundContactlistClearRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistClearRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistClearUnsupportedMediaType creates a PostOutboundContactlistClearUnsupportedMediaType with default headers values
func NewPostOutboundContactlistClearUnsupportedMediaType() *PostOutboundContactlistClearUnsupportedMediaType {
	return &PostOutboundContactlistClearUnsupportedMediaType{}
}

/*
PostOutboundContactlistClearUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostOutboundContactlistClearUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound contactlist clear unsupported media type response has a 2xx status code
func (o *PostOutboundContactlistClearUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound contactlist clear unsupported media type response has a 3xx status code
func (o *PostOutboundContactlistClearUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist clear unsupported media type response has a 4xx status code
func (o *PostOutboundContactlistClearUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound contactlist clear unsupported media type response has a 5xx status code
func (o *PostOutboundContactlistClearUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound contactlist clear unsupported media type response a status code equal to that given
func (o *PostOutboundContactlistClearUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostOutboundContactlistClearUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOutboundContactlistClearUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOutboundContactlistClearUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistClearUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistClearTooManyRequests creates a PostOutboundContactlistClearTooManyRequests with default headers values
func NewPostOutboundContactlistClearTooManyRequests() *PostOutboundContactlistClearTooManyRequests {
	return &PostOutboundContactlistClearTooManyRequests{}
}

/*
PostOutboundContactlistClearTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostOutboundContactlistClearTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound contactlist clear too many requests response has a 2xx status code
func (o *PostOutboundContactlistClearTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound contactlist clear too many requests response has a 3xx status code
func (o *PostOutboundContactlistClearTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist clear too many requests response has a 4xx status code
func (o *PostOutboundContactlistClearTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound contactlist clear too many requests response has a 5xx status code
func (o *PostOutboundContactlistClearTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound contactlist clear too many requests response a status code equal to that given
func (o *PostOutboundContactlistClearTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostOutboundContactlistClearTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOutboundContactlistClearTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOutboundContactlistClearTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistClearTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistClearInternalServerError creates a PostOutboundContactlistClearInternalServerError with default headers values
func NewPostOutboundContactlistClearInternalServerError() *PostOutboundContactlistClearInternalServerError {
	return &PostOutboundContactlistClearInternalServerError{}
}

/*
PostOutboundContactlistClearInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostOutboundContactlistClearInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound contactlist clear internal server error response has a 2xx status code
func (o *PostOutboundContactlistClearInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound contactlist clear internal server error response has a 3xx status code
func (o *PostOutboundContactlistClearInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist clear internal server error response has a 4xx status code
func (o *PostOutboundContactlistClearInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound contactlist clear internal server error response has a 5xx status code
func (o *PostOutboundContactlistClearInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post outbound contactlist clear internal server error response a status code equal to that given
func (o *PostOutboundContactlistClearInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostOutboundContactlistClearInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOutboundContactlistClearInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOutboundContactlistClearInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistClearInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistClearServiceUnavailable creates a PostOutboundContactlistClearServiceUnavailable with default headers values
func NewPostOutboundContactlistClearServiceUnavailable() *PostOutboundContactlistClearServiceUnavailable {
	return &PostOutboundContactlistClearServiceUnavailable{}
}

/*
PostOutboundContactlistClearServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostOutboundContactlistClearServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound contactlist clear service unavailable response has a 2xx status code
func (o *PostOutboundContactlistClearServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound contactlist clear service unavailable response has a 3xx status code
func (o *PostOutboundContactlistClearServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist clear service unavailable response has a 4xx status code
func (o *PostOutboundContactlistClearServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound contactlist clear service unavailable response has a 5xx status code
func (o *PostOutboundContactlistClearServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post outbound contactlist clear service unavailable response a status code equal to that given
func (o *PostOutboundContactlistClearServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostOutboundContactlistClearServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOutboundContactlistClearServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOutboundContactlistClearServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistClearServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistClearGatewayTimeout creates a PostOutboundContactlistClearGatewayTimeout with default headers values
func NewPostOutboundContactlistClearGatewayTimeout() *PostOutboundContactlistClearGatewayTimeout {
	return &PostOutboundContactlistClearGatewayTimeout{}
}

/*
PostOutboundContactlistClearGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostOutboundContactlistClearGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound contactlist clear gateway timeout response has a 2xx status code
func (o *PostOutboundContactlistClearGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound contactlist clear gateway timeout response has a 3xx status code
func (o *PostOutboundContactlistClearGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound contactlist clear gateway timeout response has a 4xx status code
func (o *PostOutboundContactlistClearGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound contactlist clear gateway timeout response has a 5xx status code
func (o *PostOutboundContactlistClearGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post outbound contactlist clear gateway timeout response a status code equal to that given
func (o *PostOutboundContactlistClearGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostOutboundContactlistClearGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOutboundContactlistClearGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOutboundContactlistClearGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistClearGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
