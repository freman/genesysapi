// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InitiatingAction initiating action
//
// swagger:model InitiatingAction
type InitiatingAction struct {

	// Action of the audit initiating the transaction
	// Enum: [Create View Update Move Copy Delete DeleteAll Fax VersionCreate Download Upload MemberAdd MemberUpdate MemberRemove ShareAdd ShareRemove TagAdd TagRemove TagUpdate Read ReadAll Execute ApplyProtection RevokeProtection UpdateRetention Abandon Archive RestoreRequest RestoreComplete Promote Publish Unpublish Activate Checkin Checkout Deactivate Debug Save Revert Transcode Enable Disable Authorize Deauthorize Authenticate ChangePassword Revoke Export Append Recycle Open Approved Rejected Rollback ImplementingChange ChangeImplemented ImplementingRollback RollbackImplemented Write Purge Processed Remove Replace UpdateInService UpdateOutOfService Cycle Scale IpAllowlistClear AddPairingRole Add Verify Assign Unassign Reassign Reschedule Cancel SoftDelete HardDelete Reset Rotate Restore Unarchive EnableCapture DownloadCapture]
	ActionContext string `json:"actionContext,omitempty"`

	// Id of the audit initiating the transaction
	TransactionID string `json:"transactionId,omitempty"`
}

// Validate validates this initiating action
func (m *InitiatingAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionContext(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var initiatingActionTypeActionContextPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Create","View","Update","Move","Copy","Delete","DeleteAll","Fax","VersionCreate","Download","Upload","MemberAdd","MemberUpdate","MemberRemove","ShareAdd","ShareRemove","TagAdd","TagRemove","TagUpdate","Read","ReadAll","Execute","ApplyProtection","RevokeProtection","UpdateRetention","Abandon","Archive","RestoreRequest","RestoreComplete","Promote","Publish","Unpublish","Activate","Checkin","Checkout","Deactivate","Debug","Save","Revert","Transcode","Enable","Disable","Authorize","Deauthorize","Authenticate","ChangePassword","Revoke","Export","Append","Recycle","Open","Approved","Rejected","Rollback","ImplementingChange","ChangeImplemented","ImplementingRollback","RollbackImplemented","Write","Purge","Processed","Remove","Replace","UpdateInService","UpdateOutOfService","Cycle","Scale","IpAllowlistClear","AddPairingRole","Add","Verify","Assign","Unassign","Reassign","Reschedule","Cancel","SoftDelete","HardDelete","Reset","Rotate","Restore","Unarchive","EnableCapture","DownloadCapture"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		initiatingActionTypeActionContextPropEnum = append(initiatingActionTypeActionContextPropEnum, v)
	}
}

const (

	// InitiatingActionActionContextCreate captures enum value "Create"
	InitiatingActionActionContextCreate string = "Create"

	// InitiatingActionActionContextView captures enum value "View"
	InitiatingActionActionContextView string = "View"

	// InitiatingActionActionContextUpdate captures enum value "Update"
	InitiatingActionActionContextUpdate string = "Update"

	// InitiatingActionActionContextMove captures enum value "Move"
	InitiatingActionActionContextMove string = "Move"

	// InitiatingActionActionContextCopy captures enum value "Copy"
	InitiatingActionActionContextCopy string = "Copy"

	// InitiatingActionActionContextDelete captures enum value "Delete"
	InitiatingActionActionContextDelete string = "Delete"

	// InitiatingActionActionContextDeleteAll captures enum value "DeleteAll"
	InitiatingActionActionContextDeleteAll string = "DeleteAll"

	// InitiatingActionActionContextFax captures enum value "Fax"
	InitiatingActionActionContextFax string = "Fax"

	// InitiatingActionActionContextVersionCreate captures enum value "VersionCreate"
	InitiatingActionActionContextVersionCreate string = "VersionCreate"

	// InitiatingActionActionContextDownload captures enum value "Download"
	InitiatingActionActionContextDownload string = "Download"

	// InitiatingActionActionContextUpload captures enum value "Upload"
	InitiatingActionActionContextUpload string = "Upload"

	// InitiatingActionActionContextMemberAdd captures enum value "MemberAdd"
	InitiatingActionActionContextMemberAdd string = "MemberAdd"

	// InitiatingActionActionContextMemberUpdate captures enum value "MemberUpdate"
	InitiatingActionActionContextMemberUpdate string = "MemberUpdate"

	// InitiatingActionActionContextMemberRemove captures enum value "MemberRemove"
	InitiatingActionActionContextMemberRemove string = "MemberRemove"

	// InitiatingActionActionContextShareAdd captures enum value "ShareAdd"
	InitiatingActionActionContextShareAdd string = "ShareAdd"

	// InitiatingActionActionContextShareRemove captures enum value "ShareRemove"
	InitiatingActionActionContextShareRemove string = "ShareRemove"

	// InitiatingActionActionContextTagAdd captures enum value "TagAdd"
	InitiatingActionActionContextTagAdd string = "TagAdd"

	// InitiatingActionActionContextTagRemove captures enum value "TagRemove"
	InitiatingActionActionContextTagRemove string = "TagRemove"

	// InitiatingActionActionContextTagUpdate captures enum value "TagUpdate"
	InitiatingActionActionContextTagUpdate string = "TagUpdate"

	// InitiatingActionActionContextRead captures enum value "Read"
	InitiatingActionActionContextRead string = "Read"

	// InitiatingActionActionContextReadAll captures enum value "ReadAll"
	InitiatingActionActionContextReadAll string = "ReadAll"

	// InitiatingActionActionContextExecute captures enum value "Execute"
	InitiatingActionActionContextExecute string = "Execute"

	// InitiatingActionActionContextApplyProtection captures enum value "ApplyProtection"
	InitiatingActionActionContextApplyProtection string = "ApplyProtection"

	// InitiatingActionActionContextRevokeProtection captures enum value "RevokeProtection"
	InitiatingActionActionContextRevokeProtection string = "RevokeProtection"

	// InitiatingActionActionContextUpdateRetention captures enum value "UpdateRetention"
	InitiatingActionActionContextUpdateRetention string = "UpdateRetention"

	// InitiatingActionActionContextAbandon captures enum value "Abandon"
	InitiatingActionActionContextAbandon string = "Abandon"

	// InitiatingActionActionContextArchive captures enum value "Archive"
	InitiatingActionActionContextArchive string = "Archive"

	// InitiatingActionActionContextRestoreRequest captures enum value "RestoreRequest"
	InitiatingActionActionContextRestoreRequest string = "RestoreRequest"

	// InitiatingActionActionContextRestoreComplete captures enum value "RestoreComplete"
	InitiatingActionActionContextRestoreComplete string = "RestoreComplete"

	// InitiatingActionActionContextPromote captures enum value "Promote"
	InitiatingActionActionContextPromote string = "Promote"

	// InitiatingActionActionContextPublish captures enum value "Publish"
	InitiatingActionActionContextPublish string = "Publish"

	// InitiatingActionActionContextUnpublish captures enum value "Unpublish"
	InitiatingActionActionContextUnpublish string = "Unpublish"

	// InitiatingActionActionContextActivate captures enum value "Activate"
	InitiatingActionActionContextActivate string = "Activate"

	// InitiatingActionActionContextCheckin captures enum value "Checkin"
	InitiatingActionActionContextCheckin string = "Checkin"

	// InitiatingActionActionContextCheckout captures enum value "Checkout"
	InitiatingActionActionContextCheckout string = "Checkout"

	// InitiatingActionActionContextDeactivate captures enum value "Deactivate"
	InitiatingActionActionContextDeactivate string = "Deactivate"

	// InitiatingActionActionContextDebug captures enum value "Debug"
	InitiatingActionActionContextDebug string = "Debug"

	// InitiatingActionActionContextSave captures enum value "Save"
	InitiatingActionActionContextSave string = "Save"

	// InitiatingActionActionContextRevert captures enum value "Revert"
	InitiatingActionActionContextRevert string = "Revert"

	// InitiatingActionActionContextTranscode captures enum value "Transcode"
	InitiatingActionActionContextTranscode string = "Transcode"

	// InitiatingActionActionContextEnable captures enum value "Enable"
	InitiatingActionActionContextEnable string = "Enable"

	// InitiatingActionActionContextDisable captures enum value "Disable"
	InitiatingActionActionContextDisable string = "Disable"

	// InitiatingActionActionContextAuthorize captures enum value "Authorize"
	InitiatingActionActionContextAuthorize string = "Authorize"

	// InitiatingActionActionContextDeauthorize captures enum value "Deauthorize"
	InitiatingActionActionContextDeauthorize string = "Deauthorize"

	// InitiatingActionActionContextAuthenticate captures enum value "Authenticate"
	InitiatingActionActionContextAuthenticate string = "Authenticate"

	// InitiatingActionActionContextChangePassword captures enum value "ChangePassword"
	InitiatingActionActionContextChangePassword string = "ChangePassword"

	// InitiatingActionActionContextRevoke captures enum value "Revoke"
	InitiatingActionActionContextRevoke string = "Revoke"

	// InitiatingActionActionContextExport captures enum value "Export"
	InitiatingActionActionContextExport string = "Export"

	// InitiatingActionActionContextAppend captures enum value "Append"
	InitiatingActionActionContextAppend string = "Append"

	// InitiatingActionActionContextRecycle captures enum value "Recycle"
	InitiatingActionActionContextRecycle string = "Recycle"

	// InitiatingActionActionContextOpen captures enum value "Open"
	InitiatingActionActionContextOpen string = "Open"

	// InitiatingActionActionContextApproved captures enum value "Approved"
	InitiatingActionActionContextApproved string = "Approved"

	// InitiatingActionActionContextRejected captures enum value "Rejected"
	InitiatingActionActionContextRejected string = "Rejected"

	// InitiatingActionActionContextRollback captures enum value "Rollback"
	InitiatingActionActionContextRollback string = "Rollback"

	// InitiatingActionActionContextImplementingChange captures enum value "ImplementingChange"
	InitiatingActionActionContextImplementingChange string = "ImplementingChange"

	// InitiatingActionActionContextChangeImplemented captures enum value "ChangeImplemented"
	InitiatingActionActionContextChangeImplemented string = "ChangeImplemented"

	// InitiatingActionActionContextImplementingRollback captures enum value "ImplementingRollback"
	InitiatingActionActionContextImplementingRollback string = "ImplementingRollback"

	// InitiatingActionActionContextRollbackImplemented captures enum value "RollbackImplemented"
	InitiatingActionActionContextRollbackImplemented string = "RollbackImplemented"

	// InitiatingActionActionContextWrite captures enum value "Write"
	InitiatingActionActionContextWrite string = "Write"

	// InitiatingActionActionContextPurge captures enum value "Purge"
	InitiatingActionActionContextPurge string = "Purge"

	// InitiatingActionActionContextProcessed captures enum value "Processed"
	InitiatingActionActionContextProcessed string = "Processed"

	// InitiatingActionActionContextRemove captures enum value "Remove"
	InitiatingActionActionContextRemove string = "Remove"

	// InitiatingActionActionContextReplace captures enum value "Replace"
	InitiatingActionActionContextReplace string = "Replace"

	// InitiatingActionActionContextUpdateInService captures enum value "UpdateInService"
	InitiatingActionActionContextUpdateInService string = "UpdateInService"

	// InitiatingActionActionContextUpdateOutOfService captures enum value "UpdateOutOfService"
	InitiatingActionActionContextUpdateOutOfService string = "UpdateOutOfService"

	// InitiatingActionActionContextCycle captures enum value "Cycle"
	InitiatingActionActionContextCycle string = "Cycle"

	// InitiatingActionActionContextScale captures enum value "Scale"
	InitiatingActionActionContextScale string = "Scale"

	// InitiatingActionActionContextIPAllowlistClear captures enum value "IpAllowlistClear"
	InitiatingActionActionContextIPAllowlistClear string = "IpAllowlistClear"

	// InitiatingActionActionContextAddPairingRole captures enum value "AddPairingRole"
	InitiatingActionActionContextAddPairingRole string = "AddPairingRole"

	// InitiatingActionActionContextAdd captures enum value "Add"
	InitiatingActionActionContextAdd string = "Add"

	// InitiatingActionActionContextVerify captures enum value "Verify"
	InitiatingActionActionContextVerify string = "Verify"

	// InitiatingActionActionContextAssign captures enum value "Assign"
	InitiatingActionActionContextAssign string = "Assign"

	// InitiatingActionActionContextUnassign captures enum value "Unassign"
	InitiatingActionActionContextUnassign string = "Unassign"

	// InitiatingActionActionContextReassign captures enum value "Reassign"
	InitiatingActionActionContextReassign string = "Reassign"

	// InitiatingActionActionContextReschedule captures enum value "Reschedule"
	InitiatingActionActionContextReschedule string = "Reschedule"

	// InitiatingActionActionContextCancel captures enum value "Cancel"
	InitiatingActionActionContextCancel string = "Cancel"

	// InitiatingActionActionContextSoftDelete captures enum value "SoftDelete"
	InitiatingActionActionContextSoftDelete string = "SoftDelete"

	// InitiatingActionActionContextHardDelete captures enum value "HardDelete"
	InitiatingActionActionContextHardDelete string = "HardDelete"

	// InitiatingActionActionContextReset captures enum value "Reset"
	InitiatingActionActionContextReset string = "Reset"

	// InitiatingActionActionContextRotate captures enum value "Rotate"
	InitiatingActionActionContextRotate string = "Rotate"

	// InitiatingActionActionContextRestore captures enum value "Restore"
	InitiatingActionActionContextRestore string = "Restore"

	// InitiatingActionActionContextUnarchive captures enum value "Unarchive"
	InitiatingActionActionContextUnarchive string = "Unarchive"

	// InitiatingActionActionContextEnableCapture captures enum value "EnableCapture"
	InitiatingActionActionContextEnableCapture string = "EnableCapture"

	// InitiatingActionActionContextDownloadCapture captures enum value "DownloadCapture"
	InitiatingActionActionContextDownloadCapture string = "DownloadCapture"
)

// prop value enum
func (m *InitiatingAction) validateActionContextEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, initiatingActionTypeActionContextPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InitiatingAction) validateActionContext(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionContext) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionContextEnum("actionContext", "body", m.ActionContext); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this initiating action based on context it is used
func (m *InitiatingAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InitiatingAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InitiatingAction) UnmarshalBinary(b []byte) error {
	var res InitiatingAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
