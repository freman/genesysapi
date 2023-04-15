// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetRoutingMessageRecipientsReader is a Reader for the GetRoutingMessageRecipients structure.
type GetRoutingMessageRecipientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingMessageRecipientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingMessageRecipientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingMessageRecipientsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingMessageRecipientsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingMessageRecipientsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingMessageRecipientsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingMessageRecipientsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingMessageRecipientsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingMessageRecipientsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingMessageRecipientsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingMessageRecipientsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingMessageRecipientsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingMessageRecipientsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingMessageRecipientsOK creates a GetRoutingMessageRecipientsOK with default headers values
func NewGetRoutingMessageRecipientsOK() *GetRoutingMessageRecipientsOK {
	return &GetRoutingMessageRecipientsOK{}
}

/*
GetRoutingMessageRecipientsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRoutingMessageRecipientsOK struct {
	Payload *models.RecipientListing
}

// IsSuccess returns true when this get routing message recipients o k response has a 2xx status code
func (o *GetRoutingMessageRecipientsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get routing message recipients o k response has a 3xx status code
func (o *GetRoutingMessageRecipientsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipients o k response has a 4xx status code
func (o *GetRoutingMessageRecipientsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing message recipients o k response has a 5xx status code
func (o *GetRoutingMessageRecipientsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing message recipients o k response a status code equal to that given
func (o *GetRoutingMessageRecipientsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRoutingMessageRecipientsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsOK  %+v", 200, o.Payload)
}

func (o *GetRoutingMessageRecipientsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsOK  %+v", 200, o.Payload)
}

func (o *GetRoutingMessageRecipientsOK) GetPayload() *models.RecipientListing {
	return o.Payload
}

func (o *GetRoutingMessageRecipientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecipientListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingMessageRecipientsBadRequest creates a GetRoutingMessageRecipientsBadRequest with default headers values
func NewGetRoutingMessageRecipientsBadRequest() *GetRoutingMessageRecipientsBadRequest {
	return &GetRoutingMessageRecipientsBadRequest{}
}

/*
GetRoutingMessageRecipientsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingMessageRecipientsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing message recipients bad request response has a 2xx status code
func (o *GetRoutingMessageRecipientsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing message recipients bad request response has a 3xx status code
func (o *GetRoutingMessageRecipientsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipients bad request response has a 4xx status code
func (o *GetRoutingMessageRecipientsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing message recipients bad request response has a 5xx status code
func (o *GetRoutingMessageRecipientsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing message recipients bad request response a status code equal to that given
func (o *GetRoutingMessageRecipientsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRoutingMessageRecipientsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingMessageRecipientsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingMessageRecipientsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingMessageRecipientsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingMessageRecipientsUnauthorized creates a GetRoutingMessageRecipientsUnauthorized with default headers values
func NewGetRoutingMessageRecipientsUnauthorized() *GetRoutingMessageRecipientsUnauthorized {
	return &GetRoutingMessageRecipientsUnauthorized{}
}

/*
GetRoutingMessageRecipientsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingMessageRecipientsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing message recipients unauthorized response has a 2xx status code
func (o *GetRoutingMessageRecipientsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing message recipients unauthorized response has a 3xx status code
func (o *GetRoutingMessageRecipientsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipients unauthorized response has a 4xx status code
func (o *GetRoutingMessageRecipientsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing message recipients unauthorized response has a 5xx status code
func (o *GetRoutingMessageRecipientsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing message recipients unauthorized response a status code equal to that given
func (o *GetRoutingMessageRecipientsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRoutingMessageRecipientsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingMessageRecipientsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingMessageRecipientsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingMessageRecipientsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingMessageRecipientsForbidden creates a GetRoutingMessageRecipientsForbidden with default headers values
func NewGetRoutingMessageRecipientsForbidden() *GetRoutingMessageRecipientsForbidden {
	return &GetRoutingMessageRecipientsForbidden{}
}

/*
GetRoutingMessageRecipientsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingMessageRecipientsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing message recipients forbidden response has a 2xx status code
func (o *GetRoutingMessageRecipientsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing message recipients forbidden response has a 3xx status code
func (o *GetRoutingMessageRecipientsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipients forbidden response has a 4xx status code
func (o *GetRoutingMessageRecipientsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing message recipients forbidden response has a 5xx status code
func (o *GetRoutingMessageRecipientsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing message recipients forbidden response a status code equal to that given
func (o *GetRoutingMessageRecipientsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRoutingMessageRecipientsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingMessageRecipientsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingMessageRecipientsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingMessageRecipientsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingMessageRecipientsNotFound creates a GetRoutingMessageRecipientsNotFound with default headers values
func NewGetRoutingMessageRecipientsNotFound() *GetRoutingMessageRecipientsNotFound {
	return &GetRoutingMessageRecipientsNotFound{}
}

/*
GetRoutingMessageRecipientsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetRoutingMessageRecipientsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing message recipients not found response has a 2xx status code
func (o *GetRoutingMessageRecipientsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing message recipients not found response has a 3xx status code
func (o *GetRoutingMessageRecipientsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipients not found response has a 4xx status code
func (o *GetRoutingMessageRecipientsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing message recipients not found response has a 5xx status code
func (o *GetRoutingMessageRecipientsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing message recipients not found response a status code equal to that given
func (o *GetRoutingMessageRecipientsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRoutingMessageRecipientsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingMessageRecipientsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingMessageRecipientsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingMessageRecipientsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingMessageRecipientsRequestTimeout creates a GetRoutingMessageRecipientsRequestTimeout with default headers values
func NewGetRoutingMessageRecipientsRequestTimeout() *GetRoutingMessageRecipientsRequestTimeout {
	return &GetRoutingMessageRecipientsRequestTimeout{}
}

/*
GetRoutingMessageRecipientsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingMessageRecipientsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing message recipients request timeout response has a 2xx status code
func (o *GetRoutingMessageRecipientsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing message recipients request timeout response has a 3xx status code
func (o *GetRoutingMessageRecipientsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipients request timeout response has a 4xx status code
func (o *GetRoutingMessageRecipientsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing message recipients request timeout response has a 5xx status code
func (o *GetRoutingMessageRecipientsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing message recipients request timeout response a status code equal to that given
func (o *GetRoutingMessageRecipientsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetRoutingMessageRecipientsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingMessageRecipientsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingMessageRecipientsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingMessageRecipientsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingMessageRecipientsRequestEntityTooLarge creates a GetRoutingMessageRecipientsRequestEntityTooLarge with default headers values
func NewGetRoutingMessageRecipientsRequestEntityTooLarge() *GetRoutingMessageRecipientsRequestEntityTooLarge {
	return &GetRoutingMessageRecipientsRequestEntityTooLarge{}
}

/*
GetRoutingMessageRecipientsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetRoutingMessageRecipientsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing message recipients request entity too large response has a 2xx status code
func (o *GetRoutingMessageRecipientsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing message recipients request entity too large response has a 3xx status code
func (o *GetRoutingMessageRecipientsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipients request entity too large response has a 4xx status code
func (o *GetRoutingMessageRecipientsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing message recipients request entity too large response has a 5xx status code
func (o *GetRoutingMessageRecipientsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing message recipients request entity too large response a status code equal to that given
func (o *GetRoutingMessageRecipientsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRoutingMessageRecipientsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingMessageRecipientsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingMessageRecipientsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingMessageRecipientsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingMessageRecipientsUnsupportedMediaType creates a GetRoutingMessageRecipientsUnsupportedMediaType with default headers values
func NewGetRoutingMessageRecipientsUnsupportedMediaType() *GetRoutingMessageRecipientsUnsupportedMediaType {
	return &GetRoutingMessageRecipientsUnsupportedMediaType{}
}

/*
GetRoutingMessageRecipientsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingMessageRecipientsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing message recipients unsupported media type response has a 2xx status code
func (o *GetRoutingMessageRecipientsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing message recipients unsupported media type response has a 3xx status code
func (o *GetRoutingMessageRecipientsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipients unsupported media type response has a 4xx status code
func (o *GetRoutingMessageRecipientsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing message recipients unsupported media type response has a 5xx status code
func (o *GetRoutingMessageRecipientsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing message recipients unsupported media type response a status code equal to that given
func (o *GetRoutingMessageRecipientsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRoutingMessageRecipientsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingMessageRecipientsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingMessageRecipientsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingMessageRecipientsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingMessageRecipientsTooManyRequests creates a GetRoutingMessageRecipientsTooManyRequests with default headers values
func NewGetRoutingMessageRecipientsTooManyRequests() *GetRoutingMessageRecipientsTooManyRequests {
	return &GetRoutingMessageRecipientsTooManyRequests{}
}

/*
GetRoutingMessageRecipientsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingMessageRecipientsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing message recipients too many requests response has a 2xx status code
func (o *GetRoutingMessageRecipientsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing message recipients too many requests response has a 3xx status code
func (o *GetRoutingMessageRecipientsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipients too many requests response has a 4xx status code
func (o *GetRoutingMessageRecipientsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing message recipients too many requests response has a 5xx status code
func (o *GetRoutingMessageRecipientsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing message recipients too many requests response a status code equal to that given
func (o *GetRoutingMessageRecipientsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRoutingMessageRecipientsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingMessageRecipientsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingMessageRecipientsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingMessageRecipientsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingMessageRecipientsInternalServerError creates a GetRoutingMessageRecipientsInternalServerError with default headers values
func NewGetRoutingMessageRecipientsInternalServerError() *GetRoutingMessageRecipientsInternalServerError {
	return &GetRoutingMessageRecipientsInternalServerError{}
}

/*
GetRoutingMessageRecipientsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingMessageRecipientsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing message recipients internal server error response has a 2xx status code
func (o *GetRoutingMessageRecipientsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing message recipients internal server error response has a 3xx status code
func (o *GetRoutingMessageRecipientsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipients internal server error response has a 4xx status code
func (o *GetRoutingMessageRecipientsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing message recipients internal server error response has a 5xx status code
func (o *GetRoutingMessageRecipientsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing message recipients internal server error response a status code equal to that given
func (o *GetRoutingMessageRecipientsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRoutingMessageRecipientsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingMessageRecipientsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingMessageRecipientsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingMessageRecipientsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingMessageRecipientsServiceUnavailable creates a GetRoutingMessageRecipientsServiceUnavailable with default headers values
func NewGetRoutingMessageRecipientsServiceUnavailable() *GetRoutingMessageRecipientsServiceUnavailable {
	return &GetRoutingMessageRecipientsServiceUnavailable{}
}

/*
GetRoutingMessageRecipientsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingMessageRecipientsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing message recipients service unavailable response has a 2xx status code
func (o *GetRoutingMessageRecipientsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing message recipients service unavailable response has a 3xx status code
func (o *GetRoutingMessageRecipientsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipients service unavailable response has a 4xx status code
func (o *GetRoutingMessageRecipientsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing message recipients service unavailable response has a 5xx status code
func (o *GetRoutingMessageRecipientsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing message recipients service unavailable response a status code equal to that given
func (o *GetRoutingMessageRecipientsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRoutingMessageRecipientsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingMessageRecipientsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingMessageRecipientsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingMessageRecipientsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingMessageRecipientsGatewayTimeout creates a GetRoutingMessageRecipientsGatewayTimeout with default headers values
func NewGetRoutingMessageRecipientsGatewayTimeout() *GetRoutingMessageRecipientsGatewayTimeout {
	return &GetRoutingMessageRecipientsGatewayTimeout{}
}

/*
GetRoutingMessageRecipientsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetRoutingMessageRecipientsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing message recipients gateway timeout response has a 2xx status code
func (o *GetRoutingMessageRecipientsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing message recipients gateway timeout response has a 3xx status code
func (o *GetRoutingMessageRecipientsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipients gateway timeout response has a 4xx status code
func (o *GetRoutingMessageRecipientsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing message recipients gateway timeout response has a 5xx status code
func (o *GetRoutingMessageRecipientsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing message recipients gateway timeout response a status code equal to that given
func (o *GetRoutingMessageRecipientsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetRoutingMessageRecipientsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingMessageRecipientsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients][%d] getRoutingMessageRecipientsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingMessageRecipientsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingMessageRecipientsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
