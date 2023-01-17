// Code generated by go-swagger; DO NOT EDIT.

package process_automation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostProcessautomationTriggerTestReader is a Reader for the PostProcessautomationTriggerTest structure.
type PostProcessautomationTriggerTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostProcessautomationTriggerTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostProcessautomationTriggerTestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostProcessautomationTriggerTestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostProcessautomationTriggerTestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostProcessautomationTriggerTestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostProcessautomationTriggerTestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostProcessautomationTriggerTestRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostProcessautomationTriggerTestRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostProcessautomationTriggerTestUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostProcessautomationTriggerTestTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostProcessautomationTriggerTestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostProcessautomationTriggerTestServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostProcessautomationTriggerTestGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostProcessautomationTriggerTestOK creates a PostProcessautomationTriggerTestOK with default headers values
func NewPostProcessautomationTriggerTestOK() *PostProcessautomationTriggerTestOK {
	return &PostProcessautomationTriggerTestOK{}
}

/*
PostProcessautomationTriggerTestOK describes a response with status code 200, with default header values.

successful operation
*/
type PostProcessautomationTriggerTestOK struct {
	Payload *models.TestModeResults
}

// IsSuccess returns true when this post processautomation trigger test o k response has a 2xx status code
func (o *PostProcessautomationTriggerTestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post processautomation trigger test o k response has a 3xx status code
func (o *PostProcessautomationTriggerTestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post processautomation trigger test o k response has a 4xx status code
func (o *PostProcessautomationTriggerTestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post processautomation trigger test o k response has a 5xx status code
func (o *PostProcessautomationTriggerTestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post processautomation trigger test o k response a status code equal to that given
func (o *PostProcessautomationTriggerTestOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostProcessautomationTriggerTestOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestOK  %+v", 200, o.Payload)
}

func (o *PostProcessautomationTriggerTestOK) String() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestOK  %+v", 200, o.Payload)
}

func (o *PostProcessautomationTriggerTestOK) GetPayload() *models.TestModeResults {
	return o.Payload
}

func (o *PostProcessautomationTriggerTestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TestModeResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProcessautomationTriggerTestBadRequest creates a PostProcessautomationTriggerTestBadRequest with default headers values
func NewPostProcessautomationTriggerTestBadRequest() *PostProcessautomationTriggerTestBadRequest {
	return &PostProcessautomationTriggerTestBadRequest{}
}

/*
PostProcessautomationTriggerTestBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostProcessautomationTriggerTestBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post processautomation trigger test bad request response has a 2xx status code
func (o *PostProcessautomationTriggerTestBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post processautomation trigger test bad request response has a 3xx status code
func (o *PostProcessautomationTriggerTestBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post processautomation trigger test bad request response has a 4xx status code
func (o *PostProcessautomationTriggerTestBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post processautomation trigger test bad request response has a 5xx status code
func (o *PostProcessautomationTriggerTestBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post processautomation trigger test bad request response a status code equal to that given
func (o *PostProcessautomationTriggerTestBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostProcessautomationTriggerTestBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestBadRequest  %+v", 400, o.Payload)
}

func (o *PostProcessautomationTriggerTestBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestBadRequest  %+v", 400, o.Payload)
}

func (o *PostProcessautomationTriggerTestBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostProcessautomationTriggerTestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProcessautomationTriggerTestUnauthorized creates a PostProcessautomationTriggerTestUnauthorized with default headers values
func NewPostProcessautomationTriggerTestUnauthorized() *PostProcessautomationTriggerTestUnauthorized {
	return &PostProcessautomationTriggerTestUnauthorized{}
}

/*
PostProcessautomationTriggerTestUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostProcessautomationTriggerTestUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post processautomation trigger test unauthorized response has a 2xx status code
func (o *PostProcessautomationTriggerTestUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post processautomation trigger test unauthorized response has a 3xx status code
func (o *PostProcessautomationTriggerTestUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post processautomation trigger test unauthorized response has a 4xx status code
func (o *PostProcessautomationTriggerTestUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post processautomation trigger test unauthorized response has a 5xx status code
func (o *PostProcessautomationTriggerTestUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post processautomation trigger test unauthorized response a status code equal to that given
func (o *PostProcessautomationTriggerTestUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostProcessautomationTriggerTestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestUnauthorized  %+v", 401, o.Payload)
}

func (o *PostProcessautomationTriggerTestUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestUnauthorized  %+v", 401, o.Payload)
}

func (o *PostProcessautomationTriggerTestUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostProcessautomationTriggerTestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProcessautomationTriggerTestForbidden creates a PostProcessautomationTriggerTestForbidden with default headers values
func NewPostProcessautomationTriggerTestForbidden() *PostProcessautomationTriggerTestForbidden {
	return &PostProcessautomationTriggerTestForbidden{}
}

/*
PostProcessautomationTriggerTestForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostProcessautomationTriggerTestForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post processautomation trigger test forbidden response has a 2xx status code
func (o *PostProcessautomationTriggerTestForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post processautomation trigger test forbidden response has a 3xx status code
func (o *PostProcessautomationTriggerTestForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post processautomation trigger test forbidden response has a 4xx status code
func (o *PostProcessautomationTriggerTestForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post processautomation trigger test forbidden response has a 5xx status code
func (o *PostProcessautomationTriggerTestForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post processautomation trigger test forbidden response a status code equal to that given
func (o *PostProcessautomationTriggerTestForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostProcessautomationTriggerTestForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestForbidden  %+v", 403, o.Payload)
}

func (o *PostProcessautomationTriggerTestForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestForbidden  %+v", 403, o.Payload)
}

func (o *PostProcessautomationTriggerTestForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostProcessautomationTriggerTestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProcessautomationTriggerTestNotFound creates a PostProcessautomationTriggerTestNotFound with default headers values
func NewPostProcessautomationTriggerTestNotFound() *PostProcessautomationTriggerTestNotFound {
	return &PostProcessautomationTriggerTestNotFound{}
}

/*
PostProcessautomationTriggerTestNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostProcessautomationTriggerTestNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post processautomation trigger test not found response has a 2xx status code
func (o *PostProcessautomationTriggerTestNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post processautomation trigger test not found response has a 3xx status code
func (o *PostProcessautomationTriggerTestNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post processautomation trigger test not found response has a 4xx status code
func (o *PostProcessautomationTriggerTestNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post processautomation trigger test not found response has a 5xx status code
func (o *PostProcessautomationTriggerTestNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post processautomation trigger test not found response a status code equal to that given
func (o *PostProcessautomationTriggerTestNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostProcessautomationTriggerTestNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestNotFound  %+v", 404, o.Payload)
}

func (o *PostProcessautomationTriggerTestNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestNotFound  %+v", 404, o.Payload)
}

func (o *PostProcessautomationTriggerTestNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostProcessautomationTriggerTestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProcessautomationTriggerTestRequestTimeout creates a PostProcessautomationTriggerTestRequestTimeout with default headers values
func NewPostProcessautomationTriggerTestRequestTimeout() *PostProcessautomationTriggerTestRequestTimeout {
	return &PostProcessautomationTriggerTestRequestTimeout{}
}

/*
PostProcessautomationTriggerTestRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostProcessautomationTriggerTestRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post processautomation trigger test request timeout response has a 2xx status code
func (o *PostProcessautomationTriggerTestRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post processautomation trigger test request timeout response has a 3xx status code
func (o *PostProcessautomationTriggerTestRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post processautomation trigger test request timeout response has a 4xx status code
func (o *PostProcessautomationTriggerTestRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post processautomation trigger test request timeout response has a 5xx status code
func (o *PostProcessautomationTriggerTestRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post processautomation trigger test request timeout response a status code equal to that given
func (o *PostProcessautomationTriggerTestRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostProcessautomationTriggerTestRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostProcessautomationTriggerTestRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostProcessautomationTriggerTestRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostProcessautomationTriggerTestRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProcessautomationTriggerTestRequestEntityTooLarge creates a PostProcessautomationTriggerTestRequestEntityTooLarge with default headers values
func NewPostProcessautomationTriggerTestRequestEntityTooLarge() *PostProcessautomationTriggerTestRequestEntityTooLarge {
	return &PostProcessautomationTriggerTestRequestEntityTooLarge{}
}

/*
PostProcessautomationTriggerTestRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostProcessautomationTriggerTestRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post processautomation trigger test request entity too large response has a 2xx status code
func (o *PostProcessautomationTriggerTestRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post processautomation trigger test request entity too large response has a 3xx status code
func (o *PostProcessautomationTriggerTestRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post processautomation trigger test request entity too large response has a 4xx status code
func (o *PostProcessautomationTriggerTestRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post processautomation trigger test request entity too large response has a 5xx status code
func (o *PostProcessautomationTriggerTestRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post processautomation trigger test request entity too large response a status code equal to that given
func (o *PostProcessautomationTriggerTestRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostProcessautomationTriggerTestRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostProcessautomationTriggerTestRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostProcessautomationTriggerTestRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostProcessautomationTriggerTestRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProcessautomationTriggerTestUnsupportedMediaType creates a PostProcessautomationTriggerTestUnsupportedMediaType with default headers values
func NewPostProcessautomationTriggerTestUnsupportedMediaType() *PostProcessautomationTriggerTestUnsupportedMediaType {
	return &PostProcessautomationTriggerTestUnsupportedMediaType{}
}

/*
PostProcessautomationTriggerTestUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostProcessautomationTriggerTestUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post processautomation trigger test unsupported media type response has a 2xx status code
func (o *PostProcessautomationTriggerTestUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post processautomation trigger test unsupported media type response has a 3xx status code
func (o *PostProcessautomationTriggerTestUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post processautomation trigger test unsupported media type response has a 4xx status code
func (o *PostProcessautomationTriggerTestUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post processautomation trigger test unsupported media type response has a 5xx status code
func (o *PostProcessautomationTriggerTestUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post processautomation trigger test unsupported media type response a status code equal to that given
func (o *PostProcessautomationTriggerTestUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostProcessautomationTriggerTestUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostProcessautomationTriggerTestUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostProcessautomationTriggerTestUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostProcessautomationTriggerTestUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProcessautomationTriggerTestTooManyRequests creates a PostProcessautomationTriggerTestTooManyRequests with default headers values
func NewPostProcessautomationTriggerTestTooManyRequests() *PostProcessautomationTriggerTestTooManyRequests {
	return &PostProcessautomationTriggerTestTooManyRequests{}
}

/*
PostProcessautomationTriggerTestTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostProcessautomationTriggerTestTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post processautomation trigger test too many requests response has a 2xx status code
func (o *PostProcessautomationTriggerTestTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post processautomation trigger test too many requests response has a 3xx status code
func (o *PostProcessautomationTriggerTestTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post processautomation trigger test too many requests response has a 4xx status code
func (o *PostProcessautomationTriggerTestTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post processautomation trigger test too many requests response has a 5xx status code
func (o *PostProcessautomationTriggerTestTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post processautomation trigger test too many requests response a status code equal to that given
func (o *PostProcessautomationTriggerTestTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostProcessautomationTriggerTestTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostProcessautomationTriggerTestTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostProcessautomationTriggerTestTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostProcessautomationTriggerTestTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProcessautomationTriggerTestInternalServerError creates a PostProcessautomationTriggerTestInternalServerError with default headers values
func NewPostProcessautomationTriggerTestInternalServerError() *PostProcessautomationTriggerTestInternalServerError {
	return &PostProcessautomationTriggerTestInternalServerError{}
}

/*
PostProcessautomationTriggerTestInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostProcessautomationTriggerTestInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post processautomation trigger test internal server error response has a 2xx status code
func (o *PostProcessautomationTriggerTestInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post processautomation trigger test internal server error response has a 3xx status code
func (o *PostProcessautomationTriggerTestInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post processautomation trigger test internal server error response has a 4xx status code
func (o *PostProcessautomationTriggerTestInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post processautomation trigger test internal server error response has a 5xx status code
func (o *PostProcessautomationTriggerTestInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post processautomation trigger test internal server error response a status code equal to that given
func (o *PostProcessautomationTriggerTestInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostProcessautomationTriggerTestInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestInternalServerError  %+v", 500, o.Payload)
}

func (o *PostProcessautomationTriggerTestInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestInternalServerError  %+v", 500, o.Payload)
}

func (o *PostProcessautomationTriggerTestInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostProcessautomationTriggerTestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProcessautomationTriggerTestServiceUnavailable creates a PostProcessautomationTriggerTestServiceUnavailable with default headers values
func NewPostProcessautomationTriggerTestServiceUnavailable() *PostProcessautomationTriggerTestServiceUnavailable {
	return &PostProcessautomationTriggerTestServiceUnavailable{}
}

/*
PostProcessautomationTriggerTestServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostProcessautomationTriggerTestServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post processautomation trigger test service unavailable response has a 2xx status code
func (o *PostProcessautomationTriggerTestServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post processautomation trigger test service unavailable response has a 3xx status code
func (o *PostProcessautomationTriggerTestServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post processautomation trigger test service unavailable response has a 4xx status code
func (o *PostProcessautomationTriggerTestServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post processautomation trigger test service unavailable response has a 5xx status code
func (o *PostProcessautomationTriggerTestServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post processautomation trigger test service unavailable response a status code equal to that given
func (o *PostProcessautomationTriggerTestServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostProcessautomationTriggerTestServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostProcessautomationTriggerTestServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostProcessautomationTriggerTestServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostProcessautomationTriggerTestServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProcessautomationTriggerTestGatewayTimeout creates a PostProcessautomationTriggerTestGatewayTimeout with default headers values
func NewPostProcessautomationTriggerTestGatewayTimeout() *PostProcessautomationTriggerTestGatewayTimeout {
	return &PostProcessautomationTriggerTestGatewayTimeout{}
}

/*
PostProcessautomationTriggerTestGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostProcessautomationTriggerTestGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post processautomation trigger test gateway timeout response has a 2xx status code
func (o *PostProcessautomationTriggerTestGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post processautomation trigger test gateway timeout response has a 3xx status code
func (o *PostProcessautomationTriggerTestGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post processautomation trigger test gateway timeout response has a 4xx status code
func (o *PostProcessautomationTriggerTestGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post processautomation trigger test gateway timeout response has a 5xx status code
func (o *PostProcessautomationTriggerTestGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post processautomation trigger test gateway timeout response a status code equal to that given
func (o *PostProcessautomationTriggerTestGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostProcessautomationTriggerTestGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostProcessautomationTriggerTestGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/processautomation/triggers/{triggerId}/test][%d] postProcessautomationTriggerTestGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostProcessautomationTriggerTestGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostProcessautomationTriggerTestGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
