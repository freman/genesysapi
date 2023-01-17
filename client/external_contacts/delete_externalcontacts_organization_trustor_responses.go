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

// DeleteExternalcontactsOrganizationTrustorReader is a Reader for the DeleteExternalcontactsOrganizationTrustor structure.
type DeleteExternalcontactsOrganizationTrustorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteExternalcontactsOrganizationTrustorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteExternalcontactsOrganizationTrustorNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteExternalcontactsOrganizationTrustorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteExternalcontactsOrganizationTrustorUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteExternalcontactsOrganizationTrustorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteExternalcontactsOrganizationTrustorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteExternalcontactsOrganizationTrustorRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteExternalcontactsOrganizationTrustorRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteExternalcontactsOrganizationTrustorUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteExternalcontactsOrganizationTrustorTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteExternalcontactsOrganizationTrustorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteExternalcontactsOrganizationTrustorServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteExternalcontactsOrganizationTrustorGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteExternalcontactsOrganizationTrustorNoContent creates a DeleteExternalcontactsOrganizationTrustorNoContent with default headers values
func NewDeleteExternalcontactsOrganizationTrustorNoContent() *DeleteExternalcontactsOrganizationTrustorNoContent {
	return &DeleteExternalcontactsOrganizationTrustorNoContent{}
}

/*
DeleteExternalcontactsOrganizationTrustorNoContent describes a response with status code 204, with default header values.

Trustor link has been deleted
*/
type DeleteExternalcontactsOrganizationTrustorNoContent struct {
}

// IsSuccess returns true when this delete externalcontacts organization trustor no content response has a 2xx status code
func (o *DeleteExternalcontactsOrganizationTrustorNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete externalcontacts organization trustor no content response has a 3xx status code
func (o *DeleteExternalcontactsOrganizationTrustorNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts organization trustor no content response has a 4xx status code
func (o *DeleteExternalcontactsOrganizationTrustorNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete externalcontacts organization trustor no content response has a 5xx status code
func (o *DeleteExternalcontactsOrganizationTrustorNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts organization trustor no content response a status code equal to that given
func (o *DeleteExternalcontactsOrganizationTrustorNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteExternalcontactsOrganizationTrustorNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorNoContent ", 204)
}

func (o *DeleteExternalcontactsOrganizationTrustorNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorNoContent ", 204)
}

func (o *DeleteExternalcontactsOrganizationTrustorNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteExternalcontactsOrganizationTrustorBadRequest creates a DeleteExternalcontactsOrganizationTrustorBadRequest with default headers values
func NewDeleteExternalcontactsOrganizationTrustorBadRequest() *DeleteExternalcontactsOrganizationTrustorBadRequest {
	return &DeleteExternalcontactsOrganizationTrustorBadRequest{}
}

/*
DeleteExternalcontactsOrganizationTrustorBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteExternalcontactsOrganizationTrustorBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts organization trustor bad request response has a 2xx status code
func (o *DeleteExternalcontactsOrganizationTrustorBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts organization trustor bad request response has a 3xx status code
func (o *DeleteExternalcontactsOrganizationTrustorBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts organization trustor bad request response has a 4xx status code
func (o *DeleteExternalcontactsOrganizationTrustorBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts organization trustor bad request response has a 5xx status code
func (o *DeleteExternalcontactsOrganizationTrustorBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts organization trustor bad request response a status code equal to that given
func (o *DeleteExternalcontactsOrganizationTrustorBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteExternalcontactsOrganizationTrustorBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteExternalcontactsOrganizationTrustorBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteExternalcontactsOrganizationTrustorBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsOrganizationTrustorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsOrganizationTrustorUnauthorized creates a DeleteExternalcontactsOrganizationTrustorUnauthorized with default headers values
func NewDeleteExternalcontactsOrganizationTrustorUnauthorized() *DeleteExternalcontactsOrganizationTrustorUnauthorized {
	return &DeleteExternalcontactsOrganizationTrustorUnauthorized{}
}

/*
DeleteExternalcontactsOrganizationTrustorUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteExternalcontactsOrganizationTrustorUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts organization trustor unauthorized response has a 2xx status code
func (o *DeleteExternalcontactsOrganizationTrustorUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts organization trustor unauthorized response has a 3xx status code
func (o *DeleteExternalcontactsOrganizationTrustorUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts organization trustor unauthorized response has a 4xx status code
func (o *DeleteExternalcontactsOrganizationTrustorUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts organization trustor unauthorized response has a 5xx status code
func (o *DeleteExternalcontactsOrganizationTrustorUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts organization trustor unauthorized response a status code equal to that given
func (o *DeleteExternalcontactsOrganizationTrustorUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteExternalcontactsOrganizationTrustorUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteExternalcontactsOrganizationTrustorUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteExternalcontactsOrganizationTrustorUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsOrganizationTrustorUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsOrganizationTrustorForbidden creates a DeleteExternalcontactsOrganizationTrustorForbidden with default headers values
func NewDeleteExternalcontactsOrganizationTrustorForbidden() *DeleteExternalcontactsOrganizationTrustorForbidden {
	return &DeleteExternalcontactsOrganizationTrustorForbidden{}
}

/*
DeleteExternalcontactsOrganizationTrustorForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteExternalcontactsOrganizationTrustorForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts organization trustor forbidden response has a 2xx status code
func (o *DeleteExternalcontactsOrganizationTrustorForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts organization trustor forbidden response has a 3xx status code
func (o *DeleteExternalcontactsOrganizationTrustorForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts organization trustor forbidden response has a 4xx status code
func (o *DeleteExternalcontactsOrganizationTrustorForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts organization trustor forbidden response has a 5xx status code
func (o *DeleteExternalcontactsOrganizationTrustorForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts organization trustor forbidden response a status code equal to that given
func (o *DeleteExternalcontactsOrganizationTrustorForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteExternalcontactsOrganizationTrustorForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorForbidden  %+v", 403, o.Payload)
}

func (o *DeleteExternalcontactsOrganizationTrustorForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorForbidden  %+v", 403, o.Payload)
}

func (o *DeleteExternalcontactsOrganizationTrustorForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsOrganizationTrustorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsOrganizationTrustorNotFound creates a DeleteExternalcontactsOrganizationTrustorNotFound with default headers values
func NewDeleteExternalcontactsOrganizationTrustorNotFound() *DeleteExternalcontactsOrganizationTrustorNotFound {
	return &DeleteExternalcontactsOrganizationTrustorNotFound{}
}

/*
DeleteExternalcontactsOrganizationTrustorNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteExternalcontactsOrganizationTrustorNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts organization trustor not found response has a 2xx status code
func (o *DeleteExternalcontactsOrganizationTrustorNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts organization trustor not found response has a 3xx status code
func (o *DeleteExternalcontactsOrganizationTrustorNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts organization trustor not found response has a 4xx status code
func (o *DeleteExternalcontactsOrganizationTrustorNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts organization trustor not found response has a 5xx status code
func (o *DeleteExternalcontactsOrganizationTrustorNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts organization trustor not found response a status code equal to that given
func (o *DeleteExternalcontactsOrganizationTrustorNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteExternalcontactsOrganizationTrustorNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorNotFound  %+v", 404, o.Payload)
}

func (o *DeleteExternalcontactsOrganizationTrustorNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorNotFound  %+v", 404, o.Payload)
}

func (o *DeleteExternalcontactsOrganizationTrustorNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsOrganizationTrustorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsOrganizationTrustorRequestTimeout creates a DeleteExternalcontactsOrganizationTrustorRequestTimeout with default headers values
func NewDeleteExternalcontactsOrganizationTrustorRequestTimeout() *DeleteExternalcontactsOrganizationTrustorRequestTimeout {
	return &DeleteExternalcontactsOrganizationTrustorRequestTimeout{}
}

/*
DeleteExternalcontactsOrganizationTrustorRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteExternalcontactsOrganizationTrustorRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts organization trustor request timeout response has a 2xx status code
func (o *DeleteExternalcontactsOrganizationTrustorRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts organization trustor request timeout response has a 3xx status code
func (o *DeleteExternalcontactsOrganizationTrustorRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts organization trustor request timeout response has a 4xx status code
func (o *DeleteExternalcontactsOrganizationTrustorRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts organization trustor request timeout response has a 5xx status code
func (o *DeleteExternalcontactsOrganizationTrustorRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts organization trustor request timeout response a status code equal to that given
func (o *DeleteExternalcontactsOrganizationTrustorRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteExternalcontactsOrganizationTrustorRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteExternalcontactsOrganizationTrustorRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteExternalcontactsOrganizationTrustorRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsOrganizationTrustorRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsOrganizationTrustorRequestEntityTooLarge creates a DeleteExternalcontactsOrganizationTrustorRequestEntityTooLarge with default headers values
func NewDeleteExternalcontactsOrganizationTrustorRequestEntityTooLarge() *DeleteExternalcontactsOrganizationTrustorRequestEntityTooLarge {
	return &DeleteExternalcontactsOrganizationTrustorRequestEntityTooLarge{}
}

/*
DeleteExternalcontactsOrganizationTrustorRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteExternalcontactsOrganizationTrustorRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts organization trustor request entity too large response has a 2xx status code
func (o *DeleteExternalcontactsOrganizationTrustorRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts organization trustor request entity too large response has a 3xx status code
func (o *DeleteExternalcontactsOrganizationTrustorRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts organization trustor request entity too large response has a 4xx status code
func (o *DeleteExternalcontactsOrganizationTrustorRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts organization trustor request entity too large response has a 5xx status code
func (o *DeleteExternalcontactsOrganizationTrustorRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts organization trustor request entity too large response a status code equal to that given
func (o *DeleteExternalcontactsOrganizationTrustorRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteExternalcontactsOrganizationTrustorRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteExternalcontactsOrganizationTrustorRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteExternalcontactsOrganizationTrustorRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsOrganizationTrustorRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsOrganizationTrustorUnsupportedMediaType creates a DeleteExternalcontactsOrganizationTrustorUnsupportedMediaType with default headers values
func NewDeleteExternalcontactsOrganizationTrustorUnsupportedMediaType() *DeleteExternalcontactsOrganizationTrustorUnsupportedMediaType {
	return &DeleteExternalcontactsOrganizationTrustorUnsupportedMediaType{}
}

/*
DeleteExternalcontactsOrganizationTrustorUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteExternalcontactsOrganizationTrustorUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts organization trustor unsupported media type response has a 2xx status code
func (o *DeleteExternalcontactsOrganizationTrustorUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts organization trustor unsupported media type response has a 3xx status code
func (o *DeleteExternalcontactsOrganizationTrustorUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts organization trustor unsupported media type response has a 4xx status code
func (o *DeleteExternalcontactsOrganizationTrustorUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts organization trustor unsupported media type response has a 5xx status code
func (o *DeleteExternalcontactsOrganizationTrustorUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts organization trustor unsupported media type response a status code equal to that given
func (o *DeleteExternalcontactsOrganizationTrustorUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteExternalcontactsOrganizationTrustorUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteExternalcontactsOrganizationTrustorUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteExternalcontactsOrganizationTrustorUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsOrganizationTrustorUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsOrganizationTrustorTooManyRequests creates a DeleteExternalcontactsOrganizationTrustorTooManyRequests with default headers values
func NewDeleteExternalcontactsOrganizationTrustorTooManyRequests() *DeleteExternalcontactsOrganizationTrustorTooManyRequests {
	return &DeleteExternalcontactsOrganizationTrustorTooManyRequests{}
}

/*
DeleteExternalcontactsOrganizationTrustorTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteExternalcontactsOrganizationTrustorTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts organization trustor too many requests response has a 2xx status code
func (o *DeleteExternalcontactsOrganizationTrustorTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts organization trustor too many requests response has a 3xx status code
func (o *DeleteExternalcontactsOrganizationTrustorTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts organization trustor too many requests response has a 4xx status code
func (o *DeleteExternalcontactsOrganizationTrustorTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts organization trustor too many requests response has a 5xx status code
func (o *DeleteExternalcontactsOrganizationTrustorTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts organization trustor too many requests response a status code equal to that given
func (o *DeleteExternalcontactsOrganizationTrustorTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteExternalcontactsOrganizationTrustorTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteExternalcontactsOrganizationTrustorTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteExternalcontactsOrganizationTrustorTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsOrganizationTrustorTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsOrganizationTrustorInternalServerError creates a DeleteExternalcontactsOrganizationTrustorInternalServerError with default headers values
func NewDeleteExternalcontactsOrganizationTrustorInternalServerError() *DeleteExternalcontactsOrganizationTrustorInternalServerError {
	return &DeleteExternalcontactsOrganizationTrustorInternalServerError{}
}

/*
DeleteExternalcontactsOrganizationTrustorInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteExternalcontactsOrganizationTrustorInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts organization trustor internal server error response has a 2xx status code
func (o *DeleteExternalcontactsOrganizationTrustorInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts organization trustor internal server error response has a 3xx status code
func (o *DeleteExternalcontactsOrganizationTrustorInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts organization trustor internal server error response has a 4xx status code
func (o *DeleteExternalcontactsOrganizationTrustorInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete externalcontacts organization trustor internal server error response has a 5xx status code
func (o *DeleteExternalcontactsOrganizationTrustorInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete externalcontacts organization trustor internal server error response a status code equal to that given
func (o *DeleteExternalcontactsOrganizationTrustorInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteExternalcontactsOrganizationTrustorInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteExternalcontactsOrganizationTrustorInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteExternalcontactsOrganizationTrustorInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsOrganizationTrustorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsOrganizationTrustorServiceUnavailable creates a DeleteExternalcontactsOrganizationTrustorServiceUnavailable with default headers values
func NewDeleteExternalcontactsOrganizationTrustorServiceUnavailable() *DeleteExternalcontactsOrganizationTrustorServiceUnavailable {
	return &DeleteExternalcontactsOrganizationTrustorServiceUnavailable{}
}

/*
DeleteExternalcontactsOrganizationTrustorServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteExternalcontactsOrganizationTrustorServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts organization trustor service unavailable response has a 2xx status code
func (o *DeleteExternalcontactsOrganizationTrustorServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts organization trustor service unavailable response has a 3xx status code
func (o *DeleteExternalcontactsOrganizationTrustorServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts organization trustor service unavailable response has a 4xx status code
func (o *DeleteExternalcontactsOrganizationTrustorServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete externalcontacts organization trustor service unavailable response has a 5xx status code
func (o *DeleteExternalcontactsOrganizationTrustorServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete externalcontacts organization trustor service unavailable response a status code equal to that given
func (o *DeleteExternalcontactsOrganizationTrustorServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteExternalcontactsOrganizationTrustorServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteExternalcontactsOrganizationTrustorServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteExternalcontactsOrganizationTrustorServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsOrganizationTrustorServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsOrganizationTrustorGatewayTimeout creates a DeleteExternalcontactsOrganizationTrustorGatewayTimeout with default headers values
func NewDeleteExternalcontactsOrganizationTrustorGatewayTimeout() *DeleteExternalcontactsOrganizationTrustorGatewayTimeout {
	return &DeleteExternalcontactsOrganizationTrustorGatewayTimeout{}
}

/*
DeleteExternalcontactsOrganizationTrustorGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteExternalcontactsOrganizationTrustorGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts organization trustor gateway timeout response has a 2xx status code
func (o *DeleteExternalcontactsOrganizationTrustorGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts organization trustor gateway timeout response has a 3xx status code
func (o *DeleteExternalcontactsOrganizationTrustorGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts organization trustor gateway timeout response has a 4xx status code
func (o *DeleteExternalcontactsOrganizationTrustorGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete externalcontacts organization trustor gateway timeout response has a 5xx status code
func (o *DeleteExternalcontactsOrganizationTrustorGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete externalcontacts organization trustor gateway timeout response a status code equal to that given
func (o *DeleteExternalcontactsOrganizationTrustorGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteExternalcontactsOrganizationTrustorGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteExternalcontactsOrganizationTrustorGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor][%d] deleteExternalcontactsOrganizationTrustorGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteExternalcontactsOrganizationTrustorGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsOrganizationTrustorGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
