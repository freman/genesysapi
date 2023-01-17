// Code generated by go-swagger; DO NOT EDIT.

package telephony_providers_edge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostTelephonyProvidersEdgeLogsJobUploadReader is a Reader for the PostTelephonyProvidersEdgeLogsJobUpload structure.
type PostTelephonyProvidersEdgeLogsJobUploadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTelephonyProvidersEdgeLogsJobUploadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostTelephonyProvidersEdgeLogsJobUploadAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTelephonyProvidersEdgeLogsJobUploadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTelephonyProvidersEdgeLogsJobUploadUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTelephonyProvidersEdgeLogsJobUploadForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTelephonyProvidersEdgeLogsJobUploadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostTelephonyProvidersEdgeLogsJobUploadRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTelephonyProvidersEdgeLogsJobUploadUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTelephonyProvidersEdgeLogsJobUploadTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTelephonyProvidersEdgeLogsJobUploadInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTelephonyProvidersEdgeLogsJobUploadServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTelephonyProvidersEdgeLogsJobUploadGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTelephonyProvidersEdgeLogsJobUploadAccepted creates a PostTelephonyProvidersEdgeLogsJobUploadAccepted with default headers values
func NewPostTelephonyProvidersEdgeLogsJobUploadAccepted() *PostTelephonyProvidersEdgeLogsJobUploadAccepted {
	return &PostTelephonyProvidersEdgeLogsJobUploadAccepted{}
}

/*
PostTelephonyProvidersEdgeLogsJobUploadAccepted describes a response with status code 202, with default header values.

Accepted - Files are being uploaded to the job.  Watch the uploadStatus property on the job files.
*/
type PostTelephonyProvidersEdgeLogsJobUploadAccepted struct {
}

// IsSuccess returns true when this post telephony providers edge logs job upload accepted response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post telephony providers edge logs job upload accepted response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logs job upload accepted response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this post telephony providers edge logs job upload accepted response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge logs job upload accepted response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogsJobUploadAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadAccepted ", 202)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadAccepted) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadAccepted ", 202)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostTelephonyProvidersEdgeLogsJobUploadBadRequest creates a PostTelephonyProvidersEdgeLogsJobUploadBadRequest with default headers values
func NewPostTelephonyProvidersEdgeLogsJobUploadBadRequest() *PostTelephonyProvidersEdgeLogsJobUploadBadRequest {
	return &PostTelephonyProvidersEdgeLogsJobUploadBadRequest{}
}

/*
PostTelephonyProvidersEdgeLogsJobUploadBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTelephonyProvidersEdgeLogsJobUploadBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge logs job upload bad request response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge logs job upload bad request response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logs job upload bad request response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge logs job upload bad request response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge logs job upload bad request response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogsJobUploadBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogsJobUploadUnauthorized creates a PostTelephonyProvidersEdgeLogsJobUploadUnauthorized with default headers values
func NewPostTelephonyProvidersEdgeLogsJobUploadUnauthorized() *PostTelephonyProvidersEdgeLogsJobUploadUnauthorized {
	return &PostTelephonyProvidersEdgeLogsJobUploadUnauthorized{}
}

/*
PostTelephonyProvidersEdgeLogsJobUploadUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTelephonyProvidersEdgeLogsJobUploadUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge logs job upload unauthorized response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge logs job upload unauthorized response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logs job upload unauthorized response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge logs job upload unauthorized response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge logs job upload unauthorized response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogsJobUploadUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogsJobUploadForbidden creates a PostTelephonyProvidersEdgeLogsJobUploadForbidden with default headers values
func NewPostTelephonyProvidersEdgeLogsJobUploadForbidden() *PostTelephonyProvidersEdgeLogsJobUploadForbidden {
	return &PostTelephonyProvidersEdgeLogsJobUploadForbidden{}
}

/*
PostTelephonyProvidersEdgeLogsJobUploadForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostTelephonyProvidersEdgeLogsJobUploadForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge logs job upload forbidden response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge logs job upload forbidden response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logs job upload forbidden response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge logs job upload forbidden response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge logs job upload forbidden response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogsJobUploadForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogsJobUploadNotFound creates a PostTelephonyProvidersEdgeLogsJobUploadNotFound with default headers values
func NewPostTelephonyProvidersEdgeLogsJobUploadNotFound() *PostTelephonyProvidersEdgeLogsJobUploadNotFound {
	return &PostTelephonyProvidersEdgeLogsJobUploadNotFound{}
}

/*
PostTelephonyProvidersEdgeLogsJobUploadNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostTelephonyProvidersEdgeLogsJobUploadNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge logs job upload not found response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge logs job upload not found response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logs job upload not found response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge logs job upload not found response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge logs job upload not found response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogsJobUploadNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogsJobUploadRequestTimeout creates a PostTelephonyProvidersEdgeLogsJobUploadRequestTimeout with default headers values
func NewPostTelephonyProvidersEdgeLogsJobUploadRequestTimeout() *PostTelephonyProvidersEdgeLogsJobUploadRequestTimeout {
	return &PostTelephonyProvidersEdgeLogsJobUploadRequestTimeout{}
}

/*
PostTelephonyProvidersEdgeLogsJobUploadRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostTelephonyProvidersEdgeLogsJobUploadRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge logs job upload request timeout response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge logs job upload request timeout response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logs job upload request timeout response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge logs job upload request timeout response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge logs job upload request timeout response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogsJobUploadRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge creates a PostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge with default headers values
func NewPostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge() *PostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge {
	return &PostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge{}
}

/*
PostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge logs job upload request entity too large response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge logs job upload request entity too large response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logs job upload request entity too large response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge logs job upload request entity too large response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge logs job upload request entity too large response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogsJobUploadUnsupportedMediaType creates a PostTelephonyProvidersEdgeLogsJobUploadUnsupportedMediaType with default headers values
func NewPostTelephonyProvidersEdgeLogsJobUploadUnsupportedMediaType() *PostTelephonyProvidersEdgeLogsJobUploadUnsupportedMediaType {
	return &PostTelephonyProvidersEdgeLogsJobUploadUnsupportedMediaType{}
}

/*
PostTelephonyProvidersEdgeLogsJobUploadUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTelephonyProvidersEdgeLogsJobUploadUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge logs job upload unsupported media type response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge logs job upload unsupported media type response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logs job upload unsupported media type response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge logs job upload unsupported media type response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge logs job upload unsupported media type response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogsJobUploadUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogsJobUploadTooManyRequests creates a PostTelephonyProvidersEdgeLogsJobUploadTooManyRequests with default headers values
func NewPostTelephonyProvidersEdgeLogsJobUploadTooManyRequests() *PostTelephonyProvidersEdgeLogsJobUploadTooManyRequests {
	return &PostTelephonyProvidersEdgeLogsJobUploadTooManyRequests{}
}

/*
PostTelephonyProvidersEdgeLogsJobUploadTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostTelephonyProvidersEdgeLogsJobUploadTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge logs job upload too many requests response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge logs job upload too many requests response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logs job upload too many requests response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post telephony providers edge logs job upload too many requests response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post telephony providers edge logs job upload too many requests response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogsJobUploadTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogsJobUploadInternalServerError creates a PostTelephonyProvidersEdgeLogsJobUploadInternalServerError with default headers values
func NewPostTelephonyProvidersEdgeLogsJobUploadInternalServerError() *PostTelephonyProvidersEdgeLogsJobUploadInternalServerError {
	return &PostTelephonyProvidersEdgeLogsJobUploadInternalServerError{}
}

/*
PostTelephonyProvidersEdgeLogsJobUploadInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTelephonyProvidersEdgeLogsJobUploadInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge logs job upload internal server error response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge logs job upload internal server error response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logs job upload internal server error response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post telephony providers edge logs job upload internal server error response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post telephony providers edge logs job upload internal server error response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogsJobUploadInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogsJobUploadServiceUnavailable creates a PostTelephonyProvidersEdgeLogsJobUploadServiceUnavailable with default headers values
func NewPostTelephonyProvidersEdgeLogsJobUploadServiceUnavailable() *PostTelephonyProvidersEdgeLogsJobUploadServiceUnavailable {
	return &PostTelephonyProvidersEdgeLogsJobUploadServiceUnavailable{}
}

/*
PostTelephonyProvidersEdgeLogsJobUploadServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTelephonyProvidersEdgeLogsJobUploadServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge logs job upload service unavailable response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge logs job upload service unavailable response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logs job upload service unavailable response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post telephony providers edge logs job upload service unavailable response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post telephony providers edge logs job upload service unavailable response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogsJobUploadServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogsJobUploadGatewayTimeout creates a PostTelephonyProvidersEdgeLogsJobUploadGatewayTimeout with default headers values
func NewPostTelephonyProvidersEdgeLogsJobUploadGatewayTimeout() *PostTelephonyProvidersEdgeLogsJobUploadGatewayTimeout {
	return &PostTelephonyProvidersEdgeLogsJobUploadGatewayTimeout{}
}

/*
PostTelephonyProvidersEdgeLogsJobUploadGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostTelephonyProvidersEdgeLogsJobUploadGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post telephony providers edge logs job upload gateway timeout response has a 2xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post telephony providers edge logs job upload gateway timeout response has a 3xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post telephony providers edge logs job upload gateway timeout response has a 4xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post telephony providers edge logs job upload gateway timeout response has a 5xx status code
func (o *PostTelephonyProvidersEdgeLogsJobUploadGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post telephony providers edge logs job upload gateway timeout response a status code equal to that given
func (o *PostTelephonyProvidersEdgeLogsJobUploadGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
