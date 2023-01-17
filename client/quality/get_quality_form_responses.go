// Code generated by go-swagger; DO NOT EDIT.

package quality

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetQualityFormReader is a Reader for the GetQualityForm structure.
type GetQualityFormReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQualityFormReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQualityFormOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetQualityFormBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetQualityFormUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetQualityFormForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetQualityFormNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetQualityFormRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetQualityFormRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetQualityFormUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetQualityFormTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetQualityFormInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetQualityFormServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetQualityFormGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetQualityFormOK creates a GetQualityFormOK with default headers values
func NewGetQualityFormOK() *GetQualityFormOK {
	return &GetQualityFormOK{}
}

/*
GetQualityFormOK describes a response with status code 200, with default header values.

successful operation
*/
type GetQualityFormOK struct {
	Payload *models.EvaluationForm
}

// IsSuccess returns true when this get quality form o k response has a 2xx status code
func (o *GetQualityFormOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get quality form o k response has a 3xx status code
func (o *GetQualityFormOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality form o k response has a 4xx status code
func (o *GetQualityFormOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality form o k response has a 5xx status code
func (o *GetQualityFormOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality form o k response a status code equal to that given
func (o *GetQualityFormOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetQualityFormOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormOK  %+v", 200, o.Payload)
}

func (o *GetQualityFormOK) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormOK  %+v", 200, o.Payload)
}

func (o *GetQualityFormOK) GetPayload() *models.EvaluationForm {
	return o.Payload
}

func (o *GetQualityFormOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EvaluationForm)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormBadRequest creates a GetQualityFormBadRequest with default headers values
func NewGetQualityFormBadRequest() *GetQualityFormBadRequest {
	return &GetQualityFormBadRequest{}
}

/*
GetQualityFormBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetQualityFormBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality form bad request response has a 2xx status code
func (o *GetQualityFormBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality form bad request response has a 3xx status code
func (o *GetQualityFormBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality form bad request response has a 4xx status code
func (o *GetQualityFormBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality form bad request response has a 5xx status code
func (o *GetQualityFormBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality form bad request response a status code equal to that given
func (o *GetQualityFormBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetQualityFormBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityFormBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityFormBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormUnauthorized creates a GetQualityFormUnauthorized with default headers values
func NewGetQualityFormUnauthorized() *GetQualityFormUnauthorized {
	return &GetQualityFormUnauthorized{}
}

/*
GetQualityFormUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetQualityFormUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality form unauthorized response has a 2xx status code
func (o *GetQualityFormUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality form unauthorized response has a 3xx status code
func (o *GetQualityFormUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality form unauthorized response has a 4xx status code
func (o *GetQualityFormUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality form unauthorized response has a 5xx status code
func (o *GetQualityFormUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality form unauthorized response a status code equal to that given
func (o *GetQualityFormUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetQualityFormUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityFormUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityFormUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormForbidden creates a GetQualityFormForbidden with default headers values
func NewGetQualityFormForbidden() *GetQualityFormForbidden {
	return &GetQualityFormForbidden{}
}

/*
GetQualityFormForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetQualityFormForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality form forbidden response has a 2xx status code
func (o *GetQualityFormForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality form forbidden response has a 3xx status code
func (o *GetQualityFormForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality form forbidden response has a 4xx status code
func (o *GetQualityFormForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality form forbidden response has a 5xx status code
func (o *GetQualityFormForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality form forbidden response a status code equal to that given
func (o *GetQualityFormForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetQualityFormForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityFormForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityFormForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormNotFound creates a GetQualityFormNotFound with default headers values
func NewGetQualityFormNotFound() *GetQualityFormNotFound {
	return &GetQualityFormNotFound{}
}

/*
GetQualityFormNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetQualityFormNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality form not found response has a 2xx status code
func (o *GetQualityFormNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality form not found response has a 3xx status code
func (o *GetQualityFormNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality form not found response has a 4xx status code
func (o *GetQualityFormNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality form not found response has a 5xx status code
func (o *GetQualityFormNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality form not found response a status code equal to that given
func (o *GetQualityFormNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetQualityFormNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityFormNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityFormNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormRequestTimeout creates a GetQualityFormRequestTimeout with default headers values
func NewGetQualityFormRequestTimeout() *GetQualityFormRequestTimeout {
	return &GetQualityFormRequestTimeout{}
}

/*
GetQualityFormRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetQualityFormRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality form request timeout response has a 2xx status code
func (o *GetQualityFormRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality form request timeout response has a 3xx status code
func (o *GetQualityFormRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality form request timeout response has a 4xx status code
func (o *GetQualityFormRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality form request timeout response has a 5xx status code
func (o *GetQualityFormRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality form request timeout response a status code equal to that given
func (o *GetQualityFormRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetQualityFormRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualityFormRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualityFormRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormRequestEntityTooLarge creates a GetQualityFormRequestEntityTooLarge with default headers values
func NewGetQualityFormRequestEntityTooLarge() *GetQualityFormRequestEntityTooLarge {
	return &GetQualityFormRequestEntityTooLarge{}
}

/*
GetQualityFormRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetQualityFormRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality form request entity too large response has a 2xx status code
func (o *GetQualityFormRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality form request entity too large response has a 3xx status code
func (o *GetQualityFormRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality form request entity too large response has a 4xx status code
func (o *GetQualityFormRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality form request entity too large response has a 5xx status code
func (o *GetQualityFormRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality form request entity too large response a status code equal to that given
func (o *GetQualityFormRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetQualityFormRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityFormRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityFormRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormUnsupportedMediaType creates a GetQualityFormUnsupportedMediaType with default headers values
func NewGetQualityFormUnsupportedMediaType() *GetQualityFormUnsupportedMediaType {
	return &GetQualityFormUnsupportedMediaType{}
}

/*
GetQualityFormUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetQualityFormUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality form unsupported media type response has a 2xx status code
func (o *GetQualityFormUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality form unsupported media type response has a 3xx status code
func (o *GetQualityFormUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality form unsupported media type response has a 4xx status code
func (o *GetQualityFormUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality form unsupported media type response has a 5xx status code
func (o *GetQualityFormUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality form unsupported media type response a status code equal to that given
func (o *GetQualityFormUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetQualityFormUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityFormUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityFormUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormTooManyRequests creates a GetQualityFormTooManyRequests with default headers values
func NewGetQualityFormTooManyRequests() *GetQualityFormTooManyRequests {
	return &GetQualityFormTooManyRequests{}
}

/*
GetQualityFormTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetQualityFormTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality form too many requests response has a 2xx status code
func (o *GetQualityFormTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality form too many requests response has a 3xx status code
func (o *GetQualityFormTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality form too many requests response has a 4xx status code
func (o *GetQualityFormTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality form too many requests response has a 5xx status code
func (o *GetQualityFormTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality form too many requests response a status code equal to that given
func (o *GetQualityFormTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetQualityFormTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityFormTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityFormTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormInternalServerError creates a GetQualityFormInternalServerError with default headers values
func NewGetQualityFormInternalServerError() *GetQualityFormInternalServerError {
	return &GetQualityFormInternalServerError{}
}

/*
GetQualityFormInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetQualityFormInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality form internal server error response has a 2xx status code
func (o *GetQualityFormInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality form internal server error response has a 3xx status code
func (o *GetQualityFormInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality form internal server error response has a 4xx status code
func (o *GetQualityFormInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality form internal server error response has a 5xx status code
func (o *GetQualityFormInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality form internal server error response a status code equal to that given
func (o *GetQualityFormInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetQualityFormInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityFormInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityFormInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormServiceUnavailable creates a GetQualityFormServiceUnavailable with default headers values
func NewGetQualityFormServiceUnavailable() *GetQualityFormServiceUnavailable {
	return &GetQualityFormServiceUnavailable{}
}

/*
GetQualityFormServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetQualityFormServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality form service unavailable response has a 2xx status code
func (o *GetQualityFormServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality form service unavailable response has a 3xx status code
func (o *GetQualityFormServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality form service unavailable response has a 4xx status code
func (o *GetQualityFormServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality form service unavailable response has a 5xx status code
func (o *GetQualityFormServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality form service unavailable response a status code equal to that given
func (o *GetQualityFormServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetQualityFormServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityFormServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityFormServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormGatewayTimeout creates a GetQualityFormGatewayTimeout with default headers values
func NewGetQualityFormGatewayTimeout() *GetQualityFormGatewayTimeout {
	return &GetQualityFormGatewayTimeout{}
}

/*
GetQualityFormGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetQualityFormGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality form gateway timeout response has a 2xx status code
func (o *GetQualityFormGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality form gateway timeout response has a 3xx status code
func (o *GetQualityFormGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality form gateway timeout response has a 4xx status code
func (o *GetQualityFormGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality form gateway timeout response has a 5xx status code
func (o *GetQualityFormGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality form gateway timeout response a status code equal to that given
func (o *GetQualityFormGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetQualityFormGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityFormGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/{formId}][%d] getQualityFormGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityFormGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
