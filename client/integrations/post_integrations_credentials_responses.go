// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostIntegrationsCredentialsReader is a Reader for the PostIntegrationsCredentials structure.
type PostIntegrationsCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIntegrationsCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostIntegrationsCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostIntegrationsCredentialsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostIntegrationsCredentialsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostIntegrationsCredentialsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostIntegrationsCredentialsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostIntegrationsCredentialsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostIntegrationsCredentialsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostIntegrationsCredentialsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostIntegrationsCredentialsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostIntegrationsCredentialsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostIntegrationsCredentialsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostIntegrationsCredentialsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostIntegrationsCredentialsOK creates a PostIntegrationsCredentialsOK with default headers values
func NewPostIntegrationsCredentialsOK() *PostIntegrationsCredentialsOK {
	return &PostIntegrationsCredentialsOK{}
}

/*
PostIntegrationsCredentialsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostIntegrationsCredentialsOK struct {
	Payload *models.CredentialInfo
}

// IsSuccess returns true when this post integrations credentials o k response has a 2xx status code
func (o *PostIntegrationsCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post integrations credentials o k response has a 3xx status code
func (o *PostIntegrationsCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations credentials o k response has a 4xx status code
func (o *PostIntegrationsCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post integrations credentials o k response has a 5xx status code
func (o *PostIntegrationsCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations credentials o k response a status code equal to that given
func (o *PostIntegrationsCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostIntegrationsCredentialsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsOK  %+v", 200, o.Payload)
}

func (o *PostIntegrationsCredentialsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsOK  %+v", 200, o.Payload)
}

func (o *PostIntegrationsCredentialsOK) GetPayload() *models.CredentialInfo {
	return o.Payload
}

func (o *PostIntegrationsCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsCredentialsBadRequest creates a PostIntegrationsCredentialsBadRequest with default headers values
func NewPostIntegrationsCredentialsBadRequest() *PostIntegrationsCredentialsBadRequest {
	return &PostIntegrationsCredentialsBadRequest{}
}

/*
PostIntegrationsCredentialsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostIntegrationsCredentialsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations credentials bad request response has a 2xx status code
func (o *PostIntegrationsCredentialsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations credentials bad request response has a 3xx status code
func (o *PostIntegrationsCredentialsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations credentials bad request response has a 4xx status code
func (o *PostIntegrationsCredentialsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations credentials bad request response has a 5xx status code
func (o *PostIntegrationsCredentialsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations credentials bad request response a status code equal to that given
func (o *PostIntegrationsCredentialsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostIntegrationsCredentialsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsBadRequest  %+v", 400, o.Payload)
}

func (o *PostIntegrationsCredentialsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsBadRequest  %+v", 400, o.Payload)
}

func (o *PostIntegrationsCredentialsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsCredentialsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsCredentialsUnauthorized creates a PostIntegrationsCredentialsUnauthorized with default headers values
func NewPostIntegrationsCredentialsUnauthorized() *PostIntegrationsCredentialsUnauthorized {
	return &PostIntegrationsCredentialsUnauthorized{}
}

/*
PostIntegrationsCredentialsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostIntegrationsCredentialsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations credentials unauthorized response has a 2xx status code
func (o *PostIntegrationsCredentialsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations credentials unauthorized response has a 3xx status code
func (o *PostIntegrationsCredentialsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations credentials unauthorized response has a 4xx status code
func (o *PostIntegrationsCredentialsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations credentials unauthorized response has a 5xx status code
func (o *PostIntegrationsCredentialsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations credentials unauthorized response a status code equal to that given
func (o *PostIntegrationsCredentialsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostIntegrationsCredentialsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostIntegrationsCredentialsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostIntegrationsCredentialsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsCredentialsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsCredentialsForbidden creates a PostIntegrationsCredentialsForbidden with default headers values
func NewPostIntegrationsCredentialsForbidden() *PostIntegrationsCredentialsForbidden {
	return &PostIntegrationsCredentialsForbidden{}
}

/*
PostIntegrationsCredentialsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostIntegrationsCredentialsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations credentials forbidden response has a 2xx status code
func (o *PostIntegrationsCredentialsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations credentials forbidden response has a 3xx status code
func (o *PostIntegrationsCredentialsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations credentials forbidden response has a 4xx status code
func (o *PostIntegrationsCredentialsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations credentials forbidden response has a 5xx status code
func (o *PostIntegrationsCredentialsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations credentials forbidden response a status code equal to that given
func (o *PostIntegrationsCredentialsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostIntegrationsCredentialsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsForbidden  %+v", 403, o.Payload)
}

func (o *PostIntegrationsCredentialsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsForbidden  %+v", 403, o.Payload)
}

func (o *PostIntegrationsCredentialsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsCredentialsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsCredentialsNotFound creates a PostIntegrationsCredentialsNotFound with default headers values
func NewPostIntegrationsCredentialsNotFound() *PostIntegrationsCredentialsNotFound {
	return &PostIntegrationsCredentialsNotFound{}
}

/*
PostIntegrationsCredentialsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostIntegrationsCredentialsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations credentials not found response has a 2xx status code
func (o *PostIntegrationsCredentialsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations credentials not found response has a 3xx status code
func (o *PostIntegrationsCredentialsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations credentials not found response has a 4xx status code
func (o *PostIntegrationsCredentialsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations credentials not found response has a 5xx status code
func (o *PostIntegrationsCredentialsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations credentials not found response a status code equal to that given
func (o *PostIntegrationsCredentialsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostIntegrationsCredentialsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsNotFound  %+v", 404, o.Payload)
}

func (o *PostIntegrationsCredentialsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsNotFound  %+v", 404, o.Payload)
}

func (o *PostIntegrationsCredentialsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsCredentialsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsCredentialsRequestTimeout creates a PostIntegrationsCredentialsRequestTimeout with default headers values
func NewPostIntegrationsCredentialsRequestTimeout() *PostIntegrationsCredentialsRequestTimeout {
	return &PostIntegrationsCredentialsRequestTimeout{}
}

/*
PostIntegrationsCredentialsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostIntegrationsCredentialsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations credentials request timeout response has a 2xx status code
func (o *PostIntegrationsCredentialsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations credentials request timeout response has a 3xx status code
func (o *PostIntegrationsCredentialsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations credentials request timeout response has a 4xx status code
func (o *PostIntegrationsCredentialsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations credentials request timeout response has a 5xx status code
func (o *PostIntegrationsCredentialsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations credentials request timeout response a status code equal to that given
func (o *PostIntegrationsCredentialsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostIntegrationsCredentialsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostIntegrationsCredentialsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostIntegrationsCredentialsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsCredentialsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsCredentialsRequestEntityTooLarge creates a PostIntegrationsCredentialsRequestEntityTooLarge with default headers values
func NewPostIntegrationsCredentialsRequestEntityTooLarge() *PostIntegrationsCredentialsRequestEntityTooLarge {
	return &PostIntegrationsCredentialsRequestEntityTooLarge{}
}

/*
PostIntegrationsCredentialsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostIntegrationsCredentialsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations credentials request entity too large response has a 2xx status code
func (o *PostIntegrationsCredentialsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations credentials request entity too large response has a 3xx status code
func (o *PostIntegrationsCredentialsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations credentials request entity too large response has a 4xx status code
func (o *PostIntegrationsCredentialsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations credentials request entity too large response has a 5xx status code
func (o *PostIntegrationsCredentialsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations credentials request entity too large response a status code equal to that given
func (o *PostIntegrationsCredentialsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostIntegrationsCredentialsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostIntegrationsCredentialsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostIntegrationsCredentialsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsCredentialsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsCredentialsUnsupportedMediaType creates a PostIntegrationsCredentialsUnsupportedMediaType with default headers values
func NewPostIntegrationsCredentialsUnsupportedMediaType() *PostIntegrationsCredentialsUnsupportedMediaType {
	return &PostIntegrationsCredentialsUnsupportedMediaType{}
}

/*
PostIntegrationsCredentialsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostIntegrationsCredentialsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations credentials unsupported media type response has a 2xx status code
func (o *PostIntegrationsCredentialsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations credentials unsupported media type response has a 3xx status code
func (o *PostIntegrationsCredentialsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations credentials unsupported media type response has a 4xx status code
func (o *PostIntegrationsCredentialsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations credentials unsupported media type response has a 5xx status code
func (o *PostIntegrationsCredentialsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations credentials unsupported media type response a status code equal to that given
func (o *PostIntegrationsCredentialsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostIntegrationsCredentialsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostIntegrationsCredentialsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostIntegrationsCredentialsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsCredentialsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsCredentialsTooManyRequests creates a PostIntegrationsCredentialsTooManyRequests with default headers values
func NewPostIntegrationsCredentialsTooManyRequests() *PostIntegrationsCredentialsTooManyRequests {
	return &PostIntegrationsCredentialsTooManyRequests{}
}

/*
PostIntegrationsCredentialsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostIntegrationsCredentialsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations credentials too many requests response has a 2xx status code
func (o *PostIntegrationsCredentialsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations credentials too many requests response has a 3xx status code
func (o *PostIntegrationsCredentialsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations credentials too many requests response has a 4xx status code
func (o *PostIntegrationsCredentialsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post integrations credentials too many requests response has a 5xx status code
func (o *PostIntegrationsCredentialsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post integrations credentials too many requests response a status code equal to that given
func (o *PostIntegrationsCredentialsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostIntegrationsCredentialsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostIntegrationsCredentialsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostIntegrationsCredentialsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsCredentialsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsCredentialsInternalServerError creates a PostIntegrationsCredentialsInternalServerError with default headers values
func NewPostIntegrationsCredentialsInternalServerError() *PostIntegrationsCredentialsInternalServerError {
	return &PostIntegrationsCredentialsInternalServerError{}
}

/*
PostIntegrationsCredentialsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostIntegrationsCredentialsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations credentials internal server error response has a 2xx status code
func (o *PostIntegrationsCredentialsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations credentials internal server error response has a 3xx status code
func (o *PostIntegrationsCredentialsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations credentials internal server error response has a 4xx status code
func (o *PostIntegrationsCredentialsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post integrations credentials internal server error response has a 5xx status code
func (o *PostIntegrationsCredentialsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post integrations credentials internal server error response a status code equal to that given
func (o *PostIntegrationsCredentialsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostIntegrationsCredentialsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostIntegrationsCredentialsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostIntegrationsCredentialsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsCredentialsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsCredentialsServiceUnavailable creates a PostIntegrationsCredentialsServiceUnavailable with default headers values
func NewPostIntegrationsCredentialsServiceUnavailable() *PostIntegrationsCredentialsServiceUnavailable {
	return &PostIntegrationsCredentialsServiceUnavailable{}
}

/*
PostIntegrationsCredentialsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostIntegrationsCredentialsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations credentials service unavailable response has a 2xx status code
func (o *PostIntegrationsCredentialsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations credentials service unavailable response has a 3xx status code
func (o *PostIntegrationsCredentialsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations credentials service unavailable response has a 4xx status code
func (o *PostIntegrationsCredentialsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post integrations credentials service unavailable response has a 5xx status code
func (o *PostIntegrationsCredentialsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post integrations credentials service unavailable response a status code equal to that given
func (o *PostIntegrationsCredentialsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostIntegrationsCredentialsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostIntegrationsCredentialsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostIntegrationsCredentialsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsCredentialsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsCredentialsGatewayTimeout creates a PostIntegrationsCredentialsGatewayTimeout with default headers values
func NewPostIntegrationsCredentialsGatewayTimeout() *PostIntegrationsCredentialsGatewayTimeout {
	return &PostIntegrationsCredentialsGatewayTimeout{}
}

/*
PostIntegrationsCredentialsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostIntegrationsCredentialsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post integrations credentials gateway timeout response has a 2xx status code
func (o *PostIntegrationsCredentialsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post integrations credentials gateway timeout response has a 3xx status code
func (o *PostIntegrationsCredentialsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post integrations credentials gateway timeout response has a 4xx status code
func (o *PostIntegrationsCredentialsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post integrations credentials gateway timeout response has a 5xx status code
func (o *PostIntegrationsCredentialsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post integrations credentials gateway timeout response a status code equal to that given
func (o *PostIntegrationsCredentialsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostIntegrationsCredentialsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostIntegrationsCredentialsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/integrations/credentials][%d] postIntegrationsCredentialsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostIntegrationsCredentialsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsCredentialsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
