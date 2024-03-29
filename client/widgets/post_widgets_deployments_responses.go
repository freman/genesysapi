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

/*
PostWidgetsDeploymentsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostWidgetsDeploymentsOK struct {
	Payload *models.WidgetDeployment
}

// IsSuccess returns true when this post widgets deployments o k response has a 2xx status code
func (o *PostWidgetsDeploymentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post widgets deployments o k response has a 3xx status code
func (o *PostWidgetsDeploymentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post widgets deployments o k response has a 4xx status code
func (o *PostWidgetsDeploymentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post widgets deployments o k response has a 5xx status code
func (o *PostWidgetsDeploymentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post widgets deployments o k response a status code equal to that given
func (o *PostWidgetsDeploymentsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostWidgetsDeploymentsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsOK  %+v", 200, o.Payload)
}

func (o *PostWidgetsDeploymentsOK) String() string {
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

/*
PostWidgetsDeploymentsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWidgetsDeploymentsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post widgets deployments bad request response has a 2xx status code
func (o *PostWidgetsDeploymentsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post widgets deployments bad request response has a 3xx status code
func (o *PostWidgetsDeploymentsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post widgets deployments bad request response has a 4xx status code
func (o *PostWidgetsDeploymentsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post widgets deployments bad request response has a 5xx status code
func (o *PostWidgetsDeploymentsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post widgets deployments bad request response a status code equal to that given
func (o *PostWidgetsDeploymentsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostWidgetsDeploymentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsBadRequest  %+v", 400, o.Payload)
}

func (o *PostWidgetsDeploymentsBadRequest) String() string {
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

/*
PostWidgetsDeploymentsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWidgetsDeploymentsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post widgets deployments unauthorized response has a 2xx status code
func (o *PostWidgetsDeploymentsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post widgets deployments unauthorized response has a 3xx status code
func (o *PostWidgetsDeploymentsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post widgets deployments unauthorized response has a 4xx status code
func (o *PostWidgetsDeploymentsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post widgets deployments unauthorized response has a 5xx status code
func (o *PostWidgetsDeploymentsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post widgets deployments unauthorized response a status code equal to that given
func (o *PostWidgetsDeploymentsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostWidgetsDeploymentsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWidgetsDeploymentsUnauthorized) String() string {
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

/*
PostWidgetsDeploymentsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostWidgetsDeploymentsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post widgets deployments forbidden response has a 2xx status code
func (o *PostWidgetsDeploymentsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post widgets deployments forbidden response has a 3xx status code
func (o *PostWidgetsDeploymentsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post widgets deployments forbidden response has a 4xx status code
func (o *PostWidgetsDeploymentsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post widgets deployments forbidden response has a 5xx status code
func (o *PostWidgetsDeploymentsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post widgets deployments forbidden response a status code equal to that given
func (o *PostWidgetsDeploymentsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostWidgetsDeploymentsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsForbidden  %+v", 403, o.Payload)
}

func (o *PostWidgetsDeploymentsForbidden) String() string {
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

/*
PostWidgetsDeploymentsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostWidgetsDeploymentsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post widgets deployments not found response has a 2xx status code
func (o *PostWidgetsDeploymentsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post widgets deployments not found response has a 3xx status code
func (o *PostWidgetsDeploymentsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post widgets deployments not found response has a 4xx status code
func (o *PostWidgetsDeploymentsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post widgets deployments not found response has a 5xx status code
func (o *PostWidgetsDeploymentsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post widgets deployments not found response a status code equal to that given
func (o *PostWidgetsDeploymentsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostWidgetsDeploymentsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsNotFound  %+v", 404, o.Payload)
}

func (o *PostWidgetsDeploymentsNotFound) String() string {
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

/*
PostWidgetsDeploymentsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostWidgetsDeploymentsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post widgets deployments request timeout response has a 2xx status code
func (o *PostWidgetsDeploymentsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post widgets deployments request timeout response has a 3xx status code
func (o *PostWidgetsDeploymentsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post widgets deployments request timeout response has a 4xx status code
func (o *PostWidgetsDeploymentsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post widgets deployments request timeout response has a 5xx status code
func (o *PostWidgetsDeploymentsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post widgets deployments request timeout response a status code equal to that given
func (o *PostWidgetsDeploymentsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostWidgetsDeploymentsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWidgetsDeploymentsRequestTimeout) String() string {
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

/*
PostWidgetsDeploymentsConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostWidgetsDeploymentsConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post widgets deployments conflict response has a 2xx status code
func (o *PostWidgetsDeploymentsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post widgets deployments conflict response has a 3xx status code
func (o *PostWidgetsDeploymentsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post widgets deployments conflict response has a 4xx status code
func (o *PostWidgetsDeploymentsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post widgets deployments conflict response has a 5xx status code
func (o *PostWidgetsDeploymentsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post widgets deployments conflict response a status code equal to that given
func (o *PostWidgetsDeploymentsConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PostWidgetsDeploymentsConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsConflict  %+v", 409, o.Payload)
}

func (o *PostWidgetsDeploymentsConflict) String() string {
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

/*
PostWidgetsDeploymentsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostWidgetsDeploymentsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post widgets deployments request entity too large response has a 2xx status code
func (o *PostWidgetsDeploymentsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post widgets deployments request entity too large response has a 3xx status code
func (o *PostWidgetsDeploymentsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post widgets deployments request entity too large response has a 4xx status code
func (o *PostWidgetsDeploymentsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post widgets deployments request entity too large response has a 5xx status code
func (o *PostWidgetsDeploymentsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post widgets deployments request entity too large response a status code equal to that given
func (o *PostWidgetsDeploymentsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostWidgetsDeploymentsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWidgetsDeploymentsRequestEntityTooLarge) String() string {
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

/*
PostWidgetsDeploymentsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWidgetsDeploymentsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post widgets deployments unsupported media type response has a 2xx status code
func (o *PostWidgetsDeploymentsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post widgets deployments unsupported media type response has a 3xx status code
func (o *PostWidgetsDeploymentsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post widgets deployments unsupported media type response has a 4xx status code
func (o *PostWidgetsDeploymentsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post widgets deployments unsupported media type response has a 5xx status code
func (o *PostWidgetsDeploymentsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post widgets deployments unsupported media type response a status code equal to that given
func (o *PostWidgetsDeploymentsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostWidgetsDeploymentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWidgetsDeploymentsUnsupportedMediaType) String() string {
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

/*
PostWidgetsDeploymentsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostWidgetsDeploymentsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post widgets deployments too many requests response has a 2xx status code
func (o *PostWidgetsDeploymentsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post widgets deployments too many requests response has a 3xx status code
func (o *PostWidgetsDeploymentsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post widgets deployments too many requests response has a 4xx status code
func (o *PostWidgetsDeploymentsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post widgets deployments too many requests response has a 5xx status code
func (o *PostWidgetsDeploymentsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post widgets deployments too many requests response a status code equal to that given
func (o *PostWidgetsDeploymentsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostWidgetsDeploymentsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWidgetsDeploymentsTooManyRequests) String() string {
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

/*
PostWidgetsDeploymentsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWidgetsDeploymentsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post widgets deployments internal server error response has a 2xx status code
func (o *PostWidgetsDeploymentsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post widgets deployments internal server error response has a 3xx status code
func (o *PostWidgetsDeploymentsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post widgets deployments internal server error response has a 4xx status code
func (o *PostWidgetsDeploymentsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post widgets deployments internal server error response has a 5xx status code
func (o *PostWidgetsDeploymentsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post widgets deployments internal server error response a status code equal to that given
func (o *PostWidgetsDeploymentsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostWidgetsDeploymentsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWidgetsDeploymentsInternalServerError) String() string {
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

/*
PostWidgetsDeploymentsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWidgetsDeploymentsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post widgets deployments service unavailable response has a 2xx status code
func (o *PostWidgetsDeploymentsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post widgets deployments service unavailable response has a 3xx status code
func (o *PostWidgetsDeploymentsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post widgets deployments service unavailable response has a 4xx status code
func (o *PostWidgetsDeploymentsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post widgets deployments service unavailable response has a 5xx status code
func (o *PostWidgetsDeploymentsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post widgets deployments service unavailable response a status code equal to that given
func (o *PostWidgetsDeploymentsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostWidgetsDeploymentsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWidgetsDeploymentsServiceUnavailable) String() string {
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

/*
PostWidgetsDeploymentsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostWidgetsDeploymentsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post widgets deployments gateway timeout response has a 2xx status code
func (o *PostWidgetsDeploymentsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post widgets deployments gateway timeout response has a 3xx status code
func (o *PostWidgetsDeploymentsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post widgets deployments gateway timeout response has a 4xx status code
func (o *PostWidgetsDeploymentsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post widgets deployments gateway timeout response has a 5xx status code
func (o *PostWidgetsDeploymentsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post widgets deployments gateway timeout response a status code equal to that given
func (o *PostWidgetsDeploymentsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostWidgetsDeploymentsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/widgets/deployments][%d] postWidgetsDeploymentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWidgetsDeploymentsGatewayTimeout) String() string {
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
