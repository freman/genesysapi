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

// GetQualityEvaluationsQueryReader is a Reader for the GetQualityEvaluationsQuery structure.
type GetQualityEvaluationsQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQualityEvaluationsQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQualityEvaluationsQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetQualityEvaluationsQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetQualityEvaluationsQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetQualityEvaluationsQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetQualityEvaluationsQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetQualityEvaluationsQueryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetQualityEvaluationsQueryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetQualityEvaluationsQueryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetQualityEvaluationsQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetQualityEvaluationsQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetQualityEvaluationsQueryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetQualityEvaluationsQueryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetQualityEvaluationsQueryOK creates a GetQualityEvaluationsQueryOK with default headers values
func NewGetQualityEvaluationsQueryOK() *GetQualityEvaluationsQueryOK {
	return &GetQualityEvaluationsQueryOK{}
}

/*
GetQualityEvaluationsQueryOK describes a response with status code 200, with default header values.

successful operation
*/
type GetQualityEvaluationsQueryOK struct {
	Payload *models.EvaluationEntityListing
}

// IsSuccess returns true when this get quality evaluations query o k response has a 2xx status code
func (o *GetQualityEvaluationsQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get quality evaluations query o k response has a 3xx status code
func (o *GetQualityEvaluationsQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluations query o k response has a 4xx status code
func (o *GetQualityEvaluationsQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality evaluations query o k response has a 5xx status code
func (o *GetQualityEvaluationsQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality evaluations query o k response a status code equal to that given
func (o *GetQualityEvaluationsQueryOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetQualityEvaluationsQueryOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryOK  %+v", 200, o.Payload)
}

func (o *GetQualityEvaluationsQueryOK) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryOK  %+v", 200, o.Payload)
}

func (o *GetQualityEvaluationsQueryOK) GetPayload() *models.EvaluationEntityListing {
	return o.Payload
}

func (o *GetQualityEvaluationsQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EvaluationEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityEvaluationsQueryBadRequest creates a GetQualityEvaluationsQueryBadRequest with default headers values
func NewGetQualityEvaluationsQueryBadRequest() *GetQualityEvaluationsQueryBadRequest {
	return &GetQualityEvaluationsQueryBadRequest{}
}

/*
GetQualityEvaluationsQueryBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetQualityEvaluationsQueryBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality evaluations query bad request response has a 2xx status code
func (o *GetQualityEvaluationsQueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality evaluations query bad request response has a 3xx status code
func (o *GetQualityEvaluationsQueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluations query bad request response has a 4xx status code
func (o *GetQualityEvaluationsQueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality evaluations query bad request response has a 5xx status code
func (o *GetQualityEvaluationsQueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality evaluations query bad request response a status code equal to that given
func (o *GetQualityEvaluationsQueryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetQualityEvaluationsQueryBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityEvaluationsQueryBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityEvaluationsQueryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityEvaluationsQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityEvaluationsQueryUnauthorized creates a GetQualityEvaluationsQueryUnauthorized with default headers values
func NewGetQualityEvaluationsQueryUnauthorized() *GetQualityEvaluationsQueryUnauthorized {
	return &GetQualityEvaluationsQueryUnauthorized{}
}

/*
GetQualityEvaluationsQueryUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetQualityEvaluationsQueryUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality evaluations query unauthorized response has a 2xx status code
func (o *GetQualityEvaluationsQueryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality evaluations query unauthorized response has a 3xx status code
func (o *GetQualityEvaluationsQueryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluations query unauthorized response has a 4xx status code
func (o *GetQualityEvaluationsQueryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality evaluations query unauthorized response has a 5xx status code
func (o *GetQualityEvaluationsQueryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality evaluations query unauthorized response a status code equal to that given
func (o *GetQualityEvaluationsQueryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetQualityEvaluationsQueryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityEvaluationsQueryUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityEvaluationsQueryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityEvaluationsQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityEvaluationsQueryForbidden creates a GetQualityEvaluationsQueryForbidden with default headers values
func NewGetQualityEvaluationsQueryForbidden() *GetQualityEvaluationsQueryForbidden {
	return &GetQualityEvaluationsQueryForbidden{}
}

/*
GetQualityEvaluationsQueryForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetQualityEvaluationsQueryForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality evaluations query forbidden response has a 2xx status code
func (o *GetQualityEvaluationsQueryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality evaluations query forbidden response has a 3xx status code
func (o *GetQualityEvaluationsQueryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluations query forbidden response has a 4xx status code
func (o *GetQualityEvaluationsQueryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality evaluations query forbidden response has a 5xx status code
func (o *GetQualityEvaluationsQueryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality evaluations query forbidden response a status code equal to that given
func (o *GetQualityEvaluationsQueryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetQualityEvaluationsQueryForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityEvaluationsQueryForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityEvaluationsQueryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityEvaluationsQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityEvaluationsQueryNotFound creates a GetQualityEvaluationsQueryNotFound with default headers values
func NewGetQualityEvaluationsQueryNotFound() *GetQualityEvaluationsQueryNotFound {
	return &GetQualityEvaluationsQueryNotFound{}
}

/*
GetQualityEvaluationsQueryNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetQualityEvaluationsQueryNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality evaluations query not found response has a 2xx status code
func (o *GetQualityEvaluationsQueryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality evaluations query not found response has a 3xx status code
func (o *GetQualityEvaluationsQueryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluations query not found response has a 4xx status code
func (o *GetQualityEvaluationsQueryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality evaluations query not found response has a 5xx status code
func (o *GetQualityEvaluationsQueryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality evaluations query not found response a status code equal to that given
func (o *GetQualityEvaluationsQueryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetQualityEvaluationsQueryNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityEvaluationsQueryNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityEvaluationsQueryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityEvaluationsQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityEvaluationsQueryRequestTimeout creates a GetQualityEvaluationsQueryRequestTimeout with default headers values
func NewGetQualityEvaluationsQueryRequestTimeout() *GetQualityEvaluationsQueryRequestTimeout {
	return &GetQualityEvaluationsQueryRequestTimeout{}
}

/*
GetQualityEvaluationsQueryRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetQualityEvaluationsQueryRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality evaluations query request timeout response has a 2xx status code
func (o *GetQualityEvaluationsQueryRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality evaluations query request timeout response has a 3xx status code
func (o *GetQualityEvaluationsQueryRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluations query request timeout response has a 4xx status code
func (o *GetQualityEvaluationsQueryRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality evaluations query request timeout response has a 5xx status code
func (o *GetQualityEvaluationsQueryRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality evaluations query request timeout response a status code equal to that given
func (o *GetQualityEvaluationsQueryRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetQualityEvaluationsQueryRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualityEvaluationsQueryRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualityEvaluationsQueryRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityEvaluationsQueryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityEvaluationsQueryRequestEntityTooLarge creates a GetQualityEvaluationsQueryRequestEntityTooLarge with default headers values
func NewGetQualityEvaluationsQueryRequestEntityTooLarge() *GetQualityEvaluationsQueryRequestEntityTooLarge {
	return &GetQualityEvaluationsQueryRequestEntityTooLarge{}
}

/*
GetQualityEvaluationsQueryRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetQualityEvaluationsQueryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality evaluations query request entity too large response has a 2xx status code
func (o *GetQualityEvaluationsQueryRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality evaluations query request entity too large response has a 3xx status code
func (o *GetQualityEvaluationsQueryRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluations query request entity too large response has a 4xx status code
func (o *GetQualityEvaluationsQueryRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality evaluations query request entity too large response has a 5xx status code
func (o *GetQualityEvaluationsQueryRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality evaluations query request entity too large response a status code equal to that given
func (o *GetQualityEvaluationsQueryRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetQualityEvaluationsQueryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityEvaluationsQueryRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityEvaluationsQueryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityEvaluationsQueryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityEvaluationsQueryUnsupportedMediaType creates a GetQualityEvaluationsQueryUnsupportedMediaType with default headers values
func NewGetQualityEvaluationsQueryUnsupportedMediaType() *GetQualityEvaluationsQueryUnsupportedMediaType {
	return &GetQualityEvaluationsQueryUnsupportedMediaType{}
}

/*
GetQualityEvaluationsQueryUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetQualityEvaluationsQueryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality evaluations query unsupported media type response has a 2xx status code
func (o *GetQualityEvaluationsQueryUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality evaluations query unsupported media type response has a 3xx status code
func (o *GetQualityEvaluationsQueryUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluations query unsupported media type response has a 4xx status code
func (o *GetQualityEvaluationsQueryUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality evaluations query unsupported media type response has a 5xx status code
func (o *GetQualityEvaluationsQueryUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality evaluations query unsupported media type response a status code equal to that given
func (o *GetQualityEvaluationsQueryUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetQualityEvaluationsQueryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityEvaluationsQueryUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityEvaluationsQueryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityEvaluationsQueryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityEvaluationsQueryTooManyRequests creates a GetQualityEvaluationsQueryTooManyRequests with default headers values
func NewGetQualityEvaluationsQueryTooManyRequests() *GetQualityEvaluationsQueryTooManyRequests {
	return &GetQualityEvaluationsQueryTooManyRequests{}
}

/*
GetQualityEvaluationsQueryTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetQualityEvaluationsQueryTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality evaluations query too many requests response has a 2xx status code
func (o *GetQualityEvaluationsQueryTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality evaluations query too many requests response has a 3xx status code
func (o *GetQualityEvaluationsQueryTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluations query too many requests response has a 4xx status code
func (o *GetQualityEvaluationsQueryTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality evaluations query too many requests response has a 5xx status code
func (o *GetQualityEvaluationsQueryTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality evaluations query too many requests response a status code equal to that given
func (o *GetQualityEvaluationsQueryTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetQualityEvaluationsQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityEvaluationsQueryTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityEvaluationsQueryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityEvaluationsQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityEvaluationsQueryInternalServerError creates a GetQualityEvaluationsQueryInternalServerError with default headers values
func NewGetQualityEvaluationsQueryInternalServerError() *GetQualityEvaluationsQueryInternalServerError {
	return &GetQualityEvaluationsQueryInternalServerError{}
}

/*
GetQualityEvaluationsQueryInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetQualityEvaluationsQueryInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality evaluations query internal server error response has a 2xx status code
func (o *GetQualityEvaluationsQueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality evaluations query internal server error response has a 3xx status code
func (o *GetQualityEvaluationsQueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluations query internal server error response has a 4xx status code
func (o *GetQualityEvaluationsQueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality evaluations query internal server error response has a 5xx status code
func (o *GetQualityEvaluationsQueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality evaluations query internal server error response a status code equal to that given
func (o *GetQualityEvaluationsQueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetQualityEvaluationsQueryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityEvaluationsQueryInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityEvaluationsQueryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityEvaluationsQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityEvaluationsQueryServiceUnavailable creates a GetQualityEvaluationsQueryServiceUnavailable with default headers values
func NewGetQualityEvaluationsQueryServiceUnavailable() *GetQualityEvaluationsQueryServiceUnavailable {
	return &GetQualityEvaluationsQueryServiceUnavailable{}
}

/*
GetQualityEvaluationsQueryServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetQualityEvaluationsQueryServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality evaluations query service unavailable response has a 2xx status code
func (o *GetQualityEvaluationsQueryServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality evaluations query service unavailable response has a 3xx status code
func (o *GetQualityEvaluationsQueryServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluations query service unavailable response has a 4xx status code
func (o *GetQualityEvaluationsQueryServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality evaluations query service unavailable response has a 5xx status code
func (o *GetQualityEvaluationsQueryServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality evaluations query service unavailable response a status code equal to that given
func (o *GetQualityEvaluationsQueryServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetQualityEvaluationsQueryServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityEvaluationsQueryServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityEvaluationsQueryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityEvaluationsQueryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityEvaluationsQueryGatewayTimeout creates a GetQualityEvaluationsQueryGatewayTimeout with default headers values
func NewGetQualityEvaluationsQueryGatewayTimeout() *GetQualityEvaluationsQueryGatewayTimeout {
	return &GetQualityEvaluationsQueryGatewayTimeout{}
}

/*
GetQualityEvaluationsQueryGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetQualityEvaluationsQueryGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality evaluations query gateway timeout response has a 2xx status code
func (o *GetQualityEvaluationsQueryGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality evaluations query gateway timeout response has a 3xx status code
func (o *GetQualityEvaluationsQueryGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality evaluations query gateway timeout response has a 4xx status code
func (o *GetQualityEvaluationsQueryGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality evaluations query gateway timeout response has a 5xx status code
func (o *GetQualityEvaluationsQueryGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality evaluations query gateway timeout response a status code equal to that given
func (o *GetQualityEvaluationsQueryGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetQualityEvaluationsQueryGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityEvaluationsQueryGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/evaluations/query][%d] getQualityEvaluationsQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityEvaluationsQueryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityEvaluationsQueryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
