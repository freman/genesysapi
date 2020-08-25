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

// GetContentmanagementSecurityprofilesReader is a Reader for the GetContentmanagementSecurityprofiles structure.
type GetContentmanagementSecurityprofilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContentmanagementSecurityprofilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContentmanagementSecurityprofilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetContentmanagementSecurityprofilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetContentmanagementSecurityprofilesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetContentmanagementSecurityprofilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetContentmanagementSecurityprofilesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetContentmanagementSecurityprofilesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetContentmanagementSecurityprofilesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetContentmanagementSecurityprofilesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetContentmanagementSecurityprofilesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetContentmanagementSecurityprofilesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetContentmanagementSecurityprofilesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetContentmanagementSecurityprofilesOK creates a GetContentmanagementSecurityprofilesOK with default headers values
func NewGetContentmanagementSecurityprofilesOK() *GetContentmanagementSecurityprofilesOK {
	return &GetContentmanagementSecurityprofilesOK{}
}

/*GetContentmanagementSecurityprofilesOK handles this case with default header values.

successful operation
*/
type GetContentmanagementSecurityprofilesOK struct {
	Payload *models.SecurityProfileEntityListing
}

func (o *GetContentmanagementSecurityprofilesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/securityprofiles][%d] getContentmanagementSecurityprofilesOK  %+v", 200, o.Payload)
}

func (o *GetContentmanagementSecurityprofilesOK) GetPayload() *models.SecurityProfileEntityListing {
	return o.Payload
}

func (o *GetContentmanagementSecurityprofilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityProfileEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSecurityprofilesBadRequest creates a GetContentmanagementSecurityprofilesBadRequest with default headers values
func NewGetContentmanagementSecurityprofilesBadRequest() *GetContentmanagementSecurityprofilesBadRequest {
	return &GetContentmanagementSecurityprofilesBadRequest{}
}

/*GetContentmanagementSecurityprofilesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetContentmanagementSecurityprofilesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementSecurityprofilesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/securityprofiles][%d] getContentmanagementSecurityprofilesBadRequest  %+v", 400, o.Payload)
}

func (o *GetContentmanagementSecurityprofilesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSecurityprofilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSecurityprofilesUnauthorized creates a GetContentmanagementSecurityprofilesUnauthorized with default headers values
func NewGetContentmanagementSecurityprofilesUnauthorized() *GetContentmanagementSecurityprofilesUnauthorized {
	return &GetContentmanagementSecurityprofilesUnauthorized{}
}

/*GetContentmanagementSecurityprofilesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetContentmanagementSecurityprofilesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementSecurityprofilesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/securityprofiles][%d] getContentmanagementSecurityprofilesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetContentmanagementSecurityprofilesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSecurityprofilesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSecurityprofilesForbidden creates a GetContentmanagementSecurityprofilesForbidden with default headers values
func NewGetContentmanagementSecurityprofilesForbidden() *GetContentmanagementSecurityprofilesForbidden {
	return &GetContentmanagementSecurityprofilesForbidden{}
}

/*GetContentmanagementSecurityprofilesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetContentmanagementSecurityprofilesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementSecurityprofilesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/securityprofiles][%d] getContentmanagementSecurityprofilesForbidden  %+v", 403, o.Payload)
}

func (o *GetContentmanagementSecurityprofilesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSecurityprofilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSecurityprofilesNotFound creates a GetContentmanagementSecurityprofilesNotFound with default headers values
func NewGetContentmanagementSecurityprofilesNotFound() *GetContentmanagementSecurityprofilesNotFound {
	return &GetContentmanagementSecurityprofilesNotFound{}
}

/*GetContentmanagementSecurityprofilesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetContentmanagementSecurityprofilesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementSecurityprofilesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/securityprofiles][%d] getContentmanagementSecurityprofilesNotFound  %+v", 404, o.Payload)
}

func (o *GetContentmanagementSecurityprofilesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSecurityprofilesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSecurityprofilesRequestEntityTooLarge creates a GetContentmanagementSecurityprofilesRequestEntityTooLarge with default headers values
func NewGetContentmanagementSecurityprofilesRequestEntityTooLarge() *GetContentmanagementSecurityprofilesRequestEntityTooLarge {
	return &GetContentmanagementSecurityprofilesRequestEntityTooLarge{}
}

/*GetContentmanagementSecurityprofilesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetContentmanagementSecurityprofilesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementSecurityprofilesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/securityprofiles][%d] getContentmanagementSecurityprofilesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetContentmanagementSecurityprofilesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSecurityprofilesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSecurityprofilesUnsupportedMediaType creates a GetContentmanagementSecurityprofilesUnsupportedMediaType with default headers values
func NewGetContentmanagementSecurityprofilesUnsupportedMediaType() *GetContentmanagementSecurityprofilesUnsupportedMediaType {
	return &GetContentmanagementSecurityprofilesUnsupportedMediaType{}
}

/*GetContentmanagementSecurityprofilesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetContentmanagementSecurityprofilesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementSecurityprofilesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/securityprofiles][%d] getContentmanagementSecurityprofilesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetContentmanagementSecurityprofilesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSecurityprofilesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSecurityprofilesTooManyRequests creates a GetContentmanagementSecurityprofilesTooManyRequests with default headers values
func NewGetContentmanagementSecurityprofilesTooManyRequests() *GetContentmanagementSecurityprofilesTooManyRequests {
	return &GetContentmanagementSecurityprofilesTooManyRequests{}
}

/*GetContentmanagementSecurityprofilesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetContentmanagementSecurityprofilesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementSecurityprofilesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/securityprofiles][%d] getContentmanagementSecurityprofilesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetContentmanagementSecurityprofilesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSecurityprofilesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSecurityprofilesInternalServerError creates a GetContentmanagementSecurityprofilesInternalServerError with default headers values
func NewGetContentmanagementSecurityprofilesInternalServerError() *GetContentmanagementSecurityprofilesInternalServerError {
	return &GetContentmanagementSecurityprofilesInternalServerError{}
}

/*GetContentmanagementSecurityprofilesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetContentmanagementSecurityprofilesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementSecurityprofilesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/securityprofiles][%d] getContentmanagementSecurityprofilesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetContentmanagementSecurityprofilesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSecurityprofilesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSecurityprofilesServiceUnavailable creates a GetContentmanagementSecurityprofilesServiceUnavailable with default headers values
func NewGetContentmanagementSecurityprofilesServiceUnavailable() *GetContentmanagementSecurityprofilesServiceUnavailable {
	return &GetContentmanagementSecurityprofilesServiceUnavailable{}
}

/*GetContentmanagementSecurityprofilesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetContentmanagementSecurityprofilesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementSecurityprofilesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/securityprofiles][%d] getContentmanagementSecurityprofilesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetContentmanagementSecurityprofilesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSecurityprofilesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementSecurityprofilesGatewayTimeout creates a GetContentmanagementSecurityprofilesGatewayTimeout with default headers values
func NewGetContentmanagementSecurityprofilesGatewayTimeout() *GetContentmanagementSecurityprofilesGatewayTimeout {
	return &GetContentmanagementSecurityprofilesGatewayTimeout{}
}

/*GetContentmanagementSecurityprofilesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetContentmanagementSecurityprofilesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementSecurityprofilesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/securityprofiles][%d] getContentmanagementSecurityprofilesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetContentmanagementSecurityprofilesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementSecurityprofilesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}