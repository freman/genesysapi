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

// DeleteOrgauthorizationTrusteeReader is a Reader for the DeleteOrgauthorizationTrustee structure.
type DeleteOrgauthorizationTrusteeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrgauthorizationTrusteeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOrgauthorizationTrusteeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteOrgauthorizationTrusteeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteOrgauthorizationTrusteeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteOrgauthorizationTrusteeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOrgauthorizationTrusteeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteOrgauthorizationTrusteeRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteOrgauthorizationTrusteeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteOrgauthorizationTrusteeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteOrgauthorizationTrusteeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOrgauthorizationTrusteeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteOrgauthorizationTrusteeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteOrgauthorizationTrusteeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOrgauthorizationTrusteeNoContent creates a DeleteOrgauthorizationTrusteeNoContent with default headers values
func NewDeleteOrgauthorizationTrusteeNoContent() *DeleteOrgauthorizationTrusteeNoContent {
	return &DeleteOrgauthorizationTrusteeNoContent{}
}

/*DeleteOrgauthorizationTrusteeNoContent handles this case with default header values.

Trust deleted
*/
type DeleteOrgauthorizationTrusteeNoContent struct {
}

func (o *DeleteOrgauthorizationTrusteeNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}][%d] deleteOrgauthorizationTrusteeNoContent ", 204)
}

func (o *DeleteOrgauthorizationTrusteeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrgauthorizationTrusteeBadRequest creates a DeleteOrgauthorizationTrusteeBadRequest with default headers values
func NewDeleteOrgauthorizationTrusteeBadRequest() *DeleteOrgauthorizationTrusteeBadRequest {
	return &DeleteOrgauthorizationTrusteeBadRequest{}
}

/*DeleteOrgauthorizationTrusteeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteOrgauthorizationTrusteeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}][%d] deleteOrgauthorizationTrusteeBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeUnauthorized creates a DeleteOrgauthorizationTrusteeUnauthorized with default headers values
func NewDeleteOrgauthorizationTrusteeUnauthorized() *DeleteOrgauthorizationTrusteeUnauthorized {
	return &DeleteOrgauthorizationTrusteeUnauthorized{}
}

/*DeleteOrgauthorizationTrusteeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteOrgauthorizationTrusteeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}][%d] deleteOrgauthorizationTrusteeUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeForbidden creates a DeleteOrgauthorizationTrusteeForbidden with default headers values
func NewDeleteOrgauthorizationTrusteeForbidden() *DeleteOrgauthorizationTrusteeForbidden {
	return &DeleteOrgauthorizationTrusteeForbidden{}
}

/*DeleteOrgauthorizationTrusteeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteOrgauthorizationTrusteeForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}][%d] deleteOrgauthorizationTrusteeForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeNotFound creates a DeleteOrgauthorizationTrusteeNotFound with default headers values
func NewDeleteOrgauthorizationTrusteeNotFound() *DeleteOrgauthorizationTrusteeNotFound {
	return &DeleteOrgauthorizationTrusteeNotFound{}
}

/*DeleteOrgauthorizationTrusteeNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteOrgauthorizationTrusteeNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}][%d] deleteOrgauthorizationTrusteeNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeRequestTimeout creates a DeleteOrgauthorizationTrusteeRequestTimeout with default headers values
func NewDeleteOrgauthorizationTrusteeRequestTimeout() *DeleteOrgauthorizationTrusteeRequestTimeout {
	return &DeleteOrgauthorizationTrusteeRequestTimeout{}
}

/*DeleteOrgauthorizationTrusteeRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteOrgauthorizationTrusteeRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}][%d] deleteOrgauthorizationTrusteeRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeRequestEntityTooLarge creates a DeleteOrgauthorizationTrusteeRequestEntityTooLarge with default headers values
func NewDeleteOrgauthorizationTrusteeRequestEntityTooLarge() *DeleteOrgauthorizationTrusteeRequestEntityTooLarge {
	return &DeleteOrgauthorizationTrusteeRequestEntityTooLarge{}
}

/*DeleteOrgauthorizationTrusteeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteOrgauthorizationTrusteeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}][%d] deleteOrgauthorizationTrusteeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeUnsupportedMediaType creates a DeleteOrgauthorizationTrusteeUnsupportedMediaType with default headers values
func NewDeleteOrgauthorizationTrusteeUnsupportedMediaType() *DeleteOrgauthorizationTrusteeUnsupportedMediaType {
	return &DeleteOrgauthorizationTrusteeUnsupportedMediaType{}
}

/*DeleteOrgauthorizationTrusteeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteOrgauthorizationTrusteeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}][%d] deleteOrgauthorizationTrusteeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeTooManyRequests creates a DeleteOrgauthorizationTrusteeTooManyRequests with default headers values
func NewDeleteOrgauthorizationTrusteeTooManyRequests() *DeleteOrgauthorizationTrusteeTooManyRequests {
	return &DeleteOrgauthorizationTrusteeTooManyRequests{}
}

/*DeleteOrgauthorizationTrusteeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteOrgauthorizationTrusteeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}][%d] deleteOrgauthorizationTrusteeTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeInternalServerError creates a DeleteOrgauthorizationTrusteeInternalServerError with default headers values
func NewDeleteOrgauthorizationTrusteeInternalServerError() *DeleteOrgauthorizationTrusteeInternalServerError {
	return &DeleteOrgauthorizationTrusteeInternalServerError{}
}

/*DeleteOrgauthorizationTrusteeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteOrgauthorizationTrusteeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}][%d] deleteOrgauthorizationTrusteeInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeServiceUnavailable creates a DeleteOrgauthorizationTrusteeServiceUnavailable with default headers values
func NewDeleteOrgauthorizationTrusteeServiceUnavailable() *DeleteOrgauthorizationTrusteeServiceUnavailable {
	return &DeleteOrgauthorizationTrusteeServiceUnavailable{}
}

/*DeleteOrgauthorizationTrusteeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteOrgauthorizationTrusteeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}][%d] deleteOrgauthorizationTrusteeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgauthorizationTrusteeGatewayTimeout creates a DeleteOrgauthorizationTrusteeGatewayTimeout with default headers values
func NewDeleteOrgauthorizationTrusteeGatewayTimeout() *DeleteOrgauthorizationTrusteeGatewayTimeout {
	return &DeleteOrgauthorizationTrusteeGatewayTimeout{}
}

/*DeleteOrgauthorizationTrusteeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteOrgauthorizationTrusteeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrgauthorizationTrusteeGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orgauthorization/trustees/{trusteeOrgId}][%d] deleteOrgauthorizationTrusteeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteOrgauthorizationTrusteeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrgauthorizationTrusteeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
