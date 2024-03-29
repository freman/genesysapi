// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PatchOutboundDnclistEmailaddressesReader is a Reader for the PatchOutboundDnclistEmailaddresses structure.
type PatchOutboundDnclistEmailaddressesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchOutboundDnclistEmailaddressesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPatchOutboundDnclistEmailaddressesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchOutboundDnclistEmailaddressesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchOutboundDnclistEmailaddressesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchOutboundDnclistEmailaddressesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchOutboundDnclistEmailaddressesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchOutboundDnclistEmailaddressesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchOutboundDnclistEmailaddressesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchOutboundDnclistEmailaddressesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchOutboundDnclistEmailaddressesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchOutboundDnclistEmailaddressesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewPatchOutboundDnclistEmailaddressesNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchOutboundDnclistEmailaddressesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchOutboundDnclistEmailaddressesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchOutboundDnclistEmailaddressesNoContent creates a PatchOutboundDnclistEmailaddressesNoContent with default headers values
func NewPatchOutboundDnclistEmailaddressesNoContent() *PatchOutboundDnclistEmailaddressesNoContent {
	return &PatchOutboundDnclistEmailaddressesNoContent{}
}

/*
PatchOutboundDnclistEmailaddressesNoContent describes a response with status code 204, with default header values.

No Content
*/
type PatchOutboundDnclistEmailaddressesNoContent struct {
}

// IsSuccess returns true when this patch outbound dnclist emailaddresses no content response has a 2xx status code
func (o *PatchOutboundDnclistEmailaddressesNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch outbound dnclist emailaddresses no content response has a 3xx status code
func (o *PatchOutboundDnclistEmailaddressesNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound dnclist emailaddresses no content response has a 4xx status code
func (o *PatchOutboundDnclistEmailaddressesNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch outbound dnclist emailaddresses no content response has a 5xx status code
func (o *PatchOutboundDnclistEmailaddressesNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this patch outbound dnclist emailaddresses no content response a status code equal to that given
func (o *PatchOutboundDnclistEmailaddressesNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *PatchOutboundDnclistEmailaddressesNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesNoContent ", 204)
}

func (o *PatchOutboundDnclistEmailaddressesNoContent) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesNoContent ", 204)
}

func (o *PatchOutboundDnclistEmailaddressesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchOutboundDnclistEmailaddressesBadRequest creates a PatchOutboundDnclistEmailaddressesBadRequest with default headers values
func NewPatchOutboundDnclistEmailaddressesBadRequest() *PatchOutboundDnclistEmailaddressesBadRequest {
	return &PatchOutboundDnclistEmailaddressesBadRequest{}
}

/*
PatchOutboundDnclistEmailaddressesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchOutboundDnclistEmailaddressesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound dnclist emailaddresses bad request response has a 2xx status code
func (o *PatchOutboundDnclistEmailaddressesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound dnclist emailaddresses bad request response has a 3xx status code
func (o *PatchOutboundDnclistEmailaddressesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound dnclist emailaddresses bad request response has a 4xx status code
func (o *PatchOutboundDnclistEmailaddressesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch outbound dnclist emailaddresses bad request response has a 5xx status code
func (o *PatchOutboundDnclistEmailaddressesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch outbound dnclist emailaddresses bad request response a status code equal to that given
func (o *PatchOutboundDnclistEmailaddressesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchOutboundDnclistEmailaddressesBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesBadRequest  %+v", 400, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesBadRequest  %+v", 400, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundDnclistEmailaddressesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOutboundDnclistEmailaddressesUnauthorized creates a PatchOutboundDnclistEmailaddressesUnauthorized with default headers values
func NewPatchOutboundDnclistEmailaddressesUnauthorized() *PatchOutboundDnclistEmailaddressesUnauthorized {
	return &PatchOutboundDnclistEmailaddressesUnauthorized{}
}

/*
PatchOutboundDnclistEmailaddressesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchOutboundDnclistEmailaddressesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound dnclist emailaddresses unauthorized response has a 2xx status code
func (o *PatchOutboundDnclistEmailaddressesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound dnclist emailaddresses unauthorized response has a 3xx status code
func (o *PatchOutboundDnclistEmailaddressesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound dnclist emailaddresses unauthorized response has a 4xx status code
func (o *PatchOutboundDnclistEmailaddressesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch outbound dnclist emailaddresses unauthorized response has a 5xx status code
func (o *PatchOutboundDnclistEmailaddressesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch outbound dnclist emailaddresses unauthorized response a status code equal to that given
func (o *PatchOutboundDnclistEmailaddressesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchOutboundDnclistEmailaddressesUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundDnclistEmailaddressesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOutboundDnclistEmailaddressesForbidden creates a PatchOutboundDnclistEmailaddressesForbidden with default headers values
func NewPatchOutboundDnclistEmailaddressesForbidden() *PatchOutboundDnclistEmailaddressesForbidden {
	return &PatchOutboundDnclistEmailaddressesForbidden{}
}

/*
PatchOutboundDnclistEmailaddressesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchOutboundDnclistEmailaddressesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound dnclist emailaddresses forbidden response has a 2xx status code
func (o *PatchOutboundDnclistEmailaddressesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound dnclist emailaddresses forbidden response has a 3xx status code
func (o *PatchOutboundDnclistEmailaddressesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound dnclist emailaddresses forbidden response has a 4xx status code
func (o *PatchOutboundDnclistEmailaddressesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch outbound dnclist emailaddresses forbidden response has a 5xx status code
func (o *PatchOutboundDnclistEmailaddressesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch outbound dnclist emailaddresses forbidden response a status code equal to that given
func (o *PatchOutboundDnclistEmailaddressesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchOutboundDnclistEmailaddressesForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesForbidden  %+v", 403, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesForbidden  %+v", 403, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundDnclistEmailaddressesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOutboundDnclistEmailaddressesNotFound creates a PatchOutboundDnclistEmailaddressesNotFound with default headers values
func NewPatchOutboundDnclistEmailaddressesNotFound() *PatchOutboundDnclistEmailaddressesNotFound {
	return &PatchOutboundDnclistEmailaddressesNotFound{}
}

/*
PatchOutboundDnclistEmailaddressesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchOutboundDnclistEmailaddressesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound dnclist emailaddresses not found response has a 2xx status code
func (o *PatchOutboundDnclistEmailaddressesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound dnclist emailaddresses not found response has a 3xx status code
func (o *PatchOutboundDnclistEmailaddressesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound dnclist emailaddresses not found response has a 4xx status code
func (o *PatchOutboundDnclistEmailaddressesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch outbound dnclist emailaddresses not found response has a 5xx status code
func (o *PatchOutboundDnclistEmailaddressesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch outbound dnclist emailaddresses not found response a status code equal to that given
func (o *PatchOutboundDnclistEmailaddressesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchOutboundDnclistEmailaddressesNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesNotFound  %+v", 404, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesNotFound  %+v", 404, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundDnclistEmailaddressesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOutboundDnclistEmailaddressesRequestTimeout creates a PatchOutboundDnclistEmailaddressesRequestTimeout with default headers values
func NewPatchOutboundDnclistEmailaddressesRequestTimeout() *PatchOutboundDnclistEmailaddressesRequestTimeout {
	return &PatchOutboundDnclistEmailaddressesRequestTimeout{}
}

/*
PatchOutboundDnclistEmailaddressesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchOutboundDnclistEmailaddressesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound dnclist emailaddresses request timeout response has a 2xx status code
func (o *PatchOutboundDnclistEmailaddressesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound dnclist emailaddresses request timeout response has a 3xx status code
func (o *PatchOutboundDnclistEmailaddressesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound dnclist emailaddresses request timeout response has a 4xx status code
func (o *PatchOutboundDnclistEmailaddressesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch outbound dnclist emailaddresses request timeout response has a 5xx status code
func (o *PatchOutboundDnclistEmailaddressesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch outbound dnclist emailaddresses request timeout response a status code equal to that given
func (o *PatchOutboundDnclistEmailaddressesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchOutboundDnclistEmailaddressesRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundDnclistEmailaddressesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOutboundDnclistEmailaddressesRequestEntityTooLarge creates a PatchOutboundDnclistEmailaddressesRequestEntityTooLarge with default headers values
func NewPatchOutboundDnclistEmailaddressesRequestEntityTooLarge() *PatchOutboundDnclistEmailaddressesRequestEntityTooLarge {
	return &PatchOutboundDnclistEmailaddressesRequestEntityTooLarge{}
}

/*
PatchOutboundDnclistEmailaddressesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PatchOutboundDnclistEmailaddressesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound dnclist emailaddresses request entity too large response has a 2xx status code
func (o *PatchOutboundDnclistEmailaddressesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound dnclist emailaddresses request entity too large response has a 3xx status code
func (o *PatchOutboundDnclistEmailaddressesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound dnclist emailaddresses request entity too large response has a 4xx status code
func (o *PatchOutboundDnclistEmailaddressesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch outbound dnclist emailaddresses request entity too large response has a 5xx status code
func (o *PatchOutboundDnclistEmailaddressesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch outbound dnclist emailaddresses request entity too large response a status code equal to that given
func (o *PatchOutboundDnclistEmailaddressesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchOutboundDnclistEmailaddressesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundDnclistEmailaddressesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOutboundDnclistEmailaddressesUnsupportedMediaType creates a PatchOutboundDnclistEmailaddressesUnsupportedMediaType with default headers values
func NewPatchOutboundDnclistEmailaddressesUnsupportedMediaType() *PatchOutboundDnclistEmailaddressesUnsupportedMediaType {
	return &PatchOutboundDnclistEmailaddressesUnsupportedMediaType{}
}

/*
PatchOutboundDnclistEmailaddressesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchOutboundDnclistEmailaddressesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound dnclist emailaddresses unsupported media type response has a 2xx status code
func (o *PatchOutboundDnclistEmailaddressesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound dnclist emailaddresses unsupported media type response has a 3xx status code
func (o *PatchOutboundDnclistEmailaddressesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound dnclist emailaddresses unsupported media type response has a 4xx status code
func (o *PatchOutboundDnclistEmailaddressesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch outbound dnclist emailaddresses unsupported media type response has a 5xx status code
func (o *PatchOutboundDnclistEmailaddressesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch outbound dnclist emailaddresses unsupported media type response a status code equal to that given
func (o *PatchOutboundDnclistEmailaddressesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchOutboundDnclistEmailaddressesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundDnclistEmailaddressesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOutboundDnclistEmailaddressesTooManyRequests creates a PatchOutboundDnclistEmailaddressesTooManyRequests with default headers values
func NewPatchOutboundDnclistEmailaddressesTooManyRequests() *PatchOutboundDnclistEmailaddressesTooManyRequests {
	return &PatchOutboundDnclistEmailaddressesTooManyRequests{}
}

/*
PatchOutboundDnclistEmailaddressesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchOutboundDnclistEmailaddressesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound dnclist emailaddresses too many requests response has a 2xx status code
func (o *PatchOutboundDnclistEmailaddressesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound dnclist emailaddresses too many requests response has a 3xx status code
func (o *PatchOutboundDnclistEmailaddressesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound dnclist emailaddresses too many requests response has a 4xx status code
func (o *PatchOutboundDnclistEmailaddressesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch outbound dnclist emailaddresses too many requests response has a 5xx status code
func (o *PatchOutboundDnclistEmailaddressesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch outbound dnclist emailaddresses too many requests response a status code equal to that given
func (o *PatchOutboundDnclistEmailaddressesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchOutboundDnclistEmailaddressesTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundDnclistEmailaddressesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOutboundDnclistEmailaddressesInternalServerError creates a PatchOutboundDnclistEmailaddressesInternalServerError with default headers values
func NewPatchOutboundDnclistEmailaddressesInternalServerError() *PatchOutboundDnclistEmailaddressesInternalServerError {
	return &PatchOutboundDnclistEmailaddressesInternalServerError{}
}

/*
PatchOutboundDnclistEmailaddressesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchOutboundDnclistEmailaddressesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound dnclist emailaddresses internal server error response has a 2xx status code
func (o *PatchOutboundDnclistEmailaddressesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound dnclist emailaddresses internal server error response has a 3xx status code
func (o *PatchOutboundDnclistEmailaddressesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound dnclist emailaddresses internal server error response has a 4xx status code
func (o *PatchOutboundDnclistEmailaddressesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch outbound dnclist emailaddresses internal server error response has a 5xx status code
func (o *PatchOutboundDnclistEmailaddressesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch outbound dnclist emailaddresses internal server error response a status code equal to that given
func (o *PatchOutboundDnclistEmailaddressesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchOutboundDnclistEmailaddressesInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundDnclistEmailaddressesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOutboundDnclistEmailaddressesNotImplemented creates a PatchOutboundDnclistEmailaddressesNotImplemented with default headers values
func NewPatchOutboundDnclistEmailaddressesNotImplemented() *PatchOutboundDnclistEmailaddressesNotImplemented {
	return &PatchOutboundDnclistEmailaddressesNotImplemented{}
}

/*
PatchOutboundDnclistEmailaddressesNotImplemented describes a response with status code 501, with default header values.

Not Implemented
*/
type PatchOutboundDnclistEmailaddressesNotImplemented struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound dnclist emailaddresses not implemented response has a 2xx status code
func (o *PatchOutboundDnclistEmailaddressesNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound dnclist emailaddresses not implemented response has a 3xx status code
func (o *PatchOutboundDnclistEmailaddressesNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound dnclist emailaddresses not implemented response has a 4xx status code
func (o *PatchOutboundDnclistEmailaddressesNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch outbound dnclist emailaddresses not implemented response has a 5xx status code
func (o *PatchOutboundDnclistEmailaddressesNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this patch outbound dnclist emailaddresses not implemented response a status code equal to that given
func (o *PatchOutboundDnclistEmailaddressesNotImplemented) IsCode(code int) bool {
	return code == 501
}

func (o *PatchOutboundDnclistEmailaddressesNotImplemented) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesNotImplemented  %+v", 501, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesNotImplemented) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesNotImplemented  %+v", 501, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesNotImplemented) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundDnclistEmailaddressesNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOutboundDnclistEmailaddressesServiceUnavailable creates a PatchOutboundDnclistEmailaddressesServiceUnavailable with default headers values
func NewPatchOutboundDnclistEmailaddressesServiceUnavailable() *PatchOutboundDnclistEmailaddressesServiceUnavailable {
	return &PatchOutboundDnclistEmailaddressesServiceUnavailable{}
}

/*
PatchOutboundDnclistEmailaddressesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchOutboundDnclistEmailaddressesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound dnclist emailaddresses service unavailable response has a 2xx status code
func (o *PatchOutboundDnclistEmailaddressesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound dnclist emailaddresses service unavailable response has a 3xx status code
func (o *PatchOutboundDnclistEmailaddressesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound dnclist emailaddresses service unavailable response has a 4xx status code
func (o *PatchOutboundDnclistEmailaddressesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch outbound dnclist emailaddresses service unavailable response has a 5xx status code
func (o *PatchOutboundDnclistEmailaddressesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch outbound dnclist emailaddresses service unavailable response a status code equal to that given
func (o *PatchOutboundDnclistEmailaddressesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchOutboundDnclistEmailaddressesServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundDnclistEmailaddressesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOutboundDnclistEmailaddressesGatewayTimeout creates a PatchOutboundDnclistEmailaddressesGatewayTimeout with default headers values
func NewPatchOutboundDnclistEmailaddressesGatewayTimeout() *PatchOutboundDnclistEmailaddressesGatewayTimeout {
	return &PatchOutboundDnclistEmailaddressesGatewayTimeout{}
}

/*
PatchOutboundDnclistEmailaddressesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchOutboundDnclistEmailaddressesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound dnclist emailaddresses gateway timeout response has a 2xx status code
func (o *PatchOutboundDnclistEmailaddressesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound dnclist emailaddresses gateway timeout response has a 3xx status code
func (o *PatchOutboundDnclistEmailaddressesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound dnclist emailaddresses gateway timeout response has a 4xx status code
func (o *PatchOutboundDnclistEmailaddressesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch outbound dnclist emailaddresses gateway timeout response has a 5xx status code
func (o *PatchOutboundDnclistEmailaddressesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch outbound dnclist emailaddresses gateway timeout response a status code equal to that given
func (o *PatchOutboundDnclistEmailaddressesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchOutboundDnclistEmailaddressesGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/dnclists/{dncListId}/emailaddresses][%d] patchOutboundDnclistEmailaddressesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchOutboundDnclistEmailaddressesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundDnclistEmailaddressesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
