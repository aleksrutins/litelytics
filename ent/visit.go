// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/aleksrutins/litelytics/ent/site"
	"github.com/aleksrutins/litelytics/ent/visit"
)

// Visit is the model entity for the Visit schema.
type Visit struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Path holds the value of the "path" field.
	Path string `json:"path,omitempty"`
	// Referrer holds the value of the "referrer" field.
	Referrer string `json:"referrer,omitempty"`
	// Timestamp holds the value of the "timestamp" field.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// IP holds the value of the "ip" field.
	IP string `json:"ip,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VisitQuery when eager-loading is set.
	Edges       VisitEdges `json:"edges"`
	site_visits *int
}

// VisitEdges holds the relations/edges for other nodes in the graph.
type VisitEdges struct {
	// Site holds the value of the site edge.
	Site *Site `json:"site,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SiteOrErr returns the Site value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VisitEdges) SiteOrErr() (*Site, error) {
	if e.loadedTypes[0] {
		if e.Site == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: site.Label}
		}
		return e.Site, nil
	}
	return nil, &NotLoadedError{edge: "site"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Visit) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case visit.FieldID:
			values[i] = new(sql.NullInt64)
		case visit.FieldPath, visit.FieldReferrer, visit.FieldIP:
			values[i] = new(sql.NullString)
		case visit.FieldTimestamp:
			values[i] = new(sql.NullTime)
		case visit.ForeignKeys[0]: // site_visits
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Visit", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Visit fields.
func (v *Visit) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case visit.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			v.ID = int(value.Int64)
		case visit.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				v.Path = value.String
			}
		case visit.FieldReferrer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field referrer", values[i])
			} else if value.Valid {
				v.Referrer = value.String
			}
		case visit.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field timestamp", values[i])
			} else if value.Valid {
				v.Timestamp = value.Time
			}
		case visit.FieldIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip", values[i])
			} else if value.Valid {
				v.IP = value.String
			}
		case visit.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field site_visits", value)
			} else if value.Valid {
				v.site_visits = new(int)
				*v.site_visits = int(value.Int64)
			}
		}
	}
	return nil
}

// QuerySite queries the "site" edge of the Visit entity.
func (v *Visit) QuerySite() *SiteQuery {
	return (&VisitClient{config: v.config}).QuerySite(v)
}

// Update returns a builder for updating this Visit.
// Note that you need to call Visit.Unwrap() before calling this method if this Visit
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Visit) Update() *VisitUpdateOne {
	return (&VisitClient{config: v.config}).UpdateOne(v)
}

// Unwrap unwraps the Visit entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Visit) Unwrap() *Visit {
	_tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Visit is not a transactional entity")
	}
	v.config.driver = _tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Visit) String() string {
	var builder strings.Builder
	builder.WriteString("Visit(")
	builder.WriteString(fmt.Sprintf("id=%v, ", v.ID))
	builder.WriteString("path=")
	builder.WriteString(v.Path)
	builder.WriteString(", ")
	builder.WriteString("referrer=")
	builder.WriteString(v.Referrer)
	builder.WriteString(", ")
	builder.WriteString("timestamp=")
	builder.WriteString(v.Timestamp.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ip=")
	builder.WriteString(v.IP)
	builder.WriteByte(')')
	return builder.String()
}

// Visits is a parsable slice of Visit.
type Visits []*Visit

func (v Visits) config(cfg config) {
	for _i := range v {
		v[_i].config = cfg
	}
}
