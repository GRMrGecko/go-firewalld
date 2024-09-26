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

	"github.com/godbus/dbus/v5"
)

type Zone struct {
	Path       string
	configPath caller
}

// Zone configuration prefix.
const configZone = "org.fedoraproject.FirewallD1.config.zone."

func (z *Zone) callWithReturn(ctx context.Context, method string, returnArg interface{}, args ...interface{}) error {
	call := newCall(configZone+method, 0).WithArguments(args...).WithReturns(returnArg)
	return z.configPath.Call(ctx, call)
}

func (z *Zone) call(ctx context.Context, method string, args ...interface{}) error {
	call := newCall(configZone+method, 0).WithArguments(args...)
	return z.configPath.Call(ctx, call)
}

// DEPRECATED: Return permanent settings of given zone.
func (z *Zone) GetSettings(
	ctx context.Context) (ZoneSettings, error) {
	var zoneSettings []interface{}
	err := z.callWithReturn(ctx, "getSettings", &zoneSettings)
	if err != nil {
		return ZoneSettings{}, err
	}

	return ZoneSettingsFromSlice(zoneSettings), nil
}

// Return permanent settings of given zone.
func (z *Zone) GetSettings2(
	ctx context.Context) (ZoneSettings, error) {
	var zoneSettings map[string]dbus.Variant
	err := z.callWithReturn(ctx, "getSettings2", &zoneSettings)
	if err != nil {
		return ZoneSettings{}, err
	}

	return ZoneSettingsFromMap2(zoneSettings), nil
}

// Add zone with given settings into permanent configuration.
// Needs https://github.com/godbus/dbus/pull/329 before this can fully function.
func (z *Zone) Update(
	ctx context.Context, settings ZoneSettings) error {
	return z.call(ctx, "update", settings.ToSlice())
}

// Add zone with given settings into permanent configuration.
// Needs https://github.com/godbus/dbus/pull/329 before this can fully function.
func (z *Zone) Update2(
	ctx context.Context, settings ZoneSettings) error {
	return z.call(ctx, "update2", settings.ToMap2())
}

// Remove zone from permanent configuration.
func (z *Zone) Remove(
	ctx context.Context) error {
	return z.call(ctx, "remove")
}

// Add port forward to zone permanent configuration.
func (z *Zone) AddPortForward(
	ctx context.Context, port, protocol, toport, toaddr string) error {
	return z.call(ctx, "addForwardPort", port, protocol, toport, toaddr)
}

// Add ICMP block to zone permanent configuration.
func (z *Zone) AddIcmpBlock(
	ctx context.Context, icmptype string) error {
	return z.call(ctx, "addIcmpBlock", icmptype)
}

// Enable ICMP block inversion on zone permanent configuration.
func (z *Zone) AddIcmpBlockInversion(
	ctx context.Context) error {
	return z.call(ctx, "addIcmpBlockInversion")
}

// Add interface to zone permanent configuration.
func (z *Zone) AddInterface(
	ctx context.Context, ifname string) error {
	return z.call(ctx, "addInterface", ifname)
}

// Enable masquerade on zone permanent configuration.
func (z *Zone) AddMasquerade(
	ctx context.Context) error {
	return z.call(ctx, "addMasquerade")
}

// Add port to zone permanent configuration.
func (z *Zone) AddPort(
	ctx context.Context, port, protocol string) error {
	return z.call(ctx, "addPort", port, protocol)
}

// Add protocol to zone permanent configuration.
func (z *Zone) AddProtocol(
	ctx context.Context, protocol string) error {
	return z.call(ctx, "addProtocol", protocol)
}

// Add rich rule to zone permanent configuration.
func (z *Zone) AddRichRule(
	ctx context.Context, rule string) error {
	return z.call(ctx, "addRichRule", rule)
}

// Add service to zone permanent configuration.
func (z *Zone) AddService(
	ctx context.Context, service string) error {
	return z.call(ctx, "addService", service)
}

// Add source to zone permanent configuration.
func (z *Zone) AddSource(
	ctx context.Context, source string) error {
	return z.call(ctx, "addSource", source)
}

// Add source oirt to zone permanent configuration.
func (z *Zone) AddSourcePort(
	ctx context.Context, port, protocol string) error {
	return z.call(ctx, "addSourcePort", port, protocol)
}

func (z *Zone) RemovePortForward(
	ctx context.Context, port, protocol, toport, toaddr string) error {
	return z.call(ctx, "removeForwardPort", port, protocol, toport, toaddr)
}

// Remove ICMP block to zone permanent configuration.
func (z *Zone) RemoveIcmpBlock(
	ctx context.Context, icmptype string) error {
	return z.call(ctx, "removeIcmpBlock", icmptype)
}

// Disable ICMP block inversion on zone permanent configuration.
func (z *Zone) RemoveIcmpBlockInversion(
	ctx context.Context) error {
	return z.call(ctx, "removeIcmpBlockInversion")
}

// Remove interface to zone permanent configuration.
func (z *Zone) RemoveInterface(
	ctx context.Context, ifname string) error {
	return z.call(ctx, "removeInterface", ifname)
}

// Disable masquerade on zone permanent configuration.
func (z *Zone) RemoveMasquerade(
	ctx context.Context) error {
	return z.call(ctx, "removeMasquerade")
}

// Remove port to zone permanent configuration.
func (z *Zone) RemovePort(
	ctx context.Context, port, protocol string) error {
	return z.call(ctx, "removePort", port, protocol)
}

// Remove protocol to zone permanent configuration.
func (z *Zone) removeProtocol(
	ctx context.Context, protocol string) error {
	return z.call(ctx, "removeProtocol", protocol)
}

// Remove rich rule to zone permanent configuration.
func (z *Zone) RemoveRichRule(
	ctx context.Context, rule string) error {
	return z.call(ctx, "removeRichRule", rule)
}

// Remove service to zone permanent configuration.
func (z *Zone) RemoveService(
	ctx context.Context, service string) error {
	return z.call(ctx, "removeService", service)
}

// Remove source to zone permanent configuration.
func (z *Zone) RemoveSource(
	ctx context.Context, source string) error {
	return z.call(ctx, "removeSource", source)
}

// Remove source oirt to zone permanent configuration.
func (z *Zone) RemoveSourcePort(
	ctx context.Context, port, protocol string) error {
	return z.call(ctx, "removeSourcePort", port, protocol)
}

// Rename the zone.
func (z *Zone) Rename(
	ctx context.Context, name string) error {
	return z.call(ctx, "rename", name)
}

// Set description for the zone permanent configuration.
func (z *Zone) SetDescription(
	ctx context.Context, description string) error {
	return z.call(ctx, "setDescription", description)
}

// Set name for the zone permanent configuration.
func (z *Zone) SetName(
	ctx context.Context, short string) error {
	return z.call(ctx, "setShort", short)
}

// Set target for the zone permanent configuration.
func (z *Zone) SetTarget(
	ctx context.Context, target string) error {
	return z.call(ctx, "setTarget", target)
}

// Set target for the zone permanent configuration.
func (z *Zone) SetVersion(
	ctx context.Context, version string) error {
	return z.call(ctx, "setVersion", version)
}
