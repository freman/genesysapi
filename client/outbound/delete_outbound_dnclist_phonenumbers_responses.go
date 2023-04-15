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

// DeleteOutboundDnclistPhonenumbersReader is a Reader for the DeleteOutboundDnclistPhonenumbers structure.
type DeleteOutboundDnclistPhonenumbersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOutboundDnclistPhonenumbersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOutboundDnclistPhonenumbersNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteOutboundDnclistPhonenumbersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteOutboundDnclistPhonenumbersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteOutboundDnclistPhonenumbersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOutboundDnclistPhonenumbersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteOutboundDnclistPhonenumbersRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteOutboundDnclistPhonenumbersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteOutboundDnclistPhonenumbersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteOutboundDnclistPhonenumbersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOutboundDnclistPhonenumbersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewDeleteOutboundDnclistPhonenumbersNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteOutboundDnclistPhonenumbersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteOutboundDnclistPhonenumbersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOutboundDnclistPhonenumbersNoContent creates a DeleteOutboundDnclistPhonenumbersNoContent with default headers values
func NewDeleteOutboundDnclistPhonenumbersNoContent() *DeleteOutboundDnclistPhonenumbersNoContent {
	return &DeleteOutboundDnclistPhonenumbersNoContent{}
}

/*
DeleteOutboundDnclistPhonenumbersNoContent describes a response with status code 204, with default header values.

DNC phone numbers deleted.
*/
type DeleteOutboundDnclistPhonenumbersNoContent struct {
}

// IsSuccess returns true when this delete outbound dnclist phonenumbers no content response has a 2xx status code
func (o *DeleteOutboundDnclistPhonenumbersNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete outbound dnclist phonenumbers no content response has a 3xx status code
func (o *DeleteOutboundDnclistPhonenumbersNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound dnclist phonenumbers no content response has a 4xx status code
func (o *DeleteOutboundDnclistPhonenumbersNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete outbound dnclist phonenumbers no content response has a 5xx status code
func (o *DeleteOutboundDnclistPhonenumbersNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound dnclist phonenumbers no content response a status code equal to that given
func (o *DeleteOutboundDnclistPhonenumbersNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteOutboundDnclistPhonenumbersNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersNoContent ", 204)
}

func (o *DeleteOutboundDnclistPhonenumbersNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersNoContent ", 204)
}

func (o *DeleteOutboundDnclistPhonenumbersNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOutboundDnclistPhonenumbersBadRequest creates a DeleteOutboundDnclistPhonenumbersBadRequest with default headers values
func NewDeleteOutboundDnclistPhonenumbersBadRequest() *DeleteOutboundDnclistPhonenumbersBadRequest {
	return &DeleteOutboundDnclistPhonenumbersBadRequest{}
}

/*
DeleteOutboundDnclistPhonenumbersBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteOutboundDnclistPhonenumbersBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound dnclist phonenumbers bad request response has a 2xx status code
func (o *DeleteOutboundDnclistPhonenumbersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound dnclist phonenumbers bad request response has a 3xx status code
func (o *DeleteOutboundDnclistPhonenumbersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound dnclist phonenumbers bad request response has a 4xx status code
func (o *DeleteOutboundDnclistPhonenumbersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound dnclist phonenumbers bad request response has a 5xx status code
func (o *DeleteOutboundDnclistPhonenumbersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound dnclist phonenumbers bad request response a status code equal to that given
func (o *DeleteOutboundDnclistPhonenumbersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteOutboundDnclistPhonenumbersBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundDnclistPhonenumbersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundDnclistPhonenumbersUnauthorized creates a DeleteOutboundDnclistPhonenumbersUnauthorized with default headers values
func NewDeleteOutboundDnclistPhonenumbersUnauthorized() *DeleteOutboundDnclistPhonenumbersUnauthorized {
	return &DeleteOutboundDnclistPhonenumbersUnauthorized{}
}

/*
DeleteOutboundDnclistPhonenumbersUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteOutboundDnclistPhonenumbersUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound dnclist phonenumbers unauthorized response has a 2xx status code
func (o *DeleteOutboundDnclistPhonenumbersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound dnclist phonenumbers unauthorized response has a 3xx status code
func (o *DeleteOutboundDnclistPhonenumbersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound dnclist phonenumbers unauthorized response has a 4xx status code
func (o *DeleteOutboundDnclistPhonenumbersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound dnclist phonenumbers unauthorized response has a 5xx status code
func (o *DeleteOutboundDnclistPhonenumbersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound dnclist phonenumbers unauthorized response a status code equal to that given
func (o *DeleteOutboundDnclistPhonenumbersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteOutboundDnclistPhonenumbersUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundDnclistPhonenumbersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundDnclistPhonenumbersForbidden creates a DeleteOutboundDnclistPhonenumbersForbidden with default headers values
func NewDeleteOutboundDnclistPhonenumbersForbidden() *DeleteOutboundDnclistPhonenumbersForbidden {
	return &DeleteOutboundDnclistPhonenumbersForbidden{}
}

/*
DeleteOutboundDnclistPhonenumbersForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteOutboundDnclistPhonenumbersForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound dnclist phonenumbers forbidden response has a 2xx status code
func (o *DeleteOutboundDnclistPhonenumbersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound dnclist phonenumbers forbidden response has a 3xx status code
func (o *DeleteOutboundDnclistPhonenumbersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound dnclist phonenumbers forbidden response has a 4xx status code
func (o *DeleteOutboundDnclistPhonenumbersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound dnclist phonenumbers forbidden response has a 5xx status code
func (o *DeleteOutboundDnclistPhonenumbersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound dnclist phonenumbers forbidden response a status code equal to that given
func (o *DeleteOutboundDnclistPhonenumbersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteOutboundDnclistPhonenumbersForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundDnclistPhonenumbersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundDnclistPhonenumbersNotFound creates a DeleteOutboundDnclistPhonenumbersNotFound with default headers values
func NewDeleteOutboundDnclistPhonenumbersNotFound() *DeleteOutboundDnclistPhonenumbersNotFound {
	return &DeleteOutboundDnclistPhonenumbersNotFound{}
}

/*
DeleteOutboundDnclistPhonenumbersNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteOutboundDnclistPhonenumbersNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound dnclist phonenumbers not found response has a 2xx status code
func (o *DeleteOutboundDnclistPhonenumbersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound dnclist phonenumbers not found response has a 3xx status code
func (o *DeleteOutboundDnclistPhonenumbersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound dnclist phonenumbers not found response has a 4xx status code
func (o *DeleteOutboundDnclistPhonenumbersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound dnclist phonenumbers not found response has a 5xx status code
func (o *DeleteOutboundDnclistPhonenumbersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound dnclist phonenumbers not found response a status code equal to that given
func (o *DeleteOutboundDnclistPhonenumbersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteOutboundDnclistPhonenumbersNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundDnclistPhonenumbersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundDnclistPhonenumbersRequestTimeout creates a DeleteOutboundDnclistPhonenumbersRequestTimeout with default headers values
func NewDeleteOutboundDnclistPhonenumbersRequestTimeout() *DeleteOutboundDnclistPhonenumbersRequestTimeout {
	return &DeleteOutboundDnclistPhonenumbersRequestTimeout{}
}

/*
DeleteOutboundDnclistPhonenumbersRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteOutboundDnclistPhonenumbersRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound dnclist phonenumbers request timeout response has a 2xx status code
func (o *DeleteOutboundDnclistPhonenumbersRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound dnclist phonenumbers request timeout response has a 3xx status code
func (o *DeleteOutboundDnclistPhonenumbersRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound dnclist phonenumbers request timeout response has a 4xx status code
func (o *DeleteOutboundDnclistPhonenumbersRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound dnclist phonenumbers request timeout response has a 5xx status code
func (o *DeleteOutboundDnclistPhonenumbersRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound dnclist phonenumbers request timeout response a status code equal to that given
func (o *DeleteOutboundDnclistPhonenumbersRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteOutboundDnclistPhonenumbersRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundDnclistPhonenumbersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundDnclistPhonenumbersRequestEntityTooLarge creates a DeleteOutboundDnclistPhonenumbersRequestEntityTooLarge with default headers values
func NewDeleteOutboundDnclistPhonenumbersRequestEntityTooLarge() *DeleteOutboundDnclistPhonenumbersRequestEntityTooLarge {
	return &DeleteOutboundDnclistPhonenumbersRequestEntityTooLarge{}
}

/*
DeleteOutboundDnclistPhonenumbersRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteOutboundDnclistPhonenumbersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound dnclist phonenumbers request entity too large response has a 2xx status code
func (o *DeleteOutboundDnclistPhonenumbersRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound dnclist phonenumbers request entity too large response has a 3xx status code
func (o *DeleteOutboundDnclistPhonenumbersRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound dnclist phonenumbers request entity too large response has a 4xx status code
func (o *DeleteOutboundDnclistPhonenumbersRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound dnclist phonenumbers request entity too large response has a 5xx status code
func (o *DeleteOutboundDnclistPhonenumbersRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound dnclist phonenumbers request entity too large response a status code equal to that given
func (o *DeleteOutboundDnclistPhonenumbersRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteOutboundDnclistPhonenumbersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundDnclistPhonenumbersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundDnclistPhonenumbersUnsupportedMediaType creates a DeleteOutboundDnclistPhonenumbersUnsupportedMediaType with default headers values
func NewDeleteOutboundDnclistPhonenumbersUnsupportedMediaType() *DeleteOutboundDnclistPhonenumbersUnsupportedMediaType {
	return &DeleteOutboundDnclistPhonenumbersUnsupportedMediaType{}
}

/*
DeleteOutboundDnclistPhonenumbersUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteOutboundDnclistPhonenumbersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound dnclist phonenumbers unsupported media type response has a 2xx status code
func (o *DeleteOutboundDnclistPhonenumbersUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound dnclist phonenumbers unsupported media type response has a 3xx status code
func (o *DeleteOutboundDnclistPhonenumbersUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound dnclist phonenumbers unsupported media type response has a 4xx status code
func (o *DeleteOutboundDnclistPhonenumbersUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound dnclist phonenumbers unsupported media type response has a 5xx status code
func (o *DeleteOutboundDnclistPhonenumbersUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound dnclist phonenumbers unsupported media type response a status code equal to that given
func (o *DeleteOutboundDnclistPhonenumbersUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteOutboundDnclistPhonenumbersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundDnclistPhonenumbersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundDnclistPhonenumbersTooManyRequests creates a DeleteOutboundDnclistPhonenumbersTooManyRequests with default headers values
func NewDeleteOutboundDnclistPhonenumbersTooManyRequests() *DeleteOutboundDnclistPhonenumbersTooManyRequests {
	return &DeleteOutboundDnclistPhonenumbersTooManyRequests{}
}

/*
DeleteOutboundDnclistPhonenumbersTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteOutboundDnclistPhonenumbersTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound dnclist phonenumbers too many requests response has a 2xx status code
func (o *DeleteOutboundDnclistPhonenumbersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound dnclist phonenumbers too many requests response has a 3xx status code
func (o *DeleteOutboundDnclistPhonenumbersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound dnclist phonenumbers too many requests response has a 4xx status code
func (o *DeleteOutboundDnclistPhonenumbersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete outbound dnclist phonenumbers too many requests response has a 5xx status code
func (o *DeleteOutboundDnclistPhonenumbersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete outbound dnclist phonenumbers too many requests response a status code equal to that given
func (o *DeleteOutboundDnclistPhonenumbersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteOutboundDnclistPhonenumbersTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundDnclistPhonenumbersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundDnclistPhonenumbersInternalServerError creates a DeleteOutboundDnclistPhonenumbersInternalServerError with default headers values
func NewDeleteOutboundDnclistPhonenumbersInternalServerError() *DeleteOutboundDnclistPhonenumbersInternalServerError {
	return &DeleteOutboundDnclistPhonenumbersInternalServerError{}
}

/*
DeleteOutboundDnclistPhonenumbersInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteOutboundDnclistPhonenumbersInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound dnclist phonenumbers internal server error response has a 2xx status code
func (o *DeleteOutboundDnclistPhonenumbersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound dnclist phonenumbers internal server error response has a 3xx status code
func (o *DeleteOutboundDnclistPhonenumbersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound dnclist phonenumbers internal server error response has a 4xx status code
func (o *DeleteOutboundDnclistPhonenumbersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete outbound dnclist phonenumbers internal server error response has a 5xx status code
func (o *DeleteOutboundDnclistPhonenumbersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete outbound dnclist phonenumbers internal server error response a status code equal to that given
func (o *DeleteOutboundDnclistPhonenumbersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteOutboundDnclistPhonenumbersInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundDnclistPhonenumbersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundDnclistPhonenumbersNotImplemented creates a DeleteOutboundDnclistPhonenumbersNotImplemented with default headers values
func NewDeleteOutboundDnclistPhonenumbersNotImplemented() *DeleteOutboundDnclistPhonenumbersNotImplemented {
	return &DeleteOutboundDnclistPhonenumbersNotImplemented{}
}

/*
DeleteOutboundDnclistPhonenumbersNotImplemented describes a response with status code 501, with default header values.

Not Implemented
*/
type DeleteOutboundDnclistPhonenumbersNotImplemented struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound dnclist phonenumbers not implemented response has a 2xx status code
func (o *DeleteOutboundDnclistPhonenumbersNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound dnclist phonenumbers not implemented response has a 3xx status code
func (o *DeleteOutboundDnclistPhonenumbersNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound dnclist phonenumbers not implemented response has a 4xx status code
func (o *DeleteOutboundDnclistPhonenumbersNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete outbound dnclist phonenumbers not implemented response has a 5xx status code
func (o *DeleteOutboundDnclistPhonenumbersNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this delete outbound dnclist phonenumbers not implemented response a status code equal to that given
func (o *DeleteOutboundDnclistPhonenumbersNotImplemented) IsCode(code int) bool {
	return code == 501
}

func (o *DeleteOutboundDnclistPhonenumbersNotImplemented) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersNotImplemented  %+v", 501, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersNotImplemented) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersNotImplemented  %+v", 501, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersNotImplemented) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundDnclistPhonenumbersNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundDnclistPhonenumbersServiceUnavailable creates a DeleteOutboundDnclistPhonenumbersServiceUnavailable with default headers values
func NewDeleteOutboundDnclistPhonenumbersServiceUnavailable() *DeleteOutboundDnclistPhonenumbersServiceUnavailable {
	return &DeleteOutboundDnclistPhonenumbersServiceUnavailable{}
}

/*
DeleteOutboundDnclistPhonenumbersServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteOutboundDnclistPhonenumbersServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound dnclist phonenumbers service unavailable response has a 2xx status code
func (o *DeleteOutboundDnclistPhonenumbersServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound dnclist phonenumbers service unavailable response has a 3xx status code
func (o *DeleteOutboundDnclistPhonenumbersServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound dnclist phonenumbers service unavailable response has a 4xx status code
func (o *DeleteOutboundDnclistPhonenumbersServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete outbound dnclist phonenumbers service unavailable response has a 5xx status code
func (o *DeleteOutboundDnclistPhonenumbersServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete outbound dnclist phonenumbers service unavailable response a status code equal to that given
func (o *DeleteOutboundDnclistPhonenumbersServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteOutboundDnclistPhonenumbersServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundDnclistPhonenumbersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundDnclistPhonenumbersGatewayTimeout creates a DeleteOutboundDnclistPhonenumbersGatewayTimeout with default headers values
func NewDeleteOutboundDnclistPhonenumbersGatewayTimeout() *DeleteOutboundDnclistPhonenumbersGatewayTimeout {
	return &DeleteOutboundDnclistPhonenumbersGatewayTimeout{}
}

/*
DeleteOutboundDnclistPhonenumbersGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteOutboundDnclistPhonenumbersGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete outbound dnclist phonenumbers gateway timeout response has a 2xx status code
func (o *DeleteOutboundDnclistPhonenumbersGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete outbound dnclist phonenumbers gateway timeout response has a 3xx status code
func (o *DeleteOutboundDnclistPhonenumbersGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete outbound dnclist phonenumbers gateway timeout response has a 4xx status code
func (o *DeleteOutboundDnclistPhonenumbersGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete outbound dnclist phonenumbers gateway timeout response has a 5xx status code
func (o *DeleteOutboundDnclistPhonenumbersGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete outbound dnclist phonenumbers gateway timeout response a status code equal to that given
func (o *DeleteOutboundDnclistPhonenumbersGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteOutboundDnclistPhonenumbersGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}/phonenumbers][%d] deleteOutboundDnclistPhonenumbersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteOutboundDnclistPhonenumbersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundDnclistPhonenumbersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
