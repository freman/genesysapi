// Code generated by go-swagger; DO NOT EDIT.

package voicemail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutVoicemailPolicyReader is a Reader for the PutVoicemailPolicy structure.
type PutVoicemailPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutVoicemailPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutVoicemailPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutVoicemailPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutVoicemailPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutVoicemailPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutVoicemailPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutVoicemailPolicyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutVoicemailPolicyRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutVoicemailPolicyUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 424:
		result := NewPutVoicemailPolicyFailedDependency()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutVoicemailPolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutVoicemailPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutVoicemailPolicyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutVoicemailPolicyGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutVoicemailPolicyOK creates a PutVoicemailPolicyOK with default headers values
func NewPutVoicemailPolicyOK() *PutVoicemailPolicyOK {
	return &PutVoicemailPolicyOK{}
}

/*PutVoicemailPolicyOK handles this case with default header values.

successful operation
*/
type PutVoicemailPolicyOK struct {
	Payload *models.VoicemailOrganizationPolicy
}

func (o *PutVoicemailPolicyOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyOK  %+v", 200, o.Payload)
}

func (o *PutVoicemailPolicyOK) GetPayload() *models.VoicemailOrganizationPolicy {
	return o.Payload
}

func (o *PutVoicemailPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VoicemailOrganizationPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyBadRequest creates a PutVoicemailPolicyBadRequest with default headers values
func NewPutVoicemailPolicyBadRequest() *PutVoicemailPolicyBadRequest {
	return &PutVoicemailPolicyBadRequest{}
}

/*PutVoicemailPolicyBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutVoicemailPolicyBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutVoicemailPolicyBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *PutVoicemailPolicyBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyUnauthorized creates a PutVoicemailPolicyUnauthorized with default headers values
func NewPutVoicemailPolicyUnauthorized() *PutVoicemailPolicyUnauthorized {
	return &PutVoicemailPolicyUnauthorized{}
}

/*PutVoicemailPolicyUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutVoicemailPolicyUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutVoicemailPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *PutVoicemailPolicyUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyForbidden creates a PutVoicemailPolicyForbidden with default headers values
func NewPutVoicemailPolicyForbidden() *PutVoicemailPolicyForbidden {
	return &PutVoicemailPolicyForbidden{}
}

/*PutVoicemailPolicyForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutVoicemailPolicyForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutVoicemailPolicyForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyForbidden  %+v", 403, o.Payload)
}

func (o *PutVoicemailPolicyForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyNotFound creates a PutVoicemailPolicyNotFound with default headers values
func NewPutVoicemailPolicyNotFound() *PutVoicemailPolicyNotFound {
	return &PutVoicemailPolicyNotFound{}
}

/*PutVoicemailPolicyNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutVoicemailPolicyNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutVoicemailPolicyNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyNotFound  %+v", 404, o.Payload)
}

func (o *PutVoicemailPolicyNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyConflict creates a PutVoicemailPolicyConflict with default headers values
func NewPutVoicemailPolicyConflict() *PutVoicemailPolicyConflict {
	return &PutVoicemailPolicyConflict{}
}

/*PutVoicemailPolicyConflict handles this case with default header values.

Conflict
*/
type PutVoicemailPolicyConflict struct {
	Payload *models.ErrorBody
}

func (o *PutVoicemailPolicyConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyConflict  %+v", 409, o.Payload)
}

func (o *PutVoicemailPolicyConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyRequestEntityTooLarge creates a PutVoicemailPolicyRequestEntityTooLarge with default headers values
func NewPutVoicemailPolicyRequestEntityTooLarge() *PutVoicemailPolicyRequestEntityTooLarge {
	return &PutVoicemailPolicyRequestEntityTooLarge{}
}

/*PutVoicemailPolicyRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutVoicemailPolicyRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutVoicemailPolicyRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutVoicemailPolicyRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyUnsupportedMediaType creates a PutVoicemailPolicyUnsupportedMediaType with default headers values
func NewPutVoicemailPolicyUnsupportedMediaType() *PutVoicemailPolicyUnsupportedMediaType {
	return &PutVoicemailPolicyUnsupportedMediaType{}
}

/*PutVoicemailPolicyUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutVoicemailPolicyUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutVoicemailPolicyUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutVoicemailPolicyUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyFailedDependency creates a PutVoicemailPolicyFailedDependency with default headers values
func NewPutVoicemailPolicyFailedDependency() *PutVoicemailPolicyFailedDependency {
	return &PutVoicemailPolicyFailedDependency{}
}

/*PutVoicemailPolicyFailedDependency handles this case with default header values.

PutVoicemailPolicyFailedDependency put voicemail policy failed dependency
*/
type PutVoicemailPolicyFailedDependency struct {
	Payload *models.ErrorBody
}

func (o *PutVoicemailPolicyFailedDependency) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyFailedDependency  %+v", 424, o.Payload)
}

func (o *PutVoicemailPolicyFailedDependency) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyFailedDependency) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyTooManyRequests creates a PutVoicemailPolicyTooManyRequests with default headers values
func NewPutVoicemailPolicyTooManyRequests() *PutVoicemailPolicyTooManyRequests {
	return &PutVoicemailPolicyTooManyRequests{}
}

/*PutVoicemailPolicyTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutVoicemailPolicyTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutVoicemailPolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutVoicemailPolicyTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyInternalServerError creates a PutVoicemailPolicyInternalServerError with default headers values
func NewPutVoicemailPolicyInternalServerError() *PutVoicemailPolicyInternalServerError {
	return &PutVoicemailPolicyInternalServerError{}
}

/*PutVoicemailPolicyInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutVoicemailPolicyInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutVoicemailPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *PutVoicemailPolicyInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyServiceUnavailable creates a PutVoicemailPolicyServiceUnavailable with default headers values
func NewPutVoicemailPolicyServiceUnavailable() *PutVoicemailPolicyServiceUnavailable {
	return &PutVoicemailPolicyServiceUnavailable{}
}

/*PutVoicemailPolicyServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutVoicemailPolicyServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutVoicemailPolicyServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutVoicemailPolicyServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyGatewayTimeout creates a PutVoicemailPolicyGatewayTimeout with default headers values
func NewPutVoicemailPolicyGatewayTimeout() *PutVoicemailPolicyGatewayTimeout {
	return &PutVoicemailPolicyGatewayTimeout{}
}

/*PutVoicemailPolicyGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutVoicemailPolicyGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutVoicemailPolicyGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutVoicemailPolicyGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
