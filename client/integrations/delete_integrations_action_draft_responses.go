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

// DeleteIntegrationsActionDraftReader is a Reader for the DeleteIntegrationsActionDraft structure.
type DeleteIntegrationsActionDraftReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIntegrationsActionDraftReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteIntegrationsActionDraftNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteIntegrationsActionDraftBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteIntegrationsActionDraftUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteIntegrationsActionDraftForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteIntegrationsActionDraftNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteIntegrationsActionDraftRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteIntegrationsActionDraftRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteIntegrationsActionDraftUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteIntegrationsActionDraftTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteIntegrationsActionDraftInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteIntegrationsActionDraftServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteIntegrationsActionDraftGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteIntegrationsActionDraftNoContent creates a DeleteIntegrationsActionDraftNoContent with default headers values
func NewDeleteIntegrationsActionDraftNoContent() *DeleteIntegrationsActionDraftNoContent {
	return &DeleteIntegrationsActionDraftNoContent{}
}

/*
DeleteIntegrationsActionDraftNoContent describes a response with status code 204, with default header values.

Delete was successful
*/
type DeleteIntegrationsActionDraftNoContent struct {
}

// IsSuccess returns true when this delete integrations action draft no content response has a 2xx status code
func (o *DeleteIntegrationsActionDraftNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete integrations action draft no content response has a 3xx status code
func (o *DeleteIntegrationsActionDraftNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integrations action draft no content response has a 4xx status code
func (o *DeleteIntegrationsActionDraftNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete integrations action draft no content response has a 5xx status code
func (o *DeleteIntegrationsActionDraftNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete integrations action draft no content response a status code equal to that given
func (o *DeleteIntegrationsActionDraftNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteIntegrationsActionDraftNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftNoContent ", 204)
}

func (o *DeleteIntegrationsActionDraftNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftNoContent ", 204)
}

func (o *DeleteIntegrationsActionDraftNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIntegrationsActionDraftBadRequest creates a DeleteIntegrationsActionDraftBadRequest with default headers values
func NewDeleteIntegrationsActionDraftBadRequest() *DeleteIntegrationsActionDraftBadRequest {
	return &DeleteIntegrationsActionDraftBadRequest{}
}

/*
DeleteIntegrationsActionDraftBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteIntegrationsActionDraftBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete integrations action draft bad request response has a 2xx status code
func (o *DeleteIntegrationsActionDraftBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete integrations action draft bad request response has a 3xx status code
func (o *DeleteIntegrationsActionDraftBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integrations action draft bad request response has a 4xx status code
func (o *DeleteIntegrationsActionDraftBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete integrations action draft bad request response has a 5xx status code
func (o *DeleteIntegrationsActionDraftBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete integrations action draft bad request response a status code equal to that given
func (o *DeleteIntegrationsActionDraftBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteIntegrationsActionDraftBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteIntegrationsActionDraftBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteIntegrationsActionDraftBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationsActionDraftBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationsActionDraftUnauthorized creates a DeleteIntegrationsActionDraftUnauthorized with default headers values
func NewDeleteIntegrationsActionDraftUnauthorized() *DeleteIntegrationsActionDraftUnauthorized {
	return &DeleteIntegrationsActionDraftUnauthorized{}
}

/*
DeleteIntegrationsActionDraftUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteIntegrationsActionDraftUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete integrations action draft unauthorized response has a 2xx status code
func (o *DeleteIntegrationsActionDraftUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete integrations action draft unauthorized response has a 3xx status code
func (o *DeleteIntegrationsActionDraftUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integrations action draft unauthorized response has a 4xx status code
func (o *DeleteIntegrationsActionDraftUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete integrations action draft unauthorized response has a 5xx status code
func (o *DeleteIntegrationsActionDraftUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete integrations action draft unauthorized response a status code equal to that given
func (o *DeleteIntegrationsActionDraftUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteIntegrationsActionDraftUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteIntegrationsActionDraftUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteIntegrationsActionDraftUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationsActionDraftUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationsActionDraftForbidden creates a DeleteIntegrationsActionDraftForbidden with default headers values
func NewDeleteIntegrationsActionDraftForbidden() *DeleteIntegrationsActionDraftForbidden {
	return &DeleteIntegrationsActionDraftForbidden{}
}

/*
DeleteIntegrationsActionDraftForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteIntegrationsActionDraftForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete integrations action draft forbidden response has a 2xx status code
func (o *DeleteIntegrationsActionDraftForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete integrations action draft forbidden response has a 3xx status code
func (o *DeleteIntegrationsActionDraftForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integrations action draft forbidden response has a 4xx status code
func (o *DeleteIntegrationsActionDraftForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete integrations action draft forbidden response has a 5xx status code
func (o *DeleteIntegrationsActionDraftForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete integrations action draft forbidden response a status code equal to that given
func (o *DeleteIntegrationsActionDraftForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteIntegrationsActionDraftForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftForbidden  %+v", 403, o.Payload)
}

func (o *DeleteIntegrationsActionDraftForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftForbidden  %+v", 403, o.Payload)
}

func (o *DeleteIntegrationsActionDraftForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationsActionDraftForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationsActionDraftNotFound creates a DeleteIntegrationsActionDraftNotFound with default headers values
func NewDeleteIntegrationsActionDraftNotFound() *DeleteIntegrationsActionDraftNotFound {
	return &DeleteIntegrationsActionDraftNotFound{}
}

/*
DeleteIntegrationsActionDraftNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteIntegrationsActionDraftNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete integrations action draft not found response has a 2xx status code
func (o *DeleteIntegrationsActionDraftNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete integrations action draft not found response has a 3xx status code
func (o *DeleteIntegrationsActionDraftNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integrations action draft not found response has a 4xx status code
func (o *DeleteIntegrationsActionDraftNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete integrations action draft not found response has a 5xx status code
func (o *DeleteIntegrationsActionDraftNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete integrations action draft not found response a status code equal to that given
func (o *DeleteIntegrationsActionDraftNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteIntegrationsActionDraftNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIntegrationsActionDraftNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIntegrationsActionDraftNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationsActionDraftNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationsActionDraftRequestTimeout creates a DeleteIntegrationsActionDraftRequestTimeout with default headers values
func NewDeleteIntegrationsActionDraftRequestTimeout() *DeleteIntegrationsActionDraftRequestTimeout {
	return &DeleteIntegrationsActionDraftRequestTimeout{}
}

/*
DeleteIntegrationsActionDraftRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteIntegrationsActionDraftRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete integrations action draft request timeout response has a 2xx status code
func (o *DeleteIntegrationsActionDraftRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete integrations action draft request timeout response has a 3xx status code
func (o *DeleteIntegrationsActionDraftRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integrations action draft request timeout response has a 4xx status code
func (o *DeleteIntegrationsActionDraftRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete integrations action draft request timeout response has a 5xx status code
func (o *DeleteIntegrationsActionDraftRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete integrations action draft request timeout response a status code equal to that given
func (o *DeleteIntegrationsActionDraftRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteIntegrationsActionDraftRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteIntegrationsActionDraftRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteIntegrationsActionDraftRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationsActionDraftRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationsActionDraftRequestEntityTooLarge creates a DeleteIntegrationsActionDraftRequestEntityTooLarge with default headers values
func NewDeleteIntegrationsActionDraftRequestEntityTooLarge() *DeleteIntegrationsActionDraftRequestEntityTooLarge {
	return &DeleteIntegrationsActionDraftRequestEntityTooLarge{}
}

/*
DeleteIntegrationsActionDraftRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteIntegrationsActionDraftRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete integrations action draft request entity too large response has a 2xx status code
func (o *DeleteIntegrationsActionDraftRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete integrations action draft request entity too large response has a 3xx status code
func (o *DeleteIntegrationsActionDraftRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integrations action draft request entity too large response has a 4xx status code
func (o *DeleteIntegrationsActionDraftRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete integrations action draft request entity too large response has a 5xx status code
func (o *DeleteIntegrationsActionDraftRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete integrations action draft request entity too large response a status code equal to that given
func (o *DeleteIntegrationsActionDraftRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteIntegrationsActionDraftRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteIntegrationsActionDraftRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteIntegrationsActionDraftRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationsActionDraftRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationsActionDraftUnsupportedMediaType creates a DeleteIntegrationsActionDraftUnsupportedMediaType with default headers values
func NewDeleteIntegrationsActionDraftUnsupportedMediaType() *DeleteIntegrationsActionDraftUnsupportedMediaType {
	return &DeleteIntegrationsActionDraftUnsupportedMediaType{}
}

/*
DeleteIntegrationsActionDraftUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteIntegrationsActionDraftUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete integrations action draft unsupported media type response has a 2xx status code
func (o *DeleteIntegrationsActionDraftUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete integrations action draft unsupported media type response has a 3xx status code
func (o *DeleteIntegrationsActionDraftUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integrations action draft unsupported media type response has a 4xx status code
func (o *DeleteIntegrationsActionDraftUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete integrations action draft unsupported media type response has a 5xx status code
func (o *DeleteIntegrationsActionDraftUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete integrations action draft unsupported media type response a status code equal to that given
func (o *DeleteIntegrationsActionDraftUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteIntegrationsActionDraftUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteIntegrationsActionDraftUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteIntegrationsActionDraftUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationsActionDraftUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationsActionDraftTooManyRequests creates a DeleteIntegrationsActionDraftTooManyRequests with default headers values
func NewDeleteIntegrationsActionDraftTooManyRequests() *DeleteIntegrationsActionDraftTooManyRequests {
	return &DeleteIntegrationsActionDraftTooManyRequests{}
}

/*
DeleteIntegrationsActionDraftTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteIntegrationsActionDraftTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete integrations action draft too many requests response has a 2xx status code
func (o *DeleteIntegrationsActionDraftTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete integrations action draft too many requests response has a 3xx status code
func (o *DeleteIntegrationsActionDraftTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integrations action draft too many requests response has a 4xx status code
func (o *DeleteIntegrationsActionDraftTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete integrations action draft too many requests response has a 5xx status code
func (o *DeleteIntegrationsActionDraftTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete integrations action draft too many requests response a status code equal to that given
func (o *DeleteIntegrationsActionDraftTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteIntegrationsActionDraftTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteIntegrationsActionDraftTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteIntegrationsActionDraftTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationsActionDraftTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationsActionDraftInternalServerError creates a DeleteIntegrationsActionDraftInternalServerError with default headers values
func NewDeleteIntegrationsActionDraftInternalServerError() *DeleteIntegrationsActionDraftInternalServerError {
	return &DeleteIntegrationsActionDraftInternalServerError{}
}

/*
DeleteIntegrationsActionDraftInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteIntegrationsActionDraftInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete integrations action draft internal server error response has a 2xx status code
func (o *DeleteIntegrationsActionDraftInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete integrations action draft internal server error response has a 3xx status code
func (o *DeleteIntegrationsActionDraftInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integrations action draft internal server error response has a 4xx status code
func (o *DeleteIntegrationsActionDraftInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete integrations action draft internal server error response has a 5xx status code
func (o *DeleteIntegrationsActionDraftInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete integrations action draft internal server error response a status code equal to that given
func (o *DeleteIntegrationsActionDraftInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteIntegrationsActionDraftInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteIntegrationsActionDraftInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteIntegrationsActionDraftInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationsActionDraftInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationsActionDraftServiceUnavailable creates a DeleteIntegrationsActionDraftServiceUnavailable with default headers values
func NewDeleteIntegrationsActionDraftServiceUnavailable() *DeleteIntegrationsActionDraftServiceUnavailable {
	return &DeleteIntegrationsActionDraftServiceUnavailable{}
}

/*
DeleteIntegrationsActionDraftServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteIntegrationsActionDraftServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete integrations action draft service unavailable response has a 2xx status code
func (o *DeleteIntegrationsActionDraftServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete integrations action draft service unavailable response has a 3xx status code
func (o *DeleteIntegrationsActionDraftServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integrations action draft service unavailable response has a 4xx status code
func (o *DeleteIntegrationsActionDraftServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete integrations action draft service unavailable response has a 5xx status code
func (o *DeleteIntegrationsActionDraftServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete integrations action draft service unavailable response a status code equal to that given
func (o *DeleteIntegrationsActionDraftServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteIntegrationsActionDraftServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteIntegrationsActionDraftServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteIntegrationsActionDraftServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationsActionDraftServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationsActionDraftGatewayTimeout creates a DeleteIntegrationsActionDraftGatewayTimeout with default headers values
func NewDeleteIntegrationsActionDraftGatewayTimeout() *DeleteIntegrationsActionDraftGatewayTimeout {
	return &DeleteIntegrationsActionDraftGatewayTimeout{}
}

/*
DeleteIntegrationsActionDraftGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteIntegrationsActionDraftGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete integrations action draft gateway timeout response has a 2xx status code
func (o *DeleteIntegrationsActionDraftGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete integrations action draft gateway timeout response has a 3xx status code
func (o *DeleteIntegrationsActionDraftGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integrations action draft gateway timeout response has a 4xx status code
func (o *DeleteIntegrationsActionDraftGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete integrations action draft gateway timeout response has a 5xx status code
func (o *DeleteIntegrationsActionDraftGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete integrations action draft gateway timeout response a status code equal to that given
func (o *DeleteIntegrationsActionDraftGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteIntegrationsActionDraftGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteIntegrationsActionDraftGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/actions/{actionId}/draft][%d] deleteIntegrationsActionDraftGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteIntegrationsActionDraftGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationsActionDraftGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
