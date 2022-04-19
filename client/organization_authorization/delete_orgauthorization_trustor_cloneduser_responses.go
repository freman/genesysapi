// Code generated by go-swagger; DO NOT EDIT.

package organization_authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteOrgauthorizationTrustorCloneduserReader is a Reader for the DeleteOrgauthorizationTrustorCloneduser structure.
type DeleteOrgauthorizationTrustorCloneduserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrgauthorizationTrustorCloneduserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOrgauthorizationTrustorCloneduserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteOrgauthorizationTrustorCloneduserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteOrgauthorizationTrustorCloneduserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteOrgauthorizationTrustorCloneduserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOrgauthorizationTrustorCloneduserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteOrgauthorizationTrustorCloneduserRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteOrgauthorizationTrustorCloneduserRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteOrgauthorizationTrustorCloneduserUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteOrgauthorizationTrustorCloneduserTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOrgauthorizationTrustorCloneduserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteOrgauthorizationTrustorCloneduserServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteOrgauthorizationTrustorCloneduserGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOrgauthorizationTrustorCloneduserNoContent creates a DeleteOrgauthorizationTrustorCloneduserNoContent with default headers values
func NewDeleteOrgauthorizationTrustorCloneduserNoContent() *DeleteOrgauthorizationTrustorCloneduserNoContent {
	return &DeleteOrgauthorizationTrustorCloneduserNoContent{}
}

/*DeleteOrgauthorizationTrustorCloneduserNoContent handles this case with default header values.

Trust deleted
*/
type DeleteOrgauthorizationTrustorCloneduserNoContent struct {
}

func (o *DeleteOrgauthorizationTrustorCloneduserNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] deleteOrgauthorizationTrustorCloneduserNoContent ", 204)
}

func (o *DeleteOrgauthorizationTrustorCloneduserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrgauthorizationTrustorCloneduserBadRequest creates a DeleteOrgauthorizationTrustorCloneduserBadRequest with default headers values
func NewDeleteOrgauthorizationTrustorCloneduserBadRequest() *DeleteOrgauthorizationTrustorCloneduserBadRequest {
	return &DeleteOrgauthorizationTrustorCloneduserBadRequest{}
}

/*DeleteOrgauthorizationTrustorCloneduserBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteOrgauthorizationTrustorCloneduserBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrustorCloneduserBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] deleteOrgauthorizationTrustorCloneduserBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOrgauthorizationTrustorCloneduserBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrustorCloneduserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrustorCloneduserUnauthorized creates a DeleteOrgauthorizationTrustorCloneduserUnauthorized with default headers values
func NewDeleteOrgauthorizationTrustorCloneduserUnauthorized() *DeleteOrgauthorizationTrustorCloneduserUnauthorized {
	return &DeleteOrgauthorizationTrustorCloneduserUnauthorized{}
}

/*DeleteOrgauthorizationTrustorCloneduserUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteOrgauthorizationTrustorCloneduserUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrustorCloneduserUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] deleteOrgauthorizationTrustorCloneduserUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOrgauthorizationTrustorCloneduserUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrustorCloneduserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrustorCloneduserForbidden creates a DeleteOrgauthorizationTrustorCloneduserForbidden with default headers values
func NewDeleteOrgauthorizationTrustorCloneduserForbidden() *DeleteOrgauthorizationTrustorCloneduserForbidden {
	return &DeleteOrgauthorizationTrustorCloneduserForbidden{}
}

/*DeleteOrgauthorizationTrustorCloneduserForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteOrgauthorizationTrustorCloneduserForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrustorCloneduserForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] deleteOrgauthorizationTrustorCloneduserForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOrgauthorizationTrustorCloneduserForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrustorCloneduserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrustorCloneduserNotFound creates a DeleteOrgauthorizationTrustorCloneduserNotFound with default headers values
func NewDeleteOrgauthorizationTrustorCloneduserNotFound() *DeleteOrgauthorizationTrustorCloneduserNotFound {
	return &DeleteOrgauthorizationTrustorCloneduserNotFound{}
}

/*DeleteOrgauthorizationTrustorCloneduserNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteOrgauthorizationTrustorCloneduserNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrustorCloneduserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] deleteOrgauthorizationTrustorCloneduserNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOrgauthorizationTrustorCloneduserNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrustorCloneduserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrustorCloneduserRequestTimeout creates a DeleteOrgauthorizationTrustorCloneduserRequestTimeout with default headers values
func NewDeleteOrgauthorizationTrustorCloneduserRequestTimeout() *DeleteOrgauthorizationTrustorCloneduserRequestTimeout {
	return &DeleteOrgauthorizationTrustorCloneduserRequestTimeout{}
}

/*DeleteOrgauthorizationTrustorCloneduserRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteOrgauthorizationTrustorCloneduserRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrustorCloneduserRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] deleteOrgauthorizationTrustorCloneduserRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteOrgauthorizationTrustorCloneduserRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrustorCloneduserRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrustorCloneduserRequestEntityTooLarge creates a DeleteOrgauthorizationTrustorCloneduserRequestEntityTooLarge with default headers values
func NewDeleteOrgauthorizationTrustorCloneduserRequestEntityTooLarge() *DeleteOrgauthorizationTrustorCloneduserRequestEntityTooLarge {
	return &DeleteOrgauthorizationTrustorCloneduserRequestEntityTooLarge{}
}

/*DeleteOrgauthorizationTrustorCloneduserRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteOrgauthorizationTrustorCloneduserRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrustorCloneduserRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] deleteOrgauthorizationTrustorCloneduserRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteOrgauthorizationTrustorCloneduserRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrustorCloneduserRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrustorCloneduserUnsupportedMediaType creates a DeleteOrgauthorizationTrustorCloneduserUnsupportedMediaType with default headers values
func NewDeleteOrgauthorizationTrustorCloneduserUnsupportedMediaType() *DeleteOrgauthorizationTrustorCloneduserUnsupportedMediaType {
	return &DeleteOrgauthorizationTrustorCloneduserUnsupportedMediaType{}
}

/*DeleteOrgauthorizationTrustorCloneduserUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteOrgauthorizationTrustorCloneduserUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrustorCloneduserUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] deleteOrgauthorizationTrustorCloneduserUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteOrgauthorizationTrustorCloneduserUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrustorCloneduserUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrustorCloneduserTooManyRequests creates a DeleteOrgauthorizationTrustorCloneduserTooManyRequests with default headers values
func NewDeleteOrgauthorizationTrustorCloneduserTooManyRequests() *DeleteOrgauthorizationTrustorCloneduserTooManyRequests {
	return &DeleteOrgauthorizationTrustorCloneduserTooManyRequests{}
}

/*DeleteOrgauthorizationTrustorCloneduserTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteOrgauthorizationTrustorCloneduserTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrustorCloneduserTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] deleteOrgauthorizationTrustorCloneduserTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteOrgauthorizationTrustorCloneduserTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrustorCloneduserTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrustorCloneduserInternalServerError creates a DeleteOrgauthorizationTrustorCloneduserInternalServerError with default headers values
func NewDeleteOrgauthorizationTrustorCloneduserInternalServerError() *DeleteOrgauthorizationTrustorCloneduserInternalServerError {
	return &DeleteOrgauthorizationTrustorCloneduserInternalServerError{}
}

/*DeleteOrgauthorizationTrustorCloneduserInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteOrgauthorizationTrustorCloneduserInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrustorCloneduserInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] deleteOrgauthorizationTrustorCloneduserInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOrgauthorizationTrustorCloneduserInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrustorCloneduserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrustorCloneduserServiceUnavailable creates a DeleteOrgauthorizationTrustorCloneduserServiceUnavailable with default headers values
func NewDeleteOrgauthorizationTrustorCloneduserServiceUnavailable() *DeleteOrgauthorizationTrustorCloneduserServiceUnavailable {
	return &DeleteOrgauthorizationTrustorCloneduserServiceUnavailable{}
}

/*DeleteOrgauthorizationTrustorCloneduserServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteOrgauthorizationTrustorCloneduserServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrustorCloneduserServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] deleteOrgauthorizationTrustorCloneduserServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteOrgauthorizationTrustorCloneduserServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrustorCloneduserServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrustorCloneduserGatewayTimeout creates a DeleteOrgauthorizationTrustorCloneduserGatewayTimeout with default headers values
func NewDeleteOrgauthorizationTrustorCloneduserGatewayTimeout() *DeleteOrgauthorizationTrustorCloneduserGatewayTimeout {
	return &DeleteOrgauthorizationTrustorCloneduserGatewayTimeout{}
}

/*DeleteOrgauthorizationTrustorCloneduserGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteOrgauthorizationTrustorCloneduserGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrustorCloneduserGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}][%d] deleteOrgauthorizationTrustorCloneduserGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteOrgauthorizationTrustorCloneduserGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrustorCloneduserGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
