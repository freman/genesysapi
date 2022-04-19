// Code generated by go-swagger; DO NOT EDIT.

package scripts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetScriptsPublishedScriptIDVariablesReader is a Reader for the GetScriptsPublishedScriptIDVariables structure.
type GetScriptsPublishedScriptIDVariablesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScriptsPublishedScriptIDVariablesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScriptsPublishedScriptIDVariablesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetScriptsPublishedScriptIDVariablesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetScriptsPublishedScriptIDVariablesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScriptsPublishedScriptIDVariablesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScriptsPublishedScriptIDVariablesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetScriptsPublishedScriptIDVariablesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetScriptsPublishedScriptIDVariablesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetScriptsPublishedScriptIDVariablesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetScriptsPublishedScriptIDVariablesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScriptsPublishedScriptIDVariablesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetScriptsPublishedScriptIDVariablesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetScriptsPublishedScriptIDVariablesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScriptsPublishedScriptIDVariablesOK creates a GetScriptsPublishedScriptIDVariablesOK with default headers values
func NewGetScriptsPublishedScriptIDVariablesOK() *GetScriptsPublishedScriptIDVariablesOK {
	return &GetScriptsPublishedScriptIDVariablesOK{}
}

/*GetScriptsPublishedScriptIDVariablesOK handles this case with default header values.

successful operation
*/
type GetScriptsPublishedScriptIDVariablesOK struct {
	Payload interface{}
}

func (o *GetScriptsPublishedScriptIDVariablesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/variables][%d] getScriptsPublishedScriptIdVariablesOK  %+v", 200, o.Payload)
}

func (o *GetScriptsPublishedScriptIDVariablesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDVariablesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDVariablesBadRequest creates a GetScriptsPublishedScriptIDVariablesBadRequest with default headers values
func NewGetScriptsPublishedScriptIDVariablesBadRequest() *GetScriptsPublishedScriptIDVariablesBadRequest {
	return &GetScriptsPublishedScriptIDVariablesBadRequest{}
}

/*GetScriptsPublishedScriptIDVariablesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetScriptsPublishedScriptIDVariablesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDVariablesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/variables][%d] getScriptsPublishedScriptIdVariablesBadRequest  %+v", 400, o.Payload)
}

func (o *GetScriptsPublishedScriptIDVariablesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDVariablesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDVariablesUnauthorized creates a GetScriptsPublishedScriptIDVariablesUnauthorized with default headers values
func NewGetScriptsPublishedScriptIDVariablesUnauthorized() *GetScriptsPublishedScriptIDVariablesUnauthorized {
	return &GetScriptsPublishedScriptIDVariablesUnauthorized{}
}

/*GetScriptsPublishedScriptIDVariablesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetScriptsPublishedScriptIDVariablesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDVariablesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/variables][%d] getScriptsPublishedScriptIdVariablesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScriptsPublishedScriptIDVariablesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDVariablesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDVariablesForbidden creates a GetScriptsPublishedScriptIDVariablesForbidden with default headers values
func NewGetScriptsPublishedScriptIDVariablesForbidden() *GetScriptsPublishedScriptIDVariablesForbidden {
	return &GetScriptsPublishedScriptIDVariablesForbidden{}
}

/*GetScriptsPublishedScriptIDVariablesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetScriptsPublishedScriptIDVariablesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDVariablesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/variables][%d] getScriptsPublishedScriptIdVariablesForbidden  %+v", 403, o.Payload)
}

func (o *GetScriptsPublishedScriptIDVariablesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDVariablesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDVariablesNotFound creates a GetScriptsPublishedScriptIDVariablesNotFound with default headers values
func NewGetScriptsPublishedScriptIDVariablesNotFound() *GetScriptsPublishedScriptIDVariablesNotFound {
	return &GetScriptsPublishedScriptIDVariablesNotFound{}
}

/*GetScriptsPublishedScriptIDVariablesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetScriptsPublishedScriptIDVariablesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDVariablesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/variables][%d] getScriptsPublishedScriptIdVariablesNotFound  %+v", 404, o.Payload)
}

func (o *GetScriptsPublishedScriptIDVariablesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDVariablesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDVariablesRequestTimeout creates a GetScriptsPublishedScriptIDVariablesRequestTimeout with default headers values
func NewGetScriptsPublishedScriptIDVariablesRequestTimeout() *GetScriptsPublishedScriptIDVariablesRequestTimeout {
	return &GetScriptsPublishedScriptIDVariablesRequestTimeout{}
}

/*GetScriptsPublishedScriptIDVariablesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetScriptsPublishedScriptIDVariablesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDVariablesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/variables][%d] getScriptsPublishedScriptIdVariablesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetScriptsPublishedScriptIDVariablesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDVariablesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDVariablesRequestEntityTooLarge creates a GetScriptsPublishedScriptIDVariablesRequestEntityTooLarge with default headers values
func NewGetScriptsPublishedScriptIDVariablesRequestEntityTooLarge() *GetScriptsPublishedScriptIDVariablesRequestEntityTooLarge {
	return &GetScriptsPublishedScriptIDVariablesRequestEntityTooLarge{}
}

/*GetScriptsPublishedScriptIDVariablesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetScriptsPublishedScriptIDVariablesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDVariablesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/variables][%d] getScriptsPublishedScriptIdVariablesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetScriptsPublishedScriptIDVariablesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDVariablesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDVariablesUnsupportedMediaType creates a GetScriptsPublishedScriptIDVariablesUnsupportedMediaType with default headers values
func NewGetScriptsPublishedScriptIDVariablesUnsupportedMediaType() *GetScriptsPublishedScriptIDVariablesUnsupportedMediaType {
	return &GetScriptsPublishedScriptIDVariablesUnsupportedMediaType{}
}

/*GetScriptsPublishedScriptIDVariablesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetScriptsPublishedScriptIDVariablesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDVariablesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/variables][%d] getScriptsPublishedScriptIdVariablesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetScriptsPublishedScriptIDVariablesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDVariablesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDVariablesTooManyRequests creates a GetScriptsPublishedScriptIDVariablesTooManyRequests with default headers values
func NewGetScriptsPublishedScriptIDVariablesTooManyRequests() *GetScriptsPublishedScriptIDVariablesTooManyRequests {
	return &GetScriptsPublishedScriptIDVariablesTooManyRequests{}
}

/*GetScriptsPublishedScriptIDVariablesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetScriptsPublishedScriptIDVariablesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDVariablesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/variables][%d] getScriptsPublishedScriptIdVariablesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScriptsPublishedScriptIDVariablesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDVariablesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDVariablesInternalServerError creates a GetScriptsPublishedScriptIDVariablesInternalServerError with default headers values
func NewGetScriptsPublishedScriptIDVariablesInternalServerError() *GetScriptsPublishedScriptIDVariablesInternalServerError {
	return &GetScriptsPublishedScriptIDVariablesInternalServerError{}
}

/*GetScriptsPublishedScriptIDVariablesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetScriptsPublishedScriptIDVariablesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDVariablesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/variables][%d] getScriptsPublishedScriptIdVariablesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScriptsPublishedScriptIDVariablesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDVariablesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDVariablesServiceUnavailable creates a GetScriptsPublishedScriptIDVariablesServiceUnavailable with default headers values
func NewGetScriptsPublishedScriptIDVariablesServiceUnavailable() *GetScriptsPublishedScriptIDVariablesServiceUnavailable {
	return &GetScriptsPublishedScriptIDVariablesServiceUnavailable{}
}

/*GetScriptsPublishedScriptIDVariablesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetScriptsPublishedScriptIDVariablesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDVariablesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/variables][%d] getScriptsPublishedScriptIdVariablesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetScriptsPublishedScriptIDVariablesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDVariablesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDVariablesGatewayTimeout creates a GetScriptsPublishedScriptIDVariablesGatewayTimeout with default headers values
func NewGetScriptsPublishedScriptIDVariablesGatewayTimeout() *GetScriptsPublishedScriptIDVariablesGatewayTimeout {
	return &GetScriptsPublishedScriptIDVariablesGatewayTimeout{}
}

/*GetScriptsPublishedScriptIDVariablesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetScriptsPublishedScriptIDVariablesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDVariablesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/variables][%d] getScriptsPublishedScriptIdVariablesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetScriptsPublishedScriptIDVariablesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDVariablesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
