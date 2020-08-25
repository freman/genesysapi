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

// PostTelephonyProvidersEdgeLogsJobsReader is a Reader for the PostTelephonyProvidersEdgeLogsJobs structure.
type PostTelephonyProvidersEdgeLogsJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTelephonyProvidersEdgeLogsJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostTelephonyProvidersEdgeLogsJobsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTelephonyProvidersEdgeLogsJobsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTelephonyProvidersEdgeLogsJobsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTelephonyProvidersEdgeLogsJobsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTelephonyProvidersEdgeLogsJobsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTelephonyProvidersEdgeLogsJobsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTelephonyProvidersEdgeLogsJobsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTelephonyProvidersEdgeLogsJobsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTelephonyProvidersEdgeLogsJobsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTelephonyProvidersEdgeLogsJobsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTelephonyProvidersEdgeLogsJobsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTelephonyProvidersEdgeLogsJobsAccepted creates a PostTelephonyProvidersEdgeLogsJobsAccepted with default headers values
func NewPostTelephonyProvidersEdgeLogsJobsAccepted() *PostTelephonyProvidersEdgeLogsJobsAccepted {
	return &PostTelephonyProvidersEdgeLogsJobsAccepted{}
}

/*PostTelephonyProvidersEdgeLogsJobsAccepted handles this case with default header values.

Accepted - Job is being processed.  The job ID is returned.
*/
type PostTelephonyProvidersEdgeLogsJobsAccepted struct {
	Payload *models.EdgeLogsJobResponse
}

func (o *PostTelephonyProvidersEdgeLogsJobsAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs][%d] postTelephonyProvidersEdgeLogsJobsAccepted  %+v", 202, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobsAccepted) GetPayload() *models.EdgeLogsJobResponse {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogsJobsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgeLogsJobResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogsJobsBadRequest creates a PostTelephonyProvidersEdgeLogsJobsBadRequest with default headers values
func NewPostTelephonyProvidersEdgeLogsJobsBadRequest() *PostTelephonyProvidersEdgeLogsJobsBadRequest {
	return &PostTelephonyProvidersEdgeLogsJobsBadRequest{}
}

/*PostTelephonyProvidersEdgeLogsJobsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTelephonyProvidersEdgeLogsJobsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogsJobsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs][%d] postTelephonyProvidersEdgeLogsJobsBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogsJobsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogsJobsUnauthorized creates a PostTelephonyProvidersEdgeLogsJobsUnauthorized with default headers values
func NewPostTelephonyProvidersEdgeLogsJobsUnauthorized() *PostTelephonyProvidersEdgeLogsJobsUnauthorized {
	return &PostTelephonyProvidersEdgeLogsJobsUnauthorized{}
}

/*PostTelephonyProvidersEdgeLogsJobsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTelephonyProvidersEdgeLogsJobsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogsJobsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs][%d] postTelephonyProvidersEdgeLogsJobsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogsJobsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogsJobsForbidden creates a PostTelephonyProvidersEdgeLogsJobsForbidden with default headers values
func NewPostTelephonyProvidersEdgeLogsJobsForbidden() *PostTelephonyProvidersEdgeLogsJobsForbidden {
	return &PostTelephonyProvidersEdgeLogsJobsForbidden{}
}

/*PostTelephonyProvidersEdgeLogsJobsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostTelephonyProvidersEdgeLogsJobsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogsJobsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs][%d] postTelephonyProvidersEdgeLogsJobsForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogsJobsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogsJobsNotFound creates a PostTelephonyProvidersEdgeLogsJobsNotFound with default headers values
func NewPostTelephonyProvidersEdgeLogsJobsNotFound() *PostTelephonyProvidersEdgeLogsJobsNotFound {
	return &PostTelephonyProvidersEdgeLogsJobsNotFound{}
}

/*PostTelephonyProvidersEdgeLogsJobsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostTelephonyProvidersEdgeLogsJobsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogsJobsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs][%d] postTelephonyProvidersEdgeLogsJobsNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogsJobsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogsJobsRequestEntityTooLarge creates a PostTelephonyProvidersEdgeLogsJobsRequestEntityTooLarge with default headers values
func NewPostTelephonyProvidersEdgeLogsJobsRequestEntityTooLarge() *PostTelephonyProvidersEdgeLogsJobsRequestEntityTooLarge {
	return &PostTelephonyProvidersEdgeLogsJobsRequestEntityTooLarge{}
}

/*PostTelephonyProvidersEdgeLogsJobsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostTelephonyProvidersEdgeLogsJobsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogsJobsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs][%d] postTelephonyProvidersEdgeLogsJobsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogsJobsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogsJobsUnsupportedMediaType creates a PostTelephonyProvidersEdgeLogsJobsUnsupportedMediaType with default headers values
func NewPostTelephonyProvidersEdgeLogsJobsUnsupportedMediaType() *PostTelephonyProvidersEdgeLogsJobsUnsupportedMediaType {
	return &PostTelephonyProvidersEdgeLogsJobsUnsupportedMediaType{}
}

/*PostTelephonyProvidersEdgeLogsJobsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTelephonyProvidersEdgeLogsJobsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogsJobsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs][%d] postTelephonyProvidersEdgeLogsJobsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogsJobsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogsJobsTooManyRequests creates a PostTelephonyProvidersEdgeLogsJobsTooManyRequests with default headers values
func NewPostTelephonyProvidersEdgeLogsJobsTooManyRequests() *PostTelephonyProvidersEdgeLogsJobsTooManyRequests {
	return &PostTelephonyProvidersEdgeLogsJobsTooManyRequests{}
}

/*PostTelephonyProvidersEdgeLogsJobsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostTelephonyProvidersEdgeLogsJobsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogsJobsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs][%d] postTelephonyProvidersEdgeLogsJobsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogsJobsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogsJobsInternalServerError creates a PostTelephonyProvidersEdgeLogsJobsInternalServerError with default headers values
func NewPostTelephonyProvidersEdgeLogsJobsInternalServerError() *PostTelephonyProvidersEdgeLogsJobsInternalServerError {
	return &PostTelephonyProvidersEdgeLogsJobsInternalServerError{}
}

/*PostTelephonyProvidersEdgeLogsJobsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTelephonyProvidersEdgeLogsJobsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogsJobsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs][%d] postTelephonyProvidersEdgeLogsJobsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogsJobsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogsJobsServiceUnavailable creates a PostTelephonyProvidersEdgeLogsJobsServiceUnavailable with default headers values
func NewPostTelephonyProvidersEdgeLogsJobsServiceUnavailable() *PostTelephonyProvidersEdgeLogsJobsServiceUnavailable {
	return &PostTelephonyProvidersEdgeLogsJobsServiceUnavailable{}
}

/*PostTelephonyProvidersEdgeLogsJobsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTelephonyProvidersEdgeLogsJobsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogsJobsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs][%d] postTelephonyProvidersEdgeLogsJobsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogsJobsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogsJobsGatewayTimeout creates a PostTelephonyProvidersEdgeLogsJobsGatewayTimeout with default headers values
func NewPostTelephonyProvidersEdgeLogsJobsGatewayTimeout() *PostTelephonyProvidersEdgeLogsJobsGatewayTimeout {
	return &PostTelephonyProvidersEdgeLogsJobsGatewayTimeout{}
}

/*PostTelephonyProvidersEdgeLogsJobsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostTelephonyProvidersEdgeLogsJobsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogsJobsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs][%d] postTelephonyProvidersEdgeLogsJobsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogsJobsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogsJobsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}