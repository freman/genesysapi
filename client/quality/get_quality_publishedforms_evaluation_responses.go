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

// GetQualityPublishedformsEvaluationReader is a Reader for the GetQualityPublishedformsEvaluation structure.
type GetQualityPublishedformsEvaluationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQualityPublishedformsEvaluationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQualityPublishedformsEvaluationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetQualityPublishedformsEvaluationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetQualityPublishedformsEvaluationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetQualityPublishedformsEvaluationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetQualityPublishedformsEvaluationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetQualityPublishedformsEvaluationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetQualityPublishedformsEvaluationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetQualityPublishedformsEvaluationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetQualityPublishedformsEvaluationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetQualityPublishedformsEvaluationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetQualityPublishedformsEvaluationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetQualityPublishedformsEvaluationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetQualityPublishedformsEvaluationOK creates a GetQualityPublishedformsEvaluationOK with default headers values
func NewGetQualityPublishedformsEvaluationOK() *GetQualityPublishedformsEvaluationOK {
	return &GetQualityPublishedformsEvaluationOK{}
}

/*
GetQualityPublishedformsEvaluationOK describes a response with status code 200, with default header values.

successful operation
*/
type GetQualityPublishedformsEvaluationOK struct {
	Payload *models.EvaluationForm
}

// IsSuccess returns true when this get quality publishedforms evaluation o k response has a 2xx status code
func (o *GetQualityPublishedformsEvaluationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get quality publishedforms evaluation o k response has a 3xx status code
func (o *GetQualityPublishedformsEvaluationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms evaluation o k response has a 4xx status code
func (o *GetQualityPublishedformsEvaluationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality publishedforms evaluation o k response has a 5xx status code
func (o *GetQualityPublishedformsEvaluationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms evaluation o k response a status code equal to that given
func (o *GetQualityPublishedformsEvaluationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetQualityPublishedformsEvaluationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationOK  %+v", 200, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationOK) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationOK  %+v", 200, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationOK) GetPayload() *models.EvaluationForm {
	return o.Payload
}

func (o *GetQualityPublishedformsEvaluationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EvaluationForm)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsEvaluationBadRequest creates a GetQualityPublishedformsEvaluationBadRequest with default headers values
func NewGetQualityPublishedformsEvaluationBadRequest() *GetQualityPublishedformsEvaluationBadRequest {
	return &GetQualityPublishedformsEvaluationBadRequest{}
}

/*
GetQualityPublishedformsEvaluationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetQualityPublishedformsEvaluationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms evaluation bad request response has a 2xx status code
func (o *GetQualityPublishedformsEvaluationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms evaluation bad request response has a 3xx status code
func (o *GetQualityPublishedformsEvaluationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms evaluation bad request response has a 4xx status code
func (o *GetQualityPublishedformsEvaluationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms evaluation bad request response has a 5xx status code
func (o *GetQualityPublishedformsEvaluationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms evaluation bad request response a status code equal to that given
func (o *GetQualityPublishedformsEvaluationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetQualityPublishedformsEvaluationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsEvaluationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsEvaluationUnauthorized creates a GetQualityPublishedformsEvaluationUnauthorized with default headers values
func NewGetQualityPublishedformsEvaluationUnauthorized() *GetQualityPublishedformsEvaluationUnauthorized {
	return &GetQualityPublishedformsEvaluationUnauthorized{}
}

/*
GetQualityPublishedformsEvaluationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetQualityPublishedformsEvaluationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms evaluation unauthorized response has a 2xx status code
func (o *GetQualityPublishedformsEvaluationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms evaluation unauthorized response has a 3xx status code
func (o *GetQualityPublishedformsEvaluationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms evaluation unauthorized response has a 4xx status code
func (o *GetQualityPublishedformsEvaluationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms evaluation unauthorized response has a 5xx status code
func (o *GetQualityPublishedformsEvaluationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms evaluation unauthorized response a status code equal to that given
func (o *GetQualityPublishedformsEvaluationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetQualityPublishedformsEvaluationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsEvaluationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsEvaluationForbidden creates a GetQualityPublishedformsEvaluationForbidden with default headers values
func NewGetQualityPublishedformsEvaluationForbidden() *GetQualityPublishedformsEvaluationForbidden {
	return &GetQualityPublishedformsEvaluationForbidden{}
}

/*
GetQualityPublishedformsEvaluationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetQualityPublishedformsEvaluationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms evaluation forbidden response has a 2xx status code
func (o *GetQualityPublishedformsEvaluationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms evaluation forbidden response has a 3xx status code
func (o *GetQualityPublishedformsEvaluationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms evaluation forbidden response has a 4xx status code
func (o *GetQualityPublishedformsEvaluationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms evaluation forbidden response has a 5xx status code
func (o *GetQualityPublishedformsEvaluationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms evaluation forbidden response a status code equal to that given
func (o *GetQualityPublishedformsEvaluationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetQualityPublishedformsEvaluationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsEvaluationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsEvaluationNotFound creates a GetQualityPublishedformsEvaluationNotFound with default headers values
func NewGetQualityPublishedformsEvaluationNotFound() *GetQualityPublishedformsEvaluationNotFound {
	return &GetQualityPublishedformsEvaluationNotFound{}
}

/*
GetQualityPublishedformsEvaluationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetQualityPublishedformsEvaluationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms evaluation not found response has a 2xx status code
func (o *GetQualityPublishedformsEvaluationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms evaluation not found response has a 3xx status code
func (o *GetQualityPublishedformsEvaluationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms evaluation not found response has a 4xx status code
func (o *GetQualityPublishedformsEvaluationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms evaluation not found response has a 5xx status code
func (o *GetQualityPublishedformsEvaluationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms evaluation not found response a status code equal to that given
func (o *GetQualityPublishedformsEvaluationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetQualityPublishedformsEvaluationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsEvaluationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsEvaluationRequestTimeout creates a GetQualityPublishedformsEvaluationRequestTimeout with default headers values
func NewGetQualityPublishedformsEvaluationRequestTimeout() *GetQualityPublishedformsEvaluationRequestTimeout {
	return &GetQualityPublishedformsEvaluationRequestTimeout{}
}

/*
GetQualityPublishedformsEvaluationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetQualityPublishedformsEvaluationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms evaluation request timeout response has a 2xx status code
func (o *GetQualityPublishedformsEvaluationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms evaluation request timeout response has a 3xx status code
func (o *GetQualityPublishedformsEvaluationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms evaluation request timeout response has a 4xx status code
func (o *GetQualityPublishedformsEvaluationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms evaluation request timeout response has a 5xx status code
func (o *GetQualityPublishedformsEvaluationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms evaluation request timeout response a status code equal to that given
func (o *GetQualityPublishedformsEvaluationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetQualityPublishedformsEvaluationRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsEvaluationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsEvaluationRequestEntityTooLarge creates a GetQualityPublishedformsEvaluationRequestEntityTooLarge with default headers values
func NewGetQualityPublishedformsEvaluationRequestEntityTooLarge() *GetQualityPublishedformsEvaluationRequestEntityTooLarge {
	return &GetQualityPublishedformsEvaluationRequestEntityTooLarge{}
}

/*
GetQualityPublishedformsEvaluationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetQualityPublishedformsEvaluationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms evaluation request entity too large response has a 2xx status code
func (o *GetQualityPublishedformsEvaluationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms evaluation request entity too large response has a 3xx status code
func (o *GetQualityPublishedformsEvaluationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms evaluation request entity too large response has a 4xx status code
func (o *GetQualityPublishedformsEvaluationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms evaluation request entity too large response has a 5xx status code
func (o *GetQualityPublishedformsEvaluationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms evaluation request entity too large response a status code equal to that given
func (o *GetQualityPublishedformsEvaluationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetQualityPublishedformsEvaluationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsEvaluationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsEvaluationUnsupportedMediaType creates a GetQualityPublishedformsEvaluationUnsupportedMediaType with default headers values
func NewGetQualityPublishedformsEvaluationUnsupportedMediaType() *GetQualityPublishedformsEvaluationUnsupportedMediaType {
	return &GetQualityPublishedformsEvaluationUnsupportedMediaType{}
}

/*
GetQualityPublishedformsEvaluationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetQualityPublishedformsEvaluationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms evaluation unsupported media type response has a 2xx status code
func (o *GetQualityPublishedformsEvaluationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms evaluation unsupported media type response has a 3xx status code
func (o *GetQualityPublishedformsEvaluationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms evaluation unsupported media type response has a 4xx status code
func (o *GetQualityPublishedformsEvaluationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms evaluation unsupported media type response has a 5xx status code
func (o *GetQualityPublishedformsEvaluationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms evaluation unsupported media type response a status code equal to that given
func (o *GetQualityPublishedformsEvaluationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetQualityPublishedformsEvaluationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsEvaluationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsEvaluationTooManyRequests creates a GetQualityPublishedformsEvaluationTooManyRequests with default headers values
func NewGetQualityPublishedformsEvaluationTooManyRequests() *GetQualityPublishedformsEvaluationTooManyRequests {
	return &GetQualityPublishedformsEvaluationTooManyRequests{}
}

/*
GetQualityPublishedformsEvaluationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetQualityPublishedformsEvaluationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms evaluation too many requests response has a 2xx status code
func (o *GetQualityPublishedformsEvaluationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms evaluation too many requests response has a 3xx status code
func (o *GetQualityPublishedformsEvaluationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms evaluation too many requests response has a 4xx status code
func (o *GetQualityPublishedformsEvaluationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality publishedforms evaluation too many requests response has a 5xx status code
func (o *GetQualityPublishedformsEvaluationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality publishedforms evaluation too many requests response a status code equal to that given
func (o *GetQualityPublishedformsEvaluationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetQualityPublishedformsEvaluationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsEvaluationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsEvaluationInternalServerError creates a GetQualityPublishedformsEvaluationInternalServerError with default headers values
func NewGetQualityPublishedformsEvaluationInternalServerError() *GetQualityPublishedformsEvaluationInternalServerError {
	return &GetQualityPublishedformsEvaluationInternalServerError{}
}

/*
GetQualityPublishedformsEvaluationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetQualityPublishedformsEvaluationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms evaluation internal server error response has a 2xx status code
func (o *GetQualityPublishedformsEvaluationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms evaluation internal server error response has a 3xx status code
func (o *GetQualityPublishedformsEvaluationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms evaluation internal server error response has a 4xx status code
func (o *GetQualityPublishedformsEvaluationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality publishedforms evaluation internal server error response has a 5xx status code
func (o *GetQualityPublishedformsEvaluationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality publishedforms evaluation internal server error response a status code equal to that given
func (o *GetQualityPublishedformsEvaluationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetQualityPublishedformsEvaluationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsEvaluationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsEvaluationServiceUnavailable creates a GetQualityPublishedformsEvaluationServiceUnavailable with default headers values
func NewGetQualityPublishedformsEvaluationServiceUnavailable() *GetQualityPublishedformsEvaluationServiceUnavailable {
	return &GetQualityPublishedformsEvaluationServiceUnavailable{}
}

/*
GetQualityPublishedformsEvaluationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetQualityPublishedformsEvaluationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms evaluation service unavailable response has a 2xx status code
func (o *GetQualityPublishedformsEvaluationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms evaluation service unavailable response has a 3xx status code
func (o *GetQualityPublishedformsEvaluationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms evaluation service unavailable response has a 4xx status code
func (o *GetQualityPublishedformsEvaluationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality publishedforms evaluation service unavailable response has a 5xx status code
func (o *GetQualityPublishedformsEvaluationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality publishedforms evaluation service unavailable response a status code equal to that given
func (o *GetQualityPublishedformsEvaluationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetQualityPublishedformsEvaluationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsEvaluationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityPublishedformsEvaluationGatewayTimeout creates a GetQualityPublishedformsEvaluationGatewayTimeout with default headers values
func NewGetQualityPublishedformsEvaluationGatewayTimeout() *GetQualityPublishedformsEvaluationGatewayTimeout {
	return &GetQualityPublishedformsEvaluationGatewayTimeout{}
}

/*
GetQualityPublishedformsEvaluationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetQualityPublishedformsEvaluationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality publishedforms evaluation gateway timeout response has a 2xx status code
func (o *GetQualityPublishedformsEvaluationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality publishedforms evaluation gateway timeout response has a 3xx status code
func (o *GetQualityPublishedformsEvaluationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality publishedforms evaluation gateway timeout response has a 4xx status code
func (o *GetQualityPublishedformsEvaluationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality publishedforms evaluation gateway timeout response has a 5xx status code
func (o *GetQualityPublishedformsEvaluationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality publishedforms evaluation gateway timeout response a status code equal to that given
func (o *GetQualityPublishedformsEvaluationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetQualityPublishedformsEvaluationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/publishedforms/evaluations/{formId}][%d] getQualityPublishedformsEvaluationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityPublishedformsEvaluationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityPublishedformsEvaluationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
