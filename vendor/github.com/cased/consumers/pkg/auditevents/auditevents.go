package auditevents

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/segmentio/ksuid"
)

const (
	// DotCasedKey is the key name for the property that encodes information about
	// an AuditEvent.
	DotCasedKey = ".cased"

	// DotlessDotCasedKey is the key name used for mapping purposes in
	// Elasticsearch. DotlessDotCasedKey is intended to be sent in place of
	// DotCasedKey when indexing AuditEvents in Elasticsearch.
	DotlessDotCasedKey = "dotcased"
)

var (
	ErrImmutableKey = errors.New("key already exists")
)

// NewID generates a new AuditEvent identifier that encodes the time the
// AuditEvent originally took place.
//
// The time that is encoded in NewID is accurate to seconds.
func NewID(t time.Time) (string, error) {
	k, err := ksuid.NewRandomWithTime(t.UTC())
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("event_%s", k.String()), nil
}

// TimeFromID decodes the event ID and extracts the time it took place.
// Its accuracy is in seconds not nanoseconds.
func TimeFromID(id string) time.Time {
	parts := strings.Split(id, "event_")
	if len(parts) <= 1 {
		return time.Time{}
	}

	kuid, err := ksuid.Parse(parts[1])
	if err != nil {
		return time.Time{}
	}

	return kuid.Time()
}

// AuditEvent represents a JSON object that can be nested.
type AuditEvent map[string]interface{}

// FlattenedAuditEvent represents a nested JSON object flattened to one top
// level property without nesting.
type FlattenedAuditEvent map[string][]string

// DotCased is a reserved property in an audit event containing the original
// event, any modifications to the event post-processing, timestamps, and more.
type DotCased struct {
	PII                map[string][]SensitiveRange `json:"pii,omitempty"`
	Event              AuditEvent                  `json:"event,omitempty"`
	PublisherUserAgent string                      `json:"publisher_user_agent,omitempty"`
	ProcessedAt        time.Time                   `json:"processed_at"`
	ReceivedAt         time.Time                   `json:"received_at"`
	PublishedAt        time.Time                   `json:"published_at"`
}

// SensitiveRange is a range that informs Cased about any sensitive information
// stored in an AuditEvent.
type SensitiveRange struct {
	Begin int    `json:"begin"`
	End   int    `json:"end"`
	Label string `json:"label"`
}

// AuditEventPayload is the wrapper struct hosting the nestable JSON AuditEvent
// with the internal `.cased` property with a rich struct.
type AuditEventPayload struct {
	DotCased DotCased `json:".cased"`
	// TODO: Make this not accessible, it's immutable
	AuditEvent AuditEvent
}

// NewAuditEventPayload takes a JSON document and prepares a AuditEventPayload
// that is safe to use for internal processing and indexing to Elasticsearch.
func NewAuditEventPayload(bytes []byte) (*AuditEventPayload, error) {
	aep := &AuditEventPayload{}
	if err := json.Unmarshal(bytes, aep); err != nil {
		return nil, err
	}

	return aep, nil
}

// Set writes the given value to the mutable event within the .cased field.
func (aep *AuditEventPayload) Set(key string, value interface{}) error {
	// We have a value that exists here, we need to make sure it's not the same
	// value but marked as sensitive.
	if aep.AuditEvent[key] != nil {
		// Check to see if the value being set is a SensitiveValue. If it is and the
		// values are exactly the same, it's okay to replace the value so
		// MarshalJSON knows to build the event differently.
		if s, ok := value.(SensitiveValue); ok && aep.AuditEvent[key] == s.Value {
			if aep.DotCased.PII == nil {
				aep.DotCased.PII = map[string][]SensitiveRange{}
			}

			aep.DotCased.PII[key] = append(aep.DotCased.PII[key], s.Ranges...)

			return nil
		}

		if aep.AuditEvent[key] != value {
			return ErrImmutableKey
		}

		return nil
	}

	if aep.DotCased.Event == nil {
		aep.DotCased.Event = AuditEvent{}
	}

	// We have a new sensitive value
	if s, ok := value.(SensitiveValue); ok {
		if aep.DotCased.PII == nil {
			aep.DotCased.PII = map[string][]SensitiveRange{}
		}

		aep.DotCased.PII[key] = append(aep.DotCased.PII[key], s.Ranges...)
		aep.DotCased.Event[key] = s.Value

		return nil
	}

	aep.DotCased.Event[key] = value

	return nil
}

// MergedEvent merges the original audit event and any fields added to the audit
// event during processing as one top level object.
func (aep *AuditEventPayload) MergedEvent() AuditEvent {
	ae := AuditEvent{}
	for key, value := range aep.DotCased.Event {
		ae[key] = value
	}

	for key, value := range aep.AuditEvent {
		ae[key] = value
	}

	return ae
}

func (aep *AuditEventPayload) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{}

	for key, value := range aep.AuditEvent {
		m[key] = value
	}

	m[DotCasedKey] = aep.DotCased

	return json.Marshal(m)
}

func (aep *AuditEventPayload) UnmarshalJSON(data []byte) error {
	// We cannot use AuditEventPayload as it doesn't know how to serialize the
	// other properties into AuditEvent. This inline struct enables us to only
	// extract the `.cased` property.
	var dc struct {
		DotCased DotCased `json:".cased"`
	}
	if err := json.Unmarshal(data, &dc); err != nil {
		return err
	}
	aep.DotCased = dc.DotCased

	// Serialize all attributes to a common interface.
	var v map[string]interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	ae := AuditEvent{}
	for key, value := range v {
		// We've already extracted out the `.cased` property in the beginning of
		// UnmarshalJSON so we need to ignore it here.
		if key == DotCasedKey {
			continue
		}

		ae[key] = value
	}
	aep.AuditEvent = ae

	return nil
}

func (aep *AuditEventPayload) Indexable() (*IndexableAuditEvent, error) {
	return &IndexableAuditEvent{
		AuditEventPayload: *aep,
		FlattenedAuditEvent: FlattenedAuditEventPayload{
			DotCased:            aep.DotCased,
			FlattenedAuditEvent: Flatten(aep.MergedEvent()),
		},
	}, nil
}

type IndexableAuditEvent struct {
	AuditEventPayload   AuditEventPayload          `json:"audit_event"`
	FlattenedAuditEvent FlattenedAuditEventPayload `json:"flattened"`
}

type FlattenedAuditEventPayload struct {
	FlattenedAuditEvent
	DotCased DotCased `json:"dotcased"`
}

// MarshalJSON takes the nested FlattenedAuditEvent along with the DotCased to
// marshal it as a single top level JSON object.
func (p *FlattenedAuditEventPayload) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{}

	for key, value := range p.FlattenedAuditEvent {
		m[key] = value
	}

	m[DotlessDotCasedKey] = DotCased{
		PII:                p.DotCased.PII,
		ProcessedAt:        p.DotCased.ProcessedAt,
		PublishedAt:        p.DotCased.PublishedAt,
		ReceivedAt:         p.DotCased.ReceivedAt,
		PublisherUserAgent: p.DotCased.PublisherUserAgent,
	}

	return json.Marshal(m)
}
