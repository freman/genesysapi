// Code generated by go-swagger; DO NOT EDIT.

package journey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostJourneyActiontemplatesReader is a Reader for the PostJourneyActiontemplates structure.
type PostJourneyActiontemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostJourneyActiontemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostJourneyActiontemplatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostJourneyActiontemplatesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostJourneyActiontemplatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostJourneyActiontemplatesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostJourneyActiontemplatesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostJourneyActiontemplatesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostJourneyActiontemplatesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostJourneyActiontemplatesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostJourneyActiontemplatesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostJourneyActiontemplatesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostJourneyActiontemplatesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostJourneyActiontemplatesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostJourneyActiontemplatesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostJourneyActiontemplatesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostJourneyActiontemplatesOK creates a PostJourneyActiontemplatesOK with default headers values
func NewPostJourneyActiontemplatesOK() *PostJourneyActiontemplatesOK {
	return &PostJourneyActiontemplatesOK{}
}

/*PostJourneyActiontemplatesOK handles this case with default header values.

successful operation
*/
type PostJourneyActiontemplatesOK struct {
	Payload *models.ActionTemplate
}

func (o *PostJourneyActiontemplatesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actiontemplates][%d] postJourneyActiontemplatesOK  %+v", 200, o.Payload)
}

func (o *PostJourneyActiontemplatesOK) GetPayload() *models.ActionTemplate {
	return o.Payload
}

func (o *PostJourneyActiontemplatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActiontemplatesCreated creates a PostJourneyActiontemplatesCreated with default headers values
func NewPostJourneyActiontemplatesCreated() *PostJourneyActiontemplatesCreated {
	return &PostJourneyActiontemplatesCreated{}
}

/*PostJourneyActiontemplatesCreated handles this case with default header values.

Action template created.
*/
type PostJourneyActiontemplatesCreated struct {
	Payload *models.ActionTemplate
}

func (o *PostJourneyActiontemplatesCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actiontemplates][%d] postJourneyActiontemplatesCreated  %+v", 201, o.Payload)
}

func (o *PostJourneyActiontemplatesCreated) GetPayload() *models.ActionTemplate {
	return o.Payload
}

func (o *PostJourneyActiontemplatesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActiontemplatesBadRequest creates a PostJourneyActiontemplatesBadRequest with default headers values
func NewPostJourneyActiontemplatesBadRequest() *PostJourneyActiontemplatesBadRequest {
	return &PostJourneyActiontemplatesBadRequest{}
}

/*PostJourneyActiontemplatesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostJourneyActiontemplatesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyActiontemplatesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actiontemplates][%d] postJourneyActiontemplatesBadRequest  %+v", 400, o.Payload)
}

func (o *PostJourneyActiontemplatesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyActiontemplatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActiontemplatesUnauthorized creates a PostJourneyActiontemplatesUnauthorized with default headers values
func NewPostJourneyActiontemplatesUnauthorized() *PostJourneyActiontemplatesUnauthorized {
	return &PostJourneyActiontemplatesUnauthorized{}
}

/*PostJourneyActiontemplatesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostJourneyActiontemplatesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyActiontemplatesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actiontemplates][%d] postJourneyActiontemplatesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostJourneyActiontemplatesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyActiontemplatesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActiontemplatesForbidden creates a PostJourneyActiontemplatesForbidden with default headers values
func NewPostJourneyActiontemplatesForbidden() *PostJourneyActiontemplatesForbidden {
	return &PostJourneyActiontemplatesForbidden{}
}

/*PostJourneyActiontemplatesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostJourneyActiontemplatesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyActiontemplatesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actiontemplates][%d] postJourneyActiontemplatesForbidden  %+v", 403, o.Payload)
}

func (o *PostJourneyActiontemplatesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyActiontemplatesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActiontemplatesNotFound creates a PostJourneyActiontemplatesNotFound with default headers values
func NewPostJourneyActiontemplatesNotFound() *PostJourneyActiontemplatesNotFound {
	return &PostJourneyActiontemplatesNotFound{}
}

/*PostJourneyActiontemplatesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostJourneyActiontemplatesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyActiontemplatesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actiontemplates][%d] postJourneyActiontemplatesNotFound  %+v", 404, o.Payload)
}

func (o *PostJourneyActiontemplatesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyActiontemplatesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActiontemplatesRequestTimeout creates a PostJourneyActiontemplatesRequestTimeout with default headers values
func NewPostJourneyActiontemplatesRequestTimeout() *PostJourneyActiontemplatesRequestTimeout {
	return &PostJourneyActiontemplatesRequestTimeout{}
}

/*PostJourneyActiontemplatesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostJourneyActiontemplatesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyActiontemplatesRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actiontemplates][%d] postJourneyActiontemplatesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostJourneyActiontemplatesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyActiontemplatesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActiontemplatesConflict creates a PostJourneyActiontemplatesConflict with default headers values
func NewPostJourneyActiontemplatesConflict() *PostJourneyActiontemplatesConflict {
	return &PostJourneyActiontemplatesConflict{}
}

/*PostJourneyActiontemplatesConflict handles this case with default header values.

Conflict
*/
type PostJourneyActiontemplatesConflict struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyActiontemplatesConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actiontemplates][%d] postJourneyActiontemplatesConflict  %+v", 409, o.Payload)
}

func (o *PostJourneyActiontemplatesConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyActiontemplatesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActiontemplatesRequestEntityTooLarge creates a PostJourneyActiontemplatesRequestEntityTooLarge with default headers values
func NewPostJourneyActiontemplatesRequestEntityTooLarge() *PostJourneyActiontemplatesRequestEntityTooLarge {
	return &PostJourneyActiontemplatesRequestEntityTooLarge{}
}

/*PostJourneyActiontemplatesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostJourneyActiontemplatesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyActiontemplatesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actiontemplates][%d] postJourneyActiontemplatesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostJourneyActiontemplatesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyActiontemplatesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActiontemplatesUnsupportedMediaType creates a PostJourneyActiontemplatesUnsupportedMediaType with default headers values
func NewPostJourneyActiontemplatesUnsupportedMediaType() *PostJourneyActiontemplatesUnsupportedMediaType {
	return &PostJourneyActiontemplatesUnsupportedMediaType{}
}

/*PostJourneyActiontemplatesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostJourneyActiontemplatesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyActiontemplatesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actiontemplates][%d] postJourneyActiontemplatesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostJourneyActiontemplatesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyActiontemplatesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActiontemplatesTooManyRequests creates a PostJourneyActiontemplatesTooManyRequests with default headers values
func NewPostJourneyActiontemplatesTooManyRequests() *PostJourneyActiontemplatesTooManyRequests {
	return &PostJourneyActiontemplatesTooManyRequests{}
}

/*PostJourneyActiontemplatesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostJourneyActiontemplatesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyActiontemplatesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actiontemplates][%d] postJourneyActiontemplatesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostJourneyActiontemplatesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyActiontemplatesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActiontemplatesInternalServerError creates a PostJourneyActiontemplatesInternalServerError with default headers values
func NewPostJourneyActiontemplatesInternalServerError() *PostJourneyActiontemplatesInternalServerError {
	return &PostJourneyActiontemplatesInternalServerError{}
}

/*PostJourneyActiontemplatesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostJourneyActiontemplatesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyActiontemplatesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actiontemplates][%d] postJourneyActiontemplatesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostJourneyActiontemplatesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyActiontemplatesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActiontemplatesServiceUnavailable creates a PostJourneyActiontemplatesServiceUnavailable with default headers values
func NewPostJourneyActiontemplatesServiceUnavailable() *PostJourneyActiontemplatesServiceUnavailable {
	return &PostJourneyActiontemplatesServiceUnavailable{}
}

/*PostJourneyActiontemplatesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostJourneyActiontemplatesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyActiontemplatesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actiontemplates][%d] postJourneyActiontemplatesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostJourneyActiontemplatesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyActiontemplatesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActiontemplatesGatewayTimeout creates a PostJourneyActiontemplatesGatewayTimeout with default headers values
func NewPostJourneyActiontemplatesGatewayTimeout() *PostJourneyActiontemplatesGatewayTimeout {
	return &PostJourneyActiontemplatesGatewayTimeout{}
}

/*PostJourneyActiontemplatesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostJourneyActiontemplatesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyActiontemplatesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actiontemplates][%d] postJourneyActiontemplatesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostJourneyActiontemplatesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyActiontemplatesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
