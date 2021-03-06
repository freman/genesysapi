// Code generated by go-swagger; DO NOT EDIT.

package telephony

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostTelephonySiptracesDownloadReader is a Reader for the PostTelephonySiptracesDownload structure.
type PostTelephonySiptracesDownloadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTelephonySiptracesDownloadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTelephonySiptracesDownloadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTelephonySiptracesDownloadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTelephonySiptracesDownloadUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTelephonySiptracesDownloadForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTelephonySiptracesDownloadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTelephonySiptracesDownloadRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTelephonySiptracesDownloadUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTelephonySiptracesDownloadTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTelephonySiptracesDownloadInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTelephonySiptracesDownloadServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTelephonySiptracesDownloadGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTelephonySiptracesDownloadOK creates a PostTelephonySiptracesDownloadOK with default headers values
func NewPostTelephonySiptracesDownloadOK() *PostTelephonySiptracesDownloadOK {
	return &PostTelephonySiptracesDownloadOK{}
}

/*PostTelephonySiptracesDownloadOK handles this case with default header values.

successful operation
*/
type PostTelephonySiptracesDownloadOK struct {
	Payload *models.SipDownloadResponse
}

func (o *PostTelephonySiptracesDownloadOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/siptraces/download][%d] postTelephonySiptracesDownloadOK  %+v", 200, o.Payload)
}

func (o *PostTelephonySiptracesDownloadOK) GetPayload() *models.SipDownloadResponse {
	return o.Payload
}

func (o *PostTelephonySiptracesDownloadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SipDownloadResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonySiptracesDownloadBadRequest creates a PostTelephonySiptracesDownloadBadRequest with default headers values
func NewPostTelephonySiptracesDownloadBadRequest() *PostTelephonySiptracesDownloadBadRequest {
	return &PostTelephonySiptracesDownloadBadRequest{}
}

/*PostTelephonySiptracesDownloadBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTelephonySiptracesDownloadBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonySiptracesDownloadBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/siptraces/download][%d] postTelephonySiptracesDownloadBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonySiptracesDownloadBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonySiptracesDownloadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonySiptracesDownloadUnauthorized creates a PostTelephonySiptracesDownloadUnauthorized with default headers values
func NewPostTelephonySiptracesDownloadUnauthorized() *PostTelephonySiptracesDownloadUnauthorized {
	return &PostTelephonySiptracesDownloadUnauthorized{}
}

/*PostTelephonySiptracesDownloadUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTelephonySiptracesDownloadUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonySiptracesDownloadUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/siptraces/download][%d] postTelephonySiptracesDownloadUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonySiptracesDownloadUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonySiptracesDownloadUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonySiptracesDownloadForbidden creates a PostTelephonySiptracesDownloadForbidden with default headers values
func NewPostTelephonySiptracesDownloadForbidden() *PostTelephonySiptracesDownloadForbidden {
	return &PostTelephonySiptracesDownloadForbidden{}
}

/*PostTelephonySiptracesDownloadForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostTelephonySiptracesDownloadForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonySiptracesDownloadForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/siptraces/download][%d] postTelephonySiptracesDownloadForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonySiptracesDownloadForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonySiptracesDownloadForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonySiptracesDownloadNotFound creates a PostTelephonySiptracesDownloadNotFound with default headers values
func NewPostTelephonySiptracesDownloadNotFound() *PostTelephonySiptracesDownloadNotFound {
	return &PostTelephonySiptracesDownloadNotFound{}
}

/*PostTelephonySiptracesDownloadNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostTelephonySiptracesDownloadNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonySiptracesDownloadNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/siptraces/download][%d] postTelephonySiptracesDownloadNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonySiptracesDownloadNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonySiptracesDownloadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonySiptracesDownloadRequestEntityTooLarge creates a PostTelephonySiptracesDownloadRequestEntityTooLarge with default headers values
func NewPostTelephonySiptracesDownloadRequestEntityTooLarge() *PostTelephonySiptracesDownloadRequestEntityTooLarge {
	return &PostTelephonySiptracesDownloadRequestEntityTooLarge{}
}

/*PostTelephonySiptracesDownloadRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostTelephonySiptracesDownloadRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonySiptracesDownloadRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/siptraces/download][%d] postTelephonySiptracesDownloadRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonySiptracesDownloadRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonySiptracesDownloadRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonySiptracesDownloadUnsupportedMediaType creates a PostTelephonySiptracesDownloadUnsupportedMediaType with default headers values
func NewPostTelephonySiptracesDownloadUnsupportedMediaType() *PostTelephonySiptracesDownloadUnsupportedMediaType {
	return &PostTelephonySiptracesDownloadUnsupportedMediaType{}
}

/*PostTelephonySiptracesDownloadUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTelephonySiptracesDownloadUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonySiptracesDownloadUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/siptraces/download][%d] postTelephonySiptracesDownloadUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonySiptracesDownloadUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonySiptracesDownloadUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonySiptracesDownloadTooManyRequests creates a PostTelephonySiptracesDownloadTooManyRequests with default headers values
func NewPostTelephonySiptracesDownloadTooManyRequests() *PostTelephonySiptracesDownloadTooManyRequests {
	return &PostTelephonySiptracesDownloadTooManyRequests{}
}

/*PostTelephonySiptracesDownloadTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostTelephonySiptracesDownloadTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonySiptracesDownloadTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/siptraces/download][%d] postTelephonySiptracesDownloadTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonySiptracesDownloadTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonySiptracesDownloadTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonySiptracesDownloadInternalServerError creates a PostTelephonySiptracesDownloadInternalServerError with default headers values
func NewPostTelephonySiptracesDownloadInternalServerError() *PostTelephonySiptracesDownloadInternalServerError {
	return &PostTelephonySiptracesDownloadInternalServerError{}
}

/*PostTelephonySiptracesDownloadInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTelephonySiptracesDownloadInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonySiptracesDownloadInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/siptraces/download][%d] postTelephonySiptracesDownloadInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonySiptracesDownloadInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonySiptracesDownloadInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonySiptracesDownloadServiceUnavailable creates a PostTelephonySiptracesDownloadServiceUnavailable with default headers values
func NewPostTelephonySiptracesDownloadServiceUnavailable() *PostTelephonySiptracesDownloadServiceUnavailable {
	return &PostTelephonySiptracesDownloadServiceUnavailable{}
}

/*PostTelephonySiptracesDownloadServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTelephonySiptracesDownloadServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonySiptracesDownloadServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/siptraces/download][%d] postTelephonySiptracesDownloadServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonySiptracesDownloadServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonySiptracesDownloadServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonySiptracesDownloadGatewayTimeout creates a PostTelephonySiptracesDownloadGatewayTimeout with default headers values
func NewPostTelephonySiptracesDownloadGatewayTimeout() *PostTelephonySiptracesDownloadGatewayTimeout {
	return &PostTelephonySiptracesDownloadGatewayTimeout{}
}

/*PostTelephonySiptracesDownloadGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostTelephonySiptracesDownloadGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonySiptracesDownloadGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/siptraces/download][%d] postTelephonySiptracesDownloadGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonySiptracesDownloadGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonySiptracesDownloadGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
