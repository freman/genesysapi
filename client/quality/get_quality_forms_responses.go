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

// GetQualityFormsReader is a Reader for the GetQualityForms structure.
type GetQualityFormsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQualityFormsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQualityFormsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetQualityFormsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetQualityFormsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetQualityFormsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetQualityFormsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetQualityFormsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetQualityFormsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetQualityFormsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetQualityFormsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetQualityFormsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetQualityFormsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetQualityFormsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetQualityFormsOK creates a GetQualityFormsOK with default headers values
func NewGetQualityFormsOK() *GetQualityFormsOK {
	return &GetQualityFormsOK{}
}

/*
GetQualityFormsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetQualityFormsOK struct {
	Payload *models.EvaluationFormEntityListing
}

// IsSuccess returns true when this get quality forms o k response has a 2xx status code
func (o *GetQualityFormsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get quality forms o k response has a 3xx status code
func (o *GetQualityFormsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms o k response has a 4xx status code
func (o *GetQualityFormsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality forms o k response has a 5xx status code
func (o *GetQualityFormsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality forms o k response a status code equal to that given
func (o *GetQualityFormsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetQualityFormsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsOK  %+v", 200, o.Payload)
}

func (o *GetQualityFormsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsOK  %+v", 200, o.Payload)
}

func (o *GetQualityFormsOK) GetPayload() *models.EvaluationFormEntityListing {
	return o.Payload
}

func (o *GetQualityFormsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EvaluationFormEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsBadRequest creates a GetQualityFormsBadRequest with default headers values
func NewGetQualityFormsBadRequest() *GetQualityFormsBadRequest {
	return &GetQualityFormsBadRequest{}
}

/*
GetQualityFormsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetQualityFormsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality forms bad request response has a 2xx status code
func (o *GetQualityFormsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality forms bad request response has a 3xx status code
func (o *GetQualityFormsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms bad request response has a 4xx status code
func (o *GetQualityFormsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality forms bad request response has a 5xx status code
func (o *GetQualityFormsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality forms bad request response a status code equal to that given
func (o *GetQualityFormsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetQualityFormsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityFormsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityFormsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsUnauthorized creates a GetQualityFormsUnauthorized with default headers values
func NewGetQualityFormsUnauthorized() *GetQualityFormsUnauthorized {
	return &GetQualityFormsUnauthorized{}
}

/*
GetQualityFormsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetQualityFormsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality forms unauthorized response has a 2xx status code
func (o *GetQualityFormsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality forms unauthorized response has a 3xx status code
func (o *GetQualityFormsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms unauthorized response has a 4xx status code
func (o *GetQualityFormsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality forms unauthorized response has a 5xx status code
func (o *GetQualityFormsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality forms unauthorized response a status code equal to that given
func (o *GetQualityFormsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetQualityFormsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityFormsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityFormsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsForbidden creates a GetQualityFormsForbidden with default headers values
func NewGetQualityFormsForbidden() *GetQualityFormsForbidden {
	return &GetQualityFormsForbidden{}
}

/*
GetQualityFormsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetQualityFormsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality forms forbidden response has a 2xx status code
func (o *GetQualityFormsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality forms forbidden response has a 3xx status code
func (o *GetQualityFormsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms forbidden response has a 4xx status code
func (o *GetQualityFormsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality forms forbidden response has a 5xx status code
func (o *GetQualityFormsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality forms forbidden response a status code equal to that given
func (o *GetQualityFormsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetQualityFormsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityFormsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityFormsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsNotFound creates a GetQualityFormsNotFound with default headers values
func NewGetQualityFormsNotFound() *GetQualityFormsNotFound {
	return &GetQualityFormsNotFound{}
}

/*
GetQualityFormsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetQualityFormsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality forms not found response has a 2xx status code
func (o *GetQualityFormsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality forms not found response has a 3xx status code
func (o *GetQualityFormsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms not found response has a 4xx status code
func (o *GetQualityFormsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality forms not found response has a 5xx status code
func (o *GetQualityFormsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality forms not found response a status code equal to that given
func (o *GetQualityFormsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetQualityFormsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityFormsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityFormsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsRequestTimeout creates a GetQualityFormsRequestTimeout with default headers values
func NewGetQualityFormsRequestTimeout() *GetQualityFormsRequestTimeout {
	return &GetQualityFormsRequestTimeout{}
}

/*
GetQualityFormsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetQualityFormsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality forms request timeout response has a 2xx status code
func (o *GetQualityFormsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality forms request timeout response has a 3xx status code
func (o *GetQualityFormsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms request timeout response has a 4xx status code
func (o *GetQualityFormsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality forms request timeout response has a 5xx status code
func (o *GetQualityFormsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality forms request timeout response a status code equal to that given
func (o *GetQualityFormsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetQualityFormsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualityFormsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualityFormsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsRequestEntityTooLarge creates a GetQualityFormsRequestEntityTooLarge with default headers values
func NewGetQualityFormsRequestEntityTooLarge() *GetQualityFormsRequestEntityTooLarge {
	return &GetQualityFormsRequestEntityTooLarge{}
}

/*
GetQualityFormsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetQualityFormsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality forms request entity too large response has a 2xx status code
func (o *GetQualityFormsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality forms request entity too large response has a 3xx status code
func (o *GetQualityFormsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms request entity too large response has a 4xx status code
func (o *GetQualityFormsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality forms request entity too large response has a 5xx status code
func (o *GetQualityFormsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality forms request entity too large response a status code equal to that given
func (o *GetQualityFormsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetQualityFormsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityFormsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityFormsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsUnsupportedMediaType creates a GetQualityFormsUnsupportedMediaType with default headers values
func NewGetQualityFormsUnsupportedMediaType() *GetQualityFormsUnsupportedMediaType {
	return &GetQualityFormsUnsupportedMediaType{}
}

/*
GetQualityFormsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetQualityFormsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality forms unsupported media type response has a 2xx status code
func (o *GetQualityFormsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality forms unsupported media type response has a 3xx status code
func (o *GetQualityFormsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms unsupported media type response has a 4xx status code
func (o *GetQualityFormsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality forms unsupported media type response has a 5xx status code
func (o *GetQualityFormsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality forms unsupported media type response a status code equal to that given
func (o *GetQualityFormsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetQualityFormsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityFormsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityFormsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsTooManyRequests creates a GetQualityFormsTooManyRequests with default headers values
func NewGetQualityFormsTooManyRequests() *GetQualityFormsTooManyRequests {
	return &GetQualityFormsTooManyRequests{}
}

/*
GetQualityFormsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetQualityFormsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality forms too many requests response has a 2xx status code
func (o *GetQualityFormsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality forms too many requests response has a 3xx status code
func (o *GetQualityFormsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms too many requests response has a 4xx status code
func (o *GetQualityFormsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality forms too many requests response has a 5xx status code
func (o *GetQualityFormsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality forms too many requests response a status code equal to that given
func (o *GetQualityFormsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetQualityFormsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityFormsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityFormsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsInternalServerError creates a GetQualityFormsInternalServerError with default headers values
func NewGetQualityFormsInternalServerError() *GetQualityFormsInternalServerError {
	return &GetQualityFormsInternalServerError{}
}

/*
GetQualityFormsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetQualityFormsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality forms internal server error response has a 2xx status code
func (o *GetQualityFormsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality forms internal server error response has a 3xx status code
func (o *GetQualityFormsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms internal server error response has a 4xx status code
func (o *GetQualityFormsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality forms internal server error response has a 5xx status code
func (o *GetQualityFormsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality forms internal server error response a status code equal to that given
func (o *GetQualityFormsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetQualityFormsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityFormsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityFormsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsServiceUnavailable creates a GetQualityFormsServiceUnavailable with default headers values
func NewGetQualityFormsServiceUnavailable() *GetQualityFormsServiceUnavailable {
	return &GetQualityFormsServiceUnavailable{}
}

/*
GetQualityFormsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetQualityFormsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality forms service unavailable response has a 2xx status code
func (o *GetQualityFormsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality forms service unavailable response has a 3xx status code
func (o *GetQualityFormsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms service unavailable response has a 4xx status code
func (o *GetQualityFormsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality forms service unavailable response has a 5xx status code
func (o *GetQualityFormsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality forms service unavailable response a status code equal to that given
func (o *GetQualityFormsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetQualityFormsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityFormsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityFormsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsGatewayTimeout creates a GetQualityFormsGatewayTimeout with default headers values
func NewGetQualityFormsGatewayTimeout() *GetQualityFormsGatewayTimeout {
	return &GetQualityFormsGatewayTimeout{}
}

/*
GetQualityFormsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetQualityFormsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality forms gateway timeout response has a 2xx status code
func (o *GetQualityFormsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality forms gateway timeout response has a 3xx status code
func (o *GetQualityFormsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms gateway timeout response has a 4xx status code
func (o *GetQualityFormsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality forms gateway timeout response has a 5xx status code
func (o *GetQualityFormsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality forms gateway timeout response a status code equal to that given
func (o *GetQualityFormsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetQualityFormsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityFormsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms][%d] getQualityFormsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityFormsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
