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

// PutQualityConversationEvaluationReader is a Reader for the PutQualityConversationEvaluation structure.
type PutQualityConversationEvaluationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutQualityConversationEvaluationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutQualityConversationEvaluationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutQualityConversationEvaluationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutQualityConversationEvaluationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutQualityConversationEvaluationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutQualityConversationEvaluationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutQualityConversationEvaluationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutQualityConversationEvaluationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutQualityConversationEvaluationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutQualityConversationEvaluationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutQualityConversationEvaluationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutQualityConversationEvaluationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutQualityConversationEvaluationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutQualityConversationEvaluationOK creates a PutQualityConversationEvaluationOK with default headers values
func NewPutQualityConversationEvaluationOK() *PutQualityConversationEvaluationOK {
	return &PutQualityConversationEvaluationOK{}
}

/*PutQualityConversationEvaluationOK handles this case with default header values.

successful operation
*/
type PutQualityConversationEvaluationOK struct {
	Payload *models.Evaluation
}

func (o *PutQualityConversationEvaluationOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/conversations/{conversationId}/evaluations/{evaluationId}][%d] putQualityConversationEvaluationOK  %+v", 200, o.Payload)
}

func (o *PutQualityConversationEvaluationOK) GetPayload() *models.Evaluation {
	return o.Payload
}

func (o *PutQualityConversationEvaluationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Evaluation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityConversationEvaluationBadRequest creates a PutQualityConversationEvaluationBadRequest with default headers values
func NewPutQualityConversationEvaluationBadRequest() *PutQualityConversationEvaluationBadRequest {
	return &PutQualityConversationEvaluationBadRequest{}
}

/*PutQualityConversationEvaluationBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutQualityConversationEvaluationBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutQualityConversationEvaluationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/conversations/{conversationId}/evaluations/{evaluationId}][%d] putQualityConversationEvaluationBadRequest  %+v", 400, o.Payload)
}

func (o *PutQualityConversationEvaluationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityConversationEvaluationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityConversationEvaluationUnauthorized creates a PutQualityConversationEvaluationUnauthorized with default headers values
func NewPutQualityConversationEvaluationUnauthorized() *PutQualityConversationEvaluationUnauthorized {
	return &PutQualityConversationEvaluationUnauthorized{}
}

/*PutQualityConversationEvaluationUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutQualityConversationEvaluationUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutQualityConversationEvaluationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/conversations/{conversationId}/evaluations/{evaluationId}][%d] putQualityConversationEvaluationUnauthorized  %+v", 401, o.Payload)
}

func (o *PutQualityConversationEvaluationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityConversationEvaluationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityConversationEvaluationForbidden creates a PutQualityConversationEvaluationForbidden with default headers values
func NewPutQualityConversationEvaluationForbidden() *PutQualityConversationEvaluationForbidden {
	return &PutQualityConversationEvaluationForbidden{}
}

/*PutQualityConversationEvaluationForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutQualityConversationEvaluationForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutQualityConversationEvaluationForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/conversations/{conversationId}/evaluations/{evaluationId}][%d] putQualityConversationEvaluationForbidden  %+v", 403, o.Payload)
}

func (o *PutQualityConversationEvaluationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityConversationEvaluationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityConversationEvaluationNotFound creates a PutQualityConversationEvaluationNotFound with default headers values
func NewPutQualityConversationEvaluationNotFound() *PutQualityConversationEvaluationNotFound {
	return &PutQualityConversationEvaluationNotFound{}
}

/*PutQualityConversationEvaluationNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutQualityConversationEvaluationNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutQualityConversationEvaluationNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/conversations/{conversationId}/evaluations/{evaluationId}][%d] putQualityConversationEvaluationNotFound  %+v", 404, o.Payload)
}

func (o *PutQualityConversationEvaluationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityConversationEvaluationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityConversationEvaluationRequestTimeout creates a PutQualityConversationEvaluationRequestTimeout with default headers values
func NewPutQualityConversationEvaluationRequestTimeout() *PutQualityConversationEvaluationRequestTimeout {
	return &PutQualityConversationEvaluationRequestTimeout{}
}

/*PutQualityConversationEvaluationRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutQualityConversationEvaluationRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutQualityConversationEvaluationRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/conversations/{conversationId}/evaluations/{evaluationId}][%d] putQualityConversationEvaluationRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutQualityConversationEvaluationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityConversationEvaluationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityConversationEvaluationRequestEntityTooLarge creates a PutQualityConversationEvaluationRequestEntityTooLarge with default headers values
func NewPutQualityConversationEvaluationRequestEntityTooLarge() *PutQualityConversationEvaluationRequestEntityTooLarge {
	return &PutQualityConversationEvaluationRequestEntityTooLarge{}
}

/*PutQualityConversationEvaluationRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutQualityConversationEvaluationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutQualityConversationEvaluationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/conversations/{conversationId}/evaluations/{evaluationId}][%d] putQualityConversationEvaluationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutQualityConversationEvaluationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityConversationEvaluationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityConversationEvaluationUnsupportedMediaType creates a PutQualityConversationEvaluationUnsupportedMediaType with default headers values
func NewPutQualityConversationEvaluationUnsupportedMediaType() *PutQualityConversationEvaluationUnsupportedMediaType {
	return &PutQualityConversationEvaluationUnsupportedMediaType{}
}

/*PutQualityConversationEvaluationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutQualityConversationEvaluationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutQualityConversationEvaluationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/conversations/{conversationId}/evaluations/{evaluationId}][%d] putQualityConversationEvaluationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutQualityConversationEvaluationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityConversationEvaluationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityConversationEvaluationTooManyRequests creates a PutQualityConversationEvaluationTooManyRequests with default headers values
func NewPutQualityConversationEvaluationTooManyRequests() *PutQualityConversationEvaluationTooManyRequests {
	return &PutQualityConversationEvaluationTooManyRequests{}
}

/*PutQualityConversationEvaluationTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutQualityConversationEvaluationTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutQualityConversationEvaluationTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/conversations/{conversationId}/evaluations/{evaluationId}][%d] putQualityConversationEvaluationTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutQualityConversationEvaluationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityConversationEvaluationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityConversationEvaluationInternalServerError creates a PutQualityConversationEvaluationInternalServerError with default headers values
func NewPutQualityConversationEvaluationInternalServerError() *PutQualityConversationEvaluationInternalServerError {
	return &PutQualityConversationEvaluationInternalServerError{}
}

/*PutQualityConversationEvaluationInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutQualityConversationEvaluationInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutQualityConversationEvaluationInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/conversations/{conversationId}/evaluations/{evaluationId}][%d] putQualityConversationEvaluationInternalServerError  %+v", 500, o.Payload)
}

func (o *PutQualityConversationEvaluationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityConversationEvaluationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityConversationEvaluationServiceUnavailable creates a PutQualityConversationEvaluationServiceUnavailable with default headers values
func NewPutQualityConversationEvaluationServiceUnavailable() *PutQualityConversationEvaluationServiceUnavailable {
	return &PutQualityConversationEvaluationServiceUnavailable{}
}

/*PutQualityConversationEvaluationServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutQualityConversationEvaluationServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutQualityConversationEvaluationServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/conversations/{conversationId}/evaluations/{evaluationId}][%d] putQualityConversationEvaluationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutQualityConversationEvaluationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityConversationEvaluationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityConversationEvaluationGatewayTimeout creates a PutQualityConversationEvaluationGatewayTimeout with default headers values
func NewPutQualityConversationEvaluationGatewayTimeout() *PutQualityConversationEvaluationGatewayTimeout {
	return &PutQualityConversationEvaluationGatewayTimeout{}
}

/*PutQualityConversationEvaluationGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutQualityConversationEvaluationGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutQualityConversationEvaluationGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/conversations/{conversationId}/evaluations/{evaluationId}][%d] putQualityConversationEvaluationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutQualityConversationEvaluationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityConversationEvaluationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
