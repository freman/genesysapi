// Code generated by go-swagger; DO NOT EDIT.

package content_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetContentmanagementSecurityprofileReader is a Reader for the GetContentmanagementSecurityprofile structure.
type GetContentmanagementSecurityprofileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContentmanagementSecurityprofileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContentmanagementSecurityprofileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetContentmanagementSecurityprofileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetContentmanagementSecurityprofileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetContentmanagementSecurityprofileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetContentmanagementSecurityprofileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetContentmanagementSecurityprofileRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetContentmanagementSecurityprofileRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetContentmanagementSecurityprofileUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetContentmanagementSecurityprofileTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetContentmanagementSecurityprofileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetContentmanagementSecurityprofileServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetContentmanagementSecurityprofileGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetContentmanagementSecurityprofileOK creates a GetContentmanagementSecurityprofileOK with default headers values
func NewGetContentmanagementSecurityprofileOK() *GetContentmanagementSecurityprofileOK {
	return &GetContentmanagementSecurityprofileOK{}
}

/*GetContentmanagementSecurityprofileOK handles this case with default header values.

successful operation
*/
type GetContentmanagementSecurityprofileOK struct {
	Payload *models.SecurityProfile
}

func (o *GetContentmanagementSecurityprofileOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/securityprofiles/{securityProfileId}][%d] getContentmanagementSecurityprofileOK  %+v", 200, o.Payload)
}

func (o *GetContentmanagementSecurityprofileOK) GetPayload() *models.SecurityProfile {
	return o.Payload
}

func (o *GetContentmanagementSecurityprofileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSecurityprofileBadRequest creates a GetContentmanagementSecurityprofileBadRequest with default headers values
func NewGetContentmanagementSecurityprofileBadRequest() *GetContentmanagementSecurityprofileBadRequest {
	return &GetContentmanagementSecurityprofileBadRequest{}
}

/*GetContentmanagementSecurityprofileBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetContentmanagementSecurityprofileBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementSecurityprofileBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/securityprofiles/{securityProfileId}][%d] getContentmanagementSecurityprofileBadRequest  %+v", 400, o.Payload)
}

func (o *GetContentmanagementSecurityprofileBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSecurityprofileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSecurityprofileUnauthorized creates a GetContentmanagementSecurityprofileUnauthorized with default headers values
func NewGetContentmanagementSecurityprofileUnauthorized() *GetContentmanagementSecurityprofileUnauthorized {
	return &GetContentmanagementSecurityprofileUnauthorized{}
}

/*GetContentmanagementSecurityprofileUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetContentmanagementSecurityprofileUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementSecurityprofileUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/securityprofiles/{securityProfileId}][%d] getContentmanagementSecurityprofileUnauthorized  %+v", 401, o.Payload)
}

func (o *GetContentmanagementSecurityprofileUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSecurityprofileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSecurityprofileForbidden creates a GetContentmanagementSecurityprofileForbidden with default headers values
func NewGetContentmanagementSecurityprofileForbidden() *GetContentmanagementSecurityprofileForbidden {
	return &GetContentmanagementSecurityprofileForbidden{}
}

/*GetContentmanagementSecurityprofileForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetContentmanagementSecurityprofileForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementSecurityprofileForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/securityprofiles/{securityProfileId}][%d] getContentmanagementSecurityprofileForbidden  %+v", 403, o.Payload)
}

func (o *GetContentmanagementSecurityprofileForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSecurityprofileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSecurityprofileNotFound creates a GetContentmanagementSecurityprofileNotFound with default headers values
func NewGetContentmanagementSecurityprofileNotFound() *GetContentmanagementSecurityprofileNotFound {
	return &GetContentmanagementSecurityprofileNotFound{}
}

/*GetContentmanagementSecurityprofileNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetContentmanagementSecurityprofileNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementSecurityprofileNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/securityprofiles/{securityProfileId}][%d] getContentmanagementSecurityprofileNotFound  %+v", 404, o.Payload)
}

func (o *GetContentmanagementSecurityprofileNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSecurityprofileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSecurityprofileRequestTimeout creates a GetContentmanagementSecurityprofileRequestTimeout with default headers values
func NewGetContentmanagementSecurityprofileRequestTimeout() *GetContentmanagementSecurityprofileRequestTimeout {
	return &GetContentmanagementSecurityprofileRequestTimeout{}
}

/*GetContentmanagementSecurityprofileRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetContentmanagementSecurityprofileRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementSecurityprofileRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/securityprofiles/{securityProfileId}][%d] getContentmanagementSecurityprofileRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetContentmanagementSecurityprofileRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSecurityprofileRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSecurityprofileRequestEntityTooLarge creates a GetContentmanagementSecurityprofileRequestEntityTooLarge with default headers values
func NewGetContentmanagementSecurityprofileRequestEntityTooLarge() *GetContentmanagementSecurityprofileRequestEntityTooLarge {
	return &GetContentmanagementSecurityprofileRequestEntityTooLarge{}
}

/*GetContentmanagementSecurityprofileRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetContentmanagementSecurityprofileRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementSecurityprofileRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/securityprofiles/{securityProfileId}][%d] getContentmanagementSecurityprofileRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetContentmanagementSecurityprofileRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSecurityprofileRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSecurityprofileUnsupportedMediaType creates a GetContentmanagementSecurityprofileUnsupportedMediaType with default headers values
func NewGetContentmanagementSecurityprofileUnsupportedMediaType() *GetContentmanagementSecurityprofileUnsupportedMediaType {
	return &GetContentmanagementSecurityprofileUnsupportedMediaType{}
}

/*GetContentmanagementSecurityprofileUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetContentmanagementSecurityprofileUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementSecurityprofileUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/securityprofiles/{securityProfileId}][%d] getContentmanagementSecurityprofileUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetContentmanagementSecurityprofileUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSecurityprofileUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSecurityprofileTooManyRequests creates a GetContentmanagementSecurityprofileTooManyRequests with default headers values
func NewGetContentmanagementSecurityprofileTooManyRequests() *GetContentmanagementSecurityprofileTooManyRequests {
	return &GetContentmanagementSecurityprofileTooManyRequests{}
}

/*GetContentmanagementSecurityprofileTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetContentmanagementSecurityprofileTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementSecurityprofileTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/securityprofiles/{securityProfileId}][%d] getContentmanagementSecurityprofileTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetContentmanagementSecurityprofileTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSecurityprofileTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSecurityprofileInternalServerError creates a GetContentmanagementSecurityprofileInternalServerError with default headers values
func NewGetContentmanagementSecurityprofileInternalServerError() *GetContentmanagementSecurityprofileInternalServerError {
	return &GetContentmanagementSecurityprofileInternalServerError{}
}

/*GetContentmanagementSecurityprofileInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetContentmanagementSecurityprofileInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementSecurityprofileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/securityprofiles/{securityProfileId}][%d] getContentmanagementSecurityprofileInternalServerError  %+v", 500, o.Payload)
}

func (o *GetContentmanagementSecurityprofileInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSecurityprofileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSecurityprofileServiceUnavailable creates a GetContentmanagementSecurityprofileServiceUnavailable with default headers values
func NewGetContentmanagementSecurityprofileServiceUnavailable() *GetContentmanagementSecurityprofileServiceUnavailable {
	return &GetContentmanagementSecurityprofileServiceUnavailable{}
}

/*GetContentmanagementSecurityprofileServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetContentmanagementSecurityprofileServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementSecurityprofileServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/securityprofiles/{securityProfileId}][%d] getContentmanagementSecurityprofileServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetContentmanagementSecurityprofileServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSecurityprofileServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSecurityprofileGatewayTimeout creates a GetContentmanagementSecurityprofileGatewayTimeout with default headers values
func NewGetContentmanagementSecurityprofileGatewayTimeout() *GetContentmanagementSecurityprofileGatewayTimeout {
	return &GetContentmanagementSecurityprofileGatewayTimeout{}
}

/*GetContentmanagementSecurityprofileGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetContentmanagementSecurityprofileGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementSecurityprofileGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/securityprofiles/{securityProfileId}][%d] getContentmanagementSecurityprofileGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetContentmanagementSecurityprofileGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSecurityprofileGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
