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

// DeleteUserRoutinglanguageReader is a Reader for the DeleteUserRoutinglanguage structure.
type DeleteUserRoutinglanguageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserRoutinglanguageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteUserRoutinglanguageNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUserRoutinglanguageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteUserRoutinglanguageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUserRoutinglanguageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUserRoutinglanguageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteUserRoutinglanguageRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteUserRoutinglanguageRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteUserRoutinglanguageUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteUserRoutinglanguageTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteUserRoutinglanguageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteUserRoutinglanguageServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteUserRoutinglanguageGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserRoutinglanguageNoContent creates a DeleteUserRoutinglanguageNoContent with default headers values
func NewDeleteUserRoutinglanguageNoContent() *DeleteUserRoutinglanguageNoContent {
	return &DeleteUserRoutinglanguageNoContent{}
}

/*DeleteUserRoutinglanguageNoContent handles this case with default header values.

Language removed
*/
type DeleteUserRoutinglanguageNoContent struct {
}

func (o *DeleteUserRoutinglanguageNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routinglanguages/{languageId}][%d] deleteUserRoutinglanguageNoContent ", 204)
}

func (o *DeleteUserRoutinglanguageNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserRoutinglanguageBadRequest creates a DeleteUserRoutinglanguageBadRequest with default headers values
func NewDeleteUserRoutinglanguageBadRequest() *DeleteUserRoutinglanguageBadRequest {
	return &DeleteUserRoutinglanguageBadRequest{}
}

/*DeleteUserRoutinglanguageBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteUserRoutinglanguageBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserRoutinglanguageBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routinglanguages/{languageId}][%d] deleteUserRoutinglanguageBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserRoutinglanguageBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserRoutinglanguageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRoutinglanguageUnauthorized creates a DeleteUserRoutinglanguageUnauthorized with default headers values
func NewDeleteUserRoutinglanguageUnauthorized() *DeleteUserRoutinglanguageUnauthorized {
	return &DeleteUserRoutinglanguageUnauthorized{}
}

/*DeleteUserRoutinglanguageUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteUserRoutinglanguageUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserRoutinglanguageUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routinglanguages/{languageId}][%d] deleteUserRoutinglanguageUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteUserRoutinglanguageUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserRoutinglanguageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRoutinglanguageForbidden creates a DeleteUserRoutinglanguageForbidden with default headers values
func NewDeleteUserRoutinglanguageForbidden() *DeleteUserRoutinglanguageForbidden {
	return &DeleteUserRoutinglanguageForbidden{}
}

/*DeleteUserRoutinglanguageForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteUserRoutinglanguageForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserRoutinglanguageForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routinglanguages/{languageId}][%d] deleteUserRoutinglanguageForbidden  %+v", 403, o.Payload)
}

func (o *DeleteUserRoutinglanguageForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserRoutinglanguageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRoutinglanguageNotFound creates a DeleteUserRoutinglanguageNotFound with default headers values
func NewDeleteUserRoutinglanguageNotFound() *DeleteUserRoutinglanguageNotFound {
	return &DeleteUserRoutinglanguageNotFound{}
}

/*DeleteUserRoutinglanguageNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteUserRoutinglanguageNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserRoutinglanguageNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routinglanguages/{languageId}][%d] deleteUserRoutinglanguageNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserRoutinglanguageNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserRoutinglanguageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRoutinglanguageRequestTimeout creates a DeleteUserRoutinglanguageRequestTimeout with default headers values
func NewDeleteUserRoutinglanguageRequestTimeout() *DeleteUserRoutinglanguageRequestTimeout {
	return &DeleteUserRoutinglanguageRequestTimeout{}
}

/*DeleteUserRoutinglanguageRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteUserRoutinglanguageRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserRoutinglanguageRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routinglanguages/{languageId}][%d] deleteUserRoutinglanguageRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteUserRoutinglanguageRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserRoutinglanguageRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRoutinglanguageRequestEntityTooLarge creates a DeleteUserRoutinglanguageRequestEntityTooLarge with default headers values
func NewDeleteUserRoutinglanguageRequestEntityTooLarge() *DeleteUserRoutinglanguageRequestEntityTooLarge {
	return &DeleteUserRoutinglanguageRequestEntityTooLarge{}
}

/*DeleteUserRoutinglanguageRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteUserRoutinglanguageRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserRoutinglanguageRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routinglanguages/{languageId}][%d] deleteUserRoutinglanguageRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteUserRoutinglanguageRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserRoutinglanguageRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRoutinglanguageUnsupportedMediaType creates a DeleteUserRoutinglanguageUnsupportedMediaType with default headers values
func NewDeleteUserRoutinglanguageUnsupportedMediaType() *DeleteUserRoutinglanguageUnsupportedMediaType {
	return &DeleteUserRoutinglanguageUnsupportedMediaType{}
}

/*DeleteUserRoutinglanguageUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteUserRoutinglanguageUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserRoutinglanguageUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routinglanguages/{languageId}][%d] deleteUserRoutinglanguageUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteUserRoutinglanguageUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserRoutinglanguageUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRoutinglanguageTooManyRequests creates a DeleteUserRoutinglanguageTooManyRequests with default headers values
func NewDeleteUserRoutinglanguageTooManyRequests() *DeleteUserRoutinglanguageTooManyRequests {
	return &DeleteUserRoutinglanguageTooManyRequests{}
}

/*DeleteUserRoutinglanguageTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteUserRoutinglanguageTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserRoutinglanguageTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routinglanguages/{languageId}][%d] deleteUserRoutinglanguageTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteUserRoutinglanguageTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserRoutinglanguageTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRoutinglanguageInternalServerError creates a DeleteUserRoutinglanguageInternalServerError with default headers values
func NewDeleteUserRoutinglanguageInternalServerError() *DeleteUserRoutinglanguageInternalServerError {
	return &DeleteUserRoutinglanguageInternalServerError{}
}

/*DeleteUserRoutinglanguageInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteUserRoutinglanguageInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserRoutinglanguageInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routinglanguages/{languageId}][%d] deleteUserRoutinglanguageInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteUserRoutinglanguageInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserRoutinglanguageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRoutinglanguageServiceUnavailable creates a DeleteUserRoutinglanguageServiceUnavailable with default headers values
func NewDeleteUserRoutinglanguageServiceUnavailable() *DeleteUserRoutinglanguageServiceUnavailable {
	return &DeleteUserRoutinglanguageServiceUnavailable{}
}

/*DeleteUserRoutinglanguageServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteUserRoutinglanguageServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserRoutinglanguageServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routinglanguages/{languageId}][%d] deleteUserRoutinglanguageServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteUserRoutinglanguageServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserRoutinglanguageServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRoutinglanguageGatewayTimeout creates a DeleteUserRoutinglanguageGatewayTimeout with default headers values
func NewDeleteUserRoutinglanguageGatewayTimeout() *DeleteUserRoutinglanguageGatewayTimeout {
	return &DeleteUserRoutinglanguageGatewayTimeout{}
}

/*DeleteUserRoutinglanguageGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteUserRoutinglanguageGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserRoutinglanguageGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routinglanguages/{languageId}][%d] deleteUserRoutinglanguageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteUserRoutinglanguageGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserRoutinglanguageGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
