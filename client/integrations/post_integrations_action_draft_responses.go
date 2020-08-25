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

// PostIntegrationsActionDraftReader is a Reader for the PostIntegrationsActionDraft structure.
type PostIntegrationsActionDraftReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIntegrationsActionDraftReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostIntegrationsActionDraftOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostIntegrationsActionDraftBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostIntegrationsActionDraftUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostIntegrationsActionDraftForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostIntegrationsActionDraftNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostIntegrationsActionDraftRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostIntegrationsActionDraftUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostIntegrationsActionDraftTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostIntegrationsActionDraftInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostIntegrationsActionDraftServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostIntegrationsActionDraftGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostIntegrationsActionDraftOK creates a PostIntegrationsActionDraftOK with default headers values
func NewPostIntegrationsActionDraftOK() *PostIntegrationsActionDraftOK {
	return &PostIntegrationsActionDraftOK{}
}

/*PostIntegrationsActionDraftOK handles this case with default header values.

successful operation
*/
type PostIntegrationsActionDraftOK struct {
	Payload *models.Action
}

func (o *PostIntegrationsActionDraftOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft][%d] postIntegrationsActionDraftOK  %+v", 200, o.Payload)
}

func (o *PostIntegrationsActionDraftOK) GetPayload() *models.Action {
	return o.Payload
}

func (o *PostIntegrationsActionDraftOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Action)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionDraftBadRequest creates a PostIntegrationsActionDraftBadRequest with default headers values
func NewPostIntegrationsActionDraftBadRequest() *PostIntegrationsActionDraftBadRequest {
	return &PostIntegrationsActionDraftBadRequest{}
}

/*PostIntegrationsActionDraftBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostIntegrationsActionDraftBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionDraftBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft][%d] postIntegrationsActionDraftBadRequest  %+v", 400, o.Payload)
}

func (o *PostIntegrationsActionDraftBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionDraftBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionDraftUnauthorized creates a PostIntegrationsActionDraftUnauthorized with default headers values
func NewPostIntegrationsActionDraftUnauthorized() *PostIntegrationsActionDraftUnauthorized {
	return &PostIntegrationsActionDraftUnauthorized{}
}

/*PostIntegrationsActionDraftUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostIntegrationsActionDraftUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionDraftUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft][%d] postIntegrationsActionDraftUnauthorized  %+v", 401, o.Payload)
}

func (o *PostIntegrationsActionDraftUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionDraftUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionDraftForbidden creates a PostIntegrationsActionDraftForbidden with default headers values
func NewPostIntegrationsActionDraftForbidden() *PostIntegrationsActionDraftForbidden {
	return &PostIntegrationsActionDraftForbidden{}
}

/*PostIntegrationsActionDraftForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostIntegrationsActionDraftForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionDraftForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft][%d] postIntegrationsActionDraftForbidden  %+v", 403, o.Payload)
}

func (o *PostIntegrationsActionDraftForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionDraftForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionDraftNotFound creates a PostIntegrationsActionDraftNotFound with default headers values
func NewPostIntegrationsActionDraftNotFound() *PostIntegrationsActionDraftNotFound {
	return &PostIntegrationsActionDraftNotFound{}
}

/*PostIntegrationsActionDraftNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostIntegrationsActionDraftNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionDraftNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft][%d] postIntegrationsActionDraftNotFound  %+v", 404, o.Payload)
}

func (o *PostIntegrationsActionDraftNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionDraftNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionDraftRequestEntityTooLarge creates a PostIntegrationsActionDraftRequestEntityTooLarge with default headers values
func NewPostIntegrationsActionDraftRequestEntityTooLarge() *PostIntegrationsActionDraftRequestEntityTooLarge {
	return &PostIntegrationsActionDraftRequestEntityTooLarge{}
}

/*PostIntegrationsActionDraftRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostIntegrationsActionDraftRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionDraftRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft][%d] postIntegrationsActionDraftRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostIntegrationsActionDraftRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionDraftRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionDraftUnsupportedMediaType creates a PostIntegrationsActionDraftUnsupportedMediaType with default headers values
func NewPostIntegrationsActionDraftUnsupportedMediaType() *PostIntegrationsActionDraftUnsupportedMediaType {
	return &PostIntegrationsActionDraftUnsupportedMediaType{}
}

/*PostIntegrationsActionDraftUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostIntegrationsActionDraftUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionDraftUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft][%d] postIntegrationsActionDraftUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostIntegrationsActionDraftUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionDraftUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionDraftTooManyRequests creates a PostIntegrationsActionDraftTooManyRequests with default headers values
func NewPostIntegrationsActionDraftTooManyRequests() *PostIntegrationsActionDraftTooManyRequests {
	return &PostIntegrationsActionDraftTooManyRequests{}
}

/*PostIntegrationsActionDraftTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostIntegrationsActionDraftTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionDraftTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft][%d] postIntegrationsActionDraftTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostIntegrationsActionDraftTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionDraftTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionDraftInternalServerError creates a PostIntegrationsActionDraftInternalServerError with default headers values
func NewPostIntegrationsActionDraftInternalServerError() *PostIntegrationsActionDraftInternalServerError {
	return &PostIntegrationsActionDraftInternalServerError{}
}

/*PostIntegrationsActionDraftInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostIntegrationsActionDraftInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionDraftInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft][%d] postIntegrationsActionDraftInternalServerError  %+v", 500, o.Payload)
}

func (o *PostIntegrationsActionDraftInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionDraftInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionDraftServiceUnavailable creates a PostIntegrationsActionDraftServiceUnavailable with default headers values
func NewPostIntegrationsActionDraftServiceUnavailable() *PostIntegrationsActionDraftServiceUnavailable {
	return &PostIntegrationsActionDraftServiceUnavailable{}
}

/*PostIntegrationsActionDraftServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostIntegrationsActionDraftServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionDraftServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft][%d] postIntegrationsActionDraftServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostIntegrationsActionDraftServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionDraftServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionDraftGatewayTimeout creates a PostIntegrationsActionDraftGatewayTimeout with default headers values
func NewPostIntegrationsActionDraftGatewayTimeout() *PostIntegrationsActionDraftGatewayTimeout {
	return &PostIntegrationsActionDraftGatewayTimeout{}
}

/*PostIntegrationsActionDraftGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostIntegrationsActionDraftGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionDraftGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/draft][%d] postIntegrationsActionDraftGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostIntegrationsActionDraftGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionDraftGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}