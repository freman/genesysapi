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

/*PostTelephonyProvidersEdgeLogsJobUploadAccepted handles this case with default header values.

Accepted - Files are being uploaded to the job.  Watch the uploadStatus property on the job files.
*/
type PostTelephonyProvidersEdgeLogsJobUploadAccepted struct {
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload][%d] postTelephonyProvidersEdgeLogsJobUploadAccepted ", 202)
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostTelephonyProvidersEdgeLogsJobUploadBadRequest creates a PostTelephonyProvidersEdgeLogsJobUploadBadRequest with default headers values
func NewPostTelephonyProvidersEdgeLogsJobUploadBadRequest() *PostTelephonyProvidersEdgeLogsJobUploadBadRequest {
	return &PostTelephonyProvidersEdgeLogsJobUploadBadRequest{}
}

/*PostTelephonyProvidersEdgeLogsJobUploadBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTelephonyProvidersEdgeLogsJobUploadBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadBadRequest) Error() string {
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

/*PostTelephonyProvidersEdgeLogsJobUploadUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTelephonyProvidersEdgeLogsJobUploadUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadUnauthorized) Error() string {
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

/*PostTelephonyProvidersEdgeLogsJobUploadForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostTelephonyProvidersEdgeLogsJobUploadForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadForbidden) Error() string {
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

/*PostTelephonyProvidersEdgeLogsJobUploadNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostTelephonyProvidersEdgeLogsJobUploadNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadNotFound) Error() string {
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

// NewPostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge creates a PostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge with default headers values
func NewPostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge() *PostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge {
	return &PostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge{}
}

/*PostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadRequestEntityTooLarge) Error() string {
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

/*PostTelephonyProvidersEdgeLogsJobUploadUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTelephonyProvidersEdgeLogsJobUploadUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadUnsupportedMediaType) Error() string {
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

/*PostTelephonyProvidersEdgeLogsJobUploadTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostTelephonyProvidersEdgeLogsJobUploadTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadTooManyRequests) Error() string {
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

/*PostTelephonyProvidersEdgeLogsJobUploadInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTelephonyProvidersEdgeLogsJobUploadInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadInternalServerError) Error() string {
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

/*PostTelephonyProvidersEdgeLogsJobUploadServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTelephonyProvidersEdgeLogsJobUploadServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadServiceUnavailable) Error() string {
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

/*PostTelephonyProvidersEdgeLogsJobUploadGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostTelephonyProvidersEdgeLogsJobUploadGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogsJobUploadGatewayTimeout) Error() string {
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
