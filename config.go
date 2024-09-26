/*
Copyright 2021 The routerd authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package firewalld

import (
	"context"
)

// Client for Firewalld org.fedoraproject.FirewallD1.config.
// Methods manipulate the persistent firewalld configuration.
type ConfigClient struct {
	conn       connection
	configPath caller
}

func NewConfigClient(conn connection) *ConfigClient {
	return &ConfigClient{
		conn:       conn,
		configPath: conn.Object(dbusDest, configPath),
	}
}

// FirewallD config prefix.
const firewalldConfig = "org.fedoraproject.FirewallD1.config."

func (c *ConfigClient) callWithReturn(ctx context.Context, method string, returnArg interface{}, args ...interface{}) error {
	call := newCall(firewalldConfig+method, 0).WithArguments(args...).WithReturns(returnArg)
	return c.configPath.Call(ctx, call)
}

// Return list of zone names (permanent configuration).
func (c *ConfigClient) GetZoneNames(
	ctx context.Context) ([]string, error) {
	var zoneNames []string
	return zoneNames, c.callWithReturn(ctx, "getZoneNames", &zoneNames)
}

// Return list of service names (permanent configuration).
func (c *ConfigClient) GetServiceNames(
	ctx context.Context) ([]string, error) {
	var serviceNames []string
	return serviceNames, c.callWithReturn(ctx, "getServiceNames", &serviceNames)
}

// List object paths of zones known to permanent environment.
func (c *ConfigClient) ListZones(
	ctx context.Context) (zonePaths []string, err error) {
	return zonePaths, c.callWithReturn(ctx, "listZones", &zonePaths)
}

// Return object path (permanent configuration) of zone with given name.
func (c *ConfigClient) GetZoneByName(
	ctx context.Context, zoneName string) (zone *Zone, err error) {
	var path string
	err = c.callWithReturn(ctx, "getZoneByName", &path, zoneName)
	if err != nil {
		return
	}
	zone = new(Zone)
	zone.Path = path
	zone.configPath = c.conn.Object(dbusDest, path)
	return
}

// Return object path (permanent configuration) of service with given name.
func (c *ConfigClient) GetServiceByName(
	ctx context.Context, serviceName string) (servicePath string, err error) {
	return servicePath, c.callWithReturn(ctx, "getServiceByName", &servicePath, serviceName)
}

// DEPRECATED: Add zone with given settings into permanent configuration.
// Needs https://github.com/godbus/dbus/pull/329 before this can fully function.
func (c *ConfigClient) AddZone(
	ctx context.Context, zoneName string, settings ZoneSettings) error {
	var z interface{}
	return c.callWithReturn(ctx, "addZone", &z, zoneName, settings.ToSlice())
}

// Add zone with given settings into permanent configuration.
// Needs https://github.com/godbus/dbus/pull/329 before this can fully function.
func (c *ConfigClient) AddZone2(
	ctx context.Context, zoneName string, settings ZoneSettings) error {
	var z interface{}
	return c.callWithReturn(ctx, "addZone2", &z, zoneName, settings.ToMap2())
}
