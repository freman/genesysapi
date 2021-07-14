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

// GetIntegrationsActionReader is a Reader for the GetIntegrationsAction structure.
type GetIntegrationsActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsActionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsActionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsActionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsActionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetIntegrationsActionRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsActionRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsActionUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsActionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsActionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsActionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsActionGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsActionOK creates a GetIntegrationsActionOK with default headers values
func NewGetIntegrationsActionOK() *GetIntegrationsActionOK {
	return &GetIntegrationsActionOK{}
}

/*GetIntegrationsActionOK handles this case with default header values.

successful operation
*/
type GetIntegrationsActionOK struct {
	Payload *models.Action
}

func (o *GetIntegrationsActionOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}][%d] getIntegrationsActionOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsActionOK) GetPayload() *models.Action {
	return o.Payload
}

func (o *GetIntegrationsActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Action)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionBadRequest creates a GetIntegrationsActionBadRequest with default headers values
func NewGetIntegrationsActionBadRequest() *GetIntegrationsActionBadRequest {
	return &GetIntegrationsActionBadRequest{}
}

/*GetIntegrationsActionBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsActionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}][%d] getIntegrationsActionBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsActionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionUnauthorized creates a GetIntegrationsActionUnauthorized with default headers values
func NewGetIntegrationsActionUnauthorized() *GetIntegrationsActionUnauthorized {
	return &GetIntegrationsActionUnauthorized{}
}

/*GetIntegrationsActionUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsActionUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}][%d] getIntegrationsActionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsActionUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionForbidden creates a GetIntegrationsActionForbidden with default headers values
func NewGetIntegrationsActionForbidden() *GetIntegrationsActionForbidden {
	return &GetIntegrationsActionForbidden{}
}

/*GetIntegrationsActionForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsActionForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}][%d] getIntegrationsActionForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsActionForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionNotFound creates a GetIntegrationsActionNotFound with default headers values
func NewGetIntegrationsActionNotFound() *GetIntegrationsActionNotFound {
	return &GetIntegrationsActionNotFound{}
}

/*GetIntegrationsActionNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetIntegrationsActionNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}][%d] getIntegrationsActionNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsActionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionRequestTimeout creates a GetIntegrationsActionRequestTimeout with default headers values
func NewGetIntegrationsActionRequestTimeout() *GetIntegrationsActionRequestTimeout {
	return &GetIntegrationsActionRequestTimeout{}
}

/*GetIntegrationsActionRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetIntegrationsActionRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}][%d] getIntegrationsActionRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsActionRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionRequestEntityTooLarge creates a GetIntegrationsActionRequestEntityTooLarge with default headers values
func NewGetIntegrationsActionRequestEntityTooLarge() *GetIntegrationsActionRequestEntityTooLarge {
	return &GetIntegrationsActionRequestEntityTooLarge{}
}

/*GetIntegrationsActionRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetIntegrationsActionRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}][%d] getIntegrationsActionRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsActionRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionUnsupportedMediaType creates a GetIntegrationsActionUnsupportedMediaType with default headers values
func NewGetIntegrationsActionUnsupportedMediaType() *GetIntegrationsActionUnsupportedMediaType {
	return &GetIntegrationsActionUnsupportedMediaType{}
}

/*GetIntegrationsActionUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsActionUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}][%d] getIntegrationsActionUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsActionUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionTooManyRequests creates a GetIntegrationsActionTooManyRequests with default headers values
func NewGetIntegrationsActionTooManyRequests() *GetIntegrationsActionTooManyRequests {
	return &GetIntegrationsActionTooManyRequests{}
}

/*GetIntegrationsActionTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetIntegrationsActionTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}][%d] getIntegrationsActionTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsActionTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionInternalServerError creates a GetIntegrationsActionInternalServerError with default headers values
func NewGetIntegrationsActionInternalServerError() *GetIntegrationsActionInternalServerError {
	return &GetIntegrationsActionInternalServerError{}
}

/*GetIntegrationsActionInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsActionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}][%d] getIntegrationsActionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsActionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionServiceUnavailable creates a GetIntegrationsActionServiceUnavailable with default headers values
func NewGetIntegrationsActionServiceUnavailable() *GetIntegrationsActionServiceUnavailable {
	return &GetIntegrationsActionServiceUnavailable{}
}

/*GetIntegrationsActionServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsActionServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}][%d] getIntegrationsActionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsActionServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionGatewayTimeout creates a GetIntegrationsActionGatewayTimeout with default headers values
func NewGetIntegrationsActionGatewayTimeout() *GetIntegrationsActionGatewayTimeout {
	return &GetIntegrationsActionGatewayTimeout{}
}

/*GetIntegrationsActionGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetIntegrationsActionGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}][%d] getIntegrationsActionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsActionGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
