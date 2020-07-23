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

// GetTelephonySiptracesDownloadDownloadIDReader is a Reader for the GetTelephonySiptracesDownloadDownloadID structure.
type GetTelephonySiptracesDownloadDownloadIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonySiptracesDownloadDownloadIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonySiptracesDownloadDownloadIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewGetTelephonySiptracesDownloadDownloadIDAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonySiptracesDownloadDownloadIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonySiptracesDownloadDownloadIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonySiptracesDownloadDownloadIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonySiptracesDownloadDownloadIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonySiptracesDownloadDownloadIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonySiptracesDownloadDownloadIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonySiptracesDownloadDownloadIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonySiptracesDownloadDownloadIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonySiptracesDownloadDownloadIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonySiptracesDownloadDownloadIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonySiptracesDownloadDownloadIDOK creates a GetTelephonySiptracesDownloadDownloadIDOK with default headers values
func NewGetTelephonySiptracesDownloadDownloadIDOK() *GetTelephonySiptracesDownloadDownloadIDOK {
	return &GetTelephonySiptracesDownloadDownloadIDOK{}
}

/*GetTelephonySiptracesDownloadDownloadIDOK handles this case with default header values.

successful operation
*/
type GetTelephonySiptracesDownloadDownloadIDOK struct {
	Payload *models.SignedURLResponse
}

func (o *GetTelephonySiptracesDownloadDownloadIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/siptraces/download/{downloadId}][%d] getTelephonySiptracesDownloadDownloadIdOK  %+v", 200, o.Payload)
}

func (o *GetTelephonySiptracesDownloadDownloadIDOK) GetPayload() *models.SignedURLResponse {
	return o.Payload
}

func (o *GetTelephonySiptracesDownloadDownloadIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SignedURLResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonySiptracesDownloadDownloadIDAccepted creates a GetTelephonySiptracesDownloadDownloadIDAccepted with default headers values
func NewGetTelephonySiptracesDownloadDownloadIDAccepted() *GetTelephonySiptracesDownloadDownloadIDAccepted {
	return &GetTelephonySiptracesDownloadDownloadIDAccepted{}
}

/*GetTelephonySiptracesDownloadDownloadIDAccepted handles this case with default header values.

Request to download pcap file has been accepted
*/
type GetTelephonySiptracesDownloadDownloadIDAccepted struct {
	Payload *models.SignedURLResponse
}

func (o *GetTelephonySiptracesDownloadDownloadIDAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/siptraces/download/{downloadId}][%d] getTelephonySiptracesDownloadDownloadIdAccepted  %+v", 202, o.Payload)
}

func (o *GetTelephonySiptracesDownloadDownloadIDAccepted) GetPayload() *models.SignedURLResponse {
	return o.Payload
}

func (o *GetTelephonySiptracesDownloadDownloadIDAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SignedURLResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonySiptracesDownloadDownloadIDBadRequest creates a GetTelephonySiptracesDownloadDownloadIDBadRequest with default headers values
func NewGetTelephonySiptracesDownloadDownloadIDBadRequest() *GetTelephonySiptracesDownloadDownloadIDBadRequest {
	return &GetTelephonySiptracesDownloadDownloadIDBadRequest{}
}

/*GetTelephonySiptracesDownloadDownloadIDBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonySiptracesDownloadDownloadIDBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonySiptracesDownloadDownloadIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/siptraces/download/{downloadId}][%d] getTelephonySiptracesDownloadDownloadIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonySiptracesDownloadDownloadIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonySiptracesDownloadDownloadIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonySiptracesDownloadDownloadIDUnauthorized creates a GetTelephonySiptracesDownloadDownloadIDUnauthorized with default headers values
func NewGetTelephonySiptracesDownloadDownloadIDUnauthorized() *GetTelephonySiptracesDownloadDownloadIDUnauthorized {
	return &GetTelephonySiptracesDownloadDownloadIDUnauthorized{}
}

/*GetTelephonySiptracesDownloadDownloadIDUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonySiptracesDownloadDownloadIDUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonySiptracesDownloadDownloadIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/siptraces/download/{downloadId}][%d] getTelephonySiptracesDownloadDownloadIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonySiptracesDownloadDownloadIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonySiptracesDownloadDownloadIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonySiptracesDownloadDownloadIDForbidden creates a GetTelephonySiptracesDownloadDownloadIDForbidden with default headers values
func NewGetTelephonySiptracesDownloadDownloadIDForbidden() *GetTelephonySiptracesDownloadDownloadIDForbidden {
	return &GetTelephonySiptracesDownloadDownloadIDForbidden{}
}

/*GetTelephonySiptracesDownloadDownloadIDForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonySiptracesDownloadDownloadIDForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonySiptracesDownloadDownloadIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/siptraces/download/{downloadId}][%d] getTelephonySiptracesDownloadDownloadIdForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonySiptracesDownloadDownloadIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonySiptracesDownloadDownloadIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonySiptracesDownloadDownloadIDNotFound creates a GetTelephonySiptracesDownloadDownloadIDNotFound with default headers values
func NewGetTelephonySiptracesDownloadDownloadIDNotFound() *GetTelephonySiptracesDownloadDownloadIDNotFound {
	return &GetTelephonySiptracesDownloadDownloadIDNotFound{}
}

/*GetTelephonySiptracesDownloadDownloadIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonySiptracesDownloadDownloadIDNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonySiptracesDownloadDownloadIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/siptraces/download/{downloadId}][%d] getTelephonySiptracesDownloadDownloadIdNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonySiptracesDownloadDownloadIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonySiptracesDownloadDownloadIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonySiptracesDownloadDownloadIDRequestEntityTooLarge creates a GetTelephonySiptracesDownloadDownloadIDRequestEntityTooLarge with default headers values
func NewGetTelephonySiptracesDownloadDownloadIDRequestEntityTooLarge() *GetTelephonySiptracesDownloadDownloadIDRequestEntityTooLarge {
	return &GetTelephonySiptracesDownloadDownloadIDRequestEntityTooLarge{}
}

/*GetTelephonySiptracesDownloadDownloadIDRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonySiptracesDownloadDownloadIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonySiptracesDownloadDownloadIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/siptraces/download/{downloadId}][%d] getTelephonySiptracesDownloadDownloadIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonySiptracesDownloadDownloadIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonySiptracesDownloadDownloadIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonySiptracesDownloadDownloadIDUnsupportedMediaType creates a GetTelephonySiptracesDownloadDownloadIDUnsupportedMediaType with default headers values
func NewGetTelephonySiptracesDownloadDownloadIDUnsupportedMediaType() *GetTelephonySiptracesDownloadDownloadIDUnsupportedMediaType {
	return &GetTelephonySiptracesDownloadDownloadIDUnsupportedMediaType{}
}

/*GetTelephonySiptracesDownloadDownloadIDUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonySiptracesDownloadDownloadIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonySiptracesDownloadDownloadIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/siptraces/download/{downloadId}][%d] getTelephonySiptracesDownloadDownloadIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonySiptracesDownloadDownloadIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonySiptracesDownloadDownloadIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonySiptracesDownloadDownloadIDTooManyRequests creates a GetTelephonySiptracesDownloadDownloadIDTooManyRequests with default headers values
func NewGetTelephonySiptracesDownloadDownloadIDTooManyRequests() *GetTelephonySiptracesDownloadDownloadIDTooManyRequests {
	return &GetTelephonySiptracesDownloadDownloadIDTooManyRequests{}
}

/*GetTelephonySiptracesDownloadDownloadIDTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetTelephonySiptracesDownloadDownloadIDTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonySiptracesDownloadDownloadIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/siptraces/download/{downloadId}][%d] getTelephonySiptracesDownloadDownloadIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonySiptracesDownloadDownloadIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonySiptracesDownloadDownloadIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonySiptracesDownloadDownloadIDInternalServerError creates a GetTelephonySiptracesDownloadDownloadIDInternalServerError with default headers values
func NewGetTelephonySiptracesDownloadDownloadIDInternalServerError() *GetTelephonySiptracesDownloadDownloadIDInternalServerError {
	return &GetTelephonySiptracesDownloadDownloadIDInternalServerError{}
}

/*GetTelephonySiptracesDownloadDownloadIDInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonySiptracesDownloadDownloadIDInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonySiptracesDownloadDownloadIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/siptraces/download/{downloadId}][%d] getTelephonySiptracesDownloadDownloadIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonySiptracesDownloadDownloadIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonySiptracesDownloadDownloadIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonySiptracesDownloadDownloadIDServiceUnavailable creates a GetTelephonySiptracesDownloadDownloadIDServiceUnavailable with default headers values
func NewGetTelephonySiptracesDownloadDownloadIDServiceUnavailable() *GetTelephonySiptracesDownloadDownloadIDServiceUnavailable {
	return &GetTelephonySiptracesDownloadDownloadIDServiceUnavailable{}
}

/*GetTelephonySiptracesDownloadDownloadIDServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonySiptracesDownloadDownloadIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonySiptracesDownloadDownloadIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/siptraces/download/{downloadId}][%d] getTelephonySiptracesDownloadDownloadIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonySiptracesDownloadDownloadIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonySiptracesDownloadDownloadIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonySiptracesDownloadDownloadIDGatewayTimeout creates a GetTelephonySiptracesDownloadDownloadIDGatewayTimeout with default headers values
func NewGetTelephonySiptracesDownloadDownloadIDGatewayTimeout() *GetTelephonySiptracesDownloadDownloadIDGatewayTimeout {
	return &GetTelephonySiptracesDownloadDownloadIDGatewayTimeout{}
}

/*GetTelephonySiptracesDownloadDownloadIDGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonySiptracesDownloadDownloadIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonySiptracesDownloadDownloadIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/siptraces/download/{downloadId}][%d] getTelephonySiptracesDownloadDownloadIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonySiptracesDownloadDownloadIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonySiptracesDownloadDownloadIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
