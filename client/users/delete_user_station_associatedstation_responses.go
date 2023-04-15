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
	case 408:
		result := NewDeleteUserStationAssociatedstationRequestTimeout()
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

/*
DeleteUserStationAssociatedstationAccepted describes a response with status code 202, with default header values.

Success
*/
type DeleteUserStationAssociatedstationAccepted struct {
}

// IsSuccess returns true when this delete user station associatedstation accepted response has a 2xx status code
func (o *DeleteUserStationAssociatedstationAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete user station associatedstation accepted response has a 3xx status code
func (o *DeleteUserStationAssociatedstationAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user station associatedstation accepted response has a 4xx status code
func (o *DeleteUserStationAssociatedstationAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user station associatedstation accepted response has a 5xx status code
func (o *DeleteUserStationAssociatedstationAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user station associatedstation accepted response a status code equal to that given
func (o *DeleteUserStationAssociatedstationAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *DeleteUserStationAssociatedstationAccepted) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationAccepted ", 202)
}

func (o *DeleteUserStationAssociatedstationAccepted) String() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationAccepted ", 202)
}

func (o *DeleteUserStationAssociatedstationAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserStationAssociatedstationBadRequest creates a DeleteUserStationAssociatedstationBadRequest with default headers values
func NewDeleteUserStationAssociatedstationBadRequest() *DeleteUserStationAssociatedstationBadRequest {
	return &DeleteUserStationAssociatedstationBadRequest{}
}

/*
DeleteUserStationAssociatedstationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteUserStationAssociatedstationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete user station associatedstation bad request response has a 2xx status code
func (o *DeleteUserStationAssociatedstationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user station associatedstation bad request response has a 3xx status code
func (o *DeleteUserStationAssociatedstationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user station associatedstation bad request response has a 4xx status code
func (o *DeleteUserStationAssociatedstationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user station associatedstation bad request response has a 5xx status code
func (o *DeleteUserStationAssociatedstationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user station associatedstation bad request response a status code equal to that given
func (o *DeleteUserStationAssociatedstationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteUserStationAssociatedstationBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserStationAssociatedstationBadRequest) String() string {
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

/*
DeleteUserStationAssociatedstationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteUserStationAssociatedstationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete user station associatedstation unauthorized response has a 2xx status code
func (o *DeleteUserStationAssociatedstationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user station associatedstation unauthorized response has a 3xx status code
func (o *DeleteUserStationAssociatedstationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user station associatedstation unauthorized response has a 4xx status code
func (o *DeleteUserStationAssociatedstationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user station associatedstation unauthorized response has a 5xx status code
func (o *DeleteUserStationAssociatedstationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user station associatedstation unauthorized response a status code equal to that given
func (o *DeleteUserStationAssociatedstationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteUserStationAssociatedstationUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteUserStationAssociatedstationUnauthorized) String() string {
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

/*
DeleteUserStationAssociatedstationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteUserStationAssociatedstationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete user station associatedstation forbidden response has a 2xx status code
func (o *DeleteUserStationAssociatedstationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user station associatedstation forbidden response has a 3xx status code
func (o *DeleteUserStationAssociatedstationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user station associatedstation forbidden response has a 4xx status code
func (o *DeleteUserStationAssociatedstationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user station associatedstation forbidden response has a 5xx status code
func (o *DeleteUserStationAssociatedstationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user station associatedstation forbidden response a status code equal to that given
func (o *DeleteUserStationAssociatedstationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteUserStationAssociatedstationForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationForbidden  %+v", 403, o.Payload)
}

func (o *DeleteUserStationAssociatedstationForbidden) String() string {
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

/*
DeleteUserStationAssociatedstationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteUserStationAssociatedstationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete user station associatedstation not found response has a 2xx status code
func (o *DeleteUserStationAssociatedstationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user station associatedstation not found response has a 3xx status code
func (o *DeleteUserStationAssociatedstationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user station associatedstation not found response has a 4xx status code
func (o *DeleteUserStationAssociatedstationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user station associatedstation not found response has a 5xx status code
func (o *DeleteUserStationAssociatedstationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user station associatedstation not found response a status code equal to that given
func (o *DeleteUserStationAssociatedstationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteUserStationAssociatedstationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserStationAssociatedstationNotFound) String() string {
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

// NewDeleteUserStationAssociatedstationRequestTimeout creates a DeleteUserStationAssociatedstationRequestTimeout with default headers values
func NewDeleteUserStationAssociatedstationRequestTimeout() *DeleteUserStationAssociatedstationRequestTimeout {
	return &DeleteUserStationAssociatedstationRequestTimeout{}
}

/*
DeleteUserStationAssociatedstationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteUserStationAssociatedstationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete user station associatedstation request timeout response has a 2xx status code
func (o *DeleteUserStationAssociatedstationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user station associatedstation request timeout response has a 3xx status code
func (o *DeleteUserStationAssociatedstationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user station associatedstation request timeout response has a 4xx status code
func (o *DeleteUserStationAssociatedstationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user station associatedstation request timeout response has a 5xx status code
func (o *DeleteUserStationAssociatedstationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user station associatedstation request timeout response a status code equal to that given
func (o *DeleteUserStationAssociatedstationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteUserStationAssociatedstationRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteUserStationAssociatedstationRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteUserStationAssociatedstationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserStationAssociatedstationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
DeleteUserStationAssociatedstationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteUserStationAssociatedstationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete user station associatedstation request entity too large response has a 2xx status code
func (o *DeleteUserStationAssociatedstationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user station associatedstation request entity too large response has a 3xx status code
func (o *DeleteUserStationAssociatedstationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user station associatedstation request entity too large response has a 4xx status code
func (o *DeleteUserStationAssociatedstationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user station associatedstation request entity too large response has a 5xx status code
func (o *DeleteUserStationAssociatedstationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user station associatedstation request entity too large response a status code equal to that given
func (o *DeleteUserStationAssociatedstationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteUserStationAssociatedstationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteUserStationAssociatedstationRequestEntityTooLarge) String() string {
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

/*
DeleteUserStationAssociatedstationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteUserStationAssociatedstationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete user station associatedstation unsupported media type response has a 2xx status code
func (o *DeleteUserStationAssociatedstationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user station associatedstation unsupported media type response has a 3xx status code
func (o *DeleteUserStationAssociatedstationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user station associatedstation unsupported media type response has a 4xx status code
func (o *DeleteUserStationAssociatedstationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user station associatedstation unsupported media type response has a 5xx status code
func (o *DeleteUserStationAssociatedstationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user station associatedstation unsupported media type response a status code equal to that given
func (o *DeleteUserStationAssociatedstationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteUserStationAssociatedstationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteUserStationAssociatedstationUnsupportedMediaType) String() string {
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

/*
DeleteUserStationAssociatedstationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteUserStationAssociatedstationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete user station associatedstation too many requests response has a 2xx status code
func (o *DeleteUserStationAssociatedstationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user station associatedstation too many requests response has a 3xx status code
func (o *DeleteUserStationAssociatedstationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user station associatedstation too many requests response has a 4xx status code
func (o *DeleteUserStationAssociatedstationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user station associatedstation too many requests response has a 5xx status code
func (o *DeleteUserStationAssociatedstationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user station associatedstation too many requests response a status code equal to that given
func (o *DeleteUserStationAssociatedstationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteUserStationAssociatedstationTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteUserStationAssociatedstationTooManyRequests) String() string {
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

/*
DeleteUserStationAssociatedstationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteUserStationAssociatedstationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete user station associatedstation internal server error response has a 2xx status code
func (o *DeleteUserStationAssociatedstationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user station associatedstation internal server error response has a 3xx status code
func (o *DeleteUserStationAssociatedstationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user station associatedstation internal server error response has a 4xx status code
func (o *DeleteUserStationAssociatedstationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user station associatedstation internal server error response has a 5xx status code
func (o *DeleteUserStationAssociatedstationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete user station associatedstation internal server error response a status code equal to that given
func (o *DeleteUserStationAssociatedstationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteUserStationAssociatedstationInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteUserStationAssociatedstationInternalServerError) String() string {
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

/*
DeleteUserStationAssociatedstationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteUserStationAssociatedstationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete user station associatedstation service unavailable response has a 2xx status code
func (o *DeleteUserStationAssociatedstationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user station associatedstation service unavailable response has a 3xx status code
func (o *DeleteUserStationAssociatedstationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user station associatedstation service unavailable response has a 4xx status code
func (o *DeleteUserStationAssociatedstationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user station associatedstation service unavailable response has a 5xx status code
func (o *DeleteUserStationAssociatedstationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete user station associatedstation service unavailable response a status code equal to that given
func (o *DeleteUserStationAssociatedstationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteUserStationAssociatedstationServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteUserStationAssociatedstationServiceUnavailable) String() string {
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

/*
DeleteUserStationAssociatedstationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteUserStationAssociatedstationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete user station associatedstation gateway timeout response has a 2xx status code
func (o *DeleteUserStationAssociatedstationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user station associatedstation gateway timeout response has a 3xx status code
func (o *DeleteUserStationAssociatedstationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user station associatedstation gateway timeout response has a 4xx status code
func (o *DeleteUserStationAssociatedstationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user station associatedstation gateway timeout response has a 5xx status code
func (o *DeleteUserStationAssociatedstationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete user station associatedstation gateway timeout response a status code equal to that given
func (o *DeleteUserStationAssociatedstationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteUserStationAssociatedstationGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/station/associatedstation][%d] deleteUserStationAssociatedstationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteUserStationAssociatedstationGatewayTimeout) String() string {
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
