/*
	PoF - This file contains all the types for service registration.
*/

package types

import (
	"time"
)

type Service struct {
	ID                  string       `json:"id" rethinkdb:"id, omitempty"`
	Name                string       `json:"name" rethinkdb:"name"`
	Tags                []string     `json:"tags" rethinkdb:"tags"`
	HTTPMethod          string       `json:"http_method" rethinkdb:"http_method"`
	IPAddress           string       `json:"ip_address" rethinkdb:"ip_address"`
	Port                int          `json:"port" rethinkdb:"port"`
	RegisterTime        time.Time    `json:"register_time" rethinkdb:"register_time"`
	LastSyncTime        time.Time    `json:"last_sync_time" rethinkdb:"last_sync_time"`
	Endpoint            string       `json:"endpoint" rethinkdb:"endpoint"`                           // This is the endpoint that will be used by client to call a microservice.
	LoadBalancingMethod string       `json:"load_balancing_method" rethinkdb:"load_balancing_method"` // RoundRobin, LeastConnections, Resource, FixedWeighting
	TotalConnections    int          `json:"total_connections" rethinkdb:"total_connections"`
	DeRegisterAfter     int          `json:"de_register_after" rethinkdb:"de_register_after"`
	Status              string       `json:"health_status" rethinkdb:"health_status"` // active, unknown, down
	HealthConfig        HealthConfig `gorm:"foreignkey:ServiceID" json:"health_config"`
}

type HealthConfig struct {
	ServiceID           string                `gorm:"primary_key" json:"-"`
	HealthCheckEndpoint string                `json:"health_url"` // <ip_address>:<port>/health
	Interval            int                   `json:"interval"`
	Options             []HealthConfigOptions `gorm:"foreignkey:HealthConfigID" json:"options"`
}

type HTTPHeader struct {
	ServiceID string `gorm:"primary_key" json:"-"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}
type HealthConfigOptions struct {
	HealthConfigID     int          `gorm:"primary_key" json:"-"`
	Headers            []HTTPHeader `gorm:"many2many:health_config_options_headers;" json:"http_headers"`
	Body               any          `json:"body"` // This must be encoded into json.
	ExpectedStatusCode int          `json:"expected_status_code"`
}
