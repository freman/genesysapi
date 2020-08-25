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

// DeleteUserStationAssociatedstationReader is a Reader for the DeleteUserStationAssociatedstation structure.
type DeleteUserStationAssociatedstationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserStationAssociatedstationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteUserStationAssociatedstationAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUserStationAssociatedstationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteUserStationAssociatedstationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUserStationAssociatedstationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUserStationAssociatedstationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteUserStationAssociatedstationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteUserStationAssociatedstationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteUserStationAssociatedstationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteUserStationAssociatedstationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteUserStationAssociatedstationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteUserStationAssociatedstationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserStationAssociatedstationAccepted creates a DeleteUserStationAssociatedstationAccepted with default headers values
func NewDeleteUserStationAssociatedstationAccepted() *DeleteUserStationAssociatedstationAccepted {
	return &DeleteUserStationAssociatedstationAccepted{}
}

/*DeleteUserStationAssociatedstationAccepted handles this case with default header values.

Success
*/
type DeleteUserStationAssociatedstationAccepted struct {
}

func (o *DeleteUserStationAssociatedstationAccepted) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationAccepted ", 202)
}

func (o *DeleteUserStationAssociatedstationAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserStationAssociatedstationBadRequest creates a DeleteUserStationAssociatedstationBadRequest with default headers values
func NewDeleteUserStationAssociatedstationBadRequest() *DeleteUserStationAssociatedstationBadRequest {
	return &DeleteUserStationAssociatedstationBadRequest{}
}

/*DeleteUserStationAssociatedstationBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteUserStationAssociatedstationBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserStationAssociatedstationBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserStationAssociatedstationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserStationAssociatedstationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserStationAssociatedstationUnauthorized creates a DeleteUserStationAssociatedstationUnauthorized with default headers values
func NewDeleteUserStationAssociatedstationUnauthorized() *DeleteUserStationAssociatedstationUnauthorized {
	return &DeleteUserStationAssociatedstationUnauthorized{}
}

/*DeleteUserStationAssociatedstationUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteUserStationAssociatedstationUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserStationAssociatedstationUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteUserStationAssociatedstationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserStationAssociatedstationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserStationAssociatedstationForbidden creates a DeleteUserStationAssociatedstationForbidden with default headers values
func NewDeleteUserStationAssociatedstationForbidden() *DeleteUserStationAssociatedstationForbidden {
	return &DeleteUserStationAssociatedstationForbidden{}
}

/*DeleteUserStationAssociatedstationForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteUserStationAssociatedstationForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserStationAssociatedstationForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationForbidden  %+v", 403, o.Payload)
}

func (o *DeleteUserStationAssociatedstationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserStationAssociatedstationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserStationAssociatedstationNotFound creates a DeleteUserStationAssociatedstationNotFound with default headers values
func NewDeleteUserStationAssociatedstationNotFound() *DeleteUserStationAssociatedstationNotFound {
	return &DeleteUserStationAssociatedstationNotFound{}
}

/*DeleteUserStationAssociatedstationNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteUserStationAssociatedstationNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserStationAssociatedstationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserStationAssociatedstationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserStationAssociatedstationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserStationAssociatedstationRequestEntityTooLarge creates a DeleteUserStationAssociatedstationRequestEntityTooLarge with default headers values
func NewDeleteUserStationAssociatedstationRequestEntityTooLarge() *DeleteUserStationAssociatedstationRequestEntityTooLarge {
	return &DeleteUserStationAssociatedstationRequestEntityTooLarge{}
}

/*DeleteUserStationAssociatedstationRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteUserStationAssociatedstationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserStationAssociatedstationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteUserStationAssociatedstationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserStationAssociatedstationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserStationAssociatedstationUnsupportedMediaType creates a DeleteUserStationAssociatedstationUnsupportedMediaType with default headers values
func NewDeleteUserStationAssociatedstationUnsupportedMediaType() *DeleteUserStationAssociatedstationUnsupportedMediaType {
	return &DeleteUserStationAssociatedstationUnsupportedMediaType{}
}

/*DeleteUserStationAssociatedstationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteUserStationAssociatedstationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserStationAssociatedstationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteUserStationAssociatedstationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserStationAssociatedstationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserStationAssociatedstationTooManyRequests creates a DeleteUserStationAssociatedstationTooManyRequests with default headers values
func NewDeleteUserStationAssociatedstationTooManyRequests() *DeleteUserStationAssociatedstationTooManyRequests {
	return &DeleteUserStationAssociatedstationTooManyRequests{}
}

/*DeleteUserStationAssociatedstationTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteUserStationAssociatedstationTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserStationAssociatedstationTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteUserStationAssociatedstationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserStationAssociatedstationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserStationAssociatedstationInternalServerError creates a DeleteUserStationAssociatedstationInternalServerError with default headers values
func NewDeleteUserStationAssociatedstationInternalServerError() *DeleteUserStationAssociatedstationInternalServerError {
	return &DeleteUserStationAssociatedstationInternalServerError{}
}

/*DeleteUserStationAssociatedstationInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteUserStationAssociatedstationInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserStationAssociatedstationInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteUserStationAssociatedstationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserStationAssociatedstationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserStationAssociatedstationServiceUnavailable creates a DeleteUserStationAssociatedstationServiceUnavailable with default headers values
func NewDeleteUserStationAssociatedstationServiceUnavailable() *DeleteUserStationAssociatedstationServiceUnavailable {
	return &DeleteUserStationAssociatedstationServiceUnavailable{}
}

/*DeleteUserStationAssociatedstationServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteUserStationAssociatedstationServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserStationAssociatedstationServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteUserStationAssociatedstationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserStationAssociatedstationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserStationAssociatedstationGatewayTimeout creates a DeleteUserStationAssociatedstationGatewayTimeout with default headers values
func NewDeleteUserStationAssociatedstationGatewayTimeout() *DeleteUserStationAssociatedstationGatewayTimeout {
	return &DeleteUserStationAssociatedstationGatewayTimeout{}
}

/*DeleteUserStationAssociatedstationGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteUserStationAssociatedstationGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserStationAssociatedstationGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteUserStationAssociatedstationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserStationAssociatedstationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}