// Code generated by go-swagger; DO NOT EDIT.

package widgets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostWidgetsDeploymentsReader is a Reader for the PostWidgetsDeployments structure.
type PostWidgetsDeploymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWidgetsDeploymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWidgetsDeploymentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWidgetsDeploymentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWidgetsDeploymentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWidgetsDeploymentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWidgetsDeploymentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostWidgetsDeploymentsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostWidgetsDeploymentsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWidgetsDeploymentsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWidgetsDeploymentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWidgetsDeploymentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWidgetsDeploymentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWidgetsDeploymentsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWidgetsDeploymentsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWidgetsDeploymentsOK creates a PostWidgetsDeploymentsOK with default headers values
func NewPostWidgetsDeploymentsOK() *PostWidgetsDeploymentsOK {
	return &PostWidgetsDeploymentsOK{}
}

/*PostWidgetsDeploymentsOK handles this case with default header values.

successful operation
*/
type PostWidgetsDeploymentsOK struct {
	Payload *models.WidgetDeployment
}

func (o *PostWidgetsDeploymentsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsOK  %+v", 200, o.Payload)
}

func (o *PostWidgetsDeploymentsOK) GetPayload() *models.WidgetDeployment {
	return o.Payload
}

func (o *PostWidgetsDeploymentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WidgetDeployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWidgetsDeploymentsBadRequest creates a PostWidgetsDeploymentsBadRequest with default headers values
func NewPostWidgetsDeploymentsBadRequest() *PostWidgetsDeploymentsBadRequest {
	return &PostWidgetsDeploymentsBadRequest{}
}

/*PostWidgetsDeploymentsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWidgetsDeploymentsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostWidgetsDeploymentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsBadRequest  %+v", 400, o.Payload)
}

func (o *PostWidgetsDeploymentsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWidgetsDeploymentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWidgetsDeploymentsUnauthorized creates a PostWidgetsDeploymentsUnauthorized with default headers values
func NewPostWidgetsDeploymentsUnauthorized() *PostWidgetsDeploymentsUnauthorized {
	return &PostWidgetsDeploymentsUnauthorized{}
}

/*PostWidgetsDeploymentsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWidgetsDeploymentsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostWidgetsDeploymentsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWidgetsDeploymentsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWidgetsDeploymentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWidgetsDeploymentsForbidden creates a PostWidgetsDeploymentsForbidden with default headers values
func NewPostWidgetsDeploymentsForbidden() *PostWidgetsDeploymentsForbidden {
	return &PostWidgetsDeploymentsForbidden{}
}

/*PostWidgetsDeploymentsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostWidgetsDeploymentsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostWidgetsDeploymentsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsForbidden  %+v", 403, o.Payload)
}

func (o *PostWidgetsDeploymentsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWidgetsDeploymentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWidgetsDeploymentsNotFound creates a PostWidgetsDeploymentsNotFound with default headers values
func NewPostWidgetsDeploymentsNotFound() *PostWidgetsDeploymentsNotFound {
	return &PostWidgetsDeploymentsNotFound{}
}

/*PostWidgetsDeploymentsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostWidgetsDeploymentsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostWidgetsDeploymentsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsNotFound  %+v", 404, o.Payload)
}

func (o *PostWidgetsDeploymentsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWidgetsDeploymentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWidgetsDeploymentsRequestTimeout creates a PostWidgetsDeploymentsRequestTimeout with default headers values
func NewPostWidgetsDeploymentsRequestTimeout() *PostWidgetsDeploymentsRequestTimeout {
	return &PostWidgetsDeploymentsRequestTimeout{}
}

/*PostWidgetsDeploymentsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostWidgetsDeploymentsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWidgetsDeploymentsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWidgetsDeploymentsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWidgetsDeploymentsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWidgetsDeploymentsConflict creates a PostWidgetsDeploymentsConflict with default headers values
func NewPostWidgetsDeploymentsConflict() *PostWidgetsDeploymentsConflict {
	return &PostWidgetsDeploymentsConflict{}
}

/*PostWidgetsDeploymentsConflict handles this case with default header values.

Conflict
*/
type PostWidgetsDeploymentsConflict struct {
	Payload *models.ErrorBody
}

func (o *PostWidgetsDeploymentsConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsConflict  %+v", 409, o.Payload)
}

func (o *PostWidgetsDeploymentsConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWidgetsDeploymentsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWidgetsDeploymentsRequestEntityTooLarge creates a PostWidgetsDeploymentsRequestEntityTooLarge with default headers values
func NewPostWidgetsDeploymentsRequestEntityTooLarge() *PostWidgetsDeploymentsRequestEntityTooLarge {
	return &PostWidgetsDeploymentsRequestEntityTooLarge{}
}

/*PostWidgetsDeploymentsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostWidgetsDeploymentsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostWidgetsDeploymentsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWidgetsDeploymentsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWidgetsDeploymentsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWidgetsDeploymentsUnsupportedMediaType creates a PostWidgetsDeploymentsUnsupportedMediaType with default headers values
func NewPostWidgetsDeploymentsUnsupportedMediaType() *PostWidgetsDeploymentsUnsupportedMediaType {
	return &PostWidgetsDeploymentsUnsupportedMediaType{}
}

/*PostWidgetsDeploymentsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWidgetsDeploymentsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostWidgetsDeploymentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWidgetsDeploymentsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWidgetsDeploymentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWidgetsDeploymentsTooManyRequests creates a PostWidgetsDeploymentsTooManyRequests with default headers values
func NewPostWidgetsDeploymentsTooManyRequests() *PostWidgetsDeploymentsTooManyRequests {
	return &PostWidgetsDeploymentsTooManyRequests{}
}

/*PostWidgetsDeploymentsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostWidgetsDeploymentsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostWidgetsDeploymentsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWidgetsDeploymentsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWidgetsDeploymentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWidgetsDeploymentsInternalServerError creates a PostWidgetsDeploymentsInternalServerError with default headers values
func NewPostWidgetsDeploymentsInternalServerError() *PostWidgetsDeploymentsInternalServerError {
	return &PostWidgetsDeploymentsInternalServerError{}
}

/*PostWidgetsDeploymentsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWidgetsDeploymentsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostWidgetsDeploymentsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWidgetsDeploymentsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWidgetsDeploymentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWidgetsDeploymentsServiceUnavailable creates a PostWidgetsDeploymentsServiceUnavailable with default headers values
func NewPostWidgetsDeploymentsServiceUnavailable() *PostWidgetsDeploymentsServiceUnavailable {
	return &PostWidgetsDeploymentsServiceUnavailable{}
}

/*PostWidgetsDeploymentsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWidgetsDeploymentsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostWidgetsDeploymentsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWidgetsDeploymentsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWidgetsDeploymentsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWidgetsDeploymentsGatewayTimeout creates a PostWidgetsDeploymentsGatewayTimeout with default headers values
func NewPostWidgetsDeploymentsGatewayTimeout() *PostWidgetsDeploymentsGatewayTimeout {
	return &PostWidgetsDeploymentsGatewayTimeout{}
}

/*PostWidgetsDeploymentsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostWidgetsDeploymentsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWidgetsDeploymentsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWidgetsDeploymentsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWidgetsDeploymentsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
