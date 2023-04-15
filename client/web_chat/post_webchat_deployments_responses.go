// Code generated by go-swagger; DO NOT EDIT.

package web_chat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostWebchatDeploymentsReader is a Reader for the PostWebchatDeployments structure.
type PostWebchatDeploymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWebchatDeploymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWebchatDeploymentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWebchatDeploymentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWebchatDeploymentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWebchatDeploymentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWebchatDeploymentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostWebchatDeploymentsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostWebchatDeploymentsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWebchatDeploymentsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWebchatDeploymentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWebchatDeploymentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWebchatDeploymentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWebchatDeploymentsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWebchatDeploymentsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWebchatDeploymentsOK creates a PostWebchatDeploymentsOK with default headers values
func NewPostWebchatDeploymentsOK() *PostWebchatDeploymentsOK {
	return &PostWebchatDeploymentsOK{}
}

/*
PostWebchatDeploymentsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostWebchatDeploymentsOK struct {
	Payload *models.WebChatDeployment
}

// IsSuccess returns true when this post webchat deployments o k response has a 2xx status code
func (o *PostWebchatDeploymentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post webchat deployments o k response has a 3xx status code
func (o *PostWebchatDeploymentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webchat deployments o k response has a 4xx status code
func (o *PostWebchatDeploymentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post webchat deployments o k response has a 5xx status code
func (o *PostWebchatDeploymentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post webchat deployments o k response a status code equal to that given
func (o *PostWebchatDeploymentsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostWebchatDeploymentsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsOK  %+v", 200, o.Payload)
}

func (o *PostWebchatDeploymentsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsOK  %+v", 200, o.Payload)
}

func (o *PostWebchatDeploymentsOK) GetPayload() *models.WebChatDeployment {
	return o.Payload
}

func (o *PostWebchatDeploymentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebChatDeployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatDeploymentsBadRequest creates a PostWebchatDeploymentsBadRequest with default headers values
func NewPostWebchatDeploymentsBadRequest() *PostWebchatDeploymentsBadRequest {
	return &PostWebchatDeploymentsBadRequest{}
}

/*
PostWebchatDeploymentsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWebchatDeploymentsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post webchat deployments bad request response has a 2xx status code
func (o *PostWebchatDeploymentsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webchat deployments bad request response has a 3xx status code
func (o *PostWebchatDeploymentsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webchat deployments bad request response has a 4xx status code
func (o *PostWebchatDeploymentsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post webchat deployments bad request response has a 5xx status code
func (o *PostWebchatDeploymentsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post webchat deployments bad request response a status code equal to that given
func (o *PostWebchatDeploymentsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostWebchatDeploymentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsBadRequest  %+v", 400, o.Payload)
}

func (o *PostWebchatDeploymentsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsBadRequest  %+v", 400, o.Payload)
}

func (o *PostWebchatDeploymentsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatDeploymentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatDeploymentsUnauthorized creates a PostWebchatDeploymentsUnauthorized with default headers values
func NewPostWebchatDeploymentsUnauthorized() *PostWebchatDeploymentsUnauthorized {
	return &PostWebchatDeploymentsUnauthorized{}
}

/*
PostWebchatDeploymentsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWebchatDeploymentsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post webchat deployments unauthorized response has a 2xx status code
func (o *PostWebchatDeploymentsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webchat deployments unauthorized response has a 3xx status code
func (o *PostWebchatDeploymentsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webchat deployments unauthorized response has a 4xx status code
func (o *PostWebchatDeploymentsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post webchat deployments unauthorized response has a 5xx status code
func (o *PostWebchatDeploymentsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post webchat deployments unauthorized response a status code equal to that given
func (o *PostWebchatDeploymentsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostWebchatDeploymentsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWebchatDeploymentsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWebchatDeploymentsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatDeploymentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatDeploymentsForbidden creates a PostWebchatDeploymentsForbidden with default headers values
func NewPostWebchatDeploymentsForbidden() *PostWebchatDeploymentsForbidden {
	return &PostWebchatDeploymentsForbidden{}
}

/*
PostWebchatDeploymentsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostWebchatDeploymentsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post webchat deployments forbidden response has a 2xx status code
func (o *PostWebchatDeploymentsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webchat deployments forbidden response has a 3xx status code
func (o *PostWebchatDeploymentsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webchat deployments forbidden response has a 4xx status code
func (o *PostWebchatDeploymentsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post webchat deployments forbidden response has a 5xx status code
func (o *PostWebchatDeploymentsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post webchat deployments forbidden response a status code equal to that given
func (o *PostWebchatDeploymentsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostWebchatDeploymentsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsForbidden  %+v", 403, o.Payload)
}

func (o *PostWebchatDeploymentsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsForbidden  %+v", 403, o.Payload)
}

func (o *PostWebchatDeploymentsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatDeploymentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatDeploymentsNotFound creates a PostWebchatDeploymentsNotFound with default headers values
func NewPostWebchatDeploymentsNotFound() *PostWebchatDeploymentsNotFound {
	return &PostWebchatDeploymentsNotFound{}
}

/*
PostWebchatDeploymentsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostWebchatDeploymentsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post webchat deployments not found response has a 2xx status code
func (o *PostWebchatDeploymentsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webchat deployments not found response has a 3xx status code
func (o *PostWebchatDeploymentsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webchat deployments not found response has a 4xx status code
func (o *PostWebchatDeploymentsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post webchat deployments not found response has a 5xx status code
func (o *PostWebchatDeploymentsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post webchat deployments not found response a status code equal to that given
func (o *PostWebchatDeploymentsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostWebchatDeploymentsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsNotFound  %+v", 404, o.Payload)
}

func (o *PostWebchatDeploymentsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsNotFound  %+v", 404, o.Payload)
}

func (o *PostWebchatDeploymentsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatDeploymentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatDeploymentsRequestTimeout creates a PostWebchatDeploymentsRequestTimeout with default headers values
func NewPostWebchatDeploymentsRequestTimeout() *PostWebchatDeploymentsRequestTimeout {
	return &PostWebchatDeploymentsRequestTimeout{}
}

/*
PostWebchatDeploymentsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostWebchatDeploymentsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post webchat deployments request timeout response has a 2xx status code
func (o *PostWebchatDeploymentsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webchat deployments request timeout response has a 3xx status code
func (o *PostWebchatDeploymentsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webchat deployments request timeout response has a 4xx status code
func (o *PostWebchatDeploymentsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post webchat deployments request timeout response has a 5xx status code
func (o *PostWebchatDeploymentsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post webchat deployments request timeout response a status code equal to that given
func (o *PostWebchatDeploymentsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostWebchatDeploymentsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWebchatDeploymentsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWebchatDeploymentsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatDeploymentsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatDeploymentsConflict creates a PostWebchatDeploymentsConflict with default headers values
func NewPostWebchatDeploymentsConflict() *PostWebchatDeploymentsConflict {
	return &PostWebchatDeploymentsConflict{}
}

/*
PostWebchatDeploymentsConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostWebchatDeploymentsConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post webchat deployments conflict response has a 2xx status code
func (o *PostWebchatDeploymentsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webchat deployments conflict response has a 3xx status code
func (o *PostWebchatDeploymentsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webchat deployments conflict response has a 4xx status code
func (o *PostWebchatDeploymentsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post webchat deployments conflict response has a 5xx status code
func (o *PostWebchatDeploymentsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post webchat deployments conflict response a status code equal to that given
func (o *PostWebchatDeploymentsConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PostWebchatDeploymentsConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsConflict  %+v", 409, o.Payload)
}

func (o *PostWebchatDeploymentsConflict) String() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsConflict  %+v", 409, o.Payload)
}

func (o *PostWebchatDeploymentsConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatDeploymentsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatDeploymentsRequestEntityTooLarge creates a PostWebchatDeploymentsRequestEntityTooLarge with default headers values
func NewPostWebchatDeploymentsRequestEntityTooLarge() *PostWebchatDeploymentsRequestEntityTooLarge {
	return &PostWebchatDeploymentsRequestEntityTooLarge{}
}

/*
PostWebchatDeploymentsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostWebchatDeploymentsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post webchat deployments request entity too large response has a 2xx status code
func (o *PostWebchatDeploymentsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webchat deployments request entity too large response has a 3xx status code
func (o *PostWebchatDeploymentsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webchat deployments request entity too large response has a 4xx status code
func (o *PostWebchatDeploymentsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post webchat deployments request entity too large response has a 5xx status code
func (o *PostWebchatDeploymentsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post webchat deployments request entity too large response a status code equal to that given
func (o *PostWebchatDeploymentsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostWebchatDeploymentsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWebchatDeploymentsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWebchatDeploymentsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatDeploymentsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatDeploymentsUnsupportedMediaType creates a PostWebchatDeploymentsUnsupportedMediaType with default headers values
func NewPostWebchatDeploymentsUnsupportedMediaType() *PostWebchatDeploymentsUnsupportedMediaType {
	return &PostWebchatDeploymentsUnsupportedMediaType{}
}

/*
PostWebchatDeploymentsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWebchatDeploymentsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post webchat deployments unsupported media type response has a 2xx status code
func (o *PostWebchatDeploymentsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webchat deployments unsupported media type response has a 3xx status code
func (o *PostWebchatDeploymentsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webchat deployments unsupported media type response has a 4xx status code
func (o *PostWebchatDeploymentsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post webchat deployments unsupported media type response has a 5xx status code
func (o *PostWebchatDeploymentsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post webchat deployments unsupported media type response a status code equal to that given
func (o *PostWebchatDeploymentsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostWebchatDeploymentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWebchatDeploymentsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWebchatDeploymentsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatDeploymentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatDeploymentsTooManyRequests creates a PostWebchatDeploymentsTooManyRequests with default headers values
func NewPostWebchatDeploymentsTooManyRequests() *PostWebchatDeploymentsTooManyRequests {
	return &PostWebchatDeploymentsTooManyRequests{}
}

/*
PostWebchatDeploymentsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostWebchatDeploymentsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post webchat deployments too many requests response has a 2xx status code
func (o *PostWebchatDeploymentsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webchat deployments too many requests response has a 3xx status code
func (o *PostWebchatDeploymentsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webchat deployments too many requests response has a 4xx status code
func (o *PostWebchatDeploymentsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post webchat deployments too many requests response has a 5xx status code
func (o *PostWebchatDeploymentsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post webchat deployments too many requests response a status code equal to that given
func (o *PostWebchatDeploymentsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostWebchatDeploymentsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWebchatDeploymentsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWebchatDeploymentsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatDeploymentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatDeploymentsInternalServerError creates a PostWebchatDeploymentsInternalServerError with default headers values
func NewPostWebchatDeploymentsInternalServerError() *PostWebchatDeploymentsInternalServerError {
	return &PostWebchatDeploymentsInternalServerError{}
}

/*
PostWebchatDeploymentsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWebchatDeploymentsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post webchat deployments internal server error response has a 2xx status code
func (o *PostWebchatDeploymentsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webchat deployments internal server error response has a 3xx status code
func (o *PostWebchatDeploymentsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webchat deployments internal server error response has a 4xx status code
func (o *PostWebchatDeploymentsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post webchat deployments internal server error response has a 5xx status code
func (o *PostWebchatDeploymentsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post webchat deployments internal server error response a status code equal to that given
func (o *PostWebchatDeploymentsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostWebchatDeploymentsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWebchatDeploymentsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWebchatDeploymentsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatDeploymentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatDeploymentsServiceUnavailable creates a PostWebchatDeploymentsServiceUnavailable with default headers values
func NewPostWebchatDeploymentsServiceUnavailable() *PostWebchatDeploymentsServiceUnavailable {
	return &PostWebchatDeploymentsServiceUnavailable{}
}

/*
PostWebchatDeploymentsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWebchatDeploymentsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post webchat deployments service unavailable response has a 2xx status code
func (o *PostWebchatDeploymentsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webchat deployments service unavailable response has a 3xx status code
func (o *PostWebchatDeploymentsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webchat deployments service unavailable response has a 4xx status code
func (o *PostWebchatDeploymentsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post webchat deployments service unavailable response has a 5xx status code
func (o *PostWebchatDeploymentsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post webchat deployments service unavailable response a status code equal to that given
func (o *PostWebchatDeploymentsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostWebchatDeploymentsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWebchatDeploymentsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWebchatDeploymentsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatDeploymentsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatDeploymentsGatewayTimeout creates a PostWebchatDeploymentsGatewayTimeout with default headers values
func NewPostWebchatDeploymentsGatewayTimeout() *PostWebchatDeploymentsGatewayTimeout {
	return &PostWebchatDeploymentsGatewayTimeout{}
}

/*
PostWebchatDeploymentsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostWebchatDeploymentsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post webchat deployments gateway timeout response has a 2xx status code
func (o *PostWebchatDeploymentsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webchat deployments gateway timeout response has a 3xx status code
func (o *PostWebchatDeploymentsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webchat deployments gateway timeout response has a 4xx status code
func (o *PostWebchatDeploymentsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post webchat deployments gateway timeout response has a 5xx status code
func (o *PostWebchatDeploymentsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post webchat deployments gateway timeout response a status code equal to that given
func (o *PostWebchatDeploymentsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostWebchatDeploymentsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWebchatDeploymentsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/webchat/deployments][%d] postWebchatDeploymentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWebchatDeploymentsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatDeploymentsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
