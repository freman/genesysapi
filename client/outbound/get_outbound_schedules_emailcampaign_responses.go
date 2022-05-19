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

// GetOutboundSchedulesEmailcampaignReader is a Reader for the GetOutboundSchedulesEmailcampaign structure.
type GetOutboundSchedulesEmailcampaignReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundSchedulesEmailcampaignReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundSchedulesEmailcampaignOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundSchedulesEmailcampaignBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundSchedulesEmailcampaignUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundSchedulesEmailcampaignForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundSchedulesEmailcampaignNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundSchedulesEmailcampaignRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundSchedulesEmailcampaignRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundSchedulesEmailcampaignUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundSchedulesEmailcampaignTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundSchedulesEmailcampaignInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundSchedulesEmailcampaignServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundSchedulesEmailcampaignGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundSchedulesEmailcampaignOK creates a GetOutboundSchedulesEmailcampaignOK with default headers values
func NewGetOutboundSchedulesEmailcampaignOK() *GetOutboundSchedulesEmailcampaignOK {
	return &GetOutboundSchedulesEmailcampaignOK{}
}

/*GetOutboundSchedulesEmailcampaignOK handles this case with default header values.

successful operation
*/
type GetOutboundSchedulesEmailcampaignOK struct {
	Payload *models.EmailCampaignSchedule
}

func (o *GetOutboundSchedulesEmailcampaignOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/emailcampaigns/{emailCampaignId}][%d] getOutboundSchedulesEmailcampaignOK  %+v", 200, o.Payload)
}

func (o *GetOutboundSchedulesEmailcampaignOK) GetPayload() *models.EmailCampaignSchedule {
	return o.Payload
}

func (o *GetOutboundSchedulesEmailcampaignOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmailCampaignSchedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesEmailcampaignBadRequest creates a GetOutboundSchedulesEmailcampaignBadRequest with default headers values
func NewGetOutboundSchedulesEmailcampaignBadRequest() *GetOutboundSchedulesEmailcampaignBadRequest {
	return &GetOutboundSchedulesEmailcampaignBadRequest{}
}

/*GetOutboundSchedulesEmailcampaignBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundSchedulesEmailcampaignBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundSchedulesEmailcampaignBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/emailcampaigns/{emailCampaignId}][%d] getOutboundSchedulesEmailcampaignBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundSchedulesEmailcampaignBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesEmailcampaignBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesEmailcampaignUnauthorized creates a GetOutboundSchedulesEmailcampaignUnauthorized with default headers values
func NewGetOutboundSchedulesEmailcampaignUnauthorized() *GetOutboundSchedulesEmailcampaignUnauthorized {
	return &GetOutboundSchedulesEmailcampaignUnauthorized{}
}

/*GetOutboundSchedulesEmailcampaignUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundSchedulesEmailcampaignUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundSchedulesEmailcampaignUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/emailcampaigns/{emailCampaignId}][%d] getOutboundSchedulesEmailcampaignUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundSchedulesEmailcampaignUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesEmailcampaignUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesEmailcampaignForbidden creates a GetOutboundSchedulesEmailcampaignForbidden with default headers values
func NewGetOutboundSchedulesEmailcampaignForbidden() *GetOutboundSchedulesEmailcampaignForbidden {
	return &GetOutboundSchedulesEmailcampaignForbidden{}
}

/*GetOutboundSchedulesEmailcampaignForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundSchedulesEmailcampaignForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundSchedulesEmailcampaignForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/emailcampaigns/{emailCampaignId}][%d] getOutboundSchedulesEmailcampaignForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundSchedulesEmailcampaignForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesEmailcampaignForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesEmailcampaignNotFound creates a GetOutboundSchedulesEmailcampaignNotFound with default headers values
func NewGetOutboundSchedulesEmailcampaignNotFound() *GetOutboundSchedulesEmailcampaignNotFound {
	return &GetOutboundSchedulesEmailcampaignNotFound{}
}

/*GetOutboundSchedulesEmailcampaignNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundSchedulesEmailcampaignNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundSchedulesEmailcampaignNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/emailcampaigns/{emailCampaignId}][%d] getOutboundSchedulesEmailcampaignNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundSchedulesEmailcampaignNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesEmailcampaignNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesEmailcampaignRequestTimeout creates a GetOutboundSchedulesEmailcampaignRequestTimeout with default headers values
func NewGetOutboundSchedulesEmailcampaignRequestTimeout() *GetOutboundSchedulesEmailcampaignRequestTimeout {
	return &GetOutboundSchedulesEmailcampaignRequestTimeout{}
}

/*GetOutboundSchedulesEmailcampaignRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundSchedulesEmailcampaignRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundSchedulesEmailcampaignRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/emailcampaigns/{emailCampaignId}][%d] getOutboundSchedulesEmailcampaignRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundSchedulesEmailcampaignRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesEmailcampaignRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesEmailcampaignRequestEntityTooLarge creates a GetOutboundSchedulesEmailcampaignRequestEntityTooLarge with default headers values
func NewGetOutboundSchedulesEmailcampaignRequestEntityTooLarge() *GetOutboundSchedulesEmailcampaignRequestEntityTooLarge {
	return &GetOutboundSchedulesEmailcampaignRequestEntityTooLarge{}
}

/*GetOutboundSchedulesEmailcampaignRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetOutboundSchedulesEmailcampaignRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundSchedulesEmailcampaignRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/emailcampaigns/{emailCampaignId}][%d] getOutboundSchedulesEmailcampaignRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundSchedulesEmailcampaignRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesEmailcampaignRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesEmailcampaignUnsupportedMediaType creates a GetOutboundSchedulesEmailcampaignUnsupportedMediaType with default headers values
func NewGetOutboundSchedulesEmailcampaignUnsupportedMediaType() *GetOutboundSchedulesEmailcampaignUnsupportedMediaType {
	return &GetOutboundSchedulesEmailcampaignUnsupportedMediaType{}
}

/*GetOutboundSchedulesEmailcampaignUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundSchedulesEmailcampaignUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundSchedulesEmailcampaignUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/emailcampaigns/{emailCampaignId}][%d] getOutboundSchedulesEmailcampaignUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundSchedulesEmailcampaignUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesEmailcampaignUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesEmailcampaignTooManyRequests creates a GetOutboundSchedulesEmailcampaignTooManyRequests with default headers values
func NewGetOutboundSchedulesEmailcampaignTooManyRequests() *GetOutboundSchedulesEmailcampaignTooManyRequests {
	return &GetOutboundSchedulesEmailcampaignTooManyRequests{}
}

/*GetOutboundSchedulesEmailcampaignTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundSchedulesEmailcampaignTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundSchedulesEmailcampaignTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/emailcampaigns/{emailCampaignId}][%d] getOutboundSchedulesEmailcampaignTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundSchedulesEmailcampaignTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesEmailcampaignTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesEmailcampaignInternalServerError creates a GetOutboundSchedulesEmailcampaignInternalServerError with default headers values
func NewGetOutboundSchedulesEmailcampaignInternalServerError() *GetOutboundSchedulesEmailcampaignInternalServerError {
	return &GetOutboundSchedulesEmailcampaignInternalServerError{}
}

/*GetOutboundSchedulesEmailcampaignInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundSchedulesEmailcampaignInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundSchedulesEmailcampaignInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/emailcampaigns/{emailCampaignId}][%d] getOutboundSchedulesEmailcampaignInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundSchedulesEmailcampaignInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesEmailcampaignInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesEmailcampaignServiceUnavailable creates a GetOutboundSchedulesEmailcampaignServiceUnavailable with default headers values
func NewGetOutboundSchedulesEmailcampaignServiceUnavailable() *GetOutboundSchedulesEmailcampaignServiceUnavailable {
	return &GetOutboundSchedulesEmailcampaignServiceUnavailable{}
}

/*GetOutboundSchedulesEmailcampaignServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundSchedulesEmailcampaignServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundSchedulesEmailcampaignServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/emailcampaigns/{emailCampaignId}][%d] getOutboundSchedulesEmailcampaignServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundSchedulesEmailcampaignServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesEmailcampaignServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundSchedulesEmailcampaignGatewayTimeout creates a GetOutboundSchedulesEmailcampaignGatewayTimeout with default headers values
func NewGetOutboundSchedulesEmailcampaignGatewayTimeout() *GetOutboundSchedulesEmailcampaignGatewayTimeout {
	return &GetOutboundSchedulesEmailcampaignGatewayTimeout{}
}

/*GetOutboundSchedulesEmailcampaignGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundSchedulesEmailcampaignGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundSchedulesEmailcampaignGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/schedules/emailcampaigns/{emailCampaignId}][%d] getOutboundSchedulesEmailcampaignGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundSchedulesEmailcampaignGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundSchedulesEmailcampaignGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
