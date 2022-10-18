// Code generated by go-swagger; DO NOT EDIT.

package journey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostJourneyActionmapsEstimatesJobsReader is a Reader for the PostJourneyActionmapsEstimatesJobs structure.
type PostJourneyActionmapsEstimatesJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostJourneyActionmapsEstimatesJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostJourneyActionmapsEstimatesJobsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostJourneyActionmapsEstimatesJobsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostJourneyActionmapsEstimatesJobsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostJourneyActionmapsEstimatesJobsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostJourneyActionmapsEstimatesJobsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostJourneyActionmapsEstimatesJobsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostJourneyActionmapsEstimatesJobsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostJourneyActionmapsEstimatesJobsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostJourneyActionmapsEstimatesJobsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostJourneyActionmapsEstimatesJobsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostJourneyActionmapsEstimatesJobsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostJourneyActionmapsEstimatesJobsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostJourneyActionmapsEstimatesJobsAccepted creates a PostJourneyActionmapsEstimatesJobsAccepted with default headers values
func NewPostJourneyActionmapsEstimatesJobsAccepted() *PostJourneyActionmapsEstimatesJobsAccepted {
	return &PostJourneyActionmapsEstimatesJobsAccepted{}
}

/*PostJourneyActionmapsEstimatesJobsAccepted handles this case with default header values.

Accepted - Running query asynchronously
*/
type PostJourneyActionmapsEstimatesJobsAccepted struct {
	Payload *models.EstimateJobAsyncResponse
}

func (o *PostJourneyActionmapsEstimatesJobsAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actionmaps/estimates/jobs][%d] postJourneyActionmapsEstimatesJobsAccepted  %+v", 202, o.Payload)
}

func (o *PostJourneyActionmapsEstimatesJobsAccepted) GetPayload() *models.EstimateJobAsyncResponse {
	return o.Payload
}

func (o *PostJourneyActionmapsEstimatesJobsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EstimateJobAsyncResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActionmapsEstimatesJobsBadRequest creates a PostJourneyActionmapsEstimatesJobsBadRequest with default headers values
func NewPostJourneyActionmapsEstimatesJobsBadRequest() *PostJourneyActionmapsEstimatesJobsBadRequest {
	return &PostJourneyActionmapsEstimatesJobsBadRequest{}
}

/*PostJourneyActionmapsEstimatesJobsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostJourneyActionmapsEstimatesJobsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyActionmapsEstimatesJobsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actionmaps/estimates/jobs][%d] postJourneyActionmapsEstimatesJobsBadRequest  %+v", 400, o.Payload)
}

func (o *PostJourneyActionmapsEstimatesJobsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyActionmapsEstimatesJobsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActionmapsEstimatesJobsUnauthorized creates a PostJourneyActionmapsEstimatesJobsUnauthorized with default headers values
func NewPostJourneyActionmapsEstimatesJobsUnauthorized() *PostJourneyActionmapsEstimatesJobsUnauthorized {
	return &PostJourneyActionmapsEstimatesJobsUnauthorized{}
}

/*PostJourneyActionmapsEstimatesJobsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostJourneyActionmapsEstimatesJobsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyActionmapsEstimatesJobsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actionmaps/estimates/jobs][%d] postJourneyActionmapsEstimatesJobsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostJourneyActionmapsEstimatesJobsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyActionmapsEstimatesJobsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActionmapsEstimatesJobsForbidden creates a PostJourneyActionmapsEstimatesJobsForbidden with default headers values
func NewPostJourneyActionmapsEstimatesJobsForbidden() *PostJourneyActionmapsEstimatesJobsForbidden {
	return &PostJourneyActionmapsEstimatesJobsForbidden{}
}

/*PostJourneyActionmapsEstimatesJobsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostJourneyActionmapsEstimatesJobsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyActionmapsEstimatesJobsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actionmaps/estimates/jobs][%d] postJourneyActionmapsEstimatesJobsForbidden  %+v", 403, o.Payload)
}

func (o *PostJourneyActionmapsEstimatesJobsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyActionmapsEstimatesJobsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActionmapsEstimatesJobsNotFound creates a PostJourneyActionmapsEstimatesJobsNotFound with default headers values
func NewPostJourneyActionmapsEstimatesJobsNotFound() *PostJourneyActionmapsEstimatesJobsNotFound {
	return &PostJourneyActionmapsEstimatesJobsNotFound{}
}

/*PostJourneyActionmapsEstimatesJobsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostJourneyActionmapsEstimatesJobsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyActionmapsEstimatesJobsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actionmaps/estimates/jobs][%d] postJourneyActionmapsEstimatesJobsNotFound  %+v", 404, o.Payload)
}

func (o *PostJourneyActionmapsEstimatesJobsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyActionmapsEstimatesJobsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActionmapsEstimatesJobsRequestTimeout creates a PostJourneyActionmapsEstimatesJobsRequestTimeout with default headers values
func NewPostJourneyActionmapsEstimatesJobsRequestTimeout() *PostJourneyActionmapsEstimatesJobsRequestTimeout {
	return &PostJourneyActionmapsEstimatesJobsRequestTimeout{}
}

/*PostJourneyActionmapsEstimatesJobsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostJourneyActionmapsEstimatesJobsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyActionmapsEstimatesJobsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actionmaps/estimates/jobs][%d] postJourneyActionmapsEstimatesJobsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostJourneyActionmapsEstimatesJobsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyActionmapsEstimatesJobsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActionmapsEstimatesJobsRequestEntityTooLarge creates a PostJourneyActionmapsEstimatesJobsRequestEntityTooLarge with default headers values
func NewPostJourneyActionmapsEstimatesJobsRequestEntityTooLarge() *PostJourneyActionmapsEstimatesJobsRequestEntityTooLarge {
	return &PostJourneyActionmapsEstimatesJobsRequestEntityTooLarge{}
}

/*PostJourneyActionmapsEstimatesJobsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostJourneyActionmapsEstimatesJobsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyActionmapsEstimatesJobsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actionmaps/estimates/jobs][%d] postJourneyActionmapsEstimatesJobsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostJourneyActionmapsEstimatesJobsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyActionmapsEstimatesJobsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActionmapsEstimatesJobsUnsupportedMediaType creates a PostJourneyActionmapsEstimatesJobsUnsupportedMediaType with default headers values
func NewPostJourneyActionmapsEstimatesJobsUnsupportedMediaType() *PostJourneyActionmapsEstimatesJobsUnsupportedMediaType {
	return &PostJourneyActionmapsEstimatesJobsUnsupportedMediaType{}
}

/*PostJourneyActionmapsEstimatesJobsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostJourneyActionmapsEstimatesJobsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyActionmapsEstimatesJobsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actionmaps/estimates/jobs][%d] postJourneyActionmapsEstimatesJobsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostJourneyActionmapsEstimatesJobsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyActionmapsEstimatesJobsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActionmapsEstimatesJobsTooManyRequests creates a PostJourneyActionmapsEstimatesJobsTooManyRequests with default headers values
func NewPostJourneyActionmapsEstimatesJobsTooManyRequests() *PostJourneyActionmapsEstimatesJobsTooManyRequests {
	return &PostJourneyActionmapsEstimatesJobsTooManyRequests{}
}

/*PostJourneyActionmapsEstimatesJobsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostJourneyActionmapsEstimatesJobsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyActionmapsEstimatesJobsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actionmaps/estimates/jobs][%d] postJourneyActionmapsEstimatesJobsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostJourneyActionmapsEstimatesJobsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyActionmapsEstimatesJobsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActionmapsEstimatesJobsInternalServerError creates a PostJourneyActionmapsEstimatesJobsInternalServerError with default headers values
func NewPostJourneyActionmapsEstimatesJobsInternalServerError() *PostJourneyActionmapsEstimatesJobsInternalServerError {
	return &PostJourneyActionmapsEstimatesJobsInternalServerError{}
}

/*PostJourneyActionmapsEstimatesJobsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostJourneyActionmapsEstimatesJobsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyActionmapsEstimatesJobsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actionmaps/estimates/jobs][%d] postJourneyActionmapsEstimatesJobsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostJourneyActionmapsEstimatesJobsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyActionmapsEstimatesJobsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActionmapsEstimatesJobsServiceUnavailable creates a PostJourneyActionmapsEstimatesJobsServiceUnavailable with default headers values
func NewPostJourneyActionmapsEstimatesJobsServiceUnavailable() *PostJourneyActionmapsEstimatesJobsServiceUnavailable {
	return &PostJourneyActionmapsEstimatesJobsServiceUnavailable{}
}

/*PostJourneyActionmapsEstimatesJobsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostJourneyActionmapsEstimatesJobsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyActionmapsEstimatesJobsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actionmaps/estimates/jobs][%d] postJourneyActionmapsEstimatesJobsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostJourneyActionmapsEstimatesJobsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyActionmapsEstimatesJobsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyActionmapsEstimatesJobsGatewayTimeout creates a PostJourneyActionmapsEstimatesJobsGatewayTimeout with default headers values
func NewPostJourneyActionmapsEstimatesJobsGatewayTimeout() *PostJourneyActionmapsEstimatesJobsGatewayTimeout {
	return &PostJourneyActionmapsEstimatesJobsGatewayTimeout{}
}

/*PostJourneyActionmapsEstimatesJobsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostJourneyActionmapsEstimatesJobsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyActionmapsEstimatesJobsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/actionmaps/estimates/jobs][%d] postJourneyActionmapsEstimatesJobsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostJourneyActionmapsEstimatesJobsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyActionmapsEstimatesJobsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
