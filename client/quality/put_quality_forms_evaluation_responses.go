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

// PutQualityFormsEvaluationReader is a Reader for the PutQualityFormsEvaluation structure.
type PutQualityFormsEvaluationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutQualityFormsEvaluationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutQualityFormsEvaluationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutQualityFormsEvaluationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutQualityFormsEvaluationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutQualityFormsEvaluationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutQualityFormsEvaluationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutQualityFormsEvaluationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutQualityFormsEvaluationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutQualityFormsEvaluationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutQualityFormsEvaluationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutQualityFormsEvaluationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutQualityFormsEvaluationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutQualityFormsEvaluationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutQualityFormsEvaluationOK creates a PutQualityFormsEvaluationOK with default headers values
func NewPutQualityFormsEvaluationOK() *PutQualityFormsEvaluationOK {
	return &PutQualityFormsEvaluationOK{}
}

/*PutQualityFormsEvaluationOK handles this case with default header values.

successful operation
*/
type PutQualityFormsEvaluationOK struct {
	Payload *models.EvaluationForm
}

func (o *PutQualityFormsEvaluationOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/evaluations/{formId}][%d] putQualityFormsEvaluationOK  %+v", 200, o.Payload)
}

func (o *PutQualityFormsEvaluationOK) GetPayload() *models.EvaluationForm {
	return o.Payload
}

func (o *PutQualityFormsEvaluationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EvaluationForm)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityFormsEvaluationBadRequest creates a PutQualityFormsEvaluationBadRequest with default headers values
func NewPutQualityFormsEvaluationBadRequest() *PutQualityFormsEvaluationBadRequest {
	return &PutQualityFormsEvaluationBadRequest{}
}

/*PutQualityFormsEvaluationBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutQualityFormsEvaluationBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutQualityFormsEvaluationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/evaluations/{formId}][%d] putQualityFormsEvaluationBadRequest  %+v", 400, o.Payload)
}

func (o *PutQualityFormsEvaluationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityFormsEvaluationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityFormsEvaluationUnauthorized creates a PutQualityFormsEvaluationUnauthorized with default headers values
func NewPutQualityFormsEvaluationUnauthorized() *PutQualityFormsEvaluationUnauthorized {
	return &PutQualityFormsEvaluationUnauthorized{}
}

/*PutQualityFormsEvaluationUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutQualityFormsEvaluationUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutQualityFormsEvaluationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/evaluations/{formId}][%d] putQualityFormsEvaluationUnauthorized  %+v", 401, o.Payload)
}

func (o *PutQualityFormsEvaluationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityFormsEvaluationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityFormsEvaluationForbidden creates a PutQualityFormsEvaluationForbidden with default headers values
func NewPutQualityFormsEvaluationForbidden() *PutQualityFormsEvaluationForbidden {
	return &PutQualityFormsEvaluationForbidden{}
}

/*PutQualityFormsEvaluationForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutQualityFormsEvaluationForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutQualityFormsEvaluationForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/evaluations/{formId}][%d] putQualityFormsEvaluationForbidden  %+v", 403, o.Payload)
}

func (o *PutQualityFormsEvaluationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityFormsEvaluationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityFormsEvaluationNotFound creates a PutQualityFormsEvaluationNotFound with default headers values
func NewPutQualityFormsEvaluationNotFound() *PutQualityFormsEvaluationNotFound {
	return &PutQualityFormsEvaluationNotFound{}
}

/*PutQualityFormsEvaluationNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutQualityFormsEvaluationNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutQualityFormsEvaluationNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/evaluations/{formId}][%d] putQualityFormsEvaluationNotFound  %+v", 404, o.Payload)
}

func (o *PutQualityFormsEvaluationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityFormsEvaluationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityFormsEvaluationRequestTimeout creates a PutQualityFormsEvaluationRequestTimeout with default headers values
func NewPutQualityFormsEvaluationRequestTimeout() *PutQualityFormsEvaluationRequestTimeout {
	return &PutQualityFormsEvaluationRequestTimeout{}
}

/*PutQualityFormsEvaluationRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutQualityFormsEvaluationRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutQualityFormsEvaluationRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/evaluations/{formId}][%d] putQualityFormsEvaluationRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutQualityFormsEvaluationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityFormsEvaluationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityFormsEvaluationRequestEntityTooLarge creates a PutQualityFormsEvaluationRequestEntityTooLarge with default headers values
func NewPutQualityFormsEvaluationRequestEntityTooLarge() *PutQualityFormsEvaluationRequestEntityTooLarge {
	return &PutQualityFormsEvaluationRequestEntityTooLarge{}
}

/*PutQualityFormsEvaluationRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutQualityFormsEvaluationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutQualityFormsEvaluationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/evaluations/{formId}][%d] putQualityFormsEvaluationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutQualityFormsEvaluationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityFormsEvaluationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityFormsEvaluationUnsupportedMediaType creates a PutQualityFormsEvaluationUnsupportedMediaType with default headers values
func NewPutQualityFormsEvaluationUnsupportedMediaType() *PutQualityFormsEvaluationUnsupportedMediaType {
	return &PutQualityFormsEvaluationUnsupportedMediaType{}
}

/*PutQualityFormsEvaluationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutQualityFormsEvaluationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutQualityFormsEvaluationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/evaluations/{formId}][%d] putQualityFormsEvaluationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutQualityFormsEvaluationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityFormsEvaluationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityFormsEvaluationTooManyRequests creates a PutQualityFormsEvaluationTooManyRequests with default headers values
func NewPutQualityFormsEvaluationTooManyRequests() *PutQualityFormsEvaluationTooManyRequests {
	return &PutQualityFormsEvaluationTooManyRequests{}
}

/*PutQualityFormsEvaluationTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutQualityFormsEvaluationTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutQualityFormsEvaluationTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/evaluations/{formId}][%d] putQualityFormsEvaluationTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutQualityFormsEvaluationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityFormsEvaluationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityFormsEvaluationInternalServerError creates a PutQualityFormsEvaluationInternalServerError with default headers values
func NewPutQualityFormsEvaluationInternalServerError() *PutQualityFormsEvaluationInternalServerError {
	return &PutQualityFormsEvaluationInternalServerError{}
}

/*PutQualityFormsEvaluationInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutQualityFormsEvaluationInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutQualityFormsEvaluationInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/evaluations/{formId}][%d] putQualityFormsEvaluationInternalServerError  %+v", 500, o.Payload)
}

func (o *PutQualityFormsEvaluationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityFormsEvaluationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityFormsEvaluationServiceUnavailable creates a PutQualityFormsEvaluationServiceUnavailable with default headers values
func NewPutQualityFormsEvaluationServiceUnavailable() *PutQualityFormsEvaluationServiceUnavailable {
	return &PutQualityFormsEvaluationServiceUnavailable{}
}

/*PutQualityFormsEvaluationServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutQualityFormsEvaluationServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutQualityFormsEvaluationServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/evaluations/{formId}][%d] putQualityFormsEvaluationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutQualityFormsEvaluationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityFormsEvaluationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityFormsEvaluationGatewayTimeout creates a PutQualityFormsEvaluationGatewayTimeout with default headers values
func NewPutQualityFormsEvaluationGatewayTimeout() *PutQualityFormsEvaluationGatewayTimeout {
	return &PutQualityFormsEvaluationGatewayTimeout{}
}

/*PutQualityFormsEvaluationGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutQualityFormsEvaluationGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutQualityFormsEvaluationGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/evaluations/{formId}][%d] putQualityFormsEvaluationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutQualityFormsEvaluationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityFormsEvaluationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
