// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostIntegrationsActionsDraftsReader is a Reader for the PostIntegrationsActionsDrafts structure.
type PostIntegrationsActionsDraftsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIntegrationsActionsDraftsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostIntegrationsActionsDraftsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostIntegrationsActionsDraftsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostIntegrationsActionsDraftsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostIntegrationsActionsDraftsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostIntegrationsActionsDraftsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostIntegrationsActionsDraftsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostIntegrationsActionsDraftsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostIntegrationsActionsDraftsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostIntegrationsActionsDraftsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostIntegrationsActionsDraftsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostIntegrationsActionsDraftsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostIntegrationsActionsDraftsOK creates a PostIntegrationsActionsDraftsOK with default headers values
func NewPostIntegrationsActionsDraftsOK() *PostIntegrationsActionsDraftsOK {
	return &PostIntegrationsActionsDraftsOK{}
}

/*PostIntegrationsActionsDraftsOK handles this case with default header values.

successful operation
*/
type PostIntegrationsActionsDraftsOK struct {
	Payload *models.Action
}

func (o *PostIntegrationsActionsDraftsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsOK  %+v", 200, o.Payload)
}

func (o *PostIntegrationsActionsDraftsOK) GetPayload() *models.Action {
	return o.Payload
}

func (o *PostIntegrationsActionsDraftsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Action)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsDraftsBadRequest creates a PostIntegrationsActionsDraftsBadRequest with default headers values
func NewPostIntegrationsActionsDraftsBadRequest() *PostIntegrationsActionsDraftsBadRequest {
	return &PostIntegrationsActionsDraftsBadRequest{}
}

/*PostIntegrationsActionsDraftsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostIntegrationsActionsDraftsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionsDraftsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsBadRequest  %+v", 400, o.Payload)
}

func (o *PostIntegrationsActionsDraftsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsDraftsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsDraftsUnauthorized creates a PostIntegrationsActionsDraftsUnauthorized with default headers values
func NewPostIntegrationsActionsDraftsUnauthorized() *PostIntegrationsActionsDraftsUnauthorized {
	return &PostIntegrationsActionsDraftsUnauthorized{}
}

/*PostIntegrationsActionsDraftsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostIntegrationsActionsDraftsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionsDraftsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostIntegrationsActionsDraftsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsDraftsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsDraftsForbidden creates a PostIntegrationsActionsDraftsForbidden with default headers values
func NewPostIntegrationsActionsDraftsForbidden() *PostIntegrationsActionsDraftsForbidden {
	return &PostIntegrationsActionsDraftsForbidden{}
}

/*PostIntegrationsActionsDraftsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostIntegrationsActionsDraftsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionsDraftsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsForbidden  %+v", 403, o.Payload)
}

func (o *PostIntegrationsActionsDraftsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsDraftsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsDraftsNotFound creates a PostIntegrationsActionsDraftsNotFound with default headers values
func NewPostIntegrationsActionsDraftsNotFound() *PostIntegrationsActionsDraftsNotFound {
	return &PostIntegrationsActionsDraftsNotFound{}
}

/*PostIntegrationsActionsDraftsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostIntegrationsActionsDraftsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionsDraftsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsNotFound  %+v", 404, o.Payload)
}

func (o *PostIntegrationsActionsDraftsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsDraftsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsDraftsRequestEntityTooLarge creates a PostIntegrationsActionsDraftsRequestEntityTooLarge with default headers values
func NewPostIntegrationsActionsDraftsRequestEntityTooLarge() *PostIntegrationsActionsDraftsRequestEntityTooLarge {
	return &PostIntegrationsActionsDraftsRequestEntityTooLarge{}
}

/*PostIntegrationsActionsDraftsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostIntegrationsActionsDraftsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionsDraftsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostIntegrationsActionsDraftsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsDraftsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsDraftsUnsupportedMediaType creates a PostIntegrationsActionsDraftsUnsupportedMediaType with default headers values
func NewPostIntegrationsActionsDraftsUnsupportedMediaType() *PostIntegrationsActionsDraftsUnsupportedMediaType {
	return &PostIntegrationsActionsDraftsUnsupportedMediaType{}
}

/*PostIntegrationsActionsDraftsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostIntegrationsActionsDraftsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionsDraftsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostIntegrationsActionsDraftsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsDraftsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsDraftsTooManyRequests creates a PostIntegrationsActionsDraftsTooManyRequests with default headers values
func NewPostIntegrationsActionsDraftsTooManyRequests() *PostIntegrationsActionsDraftsTooManyRequests {
	return &PostIntegrationsActionsDraftsTooManyRequests{}
}

/*PostIntegrationsActionsDraftsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostIntegrationsActionsDraftsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionsDraftsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostIntegrationsActionsDraftsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsDraftsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsDraftsInternalServerError creates a PostIntegrationsActionsDraftsInternalServerError with default headers values
func NewPostIntegrationsActionsDraftsInternalServerError() *PostIntegrationsActionsDraftsInternalServerError {
	return &PostIntegrationsActionsDraftsInternalServerError{}
}

/*PostIntegrationsActionsDraftsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostIntegrationsActionsDraftsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionsDraftsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostIntegrationsActionsDraftsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsDraftsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsDraftsServiceUnavailable creates a PostIntegrationsActionsDraftsServiceUnavailable with default headers values
func NewPostIntegrationsActionsDraftsServiceUnavailable() *PostIntegrationsActionsDraftsServiceUnavailable {
	return &PostIntegrationsActionsDraftsServiceUnavailable{}
}

/*PostIntegrationsActionsDraftsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostIntegrationsActionsDraftsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionsDraftsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostIntegrationsActionsDraftsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsDraftsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsDraftsGatewayTimeout creates a PostIntegrationsActionsDraftsGatewayTimeout with default headers values
func NewPostIntegrationsActionsDraftsGatewayTimeout() *PostIntegrationsActionsDraftsGatewayTimeout {
	return &PostIntegrationsActionsDraftsGatewayTimeout{}
}

/*PostIntegrationsActionsDraftsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostIntegrationsActionsDraftsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionsDraftsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/drafts][%d] postIntegrationsActionsDraftsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostIntegrationsActionsDraftsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsDraftsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}