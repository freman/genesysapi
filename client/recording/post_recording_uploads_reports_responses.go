// Code generated by go-swagger; DO NOT EDIT.

package recording

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostRecordingUploadsReportsReader is a Reader for the PostRecordingUploadsReports structure.
type PostRecordingUploadsReportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRecordingUploadsReportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostRecordingUploadsReportsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRecordingUploadsReportsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostRecordingUploadsReportsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRecordingUploadsReportsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRecordingUploadsReportsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostRecordingUploadsReportsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostRecordingUploadsReportsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostRecordingUploadsReportsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostRecordingUploadsReportsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostRecordingUploadsReportsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostRecordingUploadsReportsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostRecordingUploadsReportsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRecordingUploadsReportsAccepted creates a PostRecordingUploadsReportsAccepted with default headers values
func NewPostRecordingUploadsReportsAccepted() *PostRecordingUploadsReportsAccepted {
	return &PostRecordingUploadsReportsAccepted{}
}

/*
PostRecordingUploadsReportsAccepted describes a response with status code 202, with default header values.

Accepted - preparing report.
*/
type PostRecordingUploadsReportsAccepted struct {
	Payload *models.RecordingUploadReport
}

// IsSuccess returns true when this post recording uploads reports accepted response has a 2xx status code
func (o *PostRecordingUploadsReportsAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post recording uploads reports accepted response has a 3xx status code
func (o *PostRecordingUploadsReportsAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording uploads reports accepted response has a 4xx status code
func (o *PostRecordingUploadsReportsAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this post recording uploads reports accepted response has a 5xx status code
func (o *PostRecordingUploadsReportsAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording uploads reports accepted response a status code equal to that given
func (o *PostRecordingUploadsReportsAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PostRecordingUploadsReportsAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsAccepted  %+v", 202, o.Payload)
}

func (o *PostRecordingUploadsReportsAccepted) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsAccepted  %+v", 202, o.Payload)
}

func (o *PostRecordingUploadsReportsAccepted) GetPayload() *models.RecordingUploadReport {
	return o.Payload
}

func (o *PostRecordingUploadsReportsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecordingUploadReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingUploadsReportsBadRequest creates a PostRecordingUploadsReportsBadRequest with default headers values
func NewPostRecordingUploadsReportsBadRequest() *PostRecordingUploadsReportsBadRequest {
	return &PostRecordingUploadsReportsBadRequest{}
}

/*
PostRecordingUploadsReportsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostRecordingUploadsReportsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording uploads reports bad request response has a 2xx status code
func (o *PostRecordingUploadsReportsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording uploads reports bad request response has a 3xx status code
func (o *PostRecordingUploadsReportsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording uploads reports bad request response has a 4xx status code
func (o *PostRecordingUploadsReportsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording uploads reports bad request response has a 5xx status code
func (o *PostRecordingUploadsReportsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording uploads reports bad request response a status code equal to that given
func (o *PostRecordingUploadsReportsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostRecordingUploadsReportsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsBadRequest  %+v", 400, o.Payload)
}

func (o *PostRecordingUploadsReportsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsBadRequest  %+v", 400, o.Payload)
}

func (o *PostRecordingUploadsReportsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingUploadsReportsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingUploadsReportsUnauthorized creates a PostRecordingUploadsReportsUnauthorized with default headers values
func NewPostRecordingUploadsReportsUnauthorized() *PostRecordingUploadsReportsUnauthorized {
	return &PostRecordingUploadsReportsUnauthorized{}
}

/*
PostRecordingUploadsReportsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostRecordingUploadsReportsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording uploads reports unauthorized response has a 2xx status code
func (o *PostRecordingUploadsReportsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording uploads reports unauthorized response has a 3xx status code
func (o *PostRecordingUploadsReportsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording uploads reports unauthorized response has a 4xx status code
func (o *PostRecordingUploadsReportsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording uploads reports unauthorized response has a 5xx status code
func (o *PostRecordingUploadsReportsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording uploads reports unauthorized response a status code equal to that given
func (o *PostRecordingUploadsReportsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostRecordingUploadsReportsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRecordingUploadsReportsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRecordingUploadsReportsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingUploadsReportsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingUploadsReportsForbidden creates a PostRecordingUploadsReportsForbidden with default headers values
func NewPostRecordingUploadsReportsForbidden() *PostRecordingUploadsReportsForbidden {
	return &PostRecordingUploadsReportsForbidden{}
}

/*
PostRecordingUploadsReportsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostRecordingUploadsReportsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording uploads reports forbidden response has a 2xx status code
func (o *PostRecordingUploadsReportsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording uploads reports forbidden response has a 3xx status code
func (o *PostRecordingUploadsReportsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording uploads reports forbidden response has a 4xx status code
func (o *PostRecordingUploadsReportsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording uploads reports forbidden response has a 5xx status code
func (o *PostRecordingUploadsReportsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording uploads reports forbidden response a status code equal to that given
func (o *PostRecordingUploadsReportsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostRecordingUploadsReportsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsForbidden  %+v", 403, o.Payload)
}

func (o *PostRecordingUploadsReportsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsForbidden  %+v", 403, o.Payload)
}

func (o *PostRecordingUploadsReportsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingUploadsReportsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingUploadsReportsNotFound creates a PostRecordingUploadsReportsNotFound with default headers values
func NewPostRecordingUploadsReportsNotFound() *PostRecordingUploadsReportsNotFound {
	return &PostRecordingUploadsReportsNotFound{}
}

/*
PostRecordingUploadsReportsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostRecordingUploadsReportsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording uploads reports not found response has a 2xx status code
func (o *PostRecordingUploadsReportsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording uploads reports not found response has a 3xx status code
func (o *PostRecordingUploadsReportsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording uploads reports not found response has a 4xx status code
func (o *PostRecordingUploadsReportsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording uploads reports not found response has a 5xx status code
func (o *PostRecordingUploadsReportsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording uploads reports not found response a status code equal to that given
func (o *PostRecordingUploadsReportsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostRecordingUploadsReportsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsNotFound  %+v", 404, o.Payload)
}

func (o *PostRecordingUploadsReportsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsNotFound  %+v", 404, o.Payload)
}

func (o *PostRecordingUploadsReportsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingUploadsReportsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingUploadsReportsRequestTimeout creates a PostRecordingUploadsReportsRequestTimeout with default headers values
func NewPostRecordingUploadsReportsRequestTimeout() *PostRecordingUploadsReportsRequestTimeout {
	return &PostRecordingUploadsReportsRequestTimeout{}
}

/*
PostRecordingUploadsReportsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostRecordingUploadsReportsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording uploads reports request timeout response has a 2xx status code
func (o *PostRecordingUploadsReportsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording uploads reports request timeout response has a 3xx status code
func (o *PostRecordingUploadsReportsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording uploads reports request timeout response has a 4xx status code
func (o *PostRecordingUploadsReportsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording uploads reports request timeout response has a 5xx status code
func (o *PostRecordingUploadsReportsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording uploads reports request timeout response a status code equal to that given
func (o *PostRecordingUploadsReportsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostRecordingUploadsReportsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostRecordingUploadsReportsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostRecordingUploadsReportsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingUploadsReportsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingUploadsReportsRequestEntityTooLarge creates a PostRecordingUploadsReportsRequestEntityTooLarge with default headers values
func NewPostRecordingUploadsReportsRequestEntityTooLarge() *PostRecordingUploadsReportsRequestEntityTooLarge {
	return &PostRecordingUploadsReportsRequestEntityTooLarge{}
}

/*
PostRecordingUploadsReportsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostRecordingUploadsReportsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording uploads reports request entity too large response has a 2xx status code
func (o *PostRecordingUploadsReportsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording uploads reports request entity too large response has a 3xx status code
func (o *PostRecordingUploadsReportsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording uploads reports request entity too large response has a 4xx status code
func (o *PostRecordingUploadsReportsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording uploads reports request entity too large response has a 5xx status code
func (o *PostRecordingUploadsReportsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording uploads reports request entity too large response a status code equal to that given
func (o *PostRecordingUploadsReportsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostRecordingUploadsReportsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRecordingUploadsReportsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRecordingUploadsReportsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingUploadsReportsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingUploadsReportsUnsupportedMediaType creates a PostRecordingUploadsReportsUnsupportedMediaType with default headers values
func NewPostRecordingUploadsReportsUnsupportedMediaType() *PostRecordingUploadsReportsUnsupportedMediaType {
	return &PostRecordingUploadsReportsUnsupportedMediaType{}
}

/*
PostRecordingUploadsReportsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostRecordingUploadsReportsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording uploads reports unsupported media type response has a 2xx status code
func (o *PostRecordingUploadsReportsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording uploads reports unsupported media type response has a 3xx status code
func (o *PostRecordingUploadsReportsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording uploads reports unsupported media type response has a 4xx status code
func (o *PostRecordingUploadsReportsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording uploads reports unsupported media type response has a 5xx status code
func (o *PostRecordingUploadsReportsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording uploads reports unsupported media type response a status code equal to that given
func (o *PostRecordingUploadsReportsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostRecordingUploadsReportsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRecordingUploadsReportsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRecordingUploadsReportsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingUploadsReportsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingUploadsReportsTooManyRequests creates a PostRecordingUploadsReportsTooManyRequests with default headers values
func NewPostRecordingUploadsReportsTooManyRequests() *PostRecordingUploadsReportsTooManyRequests {
	return &PostRecordingUploadsReportsTooManyRequests{}
}

/*
PostRecordingUploadsReportsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostRecordingUploadsReportsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording uploads reports too many requests response has a 2xx status code
func (o *PostRecordingUploadsReportsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording uploads reports too many requests response has a 3xx status code
func (o *PostRecordingUploadsReportsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording uploads reports too many requests response has a 4xx status code
func (o *PostRecordingUploadsReportsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording uploads reports too many requests response has a 5xx status code
func (o *PostRecordingUploadsReportsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording uploads reports too many requests response a status code equal to that given
func (o *PostRecordingUploadsReportsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostRecordingUploadsReportsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRecordingUploadsReportsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRecordingUploadsReportsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingUploadsReportsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingUploadsReportsInternalServerError creates a PostRecordingUploadsReportsInternalServerError with default headers values
func NewPostRecordingUploadsReportsInternalServerError() *PostRecordingUploadsReportsInternalServerError {
	return &PostRecordingUploadsReportsInternalServerError{}
}

/*
PostRecordingUploadsReportsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostRecordingUploadsReportsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording uploads reports internal server error response has a 2xx status code
func (o *PostRecordingUploadsReportsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording uploads reports internal server error response has a 3xx status code
func (o *PostRecordingUploadsReportsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording uploads reports internal server error response has a 4xx status code
func (o *PostRecordingUploadsReportsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post recording uploads reports internal server error response has a 5xx status code
func (o *PostRecordingUploadsReportsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post recording uploads reports internal server error response a status code equal to that given
func (o *PostRecordingUploadsReportsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostRecordingUploadsReportsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRecordingUploadsReportsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRecordingUploadsReportsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingUploadsReportsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingUploadsReportsServiceUnavailable creates a PostRecordingUploadsReportsServiceUnavailable with default headers values
func NewPostRecordingUploadsReportsServiceUnavailable() *PostRecordingUploadsReportsServiceUnavailable {
	return &PostRecordingUploadsReportsServiceUnavailable{}
}

/*
PostRecordingUploadsReportsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostRecordingUploadsReportsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording uploads reports service unavailable response has a 2xx status code
func (o *PostRecordingUploadsReportsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording uploads reports service unavailable response has a 3xx status code
func (o *PostRecordingUploadsReportsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording uploads reports service unavailable response has a 4xx status code
func (o *PostRecordingUploadsReportsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post recording uploads reports service unavailable response has a 5xx status code
func (o *PostRecordingUploadsReportsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post recording uploads reports service unavailable response a status code equal to that given
func (o *PostRecordingUploadsReportsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostRecordingUploadsReportsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRecordingUploadsReportsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRecordingUploadsReportsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingUploadsReportsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingUploadsReportsGatewayTimeout creates a PostRecordingUploadsReportsGatewayTimeout with default headers values
func NewPostRecordingUploadsReportsGatewayTimeout() *PostRecordingUploadsReportsGatewayTimeout {
	return &PostRecordingUploadsReportsGatewayTimeout{}
}

/*
PostRecordingUploadsReportsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostRecordingUploadsReportsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording uploads reports gateway timeout response has a 2xx status code
func (o *PostRecordingUploadsReportsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording uploads reports gateway timeout response has a 3xx status code
func (o *PostRecordingUploadsReportsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording uploads reports gateway timeout response has a 4xx status code
func (o *PostRecordingUploadsReportsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post recording uploads reports gateway timeout response has a 5xx status code
func (o *PostRecordingUploadsReportsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post recording uploads reports gateway timeout response a status code equal to that given
func (o *PostRecordingUploadsReportsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostRecordingUploadsReportsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRecordingUploadsReportsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/uploads/reports][%d] postRecordingUploadsReportsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRecordingUploadsReportsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingUploadsReportsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}