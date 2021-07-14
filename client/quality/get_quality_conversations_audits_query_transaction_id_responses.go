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

// GetQualityConversationsAuditsQueryTransactionIDReader is a Reader for the GetQualityConversationsAuditsQueryTransactionID structure.
type GetQualityConversationsAuditsQueryTransactionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQualityConversationsAuditsQueryTransactionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQualityConversationsAuditsQueryTransactionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewGetQualityConversationsAuditsQueryTransactionIDAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetQualityConversationsAuditsQueryTransactionIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetQualityConversationsAuditsQueryTransactionIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetQualityConversationsAuditsQueryTransactionIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetQualityConversationsAuditsQueryTransactionIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetQualityConversationsAuditsQueryTransactionIDRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetQualityConversationsAuditsQueryTransactionIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetQualityConversationsAuditsQueryTransactionIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetQualityConversationsAuditsQueryTransactionIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetQualityConversationsAuditsQueryTransactionIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetQualityConversationsAuditsQueryTransactionIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetQualityConversationsAuditsQueryTransactionIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetQualityConversationsAuditsQueryTransactionIDOK creates a GetQualityConversationsAuditsQueryTransactionIDOK with default headers values
func NewGetQualityConversationsAuditsQueryTransactionIDOK() *GetQualityConversationsAuditsQueryTransactionIDOK {
	return &GetQualityConversationsAuditsQueryTransactionIDOK{}
}

/*GetQualityConversationsAuditsQueryTransactionIDOK handles this case with default header values.

Query execution completed.
*/
type GetQualityConversationsAuditsQueryTransactionIDOK struct {
	Payload *models.QualityAuditQueryExecutionStatusResponse
}

func (o *GetQualityConversationsAuditsQueryTransactionIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/audits/query/{transactionId}][%d] getQualityConversationsAuditsQueryTransactionIdOK  %+v", 200, o.Payload)
}

func (o *GetQualityConversationsAuditsQueryTransactionIDOK) GetPayload() *models.QualityAuditQueryExecutionStatusResponse {
	return o.Payload
}

func (o *GetQualityConversationsAuditsQueryTransactionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QualityAuditQueryExecutionStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityConversationsAuditsQueryTransactionIDAccepted creates a GetQualityConversationsAuditsQueryTransactionIDAccepted with default headers values
func NewGetQualityConversationsAuditsQueryTransactionIDAccepted() *GetQualityConversationsAuditsQueryTransactionIDAccepted {
	return &GetQualityConversationsAuditsQueryTransactionIDAccepted{}
}

/*GetQualityConversationsAuditsQueryTransactionIDAccepted handles this case with default header values.

In progress - Query execution is in progress.
*/
type GetQualityConversationsAuditsQueryTransactionIDAccepted struct {
	Payload *models.QualityAuditQueryExecutionStatusResponse
}

func (o *GetQualityConversationsAuditsQueryTransactionIDAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/audits/query/{transactionId}][%d] getQualityConversationsAuditsQueryTransactionIdAccepted  %+v", 202, o.Payload)
}

func (o *GetQualityConversationsAuditsQueryTransactionIDAccepted) GetPayload() *models.QualityAuditQueryExecutionStatusResponse {
	return o.Payload
}

func (o *GetQualityConversationsAuditsQueryTransactionIDAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QualityAuditQueryExecutionStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityConversationsAuditsQueryTransactionIDBadRequest creates a GetQualityConversationsAuditsQueryTransactionIDBadRequest with default headers values
func NewGetQualityConversationsAuditsQueryTransactionIDBadRequest() *GetQualityConversationsAuditsQueryTransactionIDBadRequest {
	return &GetQualityConversationsAuditsQueryTransactionIDBadRequest{}
}

/*GetQualityConversationsAuditsQueryTransactionIDBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetQualityConversationsAuditsQueryTransactionIDBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetQualityConversationsAuditsQueryTransactionIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/audits/query/{transactionId}][%d] getQualityConversationsAuditsQueryTransactionIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityConversationsAuditsQueryTransactionIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityConversationsAuditsQueryTransactionIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityConversationsAuditsQueryTransactionIDUnauthorized creates a GetQualityConversationsAuditsQueryTransactionIDUnauthorized with default headers values
func NewGetQualityConversationsAuditsQueryTransactionIDUnauthorized() *GetQualityConversationsAuditsQueryTransactionIDUnauthorized {
	return &GetQualityConversationsAuditsQueryTransactionIDUnauthorized{}
}

/*GetQualityConversationsAuditsQueryTransactionIDUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetQualityConversationsAuditsQueryTransactionIDUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetQualityConversationsAuditsQueryTransactionIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/audits/query/{transactionId}][%d] getQualityConversationsAuditsQueryTransactionIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityConversationsAuditsQueryTransactionIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityConversationsAuditsQueryTransactionIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityConversationsAuditsQueryTransactionIDForbidden creates a GetQualityConversationsAuditsQueryTransactionIDForbidden with default headers values
func NewGetQualityConversationsAuditsQueryTransactionIDForbidden() *GetQualityConversationsAuditsQueryTransactionIDForbidden {
	return &GetQualityConversationsAuditsQueryTransactionIDForbidden{}
}

/*GetQualityConversationsAuditsQueryTransactionIDForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetQualityConversationsAuditsQueryTransactionIDForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetQualityConversationsAuditsQueryTransactionIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/audits/query/{transactionId}][%d] getQualityConversationsAuditsQueryTransactionIdForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityConversationsAuditsQueryTransactionIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityConversationsAuditsQueryTransactionIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityConversationsAuditsQueryTransactionIDNotFound creates a GetQualityConversationsAuditsQueryTransactionIDNotFound with default headers values
func NewGetQualityConversationsAuditsQueryTransactionIDNotFound() *GetQualityConversationsAuditsQueryTransactionIDNotFound {
	return &GetQualityConversationsAuditsQueryTransactionIDNotFound{}
}

/*GetQualityConversationsAuditsQueryTransactionIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetQualityConversationsAuditsQueryTransactionIDNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetQualityConversationsAuditsQueryTransactionIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/audits/query/{transactionId}][%d] getQualityConversationsAuditsQueryTransactionIdNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityConversationsAuditsQueryTransactionIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityConversationsAuditsQueryTransactionIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityConversationsAuditsQueryTransactionIDRequestTimeout creates a GetQualityConversationsAuditsQueryTransactionIDRequestTimeout with default headers values
func NewGetQualityConversationsAuditsQueryTransactionIDRequestTimeout() *GetQualityConversationsAuditsQueryTransactionIDRequestTimeout {
	return &GetQualityConversationsAuditsQueryTransactionIDRequestTimeout{}
}

/*GetQualityConversationsAuditsQueryTransactionIDRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetQualityConversationsAuditsQueryTransactionIDRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetQualityConversationsAuditsQueryTransactionIDRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/audits/query/{transactionId}][%d] getQualityConversationsAuditsQueryTransactionIdRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualityConversationsAuditsQueryTransactionIDRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityConversationsAuditsQueryTransactionIDRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityConversationsAuditsQueryTransactionIDRequestEntityTooLarge creates a GetQualityConversationsAuditsQueryTransactionIDRequestEntityTooLarge with default headers values
func NewGetQualityConversationsAuditsQueryTransactionIDRequestEntityTooLarge() *GetQualityConversationsAuditsQueryTransactionIDRequestEntityTooLarge {
	return &GetQualityConversationsAuditsQueryTransactionIDRequestEntityTooLarge{}
}

/*GetQualityConversationsAuditsQueryTransactionIDRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetQualityConversationsAuditsQueryTransactionIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetQualityConversationsAuditsQueryTransactionIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/audits/query/{transactionId}][%d] getQualityConversationsAuditsQueryTransactionIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityConversationsAuditsQueryTransactionIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityConversationsAuditsQueryTransactionIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityConversationsAuditsQueryTransactionIDUnsupportedMediaType creates a GetQualityConversationsAuditsQueryTransactionIDUnsupportedMediaType with default headers values
func NewGetQualityConversationsAuditsQueryTransactionIDUnsupportedMediaType() *GetQualityConversationsAuditsQueryTransactionIDUnsupportedMediaType {
	return &GetQualityConversationsAuditsQueryTransactionIDUnsupportedMediaType{}
}

/*GetQualityConversationsAuditsQueryTransactionIDUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetQualityConversationsAuditsQueryTransactionIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetQualityConversationsAuditsQueryTransactionIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/audits/query/{transactionId}][%d] getQualityConversationsAuditsQueryTransactionIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityConversationsAuditsQueryTransactionIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityConversationsAuditsQueryTransactionIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityConversationsAuditsQueryTransactionIDTooManyRequests creates a GetQualityConversationsAuditsQueryTransactionIDTooManyRequests with default headers values
func NewGetQualityConversationsAuditsQueryTransactionIDTooManyRequests() *GetQualityConversationsAuditsQueryTransactionIDTooManyRequests {
	return &GetQualityConversationsAuditsQueryTransactionIDTooManyRequests{}
}

/*GetQualityConversationsAuditsQueryTransactionIDTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetQualityConversationsAuditsQueryTransactionIDTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetQualityConversationsAuditsQueryTransactionIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/audits/query/{transactionId}][%d] getQualityConversationsAuditsQueryTransactionIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityConversationsAuditsQueryTransactionIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityConversationsAuditsQueryTransactionIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityConversationsAuditsQueryTransactionIDInternalServerError creates a GetQualityConversationsAuditsQueryTransactionIDInternalServerError with default headers values
func NewGetQualityConversationsAuditsQueryTransactionIDInternalServerError() *GetQualityConversationsAuditsQueryTransactionIDInternalServerError {
	return &GetQualityConversationsAuditsQueryTransactionIDInternalServerError{}
}

/*GetQualityConversationsAuditsQueryTransactionIDInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetQualityConversationsAuditsQueryTransactionIDInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetQualityConversationsAuditsQueryTransactionIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/audits/query/{transactionId}][%d] getQualityConversationsAuditsQueryTransactionIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityConversationsAuditsQueryTransactionIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityConversationsAuditsQueryTransactionIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityConversationsAuditsQueryTransactionIDServiceUnavailable creates a GetQualityConversationsAuditsQueryTransactionIDServiceUnavailable with default headers values
func NewGetQualityConversationsAuditsQueryTransactionIDServiceUnavailable() *GetQualityConversationsAuditsQueryTransactionIDServiceUnavailable {
	return &GetQualityConversationsAuditsQueryTransactionIDServiceUnavailable{}
}

/*GetQualityConversationsAuditsQueryTransactionIDServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetQualityConversationsAuditsQueryTransactionIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetQualityConversationsAuditsQueryTransactionIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/audits/query/{transactionId}][%d] getQualityConversationsAuditsQueryTransactionIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityConversationsAuditsQueryTransactionIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityConversationsAuditsQueryTransactionIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityConversationsAuditsQueryTransactionIDGatewayTimeout creates a GetQualityConversationsAuditsQueryTransactionIDGatewayTimeout with default headers values
func NewGetQualityConversationsAuditsQueryTransactionIDGatewayTimeout() *GetQualityConversationsAuditsQueryTransactionIDGatewayTimeout {
	return &GetQualityConversationsAuditsQueryTransactionIDGatewayTimeout{}
}

/*GetQualityConversationsAuditsQueryTransactionIDGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetQualityConversationsAuditsQueryTransactionIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetQualityConversationsAuditsQueryTransactionIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/conversations/audits/query/{transactionId}][%d] getQualityConversationsAuditsQueryTransactionIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityConversationsAuditsQueryTransactionIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityConversationsAuditsQueryTransactionIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}