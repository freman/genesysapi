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

/*DeleteExternalcontactsRelationshipOK handles this case with default header values.

successful operation
*/
type DeleteExternalcontactsRelationshipOK struct {
	Payload models.Empty
}

func (o *DeleteExternalcontactsRelationshipOK) Error() string {
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

/*DeleteExternalcontactsRelationshipBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteExternalcontactsRelationshipBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsRelationshipBadRequest) Error() string {
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

/*DeleteExternalcontactsRelationshipUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteExternalcontactsRelationshipUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsRelationshipUnauthorized) Error() string {
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

/*DeleteExternalcontactsRelationshipForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteExternalcontactsRelationshipForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsRelationshipForbidden) Error() string {
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

/*DeleteExternalcontactsRelationshipNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteExternalcontactsRelationshipNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsRelationshipNotFound) Error() string {
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

// NewDeleteExternalcontactsRelationshipRequestEntityTooLarge creates a DeleteExternalcontactsRelationshipRequestEntityTooLarge with default headers values
func NewDeleteExternalcontactsRelationshipRequestEntityTooLarge() *DeleteExternalcontactsRelationshipRequestEntityTooLarge {
	return &DeleteExternalcontactsRelationshipRequestEntityTooLarge{}
}

/*DeleteExternalcontactsRelationshipRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteExternalcontactsRelationshipRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsRelationshipRequestEntityTooLarge) Error() string {
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

/*DeleteExternalcontactsRelationshipUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteExternalcontactsRelationshipUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsRelationshipUnsupportedMediaType) Error() string {
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

/*DeleteExternalcontactsRelationshipTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteExternalcontactsRelationshipTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsRelationshipTooManyRequests) Error() string {
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

/*DeleteExternalcontactsRelationshipInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteExternalcontactsRelationshipInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsRelationshipInternalServerError) Error() string {
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

/*DeleteExternalcontactsRelationshipServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteExternalcontactsRelationshipServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsRelationshipServiceUnavailable) Error() string {
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

/*DeleteExternalcontactsRelationshipGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteExternalcontactsRelationshipGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsRelationshipGatewayTimeout) Error() string {
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
