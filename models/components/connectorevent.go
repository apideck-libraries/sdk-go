// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// EventSource - Unify event source
type EventSource string

const (
	EventSourceNative  EventSource = "native"
	EventSourceVirtual EventSource = "virtual"
)

func (e EventSource) ToPointer() *EventSource {
	return &e
}
func (e *EventSource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "native":
		fallthrough
	case "virtual":
		*e = EventSource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EventSource: %v", v)
	}
}

// ConnectorEvent - Unify event that is supported on the connector. Events are delivered via Webhooks.
type ConnectorEvent struct {
	// Unify event type
	EventType *string `json:"event_type,omitempty"`
	// Unify event source
	EventSource *EventSource `json:"event_source,omitempty"`
	// Downstream event type
	DownstreamEventType *string  `json:"downstream_event_type,omitempty"`
	Resources           []string `json:"resources,omitempty"`
	// Unify entity type
	EntityType *string `json:"entity_type,omitempty"`
}

func (o *ConnectorEvent) GetEventType() *string {
	if o == nil {
		return nil
	}
	return o.EventType
}

func (o *ConnectorEvent) GetEventSource() *EventSource {
	if o == nil {
		return nil
	}
	return o.EventSource
}

func (o *ConnectorEvent) GetDownstreamEventType() *string {
	if o == nil {
		return nil
	}
	return o.DownstreamEventType
}

func (o *ConnectorEvent) GetResources() []string {
	if o == nil {
		return nil
	}
	return o.Resources
}

func (o *ConnectorEvent) GetEntityType() *string {
	if o == nil {
		return nil
	}
	return o.EntityType
}
