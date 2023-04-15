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

// PatchIntegrationsActionDraftReader is a Reader for the PatchIntegrationsActionDraft structure.
type PatchIntegrationsActionDraftReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchIntegrationsActionDraftReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchIntegrationsActionDraftOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchIntegrationsActionDraftBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchIntegrationsActionDraftUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchIntegrationsActionDraftForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchIntegrationsActionDraftNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchIntegrationsActionDraftRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchIntegrationsActionDraftRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchIntegrationsActionDraftUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchIntegrationsActionDraftTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchIntegrationsActionDraftInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchIntegrationsActionDraftServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchIntegrationsActionDraftGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchIntegrationsActionDraftOK creates a PatchIntegrationsActionDraftOK with default headers values
func NewPatchIntegrationsActionDraftOK() *PatchIntegrationsActionDraftOK {
	return &PatchIntegrationsActionDraftOK{}
}

/*
PatchIntegrationsActionDraftOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchIntegrationsActionDraftOK struct {
	Payload *models.Action
}

// IsSuccess returns true when this patch integrations action draft o k response has a 2xx status code
func (o *PatchIntegrationsActionDraftOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch integrations action draft o k response has a 3xx status code
func (o *PatchIntegrationsActionDraftOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch integrations action draft o k response has a 4xx status code
func (o *PatchIntegrationsActionDraftOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch integrations action draft o k response has a 5xx status code
func (o *PatchIntegrationsActionDraftOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch integrations action draft o k response a status code equal to that given
func (o *PatchIntegrationsActionDraftOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchIntegrationsActionDraftOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftOK  %+v", 200, o.Payload)
}

func (o *PatchIntegrationsActionDraftOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftOK  %+v", 200, o.Payload)
}

func (o *PatchIntegrationsActionDraftOK) GetPayload() *models.Action {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Action)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionDraftBadRequest creates a PatchIntegrationsActionDraftBadRequest with default headers values
func NewPatchIntegrationsActionDraftBadRequest() *PatchIntegrationsActionDraftBadRequest {
	return &PatchIntegrationsActionDraftBadRequest{}
}

/*
PatchIntegrationsActionDraftBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchIntegrationsActionDraftBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch integrations action draft bad request response has a 2xx status code
func (o *PatchIntegrationsActionDraftBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch integrations action draft bad request response has a 3xx status code
func (o *PatchIntegrationsActionDraftBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch integrations action draft bad request response has a 4xx status code
func (o *PatchIntegrationsActionDraftBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch integrations action draft bad request response has a 5xx status code
func (o *PatchIntegrationsActionDraftBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch integrations action draft bad request response a status code equal to that given
func (o *PatchIntegrationsActionDraftBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchIntegrationsActionDraftBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftBadRequest  %+v", 400, o.Payload)
}

func (o *PatchIntegrationsActionDraftBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftBadRequest  %+v", 400, o.Payload)
}

func (o *PatchIntegrationsActionDraftBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionDraftUnauthorized creates a PatchIntegrationsActionDraftUnauthorized with default headers values
func NewPatchIntegrationsActionDraftUnauthorized() *PatchIntegrationsActionDraftUnauthorized {
	return &PatchIntegrationsActionDraftUnauthorized{}
}

/*
PatchIntegrationsActionDraftUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchIntegrationsActionDraftUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch integrations action draft unauthorized response has a 2xx status code
func (o *PatchIntegrationsActionDraftUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch integrations action draft unauthorized response has a 3xx status code
func (o *PatchIntegrationsActionDraftUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch integrations action draft unauthorized response has a 4xx status code
func (o *PatchIntegrationsActionDraftUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch integrations action draft unauthorized response has a 5xx status code
func (o *PatchIntegrationsActionDraftUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch integrations action draft unauthorized response a status code equal to that given
func (o *PatchIntegrationsActionDraftUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchIntegrationsActionDraftUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchIntegrationsActionDraftUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchIntegrationsActionDraftUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionDraftForbidden creates a PatchIntegrationsActionDraftForbidden with default headers values
func NewPatchIntegrationsActionDraftForbidden() *PatchIntegrationsActionDraftForbidden {
	return &PatchIntegrationsActionDraftForbidden{}
}

/*
PatchIntegrationsActionDraftForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchIntegrationsActionDraftForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch integrations action draft forbidden response has a 2xx status code
func (o *PatchIntegrationsActionDraftForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch integrations action draft forbidden response has a 3xx status code
func (o *PatchIntegrationsActionDraftForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch integrations action draft forbidden response has a 4xx status code
func (o *PatchIntegrationsActionDraftForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch integrations action draft forbidden response has a 5xx status code
func (o *PatchIntegrationsActionDraftForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch integrations action draft forbidden response a status code equal to that given
func (o *PatchIntegrationsActionDraftForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchIntegrationsActionDraftForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftForbidden  %+v", 403, o.Payload)
}

func (o *PatchIntegrationsActionDraftForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftForbidden  %+v", 403, o.Payload)
}

func (o *PatchIntegrationsActionDraftForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionDraftNotFound creates a PatchIntegrationsActionDraftNotFound with default headers values
func NewPatchIntegrationsActionDraftNotFound() *PatchIntegrationsActionDraftNotFound {
	return &PatchIntegrationsActionDraftNotFound{}
}

/*
PatchIntegrationsActionDraftNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchIntegrationsActionDraftNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch integrations action draft not found response has a 2xx status code
func (o *PatchIntegrationsActionDraftNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch integrations action draft not found response has a 3xx status code
func (o *PatchIntegrationsActionDraftNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch integrations action draft not found response has a 4xx status code
func (o *PatchIntegrationsActionDraftNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch integrations action draft not found response has a 5xx status code
func (o *PatchIntegrationsActionDraftNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch integrations action draft not found response a status code equal to that given
func (o *PatchIntegrationsActionDraftNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchIntegrationsActionDraftNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftNotFound  %+v", 404, o.Payload)
}

func (o *PatchIntegrationsActionDraftNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftNotFound  %+v", 404, o.Payload)
}

func (o *PatchIntegrationsActionDraftNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionDraftRequestTimeout creates a PatchIntegrationsActionDraftRequestTimeout with default headers values
func NewPatchIntegrationsActionDraftRequestTimeout() *PatchIntegrationsActionDraftRequestTimeout {
	return &PatchIntegrationsActionDraftRequestTimeout{}
}

/*
PatchIntegrationsActionDraftRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchIntegrationsActionDraftRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch integrations action draft request timeout response has a 2xx status code
func (o *PatchIntegrationsActionDraftRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch integrations action draft request timeout response has a 3xx status code
func (o *PatchIntegrationsActionDraftRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch integrations action draft request timeout response has a 4xx status code
func (o *PatchIntegrationsActionDraftRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch integrations action draft request timeout response has a 5xx status code
func (o *PatchIntegrationsActionDraftRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch integrations action draft request timeout response a status code equal to that given
func (o *PatchIntegrationsActionDraftRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchIntegrationsActionDraftRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchIntegrationsActionDraftRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchIntegrationsActionDraftRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionDraftRequestEntityTooLarge creates a PatchIntegrationsActionDraftRequestEntityTooLarge with default headers values
func NewPatchIntegrationsActionDraftRequestEntityTooLarge() *PatchIntegrationsActionDraftRequestEntityTooLarge {
	return &PatchIntegrationsActionDraftRequestEntityTooLarge{}
}

/*
PatchIntegrationsActionDraftRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PatchIntegrationsActionDraftRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch integrations action draft request entity too large response has a 2xx status code
func (o *PatchIntegrationsActionDraftRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch integrations action draft request entity too large response has a 3xx status code
func (o *PatchIntegrationsActionDraftRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch integrations action draft request entity too large response has a 4xx status code
func (o *PatchIntegrationsActionDraftRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch integrations action draft request entity too large response has a 5xx status code
func (o *PatchIntegrationsActionDraftRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch integrations action draft request entity too large response a status code equal to that given
func (o *PatchIntegrationsActionDraftRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchIntegrationsActionDraftRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchIntegrationsActionDraftRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchIntegrationsActionDraftRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionDraftUnsupportedMediaType creates a PatchIntegrationsActionDraftUnsupportedMediaType with default headers values
func NewPatchIntegrationsActionDraftUnsupportedMediaType() *PatchIntegrationsActionDraftUnsupportedMediaType {
	return &PatchIntegrationsActionDraftUnsupportedMediaType{}
}

/*
PatchIntegrationsActionDraftUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchIntegrationsActionDraftUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch integrations action draft unsupported media type response has a 2xx status code
func (o *PatchIntegrationsActionDraftUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch integrations action draft unsupported media type response has a 3xx status code
func (o *PatchIntegrationsActionDraftUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch integrations action draft unsupported media type response has a 4xx status code
func (o *PatchIntegrationsActionDraftUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch integrations action draft unsupported media type response has a 5xx status code
func (o *PatchIntegrationsActionDraftUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch integrations action draft unsupported media type response a status code equal to that given
func (o *PatchIntegrationsActionDraftUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchIntegrationsActionDraftUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchIntegrationsActionDraftUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchIntegrationsActionDraftUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionDraftTooManyRequests creates a PatchIntegrationsActionDraftTooManyRequests with default headers values
func NewPatchIntegrationsActionDraftTooManyRequests() *PatchIntegrationsActionDraftTooManyRequests {
	return &PatchIntegrationsActionDraftTooManyRequests{}
}

/*
PatchIntegrationsActionDraftTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchIntegrationsActionDraftTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch integrations action draft too many requests response has a 2xx status code
func (o *PatchIntegrationsActionDraftTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch integrations action draft too many requests response has a 3xx status code
func (o *PatchIntegrationsActionDraftTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch integrations action draft too many requests response has a 4xx status code
func (o *PatchIntegrationsActionDraftTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch integrations action draft too many requests response has a 5xx status code
func (o *PatchIntegrationsActionDraftTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch integrations action draft too many requests response a status code equal to that given
func (o *PatchIntegrationsActionDraftTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchIntegrationsActionDraftTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchIntegrationsActionDraftTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchIntegrationsActionDraftTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionDraftInternalServerError creates a PatchIntegrationsActionDraftInternalServerError with default headers values
func NewPatchIntegrationsActionDraftInternalServerError() *PatchIntegrationsActionDraftInternalServerError {
	return &PatchIntegrationsActionDraftInternalServerError{}
}

/*
PatchIntegrationsActionDraftInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchIntegrationsActionDraftInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch integrations action draft internal server error response has a 2xx status code
func (o *PatchIntegrationsActionDraftInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch integrations action draft internal server error response has a 3xx status code
func (o *PatchIntegrationsActionDraftInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch integrations action draft internal server error response has a 4xx status code
func (o *PatchIntegrationsActionDraftInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch integrations action draft internal server error response has a 5xx status code
func (o *PatchIntegrationsActionDraftInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch integrations action draft internal server error response a status code equal to that given
func (o *PatchIntegrationsActionDraftInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchIntegrationsActionDraftInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchIntegrationsActionDraftInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchIntegrationsActionDraftInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionDraftServiceUnavailable creates a PatchIntegrationsActionDraftServiceUnavailable with default headers values
func NewPatchIntegrationsActionDraftServiceUnavailable() *PatchIntegrationsActionDraftServiceUnavailable {
	return &PatchIntegrationsActionDraftServiceUnavailable{}
}

/*
PatchIntegrationsActionDraftServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchIntegrationsActionDraftServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch integrations action draft service unavailable response has a 2xx status code
func (o *PatchIntegrationsActionDraftServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch integrations action draft service unavailable response has a 3xx status code
func (o *PatchIntegrationsActionDraftServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch integrations action draft service unavailable response has a 4xx status code
func (o *PatchIntegrationsActionDraftServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch integrations action draft service unavailable response has a 5xx status code
func (o *PatchIntegrationsActionDraftServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch integrations action draft service unavailable response a status code equal to that given
func (o *PatchIntegrationsActionDraftServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchIntegrationsActionDraftServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchIntegrationsActionDraftServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchIntegrationsActionDraftServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionDraftGatewayTimeout creates a PatchIntegrationsActionDraftGatewayTimeout with default headers values
func NewPatchIntegrationsActionDraftGatewayTimeout() *PatchIntegrationsActionDraftGatewayTimeout {
	return &PatchIntegrationsActionDraftGatewayTimeout{}
}

/*
PatchIntegrationsActionDraftGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchIntegrationsActionDraftGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch integrations action draft gateway timeout response has a 2xx status code
func (o *PatchIntegrationsActionDraftGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch integrations action draft gateway timeout response has a 3xx status code
func (o *PatchIntegrationsActionDraftGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch integrations action draft gateway timeout response has a 4xx status code
func (o *PatchIntegrationsActionDraftGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch integrations action draft gateway timeout response has a 5xx status code
func (o *PatchIntegrationsActionDraftGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch integrations action draft gateway timeout response a status code equal to that given
func (o *PatchIntegrationsActionDraftGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchIntegrationsActionDraftGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchIntegrationsActionDraftGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchIntegrationsActionDraftGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
