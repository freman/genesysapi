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

// PostIntegrationsActionTestReader is a Reader for the PostIntegrationsActionTest structure.
type PostIntegrationsActionTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIntegrationsActionTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostIntegrationsActionTestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostIntegrationsActionTestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostIntegrationsActionTestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostIntegrationsActionTestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostIntegrationsActionTestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostIntegrationsActionTestRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostIntegrationsActionTestRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostIntegrationsActionTestUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostIntegrationsActionTestTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostIntegrationsActionTestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostIntegrationsActionTestServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostIntegrationsActionTestGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostIntegrationsActionTestOK creates a PostIntegrationsActionTestOK with default headers values
func NewPostIntegrationsActionTestOK() *PostIntegrationsActionTestOK {
	return &PostIntegrationsActionTestOK{}
}

/*PostIntegrationsActionTestOK handles this case with default header values.

successful operation
*/
type PostIntegrationsActionTestOK struct {
	Payload *models.TestExecutionResult
}

func (o *PostIntegrationsActionTestOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/test][%d] postIntegrationsActionTestOK  %+v", 200, o.Payload)
}

func (o *PostIntegrationsActionTestOK) GetPayload() *models.TestExecutionResult {
	return o.Payload
}

func (o *PostIntegrationsActionTestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TestExecutionResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionTestBadRequest creates a PostIntegrationsActionTestBadRequest with default headers values
func NewPostIntegrationsActionTestBadRequest() *PostIntegrationsActionTestBadRequest {
	return &PostIntegrationsActionTestBadRequest{}
}

/*PostIntegrationsActionTestBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostIntegrationsActionTestBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionTestBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/test][%d] postIntegrationsActionTestBadRequest  %+v", 400, o.Payload)
}

func (o *PostIntegrationsActionTestBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionTestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionTestUnauthorized creates a PostIntegrationsActionTestUnauthorized with default headers values
func NewPostIntegrationsActionTestUnauthorized() *PostIntegrationsActionTestUnauthorized {
	return &PostIntegrationsActionTestUnauthorized{}
}

/*PostIntegrationsActionTestUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostIntegrationsActionTestUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionTestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/test][%d] postIntegrationsActionTestUnauthorized  %+v", 401, o.Payload)
}

func (o *PostIntegrationsActionTestUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionTestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionTestForbidden creates a PostIntegrationsActionTestForbidden with default headers values
func NewPostIntegrationsActionTestForbidden() *PostIntegrationsActionTestForbidden {
	return &PostIntegrationsActionTestForbidden{}
}

/*PostIntegrationsActionTestForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostIntegrationsActionTestForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionTestForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/test][%d] postIntegrationsActionTestForbidden  %+v", 403, o.Payload)
}

func (o *PostIntegrationsActionTestForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionTestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionTestNotFound creates a PostIntegrationsActionTestNotFound with default headers values
func NewPostIntegrationsActionTestNotFound() *PostIntegrationsActionTestNotFound {
	return &PostIntegrationsActionTestNotFound{}
}

/*PostIntegrationsActionTestNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostIntegrationsActionTestNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionTestNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/test][%d] postIntegrationsActionTestNotFound  %+v", 404, o.Payload)
}

func (o *PostIntegrationsActionTestNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionTestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionTestRequestTimeout creates a PostIntegrationsActionTestRequestTimeout with default headers values
func NewPostIntegrationsActionTestRequestTimeout() *PostIntegrationsActionTestRequestTimeout {
	return &PostIntegrationsActionTestRequestTimeout{}
}

/*PostIntegrationsActionTestRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostIntegrationsActionTestRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionTestRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/test][%d] postIntegrationsActionTestRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostIntegrationsActionTestRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionTestRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionTestRequestEntityTooLarge creates a PostIntegrationsActionTestRequestEntityTooLarge with default headers values
func NewPostIntegrationsActionTestRequestEntityTooLarge() *PostIntegrationsActionTestRequestEntityTooLarge {
	return &PostIntegrationsActionTestRequestEntityTooLarge{}
}

/*PostIntegrationsActionTestRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostIntegrationsActionTestRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionTestRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/test][%d] postIntegrationsActionTestRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostIntegrationsActionTestRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionTestRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionTestUnsupportedMediaType creates a PostIntegrationsActionTestUnsupportedMediaType with default headers values
func NewPostIntegrationsActionTestUnsupportedMediaType() *PostIntegrationsActionTestUnsupportedMediaType {
	return &PostIntegrationsActionTestUnsupportedMediaType{}
}

/*PostIntegrationsActionTestUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostIntegrationsActionTestUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionTestUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/test][%d] postIntegrationsActionTestUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostIntegrationsActionTestUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionTestUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionTestTooManyRequests creates a PostIntegrationsActionTestTooManyRequests with default headers values
func NewPostIntegrationsActionTestTooManyRequests() *PostIntegrationsActionTestTooManyRequests {
	return &PostIntegrationsActionTestTooManyRequests{}
}

/*PostIntegrationsActionTestTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostIntegrationsActionTestTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionTestTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/test][%d] postIntegrationsActionTestTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostIntegrationsActionTestTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionTestTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionTestInternalServerError creates a PostIntegrationsActionTestInternalServerError with default headers values
func NewPostIntegrationsActionTestInternalServerError() *PostIntegrationsActionTestInternalServerError {
	return &PostIntegrationsActionTestInternalServerError{}
}

/*PostIntegrationsActionTestInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostIntegrationsActionTestInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionTestInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/test][%d] postIntegrationsActionTestInternalServerError  %+v", 500, o.Payload)
}

func (o *PostIntegrationsActionTestInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionTestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionTestServiceUnavailable creates a PostIntegrationsActionTestServiceUnavailable with default headers values
func NewPostIntegrationsActionTestServiceUnavailable() *PostIntegrationsActionTestServiceUnavailable {
	return &PostIntegrationsActionTestServiceUnavailable{}
}

/*PostIntegrationsActionTestServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostIntegrationsActionTestServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionTestServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/test][%d] postIntegrationsActionTestServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostIntegrationsActionTestServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionTestServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionTestGatewayTimeout creates a PostIntegrationsActionTestGatewayTimeout with default headers values
func NewPostIntegrationsActionTestGatewayTimeout() *PostIntegrationsActionTestGatewayTimeout {
	return &PostIntegrationsActionTestGatewayTimeout{}
}

/*PostIntegrationsActionTestGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostIntegrationsActionTestGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionTestGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/test][%d] postIntegrationsActionTestGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostIntegrationsActionTestGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionTestGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
