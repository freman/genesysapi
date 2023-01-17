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

// GetRecordingJobFailedrecordingsReader is a Reader for the GetRecordingJobFailedrecordings structure.
type GetRecordingJobFailedrecordingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRecordingJobFailedrecordingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRecordingJobFailedrecordingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRecordingJobFailedrecordingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRecordingJobFailedrecordingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRecordingJobFailedrecordingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRecordingJobFailedrecordingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRecordingJobFailedrecordingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRecordingJobFailedrecordingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRecordingJobFailedrecordingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRecordingJobFailedrecordingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRecordingJobFailedrecordingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRecordingJobFailedrecordingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRecordingJobFailedrecordingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRecordingJobFailedrecordingsOK creates a GetRecordingJobFailedrecordingsOK with default headers values
func NewGetRecordingJobFailedrecordingsOK() *GetRecordingJobFailedrecordingsOK {
	return &GetRecordingJobFailedrecordingsOK{}
}

/*
GetRecordingJobFailedrecordingsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRecordingJobFailedrecordingsOK struct {
	Payload *models.FailedRecordingEntityListing
}

// IsSuccess returns true when this get recording job failedrecordings o k response has a 2xx status code
func (o *GetRecordingJobFailedrecordingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get recording job failedrecordings o k response has a 3xx status code
func (o *GetRecordingJobFailedrecordingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording job failedrecordings o k response has a 4xx status code
func (o *GetRecordingJobFailedrecordingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get recording job failedrecordings o k response has a 5xx status code
func (o *GetRecordingJobFailedrecordingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording job failedrecordings o k response a status code equal to that given
func (o *GetRecordingJobFailedrecordingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRecordingJobFailedrecordingsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsOK  %+v", 200, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsOK  %+v", 200, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsOK) GetPayload() *models.FailedRecordingEntityListing {
	return o.Payload
}

func (o *GetRecordingJobFailedrecordingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FailedRecordingEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingJobFailedrecordingsBadRequest creates a GetRecordingJobFailedrecordingsBadRequest with default headers values
func NewGetRecordingJobFailedrecordingsBadRequest() *GetRecordingJobFailedrecordingsBadRequest {
	return &GetRecordingJobFailedrecordingsBadRequest{}
}

/*
GetRecordingJobFailedrecordingsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRecordingJobFailedrecordingsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording job failedrecordings bad request response has a 2xx status code
func (o *GetRecordingJobFailedrecordingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording job failedrecordings bad request response has a 3xx status code
func (o *GetRecordingJobFailedrecordingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording job failedrecordings bad request response has a 4xx status code
func (o *GetRecordingJobFailedrecordingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording job failedrecordings bad request response has a 5xx status code
func (o *GetRecordingJobFailedrecordingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording job failedrecordings bad request response a status code equal to that given
func (o *GetRecordingJobFailedrecordingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRecordingJobFailedrecordingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingJobFailedrecordingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingJobFailedrecordingsUnauthorized creates a GetRecordingJobFailedrecordingsUnauthorized with default headers values
func NewGetRecordingJobFailedrecordingsUnauthorized() *GetRecordingJobFailedrecordingsUnauthorized {
	return &GetRecordingJobFailedrecordingsUnauthorized{}
}

/*
GetRecordingJobFailedrecordingsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRecordingJobFailedrecordingsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording job failedrecordings unauthorized response has a 2xx status code
func (o *GetRecordingJobFailedrecordingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording job failedrecordings unauthorized response has a 3xx status code
func (o *GetRecordingJobFailedrecordingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording job failedrecordings unauthorized response has a 4xx status code
func (o *GetRecordingJobFailedrecordingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording job failedrecordings unauthorized response has a 5xx status code
func (o *GetRecordingJobFailedrecordingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording job failedrecordings unauthorized response a status code equal to that given
func (o *GetRecordingJobFailedrecordingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRecordingJobFailedrecordingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingJobFailedrecordingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingJobFailedrecordingsForbidden creates a GetRecordingJobFailedrecordingsForbidden with default headers values
func NewGetRecordingJobFailedrecordingsForbidden() *GetRecordingJobFailedrecordingsForbidden {
	return &GetRecordingJobFailedrecordingsForbidden{}
}

/*
GetRecordingJobFailedrecordingsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetRecordingJobFailedrecordingsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording job failedrecordings forbidden response has a 2xx status code
func (o *GetRecordingJobFailedrecordingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording job failedrecordings forbidden response has a 3xx status code
func (o *GetRecordingJobFailedrecordingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording job failedrecordings forbidden response has a 4xx status code
func (o *GetRecordingJobFailedrecordingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording job failedrecordings forbidden response has a 5xx status code
func (o *GetRecordingJobFailedrecordingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording job failedrecordings forbidden response a status code equal to that given
func (o *GetRecordingJobFailedrecordingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRecordingJobFailedrecordingsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsForbidden  %+v", 403, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsForbidden  %+v", 403, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingJobFailedrecordingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingJobFailedrecordingsNotFound creates a GetRecordingJobFailedrecordingsNotFound with default headers values
func NewGetRecordingJobFailedrecordingsNotFound() *GetRecordingJobFailedrecordingsNotFound {
	return &GetRecordingJobFailedrecordingsNotFound{}
}

/*
GetRecordingJobFailedrecordingsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetRecordingJobFailedrecordingsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording job failedrecordings not found response has a 2xx status code
func (o *GetRecordingJobFailedrecordingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording job failedrecordings not found response has a 3xx status code
func (o *GetRecordingJobFailedrecordingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording job failedrecordings not found response has a 4xx status code
func (o *GetRecordingJobFailedrecordingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording job failedrecordings not found response has a 5xx status code
func (o *GetRecordingJobFailedrecordingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording job failedrecordings not found response a status code equal to that given
func (o *GetRecordingJobFailedrecordingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRecordingJobFailedrecordingsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsNotFound  %+v", 404, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsNotFound  %+v", 404, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingJobFailedrecordingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingJobFailedrecordingsRequestTimeout creates a GetRecordingJobFailedrecordingsRequestTimeout with default headers values
func NewGetRecordingJobFailedrecordingsRequestTimeout() *GetRecordingJobFailedrecordingsRequestTimeout {
	return &GetRecordingJobFailedrecordingsRequestTimeout{}
}

/*
GetRecordingJobFailedrecordingsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRecordingJobFailedrecordingsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording job failedrecordings request timeout response has a 2xx status code
func (o *GetRecordingJobFailedrecordingsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording job failedrecordings request timeout response has a 3xx status code
func (o *GetRecordingJobFailedrecordingsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording job failedrecordings request timeout response has a 4xx status code
func (o *GetRecordingJobFailedrecordingsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording job failedrecordings request timeout response has a 5xx status code
func (o *GetRecordingJobFailedrecordingsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording job failedrecordings request timeout response a status code equal to that given
func (o *GetRecordingJobFailedrecordingsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetRecordingJobFailedrecordingsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingJobFailedrecordingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingJobFailedrecordingsRequestEntityTooLarge creates a GetRecordingJobFailedrecordingsRequestEntityTooLarge with default headers values
func NewGetRecordingJobFailedrecordingsRequestEntityTooLarge() *GetRecordingJobFailedrecordingsRequestEntityTooLarge {
	return &GetRecordingJobFailedrecordingsRequestEntityTooLarge{}
}

/*
GetRecordingJobFailedrecordingsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetRecordingJobFailedrecordingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording job failedrecordings request entity too large response has a 2xx status code
func (o *GetRecordingJobFailedrecordingsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording job failedrecordings request entity too large response has a 3xx status code
func (o *GetRecordingJobFailedrecordingsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording job failedrecordings request entity too large response has a 4xx status code
func (o *GetRecordingJobFailedrecordingsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording job failedrecordings request entity too large response has a 5xx status code
func (o *GetRecordingJobFailedrecordingsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording job failedrecordings request entity too large response a status code equal to that given
func (o *GetRecordingJobFailedrecordingsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRecordingJobFailedrecordingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingJobFailedrecordingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingJobFailedrecordingsUnsupportedMediaType creates a GetRecordingJobFailedrecordingsUnsupportedMediaType with default headers values
func NewGetRecordingJobFailedrecordingsUnsupportedMediaType() *GetRecordingJobFailedrecordingsUnsupportedMediaType {
	return &GetRecordingJobFailedrecordingsUnsupportedMediaType{}
}

/*
GetRecordingJobFailedrecordingsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRecordingJobFailedrecordingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording job failedrecordings unsupported media type response has a 2xx status code
func (o *GetRecordingJobFailedrecordingsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording job failedrecordings unsupported media type response has a 3xx status code
func (o *GetRecordingJobFailedrecordingsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording job failedrecordings unsupported media type response has a 4xx status code
func (o *GetRecordingJobFailedrecordingsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording job failedrecordings unsupported media type response has a 5xx status code
func (o *GetRecordingJobFailedrecordingsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording job failedrecordings unsupported media type response a status code equal to that given
func (o *GetRecordingJobFailedrecordingsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRecordingJobFailedrecordingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingJobFailedrecordingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingJobFailedrecordingsTooManyRequests creates a GetRecordingJobFailedrecordingsTooManyRequests with default headers values
func NewGetRecordingJobFailedrecordingsTooManyRequests() *GetRecordingJobFailedrecordingsTooManyRequests {
	return &GetRecordingJobFailedrecordingsTooManyRequests{}
}

/*
GetRecordingJobFailedrecordingsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRecordingJobFailedrecordingsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording job failedrecordings too many requests response has a 2xx status code
func (o *GetRecordingJobFailedrecordingsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording job failedrecordings too many requests response has a 3xx status code
func (o *GetRecordingJobFailedrecordingsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording job failedrecordings too many requests response has a 4xx status code
func (o *GetRecordingJobFailedrecordingsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording job failedrecordings too many requests response has a 5xx status code
func (o *GetRecordingJobFailedrecordingsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording job failedrecordings too many requests response a status code equal to that given
func (o *GetRecordingJobFailedrecordingsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRecordingJobFailedrecordingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingJobFailedrecordingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingJobFailedrecordingsInternalServerError creates a GetRecordingJobFailedrecordingsInternalServerError with default headers values
func NewGetRecordingJobFailedrecordingsInternalServerError() *GetRecordingJobFailedrecordingsInternalServerError {
	return &GetRecordingJobFailedrecordingsInternalServerError{}
}

/*
GetRecordingJobFailedrecordingsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRecordingJobFailedrecordingsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording job failedrecordings internal server error response has a 2xx status code
func (o *GetRecordingJobFailedrecordingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording job failedrecordings internal server error response has a 3xx status code
func (o *GetRecordingJobFailedrecordingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording job failedrecordings internal server error response has a 4xx status code
func (o *GetRecordingJobFailedrecordingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get recording job failedrecordings internal server error response has a 5xx status code
func (o *GetRecordingJobFailedrecordingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get recording job failedrecordings internal server error response a status code equal to that given
func (o *GetRecordingJobFailedrecordingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRecordingJobFailedrecordingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingJobFailedrecordingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingJobFailedrecordingsServiceUnavailable creates a GetRecordingJobFailedrecordingsServiceUnavailable with default headers values
func NewGetRecordingJobFailedrecordingsServiceUnavailable() *GetRecordingJobFailedrecordingsServiceUnavailable {
	return &GetRecordingJobFailedrecordingsServiceUnavailable{}
}

/*
GetRecordingJobFailedrecordingsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRecordingJobFailedrecordingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording job failedrecordings service unavailable response has a 2xx status code
func (o *GetRecordingJobFailedrecordingsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording job failedrecordings service unavailable response has a 3xx status code
func (o *GetRecordingJobFailedrecordingsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording job failedrecordings service unavailable response has a 4xx status code
func (o *GetRecordingJobFailedrecordingsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get recording job failedrecordings service unavailable response has a 5xx status code
func (o *GetRecordingJobFailedrecordingsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get recording job failedrecordings service unavailable response a status code equal to that given
func (o *GetRecordingJobFailedrecordingsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRecordingJobFailedrecordingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingJobFailedrecordingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingJobFailedrecordingsGatewayTimeout creates a GetRecordingJobFailedrecordingsGatewayTimeout with default headers values
func NewGetRecordingJobFailedrecordingsGatewayTimeout() *GetRecordingJobFailedrecordingsGatewayTimeout {
	return &GetRecordingJobFailedrecordingsGatewayTimeout{}
}

/*
GetRecordingJobFailedrecordingsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetRecordingJobFailedrecordingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording job failedrecordings gateway timeout response has a 2xx status code
func (o *GetRecordingJobFailedrecordingsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording job failedrecordings gateway timeout response has a 3xx status code
func (o *GetRecordingJobFailedrecordingsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording job failedrecordings gateway timeout response has a 4xx status code
func (o *GetRecordingJobFailedrecordingsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get recording job failedrecordings gateway timeout response has a 5xx status code
func (o *GetRecordingJobFailedrecordingsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get recording job failedrecordings gateway timeout response a status code equal to that given
func (o *GetRecordingJobFailedrecordingsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetRecordingJobFailedrecordingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/jobs/{jobId}/failedrecordings][%d] getRecordingJobFailedrecordingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRecordingJobFailedrecordingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingJobFailedrecordingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
