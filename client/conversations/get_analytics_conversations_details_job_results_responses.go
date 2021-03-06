// Code generated by go-swagger; DO NOT EDIT.

package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetAnalyticsConversationsDetailsJobResultsReader is a Reader for the GetAnalyticsConversationsDetailsJobResults structure.
type GetAnalyticsConversationsDetailsJobResultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnalyticsConversationsDetailsJobResultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnalyticsConversationsDetailsJobResultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAnalyticsConversationsDetailsJobResultsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAnalyticsConversationsDetailsJobResultsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAnalyticsConversationsDetailsJobResultsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAnalyticsConversationsDetailsJobResultsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAnalyticsConversationsDetailsJobResultsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAnalyticsConversationsDetailsJobResultsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAnalyticsConversationsDetailsJobResultsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAnalyticsConversationsDetailsJobResultsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAnalyticsConversationsDetailsJobResultsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAnalyticsConversationsDetailsJobResultsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAnalyticsConversationsDetailsJobResultsOK creates a GetAnalyticsConversationsDetailsJobResultsOK with default headers values
func NewGetAnalyticsConversationsDetailsJobResultsOK() *GetAnalyticsConversationsDetailsJobResultsOK {
	return &GetAnalyticsConversationsDetailsJobResultsOK{}
}

/*GetAnalyticsConversationsDetailsJobResultsOK handles this case with default header values.

successful operation
*/
type GetAnalyticsConversationsDetailsJobResultsOK struct {
	Payload *models.AnalyticsConversationAsyncQueryResponse
}

func (o *GetAnalyticsConversationsDetailsJobResultsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/{jobId}/results][%d] getAnalyticsConversationsDetailsJobResultsOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobResultsOK) GetPayload() *models.AnalyticsConversationAsyncQueryResponse {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsJobResultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AnalyticsConversationAsyncQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsJobResultsBadRequest creates a GetAnalyticsConversationsDetailsJobResultsBadRequest with default headers values
func NewGetAnalyticsConversationsDetailsJobResultsBadRequest() *GetAnalyticsConversationsDetailsJobResultsBadRequest {
	return &GetAnalyticsConversationsDetailsJobResultsBadRequest{}
}

/*GetAnalyticsConversationsDetailsJobResultsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAnalyticsConversationsDetailsJobResultsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsJobResultsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/{jobId}/results][%d] getAnalyticsConversationsDetailsJobResultsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobResultsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsJobResultsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsJobResultsUnauthorized creates a GetAnalyticsConversationsDetailsJobResultsUnauthorized with default headers values
func NewGetAnalyticsConversationsDetailsJobResultsUnauthorized() *GetAnalyticsConversationsDetailsJobResultsUnauthorized {
	return &GetAnalyticsConversationsDetailsJobResultsUnauthorized{}
}

/*GetAnalyticsConversationsDetailsJobResultsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAnalyticsConversationsDetailsJobResultsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsJobResultsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/{jobId}/results][%d] getAnalyticsConversationsDetailsJobResultsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobResultsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsJobResultsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsJobResultsForbidden creates a GetAnalyticsConversationsDetailsJobResultsForbidden with default headers values
func NewGetAnalyticsConversationsDetailsJobResultsForbidden() *GetAnalyticsConversationsDetailsJobResultsForbidden {
	return &GetAnalyticsConversationsDetailsJobResultsForbidden{}
}

/*GetAnalyticsConversationsDetailsJobResultsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAnalyticsConversationsDetailsJobResultsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsJobResultsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/{jobId}/results][%d] getAnalyticsConversationsDetailsJobResultsForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobResultsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsJobResultsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsJobResultsNotFound creates a GetAnalyticsConversationsDetailsJobResultsNotFound with default headers values
func NewGetAnalyticsConversationsDetailsJobResultsNotFound() *GetAnalyticsConversationsDetailsJobResultsNotFound {
	return &GetAnalyticsConversationsDetailsJobResultsNotFound{}
}

/*GetAnalyticsConversationsDetailsJobResultsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAnalyticsConversationsDetailsJobResultsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsJobResultsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/{jobId}/results][%d] getAnalyticsConversationsDetailsJobResultsNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobResultsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsJobResultsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsJobResultsRequestEntityTooLarge creates a GetAnalyticsConversationsDetailsJobResultsRequestEntityTooLarge with default headers values
func NewGetAnalyticsConversationsDetailsJobResultsRequestEntityTooLarge() *GetAnalyticsConversationsDetailsJobResultsRequestEntityTooLarge {
	return &GetAnalyticsConversationsDetailsJobResultsRequestEntityTooLarge{}
}

/*GetAnalyticsConversationsDetailsJobResultsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetAnalyticsConversationsDetailsJobResultsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsJobResultsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/{jobId}/results][%d] getAnalyticsConversationsDetailsJobResultsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobResultsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsJobResultsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsJobResultsUnsupportedMediaType creates a GetAnalyticsConversationsDetailsJobResultsUnsupportedMediaType with default headers values
func NewGetAnalyticsConversationsDetailsJobResultsUnsupportedMediaType() *GetAnalyticsConversationsDetailsJobResultsUnsupportedMediaType {
	return &GetAnalyticsConversationsDetailsJobResultsUnsupportedMediaType{}
}

/*GetAnalyticsConversationsDetailsJobResultsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAnalyticsConversationsDetailsJobResultsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsJobResultsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/{jobId}/results][%d] getAnalyticsConversationsDetailsJobResultsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobResultsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsJobResultsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsJobResultsTooManyRequests creates a GetAnalyticsConversationsDetailsJobResultsTooManyRequests with default headers values
func NewGetAnalyticsConversationsDetailsJobResultsTooManyRequests() *GetAnalyticsConversationsDetailsJobResultsTooManyRequests {
	return &GetAnalyticsConversationsDetailsJobResultsTooManyRequests{}
}

/*GetAnalyticsConversationsDetailsJobResultsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetAnalyticsConversationsDetailsJobResultsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsJobResultsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/{jobId}/results][%d] getAnalyticsConversationsDetailsJobResultsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobResultsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsJobResultsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsJobResultsInternalServerError creates a GetAnalyticsConversationsDetailsJobResultsInternalServerError with default headers values
func NewGetAnalyticsConversationsDetailsJobResultsInternalServerError() *GetAnalyticsConversationsDetailsJobResultsInternalServerError {
	return &GetAnalyticsConversationsDetailsJobResultsInternalServerError{}
}

/*GetAnalyticsConversationsDetailsJobResultsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAnalyticsConversationsDetailsJobResultsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsJobResultsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/{jobId}/results][%d] getAnalyticsConversationsDetailsJobResultsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobResultsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsJobResultsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsJobResultsServiceUnavailable creates a GetAnalyticsConversationsDetailsJobResultsServiceUnavailable with default headers values
func NewGetAnalyticsConversationsDetailsJobResultsServiceUnavailable() *GetAnalyticsConversationsDetailsJobResultsServiceUnavailable {
	return &GetAnalyticsConversationsDetailsJobResultsServiceUnavailable{}
}

/*GetAnalyticsConversationsDetailsJobResultsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAnalyticsConversationsDetailsJobResultsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsJobResultsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/{jobId}/results][%d] getAnalyticsConversationsDetailsJobResultsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobResultsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsJobResultsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsJobResultsGatewayTimeout creates a GetAnalyticsConversationsDetailsJobResultsGatewayTimeout with default headers values
func NewGetAnalyticsConversationsDetailsJobResultsGatewayTimeout() *GetAnalyticsConversationsDetailsJobResultsGatewayTimeout {
	return &GetAnalyticsConversationsDetailsJobResultsGatewayTimeout{}
}

/*GetAnalyticsConversationsDetailsJobResultsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAnalyticsConversationsDetailsJobResultsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsJobResultsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/{jobId}/results][%d] getAnalyticsConversationsDetailsJobResultsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobResultsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsJobResultsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
