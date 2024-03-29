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

// GetQualityPublishedformReader is a Reader for the GetQualityPublishedform structure.
type GetQualityPublishedformReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQualityPublishedformReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQualityPublishedformOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetQualityPublishedformBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetQualityPublishedformUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetQualityPublishedformForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetQualityPublishedformNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetQualityPublishedformRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetQualityPublishedformRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetQualityPublishedformUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetQualityPublishedformTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetQualityPublishedformInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetQualityPublishedformServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetQualityPublishedformGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetQualityPublishedformOK creates a GetQualityPublishedformOK with default headers values
func NewGetQualityPublishedformOK() *GetQualityPublishedformOK {
	return &GetQualityPublishedformOK{}
}

/*
GetQualityPublishedformOK describes a response with status code 200, with default header values.

successful operation
*/
type GetQualityPublishedformOK struct {
	Payload *models.EvaluationForm
}

// IsSuccess returns true when this get quality publishedform o k response has a 2xx status code
func (o *GetQualityPublishedformOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get quality publishedform o k response has a 3xx status code
func (o *GetQualityPublishedformOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedform o k response has a 4xx status code
func (o *GetQualityPublishedformOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality publishedform o k response has a 5xx status code
func (o *GetQualityPublishedformOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedform o k response a status code equal to that given
func (o *GetQualityPublishedformOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetQualityPublishedformOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformOK  %+v", 200, o.Payload)
}

func (o *GetQualityPublishedformOK) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformOK  %+v", 200, o.Payload)
}

func (o *GetQualityPublishedformOK) GetPayload() *models.EvaluationForm {
	return o.Payload
}

func (o *GetQualityPublishedformOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EvaluationForm)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformBadRequest creates a GetQualityPublishedformBadRequest with default headers values
func NewGetQualityPublishedformBadRequest() *GetQualityPublishedformBadRequest {
	return &GetQualityPublishedformBadRequest{}
}

/*
GetQualityPublishedformBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetQualityPublishedformBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedform bad request response has a 2xx status code
func (o *GetQualityPublishedformBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedform bad request response has a 3xx status code
func (o *GetQualityPublishedformBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedform bad request response has a 4xx status code
func (o *GetQualityPublishedformBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedform bad request response has a 5xx status code
func (o *GetQualityPublishedformBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedform bad request response a status code equal to that given
func (o *GetQualityPublishedformBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetQualityPublishedformBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityPublishedformBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityPublishedformBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformUnauthorized creates a GetQualityPublishedformUnauthorized with default headers values
func NewGetQualityPublishedformUnauthorized() *GetQualityPublishedformUnauthorized {
	return &GetQualityPublishedformUnauthorized{}
}

/*
GetQualityPublishedformUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetQualityPublishedformUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedform unauthorized response has a 2xx status code
func (o *GetQualityPublishedformUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedform unauthorized response has a 3xx status code
func (o *GetQualityPublishedformUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedform unauthorized response has a 4xx status code
func (o *GetQualityPublishedformUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedform unauthorized response has a 5xx status code
func (o *GetQualityPublishedformUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedform unauthorized response a status code equal to that given
func (o *GetQualityPublishedformUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetQualityPublishedformUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityPublishedformUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityPublishedformUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformForbidden creates a GetQualityPublishedformForbidden with default headers values
func NewGetQualityPublishedformForbidden() *GetQualityPublishedformForbidden {
	return &GetQualityPublishedformForbidden{}
}

/*
GetQualityPublishedformForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetQualityPublishedformForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedform forbidden response has a 2xx status code
func (o *GetQualityPublishedformForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedform forbidden response has a 3xx status code
func (o *GetQualityPublishedformForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedform forbidden response has a 4xx status code
func (o *GetQualityPublishedformForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedform forbidden response has a 5xx status code
func (o *GetQualityPublishedformForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedform forbidden response a status code equal to that given
func (o *GetQualityPublishedformForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetQualityPublishedformForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityPublishedformForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityPublishedformForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformNotFound creates a GetQualityPublishedformNotFound with default headers values
func NewGetQualityPublishedformNotFound() *GetQualityPublishedformNotFound {
	return &GetQualityPublishedformNotFound{}
}

/*
GetQualityPublishedformNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetQualityPublishedformNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedform not found response has a 2xx status code
func (o *GetQualityPublishedformNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedform not found response has a 3xx status code
func (o *GetQualityPublishedformNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedform not found response has a 4xx status code
func (o *GetQualityPublishedformNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedform not found response has a 5xx status code
func (o *GetQualityPublishedformNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedform not found response a status code equal to that given
func (o *GetQualityPublishedformNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetQualityPublishedformNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityPublishedformNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityPublishedformNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformRequestTimeout creates a GetQualityPublishedformRequestTimeout with default headers values
func NewGetQualityPublishedformRequestTimeout() *GetQualityPublishedformRequestTimeout {
	return &GetQualityPublishedformRequestTimeout{}
}

/*
GetQualityPublishedformRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetQualityPublishedformRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedform request timeout response has a 2xx status code
func (o *GetQualityPublishedformRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedform request timeout response has a 3xx status code
func (o *GetQualityPublishedformRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedform request timeout response has a 4xx status code
func (o *GetQualityPublishedformRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedform request timeout response has a 5xx status code
func (o *GetQualityPublishedformRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedform request timeout response a status code equal to that given
func (o *GetQualityPublishedformRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetQualityPublishedformRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualityPublishedformRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualityPublishedformRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformRequestEntityTooLarge creates a GetQualityPublishedformRequestEntityTooLarge with default headers values
func NewGetQualityPublishedformRequestEntityTooLarge() *GetQualityPublishedformRequestEntityTooLarge {
	return &GetQualityPublishedformRequestEntityTooLarge{}
}

/*
GetQualityPublishedformRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetQualityPublishedformRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedform request entity too large response has a 2xx status code
func (o *GetQualityPublishedformRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedform request entity too large response has a 3xx status code
func (o *GetQualityPublishedformRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedform request entity too large response has a 4xx status code
func (o *GetQualityPublishedformRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedform request entity too large response has a 5xx status code
func (o *GetQualityPublishedformRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedform request entity too large response a status code equal to that given
func (o *GetQualityPublishedformRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetQualityPublishedformRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityPublishedformRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityPublishedformRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformUnsupportedMediaType creates a GetQualityPublishedformUnsupportedMediaType with default headers values
func NewGetQualityPublishedformUnsupportedMediaType() *GetQualityPublishedformUnsupportedMediaType {
	return &GetQualityPublishedformUnsupportedMediaType{}
}

/*
GetQualityPublishedformUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetQualityPublishedformUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedform unsupported media type response has a 2xx status code
func (o *GetQualityPublishedformUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedform unsupported media type response has a 3xx status code
func (o *GetQualityPublishedformUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedform unsupported media type response has a 4xx status code
func (o *GetQualityPublishedformUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedform unsupported media type response has a 5xx status code
func (o *GetQualityPublishedformUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedform unsupported media type response a status code equal to that given
func (o *GetQualityPublishedformUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetQualityPublishedformUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityPublishedformUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityPublishedformUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformTooManyRequests creates a GetQualityPublishedformTooManyRequests with default headers values
func NewGetQualityPublishedformTooManyRequests() *GetQualityPublishedformTooManyRequests {
	return &GetQualityPublishedformTooManyRequests{}
}

/*
GetQualityPublishedformTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetQualityPublishedformTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedform too many requests response has a 2xx status code
func (o *GetQualityPublishedformTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedform too many requests response has a 3xx status code
func (o *GetQualityPublishedformTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedform too many requests response has a 4xx status code
func (o *GetQualityPublishedformTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedform too many requests response has a 5xx status code
func (o *GetQualityPublishedformTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedform too many requests response a status code equal to that given
func (o *GetQualityPublishedformTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetQualityPublishedformTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityPublishedformTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityPublishedformTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformInternalServerError creates a GetQualityPublishedformInternalServerError with default headers values
func NewGetQualityPublishedformInternalServerError() *GetQualityPublishedformInternalServerError {
	return &GetQualityPublishedformInternalServerError{}
}

/*
GetQualityPublishedformInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetQualityPublishedformInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedform internal server error response has a 2xx status code
func (o *GetQualityPublishedformInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedform internal server error response has a 3xx status code
func (o *GetQualityPublishedformInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedform internal server error response has a 4xx status code
func (o *GetQualityPublishedformInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality publishedform internal server error response has a 5xx status code
func (o *GetQualityPublishedformInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality publishedform internal server error response a status code equal to that given
func (o *GetQualityPublishedformInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetQualityPublishedformInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityPublishedformInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityPublishedformInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformServiceUnavailable creates a GetQualityPublishedformServiceUnavailable with default headers values
func NewGetQualityPublishedformServiceUnavailable() *GetQualityPublishedformServiceUnavailable {
	return &GetQualityPublishedformServiceUnavailable{}
}

/*
GetQualityPublishedformServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetQualityPublishedformServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedform service unavailable response has a 2xx status code
func (o *GetQualityPublishedformServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedform service unavailable response has a 3xx status code
func (o *GetQualityPublishedformServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedform service unavailable response has a 4xx status code
func (o *GetQualityPublishedformServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality publishedform service unavailable response has a 5xx status code
func (o *GetQualityPublishedformServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality publishedform service unavailable response a status code equal to that given
func (o *GetQualityPublishedformServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetQualityPublishedformServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityPublishedformServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityPublishedformServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformGatewayTimeout creates a GetQualityPublishedformGatewayTimeout with default headers values
func NewGetQualityPublishedformGatewayTimeout() *GetQualityPublishedformGatewayTimeout {
	return &GetQualityPublishedformGatewayTimeout{}
}

/*
GetQualityPublishedformGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetQualityPublishedformGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedform gateway timeout response has a 2xx status code
func (o *GetQualityPublishedformGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedform gateway timeout response has a 3xx status code
func (o *GetQualityPublishedformGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedform gateway timeout response has a 4xx status code
func (o *GetQualityPublishedformGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality publishedform gateway timeout response has a 5xx status code
func (o *GetQualityPublishedformGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality publishedform gateway timeout response a status code equal to that given
func (o *GetQualityPublishedformGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetQualityPublishedformGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityPublishedformGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/{formId}][%d] getQualityPublishedformGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityPublishedformGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
