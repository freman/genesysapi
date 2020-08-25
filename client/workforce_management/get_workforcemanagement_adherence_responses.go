// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetWorkforcemanagementAdherenceReader is a Reader for the GetWorkforcemanagementAdherence structure.
type GetWorkforcemanagementAdherenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementAdherenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementAdherenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementAdherenceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementAdherenceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementAdherenceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementAdherenceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementAdherenceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementAdherenceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementAdherenceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementAdherenceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementAdherenceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementAdherenceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementAdherenceOK creates a GetWorkforcemanagementAdherenceOK with default headers values
func NewGetWorkforcemanagementAdherenceOK() *GetWorkforcemanagementAdherenceOK {
	return &GetWorkforcemanagementAdherenceOK{}
}

/*GetWorkforcemanagementAdherenceOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementAdherenceOK struct {
	Payload []*models.UserScheduleAdherence
}

func (o *GetWorkforcemanagementAdherenceOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence][%d] getWorkforcemanagementAdherenceOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceOK) GetPayload() []*models.UserScheduleAdherence {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceBadRequest creates a GetWorkforcemanagementAdherenceBadRequest with default headers values
func NewGetWorkforcemanagementAdherenceBadRequest() *GetWorkforcemanagementAdherenceBadRequest {
	return &GetWorkforcemanagementAdherenceBadRequest{}
}

/*GetWorkforcemanagementAdherenceBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementAdherenceBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementAdherenceBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence][%d] getWorkforcemanagementAdherenceBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceUnauthorized creates a GetWorkforcemanagementAdherenceUnauthorized with default headers values
func NewGetWorkforcemanagementAdherenceUnauthorized() *GetWorkforcemanagementAdherenceUnauthorized {
	return &GetWorkforcemanagementAdherenceUnauthorized{}
}

/*GetWorkforcemanagementAdherenceUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementAdherenceUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementAdherenceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence][%d] getWorkforcemanagementAdherenceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceForbidden creates a GetWorkforcemanagementAdherenceForbidden with default headers values
func NewGetWorkforcemanagementAdherenceForbidden() *GetWorkforcemanagementAdherenceForbidden {
	return &GetWorkforcemanagementAdherenceForbidden{}
}

/*GetWorkforcemanagementAdherenceForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementAdherenceForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementAdherenceForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence][%d] getWorkforcemanagementAdherenceForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceNotFound creates a GetWorkforcemanagementAdherenceNotFound with default headers values
func NewGetWorkforcemanagementAdherenceNotFound() *GetWorkforcemanagementAdherenceNotFound {
	return &GetWorkforcemanagementAdherenceNotFound{}
}

/*GetWorkforcemanagementAdherenceNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementAdherenceNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementAdherenceNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence][%d] getWorkforcemanagementAdherenceNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceRequestEntityTooLarge creates a GetWorkforcemanagementAdherenceRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementAdherenceRequestEntityTooLarge() *GetWorkforcemanagementAdherenceRequestEntityTooLarge {
	return &GetWorkforcemanagementAdherenceRequestEntityTooLarge{}
}

/*GetWorkforcemanagementAdherenceRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetWorkforcemanagementAdherenceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementAdherenceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence][%d] getWorkforcemanagementAdherenceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceUnsupportedMediaType creates a GetWorkforcemanagementAdherenceUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementAdherenceUnsupportedMediaType() *GetWorkforcemanagementAdherenceUnsupportedMediaType {
	return &GetWorkforcemanagementAdherenceUnsupportedMediaType{}
}

/*GetWorkforcemanagementAdherenceUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementAdherenceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementAdherenceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence][%d] getWorkforcemanagementAdherenceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceTooManyRequests creates a GetWorkforcemanagementAdherenceTooManyRequests with default headers values
func NewGetWorkforcemanagementAdherenceTooManyRequests() *GetWorkforcemanagementAdherenceTooManyRequests {
	return &GetWorkforcemanagementAdherenceTooManyRequests{}
}

/*GetWorkforcemanagementAdherenceTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetWorkforcemanagementAdherenceTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementAdherenceTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence][%d] getWorkforcemanagementAdherenceTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceInternalServerError creates a GetWorkforcemanagementAdherenceInternalServerError with default headers values
func NewGetWorkforcemanagementAdherenceInternalServerError() *GetWorkforcemanagementAdherenceInternalServerError {
	return &GetWorkforcemanagementAdherenceInternalServerError{}
}

/*GetWorkforcemanagementAdherenceInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementAdherenceInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementAdherenceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence][%d] getWorkforcemanagementAdherenceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceServiceUnavailable creates a GetWorkforcemanagementAdherenceServiceUnavailable with default headers values
func NewGetWorkforcemanagementAdherenceServiceUnavailable() *GetWorkforcemanagementAdherenceServiceUnavailable {
	return &GetWorkforcemanagementAdherenceServiceUnavailable{}
}

/*GetWorkforcemanagementAdherenceServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementAdherenceServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementAdherenceServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence][%d] getWorkforcemanagementAdherenceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementAdherenceGatewayTimeout creates a GetWorkforcemanagementAdherenceGatewayTimeout with default headers values
func NewGetWorkforcemanagementAdherenceGatewayTimeout() *GetWorkforcemanagementAdherenceGatewayTimeout {
	return &GetWorkforcemanagementAdherenceGatewayTimeout{}
}

/*GetWorkforcemanagementAdherenceGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementAdherenceGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementAdherenceGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/adherence][%d] getWorkforcemanagementAdherenceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementAdherenceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementAdherenceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}