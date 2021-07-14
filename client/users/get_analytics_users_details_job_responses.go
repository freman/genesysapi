// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetAnalyticsUsersDetailsJobReader is a Reader for the GetAnalyticsUsersDetailsJob structure.
type GetAnalyticsUsersDetailsJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnalyticsUsersDetailsJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnalyticsUsersDetailsJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewGetAnalyticsUsersDetailsJobAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAnalyticsUsersDetailsJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAnalyticsUsersDetailsJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAnalyticsUsersDetailsJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAnalyticsUsersDetailsJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAnalyticsUsersDetailsJobRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAnalyticsUsersDetailsJobRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAnalyticsUsersDetailsJobUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAnalyticsUsersDetailsJobTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAnalyticsUsersDetailsJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAnalyticsUsersDetailsJobServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAnalyticsUsersDetailsJobGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAnalyticsUsersDetailsJobOK creates a GetAnalyticsUsersDetailsJobOK with default headers values
func NewGetAnalyticsUsersDetailsJobOK() *GetAnalyticsUsersDetailsJobOK {
	return &GetAnalyticsUsersDetailsJobOK{}
}

/*GetAnalyticsUsersDetailsJobOK handles this case with default header values.

successful operation
*/
type GetAnalyticsUsersDetailsJobOK struct {
	Payload *models.AsyncQueryStatus
}

func (o *GetAnalyticsUsersDetailsJobOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}][%d] getAnalyticsUsersDetailsJobOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobOK) GetPayload() *models.AsyncQueryStatus {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AsyncQueryStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobAccepted creates a GetAnalyticsUsersDetailsJobAccepted with default headers values
func NewGetAnalyticsUsersDetailsJobAccepted() *GetAnalyticsUsersDetailsJobAccepted {
	return &GetAnalyticsUsersDetailsJobAccepted{}
}

/*GetAnalyticsUsersDetailsJobAccepted handles this case with default header values.

Accepted - Running query asynchronously
*/
type GetAnalyticsUsersDetailsJobAccepted struct {
	Payload *models.AsyncQueryStatus
}

func (o *GetAnalyticsUsersDetailsJobAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}][%d] getAnalyticsUsersDetailsJobAccepted  %+v", 202, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobAccepted) GetPayload() *models.AsyncQueryStatus {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AsyncQueryStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobBadRequest creates a GetAnalyticsUsersDetailsJobBadRequest with default headers values
func NewGetAnalyticsUsersDetailsJobBadRequest() *GetAnalyticsUsersDetailsJobBadRequest {
	return &GetAnalyticsUsersDetailsJobBadRequest{}
}

/*GetAnalyticsUsersDetailsJobBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAnalyticsUsersDetailsJobBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsUsersDetailsJobBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}][%d] getAnalyticsUsersDetailsJobBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobUnauthorized creates a GetAnalyticsUsersDetailsJobUnauthorized with default headers values
func NewGetAnalyticsUsersDetailsJobUnauthorized() *GetAnalyticsUsersDetailsJobUnauthorized {
	return &GetAnalyticsUsersDetailsJobUnauthorized{}
}

/*GetAnalyticsUsersDetailsJobUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAnalyticsUsersDetailsJobUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsUsersDetailsJobUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}][%d] getAnalyticsUsersDetailsJobUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobForbidden creates a GetAnalyticsUsersDetailsJobForbidden with default headers values
func NewGetAnalyticsUsersDetailsJobForbidden() *GetAnalyticsUsersDetailsJobForbidden {
	return &GetAnalyticsUsersDetailsJobForbidden{}
}

/*GetAnalyticsUsersDetailsJobForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAnalyticsUsersDetailsJobForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsUsersDetailsJobForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}][%d] getAnalyticsUsersDetailsJobForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobNotFound creates a GetAnalyticsUsersDetailsJobNotFound with default headers values
func NewGetAnalyticsUsersDetailsJobNotFound() *GetAnalyticsUsersDetailsJobNotFound {
	return &GetAnalyticsUsersDetailsJobNotFound{}
}

/*GetAnalyticsUsersDetailsJobNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAnalyticsUsersDetailsJobNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsUsersDetailsJobNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}][%d] getAnalyticsUsersDetailsJobNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobRequestTimeout creates a GetAnalyticsUsersDetailsJobRequestTimeout with default headers values
func NewGetAnalyticsUsersDetailsJobRequestTimeout() *GetAnalyticsUsersDetailsJobRequestTimeout {
	return &GetAnalyticsUsersDetailsJobRequestTimeout{}
}

/*GetAnalyticsUsersDetailsJobRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAnalyticsUsersDetailsJobRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsUsersDetailsJobRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}][%d] getAnalyticsUsersDetailsJobRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobRequestEntityTooLarge creates a GetAnalyticsUsersDetailsJobRequestEntityTooLarge with default headers values
func NewGetAnalyticsUsersDetailsJobRequestEntityTooLarge() *GetAnalyticsUsersDetailsJobRequestEntityTooLarge {
	return &GetAnalyticsUsersDetailsJobRequestEntityTooLarge{}
}

/*GetAnalyticsUsersDetailsJobRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetAnalyticsUsersDetailsJobRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsUsersDetailsJobRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}][%d] getAnalyticsUsersDetailsJobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobUnsupportedMediaType creates a GetAnalyticsUsersDetailsJobUnsupportedMediaType with default headers values
func NewGetAnalyticsUsersDetailsJobUnsupportedMediaType() *GetAnalyticsUsersDetailsJobUnsupportedMediaType {
	return &GetAnalyticsUsersDetailsJobUnsupportedMediaType{}
}

/*GetAnalyticsUsersDetailsJobUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAnalyticsUsersDetailsJobUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsUsersDetailsJobUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}][%d] getAnalyticsUsersDetailsJobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobTooManyRequests creates a GetAnalyticsUsersDetailsJobTooManyRequests with default headers values
func NewGetAnalyticsUsersDetailsJobTooManyRequests() *GetAnalyticsUsersDetailsJobTooManyRequests {
	return &GetAnalyticsUsersDetailsJobTooManyRequests{}
}

/*GetAnalyticsUsersDetailsJobTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAnalyticsUsersDetailsJobTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsUsersDetailsJobTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}][%d] getAnalyticsUsersDetailsJobTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobInternalServerError creates a GetAnalyticsUsersDetailsJobInternalServerError with default headers values
func NewGetAnalyticsUsersDetailsJobInternalServerError() *GetAnalyticsUsersDetailsJobInternalServerError {
	return &GetAnalyticsUsersDetailsJobInternalServerError{}
}

/*GetAnalyticsUsersDetailsJobInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAnalyticsUsersDetailsJobInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsUsersDetailsJobInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}][%d] getAnalyticsUsersDetailsJobInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobServiceUnavailable creates a GetAnalyticsUsersDetailsJobServiceUnavailable with default headers values
func NewGetAnalyticsUsersDetailsJobServiceUnavailable() *GetAnalyticsUsersDetailsJobServiceUnavailable {
	return &GetAnalyticsUsersDetailsJobServiceUnavailable{}
}

/*GetAnalyticsUsersDetailsJobServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAnalyticsUsersDetailsJobServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsUsersDetailsJobServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}][%d] getAnalyticsUsersDetailsJobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobGatewayTimeout creates a GetAnalyticsUsersDetailsJobGatewayTimeout with default headers values
func NewGetAnalyticsUsersDetailsJobGatewayTimeout() *GetAnalyticsUsersDetailsJobGatewayTimeout {
	return &GetAnalyticsUsersDetailsJobGatewayTimeout{}
}

/*GetAnalyticsUsersDetailsJobGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAnalyticsUsersDetailsJobGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsUsersDetailsJobGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/{jobId}][%d] getAnalyticsUsersDetailsJobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
