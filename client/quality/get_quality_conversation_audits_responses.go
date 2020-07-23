// Code generated by go-swagger; DO NOT EDIT.

package quality

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetQualityConversationAuditsReader is a Reader for the GetQualityConversationAudits structure.
type GetQualityConversationAuditsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQualityConversationAuditsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQualityConversationAuditsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetQualityConversationAuditsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetQualityConversationAuditsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetQualityConversationAuditsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetQualityConversationAuditsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetQualityConversationAuditsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetQualityConversationAuditsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetQualityConversationAuditsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetQualityConversationAuditsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetQualityConversationAuditsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetQualityConversationAuditsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetQualityConversationAuditsOK creates a GetQualityConversationAuditsOK with default headers values
func NewGetQualityConversationAuditsOK() *GetQualityConversationAuditsOK {
	return &GetQualityConversationAuditsOK{}
}

/*GetQualityConversationAuditsOK handles this case with default header values.

successful operation
*/
type GetQualityConversationAuditsOK struct {
	Payload *models.QualityAuditPage
}

func (o *GetQualityConversationAuditsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/{conversationId}/audits][%d] getQualityConversationAuditsOK  %+v", 200, o.Payload)
}

func (o *GetQualityConversationAuditsOK) GetPayload() *models.QualityAuditPage {
	return o.Payload
}

func (o *GetQualityConversationAuditsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QualityAuditPage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityConversationAuditsBadRequest creates a GetQualityConversationAuditsBadRequest with default headers values
func NewGetQualityConversationAuditsBadRequest() *GetQualityConversationAuditsBadRequest {
	return &GetQualityConversationAuditsBadRequest{}
}

/*GetQualityConversationAuditsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetQualityConversationAuditsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetQualityConversationAuditsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/{conversationId}/audits][%d] getQualityConversationAuditsBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityConversationAuditsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityConversationAuditsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityConversationAuditsUnauthorized creates a GetQualityConversationAuditsUnauthorized with default headers values
func NewGetQualityConversationAuditsUnauthorized() *GetQualityConversationAuditsUnauthorized {
	return &GetQualityConversationAuditsUnauthorized{}
}

/*GetQualityConversationAuditsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetQualityConversationAuditsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetQualityConversationAuditsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/{conversationId}/audits][%d] getQualityConversationAuditsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityConversationAuditsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityConversationAuditsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityConversationAuditsForbidden creates a GetQualityConversationAuditsForbidden with default headers values
func NewGetQualityConversationAuditsForbidden() *GetQualityConversationAuditsForbidden {
	return &GetQualityConversationAuditsForbidden{}
}

/*GetQualityConversationAuditsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetQualityConversationAuditsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetQualityConversationAuditsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/{conversationId}/audits][%d] getQualityConversationAuditsForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityConversationAuditsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityConversationAuditsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityConversationAuditsNotFound creates a GetQualityConversationAuditsNotFound with default headers values
func NewGetQualityConversationAuditsNotFound() *GetQualityConversationAuditsNotFound {
	return &GetQualityConversationAuditsNotFound{}
}

/*GetQualityConversationAuditsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetQualityConversationAuditsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetQualityConversationAuditsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/{conversationId}/audits][%d] getQualityConversationAuditsNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityConversationAuditsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityConversationAuditsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityConversationAuditsRequestEntityTooLarge creates a GetQualityConversationAuditsRequestEntityTooLarge with default headers values
func NewGetQualityConversationAuditsRequestEntityTooLarge() *GetQualityConversationAuditsRequestEntityTooLarge {
	return &GetQualityConversationAuditsRequestEntityTooLarge{}
}

/*GetQualityConversationAuditsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetQualityConversationAuditsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetQualityConversationAuditsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/{conversationId}/audits][%d] getQualityConversationAuditsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityConversationAuditsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityConversationAuditsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityConversationAuditsUnsupportedMediaType creates a GetQualityConversationAuditsUnsupportedMediaType with default headers values
func NewGetQualityConversationAuditsUnsupportedMediaType() *GetQualityConversationAuditsUnsupportedMediaType {
	return &GetQualityConversationAuditsUnsupportedMediaType{}
}

/*GetQualityConversationAuditsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetQualityConversationAuditsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetQualityConversationAuditsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/{conversationId}/audits][%d] getQualityConversationAuditsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityConversationAuditsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityConversationAuditsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityConversationAuditsTooManyRequests creates a GetQualityConversationAuditsTooManyRequests with default headers values
func NewGetQualityConversationAuditsTooManyRequests() *GetQualityConversationAuditsTooManyRequests {
	return &GetQualityConversationAuditsTooManyRequests{}
}

/*GetQualityConversationAuditsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetQualityConversationAuditsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetQualityConversationAuditsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/{conversationId}/audits][%d] getQualityConversationAuditsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityConversationAuditsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityConversationAuditsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityConversationAuditsInternalServerError creates a GetQualityConversationAuditsInternalServerError with default headers values
func NewGetQualityConversationAuditsInternalServerError() *GetQualityConversationAuditsInternalServerError {
	return &GetQualityConversationAuditsInternalServerError{}
}

/*GetQualityConversationAuditsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetQualityConversationAuditsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetQualityConversationAuditsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/{conversationId}/audits][%d] getQualityConversationAuditsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityConversationAuditsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityConversationAuditsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityConversationAuditsServiceUnavailable creates a GetQualityConversationAuditsServiceUnavailable with default headers values
func NewGetQualityConversationAuditsServiceUnavailable() *GetQualityConversationAuditsServiceUnavailable {
	return &GetQualityConversationAuditsServiceUnavailable{}
}

/*GetQualityConversationAuditsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetQualityConversationAuditsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetQualityConversationAuditsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/{conversationId}/audits][%d] getQualityConversationAuditsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityConversationAuditsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityConversationAuditsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityConversationAuditsGatewayTimeout creates a GetQualityConversationAuditsGatewayTimeout with default headers values
func NewGetQualityConversationAuditsGatewayTimeout() *GetQualityConversationAuditsGatewayTimeout {
	return &GetQualityConversationAuditsGatewayTimeout{}
}

/*GetQualityConversationAuditsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetQualityConversationAuditsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetQualityConversationAuditsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/{conversationId}/audits][%d] getQualityConversationAuditsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityConversationAuditsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityConversationAuditsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
