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

// GetQualityEvaluatorsActivityReader is a Reader for the GetQualityEvaluatorsActivity structure.
type GetQualityEvaluatorsActivityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQualityEvaluatorsActivityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQualityEvaluatorsActivityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetQualityEvaluatorsActivityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetQualityEvaluatorsActivityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetQualityEvaluatorsActivityForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetQualityEvaluatorsActivityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetQualityEvaluatorsActivityRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetQualityEvaluatorsActivityRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetQualityEvaluatorsActivityUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetQualityEvaluatorsActivityTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetQualityEvaluatorsActivityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetQualityEvaluatorsActivityServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetQualityEvaluatorsActivityGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetQualityEvaluatorsActivityOK creates a GetQualityEvaluatorsActivityOK with default headers values
func NewGetQualityEvaluatorsActivityOK() *GetQualityEvaluatorsActivityOK {
	return &GetQualityEvaluatorsActivityOK{}
}

/*
GetQualityEvaluatorsActivityOK describes a response with status code 200, with default header values.

successful operation
*/
type GetQualityEvaluatorsActivityOK struct {
	Payload *models.EvaluatorActivityEntityListing
}

// IsSuccess returns true when this get quality evaluators activity o k response has a 2xx status code
func (o *GetQualityEvaluatorsActivityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get quality evaluators activity o k response has a 3xx status code
func (o *GetQualityEvaluatorsActivityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluators activity o k response has a 4xx status code
func (o *GetQualityEvaluatorsActivityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality evaluators activity o k response has a 5xx status code
func (o *GetQualityEvaluatorsActivityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality evaluators activity o k response a status code equal to that given
func (o *GetQualityEvaluatorsActivityOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetQualityEvaluatorsActivityOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityOK  %+v", 200, o.Payload)
}

func (o *GetQualityEvaluatorsActivityOK) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityOK  %+v", 200, o.Payload)
}

func (o *GetQualityEvaluatorsActivityOK) GetPayload() *models.EvaluatorActivityEntityListing {
	return o.Payload
}

func (o *GetQualityEvaluatorsActivityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EvaluatorActivityEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityEvaluatorsActivityBadRequest creates a GetQualityEvaluatorsActivityBadRequest with default headers values
func NewGetQualityEvaluatorsActivityBadRequest() *GetQualityEvaluatorsActivityBadRequest {
	return &GetQualityEvaluatorsActivityBadRequest{}
}

/*
GetQualityEvaluatorsActivityBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetQualityEvaluatorsActivityBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality evaluators activity bad request response has a 2xx status code
func (o *GetQualityEvaluatorsActivityBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality evaluators activity bad request response has a 3xx status code
func (o *GetQualityEvaluatorsActivityBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluators activity bad request response has a 4xx status code
func (o *GetQualityEvaluatorsActivityBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality evaluators activity bad request response has a 5xx status code
func (o *GetQualityEvaluatorsActivityBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality evaluators activity bad request response a status code equal to that given
func (o *GetQualityEvaluatorsActivityBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetQualityEvaluatorsActivityBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityEvaluatorsActivityBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityEvaluatorsActivityBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityEvaluatorsActivityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityEvaluatorsActivityUnauthorized creates a GetQualityEvaluatorsActivityUnauthorized with default headers values
func NewGetQualityEvaluatorsActivityUnauthorized() *GetQualityEvaluatorsActivityUnauthorized {
	return &GetQualityEvaluatorsActivityUnauthorized{}
}

/*
GetQualityEvaluatorsActivityUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetQualityEvaluatorsActivityUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality evaluators activity unauthorized response has a 2xx status code
func (o *GetQualityEvaluatorsActivityUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality evaluators activity unauthorized response has a 3xx status code
func (o *GetQualityEvaluatorsActivityUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluators activity unauthorized response has a 4xx status code
func (o *GetQualityEvaluatorsActivityUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality evaluators activity unauthorized response has a 5xx status code
func (o *GetQualityEvaluatorsActivityUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality evaluators activity unauthorized response a status code equal to that given
func (o *GetQualityEvaluatorsActivityUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetQualityEvaluatorsActivityUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityEvaluatorsActivityUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityEvaluatorsActivityUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityEvaluatorsActivityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityEvaluatorsActivityForbidden creates a GetQualityEvaluatorsActivityForbidden with default headers values
func NewGetQualityEvaluatorsActivityForbidden() *GetQualityEvaluatorsActivityForbidden {
	return &GetQualityEvaluatorsActivityForbidden{}
}

/*
GetQualityEvaluatorsActivityForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetQualityEvaluatorsActivityForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality evaluators activity forbidden response has a 2xx status code
func (o *GetQualityEvaluatorsActivityForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality evaluators activity forbidden response has a 3xx status code
func (o *GetQualityEvaluatorsActivityForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluators activity forbidden response has a 4xx status code
func (o *GetQualityEvaluatorsActivityForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality evaluators activity forbidden response has a 5xx status code
func (o *GetQualityEvaluatorsActivityForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality evaluators activity forbidden response a status code equal to that given
func (o *GetQualityEvaluatorsActivityForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetQualityEvaluatorsActivityForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityEvaluatorsActivityForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityEvaluatorsActivityForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityEvaluatorsActivityForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityEvaluatorsActivityNotFound creates a GetQualityEvaluatorsActivityNotFound with default headers values
func NewGetQualityEvaluatorsActivityNotFound() *GetQualityEvaluatorsActivityNotFound {
	return &GetQualityEvaluatorsActivityNotFound{}
}

/*
GetQualityEvaluatorsActivityNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetQualityEvaluatorsActivityNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality evaluators activity not found response has a 2xx status code
func (o *GetQualityEvaluatorsActivityNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality evaluators activity not found response has a 3xx status code
func (o *GetQualityEvaluatorsActivityNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluators activity not found response has a 4xx status code
func (o *GetQualityEvaluatorsActivityNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality evaluators activity not found response has a 5xx status code
func (o *GetQualityEvaluatorsActivityNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality evaluators activity not found response a status code equal to that given
func (o *GetQualityEvaluatorsActivityNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetQualityEvaluatorsActivityNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityEvaluatorsActivityNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityEvaluatorsActivityNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityEvaluatorsActivityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityEvaluatorsActivityRequestTimeout creates a GetQualityEvaluatorsActivityRequestTimeout with default headers values
func NewGetQualityEvaluatorsActivityRequestTimeout() *GetQualityEvaluatorsActivityRequestTimeout {
	return &GetQualityEvaluatorsActivityRequestTimeout{}
}

/*
GetQualityEvaluatorsActivityRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetQualityEvaluatorsActivityRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality evaluators activity request timeout response has a 2xx status code
func (o *GetQualityEvaluatorsActivityRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality evaluators activity request timeout response has a 3xx status code
func (o *GetQualityEvaluatorsActivityRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluators activity request timeout response has a 4xx status code
func (o *GetQualityEvaluatorsActivityRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality evaluators activity request timeout response has a 5xx status code
func (o *GetQualityEvaluatorsActivityRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality evaluators activity request timeout response a status code equal to that given
func (o *GetQualityEvaluatorsActivityRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetQualityEvaluatorsActivityRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualityEvaluatorsActivityRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualityEvaluatorsActivityRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityEvaluatorsActivityRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityEvaluatorsActivityRequestEntityTooLarge creates a GetQualityEvaluatorsActivityRequestEntityTooLarge with default headers values
func NewGetQualityEvaluatorsActivityRequestEntityTooLarge() *GetQualityEvaluatorsActivityRequestEntityTooLarge {
	return &GetQualityEvaluatorsActivityRequestEntityTooLarge{}
}

/*
GetQualityEvaluatorsActivityRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetQualityEvaluatorsActivityRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality evaluators activity request entity too large response has a 2xx status code
func (o *GetQualityEvaluatorsActivityRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality evaluators activity request entity too large response has a 3xx status code
func (o *GetQualityEvaluatorsActivityRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluators activity request entity too large response has a 4xx status code
func (o *GetQualityEvaluatorsActivityRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality evaluators activity request entity too large response has a 5xx status code
func (o *GetQualityEvaluatorsActivityRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality evaluators activity request entity too large response a status code equal to that given
func (o *GetQualityEvaluatorsActivityRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetQualityEvaluatorsActivityRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityEvaluatorsActivityRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityEvaluatorsActivityRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityEvaluatorsActivityRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityEvaluatorsActivityUnsupportedMediaType creates a GetQualityEvaluatorsActivityUnsupportedMediaType with default headers values
func NewGetQualityEvaluatorsActivityUnsupportedMediaType() *GetQualityEvaluatorsActivityUnsupportedMediaType {
	return &GetQualityEvaluatorsActivityUnsupportedMediaType{}
}

/*
GetQualityEvaluatorsActivityUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetQualityEvaluatorsActivityUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality evaluators activity unsupported media type response has a 2xx status code
func (o *GetQualityEvaluatorsActivityUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality evaluators activity unsupported media type response has a 3xx status code
func (o *GetQualityEvaluatorsActivityUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluators activity unsupported media type response has a 4xx status code
func (o *GetQualityEvaluatorsActivityUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality evaluators activity unsupported media type response has a 5xx status code
func (o *GetQualityEvaluatorsActivityUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality evaluators activity unsupported media type response a status code equal to that given
func (o *GetQualityEvaluatorsActivityUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetQualityEvaluatorsActivityUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityEvaluatorsActivityUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityEvaluatorsActivityUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityEvaluatorsActivityUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityEvaluatorsActivityTooManyRequests creates a GetQualityEvaluatorsActivityTooManyRequests with default headers values
func NewGetQualityEvaluatorsActivityTooManyRequests() *GetQualityEvaluatorsActivityTooManyRequests {
	return &GetQualityEvaluatorsActivityTooManyRequests{}
}

/*
GetQualityEvaluatorsActivityTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetQualityEvaluatorsActivityTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality evaluators activity too many requests response has a 2xx status code
func (o *GetQualityEvaluatorsActivityTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality evaluators activity too many requests response has a 3xx status code
func (o *GetQualityEvaluatorsActivityTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluators activity too many requests response has a 4xx status code
func (o *GetQualityEvaluatorsActivityTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality evaluators activity too many requests response has a 5xx status code
func (o *GetQualityEvaluatorsActivityTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality evaluators activity too many requests response a status code equal to that given
func (o *GetQualityEvaluatorsActivityTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetQualityEvaluatorsActivityTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityEvaluatorsActivityTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityEvaluatorsActivityTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityEvaluatorsActivityTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityEvaluatorsActivityInternalServerError creates a GetQualityEvaluatorsActivityInternalServerError with default headers values
func NewGetQualityEvaluatorsActivityInternalServerError() *GetQualityEvaluatorsActivityInternalServerError {
	return &GetQualityEvaluatorsActivityInternalServerError{}
}

/*
GetQualityEvaluatorsActivityInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetQualityEvaluatorsActivityInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality evaluators activity internal server error response has a 2xx status code
func (o *GetQualityEvaluatorsActivityInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality evaluators activity internal server error response has a 3xx status code
func (o *GetQualityEvaluatorsActivityInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluators activity internal server error response has a 4xx status code
func (o *GetQualityEvaluatorsActivityInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality evaluators activity internal server error response has a 5xx status code
func (o *GetQualityEvaluatorsActivityInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality evaluators activity internal server error response a status code equal to that given
func (o *GetQualityEvaluatorsActivityInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetQualityEvaluatorsActivityInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityEvaluatorsActivityInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityEvaluatorsActivityInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityEvaluatorsActivityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityEvaluatorsActivityServiceUnavailable creates a GetQualityEvaluatorsActivityServiceUnavailable with default headers values
func NewGetQualityEvaluatorsActivityServiceUnavailable() *GetQualityEvaluatorsActivityServiceUnavailable {
	return &GetQualityEvaluatorsActivityServiceUnavailable{}
}

/*
GetQualityEvaluatorsActivityServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetQualityEvaluatorsActivityServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality evaluators activity service unavailable response has a 2xx status code
func (o *GetQualityEvaluatorsActivityServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality evaluators activity service unavailable response has a 3xx status code
func (o *GetQualityEvaluatorsActivityServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluators activity service unavailable response has a 4xx status code
func (o *GetQualityEvaluatorsActivityServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality evaluators activity service unavailable response has a 5xx status code
func (o *GetQualityEvaluatorsActivityServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality evaluators activity service unavailable response a status code equal to that given
func (o *GetQualityEvaluatorsActivityServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetQualityEvaluatorsActivityServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityEvaluatorsActivityServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityEvaluatorsActivityServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityEvaluatorsActivityServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityEvaluatorsActivityGatewayTimeout creates a GetQualityEvaluatorsActivityGatewayTimeout with default headers values
func NewGetQualityEvaluatorsActivityGatewayTimeout() *GetQualityEvaluatorsActivityGatewayTimeout {
	return &GetQualityEvaluatorsActivityGatewayTimeout{}
}

/*
GetQualityEvaluatorsActivityGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetQualityEvaluatorsActivityGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality evaluators activity gateway timeout response has a 2xx status code
func (o *GetQualityEvaluatorsActivityGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality evaluators activity gateway timeout response has a 3xx status code
func (o *GetQualityEvaluatorsActivityGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluators activity gateway timeout response has a 4xx status code
func (o *GetQualityEvaluatorsActivityGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality evaluators activity gateway timeout response has a 5xx status code
func (o *GetQualityEvaluatorsActivityGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality evaluators activity gateway timeout response a status code equal to that given
func (o *GetQualityEvaluatorsActivityGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetQualityEvaluatorsActivityGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityEvaluatorsActivityGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluators/activity][%d] getQualityEvaluatorsActivityGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityEvaluatorsActivityGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityEvaluatorsActivityGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
