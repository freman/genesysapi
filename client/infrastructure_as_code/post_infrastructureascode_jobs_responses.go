// Code generated by go-swagger; DO NOT EDIT.

package infrastructure_as_code

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostInfrastructureascodeJobsReader is a Reader for the PostInfrastructureascodeJobs structure.
type PostInfrastructureascodeJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostInfrastructureascodeJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostInfrastructureascodeJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPostInfrastructureascodeJobsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostInfrastructureascodeJobsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostInfrastructureascodeJobsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostInfrastructureascodeJobsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostInfrastructureascodeJobsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostInfrastructureascodeJobsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostInfrastructureascodeJobsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostInfrastructureascodeJobsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostInfrastructureascodeJobsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostInfrastructureascodeJobsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostInfrastructureascodeJobsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostInfrastructureascodeJobsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostInfrastructureascodeJobsOK creates a PostInfrastructureascodeJobsOK with default headers values
func NewPostInfrastructureascodeJobsOK() *PostInfrastructureascodeJobsOK {
	return &PostInfrastructureascodeJobsOK{}
}

/*
PostInfrastructureascodeJobsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostInfrastructureascodeJobsOK struct {
	Payload *models.InfrastructureascodeJob
}

// IsSuccess returns true when this post infrastructureascode jobs o k response has a 2xx status code
func (o *PostInfrastructureascodeJobsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post infrastructureascode jobs o k response has a 3xx status code
func (o *PostInfrastructureascodeJobsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post infrastructureascode jobs o k response has a 4xx status code
func (o *PostInfrastructureascodeJobsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post infrastructureascode jobs o k response has a 5xx status code
func (o *PostInfrastructureascodeJobsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post infrastructureascode jobs o k response a status code equal to that given
func (o *PostInfrastructureascodeJobsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostInfrastructureascodeJobsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsOK  %+v", 200, o.Payload)
}

func (o *PostInfrastructureascodeJobsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsOK  %+v", 200, o.Payload)
}

func (o *PostInfrastructureascodeJobsOK) GetPayload() *models.InfrastructureascodeJob {
	return o.Payload
}

func (o *PostInfrastructureascodeJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfrastructureascodeJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInfrastructureascodeJobsAccepted creates a PostInfrastructureascodeJobsAccepted with default headers values
func NewPostInfrastructureascodeJobsAccepted() *PostInfrastructureascodeJobsAccepted {
	return &PostInfrastructureascodeJobsAccepted{}
}

/*
PostInfrastructureascodeJobsAccepted describes a response with status code 202, with default header values.

Job submitted for execution.
*/
type PostInfrastructureascodeJobsAccepted struct {
	Payload *models.InfrastructureascodeJob
}

// IsSuccess returns true when this post infrastructureascode jobs accepted response has a 2xx status code
func (o *PostInfrastructureascodeJobsAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post infrastructureascode jobs accepted response has a 3xx status code
func (o *PostInfrastructureascodeJobsAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post infrastructureascode jobs accepted response has a 4xx status code
func (o *PostInfrastructureascodeJobsAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this post infrastructureascode jobs accepted response has a 5xx status code
func (o *PostInfrastructureascodeJobsAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this post infrastructureascode jobs accepted response a status code equal to that given
func (o *PostInfrastructureascodeJobsAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PostInfrastructureascodeJobsAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsAccepted  %+v", 202, o.Payload)
}

func (o *PostInfrastructureascodeJobsAccepted) String() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsAccepted  %+v", 202, o.Payload)
}

func (o *PostInfrastructureascodeJobsAccepted) GetPayload() *models.InfrastructureascodeJob {
	return o.Payload
}

func (o *PostInfrastructureascodeJobsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfrastructureascodeJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInfrastructureascodeJobsBadRequest creates a PostInfrastructureascodeJobsBadRequest with default headers values
func NewPostInfrastructureascodeJobsBadRequest() *PostInfrastructureascodeJobsBadRequest {
	return &PostInfrastructureascodeJobsBadRequest{}
}

/*
PostInfrastructureascodeJobsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostInfrastructureascodeJobsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post infrastructureascode jobs bad request response has a 2xx status code
func (o *PostInfrastructureascodeJobsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post infrastructureascode jobs bad request response has a 3xx status code
func (o *PostInfrastructureascodeJobsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post infrastructureascode jobs bad request response has a 4xx status code
func (o *PostInfrastructureascodeJobsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post infrastructureascode jobs bad request response has a 5xx status code
func (o *PostInfrastructureascodeJobsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post infrastructureascode jobs bad request response a status code equal to that given
func (o *PostInfrastructureascodeJobsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostInfrastructureascodeJobsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsBadRequest  %+v", 400, o.Payload)
}

func (o *PostInfrastructureascodeJobsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsBadRequest  %+v", 400, o.Payload)
}

func (o *PostInfrastructureascodeJobsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostInfrastructureascodeJobsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInfrastructureascodeJobsUnauthorized creates a PostInfrastructureascodeJobsUnauthorized with default headers values
func NewPostInfrastructureascodeJobsUnauthorized() *PostInfrastructureascodeJobsUnauthorized {
	return &PostInfrastructureascodeJobsUnauthorized{}
}

/*
PostInfrastructureascodeJobsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostInfrastructureascodeJobsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post infrastructureascode jobs unauthorized response has a 2xx status code
func (o *PostInfrastructureascodeJobsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post infrastructureascode jobs unauthorized response has a 3xx status code
func (o *PostInfrastructureascodeJobsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post infrastructureascode jobs unauthorized response has a 4xx status code
func (o *PostInfrastructureascodeJobsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post infrastructureascode jobs unauthorized response has a 5xx status code
func (o *PostInfrastructureascodeJobsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post infrastructureascode jobs unauthorized response a status code equal to that given
func (o *PostInfrastructureascodeJobsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostInfrastructureascodeJobsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostInfrastructureascodeJobsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostInfrastructureascodeJobsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostInfrastructureascodeJobsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInfrastructureascodeJobsForbidden creates a PostInfrastructureascodeJobsForbidden with default headers values
func NewPostInfrastructureascodeJobsForbidden() *PostInfrastructureascodeJobsForbidden {
	return &PostInfrastructureascodeJobsForbidden{}
}

/*
PostInfrastructureascodeJobsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostInfrastructureascodeJobsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post infrastructureascode jobs forbidden response has a 2xx status code
func (o *PostInfrastructureascodeJobsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post infrastructureascode jobs forbidden response has a 3xx status code
func (o *PostInfrastructureascodeJobsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post infrastructureascode jobs forbidden response has a 4xx status code
func (o *PostInfrastructureascodeJobsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post infrastructureascode jobs forbidden response has a 5xx status code
func (o *PostInfrastructureascodeJobsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post infrastructureascode jobs forbidden response a status code equal to that given
func (o *PostInfrastructureascodeJobsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostInfrastructureascodeJobsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsForbidden  %+v", 403, o.Payload)
}

func (o *PostInfrastructureascodeJobsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsForbidden  %+v", 403, o.Payload)
}

func (o *PostInfrastructureascodeJobsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostInfrastructureascodeJobsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInfrastructureascodeJobsNotFound creates a PostInfrastructureascodeJobsNotFound with default headers values
func NewPostInfrastructureascodeJobsNotFound() *PostInfrastructureascodeJobsNotFound {
	return &PostInfrastructureascodeJobsNotFound{}
}

/*
PostInfrastructureascodeJobsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostInfrastructureascodeJobsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post infrastructureascode jobs not found response has a 2xx status code
func (o *PostInfrastructureascodeJobsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post infrastructureascode jobs not found response has a 3xx status code
func (o *PostInfrastructureascodeJobsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post infrastructureascode jobs not found response has a 4xx status code
func (o *PostInfrastructureascodeJobsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post infrastructureascode jobs not found response has a 5xx status code
func (o *PostInfrastructureascodeJobsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post infrastructureascode jobs not found response a status code equal to that given
func (o *PostInfrastructureascodeJobsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostInfrastructureascodeJobsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsNotFound  %+v", 404, o.Payload)
}

func (o *PostInfrastructureascodeJobsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsNotFound  %+v", 404, o.Payload)
}

func (o *PostInfrastructureascodeJobsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostInfrastructureascodeJobsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInfrastructureascodeJobsRequestTimeout creates a PostInfrastructureascodeJobsRequestTimeout with default headers values
func NewPostInfrastructureascodeJobsRequestTimeout() *PostInfrastructureascodeJobsRequestTimeout {
	return &PostInfrastructureascodeJobsRequestTimeout{}
}

/*
PostInfrastructureascodeJobsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostInfrastructureascodeJobsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post infrastructureascode jobs request timeout response has a 2xx status code
func (o *PostInfrastructureascodeJobsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post infrastructureascode jobs request timeout response has a 3xx status code
func (o *PostInfrastructureascodeJobsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post infrastructureascode jobs request timeout response has a 4xx status code
func (o *PostInfrastructureascodeJobsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post infrastructureascode jobs request timeout response has a 5xx status code
func (o *PostInfrastructureascodeJobsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post infrastructureascode jobs request timeout response a status code equal to that given
func (o *PostInfrastructureascodeJobsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostInfrastructureascodeJobsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostInfrastructureascodeJobsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostInfrastructureascodeJobsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostInfrastructureascodeJobsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInfrastructureascodeJobsRequestEntityTooLarge creates a PostInfrastructureascodeJobsRequestEntityTooLarge with default headers values
func NewPostInfrastructureascodeJobsRequestEntityTooLarge() *PostInfrastructureascodeJobsRequestEntityTooLarge {
	return &PostInfrastructureascodeJobsRequestEntityTooLarge{}
}

/*
PostInfrastructureascodeJobsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostInfrastructureascodeJobsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post infrastructureascode jobs request entity too large response has a 2xx status code
func (o *PostInfrastructureascodeJobsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post infrastructureascode jobs request entity too large response has a 3xx status code
func (o *PostInfrastructureascodeJobsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post infrastructureascode jobs request entity too large response has a 4xx status code
func (o *PostInfrastructureascodeJobsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post infrastructureascode jobs request entity too large response has a 5xx status code
func (o *PostInfrastructureascodeJobsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post infrastructureascode jobs request entity too large response a status code equal to that given
func (o *PostInfrastructureascodeJobsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostInfrastructureascodeJobsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostInfrastructureascodeJobsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostInfrastructureascodeJobsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostInfrastructureascodeJobsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInfrastructureascodeJobsUnsupportedMediaType creates a PostInfrastructureascodeJobsUnsupportedMediaType with default headers values
func NewPostInfrastructureascodeJobsUnsupportedMediaType() *PostInfrastructureascodeJobsUnsupportedMediaType {
	return &PostInfrastructureascodeJobsUnsupportedMediaType{}
}

/*
PostInfrastructureascodeJobsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostInfrastructureascodeJobsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post infrastructureascode jobs unsupported media type response has a 2xx status code
func (o *PostInfrastructureascodeJobsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post infrastructureascode jobs unsupported media type response has a 3xx status code
func (o *PostInfrastructureascodeJobsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post infrastructureascode jobs unsupported media type response has a 4xx status code
func (o *PostInfrastructureascodeJobsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post infrastructureascode jobs unsupported media type response has a 5xx status code
func (o *PostInfrastructureascodeJobsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post infrastructureascode jobs unsupported media type response a status code equal to that given
func (o *PostInfrastructureascodeJobsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostInfrastructureascodeJobsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostInfrastructureascodeJobsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostInfrastructureascodeJobsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostInfrastructureascodeJobsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInfrastructureascodeJobsTooManyRequests creates a PostInfrastructureascodeJobsTooManyRequests with default headers values
func NewPostInfrastructureascodeJobsTooManyRequests() *PostInfrastructureascodeJobsTooManyRequests {
	return &PostInfrastructureascodeJobsTooManyRequests{}
}

/*
PostInfrastructureascodeJobsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostInfrastructureascodeJobsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post infrastructureascode jobs too many requests response has a 2xx status code
func (o *PostInfrastructureascodeJobsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post infrastructureascode jobs too many requests response has a 3xx status code
func (o *PostInfrastructureascodeJobsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post infrastructureascode jobs too many requests response has a 4xx status code
func (o *PostInfrastructureascodeJobsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post infrastructureascode jobs too many requests response has a 5xx status code
func (o *PostInfrastructureascodeJobsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post infrastructureascode jobs too many requests response a status code equal to that given
func (o *PostInfrastructureascodeJobsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostInfrastructureascodeJobsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostInfrastructureascodeJobsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostInfrastructureascodeJobsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostInfrastructureascodeJobsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInfrastructureascodeJobsInternalServerError creates a PostInfrastructureascodeJobsInternalServerError with default headers values
func NewPostInfrastructureascodeJobsInternalServerError() *PostInfrastructureascodeJobsInternalServerError {
	return &PostInfrastructureascodeJobsInternalServerError{}
}

/*
PostInfrastructureascodeJobsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostInfrastructureascodeJobsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post infrastructureascode jobs internal server error response has a 2xx status code
func (o *PostInfrastructureascodeJobsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post infrastructureascode jobs internal server error response has a 3xx status code
func (o *PostInfrastructureascodeJobsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post infrastructureascode jobs internal server error response has a 4xx status code
func (o *PostInfrastructureascodeJobsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post infrastructureascode jobs internal server error response has a 5xx status code
func (o *PostInfrastructureascodeJobsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post infrastructureascode jobs internal server error response a status code equal to that given
func (o *PostInfrastructureascodeJobsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostInfrastructureascodeJobsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostInfrastructureascodeJobsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostInfrastructureascodeJobsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostInfrastructureascodeJobsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInfrastructureascodeJobsServiceUnavailable creates a PostInfrastructureascodeJobsServiceUnavailable with default headers values
func NewPostInfrastructureascodeJobsServiceUnavailable() *PostInfrastructureascodeJobsServiceUnavailable {
	return &PostInfrastructureascodeJobsServiceUnavailable{}
}

/*
PostInfrastructureascodeJobsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostInfrastructureascodeJobsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post infrastructureascode jobs service unavailable response has a 2xx status code
func (o *PostInfrastructureascodeJobsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post infrastructureascode jobs service unavailable response has a 3xx status code
func (o *PostInfrastructureascodeJobsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post infrastructureascode jobs service unavailable response has a 4xx status code
func (o *PostInfrastructureascodeJobsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post infrastructureascode jobs service unavailable response has a 5xx status code
func (o *PostInfrastructureascodeJobsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post infrastructureascode jobs service unavailable response a status code equal to that given
func (o *PostInfrastructureascodeJobsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostInfrastructureascodeJobsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostInfrastructureascodeJobsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostInfrastructureascodeJobsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostInfrastructureascodeJobsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInfrastructureascodeJobsGatewayTimeout creates a PostInfrastructureascodeJobsGatewayTimeout with default headers values
func NewPostInfrastructureascodeJobsGatewayTimeout() *PostInfrastructureascodeJobsGatewayTimeout {
	return &PostInfrastructureascodeJobsGatewayTimeout{}
}

/*
PostInfrastructureascodeJobsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostInfrastructureascodeJobsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post infrastructureascode jobs gateway timeout response has a 2xx status code
func (o *PostInfrastructureascodeJobsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post infrastructureascode jobs gateway timeout response has a 3xx status code
func (o *PostInfrastructureascodeJobsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post infrastructureascode jobs gateway timeout response has a 4xx status code
func (o *PostInfrastructureascodeJobsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post infrastructureascode jobs gateway timeout response has a 5xx status code
func (o *PostInfrastructureascodeJobsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post infrastructureascode jobs gateway timeout response a status code equal to that given
func (o *PostInfrastructureascodeJobsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostInfrastructureascodeJobsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostInfrastructureascodeJobsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/infrastructureascode/jobs][%d] postInfrastructureascodeJobsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostInfrastructureascodeJobsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostInfrastructureascodeJobsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
