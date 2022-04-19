// Code generated by go-swagger; DO NOT EDIT.

package quality

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetQualityAgentsActivityReader is a Reader for the GetQualityAgentsActivity structure.
type GetQualityAgentsActivityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQualityAgentsActivityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQualityAgentsActivityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetQualityAgentsActivityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetQualityAgentsActivityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetQualityAgentsActivityForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetQualityAgentsActivityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetQualityAgentsActivityRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetQualityAgentsActivityRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetQualityAgentsActivityUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetQualityAgentsActivityTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetQualityAgentsActivityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetQualityAgentsActivityServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetQualityAgentsActivityGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetQualityAgentsActivityOK creates a GetQualityAgentsActivityOK with default headers values
func NewGetQualityAgentsActivityOK() *GetQualityAgentsActivityOK {
	return &GetQualityAgentsActivityOK{}
}

/*GetQualityAgentsActivityOK handles this case with default header values.

successful operation
*/
type GetQualityAgentsActivityOK struct {
	Payload *models.AgentActivityEntityListing
}

func (o *GetQualityAgentsActivityOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/agents/activity][%d] getQualityAgentsActivityOK  %+v", 200, o.Payload)
}

func (o *GetQualityAgentsActivityOK) GetPayload() *models.AgentActivityEntityListing {
	return o.Payload
}

func (o *GetQualityAgentsActivityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AgentActivityEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityAgentsActivityBadRequest creates a GetQualityAgentsActivityBadRequest with default headers values
func NewGetQualityAgentsActivityBadRequest() *GetQualityAgentsActivityBadRequest {
	return &GetQualityAgentsActivityBadRequest{}
}

/*GetQualityAgentsActivityBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetQualityAgentsActivityBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetQualityAgentsActivityBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/agents/activity][%d] getQualityAgentsActivityBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityAgentsActivityBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityAgentsActivityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityAgentsActivityUnauthorized creates a GetQualityAgentsActivityUnauthorized with default headers values
func NewGetQualityAgentsActivityUnauthorized() *GetQualityAgentsActivityUnauthorized {
	return &GetQualityAgentsActivityUnauthorized{}
}

/*GetQualityAgentsActivityUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetQualityAgentsActivityUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetQualityAgentsActivityUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/agents/activity][%d] getQualityAgentsActivityUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityAgentsActivityUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityAgentsActivityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityAgentsActivityForbidden creates a GetQualityAgentsActivityForbidden with default headers values
func NewGetQualityAgentsActivityForbidden() *GetQualityAgentsActivityForbidden {
	return &GetQualityAgentsActivityForbidden{}
}

/*GetQualityAgentsActivityForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetQualityAgentsActivityForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetQualityAgentsActivityForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/agents/activity][%d] getQualityAgentsActivityForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityAgentsActivityForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityAgentsActivityForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityAgentsActivityNotFound creates a GetQualityAgentsActivityNotFound with default headers values
func NewGetQualityAgentsActivityNotFound() *GetQualityAgentsActivityNotFound {
	return &GetQualityAgentsActivityNotFound{}
}

/*GetQualityAgentsActivityNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetQualityAgentsActivityNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetQualityAgentsActivityNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/agents/activity][%d] getQualityAgentsActivityNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityAgentsActivityNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityAgentsActivityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityAgentsActivityRequestTimeout creates a GetQualityAgentsActivityRequestTimeout with default headers values
func NewGetQualityAgentsActivityRequestTimeout() *GetQualityAgentsActivityRequestTimeout {
	return &GetQualityAgentsActivityRequestTimeout{}
}

/*GetQualityAgentsActivityRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetQualityAgentsActivityRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetQualityAgentsActivityRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/agents/activity][%d] getQualityAgentsActivityRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualityAgentsActivityRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityAgentsActivityRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityAgentsActivityRequestEntityTooLarge creates a GetQualityAgentsActivityRequestEntityTooLarge with default headers values
func NewGetQualityAgentsActivityRequestEntityTooLarge() *GetQualityAgentsActivityRequestEntityTooLarge {
	return &GetQualityAgentsActivityRequestEntityTooLarge{}
}

/*GetQualityAgentsActivityRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetQualityAgentsActivityRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetQualityAgentsActivityRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/agents/activity][%d] getQualityAgentsActivityRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityAgentsActivityRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityAgentsActivityRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityAgentsActivityUnsupportedMediaType creates a GetQualityAgentsActivityUnsupportedMediaType with default headers values
func NewGetQualityAgentsActivityUnsupportedMediaType() *GetQualityAgentsActivityUnsupportedMediaType {
	return &GetQualityAgentsActivityUnsupportedMediaType{}
}

/*GetQualityAgentsActivityUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetQualityAgentsActivityUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetQualityAgentsActivityUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/agents/activity][%d] getQualityAgentsActivityUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityAgentsActivityUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityAgentsActivityUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityAgentsActivityTooManyRequests creates a GetQualityAgentsActivityTooManyRequests with default headers values
func NewGetQualityAgentsActivityTooManyRequests() *GetQualityAgentsActivityTooManyRequests {
	return &GetQualityAgentsActivityTooManyRequests{}
}

/*GetQualityAgentsActivityTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetQualityAgentsActivityTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetQualityAgentsActivityTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/agents/activity][%d] getQualityAgentsActivityTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityAgentsActivityTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityAgentsActivityTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityAgentsActivityInternalServerError creates a GetQualityAgentsActivityInternalServerError with default headers values
func NewGetQualityAgentsActivityInternalServerError() *GetQualityAgentsActivityInternalServerError {
	return &GetQualityAgentsActivityInternalServerError{}
}

/*GetQualityAgentsActivityInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetQualityAgentsActivityInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetQualityAgentsActivityInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/agents/activity][%d] getQualityAgentsActivityInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityAgentsActivityInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityAgentsActivityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityAgentsActivityServiceUnavailable creates a GetQualityAgentsActivityServiceUnavailable with default headers values
func NewGetQualityAgentsActivityServiceUnavailable() *GetQualityAgentsActivityServiceUnavailable {
	return &GetQualityAgentsActivityServiceUnavailable{}
}

/*GetQualityAgentsActivityServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetQualityAgentsActivityServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetQualityAgentsActivityServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/agents/activity][%d] getQualityAgentsActivityServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityAgentsActivityServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityAgentsActivityServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityAgentsActivityGatewayTimeout creates a GetQualityAgentsActivityGatewayTimeout with default headers values
func NewGetQualityAgentsActivityGatewayTimeout() *GetQualityAgentsActivityGatewayTimeout {
	return &GetQualityAgentsActivityGatewayTimeout{}
}

/*GetQualityAgentsActivityGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetQualityAgentsActivityGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetQualityAgentsActivityGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/agents/activity][%d] getQualityAgentsActivityGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityAgentsActivityGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityAgentsActivityGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
