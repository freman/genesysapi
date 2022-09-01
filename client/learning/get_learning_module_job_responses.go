// Code generated by go-swagger; DO NOT EDIT.

package learning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetLearningModuleJobReader is a Reader for the GetLearningModuleJob structure.
type GetLearningModuleJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearningModuleJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearningModuleJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLearningModuleJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLearningModuleJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLearningModuleJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLearningModuleJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetLearningModuleJobRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLearningModuleJobRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLearningModuleJobUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLearningModuleJobTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLearningModuleJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLearningModuleJobServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLearningModuleJobGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearningModuleJobOK creates a GetLearningModuleJobOK with default headers values
func NewGetLearningModuleJobOK() *GetLearningModuleJobOK {
	return &GetLearningModuleJobOK{}
}

/*GetLearningModuleJobOK handles this case with default header values.

successful operation
*/
type GetLearningModuleJobOK struct {
	Payload *models.LearningModuleJobResponse
}

func (o *GetLearningModuleJobOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/jobs/{jobId}][%d] getLearningModuleJobOK  %+v", 200, o.Payload)
}

func (o *GetLearningModuleJobOK) GetPayload() *models.LearningModuleJobResponse {
	return o.Payload
}

func (o *GetLearningModuleJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LearningModuleJobResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningModuleJobBadRequest creates a GetLearningModuleJobBadRequest with default headers values
func NewGetLearningModuleJobBadRequest() *GetLearningModuleJobBadRequest {
	return &GetLearningModuleJobBadRequest{}
}

/*GetLearningModuleJobBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLearningModuleJobBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetLearningModuleJobBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/jobs/{jobId}][%d] getLearningModuleJobBadRequest  %+v", 400, o.Payload)
}

func (o *GetLearningModuleJobBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningModuleJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningModuleJobUnauthorized creates a GetLearningModuleJobUnauthorized with default headers values
func NewGetLearningModuleJobUnauthorized() *GetLearningModuleJobUnauthorized {
	return &GetLearningModuleJobUnauthorized{}
}

/*GetLearningModuleJobUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLearningModuleJobUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetLearningModuleJobUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/jobs/{jobId}][%d] getLearningModuleJobUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLearningModuleJobUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningModuleJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningModuleJobForbidden creates a GetLearningModuleJobForbidden with default headers values
func NewGetLearningModuleJobForbidden() *GetLearningModuleJobForbidden {
	return &GetLearningModuleJobForbidden{}
}

/*GetLearningModuleJobForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetLearningModuleJobForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetLearningModuleJobForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/jobs/{jobId}][%d] getLearningModuleJobForbidden  %+v", 403, o.Payload)
}

func (o *GetLearningModuleJobForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningModuleJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningModuleJobNotFound creates a GetLearningModuleJobNotFound with default headers values
func NewGetLearningModuleJobNotFound() *GetLearningModuleJobNotFound {
	return &GetLearningModuleJobNotFound{}
}

/*GetLearningModuleJobNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetLearningModuleJobNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetLearningModuleJobNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/jobs/{jobId}][%d] getLearningModuleJobNotFound  %+v", 404, o.Payload)
}

func (o *GetLearningModuleJobNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningModuleJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningModuleJobRequestTimeout creates a GetLearningModuleJobRequestTimeout with default headers values
func NewGetLearningModuleJobRequestTimeout() *GetLearningModuleJobRequestTimeout {
	return &GetLearningModuleJobRequestTimeout{}
}

/*GetLearningModuleJobRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetLearningModuleJobRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetLearningModuleJobRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/jobs/{jobId}][%d] getLearningModuleJobRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLearningModuleJobRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningModuleJobRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningModuleJobRequestEntityTooLarge creates a GetLearningModuleJobRequestEntityTooLarge with default headers values
func NewGetLearningModuleJobRequestEntityTooLarge() *GetLearningModuleJobRequestEntityTooLarge {
	return &GetLearningModuleJobRequestEntityTooLarge{}
}

/*GetLearningModuleJobRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetLearningModuleJobRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetLearningModuleJobRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/jobs/{jobId}][%d] getLearningModuleJobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLearningModuleJobRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningModuleJobRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningModuleJobUnsupportedMediaType creates a GetLearningModuleJobUnsupportedMediaType with default headers values
func NewGetLearningModuleJobUnsupportedMediaType() *GetLearningModuleJobUnsupportedMediaType {
	return &GetLearningModuleJobUnsupportedMediaType{}
}

/*GetLearningModuleJobUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLearningModuleJobUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetLearningModuleJobUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/jobs/{jobId}][%d] getLearningModuleJobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLearningModuleJobUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningModuleJobUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningModuleJobTooManyRequests creates a GetLearningModuleJobTooManyRequests with default headers values
func NewGetLearningModuleJobTooManyRequests() *GetLearningModuleJobTooManyRequests {
	return &GetLearningModuleJobTooManyRequests{}
}

/*GetLearningModuleJobTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetLearningModuleJobTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetLearningModuleJobTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/jobs/{jobId}][%d] getLearningModuleJobTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLearningModuleJobTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningModuleJobTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningModuleJobInternalServerError creates a GetLearningModuleJobInternalServerError with default headers values
func NewGetLearningModuleJobInternalServerError() *GetLearningModuleJobInternalServerError {
	return &GetLearningModuleJobInternalServerError{}
}

/*GetLearningModuleJobInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLearningModuleJobInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetLearningModuleJobInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/jobs/{jobId}][%d] getLearningModuleJobInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLearningModuleJobInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningModuleJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningModuleJobServiceUnavailable creates a GetLearningModuleJobServiceUnavailable with default headers values
func NewGetLearningModuleJobServiceUnavailable() *GetLearningModuleJobServiceUnavailable {
	return &GetLearningModuleJobServiceUnavailable{}
}

/*GetLearningModuleJobServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLearningModuleJobServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetLearningModuleJobServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/jobs/{jobId}][%d] getLearningModuleJobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLearningModuleJobServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningModuleJobServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLearningModuleJobGatewayTimeout creates a GetLearningModuleJobGatewayTimeout with default headers values
func NewGetLearningModuleJobGatewayTimeout() *GetLearningModuleJobGatewayTimeout {
	return &GetLearningModuleJobGatewayTimeout{}
}

/*GetLearningModuleJobGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetLearningModuleJobGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetLearningModuleJobGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/learning/modules/{moduleId}/jobs/{jobId}][%d] getLearningModuleJobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLearningModuleJobGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLearningModuleJobGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}