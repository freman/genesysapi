// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostArchitectSystempromptResourcesReader is a Reader for the PostArchitectSystempromptResources structure.
type PostArchitectSystempromptResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostArchitectSystempromptResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostArchitectSystempromptResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostArchitectSystempromptResourcesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostArchitectSystempromptResourcesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostArchitectSystempromptResourcesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostArchitectSystempromptResourcesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostArchitectSystempromptResourcesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostArchitectSystempromptResourcesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostArchitectSystempromptResourcesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostArchitectSystempromptResourcesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostArchitectSystempromptResourcesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostArchitectSystempromptResourcesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostArchitectSystempromptResourcesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostArchitectSystempromptResourcesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostArchitectSystempromptResourcesOK creates a PostArchitectSystempromptResourcesOK with default headers values
func NewPostArchitectSystempromptResourcesOK() *PostArchitectSystempromptResourcesOK {
	return &PostArchitectSystempromptResourcesOK{}
}

/*
PostArchitectSystempromptResourcesOK describes a response with status code 200, with default header values.

successful operation
*/
type PostArchitectSystempromptResourcesOK struct {
	Payload *models.SystemPromptAsset
}

// IsSuccess returns true when this post architect systemprompt resources o k response has a 2xx status code
func (o *PostArchitectSystempromptResourcesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post architect systemprompt resources o k response has a 3xx status code
func (o *PostArchitectSystempromptResourcesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect systemprompt resources o k response has a 4xx status code
func (o *PostArchitectSystempromptResourcesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post architect systemprompt resources o k response has a 5xx status code
func (o *PostArchitectSystempromptResourcesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect systemprompt resources o k response a status code equal to that given
func (o *PostArchitectSystempromptResourcesOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostArchitectSystempromptResourcesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesOK  %+v", 200, o.Payload)
}

func (o *PostArchitectSystempromptResourcesOK) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesOK  %+v", 200, o.Payload)
}

func (o *PostArchitectSystempromptResourcesOK) GetPayload() *models.SystemPromptAsset {
	return o.Payload
}

func (o *PostArchitectSystempromptResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemPromptAsset)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSystempromptResourcesBadRequest creates a PostArchitectSystempromptResourcesBadRequest with default headers values
func NewPostArchitectSystempromptResourcesBadRequest() *PostArchitectSystempromptResourcesBadRequest {
	return &PostArchitectSystempromptResourcesBadRequest{}
}

/*
PostArchitectSystempromptResourcesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostArchitectSystempromptResourcesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect systemprompt resources bad request response has a 2xx status code
func (o *PostArchitectSystempromptResourcesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect systemprompt resources bad request response has a 3xx status code
func (o *PostArchitectSystempromptResourcesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect systemprompt resources bad request response has a 4xx status code
func (o *PostArchitectSystempromptResourcesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect systemprompt resources bad request response has a 5xx status code
func (o *PostArchitectSystempromptResourcesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect systemprompt resources bad request response a status code equal to that given
func (o *PostArchitectSystempromptResourcesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostArchitectSystempromptResourcesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesBadRequest  %+v", 400, o.Payload)
}

func (o *PostArchitectSystempromptResourcesBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesBadRequest  %+v", 400, o.Payload)
}

func (o *PostArchitectSystempromptResourcesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSystempromptResourcesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSystempromptResourcesUnauthorized creates a PostArchitectSystempromptResourcesUnauthorized with default headers values
func NewPostArchitectSystempromptResourcesUnauthorized() *PostArchitectSystempromptResourcesUnauthorized {
	return &PostArchitectSystempromptResourcesUnauthorized{}
}

/*
PostArchitectSystempromptResourcesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostArchitectSystempromptResourcesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect systemprompt resources unauthorized response has a 2xx status code
func (o *PostArchitectSystempromptResourcesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect systemprompt resources unauthorized response has a 3xx status code
func (o *PostArchitectSystempromptResourcesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect systemprompt resources unauthorized response has a 4xx status code
func (o *PostArchitectSystempromptResourcesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect systemprompt resources unauthorized response has a 5xx status code
func (o *PostArchitectSystempromptResourcesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect systemprompt resources unauthorized response a status code equal to that given
func (o *PostArchitectSystempromptResourcesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostArchitectSystempromptResourcesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostArchitectSystempromptResourcesUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostArchitectSystempromptResourcesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSystempromptResourcesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSystempromptResourcesForbidden creates a PostArchitectSystempromptResourcesForbidden with default headers values
func NewPostArchitectSystempromptResourcesForbidden() *PostArchitectSystempromptResourcesForbidden {
	return &PostArchitectSystempromptResourcesForbidden{}
}

/*
PostArchitectSystempromptResourcesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostArchitectSystempromptResourcesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect systemprompt resources forbidden response has a 2xx status code
func (o *PostArchitectSystempromptResourcesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect systemprompt resources forbidden response has a 3xx status code
func (o *PostArchitectSystempromptResourcesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect systemprompt resources forbidden response has a 4xx status code
func (o *PostArchitectSystempromptResourcesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect systemprompt resources forbidden response has a 5xx status code
func (o *PostArchitectSystempromptResourcesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect systemprompt resources forbidden response a status code equal to that given
func (o *PostArchitectSystempromptResourcesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostArchitectSystempromptResourcesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesForbidden  %+v", 403, o.Payload)
}

func (o *PostArchitectSystempromptResourcesForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesForbidden  %+v", 403, o.Payload)
}

func (o *PostArchitectSystempromptResourcesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSystempromptResourcesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSystempromptResourcesNotFound creates a PostArchitectSystempromptResourcesNotFound with default headers values
func NewPostArchitectSystempromptResourcesNotFound() *PostArchitectSystempromptResourcesNotFound {
	return &PostArchitectSystempromptResourcesNotFound{}
}

/*
PostArchitectSystempromptResourcesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostArchitectSystempromptResourcesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect systemprompt resources not found response has a 2xx status code
func (o *PostArchitectSystempromptResourcesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect systemprompt resources not found response has a 3xx status code
func (o *PostArchitectSystempromptResourcesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect systemprompt resources not found response has a 4xx status code
func (o *PostArchitectSystempromptResourcesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect systemprompt resources not found response has a 5xx status code
func (o *PostArchitectSystempromptResourcesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect systemprompt resources not found response a status code equal to that given
func (o *PostArchitectSystempromptResourcesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostArchitectSystempromptResourcesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesNotFound  %+v", 404, o.Payload)
}

func (o *PostArchitectSystempromptResourcesNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesNotFound  %+v", 404, o.Payload)
}

func (o *PostArchitectSystempromptResourcesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSystempromptResourcesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSystempromptResourcesRequestTimeout creates a PostArchitectSystempromptResourcesRequestTimeout with default headers values
func NewPostArchitectSystempromptResourcesRequestTimeout() *PostArchitectSystempromptResourcesRequestTimeout {
	return &PostArchitectSystempromptResourcesRequestTimeout{}
}

/*
PostArchitectSystempromptResourcesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostArchitectSystempromptResourcesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect systemprompt resources request timeout response has a 2xx status code
func (o *PostArchitectSystempromptResourcesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect systemprompt resources request timeout response has a 3xx status code
func (o *PostArchitectSystempromptResourcesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect systemprompt resources request timeout response has a 4xx status code
func (o *PostArchitectSystempromptResourcesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect systemprompt resources request timeout response has a 5xx status code
func (o *PostArchitectSystempromptResourcesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect systemprompt resources request timeout response a status code equal to that given
func (o *PostArchitectSystempromptResourcesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostArchitectSystempromptResourcesRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostArchitectSystempromptResourcesRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostArchitectSystempromptResourcesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSystempromptResourcesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSystempromptResourcesConflict creates a PostArchitectSystempromptResourcesConflict with default headers values
func NewPostArchitectSystempromptResourcesConflict() *PostArchitectSystempromptResourcesConflict {
	return &PostArchitectSystempromptResourcesConflict{}
}

/*
PostArchitectSystempromptResourcesConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostArchitectSystempromptResourcesConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect systemprompt resources conflict response has a 2xx status code
func (o *PostArchitectSystempromptResourcesConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect systemprompt resources conflict response has a 3xx status code
func (o *PostArchitectSystempromptResourcesConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect systemprompt resources conflict response has a 4xx status code
func (o *PostArchitectSystempromptResourcesConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect systemprompt resources conflict response has a 5xx status code
func (o *PostArchitectSystempromptResourcesConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect systemprompt resources conflict response a status code equal to that given
func (o *PostArchitectSystempromptResourcesConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PostArchitectSystempromptResourcesConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesConflict  %+v", 409, o.Payload)
}

func (o *PostArchitectSystempromptResourcesConflict) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesConflict  %+v", 409, o.Payload)
}

func (o *PostArchitectSystempromptResourcesConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSystempromptResourcesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSystempromptResourcesRequestEntityTooLarge creates a PostArchitectSystempromptResourcesRequestEntityTooLarge with default headers values
func NewPostArchitectSystempromptResourcesRequestEntityTooLarge() *PostArchitectSystempromptResourcesRequestEntityTooLarge {
	return &PostArchitectSystempromptResourcesRequestEntityTooLarge{}
}

/*
PostArchitectSystempromptResourcesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostArchitectSystempromptResourcesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect systemprompt resources request entity too large response has a 2xx status code
func (o *PostArchitectSystempromptResourcesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect systemprompt resources request entity too large response has a 3xx status code
func (o *PostArchitectSystempromptResourcesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect systemprompt resources request entity too large response has a 4xx status code
func (o *PostArchitectSystempromptResourcesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect systemprompt resources request entity too large response has a 5xx status code
func (o *PostArchitectSystempromptResourcesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect systemprompt resources request entity too large response a status code equal to that given
func (o *PostArchitectSystempromptResourcesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostArchitectSystempromptResourcesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostArchitectSystempromptResourcesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostArchitectSystempromptResourcesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSystempromptResourcesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSystempromptResourcesUnsupportedMediaType creates a PostArchitectSystempromptResourcesUnsupportedMediaType with default headers values
func NewPostArchitectSystempromptResourcesUnsupportedMediaType() *PostArchitectSystempromptResourcesUnsupportedMediaType {
	return &PostArchitectSystempromptResourcesUnsupportedMediaType{}
}

/*
PostArchitectSystempromptResourcesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostArchitectSystempromptResourcesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect systemprompt resources unsupported media type response has a 2xx status code
func (o *PostArchitectSystempromptResourcesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect systemprompt resources unsupported media type response has a 3xx status code
func (o *PostArchitectSystempromptResourcesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect systemprompt resources unsupported media type response has a 4xx status code
func (o *PostArchitectSystempromptResourcesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect systemprompt resources unsupported media type response has a 5xx status code
func (o *PostArchitectSystempromptResourcesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect systemprompt resources unsupported media type response a status code equal to that given
func (o *PostArchitectSystempromptResourcesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostArchitectSystempromptResourcesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostArchitectSystempromptResourcesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostArchitectSystempromptResourcesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSystempromptResourcesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSystempromptResourcesTooManyRequests creates a PostArchitectSystempromptResourcesTooManyRequests with default headers values
func NewPostArchitectSystempromptResourcesTooManyRequests() *PostArchitectSystempromptResourcesTooManyRequests {
	return &PostArchitectSystempromptResourcesTooManyRequests{}
}

/*
PostArchitectSystempromptResourcesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostArchitectSystempromptResourcesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect systemprompt resources too many requests response has a 2xx status code
func (o *PostArchitectSystempromptResourcesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect systemprompt resources too many requests response has a 3xx status code
func (o *PostArchitectSystempromptResourcesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect systemprompt resources too many requests response has a 4xx status code
func (o *PostArchitectSystempromptResourcesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post architect systemprompt resources too many requests response has a 5xx status code
func (o *PostArchitectSystempromptResourcesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post architect systemprompt resources too many requests response a status code equal to that given
func (o *PostArchitectSystempromptResourcesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostArchitectSystempromptResourcesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostArchitectSystempromptResourcesTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostArchitectSystempromptResourcesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSystempromptResourcesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSystempromptResourcesInternalServerError creates a PostArchitectSystempromptResourcesInternalServerError with default headers values
func NewPostArchitectSystempromptResourcesInternalServerError() *PostArchitectSystempromptResourcesInternalServerError {
	return &PostArchitectSystempromptResourcesInternalServerError{}
}

/*
PostArchitectSystempromptResourcesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostArchitectSystempromptResourcesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect systemprompt resources internal server error response has a 2xx status code
func (o *PostArchitectSystempromptResourcesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect systemprompt resources internal server error response has a 3xx status code
func (o *PostArchitectSystempromptResourcesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect systemprompt resources internal server error response has a 4xx status code
func (o *PostArchitectSystempromptResourcesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post architect systemprompt resources internal server error response has a 5xx status code
func (o *PostArchitectSystempromptResourcesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post architect systemprompt resources internal server error response a status code equal to that given
func (o *PostArchitectSystempromptResourcesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostArchitectSystempromptResourcesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostArchitectSystempromptResourcesInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostArchitectSystempromptResourcesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSystempromptResourcesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSystempromptResourcesServiceUnavailable creates a PostArchitectSystempromptResourcesServiceUnavailable with default headers values
func NewPostArchitectSystempromptResourcesServiceUnavailable() *PostArchitectSystempromptResourcesServiceUnavailable {
	return &PostArchitectSystempromptResourcesServiceUnavailable{}
}

/*
PostArchitectSystempromptResourcesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostArchitectSystempromptResourcesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect systemprompt resources service unavailable response has a 2xx status code
func (o *PostArchitectSystempromptResourcesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect systemprompt resources service unavailable response has a 3xx status code
func (o *PostArchitectSystempromptResourcesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect systemprompt resources service unavailable response has a 4xx status code
func (o *PostArchitectSystempromptResourcesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post architect systemprompt resources service unavailable response has a 5xx status code
func (o *PostArchitectSystempromptResourcesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post architect systemprompt resources service unavailable response a status code equal to that given
func (o *PostArchitectSystempromptResourcesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostArchitectSystempromptResourcesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostArchitectSystempromptResourcesServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostArchitectSystempromptResourcesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSystempromptResourcesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectSystempromptResourcesGatewayTimeout creates a PostArchitectSystempromptResourcesGatewayTimeout with default headers values
func NewPostArchitectSystempromptResourcesGatewayTimeout() *PostArchitectSystempromptResourcesGatewayTimeout {
	return &PostArchitectSystempromptResourcesGatewayTimeout{}
}

/*
PostArchitectSystempromptResourcesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostArchitectSystempromptResourcesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post architect systemprompt resources gateway timeout response has a 2xx status code
func (o *PostArchitectSystempromptResourcesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post architect systemprompt resources gateway timeout response has a 3xx status code
func (o *PostArchitectSystempromptResourcesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post architect systemprompt resources gateway timeout response has a 4xx status code
func (o *PostArchitectSystempromptResourcesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post architect systemprompt resources gateway timeout response has a 5xx status code
func (o *PostArchitectSystempromptResourcesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post architect systemprompt resources gateway timeout response a status code equal to that given
func (o *PostArchitectSystempromptResourcesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostArchitectSystempromptResourcesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostArchitectSystempromptResourcesGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/architect/systemprompts/{promptId}/resources][%d] postArchitectSystempromptResourcesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostArchitectSystempromptResourcesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectSystempromptResourcesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
