// Code generated by go-swagger; DO NOT EDIT.

package recording

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteRecordingMediaretentionpolicyReader is a Reader for the DeleteRecordingMediaretentionpolicy structure.
type DeleteRecordingMediaretentionpolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRecordingMediaretentionpolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRecordingMediaretentionpolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRecordingMediaretentionpolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRecordingMediaretentionpolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRecordingMediaretentionpolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRecordingMediaretentionpolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteRecordingMediaretentionpolicyRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteRecordingMediaretentionpolicyRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteRecordingMediaretentionpolicyUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteRecordingMediaretentionpolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRecordingMediaretentionpolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteRecordingMediaretentionpolicyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteRecordingMediaretentionpolicyGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRecordingMediaretentionpolicyOK creates a DeleteRecordingMediaretentionpolicyOK with default headers values
func NewDeleteRecordingMediaretentionpolicyOK() *DeleteRecordingMediaretentionpolicyOK {
	return &DeleteRecordingMediaretentionpolicyOK{}
}

/*DeleteRecordingMediaretentionpolicyOK handles this case with default header values.

Operation was successful.
*/
type DeleteRecordingMediaretentionpolicyOK struct {
}

func (o *DeleteRecordingMediaretentionpolicyOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/mediaretentionpolicies/{policyId}][%d] deleteRecordingMediaretentionpolicyOK ", 200)
}

func (o *DeleteRecordingMediaretentionpolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRecordingMediaretentionpolicyBadRequest creates a DeleteRecordingMediaretentionpolicyBadRequest with default headers values
func NewDeleteRecordingMediaretentionpolicyBadRequest() *DeleteRecordingMediaretentionpolicyBadRequest {
	return &DeleteRecordingMediaretentionpolicyBadRequest{}
}

/*DeleteRecordingMediaretentionpolicyBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteRecordingMediaretentionpolicyBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingMediaretentionpolicyBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/mediaretentionpolicies/{policyId}][%d] deleteRecordingMediaretentionpolicyBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRecordingMediaretentionpolicyBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingMediaretentionpolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingMediaretentionpolicyUnauthorized creates a DeleteRecordingMediaretentionpolicyUnauthorized with default headers values
func NewDeleteRecordingMediaretentionpolicyUnauthorized() *DeleteRecordingMediaretentionpolicyUnauthorized {
	return &DeleteRecordingMediaretentionpolicyUnauthorized{}
}

/*DeleteRecordingMediaretentionpolicyUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteRecordingMediaretentionpolicyUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingMediaretentionpolicyUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/mediaretentionpolicies/{policyId}][%d] deleteRecordingMediaretentionpolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRecordingMediaretentionpolicyUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingMediaretentionpolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingMediaretentionpolicyForbidden creates a DeleteRecordingMediaretentionpolicyForbidden with default headers values
func NewDeleteRecordingMediaretentionpolicyForbidden() *DeleteRecordingMediaretentionpolicyForbidden {
	return &DeleteRecordingMediaretentionpolicyForbidden{}
}

/*DeleteRecordingMediaretentionpolicyForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteRecordingMediaretentionpolicyForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingMediaretentionpolicyForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/mediaretentionpolicies/{policyId}][%d] deleteRecordingMediaretentionpolicyForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRecordingMediaretentionpolicyForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingMediaretentionpolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingMediaretentionpolicyNotFound creates a DeleteRecordingMediaretentionpolicyNotFound with default headers values
func NewDeleteRecordingMediaretentionpolicyNotFound() *DeleteRecordingMediaretentionpolicyNotFound {
	return &DeleteRecordingMediaretentionpolicyNotFound{}
}

/*DeleteRecordingMediaretentionpolicyNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteRecordingMediaretentionpolicyNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingMediaretentionpolicyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/mediaretentionpolicies/{policyId}][%d] deleteRecordingMediaretentionpolicyNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRecordingMediaretentionpolicyNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingMediaretentionpolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingMediaretentionpolicyRequestTimeout creates a DeleteRecordingMediaretentionpolicyRequestTimeout with default headers values
func NewDeleteRecordingMediaretentionpolicyRequestTimeout() *DeleteRecordingMediaretentionpolicyRequestTimeout {
	return &DeleteRecordingMediaretentionpolicyRequestTimeout{}
}

/*DeleteRecordingMediaretentionpolicyRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteRecordingMediaretentionpolicyRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingMediaretentionpolicyRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/mediaretentionpolicies/{policyId}][%d] deleteRecordingMediaretentionpolicyRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteRecordingMediaretentionpolicyRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingMediaretentionpolicyRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingMediaretentionpolicyRequestEntityTooLarge creates a DeleteRecordingMediaretentionpolicyRequestEntityTooLarge with default headers values
func NewDeleteRecordingMediaretentionpolicyRequestEntityTooLarge() *DeleteRecordingMediaretentionpolicyRequestEntityTooLarge {
	return &DeleteRecordingMediaretentionpolicyRequestEntityTooLarge{}
}

/*DeleteRecordingMediaretentionpolicyRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteRecordingMediaretentionpolicyRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingMediaretentionpolicyRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/mediaretentionpolicies/{policyId}][%d] deleteRecordingMediaretentionpolicyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteRecordingMediaretentionpolicyRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingMediaretentionpolicyRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingMediaretentionpolicyUnsupportedMediaType creates a DeleteRecordingMediaretentionpolicyUnsupportedMediaType with default headers values
func NewDeleteRecordingMediaretentionpolicyUnsupportedMediaType() *DeleteRecordingMediaretentionpolicyUnsupportedMediaType {
	return &DeleteRecordingMediaretentionpolicyUnsupportedMediaType{}
}

/*DeleteRecordingMediaretentionpolicyUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteRecordingMediaretentionpolicyUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingMediaretentionpolicyUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/mediaretentionpolicies/{policyId}][%d] deleteRecordingMediaretentionpolicyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteRecordingMediaretentionpolicyUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingMediaretentionpolicyUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingMediaretentionpolicyTooManyRequests creates a DeleteRecordingMediaretentionpolicyTooManyRequests with default headers values
func NewDeleteRecordingMediaretentionpolicyTooManyRequests() *DeleteRecordingMediaretentionpolicyTooManyRequests {
	return &DeleteRecordingMediaretentionpolicyTooManyRequests{}
}

/*DeleteRecordingMediaretentionpolicyTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteRecordingMediaretentionpolicyTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingMediaretentionpolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/mediaretentionpolicies/{policyId}][%d] deleteRecordingMediaretentionpolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRecordingMediaretentionpolicyTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingMediaretentionpolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingMediaretentionpolicyInternalServerError creates a DeleteRecordingMediaretentionpolicyInternalServerError with default headers values
func NewDeleteRecordingMediaretentionpolicyInternalServerError() *DeleteRecordingMediaretentionpolicyInternalServerError {
	return &DeleteRecordingMediaretentionpolicyInternalServerError{}
}

/*DeleteRecordingMediaretentionpolicyInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteRecordingMediaretentionpolicyInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingMediaretentionpolicyInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/mediaretentionpolicies/{policyId}][%d] deleteRecordingMediaretentionpolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRecordingMediaretentionpolicyInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingMediaretentionpolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingMediaretentionpolicyServiceUnavailable creates a DeleteRecordingMediaretentionpolicyServiceUnavailable with default headers values
func NewDeleteRecordingMediaretentionpolicyServiceUnavailable() *DeleteRecordingMediaretentionpolicyServiceUnavailable {
	return &DeleteRecordingMediaretentionpolicyServiceUnavailable{}
}

/*DeleteRecordingMediaretentionpolicyServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteRecordingMediaretentionpolicyServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingMediaretentionpolicyServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/mediaretentionpolicies/{policyId}][%d] deleteRecordingMediaretentionpolicyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteRecordingMediaretentionpolicyServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingMediaretentionpolicyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingMediaretentionpolicyGatewayTimeout creates a DeleteRecordingMediaretentionpolicyGatewayTimeout with default headers values
func NewDeleteRecordingMediaretentionpolicyGatewayTimeout() *DeleteRecordingMediaretentionpolicyGatewayTimeout {
	return &DeleteRecordingMediaretentionpolicyGatewayTimeout{}
}

/*DeleteRecordingMediaretentionpolicyGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteRecordingMediaretentionpolicyGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingMediaretentionpolicyGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/mediaretentionpolicies/{policyId}][%d] deleteRecordingMediaretentionpolicyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteRecordingMediaretentionpolicyGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingMediaretentionpolicyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
