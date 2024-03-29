// Code generated by go-swagger; DO NOT EDIT.

package web_deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostWebdeploymentsConfigurationsReader is a Reader for the PostWebdeploymentsConfigurations structure.
type PostWebdeploymentsConfigurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWebdeploymentsConfigurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWebdeploymentsConfigurationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostWebdeploymentsConfigurationsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWebdeploymentsConfigurationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWebdeploymentsConfigurationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWebdeploymentsConfigurationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWebdeploymentsConfigurationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostWebdeploymentsConfigurationsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWebdeploymentsConfigurationsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWebdeploymentsConfigurationsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWebdeploymentsConfigurationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWebdeploymentsConfigurationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWebdeploymentsConfigurationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWebdeploymentsConfigurationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWebdeploymentsConfigurationsOK creates a PostWebdeploymentsConfigurationsOK with default headers values
func NewPostWebdeploymentsConfigurationsOK() *PostWebdeploymentsConfigurationsOK {
	return &PostWebdeploymentsConfigurationsOK{}
}

/*
PostWebdeploymentsConfigurationsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostWebdeploymentsConfigurationsOK struct {
	Payload *models.WebDeploymentConfigurationVersion
}

// IsSuccess returns true when this post webdeployments configurations o k response has a 2xx status code
func (o *PostWebdeploymentsConfigurationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post webdeployments configurations o k response has a 3xx status code
func (o *PostWebdeploymentsConfigurationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webdeployments configurations o k response has a 4xx status code
func (o *PostWebdeploymentsConfigurationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post webdeployments configurations o k response has a 5xx status code
func (o *PostWebdeploymentsConfigurationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post webdeployments configurations o k response a status code equal to that given
func (o *PostWebdeploymentsConfigurationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostWebdeploymentsConfigurationsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsOK  %+v", 200, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsOK  %+v", 200, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsOK) GetPayload() *models.WebDeploymentConfigurationVersion {
	return o.Payload
}

func (o *PostWebdeploymentsConfigurationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebDeploymentConfigurationVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebdeploymentsConfigurationsCreated creates a PostWebdeploymentsConfigurationsCreated with default headers values
func NewPostWebdeploymentsConfigurationsCreated() *PostWebdeploymentsConfigurationsCreated {
	return &PostWebdeploymentsConfigurationsCreated{}
}

/*
PostWebdeploymentsConfigurationsCreated describes a response with status code 201, with default header values.

The configuration version draft was created successfully
*/
type PostWebdeploymentsConfigurationsCreated struct {
	Payload *models.WebDeploymentConfigurationVersion
}

// IsSuccess returns true when this post webdeployments configurations created response has a 2xx status code
func (o *PostWebdeploymentsConfigurationsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post webdeployments configurations created response has a 3xx status code
func (o *PostWebdeploymentsConfigurationsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webdeployments configurations created response has a 4xx status code
func (o *PostWebdeploymentsConfigurationsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post webdeployments configurations created response has a 5xx status code
func (o *PostWebdeploymentsConfigurationsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post webdeployments configurations created response a status code equal to that given
func (o *PostWebdeploymentsConfigurationsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostWebdeploymentsConfigurationsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsCreated  %+v", 201, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsCreated  %+v", 201, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsCreated) GetPayload() *models.WebDeploymentConfigurationVersion {
	return o.Payload
}

func (o *PostWebdeploymentsConfigurationsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebDeploymentConfigurationVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebdeploymentsConfigurationsBadRequest creates a PostWebdeploymentsConfigurationsBadRequest with default headers values
func NewPostWebdeploymentsConfigurationsBadRequest() *PostWebdeploymentsConfigurationsBadRequest {
	return &PostWebdeploymentsConfigurationsBadRequest{}
}

/*
PostWebdeploymentsConfigurationsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWebdeploymentsConfigurationsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post webdeployments configurations bad request response has a 2xx status code
func (o *PostWebdeploymentsConfigurationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webdeployments configurations bad request response has a 3xx status code
func (o *PostWebdeploymentsConfigurationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webdeployments configurations bad request response has a 4xx status code
func (o *PostWebdeploymentsConfigurationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post webdeployments configurations bad request response has a 5xx status code
func (o *PostWebdeploymentsConfigurationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post webdeployments configurations bad request response a status code equal to that given
func (o *PostWebdeploymentsConfigurationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostWebdeploymentsConfigurationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsBadRequest  %+v", 400, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsBadRequest  %+v", 400, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebdeploymentsConfigurationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebdeploymentsConfigurationsUnauthorized creates a PostWebdeploymentsConfigurationsUnauthorized with default headers values
func NewPostWebdeploymentsConfigurationsUnauthorized() *PostWebdeploymentsConfigurationsUnauthorized {
	return &PostWebdeploymentsConfigurationsUnauthorized{}
}

/*
PostWebdeploymentsConfigurationsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWebdeploymentsConfigurationsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post webdeployments configurations unauthorized response has a 2xx status code
func (o *PostWebdeploymentsConfigurationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webdeployments configurations unauthorized response has a 3xx status code
func (o *PostWebdeploymentsConfigurationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webdeployments configurations unauthorized response has a 4xx status code
func (o *PostWebdeploymentsConfigurationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post webdeployments configurations unauthorized response has a 5xx status code
func (o *PostWebdeploymentsConfigurationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post webdeployments configurations unauthorized response a status code equal to that given
func (o *PostWebdeploymentsConfigurationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostWebdeploymentsConfigurationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebdeploymentsConfigurationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebdeploymentsConfigurationsForbidden creates a PostWebdeploymentsConfigurationsForbidden with default headers values
func NewPostWebdeploymentsConfigurationsForbidden() *PostWebdeploymentsConfigurationsForbidden {
	return &PostWebdeploymentsConfigurationsForbidden{}
}

/*
PostWebdeploymentsConfigurationsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostWebdeploymentsConfigurationsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post webdeployments configurations forbidden response has a 2xx status code
func (o *PostWebdeploymentsConfigurationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webdeployments configurations forbidden response has a 3xx status code
func (o *PostWebdeploymentsConfigurationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webdeployments configurations forbidden response has a 4xx status code
func (o *PostWebdeploymentsConfigurationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post webdeployments configurations forbidden response has a 5xx status code
func (o *PostWebdeploymentsConfigurationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post webdeployments configurations forbidden response a status code equal to that given
func (o *PostWebdeploymentsConfigurationsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostWebdeploymentsConfigurationsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsForbidden  %+v", 403, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsForbidden  %+v", 403, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebdeploymentsConfigurationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebdeploymentsConfigurationsNotFound creates a PostWebdeploymentsConfigurationsNotFound with default headers values
func NewPostWebdeploymentsConfigurationsNotFound() *PostWebdeploymentsConfigurationsNotFound {
	return &PostWebdeploymentsConfigurationsNotFound{}
}

/*
PostWebdeploymentsConfigurationsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostWebdeploymentsConfigurationsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post webdeployments configurations not found response has a 2xx status code
func (o *PostWebdeploymentsConfigurationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webdeployments configurations not found response has a 3xx status code
func (o *PostWebdeploymentsConfigurationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webdeployments configurations not found response has a 4xx status code
func (o *PostWebdeploymentsConfigurationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post webdeployments configurations not found response has a 5xx status code
func (o *PostWebdeploymentsConfigurationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post webdeployments configurations not found response a status code equal to that given
func (o *PostWebdeploymentsConfigurationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostWebdeploymentsConfigurationsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsNotFound  %+v", 404, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsNotFound  %+v", 404, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebdeploymentsConfigurationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebdeploymentsConfigurationsRequestTimeout creates a PostWebdeploymentsConfigurationsRequestTimeout with default headers values
func NewPostWebdeploymentsConfigurationsRequestTimeout() *PostWebdeploymentsConfigurationsRequestTimeout {
	return &PostWebdeploymentsConfigurationsRequestTimeout{}
}

/*
PostWebdeploymentsConfigurationsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostWebdeploymentsConfigurationsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post webdeployments configurations request timeout response has a 2xx status code
func (o *PostWebdeploymentsConfigurationsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webdeployments configurations request timeout response has a 3xx status code
func (o *PostWebdeploymentsConfigurationsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webdeployments configurations request timeout response has a 4xx status code
func (o *PostWebdeploymentsConfigurationsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post webdeployments configurations request timeout response has a 5xx status code
func (o *PostWebdeploymentsConfigurationsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post webdeployments configurations request timeout response a status code equal to that given
func (o *PostWebdeploymentsConfigurationsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostWebdeploymentsConfigurationsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebdeploymentsConfigurationsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebdeploymentsConfigurationsRequestEntityTooLarge creates a PostWebdeploymentsConfigurationsRequestEntityTooLarge with default headers values
func NewPostWebdeploymentsConfigurationsRequestEntityTooLarge() *PostWebdeploymentsConfigurationsRequestEntityTooLarge {
	return &PostWebdeploymentsConfigurationsRequestEntityTooLarge{}
}

/*
PostWebdeploymentsConfigurationsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostWebdeploymentsConfigurationsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post webdeployments configurations request entity too large response has a 2xx status code
func (o *PostWebdeploymentsConfigurationsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webdeployments configurations request entity too large response has a 3xx status code
func (o *PostWebdeploymentsConfigurationsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webdeployments configurations request entity too large response has a 4xx status code
func (o *PostWebdeploymentsConfigurationsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post webdeployments configurations request entity too large response has a 5xx status code
func (o *PostWebdeploymentsConfigurationsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post webdeployments configurations request entity too large response a status code equal to that given
func (o *PostWebdeploymentsConfigurationsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostWebdeploymentsConfigurationsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebdeploymentsConfigurationsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebdeploymentsConfigurationsUnsupportedMediaType creates a PostWebdeploymentsConfigurationsUnsupportedMediaType with default headers values
func NewPostWebdeploymentsConfigurationsUnsupportedMediaType() *PostWebdeploymentsConfigurationsUnsupportedMediaType {
	return &PostWebdeploymentsConfigurationsUnsupportedMediaType{}
}

/*
PostWebdeploymentsConfigurationsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWebdeploymentsConfigurationsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post webdeployments configurations unsupported media type response has a 2xx status code
func (o *PostWebdeploymentsConfigurationsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webdeployments configurations unsupported media type response has a 3xx status code
func (o *PostWebdeploymentsConfigurationsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webdeployments configurations unsupported media type response has a 4xx status code
func (o *PostWebdeploymentsConfigurationsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post webdeployments configurations unsupported media type response has a 5xx status code
func (o *PostWebdeploymentsConfigurationsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post webdeployments configurations unsupported media type response a status code equal to that given
func (o *PostWebdeploymentsConfigurationsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostWebdeploymentsConfigurationsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebdeploymentsConfigurationsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebdeploymentsConfigurationsTooManyRequests creates a PostWebdeploymentsConfigurationsTooManyRequests with default headers values
func NewPostWebdeploymentsConfigurationsTooManyRequests() *PostWebdeploymentsConfigurationsTooManyRequests {
	return &PostWebdeploymentsConfigurationsTooManyRequests{}
}

/*
PostWebdeploymentsConfigurationsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostWebdeploymentsConfigurationsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post webdeployments configurations too many requests response has a 2xx status code
func (o *PostWebdeploymentsConfigurationsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webdeployments configurations too many requests response has a 3xx status code
func (o *PostWebdeploymentsConfigurationsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webdeployments configurations too many requests response has a 4xx status code
func (o *PostWebdeploymentsConfigurationsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post webdeployments configurations too many requests response has a 5xx status code
func (o *PostWebdeploymentsConfigurationsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post webdeployments configurations too many requests response a status code equal to that given
func (o *PostWebdeploymentsConfigurationsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostWebdeploymentsConfigurationsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebdeploymentsConfigurationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebdeploymentsConfigurationsInternalServerError creates a PostWebdeploymentsConfigurationsInternalServerError with default headers values
func NewPostWebdeploymentsConfigurationsInternalServerError() *PostWebdeploymentsConfigurationsInternalServerError {
	return &PostWebdeploymentsConfigurationsInternalServerError{}
}

/*
PostWebdeploymentsConfigurationsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWebdeploymentsConfigurationsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post webdeployments configurations internal server error response has a 2xx status code
func (o *PostWebdeploymentsConfigurationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webdeployments configurations internal server error response has a 3xx status code
func (o *PostWebdeploymentsConfigurationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webdeployments configurations internal server error response has a 4xx status code
func (o *PostWebdeploymentsConfigurationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post webdeployments configurations internal server error response has a 5xx status code
func (o *PostWebdeploymentsConfigurationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post webdeployments configurations internal server error response a status code equal to that given
func (o *PostWebdeploymentsConfigurationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostWebdeploymentsConfigurationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebdeploymentsConfigurationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebdeploymentsConfigurationsServiceUnavailable creates a PostWebdeploymentsConfigurationsServiceUnavailable with default headers values
func NewPostWebdeploymentsConfigurationsServiceUnavailable() *PostWebdeploymentsConfigurationsServiceUnavailable {
	return &PostWebdeploymentsConfigurationsServiceUnavailable{}
}

/*
PostWebdeploymentsConfigurationsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWebdeploymentsConfigurationsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post webdeployments configurations service unavailable response has a 2xx status code
func (o *PostWebdeploymentsConfigurationsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webdeployments configurations service unavailable response has a 3xx status code
func (o *PostWebdeploymentsConfigurationsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webdeployments configurations service unavailable response has a 4xx status code
func (o *PostWebdeploymentsConfigurationsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post webdeployments configurations service unavailable response has a 5xx status code
func (o *PostWebdeploymentsConfigurationsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post webdeployments configurations service unavailable response a status code equal to that given
func (o *PostWebdeploymentsConfigurationsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostWebdeploymentsConfigurationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebdeploymentsConfigurationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebdeploymentsConfigurationsGatewayTimeout creates a PostWebdeploymentsConfigurationsGatewayTimeout with default headers values
func NewPostWebdeploymentsConfigurationsGatewayTimeout() *PostWebdeploymentsConfigurationsGatewayTimeout {
	return &PostWebdeploymentsConfigurationsGatewayTimeout{}
}

/*
PostWebdeploymentsConfigurationsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostWebdeploymentsConfigurationsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post webdeployments configurations gateway timeout response has a 2xx status code
func (o *PostWebdeploymentsConfigurationsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post webdeployments configurations gateway timeout response has a 3xx status code
func (o *PostWebdeploymentsConfigurationsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post webdeployments configurations gateway timeout response has a 4xx status code
func (o *PostWebdeploymentsConfigurationsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post webdeployments configurations gateway timeout response has a 5xx status code
func (o *PostWebdeploymentsConfigurationsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post webdeployments configurations gateway timeout response a status code equal to that given
func (o *PostWebdeploymentsConfigurationsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostWebdeploymentsConfigurationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/webdeployments/configurations][%d] postWebdeploymentsConfigurationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWebdeploymentsConfigurationsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebdeploymentsConfigurationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
