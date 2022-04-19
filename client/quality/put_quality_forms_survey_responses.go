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

// PutQualityFormsSurveyReader is a Reader for the PutQualityFormsSurvey structure.
type PutQualityFormsSurveyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutQualityFormsSurveyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutQualityFormsSurveyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutQualityFormsSurveyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutQualityFormsSurveyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutQualityFormsSurveyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutQualityFormsSurveyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutQualityFormsSurveyRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutQualityFormsSurveyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutQualityFormsSurveyRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutQualityFormsSurveyUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutQualityFormsSurveyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutQualityFormsSurveyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutQualityFormsSurveyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutQualityFormsSurveyGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutQualityFormsSurveyOK creates a PutQualityFormsSurveyOK with default headers values
func NewPutQualityFormsSurveyOK() *PutQualityFormsSurveyOK {
	return &PutQualityFormsSurveyOK{}
}

/*PutQualityFormsSurveyOK handles this case with default header values.

successful operation
*/
type PutQualityFormsSurveyOK struct {
	Payload *models.SurveyForm
}

func (o *PutQualityFormsSurveyOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/surveys/{formId}][%d] putQualityFormsSurveyOK  %+v", 200, o.Payload)
}

func (o *PutQualityFormsSurveyOK) GetPayload() *models.SurveyForm {
	return o.Payload
}

func (o *PutQualityFormsSurveyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SurveyForm)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityFormsSurveyBadRequest creates a PutQualityFormsSurveyBadRequest with default headers values
func NewPutQualityFormsSurveyBadRequest() *PutQualityFormsSurveyBadRequest {
	return &PutQualityFormsSurveyBadRequest{}
}

/*PutQualityFormsSurveyBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutQualityFormsSurveyBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutQualityFormsSurveyBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/surveys/{formId}][%d] putQualityFormsSurveyBadRequest  %+v", 400, o.Payload)
}

func (o *PutQualityFormsSurveyBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityFormsSurveyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityFormsSurveyUnauthorized creates a PutQualityFormsSurveyUnauthorized with default headers values
func NewPutQualityFormsSurveyUnauthorized() *PutQualityFormsSurveyUnauthorized {
	return &PutQualityFormsSurveyUnauthorized{}
}

/*PutQualityFormsSurveyUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutQualityFormsSurveyUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutQualityFormsSurveyUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/surveys/{formId}][%d] putQualityFormsSurveyUnauthorized  %+v", 401, o.Payload)
}

func (o *PutQualityFormsSurveyUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityFormsSurveyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityFormsSurveyForbidden creates a PutQualityFormsSurveyForbidden with default headers values
func NewPutQualityFormsSurveyForbidden() *PutQualityFormsSurveyForbidden {
	return &PutQualityFormsSurveyForbidden{}
}

/*PutQualityFormsSurveyForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutQualityFormsSurveyForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutQualityFormsSurveyForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/surveys/{formId}][%d] putQualityFormsSurveyForbidden  %+v", 403, o.Payload)
}

func (o *PutQualityFormsSurveyForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityFormsSurveyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityFormsSurveyNotFound creates a PutQualityFormsSurveyNotFound with default headers values
func NewPutQualityFormsSurveyNotFound() *PutQualityFormsSurveyNotFound {
	return &PutQualityFormsSurveyNotFound{}
}

/*PutQualityFormsSurveyNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutQualityFormsSurveyNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutQualityFormsSurveyNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/surveys/{formId}][%d] putQualityFormsSurveyNotFound  %+v", 404, o.Payload)
}

func (o *PutQualityFormsSurveyNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityFormsSurveyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityFormsSurveyRequestTimeout creates a PutQualityFormsSurveyRequestTimeout with default headers values
func NewPutQualityFormsSurveyRequestTimeout() *PutQualityFormsSurveyRequestTimeout {
	return &PutQualityFormsSurveyRequestTimeout{}
}

/*PutQualityFormsSurveyRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutQualityFormsSurveyRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutQualityFormsSurveyRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/surveys/{formId}][%d] putQualityFormsSurveyRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutQualityFormsSurveyRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityFormsSurveyRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityFormsSurveyConflict creates a PutQualityFormsSurveyConflict with default headers values
func NewPutQualityFormsSurveyConflict() *PutQualityFormsSurveyConflict {
	return &PutQualityFormsSurveyConflict{}
}

/*PutQualityFormsSurveyConflict handles this case with default header values.

Conflict
*/
type PutQualityFormsSurveyConflict struct {
	Payload *models.ErrorBody
}

func (o *PutQualityFormsSurveyConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/surveys/{formId}][%d] putQualityFormsSurveyConflict  %+v", 409, o.Payload)
}

func (o *PutQualityFormsSurveyConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityFormsSurveyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityFormsSurveyRequestEntityTooLarge creates a PutQualityFormsSurveyRequestEntityTooLarge with default headers values
func NewPutQualityFormsSurveyRequestEntityTooLarge() *PutQualityFormsSurveyRequestEntityTooLarge {
	return &PutQualityFormsSurveyRequestEntityTooLarge{}
}

/*PutQualityFormsSurveyRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutQualityFormsSurveyRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutQualityFormsSurveyRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/surveys/{formId}][%d] putQualityFormsSurveyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutQualityFormsSurveyRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityFormsSurveyRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityFormsSurveyUnsupportedMediaType creates a PutQualityFormsSurveyUnsupportedMediaType with default headers values
func NewPutQualityFormsSurveyUnsupportedMediaType() *PutQualityFormsSurveyUnsupportedMediaType {
	return &PutQualityFormsSurveyUnsupportedMediaType{}
}

/*PutQualityFormsSurveyUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutQualityFormsSurveyUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutQualityFormsSurveyUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/surveys/{formId}][%d] putQualityFormsSurveyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutQualityFormsSurveyUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityFormsSurveyUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityFormsSurveyTooManyRequests creates a PutQualityFormsSurveyTooManyRequests with default headers values
func NewPutQualityFormsSurveyTooManyRequests() *PutQualityFormsSurveyTooManyRequests {
	return &PutQualityFormsSurveyTooManyRequests{}
}

/*PutQualityFormsSurveyTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutQualityFormsSurveyTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutQualityFormsSurveyTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/surveys/{formId}][%d] putQualityFormsSurveyTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutQualityFormsSurveyTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityFormsSurveyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityFormsSurveyInternalServerError creates a PutQualityFormsSurveyInternalServerError with default headers values
func NewPutQualityFormsSurveyInternalServerError() *PutQualityFormsSurveyInternalServerError {
	return &PutQualityFormsSurveyInternalServerError{}
}

/*PutQualityFormsSurveyInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutQualityFormsSurveyInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutQualityFormsSurveyInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/surveys/{formId}][%d] putQualityFormsSurveyInternalServerError  %+v", 500, o.Payload)
}

func (o *PutQualityFormsSurveyInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityFormsSurveyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityFormsSurveyServiceUnavailable creates a PutQualityFormsSurveyServiceUnavailable with default headers values
func NewPutQualityFormsSurveyServiceUnavailable() *PutQualityFormsSurveyServiceUnavailable {
	return &PutQualityFormsSurveyServiceUnavailable{}
}

/*PutQualityFormsSurveyServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutQualityFormsSurveyServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutQualityFormsSurveyServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/surveys/{formId}][%d] putQualityFormsSurveyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutQualityFormsSurveyServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityFormsSurveyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutQualityFormsSurveyGatewayTimeout creates a PutQualityFormsSurveyGatewayTimeout with default headers values
func NewPutQualityFormsSurveyGatewayTimeout() *PutQualityFormsSurveyGatewayTimeout {
	return &PutQualityFormsSurveyGatewayTimeout{}
}

/*PutQualityFormsSurveyGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutQualityFormsSurveyGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutQualityFormsSurveyGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/quality/forms/surveys/{formId}][%d] putQualityFormsSurveyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutQualityFormsSurveyGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutQualityFormsSurveyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
