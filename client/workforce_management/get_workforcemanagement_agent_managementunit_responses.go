// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetWorkforcemanagementAgentManagementunitReader is a Reader for the GetWorkforcemanagementAgentManagementunit structure.
type GetWorkforcemanagementAgentManagementunitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementAgentManagementunitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementAgentManagementunitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementAgentManagementunitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementAgentManagementunitUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementAgentManagementunitForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementAgentManagementunitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementAgentManagementunitRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementAgentManagementunitRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementAgentManagementunitUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementAgentManagementunitTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementAgentManagementunitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementAgentManagementunitServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementAgentManagementunitGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementAgentManagementunitOK creates a GetWorkforcemanagementAgentManagementunitOK with default headers values
func NewGetWorkforcemanagementAgentManagementunitOK() *GetWorkforcemanagementAgentManagementunitOK {
	return &GetWorkforcemanagementAgentManagementunitOK{}
}

/*GetWorkforcemanagementAgentManagementunitOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementAgentManagementunitOK struct {
	Payload *models.AgentManagementUnitReference
}

func (o *GetWorkforcemanagementAgentManagementunitOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/managementunit][%d] getWorkforcemanagementAgentManagementunitOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementAgentManagementunitOK) GetPayload() *models.AgentManagementUnitReference {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentManagementunitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AgentManagementUnitReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentManagementunitBadRequest creates a GetWorkforcemanagementAgentManagementunitBadRequest with default headers values
func NewGetWorkforcemanagementAgentManagementunitBadRequest() *GetWorkforcemanagementAgentManagementunitBadRequest {
	return &GetWorkforcemanagementAgentManagementunitBadRequest{}
}

/*GetWorkforcemanagementAgentManagementunitBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementAgentManagementunitBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementAgentManagementunitBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/managementunit][%d] getWorkforcemanagementAgentManagementunitBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementAgentManagementunitBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentManagementunitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentManagementunitUnauthorized creates a GetWorkforcemanagementAgentManagementunitUnauthorized with default headers values
func NewGetWorkforcemanagementAgentManagementunitUnauthorized() *GetWorkforcemanagementAgentManagementunitUnauthorized {
	return &GetWorkforcemanagementAgentManagementunitUnauthorized{}
}

/*GetWorkforcemanagementAgentManagementunitUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementAgentManagementunitUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementAgentManagementunitUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/managementunit][%d] getWorkforcemanagementAgentManagementunitUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementAgentManagementunitUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentManagementunitUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentManagementunitForbidden creates a GetWorkforcemanagementAgentManagementunitForbidden with default headers values
func NewGetWorkforcemanagementAgentManagementunitForbidden() *GetWorkforcemanagementAgentManagementunitForbidden {
	return &GetWorkforcemanagementAgentManagementunitForbidden{}
}

/*GetWorkforcemanagementAgentManagementunitForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementAgentManagementunitForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementAgentManagementunitForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/managementunit][%d] getWorkforcemanagementAgentManagementunitForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementAgentManagementunitForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentManagementunitForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentManagementunitNotFound creates a GetWorkforcemanagementAgentManagementunitNotFound with default headers values
func NewGetWorkforcemanagementAgentManagementunitNotFound() *GetWorkforcemanagementAgentManagementunitNotFound {
	return &GetWorkforcemanagementAgentManagementunitNotFound{}
}

/*GetWorkforcemanagementAgentManagementunitNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementAgentManagementunitNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementAgentManagementunitNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/managementunit][%d] getWorkforcemanagementAgentManagementunitNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementAgentManagementunitNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentManagementunitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentManagementunitRequestTimeout creates a GetWorkforcemanagementAgentManagementunitRequestTimeout with default headers values
func NewGetWorkforcemanagementAgentManagementunitRequestTimeout() *GetWorkforcemanagementAgentManagementunitRequestTimeout {
	return &GetWorkforcemanagementAgentManagementunitRequestTimeout{}
}

/*GetWorkforcemanagementAgentManagementunitRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementAgentManagementunitRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementAgentManagementunitRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/managementunit][%d] getWorkforcemanagementAgentManagementunitRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementAgentManagementunitRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentManagementunitRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentManagementunitRequestEntityTooLarge creates a GetWorkforcemanagementAgentManagementunitRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementAgentManagementunitRequestEntityTooLarge() *GetWorkforcemanagementAgentManagementunitRequestEntityTooLarge {
	return &GetWorkforcemanagementAgentManagementunitRequestEntityTooLarge{}
}

/*GetWorkforcemanagementAgentManagementunitRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetWorkforcemanagementAgentManagementunitRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementAgentManagementunitRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/managementunit][%d] getWorkforcemanagementAgentManagementunitRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementAgentManagementunitRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentManagementunitRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentManagementunitUnsupportedMediaType creates a GetWorkforcemanagementAgentManagementunitUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementAgentManagementunitUnsupportedMediaType() *GetWorkforcemanagementAgentManagementunitUnsupportedMediaType {
	return &GetWorkforcemanagementAgentManagementunitUnsupportedMediaType{}
}

/*GetWorkforcemanagementAgentManagementunitUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementAgentManagementunitUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementAgentManagementunitUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/managementunit][%d] getWorkforcemanagementAgentManagementunitUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementAgentManagementunitUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentManagementunitUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentManagementunitTooManyRequests creates a GetWorkforcemanagementAgentManagementunitTooManyRequests with default headers values
func NewGetWorkforcemanagementAgentManagementunitTooManyRequests() *GetWorkforcemanagementAgentManagementunitTooManyRequests {
	return &GetWorkforcemanagementAgentManagementunitTooManyRequests{}
}

/*GetWorkforcemanagementAgentManagementunitTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementAgentManagementunitTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementAgentManagementunitTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/managementunit][%d] getWorkforcemanagementAgentManagementunitTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementAgentManagementunitTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentManagementunitTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentManagementunitInternalServerError creates a GetWorkforcemanagementAgentManagementunitInternalServerError with default headers values
func NewGetWorkforcemanagementAgentManagementunitInternalServerError() *GetWorkforcemanagementAgentManagementunitInternalServerError {
	return &GetWorkforcemanagementAgentManagementunitInternalServerError{}
}

/*GetWorkforcemanagementAgentManagementunitInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementAgentManagementunitInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementAgentManagementunitInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/managementunit][%d] getWorkforcemanagementAgentManagementunitInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementAgentManagementunitInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentManagementunitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentManagementunitServiceUnavailable creates a GetWorkforcemanagementAgentManagementunitServiceUnavailable with default headers values
func NewGetWorkforcemanagementAgentManagementunitServiceUnavailable() *GetWorkforcemanagementAgentManagementunitServiceUnavailable {
	return &GetWorkforcemanagementAgentManagementunitServiceUnavailable{}
}

/*GetWorkforcemanagementAgentManagementunitServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementAgentManagementunitServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementAgentManagementunitServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/managementunit][%d] getWorkforcemanagementAgentManagementunitServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementAgentManagementunitServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentManagementunitServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAgentManagementunitGatewayTimeout creates a GetWorkforcemanagementAgentManagementunitGatewayTimeout with default headers values
func NewGetWorkforcemanagementAgentManagementunitGatewayTimeout() *GetWorkforcemanagementAgentManagementunitGatewayTimeout {
	return &GetWorkforcemanagementAgentManagementunitGatewayTimeout{}
}

/*GetWorkforcemanagementAgentManagementunitGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementAgentManagementunitGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementAgentManagementunitGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/agents/{agentId}/managementunit][%d] getWorkforcemanagementAgentManagementunitGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementAgentManagementunitGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAgentManagementunitGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
