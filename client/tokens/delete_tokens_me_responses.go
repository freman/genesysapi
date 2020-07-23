// Code generated by go-swagger; DO NOT EDIT.

package tokens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteTokensMeReader is a Reader for the DeleteTokensMe structure.
type DeleteTokensMeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTokensMeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTokensMeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTokensMeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteTokensMeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTokensMeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTokensMeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteTokensMeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteTokensMeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteTokensMeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTokensMeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteTokensMeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteTokensMeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTokensMeOK creates a DeleteTokensMeOK with default headers values
func NewDeleteTokensMeOK() *DeleteTokensMeOK {
	return &DeleteTokensMeOK{}
}

/*DeleteTokensMeOK handles this case with default header values.

Operation was successful.
*/
type DeleteTokensMeOK struct {
}

func (o *DeleteTokensMeOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/tokens/me][%d] deleteTokensMeOK ", 200)
}

func (o *DeleteTokensMeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTokensMeBadRequest creates a DeleteTokensMeBadRequest with default headers values
func NewDeleteTokensMeBadRequest() *DeleteTokensMeBadRequest {
	return &DeleteTokensMeBadRequest{}
}

/*DeleteTokensMeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteTokensMeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteTokensMeBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/tokens/me][%d] deleteTokensMeBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTokensMeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTokensMeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTokensMeUnauthorized creates a DeleteTokensMeUnauthorized with default headers values
func NewDeleteTokensMeUnauthorized() *DeleteTokensMeUnauthorized {
	return &DeleteTokensMeUnauthorized{}
}

/*DeleteTokensMeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteTokensMeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteTokensMeUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/tokens/me][%d] deleteTokensMeUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTokensMeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTokensMeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTokensMeForbidden creates a DeleteTokensMeForbidden with default headers values
func NewDeleteTokensMeForbidden() *DeleteTokensMeForbidden {
	return &DeleteTokensMeForbidden{}
}

/*DeleteTokensMeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteTokensMeForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteTokensMeForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/tokens/me][%d] deleteTokensMeForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTokensMeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTokensMeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTokensMeNotFound creates a DeleteTokensMeNotFound with default headers values
func NewDeleteTokensMeNotFound() *DeleteTokensMeNotFound {
	return &DeleteTokensMeNotFound{}
}

/*DeleteTokensMeNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteTokensMeNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteTokensMeNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/tokens/me][%d] deleteTokensMeNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTokensMeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTokensMeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTokensMeRequestEntityTooLarge creates a DeleteTokensMeRequestEntityTooLarge with default headers values
func NewDeleteTokensMeRequestEntityTooLarge() *DeleteTokensMeRequestEntityTooLarge {
	return &DeleteTokensMeRequestEntityTooLarge{}
}

/*DeleteTokensMeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteTokensMeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteTokensMeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/tokens/me][%d] deleteTokensMeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteTokensMeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTokensMeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTokensMeUnsupportedMediaType creates a DeleteTokensMeUnsupportedMediaType with default headers values
func NewDeleteTokensMeUnsupportedMediaType() *DeleteTokensMeUnsupportedMediaType {
	return &DeleteTokensMeUnsupportedMediaType{}
}

/*DeleteTokensMeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteTokensMeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteTokensMeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/tokens/me][%d] deleteTokensMeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteTokensMeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTokensMeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTokensMeTooManyRequests creates a DeleteTokensMeTooManyRequests with default headers values
func NewDeleteTokensMeTooManyRequests() *DeleteTokensMeTooManyRequests {
	return &DeleteTokensMeTooManyRequests{}
}

/*DeleteTokensMeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteTokensMeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteTokensMeTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/tokens/me][%d] deleteTokensMeTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteTokensMeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTokensMeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTokensMeInternalServerError creates a DeleteTokensMeInternalServerError with default headers values
func NewDeleteTokensMeInternalServerError() *DeleteTokensMeInternalServerError {
	return &DeleteTokensMeInternalServerError{}
}

/*DeleteTokensMeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteTokensMeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteTokensMeInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/tokens/me][%d] deleteTokensMeInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTokensMeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTokensMeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTokensMeServiceUnavailable creates a DeleteTokensMeServiceUnavailable with default headers values
func NewDeleteTokensMeServiceUnavailable() *DeleteTokensMeServiceUnavailable {
	return &DeleteTokensMeServiceUnavailable{}
}

/*DeleteTokensMeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteTokensMeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteTokensMeServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/tokens/me][%d] deleteTokensMeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteTokensMeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTokensMeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTokensMeGatewayTimeout creates a DeleteTokensMeGatewayTimeout with default headers values
func NewDeleteTokensMeGatewayTimeout() *DeleteTokensMeGatewayTimeout {
	return &DeleteTokensMeGatewayTimeout{}
}

/*DeleteTokensMeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteTokensMeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteTokensMeGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/tokens/me][%d] deleteTokensMeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteTokensMeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTokensMeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
