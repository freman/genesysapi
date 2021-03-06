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

// PostQualityFormsReader is a Reader for the PostQualityForms structure.
type PostQualityFormsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostQualityFormsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostQualityFormsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostQualityFormsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostQualityFormsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostQualityFormsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostQualityFormsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostQualityFormsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostQualityFormsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostQualityFormsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostQualityFormsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostQualityFormsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostQualityFormsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostQualityFormsOK creates a PostQualityFormsOK with default headers values
func NewPostQualityFormsOK() *PostQualityFormsOK {
	return &PostQualityFormsOK{}
}

/*PostQualityFormsOK handles this case with default header values.

successful operation
*/
type PostQualityFormsOK struct {
	Payload *models.EvaluationForm
}

func (o *PostQualityFormsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms][%d] postQualityFormsOK  %+v", 200, o.Payload)
}

func (o *PostQualityFormsOK) GetPayload() *models.EvaluationForm {
	return o.Payload
}

func (o *PostQualityFormsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EvaluationForm)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsBadRequest creates a PostQualityFormsBadRequest with default headers values
func NewPostQualityFormsBadRequest() *PostQualityFormsBadRequest {
	return &PostQualityFormsBadRequest{}
}

/*PostQualityFormsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostQualityFormsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms][%d] postQualityFormsBadRequest  %+v", 400, o.Payload)
}

func (o *PostQualityFormsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsUnauthorized creates a PostQualityFormsUnauthorized with default headers values
func NewPostQualityFormsUnauthorized() *PostQualityFormsUnauthorized {
	return &PostQualityFormsUnauthorized{}
}

/*PostQualityFormsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostQualityFormsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms][%d] postQualityFormsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostQualityFormsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsForbidden creates a PostQualityFormsForbidden with default headers values
func NewPostQualityFormsForbidden() *PostQualityFormsForbidden {
	return &PostQualityFormsForbidden{}
}

/*PostQualityFormsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostQualityFormsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms][%d] postQualityFormsForbidden  %+v", 403, o.Payload)
}

func (o *PostQualityFormsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsNotFound creates a PostQualityFormsNotFound with default headers values
func NewPostQualityFormsNotFound() *PostQualityFormsNotFound {
	return &PostQualityFormsNotFound{}
}

/*PostQualityFormsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostQualityFormsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms][%d] postQualityFormsNotFound  %+v", 404, o.Payload)
}

func (o *PostQualityFormsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsRequestEntityTooLarge creates a PostQualityFormsRequestEntityTooLarge with default headers values
func NewPostQualityFormsRequestEntityTooLarge() *PostQualityFormsRequestEntityTooLarge {
	return &PostQualityFormsRequestEntityTooLarge{}
}

/*PostQualityFormsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostQualityFormsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms][%d] postQualityFormsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostQualityFormsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsUnsupportedMediaType creates a PostQualityFormsUnsupportedMediaType with default headers values
func NewPostQualityFormsUnsupportedMediaType() *PostQualityFormsUnsupportedMediaType {
	return &PostQualityFormsUnsupportedMediaType{}
}

/*PostQualityFormsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostQualityFormsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms][%d] postQualityFormsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostQualityFormsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsTooManyRequests creates a PostQualityFormsTooManyRequests with default headers values
func NewPostQualityFormsTooManyRequests() *PostQualityFormsTooManyRequests {
	return &PostQualityFormsTooManyRequests{}
}

/*PostQualityFormsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostQualityFormsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms][%d] postQualityFormsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostQualityFormsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsInternalServerError creates a PostQualityFormsInternalServerError with default headers values
func NewPostQualityFormsInternalServerError() *PostQualityFormsInternalServerError {
	return &PostQualityFormsInternalServerError{}
}

/*PostQualityFormsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostQualityFormsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms][%d] postQualityFormsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostQualityFormsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsServiceUnavailable creates a PostQualityFormsServiceUnavailable with default headers values
func NewPostQualityFormsServiceUnavailable() *PostQualityFormsServiceUnavailable {
	return &PostQualityFormsServiceUnavailable{}
}

/*PostQualityFormsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostQualityFormsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms][%d] postQualityFormsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostQualityFormsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsGatewayTimeout creates a PostQualityFormsGatewayTimeout with default headers values
func NewPostQualityFormsGatewayTimeout() *PostQualityFormsGatewayTimeout {
	return &PostQualityFormsGatewayTimeout{}
}

/*PostQualityFormsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostQualityFormsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms][%d] postQualityFormsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostQualityFormsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
