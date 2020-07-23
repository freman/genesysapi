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

// GetVoicemailGroupPolicyReader is a Reader for the GetVoicemailGroupPolicy structure.
type GetVoicemailGroupPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVoicemailGroupPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVoicemailGroupPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVoicemailGroupPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetVoicemailGroupPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVoicemailGroupPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVoicemailGroupPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetVoicemailGroupPolicyRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetVoicemailGroupPolicyUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetVoicemailGroupPolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVoicemailGroupPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetVoicemailGroupPolicyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetVoicemailGroupPolicyGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVoicemailGroupPolicyOK creates a GetVoicemailGroupPolicyOK with default headers values
func NewGetVoicemailGroupPolicyOK() *GetVoicemailGroupPolicyOK {
	return &GetVoicemailGroupPolicyOK{}
}

/*GetVoicemailGroupPolicyOK handles this case with default header values.

successful operation
*/
type GetVoicemailGroupPolicyOK struct {
	Payload *models.VoicemailGroupPolicy
}

func (o *GetVoicemailGroupPolicyOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/policy][%d] getVoicemailGroupPolicyOK  %+v", 200, o.Payload)
}

func (o *GetVoicemailGroupPolicyOK) GetPayload() *models.VoicemailGroupPolicy {
	return o.Payload
}

func (o *GetVoicemailGroupPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VoicemailGroupPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupPolicyBadRequest creates a GetVoicemailGroupPolicyBadRequest with default headers values
func NewGetVoicemailGroupPolicyBadRequest() *GetVoicemailGroupPolicyBadRequest {
	return &GetVoicemailGroupPolicyBadRequest{}
}

/*GetVoicemailGroupPolicyBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetVoicemailGroupPolicyBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailGroupPolicyBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/policy][%d] getVoicemailGroupPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *GetVoicemailGroupPolicyBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupPolicyUnauthorized creates a GetVoicemailGroupPolicyUnauthorized with default headers values
func NewGetVoicemailGroupPolicyUnauthorized() *GetVoicemailGroupPolicyUnauthorized {
	return &GetVoicemailGroupPolicyUnauthorized{}
}

/*GetVoicemailGroupPolicyUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetVoicemailGroupPolicyUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailGroupPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/policy][%d] getVoicemailGroupPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *GetVoicemailGroupPolicyUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupPolicyForbidden creates a GetVoicemailGroupPolicyForbidden with default headers values
func NewGetVoicemailGroupPolicyForbidden() *GetVoicemailGroupPolicyForbidden {
	return &GetVoicemailGroupPolicyForbidden{}
}

/*GetVoicemailGroupPolicyForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetVoicemailGroupPolicyForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailGroupPolicyForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/policy][%d] getVoicemailGroupPolicyForbidden  %+v", 403, o.Payload)
}

func (o *GetVoicemailGroupPolicyForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupPolicyNotFound creates a GetVoicemailGroupPolicyNotFound with default headers values
func NewGetVoicemailGroupPolicyNotFound() *GetVoicemailGroupPolicyNotFound {
	return &GetVoicemailGroupPolicyNotFound{}
}

/*GetVoicemailGroupPolicyNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetVoicemailGroupPolicyNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailGroupPolicyNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/policy][%d] getVoicemailGroupPolicyNotFound  %+v", 404, o.Payload)
}

func (o *GetVoicemailGroupPolicyNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupPolicyRequestEntityTooLarge creates a GetVoicemailGroupPolicyRequestEntityTooLarge with default headers values
func NewGetVoicemailGroupPolicyRequestEntityTooLarge() *GetVoicemailGroupPolicyRequestEntityTooLarge {
	return &GetVoicemailGroupPolicyRequestEntityTooLarge{}
}

/*GetVoicemailGroupPolicyRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetVoicemailGroupPolicyRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailGroupPolicyRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/policy][%d] getVoicemailGroupPolicyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetVoicemailGroupPolicyRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupPolicyRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupPolicyUnsupportedMediaType creates a GetVoicemailGroupPolicyUnsupportedMediaType with default headers values
func NewGetVoicemailGroupPolicyUnsupportedMediaType() *GetVoicemailGroupPolicyUnsupportedMediaType {
	return &GetVoicemailGroupPolicyUnsupportedMediaType{}
}

/*GetVoicemailGroupPolicyUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetVoicemailGroupPolicyUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailGroupPolicyUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/policy][%d] getVoicemailGroupPolicyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetVoicemailGroupPolicyUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupPolicyUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupPolicyTooManyRequests creates a GetVoicemailGroupPolicyTooManyRequests with default headers values
func NewGetVoicemailGroupPolicyTooManyRequests() *GetVoicemailGroupPolicyTooManyRequests {
	return &GetVoicemailGroupPolicyTooManyRequests{}
}

/*GetVoicemailGroupPolicyTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetVoicemailGroupPolicyTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailGroupPolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/policy][%d] getVoicemailGroupPolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetVoicemailGroupPolicyTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupPolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupPolicyInternalServerError creates a GetVoicemailGroupPolicyInternalServerError with default headers values
func NewGetVoicemailGroupPolicyInternalServerError() *GetVoicemailGroupPolicyInternalServerError {
	return &GetVoicemailGroupPolicyInternalServerError{}
}

/*GetVoicemailGroupPolicyInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetVoicemailGroupPolicyInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailGroupPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/policy][%d] getVoicemailGroupPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVoicemailGroupPolicyInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupPolicyServiceUnavailable creates a GetVoicemailGroupPolicyServiceUnavailable with default headers values
func NewGetVoicemailGroupPolicyServiceUnavailable() *GetVoicemailGroupPolicyServiceUnavailable {
	return &GetVoicemailGroupPolicyServiceUnavailable{}
}

/*GetVoicemailGroupPolicyServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetVoicemailGroupPolicyServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailGroupPolicyServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/policy][%d] getVoicemailGroupPolicyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetVoicemailGroupPolicyServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupPolicyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupPolicyGatewayTimeout creates a GetVoicemailGroupPolicyGatewayTimeout with default headers values
func NewGetVoicemailGroupPolicyGatewayTimeout() *GetVoicemailGroupPolicyGatewayTimeout {
	return &GetVoicemailGroupPolicyGatewayTimeout{}
}

/*GetVoicemailGroupPolicyGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetVoicemailGroupPolicyGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailGroupPolicyGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/policy][%d] getVoicemailGroupPolicyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetVoicemailGroupPolicyGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupPolicyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
