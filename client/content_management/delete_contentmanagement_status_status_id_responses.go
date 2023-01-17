// Code generated by go-swagger; DO NOT EDIT.

package content_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteContentmanagementStatusStatusIDReader is a Reader for the DeleteContentmanagementStatusStatusID structure.
type DeleteContentmanagementStatusStatusIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteContentmanagementStatusStatusIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewDeleteContentmanagementStatusStatusIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteContentmanagementStatusStatusIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteContentmanagementStatusStatusIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteContentmanagementStatusStatusIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteContentmanagementStatusStatusIDRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteContentmanagementStatusStatusIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteContentmanagementStatusStatusIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteContentmanagementStatusStatusIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteContentmanagementStatusStatusIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteContentmanagementStatusStatusIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteContentmanagementStatusStatusIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteContentmanagementStatusStatusIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteContentmanagementStatusStatusIDBadRequest creates a DeleteContentmanagementStatusStatusIDBadRequest with default headers values
func NewDeleteContentmanagementStatusStatusIDBadRequest() *DeleteContentmanagementStatusStatusIDBadRequest {
	return &DeleteContentmanagementStatusStatusIDBadRequest{}
}

/*
DeleteContentmanagementStatusStatusIDBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteContentmanagementStatusStatusIDBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete contentmanagement status status Id bad request response has a 2xx status code
func (o *DeleteContentmanagementStatusStatusIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete contentmanagement status status Id bad request response has a 3xx status code
func (o *DeleteContentmanagementStatusStatusIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete contentmanagement status status Id bad request response has a 4xx status code
func (o *DeleteContentmanagementStatusStatusIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete contentmanagement status status Id bad request response has a 5xx status code
func (o *DeleteContentmanagementStatusStatusIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete contentmanagement status status Id bad request response a status code equal to that given
func (o *DeleteContentmanagementStatusStatusIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteContentmanagementStatusStatusIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteContentmanagementStatusStatusIDBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteContentmanagementStatusStatusIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentmanagementStatusStatusIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentmanagementStatusStatusIDUnauthorized creates a DeleteContentmanagementStatusStatusIDUnauthorized with default headers values
func NewDeleteContentmanagementStatusStatusIDUnauthorized() *DeleteContentmanagementStatusStatusIDUnauthorized {
	return &DeleteContentmanagementStatusStatusIDUnauthorized{}
}

/*
DeleteContentmanagementStatusStatusIDUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteContentmanagementStatusStatusIDUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete contentmanagement status status Id unauthorized response has a 2xx status code
func (o *DeleteContentmanagementStatusStatusIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete contentmanagement status status Id unauthorized response has a 3xx status code
func (o *DeleteContentmanagementStatusStatusIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete contentmanagement status status Id unauthorized response has a 4xx status code
func (o *DeleteContentmanagementStatusStatusIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete contentmanagement status status Id unauthorized response has a 5xx status code
func (o *DeleteContentmanagementStatusStatusIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete contentmanagement status status Id unauthorized response a status code equal to that given
func (o *DeleteContentmanagementStatusStatusIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteContentmanagementStatusStatusIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteContentmanagementStatusStatusIDUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteContentmanagementStatusStatusIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentmanagementStatusStatusIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentmanagementStatusStatusIDForbidden creates a DeleteContentmanagementStatusStatusIDForbidden with default headers values
func NewDeleteContentmanagementStatusStatusIDForbidden() *DeleteContentmanagementStatusStatusIDForbidden {
	return &DeleteContentmanagementStatusStatusIDForbidden{}
}

/*
DeleteContentmanagementStatusStatusIDForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteContentmanagementStatusStatusIDForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete contentmanagement status status Id forbidden response has a 2xx status code
func (o *DeleteContentmanagementStatusStatusIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete contentmanagement status status Id forbidden response has a 3xx status code
func (o *DeleteContentmanagementStatusStatusIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete contentmanagement status status Id forbidden response has a 4xx status code
func (o *DeleteContentmanagementStatusStatusIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete contentmanagement status status Id forbidden response has a 5xx status code
func (o *DeleteContentmanagementStatusStatusIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete contentmanagement status status Id forbidden response a status code equal to that given
func (o *DeleteContentmanagementStatusStatusIDForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteContentmanagementStatusStatusIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteContentmanagementStatusStatusIDForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteContentmanagementStatusStatusIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentmanagementStatusStatusIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentmanagementStatusStatusIDNotFound creates a DeleteContentmanagementStatusStatusIDNotFound with default headers values
func NewDeleteContentmanagementStatusStatusIDNotFound() *DeleteContentmanagementStatusStatusIDNotFound {
	return &DeleteContentmanagementStatusStatusIDNotFound{}
}

/*
DeleteContentmanagementStatusStatusIDNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteContentmanagementStatusStatusIDNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete contentmanagement status status Id not found response has a 2xx status code
func (o *DeleteContentmanagementStatusStatusIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete contentmanagement status status Id not found response has a 3xx status code
func (o *DeleteContentmanagementStatusStatusIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete contentmanagement status status Id not found response has a 4xx status code
func (o *DeleteContentmanagementStatusStatusIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete contentmanagement status status Id not found response has a 5xx status code
func (o *DeleteContentmanagementStatusStatusIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete contentmanagement status status Id not found response a status code equal to that given
func (o *DeleteContentmanagementStatusStatusIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteContentmanagementStatusStatusIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteContentmanagementStatusStatusIDNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteContentmanagementStatusStatusIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentmanagementStatusStatusIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentmanagementStatusStatusIDRequestTimeout creates a DeleteContentmanagementStatusStatusIDRequestTimeout with default headers values
func NewDeleteContentmanagementStatusStatusIDRequestTimeout() *DeleteContentmanagementStatusStatusIDRequestTimeout {
	return &DeleteContentmanagementStatusStatusIDRequestTimeout{}
}

/*
DeleteContentmanagementStatusStatusIDRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteContentmanagementStatusStatusIDRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete contentmanagement status status Id request timeout response has a 2xx status code
func (o *DeleteContentmanagementStatusStatusIDRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete contentmanagement status status Id request timeout response has a 3xx status code
func (o *DeleteContentmanagementStatusStatusIDRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete contentmanagement status status Id request timeout response has a 4xx status code
func (o *DeleteContentmanagementStatusStatusIDRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete contentmanagement status status Id request timeout response has a 5xx status code
func (o *DeleteContentmanagementStatusStatusIDRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete contentmanagement status status Id request timeout response a status code equal to that given
func (o *DeleteContentmanagementStatusStatusIDRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteContentmanagementStatusStatusIDRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusIdRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteContentmanagementStatusStatusIDRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusIdRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteContentmanagementStatusStatusIDRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentmanagementStatusStatusIDRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentmanagementStatusStatusIDRequestEntityTooLarge creates a DeleteContentmanagementStatusStatusIDRequestEntityTooLarge with default headers values
func NewDeleteContentmanagementStatusStatusIDRequestEntityTooLarge() *DeleteContentmanagementStatusStatusIDRequestEntityTooLarge {
	return &DeleteContentmanagementStatusStatusIDRequestEntityTooLarge{}
}

/*
DeleteContentmanagementStatusStatusIDRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteContentmanagementStatusStatusIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete contentmanagement status status Id request entity too large response has a 2xx status code
func (o *DeleteContentmanagementStatusStatusIDRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete contentmanagement status status Id request entity too large response has a 3xx status code
func (o *DeleteContentmanagementStatusStatusIDRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete contentmanagement status status Id request entity too large response has a 4xx status code
func (o *DeleteContentmanagementStatusStatusIDRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete contentmanagement status status Id request entity too large response has a 5xx status code
func (o *DeleteContentmanagementStatusStatusIDRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete contentmanagement status status Id request entity too large response a status code equal to that given
func (o *DeleteContentmanagementStatusStatusIDRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteContentmanagementStatusStatusIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteContentmanagementStatusStatusIDRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteContentmanagementStatusStatusIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentmanagementStatusStatusIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentmanagementStatusStatusIDUnsupportedMediaType creates a DeleteContentmanagementStatusStatusIDUnsupportedMediaType with default headers values
func NewDeleteContentmanagementStatusStatusIDUnsupportedMediaType() *DeleteContentmanagementStatusStatusIDUnsupportedMediaType {
	return &DeleteContentmanagementStatusStatusIDUnsupportedMediaType{}
}

/*
DeleteContentmanagementStatusStatusIDUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteContentmanagementStatusStatusIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete contentmanagement status status Id unsupported media type response has a 2xx status code
func (o *DeleteContentmanagementStatusStatusIDUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete contentmanagement status status Id unsupported media type response has a 3xx status code
func (o *DeleteContentmanagementStatusStatusIDUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete contentmanagement status status Id unsupported media type response has a 4xx status code
func (o *DeleteContentmanagementStatusStatusIDUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete contentmanagement status status Id unsupported media type response has a 5xx status code
func (o *DeleteContentmanagementStatusStatusIDUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete contentmanagement status status Id unsupported media type response a status code equal to that given
func (o *DeleteContentmanagementStatusStatusIDUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteContentmanagementStatusStatusIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteContentmanagementStatusStatusIDUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteContentmanagementStatusStatusIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentmanagementStatusStatusIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentmanagementStatusStatusIDTooManyRequests creates a DeleteContentmanagementStatusStatusIDTooManyRequests with default headers values
func NewDeleteContentmanagementStatusStatusIDTooManyRequests() *DeleteContentmanagementStatusStatusIDTooManyRequests {
	return &DeleteContentmanagementStatusStatusIDTooManyRequests{}
}

/*
DeleteContentmanagementStatusStatusIDTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteContentmanagementStatusStatusIDTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete contentmanagement status status Id too many requests response has a 2xx status code
func (o *DeleteContentmanagementStatusStatusIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete contentmanagement status status Id too many requests response has a 3xx status code
func (o *DeleteContentmanagementStatusStatusIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete contentmanagement status status Id too many requests response has a 4xx status code
func (o *DeleteContentmanagementStatusStatusIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete contentmanagement status status Id too many requests response has a 5xx status code
func (o *DeleteContentmanagementStatusStatusIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete contentmanagement status status Id too many requests response a status code equal to that given
func (o *DeleteContentmanagementStatusStatusIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteContentmanagementStatusStatusIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteContentmanagementStatusStatusIDTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteContentmanagementStatusStatusIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentmanagementStatusStatusIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentmanagementStatusStatusIDInternalServerError creates a DeleteContentmanagementStatusStatusIDInternalServerError with default headers values
func NewDeleteContentmanagementStatusStatusIDInternalServerError() *DeleteContentmanagementStatusStatusIDInternalServerError {
	return &DeleteContentmanagementStatusStatusIDInternalServerError{}
}

/*
DeleteContentmanagementStatusStatusIDInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteContentmanagementStatusStatusIDInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete contentmanagement status status Id internal server error response has a 2xx status code
func (o *DeleteContentmanagementStatusStatusIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete contentmanagement status status Id internal server error response has a 3xx status code
func (o *DeleteContentmanagementStatusStatusIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete contentmanagement status status Id internal server error response has a 4xx status code
func (o *DeleteContentmanagementStatusStatusIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete contentmanagement status status Id internal server error response has a 5xx status code
func (o *DeleteContentmanagementStatusStatusIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete contentmanagement status status Id internal server error response a status code equal to that given
func (o *DeleteContentmanagementStatusStatusIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteContentmanagementStatusStatusIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteContentmanagementStatusStatusIDInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteContentmanagementStatusStatusIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentmanagementStatusStatusIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentmanagementStatusStatusIDServiceUnavailable creates a DeleteContentmanagementStatusStatusIDServiceUnavailable with default headers values
func NewDeleteContentmanagementStatusStatusIDServiceUnavailable() *DeleteContentmanagementStatusStatusIDServiceUnavailable {
	return &DeleteContentmanagementStatusStatusIDServiceUnavailable{}
}

/*
DeleteContentmanagementStatusStatusIDServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteContentmanagementStatusStatusIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete contentmanagement status status Id service unavailable response has a 2xx status code
func (o *DeleteContentmanagementStatusStatusIDServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete contentmanagement status status Id service unavailable response has a 3xx status code
func (o *DeleteContentmanagementStatusStatusIDServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete contentmanagement status status Id service unavailable response has a 4xx status code
func (o *DeleteContentmanagementStatusStatusIDServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete contentmanagement status status Id service unavailable response has a 5xx status code
func (o *DeleteContentmanagementStatusStatusIDServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete contentmanagement status status Id service unavailable response a status code equal to that given
func (o *DeleteContentmanagementStatusStatusIDServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteContentmanagementStatusStatusIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteContentmanagementStatusStatusIDServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteContentmanagementStatusStatusIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentmanagementStatusStatusIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentmanagementStatusStatusIDGatewayTimeout creates a DeleteContentmanagementStatusStatusIDGatewayTimeout with default headers values
func NewDeleteContentmanagementStatusStatusIDGatewayTimeout() *DeleteContentmanagementStatusStatusIDGatewayTimeout {
	return &DeleteContentmanagementStatusStatusIDGatewayTimeout{}
}

/*
DeleteContentmanagementStatusStatusIDGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteContentmanagementStatusStatusIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete contentmanagement status status Id gateway timeout response has a 2xx status code
func (o *DeleteContentmanagementStatusStatusIDGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete contentmanagement status status Id gateway timeout response has a 3xx status code
func (o *DeleteContentmanagementStatusStatusIDGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete contentmanagement status status Id gateway timeout response has a 4xx status code
func (o *DeleteContentmanagementStatusStatusIDGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete contentmanagement status status Id gateway timeout response has a 5xx status code
func (o *DeleteContentmanagementStatusStatusIDGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete contentmanagement status status Id gateway timeout response a status code equal to that given
func (o *DeleteContentmanagementStatusStatusIDGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteContentmanagementStatusStatusIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteContentmanagementStatusStatusIDGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteContentmanagementStatusStatusIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentmanagementStatusStatusIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentmanagementStatusStatusIDDefault creates a DeleteContentmanagementStatusStatusIDDefault with default headers values
func NewDeleteContentmanagementStatusStatusIDDefault(code int) *DeleteContentmanagementStatusStatusIDDefault {
	return &DeleteContentmanagementStatusStatusIDDefault{
		_statusCode: code,
	}
}

/*
DeleteContentmanagementStatusStatusIDDefault describes a response with status code -1, with default header values.

successful operation
*/
type DeleteContentmanagementStatusStatusIDDefault struct {
	_statusCode int
}

// Code gets the status code for the delete contentmanagement status status Id default response
func (o *DeleteContentmanagementStatusStatusIDDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete contentmanagement status status Id default response has a 2xx status code
func (o *DeleteContentmanagementStatusStatusIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete contentmanagement status status Id default response has a 3xx status code
func (o *DeleteContentmanagementStatusStatusIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete contentmanagement status status Id default response has a 4xx status code
func (o *DeleteContentmanagementStatusStatusIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete contentmanagement status status Id default response has a 5xx status code
func (o *DeleteContentmanagementStatusStatusIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete contentmanagement status status Id default response a status code equal to that given
func (o *DeleteContentmanagementStatusStatusIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteContentmanagementStatusStatusIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusId default ", o._statusCode)
}

func (o *DeleteContentmanagementStatusStatusIDDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/status/{statusId}][%d] deleteContentmanagementStatusStatusId default ", o._statusCode)
}

func (o *DeleteContentmanagementStatusStatusIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
