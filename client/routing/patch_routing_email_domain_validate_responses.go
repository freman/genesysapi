// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PatchRoutingEmailDomainValidateReader is a Reader for the PatchRoutingEmailDomainValidate structure.
type PatchRoutingEmailDomainValidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchRoutingEmailDomainValidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchRoutingEmailDomainValidateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchRoutingEmailDomainValidateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchRoutingEmailDomainValidateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchRoutingEmailDomainValidateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchRoutingEmailDomainValidateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchRoutingEmailDomainValidateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchRoutingEmailDomainValidateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchRoutingEmailDomainValidateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchRoutingEmailDomainValidateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchRoutingEmailDomainValidateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchRoutingEmailDomainValidateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchRoutingEmailDomainValidateOK creates a PatchRoutingEmailDomainValidateOK with default headers values
func NewPatchRoutingEmailDomainValidateOK() *PatchRoutingEmailDomainValidateOK {
	return &PatchRoutingEmailDomainValidateOK{}
}

/*PatchRoutingEmailDomainValidateOK handles this case with default header values.

successful operation
*/
type PatchRoutingEmailDomainValidateOK struct {
	Payload *models.InboundDomain
}

func (o *PatchRoutingEmailDomainValidateOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/email/domains/{domainId}/validate][%d] patchRoutingEmailDomainValidateOK  %+v", 200, o.Payload)
}

func (o *PatchRoutingEmailDomainValidateOK) GetPayload() *models.InboundDomain {
	return o.Payload
}

func (o *PatchRoutingEmailDomainValidateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InboundDomain)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingEmailDomainValidateBadRequest creates a PatchRoutingEmailDomainValidateBadRequest with default headers values
func NewPatchRoutingEmailDomainValidateBadRequest() *PatchRoutingEmailDomainValidateBadRequest {
	return &PatchRoutingEmailDomainValidateBadRequest{}
}

/*PatchRoutingEmailDomainValidateBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchRoutingEmailDomainValidateBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingEmailDomainValidateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/email/domains/{domainId}/validate][%d] patchRoutingEmailDomainValidateBadRequest  %+v", 400, o.Payload)
}

func (o *PatchRoutingEmailDomainValidateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingEmailDomainValidateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingEmailDomainValidateUnauthorized creates a PatchRoutingEmailDomainValidateUnauthorized with default headers values
func NewPatchRoutingEmailDomainValidateUnauthorized() *PatchRoutingEmailDomainValidateUnauthorized {
	return &PatchRoutingEmailDomainValidateUnauthorized{}
}

/*PatchRoutingEmailDomainValidateUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchRoutingEmailDomainValidateUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingEmailDomainValidateUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/email/domains/{domainId}/validate][%d] patchRoutingEmailDomainValidateUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchRoutingEmailDomainValidateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingEmailDomainValidateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingEmailDomainValidateForbidden creates a PatchRoutingEmailDomainValidateForbidden with default headers values
func NewPatchRoutingEmailDomainValidateForbidden() *PatchRoutingEmailDomainValidateForbidden {
	return &PatchRoutingEmailDomainValidateForbidden{}
}

/*PatchRoutingEmailDomainValidateForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchRoutingEmailDomainValidateForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingEmailDomainValidateForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/email/domains/{domainId}/validate][%d] patchRoutingEmailDomainValidateForbidden  %+v", 403, o.Payload)
}

func (o *PatchRoutingEmailDomainValidateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingEmailDomainValidateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingEmailDomainValidateNotFound creates a PatchRoutingEmailDomainValidateNotFound with default headers values
func NewPatchRoutingEmailDomainValidateNotFound() *PatchRoutingEmailDomainValidateNotFound {
	return &PatchRoutingEmailDomainValidateNotFound{}
}

/*PatchRoutingEmailDomainValidateNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchRoutingEmailDomainValidateNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingEmailDomainValidateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/email/domains/{domainId}/validate][%d] patchRoutingEmailDomainValidateNotFound  %+v", 404, o.Payload)
}

func (o *PatchRoutingEmailDomainValidateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingEmailDomainValidateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingEmailDomainValidateRequestEntityTooLarge creates a PatchRoutingEmailDomainValidateRequestEntityTooLarge with default headers values
func NewPatchRoutingEmailDomainValidateRequestEntityTooLarge() *PatchRoutingEmailDomainValidateRequestEntityTooLarge {
	return &PatchRoutingEmailDomainValidateRequestEntityTooLarge{}
}

/*PatchRoutingEmailDomainValidateRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchRoutingEmailDomainValidateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingEmailDomainValidateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/email/domains/{domainId}/validate][%d] patchRoutingEmailDomainValidateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchRoutingEmailDomainValidateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingEmailDomainValidateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingEmailDomainValidateUnsupportedMediaType creates a PatchRoutingEmailDomainValidateUnsupportedMediaType with default headers values
func NewPatchRoutingEmailDomainValidateUnsupportedMediaType() *PatchRoutingEmailDomainValidateUnsupportedMediaType {
	return &PatchRoutingEmailDomainValidateUnsupportedMediaType{}
}

/*PatchRoutingEmailDomainValidateUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchRoutingEmailDomainValidateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingEmailDomainValidateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/email/domains/{domainId}/validate][%d] patchRoutingEmailDomainValidateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchRoutingEmailDomainValidateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingEmailDomainValidateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingEmailDomainValidateTooManyRequests creates a PatchRoutingEmailDomainValidateTooManyRequests with default headers values
func NewPatchRoutingEmailDomainValidateTooManyRequests() *PatchRoutingEmailDomainValidateTooManyRequests {
	return &PatchRoutingEmailDomainValidateTooManyRequests{}
}

/*PatchRoutingEmailDomainValidateTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchRoutingEmailDomainValidateTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingEmailDomainValidateTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/email/domains/{domainId}/validate][%d] patchRoutingEmailDomainValidateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchRoutingEmailDomainValidateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingEmailDomainValidateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingEmailDomainValidateInternalServerError creates a PatchRoutingEmailDomainValidateInternalServerError with default headers values
func NewPatchRoutingEmailDomainValidateInternalServerError() *PatchRoutingEmailDomainValidateInternalServerError {
	return &PatchRoutingEmailDomainValidateInternalServerError{}
}

/*PatchRoutingEmailDomainValidateInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchRoutingEmailDomainValidateInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingEmailDomainValidateInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/email/domains/{domainId}/validate][%d] patchRoutingEmailDomainValidateInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchRoutingEmailDomainValidateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingEmailDomainValidateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingEmailDomainValidateServiceUnavailable creates a PatchRoutingEmailDomainValidateServiceUnavailable with default headers values
func NewPatchRoutingEmailDomainValidateServiceUnavailable() *PatchRoutingEmailDomainValidateServiceUnavailable {
	return &PatchRoutingEmailDomainValidateServiceUnavailable{}
}

/*PatchRoutingEmailDomainValidateServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchRoutingEmailDomainValidateServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingEmailDomainValidateServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/email/domains/{domainId}/validate][%d] patchRoutingEmailDomainValidateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchRoutingEmailDomainValidateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingEmailDomainValidateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingEmailDomainValidateGatewayTimeout creates a PatchRoutingEmailDomainValidateGatewayTimeout with default headers values
func NewPatchRoutingEmailDomainValidateGatewayTimeout() *PatchRoutingEmailDomainValidateGatewayTimeout {
	return &PatchRoutingEmailDomainValidateGatewayTimeout{}
}

/*PatchRoutingEmailDomainValidateGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchRoutingEmailDomainValidateGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingEmailDomainValidateGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/email/domains/{domainId}/validate][%d] patchRoutingEmailDomainValidateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchRoutingEmailDomainValidateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingEmailDomainValidateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
