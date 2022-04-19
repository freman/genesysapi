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

// GetIntegrationsActionDraftValidationReader is a Reader for the GetIntegrationsActionDraftValidation structure.
type GetIntegrationsActionDraftValidationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsActionDraftValidationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsActionDraftValidationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsActionDraftValidationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsActionDraftValidationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsActionDraftValidationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsActionDraftValidationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetIntegrationsActionDraftValidationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsActionDraftValidationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsActionDraftValidationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsActionDraftValidationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsActionDraftValidationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsActionDraftValidationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsActionDraftValidationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsActionDraftValidationOK creates a GetIntegrationsActionDraftValidationOK with default headers values
func NewGetIntegrationsActionDraftValidationOK() *GetIntegrationsActionDraftValidationOK {
	return &GetIntegrationsActionDraftValidationOK{}
}

/*GetIntegrationsActionDraftValidationOK handles this case with default header values.

successful operation
*/
type GetIntegrationsActionDraftValidationOK struct {
	Payload *models.DraftValidationResult
}

func (o *GetIntegrationsActionDraftValidationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/validation][%d] getIntegrationsActionDraftValidationOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsActionDraftValidationOK) GetPayload() *models.DraftValidationResult {
	return o.Payload
}

func (o *GetIntegrationsActionDraftValidationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DraftValidationResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftValidationBadRequest creates a GetIntegrationsActionDraftValidationBadRequest with default headers values
func NewGetIntegrationsActionDraftValidationBadRequest() *GetIntegrationsActionDraftValidationBadRequest {
	return &GetIntegrationsActionDraftValidationBadRequest{}
}

/*GetIntegrationsActionDraftValidationBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsActionDraftValidationBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftValidationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/validation][%d] getIntegrationsActionDraftValidationBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsActionDraftValidationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftValidationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftValidationUnauthorized creates a GetIntegrationsActionDraftValidationUnauthorized with default headers values
func NewGetIntegrationsActionDraftValidationUnauthorized() *GetIntegrationsActionDraftValidationUnauthorized {
	return &GetIntegrationsActionDraftValidationUnauthorized{}
}

/*GetIntegrationsActionDraftValidationUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsActionDraftValidationUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftValidationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/validation][%d] getIntegrationsActionDraftValidationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsActionDraftValidationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftValidationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftValidationForbidden creates a GetIntegrationsActionDraftValidationForbidden with default headers values
func NewGetIntegrationsActionDraftValidationForbidden() *GetIntegrationsActionDraftValidationForbidden {
	return &GetIntegrationsActionDraftValidationForbidden{}
}

/*GetIntegrationsActionDraftValidationForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsActionDraftValidationForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftValidationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/validation][%d] getIntegrationsActionDraftValidationForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsActionDraftValidationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftValidationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftValidationNotFound creates a GetIntegrationsActionDraftValidationNotFound with default headers values
func NewGetIntegrationsActionDraftValidationNotFound() *GetIntegrationsActionDraftValidationNotFound {
	return &GetIntegrationsActionDraftValidationNotFound{}
}

/*GetIntegrationsActionDraftValidationNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetIntegrationsActionDraftValidationNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftValidationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/validation][%d] getIntegrationsActionDraftValidationNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsActionDraftValidationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftValidationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftValidationRequestTimeout creates a GetIntegrationsActionDraftValidationRequestTimeout with default headers values
func NewGetIntegrationsActionDraftValidationRequestTimeout() *GetIntegrationsActionDraftValidationRequestTimeout {
	return &GetIntegrationsActionDraftValidationRequestTimeout{}
}

/*GetIntegrationsActionDraftValidationRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetIntegrationsActionDraftValidationRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftValidationRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/validation][%d] getIntegrationsActionDraftValidationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsActionDraftValidationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftValidationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftValidationRequestEntityTooLarge creates a GetIntegrationsActionDraftValidationRequestEntityTooLarge with default headers values
func NewGetIntegrationsActionDraftValidationRequestEntityTooLarge() *GetIntegrationsActionDraftValidationRequestEntityTooLarge {
	return &GetIntegrationsActionDraftValidationRequestEntityTooLarge{}
}

/*GetIntegrationsActionDraftValidationRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetIntegrationsActionDraftValidationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftValidationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/validation][%d] getIntegrationsActionDraftValidationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsActionDraftValidationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftValidationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftValidationUnsupportedMediaType creates a GetIntegrationsActionDraftValidationUnsupportedMediaType with default headers values
func NewGetIntegrationsActionDraftValidationUnsupportedMediaType() *GetIntegrationsActionDraftValidationUnsupportedMediaType {
	return &GetIntegrationsActionDraftValidationUnsupportedMediaType{}
}

/*GetIntegrationsActionDraftValidationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsActionDraftValidationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftValidationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/validation][%d] getIntegrationsActionDraftValidationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsActionDraftValidationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftValidationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftValidationTooManyRequests creates a GetIntegrationsActionDraftValidationTooManyRequests with default headers values
func NewGetIntegrationsActionDraftValidationTooManyRequests() *GetIntegrationsActionDraftValidationTooManyRequests {
	return &GetIntegrationsActionDraftValidationTooManyRequests{}
}

/*GetIntegrationsActionDraftValidationTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetIntegrationsActionDraftValidationTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftValidationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/validation][%d] getIntegrationsActionDraftValidationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsActionDraftValidationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftValidationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftValidationInternalServerError creates a GetIntegrationsActionDraftValidationInternalServerError with default headers values
func NewGetIntegrationsActionDraftValidationInternalServerError() *GetIntegrationsActionDraftValidationInternalServerError {
	return &GetIntegrationsActionDraftValidationInternalServerError{}
}

/*GetIntegrationsActionDraftValidationInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsActionDraftValidationInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftValidationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/validation][%d] getIntegrationsActionDraftValidationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsActionDraftValidationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftValidationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftValidationServiceUnavailable creates a GetIntegrationsActionDraftValidationServiceUnavailable with default headers values
func NewGetIntegrationsActionDraftValidationServiceUnavailable() *GetIntegrationsActionDraftValidationServiceUnavailable {
	return &GetIntegrationsActionDraftValidationServiceUnavailable{}
}

/*GetIntegrationsActionDraftValidationServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsActionDraftValidationServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftValidationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/validation][%d] getIntegrationsActionDraftValidationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsActionDraftValidationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftValidationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftValidationGatewayTimeout creates a GetIntegrationsActionDraftValidationGatewayTimeout with default headers values
func NewGetIntegrationsActionDraftValidationGatewayTimeout() *GetIntegrationsActionDraftValidationGatewayTimeout {
	return &GetIntegrationsActionDraftValidationGatewayTimeout{}
}

/*GetIntegrationsActionDraftValidationGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetIntegrationsActionDraftValidationGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftValidationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/validation][%d] getIntegrationsActionDraftValidationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsActionDraftValidationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftValidationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
