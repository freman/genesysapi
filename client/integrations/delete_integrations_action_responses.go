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

// DeleteIntegrationsActionReader is a Reader for the DeleteIntegrationsAction structure.
type DeleteIntegrationsActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIntegrationsActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteIntegrationsActionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteIntegrationsActionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteIntegrationsActionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteIntegrationsActionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteIntegrationsActionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteIntegrationsActionRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteIntegrationsActionUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteIntegrationsActionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteIntegrationsActionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteIntegrationsActionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteIntegrationsActionGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteIntegrationsActionNoContent creates a DeleteIntegrationsActionNoContent with default headers values
func NewDeleteIntegrationsActionNoContent() *DeleteIntegrationsActionNoContent {
	return &DeleteIntegrationsActionNoContent{}
}

/*DeleteIntegrationsActionNoContent handles this case with default header values.

Delete was successful
*/
type DeleteIntegrationsActionNoContent struct {
}

func (o *DeleteIntegrationsActionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}][%d] deleteIntegrationsActionNoContent ", 204)
}

func (o *DeleteIntegrationsActionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIntegrationsActionBadRequest creates a DeleteIntegrationsActionBadRequest with default headers values
func NewDeleteIntegrationsActionBadRequest() *DeleteIntegrationsActionBadRequest {
	return &DeleteIntegrationsActionBadRequest{}
}

/*DeleteIntegrationsActionBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteIntegrationsActionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteIntegrationsActionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}][%d] deleteIntegrationsActionBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteIntegrationsActionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationsActionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationsActionUnauthorized creates a DeleteIntegrationsActionUnauthorized with default headers values
func NewDeleteIntegrationsActionUnauthorized() *DeleteIntegrationsActionUnauthorized {
	return &DeleteIntegrationsActionUnauthorized{}
}

/*DeleteIntegrationsActionUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteIntegrationsActionUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteIntegrationsActionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}][%d] deleteIntegrationsActionUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteIntegrationsActionUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationsActionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationsActionForbidden creates a DeleteIntegrationsActionForbidden with default headers values
func NewDeleteIntegrationsActionForbidden() *DeleteIntegrationsActionForbidden {
	return &DeleteIntegrationsActionForbidden{}
}

/*DeleteIntegrationsActionForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteIntegrationsActionForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteIntegrationsActionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}][%d] deleteIntegrationsActionForbidden  %+v", 403, o.Payload)
}

func (o *DeleteIntegrationsActionForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationsActionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationsActionNotFound creates a DeleteIntegrationsActionNotFound with default headers values
func NewDeleteIntegrationsActionNotFound() *DeleteIntegrationsActionNotFound {
	return &DeleteIntegrationsActionNotFound{}
}

/*DeleteIntegrationsActionNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteIntegrationsActionNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteIntegrationsActionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}][%d] deleteIntegrationsActionNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIntegrationsActionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationsActionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationsActionRequestEntityTooLarge creates a DeleteIntegrationsActionRequestEntityTooLarge with default headers values
func NewDeleteIntegrationsActionRequestEntityTooLarge() *DeleteIntegrationsActionRequestEntityTooLarge {
	return &DeleteIntegrationsActionRequestEntityTooLarge{}
}

/*DeleteIntegrationsActionRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteIntegrationsActionRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteIntegrationsActionRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}][%d] deleteIntegrationsActionRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteIntegrationsActionRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationsActionRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationsActionUnsupportedMediaType creates a DeleteIntegrationsActionUnsupportedMediaType with default headers values
func NewDeleteIntegrationsActionUnsupportedMediaType() *DeleteIntegrationsActionUnsupportedMediaType {
	return &DeleteIntegrationsActionUnsupportedMediaType{}
}

/*DeleteIntegrationsActionUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteIntegrationsActionUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteIntegrationsActionUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}][%d] deleteIntegrationsActionUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteIntegrationsActionUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationsActionUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationsActionTooManyRequests creates a DeleteIntegrationsActionTooManyRequests with default headers values
func NewDeleteIntegrationsActionTooManyRequests() *DeleteIntegrationsActionTooManyRequests {
	return &DeleteIntegrationsActionTooManyRequests{}
}

/*DeleteIntegrationsActionTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteIntegrationsActionTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteIntegrationsActionTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}][%d] deleteIntegrationsActionTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteIntegrationsActionTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationsActionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationsActionInternalServerError creates a DeleteIntegrationsActionInternalServerError with default headers values
func NewDeleteIntegrationsActionInternalServerError() *DeleteIntegrationsActionInternalServerError {
	return &DeleteIntegrationsActionInternalServerError{}
}

/*DeleteIntegrationsActionInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteIntegrationsActionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteIntegrationsActionInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}][%d] deleteIntegrationsActionInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteIntegrationsActionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationsActionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationsActionServiceUnavailable creates a DeleteIntegrationsActionServiceUnavailable with default headers values
func NewDeleteIntegrationsActionServiceUnavailable() *DeleteIntegrationsActionServiceUnavailable {
	return &DeleteIntegrationsActionServiceUnavailable{}
}

/*DeleteIntegrationsActionServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteIntegrationsActionServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteIntegrationsActionServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}][%d] deleteIntegrationsActionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteIntegrationsActionServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationsActionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationsActionGatewayTimeout creates a DeleteIntegrationsActionGatewayTimeout with default headers values
func NewDeleteIntegrationsActionGatewayTimeout() *DeleteIntegrationsActionGatewayTimeout {
	return &DeleteIntegrationsActionGatewayTimeout{}
}

/*DeleteIntegrationsActionGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteIntegrationsActionGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteIntegrationsActionGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}][%d] deleteIntegrationsActionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteIntegrationsActionGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationsActionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
