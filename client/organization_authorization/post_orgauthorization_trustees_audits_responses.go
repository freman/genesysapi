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

// PostOrgauthorizationTrusteesAuditsReader is a Reader for the PostOrgauthorizationTrusteesAudits structure.
type PostOrgauthorizationTrusteesAuditsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOrgauthorizationTrusteesAuditsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOrgauthorizationTrusteesAuditsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOrgauthorizationTrusteesAuditsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostOrgauthorizationTrusteesAuditsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostOrgauthorizationTrusteesAuditsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostOrgauthorizationTrusteesAuditsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostOrgauthorizationTrusteesAuditsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostOrgauthorizationTrusteesAuditsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostOrgauthorizationTrusteesAuditsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOrgauthorizationTrusteesAuditsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostOrgauthorizationTrusteesAuditsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostOrgauthorizationTrusteesAuditsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOrgauthorizationTrusteesAuditsOK creates a PostOrgauthorizationTrusteesAuditsOK with default headers values
func NewPostOrgauthorizationTrusteesAuditsOK() *PostOrgauthorizationTrusteesAuditsOK {
	return &PostOrgauthorizationTrusteesAuditsOK{}
}

/*PostOrgauthorizationTrusteesAuditsOK handles this case with default header values.

successful operation
*/
type PostOrgauthorizationTrusteesAuditsOK struct {
	Payload models.AuditQueryResponse
}

func (o *PostOrgauthorizationTrusteesAuditsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/audits][%d] postOrgauthorizationTrusteesAuditsOK  %+v", 200, o.Payload)
}

func (o *PostOrgauthorizationTrusteesAuditsOK) GetPayload() models.AuditQueryResponse {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteesAuditsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteesAuditsBadRequest creates a PostOrgauthorizationTrusteesAuditsBadRequest with default headers values
func NewPostOrgauthorizationTrusteesAuditsBadRequest() *PostOrgauthorizationTrusteesAuditsBadRequest {
	return &PostOrgauthorizationTrusteesAuditsBadRequest{}
}

/*PostOrgauthorizationTrusteesAuditsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostOrgauthorizationTrusteesAuditsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrusteesAuditsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/audits][%d] postOrgauthorizationTrusteesAuditsBadRequest  %+v", 400, o.Payload)
}

func (o *PostOrgauthorizationTrusteesAuditsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteesAuditsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteesAuditsUnauthorized creates a PostOrgauthorizationTrusteesAuditsUnauthorized with default headers values
func NewPostOrgauthorizationTrusteesAuditsUnauthorized() *PostOrgauthorizationTrusteesAuditsUnauthorized {
	return &PostOrgauthorizationTrusteesAuditsUnauthorized{}
}

/*PostOrgauthorizationTrusteesAuditsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostOrgauthorizationTrusteesAuditsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrusteesAuditsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/audits][%d] postOrgauthorizationTrusteesAuditsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOrgauthorizationTrusteesAuditsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteesAuditsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteesAuditsForbidden creates a PostOrgauthorizationTrusteesAuditsForbidden with default headers values
func NewPostOrgauthorizationTrusteesAuditsForbidden() *PostOrgauthorizationTrusteesAuditsForbidden {
	return &PostOrgauthorizationTrusteesAuditsForbidden{}
}

/*PostOrgauthorizationTrusteesAuditsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostOrgauthorizationTrusteesAuditsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrusteesAuditsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/audits][%d] postOrgauthorizationTrusteesAuditsForbidden  %+v", 403, o.Payload)
}

func (o *PostOrgauthorizationTrusteesAuditsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteesAuditsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteesAuditsNotFound creates a PostOrgauthorizationTrusteesAuditsNotFound with default headers values
func NewPostOrgauthorizationTrusteesAuditsNotFound() *PostOrgauthorizationTrusteesAuditsNotFound {
	return &PostOrgauthorizationTrusteesAuditsNotFound{}
}

/*PostOrgauthorizationTrusteesAuditsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostOrgauthorizationTrusteesAuditsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrusteesAuditsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/audits][%d] postOrgauthorizationTrusteesAuditsNotFound  %+v", 404, o.Payload)
}

func (o *PostOrgauthorizationTrusteesAuditsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteesAuditsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteesAuditsRequestEntityTooLarge creates a PostOrgauthorizationTrusteesAuditsRequestEntityTooLarge with default headers values
func NewPostOrgauthorizationTrusteesAuditsRequestEntityTooLarge() *PostOrgauthorizationTrusteesAuditsRequestEntityTooLarge {
	return &PostOrgauthorizationTrusteesAuditsRequestEntityTooLarge{}
}

/*PostOrgauthorizationTrusteesAuditsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostOrgauthorizationTrusteesAuditsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrusteesAuditsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/audits][%d] postOrgauthorizationTrusteesAuditsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOrgauthorizationTrusteesAuditsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteesAuditsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteesAuditsUnsupportedMediaType creates a PostOrgauthorizationTrusteesAuditsUnsupportedMediaType with default headers values
func NewPostOrgauthorizationTrusteesAuditsUnsupportedMediaType() *PostOrgauthorizationTrusteesAuditsUnsupportedMediaType {
	return &PostOrgauthorizationTrusteesAuditsUnsupportedMediaType{}
}

/*PostOrgauthorizationTrusteesAuditsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostOrgauthorizationTrusteesAuditsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrusteesAuditsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/audits][%d] postOrgauthorizationTrusteesAuditsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOrgauthorizationTrusteesAuditsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteesAuditsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteesAuditsTooManyRequests creates a PostOrgauthorizationTrusteesAuditsTooManyRequests with default headers values
func NewPostOrgauthorizationTrusteesAuditsTooManyRequests() *PostOrgauthorizationTrusteesAuditsTooManyRequests {
	return &PostOrgauthorizationTrusteesAuditsTooManyRequests{}
}

/*PostOrgauthorizationTrusteesAuditsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostOrgauthorizationTrusteesAuditsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrusteesAuditsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/audits][%d] postOrgauthorizationTrusteesAuditsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOrgauthorizationTrusteesAuditsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteesAuditsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteesAuditsInternalServerError creates a PostOrgauthorizationTrusteesAuditsInternalServerError with default headers values
func NewPostOrgauthorizationTrusteesAuditsInternalServerError() *PostOrgauthorizationTrusteesAuditsInternalServerError {
	return &PostOrgauthorizationTrusteesAuditsInternalServerError{}
}

/*PostOrgauthorizationTrusteesAuditsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostOrgauthorizationTrusteesAuditsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrusteesAuditsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/audits][%d] postOrgauthorizationTrusteesAuditsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOrgauthorizationTrusteesAuditsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteesAuditsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteesAuditsServiceUnavailable creates a PostOrgauthorizationTrusteesAuditsServiceUnavailable with default headers values
func NewPostOrgauthorizationTrusteesAuditsServiceUnavailable() *PostOrgauthorizationTrusteesAuditsServiceUnavailable {
	return &PostOrgauthorizationTrusteesAuditsServiceUnavailable{}
}

/*PostOrgauthorizationTrusteesAuditsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostOrgauthorizationTrusteesAuditsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrusteesAuditsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/audits][%d] postOrgauthorizationTrusteesAuditsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOrgauthorizationTrusteesAuditsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteesAuditsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteesAuditsGatewayTimeout creates a PostOrgauthorizationTrusteesAuditsGatewayTimeout with default headers values
func NewPostOrgauthorizationTrusteesAuditsGatewayTimeout() *PostOrgauthorizationTrusteesAuditsGatewayTimeout {
	return &PostOrgauthorizationTrusteesAuditsGatewayTimeout{}
}

/*PostOrgauthorizationTrusteesAuditsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostOrgauthorizationTrusteesAuditsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrusteesAuditsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/audits][%d] postOrgauthorizationTrusteesAuditsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOrgauthorizationTrusteesAuditsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteesAuditsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
