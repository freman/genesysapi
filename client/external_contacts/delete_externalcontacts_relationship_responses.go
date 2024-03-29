// Code generated by go-swagger; DO NOT EDIT.

package external_contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteExternalcontactsRelationshipReader is a Reader for the DeleteExternalcontactsRelationship structure.
type DeleteExternalcontactsRelationshipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteExternalcontactsRelationshipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteExternalcontactsRelationshipOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteExternalcontactsRelationshipBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteExternalcontactsRelationshipUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteExternalcontactsRelationshipForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteExternalcontactsRelationshipNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteExternalcontactsRelationshipRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteExternalcontactsRelationshipRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteExternalcontactsRelationshipUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteExternalcontactsRelationshipTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteExternalcontactsRelationshipInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteExternalcontactsRelationshipServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteExternalcontactsRelationshipGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteExternalcontactsRelationshipOK creates a DeleteExternalcontactsRelationshipOK with default headers values
func NewDeleteExternalcontactsRelationshipOK() *DeleteExternalcontactsRelationshipOK {
	return &DeleteExternalcontactsRelationshipOK{}
}

/*
DeleteExternalcontactsRelationshipOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteExternalcontactsRelationshipOK struct {
	Payload models.Empty
}

// IsSuccess returns true when this delete externalcontacts relationship o k response has a 2xx status code
func (o *DeleteExternalcontactsRelationshipOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete externalcontacts relationship o k response has a 3xx status code
func (o *DeleteExternalcontactsRelationshipOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts relationship o k response has a 4xx status code
func (o *DeleteExternalcontactsRelationshipOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete externalcontacts relationship o k response has a 5xx status code
func (o *DeleteExternalcontactsRelationshipOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts relationship o k response a status code equal to that given
func (o *DeleteExternalcontactsRelationshipOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteExternalcontactsRelationshipOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipOK  %+v", 200, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipOK  %+v", 200, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipOK) GetPayload() models.Empty {
	return o.Payload
}

func (o *DeleteExternalcontactsRelationshipOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsRelationshipBadRequest creates a DeleteExternalcontactsRelationshipBadRequest with default headers values
func NewDeleteExternalcontactsRelationshipBadRequest() *DeleteExternalcontactsRelationshipBadRequest {
	return &DeleteExternalcontactsRelationshipBadRequest{}
}

/*
DeleteExternalcontactsRelationshipBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteExternalcontactsRelationshipBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts relationship bad request response has a 2xx status code
func (o *DeleteExternalcontactsRelationshipBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts relationship bad request response has a 3xx status code
func (o *DeleteExternalcontactsRelationshipBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts relationship bad request response has a 4xx status code
func (o *DeleteExternalcontactsRelationshipBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts relationship bad request response has a 5xx status code
func (o *DeleteExternalcontactsRelationshipBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts relationship bad request response a status code equal to that given
func (o *DeleteExternalcontactsRelationshipBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteExternalcontactsRelationshipBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsRelationshipBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsRelationshipUnauthorized creates a DeleteExternalcontactsRelationshipUnauthorized with default headers values
func NewDeleteExternalcontactsRelationshipUnauthorized() *DeleteExternalcontactsRelationshipUnauthorized {
	return &DeleteExternalcontactsRelationshipUnauthorized{}
}

/*
DeleteExternalcontactsRelationshipUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteExternalcontactsRelationshipUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts relationship unauthorized response has a 2xx status code
func (o *DeleteExternalcontactsRelationshipUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts relationship unauthorized response has a 3xx status code
func (o *DeleteExternalcontactsRelationshipUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts relationship unauthorized response has a 4xx status code
func (o *DeleteExternalcontactsRelationshipUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts relationship unauthorized response has a 5xx status code
func (o *DeleteExternalcontactsRelationshipUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts relationship unauthorized response a status code equal to that given
func (o *DeleteExternalcontactsRelationshipUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteExternalcontactsRelationshipUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsRelationshipUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsRelationshipForbidden creates a DeleteExternalcontactsRelationshipForbidden with default headers values
func NewDeleteExternalcontactsRelationshipForbidden() *DeleteExternalcontactsRelationshipForbidden {
	return &DeleteExternalcontactsRelationshipForbidden{}
}

/*
DeleteExternalcontactsRelationshipForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteExternalcontactsRelationshipForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts relationship forbidden response has a 2xx status code
func (o *DeleteExternalcontactsRelationshipForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts relationship forbidden response has a 3xx status code
func (o *DeleteExternalcontactsRelationshipForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts relationship forbidden response has a 4xx status code
func (o *DeleteExternalcontactsRelationshipForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts relationship forbidden response has a 5xx status code
func (o *DeleteExternalcontactsRelationshipForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts relationship forbidden response a status code equal to that given
func (o *DeleteExternalcontactsRelationshipForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteExternalcontactsRelationshipForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipForbidden  %+v", 403, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipForbidden  %+v", 403, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsRelationshipForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsRelationshipNotFound creates a DeleteExternalcontactsRelationshipNotFound with default headers values
func NewDeleteExternalcontactsRelationshipNotFound() *DeleteExternalcontactsRelationshipNotFound {
	return &DeleteExternalcontactsRelationshipNotFound{}
}

/*
DeleteExternalcontactsRelationshipNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteExternalcontactsRelationshipNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts relationship not found response has a 2xx status code
func (o *DeleteExternalcontactsRelationshipNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts relationship not found response has a 3xx status code
func (o *DeleteExternalcontactsRelationshipNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts relationship not found response has a 4xx status code
func (o *DeleteExternalcontactsRelationshipNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts relationship not found response has a 5xx status code
func (o *DeleteExternalcontactsRelationshipNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts relationship not found response a status code equal to that given
func (o *DeleteExternalcontactsRelationshipNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteExternalcontactsRelationshipNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipNotFound  %+v", 404, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipNotFound  %+v", 404, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsRelationshipNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsRelationshipRequestTimeout creates a DeleteExternalcontactsRelationshipRequestTimeout with default headers values
func NewDeleteExternalcontactsRelationshipRequestTimeout() *DeleteExternalcontactsRelationshipRequestTimeout {
	return &DeleteExternalcontactsRelationshipRequestTimeout{}
}

/*
DeleteExternalcontactsRelationshipRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteExternalcontactsRelationshipRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts relationship request timeout response has a 2xx status code
func (o *DeleteExternalcontactsRelationshipRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts relationship request timeout response has a 3xx status code
func (o *DeleteExternalcontactsRelationshipRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts relationship request timeout response has a 4xx status code
func (o *DeleteExternalcontactsRelationshipRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts relationship request timeout response has a 5xx status code
func (o *DeleteExternalcontactsRelationshipRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts relationship request timeout response a status code equal to that given
func (o *DeleteExternalcontactsRelationshipRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteExternalcontactsRelationshipRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsRelationshipRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsRelationshipRequestEntityTooLarge creates a DeleteExternalcontactsRelationshipRequestEntityTooLarge with default headers values
func NewDeleteExternalcontactsRelationshipRequestEntityTooLarge() *DeleteExternalcontactsRelationshipRequestEntityTooLarge {
	return &DeleteExternalcontactsRelationshipRequestEntityTooLarge{}
}

/*
DeleteExternalcontactsRelationshipRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteExternalcontactsRelationshipRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts relationship request entity too large response has a 2xx status code
func (o *DeleteExternalcontactsRelationshipRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts relationship request entity too large response has a 3xx status code
func (o *DeleteExternalcontactsRelationshipRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts relationship request entity too large response has a 4xx status code
func (o *DeleteExternalcontactsRelationshipRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts relationship request entity too large response has a 5xx status code
func (o *DeleteExternalcontactsRelationshipRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts relationship request entity too large response a status code equal to that given
func (o *DeleteExternalcontactsRelationshipRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteExternalcontactsRelationshipRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsRelationshipRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsRelationshipUnsupportedMediaType creates a DeleteExternalcontactsRelationshipUnsupportedMediaType with default headers values
func NewDeleteExternalcontactsRelationshipUnsupportedMediaType() *DeleteExternalcontactsRelationshipUnsupportedMediaType {
	return &DeleteExternalcontactsRelationshipUnsupportedMediaType{}
}

/*
DeleteExternalcontactsRelationshipUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteExternalcontactsRelationshipUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts relationship unsupported media type response has a 2xx status code
func (o *DeleteExternalcontactsRelationshipUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts relationship unsupported media type response has a 3xx status code
func (o *DeleteExternalcontactsRelationshipUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts relationship unsupported media type response has a 4xx status code
func (o *DeleteExternalcontactsRelationshipUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts relationship unsupported media type response has a 5xx status code
func (o *DeleteExternalcontactsRelationshipUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts relationship unsupported media type response a status code equal to that given
func (o *DeleteExternalcontactsRelationshipUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteExternalcontactsRelationshipUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsRelationshipUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsRelationshipTooManyRequests creates a DeleteExternalcontactsRelationshipTooManyRequests with default headers values
func NewDeleteExternalcontactsRelationshipTooManyRequests() *DeleteExternalcontactsRelationshipTooManyRequests {
	return &DeleteExternalcontactsRelationshipTooManyRequests{}
}

/*
DeleteExternalcontactsRelationshipTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteExternalcontactsRelationshipTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts relationship too many requests response has a 2xx status code
func (o *DeleteExternalcontactsRelationshipTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts relationship too many requests response has a 3xx status code
func (o *DeleteExternalcontactsRelationshipTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts relationship too many requests response has a 4xx status code
func (o *DeleteExternalcontactsRelationshipTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts relationship too many requests response has a 5xx status code
func (o *DeleteExternalcontactsRelationshipTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts relationship too many requests response a status code equal to that given
func (o *DeleteExternalcontactsRelationshipTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteExternalcontactsRelationshipTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsRelationshipTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsRelationshipInternalServerError creates a DeleteExternalcontactsRelationshipInternalServerError with default headers values
func NewDeleteExternalcontactsRelationshipInternalServerError() *DeleteExternalcontactsRelationshipInternalServerError {
	return &DeleteExternalcontactsRelationshipInternalServerError{}
}

/*
DeleteExternalcontactsRelationshipInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteExternalcontactsRelationshipInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts relationship internal server error response has a 2xx status code
func (o *DeleteExternalcontactsRelationshipInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts relationship internal server error response has a 3xx status code
func (o *DeleteExternalcontactsRelationshipInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts relationship internal server error response has a 4xx status code
func (o *DeleteExternalcontactsRelationshipInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete externalcontacts relationship internal server error response has a 5xx status code
func (o *DeleteExternalcontactsRelationshipInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete externalcontacts relationship internal server error response a status code equal to that given
func (o *DeleteExternalcontactsRelationshipInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteExternalcontactsRelationshipInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsRelationshipInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsRelationshipServiceUnavailable creates a DeleteExternalcontactsRelationshipServiceUnavailable with default headers values
func NewDeleteExternalcontactsRelationshipServiceUnavailable() *DeleteExternalcontactsRelationshipServiceUnavailable {
	return &DeleteExternalcontactsRelationshipServiceUnavailable{}
}

/*
DeleteExternalcontactsRelationshipServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteExternalcontactsRelationshipServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts relationship service unavailable response has a 2xx status code
func (o *DeleteExternalcontactsRelationshipServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts relationship service unavailable response has a 3xx status code
func (o *DeleteExternalcontactsRelationshipServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts relationship service unavailable response has a 4xx status code
func (o *DeleteExternalcontactsRelationshipServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete externalcontacts relationship service unavailable response has a 5xx status code
func (o *DeleteExternalcontactsRelationshipServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete externalcontacts relationship service unavailable response a status code equal to that given
func (o *DeleteExternalcontactsRelationshipServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteExternalcontactsRelationshipServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsRelationshipServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsRelationshipGatewayTimeout creates a DeleteExternalcontactsRelationshipGatewayTimeout with default headers values
func NewDeleteExternalcontactsRelationshipGatewayTimeout() *DeleteExternalcontactsRelationshipGatewayTimeout {
	return &DeleteExternalcontactsRelationshipGatewayTimeout{}
}

/*
DeleteExternalcontactsRelationshipGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteExternalcontactsRelationshipGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts relationship gateway timeout response has a 2xx status code
func (o *DeleteExternalcontactsRelationshipGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts relationship gateway timeout response has a 3xx status code
func (o *DeleteExternalcontactsRelationshipGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts relationship gateway timeout response has a 4xx status code
func (o *DeleteExternalcontactsRelationshipGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete externalcontacts relationship gateway timeout response has a 5xx status code
func (o *DeleteExternalcontactsRelationshipGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete externalcontacts relationship gateway timeout response a status code equal to that given
func (o *DeleteExternalcontactsRelationshipGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteExternalcontactsRelationshipGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/relationships/{relationshipId}][%d] deleteExternalcontactsRelationshipGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteExternalcontactsRelationshipGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsRelationshipGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
