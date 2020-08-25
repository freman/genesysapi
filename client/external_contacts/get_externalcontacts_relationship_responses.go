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

// GetExternalcontactsRelationshipReader is a Reader for the GetExternalcontactsRelationship structure.
type GetExternalcontactsRelationshipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalcontactsRelationshipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalcontactsRelationshipOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetExternalcontactsRelationshipBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetExternalcontactsRelationshipUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExternalcontactsRelationshipForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExternalcontactsRelationshipNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetExternalcontactsRelationshipRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetExternalcontactsRelationshipUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetExternalcontactsRelationshipTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetExternalcontactsRelationshipInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetExternalcontactsRelationshipServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetExternalcontactsRelationshipGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExternalcontactsRelationshipOK creates a GetExternalcontactsRelationshipOK with default headers values
func NewGetExternalcontactsRelationshipOK() *GetExternalcontactsRelationshipOK {
	return &GetExternalcontactsRelationshipOK{}
}

/*GetExternalcontactsRelationshipOK handles this case with default header values.

successful operation
*/
type GetExternalcontactsRelationshipOK struct {
	Payload *models.Relationship
}

func (o *GetExternalcontactsRelationshipOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/relationships/{relationshipId}][%d] getExternalcontactsRelationshipOK  %+v", 200, o.Payload)
}

func (o *GetExternalcontactsRelationshipOK) GetPayload() *models.Relationship {
	return o.Payload
}

func (o *GetExternalcontactsRelationshipOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Relationship)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsRelationshipBadRequest creates a GetExternalcontactsRelationshipBadRequest with default headers values
func NewGetExternalcontactsRelationshipBadRequest() *GetExternalcontactsRelationshipBadRequest {
	return &GetExternalcontactsRelationshipBadRequest{}
}

/*GetExternalcontactsRelationshipBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetExternalcontactsRelationshipBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsRelationshipBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/relationships/{relationshipId}][%d] getExternalcontactsRelationshipBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalcontactsRelationshipBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsRelationshipBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsRelationshipUnauthorized creates a GetExternalcontactsRelationshipUnauthorized with default headers values
func NewGetExternalcontactsRelationshipUnauthorized() *GetExternalcontactsRelationshipUnauthorized {
	return &GetExternalcontactsRelationshipUnauthorized{}
}

/*GetExternalcontactsRelationshipUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetExternalcontactsRelationshipUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsRelationshipUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/relationships/{relationshipId}][%d] getExternalcontactsRelationshipUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExternalcontactsRelationshipUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsRelationshipUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsRelationshipForbidden creates a GetExternalcontactsRelationshipForbidden with default headers values
func NewGetExternalcontactsRelationshipForbidden() *GetExternalcontactsRelationshipForbidden {
	return &GetExternalcontactsRelationshipForbidden{}
}

/*GetExternalcontactsRelationshipForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetExternalcontactsRelationshipForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsRelationshipForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/relationships/{relationshipId}][%d] getExternalcontactsRelationshipForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalcontactsRelationshipForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsRelationshipForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsRelationshipNotFound creates a GetExternalcontactsRelationshipNotFound with default headers values
func NewGetExternalcontactsRelationshipNotFound() *GetExternalcontactsRelationshipNotFound {
	return &GetExternalcontactsRelationshipNotFound{}
}

/*GetExternalcontactsRelationshipNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetExternalcontactsRelationshipNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsRelationshipNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/relationships/{relationshipId}][%d] getExternalcontactsRelationshipNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalcontactsRelationshipNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsRelationshipNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsRelationshipRequestEntityTooLarge creates a GetExternalcontactsRelationshipRequestEntityTooLarge with default headers values
func NewGetExternalcontactsRelationshipRequestEntityTooLarge() *GetExternalcontactsRelationshipRequestEntityTooLarge {
	return &GetExternalcontactsRelationshipRequestEntityTooLarge{}
}

/*GetExternalcontactsRelationshipRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetExternalcontactsRelationshipRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsRelationshipRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/relationships/{relationshipId}][%d] getExternalcontactsRelationshipRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetExternalcontactsRelationshipRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsRelationshipRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsRelationshipUnsupportedMediaType creates a GetExternalcontactsRelationshipUnsupportedMediaType with default headers values
func NewGetExternalcontactsRelationshipUnsupportedMediaType() *GetExternalcontactsRelationshipUnsupportedMediaType {
	return &GetExternalcontactsRelationshipUnsupportedMediaType{}
}

/*GetExternalcontactsRelationshipUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetExternalcontactsRelationshipUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsRelationshipUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/relationships/{relationshipId}][%d] getExternalcontactsRelationshipUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetExternalcontactsRelationshipUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsRelationshipUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsRelationshipTooManyRequests creates a GetExternalcontactsRelationshipTooManyRequests with default headers values
func NewGetExternalcontactsRelationshipTooManyRequests() *GetExternalcontactsRelationshipTooManyRequests {
	return &GetExternalcontactsRelationshipTooManyRequests{}
}

/*GetExternalcontactsRelationshipTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetExternalcontactsRelationshipTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsRelationshipTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/relationships/{relationshipId}][%d] getExternalcontactsRelationshipTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExternalcontactsRelationshipTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsRelationshipTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsRelationshipInternalServerError creates a GetExternalcontactsRelationshipInternalServerError with default headers values
func NewGetExternalcontactsRelationshipInternalServerError() *GetExternalcontactsRelationshipInternalServerError {
	return &GetExternalcontactsRelationshipInternalServerError{}
}

/*GetExternalcontactsRelationshipInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetExternalcontactsRelationshipInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsRelationshipInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/relationships/{relationshipId}][%d] getExternalcontactsRelationshipInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalcontactsRelationshipInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsRelationshipInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsRelationshipServiceUnavailable creates a GetExternalcontactsRelationshipServiceUnavailable with default headers values
func NewGetExternalcontactsRelationshipServiceUnavailable() *GetExternalcontactsRelationshipServiceUnavailable {
	return &GetExternalcontactsRelationshipServiceUnavailable{}
}

/*GetExternalcontactsRelationshipServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetExternalcontactsRelationshipServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsRelationshipServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/relationships/{relationshipId}][%d] getExternalcontactsRelationshipServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetExternalcontactsRelationshipServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsRelationshipServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsRelationshipGatewayTimeout creates a GetExternalcontactsRelationshipGatewayTimeout with default headers values
func NewGetExternalcontactsRelationshipGatewayTimeout() *GetExternalcontactsRelationshipGatewayTimeout {
	return &GetExternalcontactsRelationshipGatewayTimeout{}
}

/*GetExternalcontactsRelationshipGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetExternalcontactsRelationshipGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsRelationshipGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/relationships/{relationshipId}][%d] getExternalcontactsRelationshipGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetExternalcontactsRelationshipGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsRelationshipGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}