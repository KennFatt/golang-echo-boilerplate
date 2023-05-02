// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"myapp/ent/userdevice"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserDevice is the model entity for the UserDevice schema.
type UserDevice struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uint64 `json:"user_id,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// Platform holds the value of the "platform" field.
	Platform string `json:"platform,omitempty"`
	// LatestSkipUpdate holds the value of the "latest_skip_update" field.
	LatestSkipUpdate time.Time `json:"latest_skip_update,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeviceID holds the value of the "device_id" field.
	DeviceID     string `json:"device_id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserDevice) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userdevice.FieldID, userdevice.FieldUserID:
			values[i] = new(sql.NullInt64)
		case userdevice.FieldVersion, userdevice.FieldPlatform, userdevice.FieldDeviceID:
			values[i] = new(sql.NullString)
		case userdevice.FieldLatestSkipUpdate, userdevice.FieldCreatedAt, userdevice.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserDevice fields.
func (ud *UserDevice) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userdevice.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ud.ID = uint64(value.Int64)
		case userdevice.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ud.UserID = uint64(value.Int64)
			}
		case userdevice.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				ud.Version = value.String
			}
		case userdevice.FieldPlatform:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field platform", values[i])
			} else if value.Valid {
				ud.Platform = value.String
			}
		case userdevice.FieldLatestSkipUpdate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field latest_skip_update", values[i])
			} else if value.Valid {
				ud.LatestSkipUpdate = value.Time
			}
		case userdevice.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ud.CreatedAt = value.Time
			}
		case userdevice.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ud.UpdatedAt = value.Time
			}
		case userdevice.FieldDeviceID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device_id", values[i])
			} else if value.Valid {
				ud.DeviceID = value.String
			}
		default:
			ud.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserDevice.
// This includes values selected through modifiers, order, etc.
func (ud *UserDevice) Value(name string) (ent.Value, error) {
	return ud.selectValues.Get(name)
}

// Update returns a builder for updating this UserDevice.
// Note that you need to call UserDevice.Unwrap() before calling this method if this UserDevice
// was returned from a transaction, and the transaction was committed or rolled back.
func (ud *UserDevice) Update() *UserDeviceUpdateOne {
	return NewUserDeviceClient(ud.config).UpdateOne(ud)
}

// Unwrap unwraps the UserDevice entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ud *UserDevice) Unwrap() *UserDevice {
	_tx, ok := ud.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserDevice is not a transactional entity")
	}
	ud.config.driver = _tx.drv
	return ud
}

// String implements the fmt.Stringer.
func (ud *UserDevice) String() string {
	var builder strings.Builder
	builder.WriteString("UserDevice(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ud.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ud.UserID))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(ud.Version)
	builder.WriteString(", ")
	builder.WriteString("platform=")
	builder.WriteString(ud.Platform)
	builder.WriteString(", ")
	builder.WriteString("latest_skip_update=")
	builder.WriteString(ud.LatestSkipUpdate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ud.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ud.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("device_id=")
	builder.WriteString(ud.DeviceID)
	builder.WriteByte(')')
	return builder.String()
}

// UserDevices is a parsable slice of UserDevice.
type UserDevices []*UserDevice
