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

// PostOutboundDnclistPhonenumbersReader is a Reader for the PostOutboundDnclistPhonenumbers structure.
type PostOutboundDnclistPhonenumbersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOutboundDnclistPhonenumbersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewPostOutboundDnclistPhonenumbersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostOutboundDnclistPhonenumbersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostOutboundDnclistPhonenumbersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostOutboundDnclistPhonenumbersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostOutboundDnclistPhonenumbersRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostOutboundDnclistPhonenumbersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostOutboundDnclistPhonenumbersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostOutboundDnclistPhonenumbersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOutboundDnclistPhonenumbersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostOutboundDnclistPhonenumbersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostOutboundDnclistPhonenumbersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPostOutboundDnclistPhonenumbersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostOutboundDnclistPhonenumbersBadRequest creates a PostOutboundDnclistPhonenumbersBadRequest with default headers values
func NewPostOutboundDnclistPhonenumbersBadRequest() *PostOutboundDnclistPhonenumbersBadRequest {
	return &PostOutboundDnclistPhonenumbersBadRequest{}
}

/*
PostOutboundDnclistPhonenumbersBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostOutboundDnclistPhonenumbersBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound dnclist phonenumbers bad request response has a 2xx status code
func (o *PostOutboundDnclistPhonenumbersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound dnclist phonenumbers bad request response has a 3xx status code
func (o *PostOutboundDnclistPhonenumbersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclist phonenumbers bad request response has a 4xx status code
func (o *PostOutboundDnclistPhonenumbersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound dnclist phonenumbers bad request response has a 5xx status code
func (o *PostOutboundDnclistPhonenumbersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound dnclist phonenumbers bad request response a status code equal to that given
func (o *PostOutboundDnclistPhonenumbersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostOutboundDnclistPhonenumbersBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbersBadRequest  %+v", 400, o.Payload)
}

func (o *PostOutboundDnclistPhonenumbersBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbersBadRequest  %+v", 400, o.Payload)
}

func (o *PostOutboundDnclistPhonenumbersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistPhonenumbersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistPhonenumbersUnauthorized creates a PostOutboundDnclistPhonenumbersUnauthorized with default headers values
func NewPostOutboundDnclistPhonenumbersUnauthorized() *PostOutboundDnclistPhonenumbersUnauthorized {
	return &PostOutboundDnclistPhonenumbersUnauthorized{}
}

/*
PostOutboundDnclistPhonenumbersUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostOutboundDnclistPhonenumbersUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound dnclist phonenumbers unauthorized response has a 2xx status code
func (o *PostOutboundDnclistPhonenumbersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound dnclist phonenumbers unauthorized response has a 3xx status code
func (o *PostOutboundDnclistPhonenumbersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclist phonenumbers unauthorized response has a 4xx status code
func (o *PostOutboundDnclistPhonenumbersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound dnclist phonenumbers unauthorized response has a 5xx status code
func (o *PostOutboundDnclistPhonenumbersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound dnclist phonenumbers unauthorized response a status code equal to that given
func (o *PostOutboundDnclistPhonenumbersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostOutboundDnclistPhonenumbersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbersUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOutboundDnclistPhonenumbersUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbersUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOutboundDnclistPhonenumbersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistPhonenumbersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistPhonenumbersForbidden creates a PostOutboundDnclistPhonenumbersForbidden with default headers values
func NewPostOutboundDnclistPhonenumbersForbidden() *PostOutboundDnclistPhonenumbersForbidden {
	return &PostOutboundDnclistPhonenumbersForbidden{}
}

/*
PostOutboundDnclistPhonenumbersForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostOutboundDnclistPhonenumbersForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound dnclist phonenumbers forbidden response has a 2xx status code
func (o *PostOutboundDnclistPhonenumbersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound dnclist phonenumbers forbidden response has a 3xx status code
func (o *PostOutboundDnclistPhonenumbersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclist phonenumbers forbidden response has a 4xx status code
func (o *PostOutboundDnclistPhonenumbersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound dnclist phonenumbers forbidden response has a 5xx status code
func (o *PostOutboundDnclistPhonenumbersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound dnclist phonenumbers forbidden response a status code equal to that given
func (o *PostOutboundDnclistPhonenumbersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostOutboundDnclistPhonenumbersForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbersForbidden  %+v", 403, o.Payload)
}

func (o *PostOutboundDnclistPhonenumbersForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbersForbidden  %+v", 403, o.Payload)
}

func (o *PostOutboundDnclistPhonenumbersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistPhonenumbersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistPhonenumbersNotFound creates a PostOutboundDnclistPhonenumbersNotFound with default headers values
func NewPostOutboundDnclistPhonenumbersNotFound() *PostOutboundDnclistPhonenumbersNotFound {
	return &PostOutboundDnclistPhonenumbersNotFound{}
}

/*
PostOutboundDnclistPhonenumbersNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostOutboundDnclistPhonenumbersNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound dnclist phonenumbers not found response has a 2xx status code
func (o *PostOutboundDnclistPhonenumbersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound dnclist phonenumbers not found response has a 3xx status code
func (o *PostOutboundDnclistPhonenumbersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclist phonenumbers not found response has a 4xx status code
func (o *PostOutboundDnclistPhonenumbersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound dnclist phonenumbers not found response has a 5xx status code
func (o *PostOutboundDnclistPhonenumbersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound dnclist phonenumbers not found response a status code equal to that given
func (o *PostOutboundDnclistPhonenumbersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostOutboundDnclistPhonenumbersNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbersNotFound  %+v", 404, o.Payload)
}

func (o *PostOutboundDnclistPhonenumbersNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbersNotFound  %+v", 404, o.Payload)
}

func (o *PostOutboundDnclistPhonenumbersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistPhonenumbersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistPhonenumbersRequestTimeout creates a PostOutboundDnclistPhonenumbersRequestTimeout with default headers values
func NewPostOutboundDnclistPhonenumbersRequestTimeout() *PostOutboundDnclistPhonenumbersRequestTimeout {
	return &PostOutboundDnclistPhonenumbersRequestTimeout{}
}

/*
PostOutboundDnclistPhonenumbersRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostOutboundDnclistPhonenumbersRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound dnclist phonenumbers request timeout response has a 2xx status code
func (o *PostOutboundDnclistPhonenumbersRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound dnclist phonenumbers request timeout response has a 3xx status code
func (o *PostOutboundDnclistPhonenumbersRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclist phonenumbers request timeout response has a 4xx status code
func (o *PostOutboundDnclistPhonenumbersRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound dnclist phonenumbers request timeout response has a 5xx status code
func (o *PostOutboundDnclistPhonenumbersRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound dnclist phonenumbers request timeout response a status code equal to that given
func (o *PostOutboundDnclistPhonenumbersRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostOutboundDnclistPhonenumbersRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbersRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOutboundDnclistPhonenumbersRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbersRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOutboundDnclistPhonenumbersRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistPhonenumbersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistPhonenumbersRequestEntityTooLarge creates a PostOutboundDnclistPhonenumbersRequestEntityTooLarge with default headers values
func NewPostOutboundDnclistPhonenumbersRequestEntityTooLarge() *PostOutboundDnclistPhonenumbersRequestEntityTooLarge {
	return &PostOutboundDnclistPhonenumbersRequestEntityTooLarge{}
}

/*
PostOutboundDnclistPhonenumbersRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostOutboundDnclistPhonenumbersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound dnclist phonenumbers request entity too large response has a 2xx status code
func (o *PostOutboundDnclistPhonenumbersRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound dnclist phonenumbers request entity too large response has a 3xx status code
func (o *PostOutboundDnclistPhonenumbersRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclist phonenumbers request entity too large response has a 4xx status code
func (o *PostOutboundDnclistPhonenumbersRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound dnclist phonenumbers request entity too large response has a 5xx status code
func (o *PostOutboundDnclistPhonenumbersRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound dnclist phonenumbers request entity too large response a status code equal to that given
func (o *PostOutboundDnclistPhonenumbersRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostOutboundDnclistPhonenumbersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOutboundDnclistPhonenumbersRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOutboundDnclistPhonenumbersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistPhonenumbersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistPhonenumbersUnsupportedMediaType creates a PostOutboundDnclistPhonenumbersUnsupportedMediaType with default headers values
func NewPostOutboundDnclistPhonenumbersUnsupportedMediaType() *PostOutboundDnclistPhonenumbersUnsupportedMediaType {
	return &PostOutboundDnclistPhonenumbersUnsupportedMediaType{}
}

/*
PostOutboundDnclistPhonenumbersUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostOutboundDnclistPhonenumbersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound dnclist phonenumbers unsupported media type response has a 2xx status code
func (o *PostOutboundDnclistPhonenumbersUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound dnclist phonenumbers unsupported media type response has a 3xx status code
func (o *PostOutboundDnclistPhonenumbersUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclist phonenumbers unsupported media type response has a 4xx status code
func (o *PostOutboundDnclistPhonenumbersUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound dnclist phonenumbers unsupported media type response has a 5xx status code
func (o *PostOutboundDnclistPhonenumbersUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound dnclist phonenumbers unsupported media type response a status code equal to that given
func (o *PostOutboundDnclistPhonenumbersUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostOutboundDnclistPhonenumbersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOutboundDnclistPhonenumbersUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOutboundDnclistPhonenumbersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistPhonenumbersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistPhonenumbersTooManyRequests creates a PostOutboundDnclistPhonenumbersTooManyRequests with default headers values
func NewPostOutboundDnclistPhonenumbersTooManyRequests() *PostOutboundDnclistPhonenumbersTooManyRequests {
	return &PostOutboundDnclistPhonenumbersTooManyRequests{}
}

/*
PostOutboundDnclistPhonenumbersTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostOutboundDnclistPhonenumbersTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound dnclist phonenumbers too many requests response has a 2xx status code
func (o *PostOutboundDnclistPhonenumbersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound dnclist phonenumbers too many requests response has a 3xx status code
func (o *PostOutboundDnclistPhonenumbersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclist phonenumbers too many requests response has a 4xx status code
func (o *PostOutboundDnclistPhonenumbersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post outbound dnclist phonenumbers too many requests response has a 5xx status code
func (o *PostOutboundDnclistPhonenumbersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post outbound dnclist phonenumbers too many requests response a status code equal to that given
func (o *PostOutboundDnclistPhonenumbersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostOutboundDnclistPhonenumbersTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbersTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOutboundDnclistPhonenumbersTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbersTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOutboundDnclistPhonenumbersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistPhonenumbersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistPhonenumbersInternalServerError creates a PostOutboundDnclistPhonenumbersInternalServerError with default headers values
func NewPostOutboundDnclistPhonenumbersInternalServerError() *PostOutboundDnclistPhonenumbersInternalServerError {
	return &PostOutboundDnclistPhonenumbersInternalServerError{}
}

/*
PostOutboundDnclistPhonenumbersInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostOutboundDnclistPhonenumbersInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound dnclist phonenumbers internal server error response has a 2xx status code
func (o *PostOutboundDnclistPhonenumbersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound dnclist phonenumbers internal server error response has a 3xx status code
func (o *PostOutboundDnclistPhonenumbersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclist phonenumbers internal server error response has a 4xx status code
func (o *PostOutboundDnclistPhonenumbersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound dnclist phonenumbers internal server error response has a 5xx status code
func (o *PostOutboundDnclistPhonenumbersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post outbound dnclist phonenumbers internal server error response a status code equal to that given
func (o *PostOutboundDnclistPhonenumbersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostOutboundDnclistPhonenumbersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbersInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOutboundDnclistPhonenumbersInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbersInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOutboundDnclistPhonenumbersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistPhonenumbersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistPhonenumbersServiceUnavailable creates a PostOutboundDnclistPhonenumbersServiceUnavailable with default headers values
func NewPostOutboundDnclistPhonenumbersServiceUnavailable() *PostOutboundDnclistPhonenumbersServiceUnavailable {
	return &PostOutboundDnclistPhonenumbersServiceUnavailable{}
}

/*
PostOutboundDnclistPhonenumbersServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostOutboundDnclistPhonenumbersServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound dnclist phonenumbers service unavailable response has a 2xx status code
func (o *PostOutboundDnclistPhonenumbersServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound dnclist phonenumbers service unavailable response has a 3xx status code
func (o *PostOutboundDnclistPhonenumbersServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclist phonenumbers service unavailable response has a 4xx status code
func (o *PostOutboundDnclistPhonenumbersServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound dnclist phonenumbers service unavailable response has a 5xx status code
func (o *PostOutboundDnclistPhonenumbersServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post outbound dnclist phonenumbers service unavailable response a status code equal to that given
func (o *PostOutboundDnclistPhonenumbersServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostOutboundDnclistPhonenumbersServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOutboundDnclistPhonenumbersServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOutboundDnclistPhonenumbersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistPhonenumbersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistPhonenumbersGatewayTimeout creates a PostOutboundDnclistPhonenumbersGatewayTimeout with default headers values
func NewPostOutboundDnclistPhonenumbersGatewayTimeout() *PostOutboundDnclistPhonenumbersGatewayTimeout {
	return &PostOutboundDnclistPhonenumbersGatewayTimeout{}
}

/*
PostOutboundDnclistPhonenumbersGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostOutboundDnclistPhonenumbersGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post outbound dnclist phonenumbers gateway timeout response has a 2xx status code
func (o *PostOutboundDnclistPhonenumbersGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post outbound dnclist phonenumbers gateway timeout response has a 3xx status code
func (o *PostOutboundDnclistPhonenumbersGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post outbound dnclist phonenumbers gateway timeout response has a 4xx status code
func (o *PostOutboundDnclistPhonenumbersGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post outbound dnclist phonenumbers gateway timeout response has a 5xx status code
func (o *PostOutboundDnclistPhonenumbersGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post outbound dnclist phonenumbers gateway timeout response a status code equal to that given
func (o *PostOutboundDnclistPhonenumbersGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostOutboundDnclistPhonenumbersGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOutboundDnclistPhonenumbersGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOutboundDnclistPhonenumbersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistPhonenumbersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistPhonenumbersDefault creates a PostOutboundDnclistPhonenumbersDefault with default headers values
func NewPostOutboundDnclistPhonenumbersDefault(code int) *PostOutboundDnclistPhonenumbersDefault {
	return &PostOutboundDnclistPhonenumbersDefault{
		_statusCode: code,
	}
}

/*
PostOutboundDnclistPhonenumbersDefault describes a response with status code -1, with default header values.

successful operation
*/
type PostOutboundDnclistPhonenumbersDefault struct {
	_statusCode int
}

// Code gets the status code for the post outbound dnclist phonenumbers default response
func (o *PostOutboundDnclistPhonenumbersDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this post outbound dnclist phonenumbers default response has a 2xx status code
func (o *PostOutboundDnclistPhonenumbersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this post outbound dnclist phonenumbers default response has a 3xx status code
func (o *PostOutboundDnclistPhonenumbersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this post outbound dnclist phonenumbers default response has a 4xx status code
func (o *PostOutboundDnclistPhonenumbersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this post outbound dnclist phonenumbers default response has a 5xx status code
func (o *PostOutboundDnclistPhonenumbersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this post outbound dnclist phonenumbers default response a status code equal to that given
func (o *PostOutboundDnclistPhonenumbersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PostOutboundDnclistPhonenumbersDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbers default ", o._statusCode)
}

func (o *PostOutboundDnclistPhonenumbersDefault) String() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] postOutboundDnclistPhonenumbers default ", o._statusCode)
}

func (o *PostOutboundDnclistPhonenumbersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
