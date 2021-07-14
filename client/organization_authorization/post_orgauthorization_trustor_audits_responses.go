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

// PostOrgauthorizationTrustorAuditsReader is a Reader for the PostOrgauthorizationTrustorAudits structure.
type PostOrgauthorizationTrustorAuditsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOrgauthorizationTrustorAuditsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOrgauthorizationTrustorAuditsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOrgauthorizationTrustorAuditsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostOrgauthorizationTrustorAuditsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostOrgauthorizationTrustorAuditsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostOrgauthorizationTrustorAuditsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostOrgauthorizationTrustorAuditsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostOrgauthorizationTrustorAuditsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostOrgauthorizationTrustorAuditsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostOrgauthorizationTrustorAuditsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOrgauthorizationTrustorAuditsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostOrgauthorizationTrustorAuditsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostOrgauthorizationTrustorAuditsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOrgauthorizationTrustorAuditsOK creates a PostOrgauthorizationTrustorAuditsOK with default headers values
func NewPostOrgauthorizationTrustorAuditsOK() *PostOrgauthorizationTrustorAuditsOK {
	return &PostOrgauthorizationTrustorAuditsOK{}
}

/*PostOrgauthorizationTrustorAuditsOK handles this case with default header values.

successful operation
*/
type PostOrgauthorizationTrustorAuditsOK struct {
	Payload models.AuditQueryResponse
}

func (o *PostOrgauthorizationTrustorAuditsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustor/audits][%d] postOrgauthorizationTrustorAuditsOK  %+v", 200, o.Payload)
}

func (o *PostOrgauthorizationTrustorAuditsOK) GetPayload() models.AuditQueryResponse {
	return o.Payload
}

func (o *PostOrgauthorizationTrustorAuditsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrustorAuditsBadRequest creates a PostOrgauthorizationTrustorAuditsBadRequest with default headers values
func NewPostOrgauthorizationTrustorAuditsBadRequest() *PostOrgauthorizationTrustorAuditsBadRequest {
	return &PostOrgauthorizationTrustorAuditsBadRequest{}
}

/*PostOrgauthorizationTrustorAuditsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostOrgauthorizationTrustorAuditsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrustorAuditsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustor/audits][%d] postOrgauthorizationTrustorAuditsBadRequest  %+v", 400, o.Payload)
}

func (o *PostOrgauthorizationTrustorAuditsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrustorAuditsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrustorAuditsUnauthorized creates a PostOrgauthorizationTrustorAuditsUnauthorized with default headers values
func NewPostOrgauthorizationTrustorAuditsUnauthorized() *PostOrgauthorizationTrustorAuditsUnauthorized {
	return &PostOrgauthorizationTrustorAuditsUnauthorized{}
}

/*PostOrgauthorizationTrustorAuditsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostOrgauthorizationTrustorAuditsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrustorAuditsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustor/audits][%d] postOrgauthorizationTrustorAuditsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOrgauthorizationTrustorAuditsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrustorAuditsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrustorAuditsForbidden creates a PostOrgauthorizationTrustorAuditsForbidden with default headers values
func NewPostOrgauthorizationTrustorAuditsForbidden() *PostOrgauthorizationTrustorAuditsForbidden {
	return &PostOrgauthorizationTrustorAuditsForbidden{}
}

/*PostOrgauthorizationTrustorAuditsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostOrgauthorizationTrustorAuditsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrustorAuditsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustor/audits][%d] postOrgauthorizationTrustorAuditsForbidden  %+v", 403, o.Payload)
}

func (o *PostOrgauthorizationTrustorAuditsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrustorAuditsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrustorAuditsNotFound creates a PostOrgauthorizationTrustorAuditsNotFound with default headers values
func NewPostOrgauthorizationTrustorAuditsNotFound() *PostOrgauthorizationTrustorAuditsNotFound {
	return &PostOrgauthorizationTrustorAuditsNotFound{}
}

/*PostOrgauthorizationTrustorAuditsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostOrgauthorizationTrustorAuditsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrustorAuditsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustor/audits][%d] postOrgauthorizationTrustorAuditsNotFound  %+v", 404, o.Payload)
}

func (o *PostOrgauthorizationTrustorAuditsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrustorAuditsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrustorAuditsRequestTimeout creates a PostOrgauthorizationTrustorAuditsRequestTimeout with default headers values
func NewPostOrgauthorizationTrustorAuditsRequestTimeout() *PostOrgauthorizationTrustorAuditsRequestTimeout {
	return &PostOrgauthorizationTrustorAuditsRequestTimeout{}
}

/*PostOrgauthorizationTrustorAuditsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostOrgauthorizationTrustorAuditsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrustorAuditsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustor/audits][%d] postOrgauthorizationTrustorAuditsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOrgauthorizationTrustorAuditsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrustorAuditsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrustorAuditsRequestEntityTooLarge creates a PostOrgauthorizationTrustorAuditsRequestEntityTooLarge with default headers values
func NewPostOrgauthorizationTrustorAuditsRequestEntityTooLarge() *PostOrgauthorizationTrustorAuditsRequestEntityTooLarge {
	return &PostOrgauthorizationTrustorAuditsRequestEntityTooLarge{}
}

/*PostOrgauthorizationTrustorAuditsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostOrgauthorizationTrustorAuditsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrustorAuditsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustor/audits][%d] postOrgauthorizationTrustorAuditsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOrgauthorizationTrustorAuditsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrustorAuditsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrustorAuditsUnsupportedMediaType creates a PostOrgauthorizationTrustorAuditsUnsupportedMediaType with default headers values
func NewPostOrgauthorizationTrustorAuditsUnsupportedMediaType() *PostOrgauthorizationTrustorAuditsUnsupportedMediaType {
	return &PostOrgauthorizationTrustorAuditsUnsupportedMediaType{}
}

/*PostOrgauthorizationTrustorAuditsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostOrgauthorizationTrustorAuditsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrustorAuditsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustor/audits][%d] postOrgauthorizationTrustorAuditsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOrgauthorizationTrustorAuditsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrustorAuditsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrustorAuditsTooManyRequests creates a PostOrgauthorizationTrustorAuditsTooManyRequests with default headers values
func NewPostOrgauthorizationTrustorAuditsTooManyRequests() *PostOrgauthorizationTrustorAuditsTooManyRequests {
	return &PostOrgauthorizationTrustorAuditsTooManyRequests{}
}

/*PostOrgauthorizationTrustorAuditsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostOrgauthorizationTrustorAuditsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrustorAuditsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustor/audits][%d] postOrgauthorizationTrustorAuditsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOrgauthorizationTrustorAuditsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrustorAuditsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrustorAuditsInternalServerError creates a PostOrgauthorizationTrustorAuditsInternalServerError with default headers values
func NewPostOrgauthorizationTrustorAuditsInternalServerError() *PostOrgauthorizationTrustorAuditsInternalServerError {
	return &PostOrgauthorizationTrustorAuditsInternalServerError{}
}

/*PostOrgauthorizationTrustorAuditsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostOrgauthorizationTrustorAuditsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrustorAuditsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustor/audits][%d] postOrgauthorizationTrustorAuditsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOrgauthorizationTrustorAuditsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrustorAuditsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrustorAuditsServiceUnavailable creates a PostOrgauthorizationTrustorAuditsServiceUnavailable with default headers values
func NewPostOrgauthorizationTrustorAuditsServiceUnavailable() *PostOrgauthorizationTrustorAuditsServiceUnavailable {
	return &PostOrgauthorizationTrustorAuditsServiceUnavailable{}
}

/*PostOrgauthorizationTrustorAuditsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostOrgauthorizationTrustorAuditsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrustorAuditsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustor/audits][%d] postOrgauthorizationTrustorAuditsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOrgauthorizationTrustorAuditsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrustorAuditsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrustorAuditsGatewayTimeout creates a PostOrgauthorizationTrustorAuditsGatewayTimeout with default headers values
func NewPostOrgauthorizationTrustorAuditsGatewayTimeout() *PostOrgauthorizationTrustorAuditsGatewayTimeout {
	return &PostOrgauthorizationTrustorAuditsGatewayTimeout{}
}

/*PostOrgauthorizationTrustorAuditsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostOrgauthorizationTrustorAuditsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostOrgauthorizationTrustorAuditsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustor/audits][%d] postOrgauthorizationTrustorAuditsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOrgauthorizationTrustorAuditsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrustorAuditsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
