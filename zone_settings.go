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
	"github.com/godbus/dbus/v5"
)

type ZoneSettings struct {
	Version         string
	Name            string
	Description     string
	Target          string
	Services        []string
	Ports           []Port
	ICMPBlocks      []string
	Masquerade      bool
	ForwardPorts    []ForwardPort
	Interfaces      []string
	SourceAddresses []string
	RichRules       []string
	Protocols       []string
	SourcePorts     []Port
	// Only compatible with version 2 of zone settings.
	ICMPBlockInversion bool
	Forwarded          bool
	EgressPriority     int32
	IngressPriority    int32
}

func ZoneSettingsFromSlice(s []interface{}) ZoneSettings {
	return ZoneSettings{
		Version:     s[0].(string),
		Name:        s[1].(string),
		Description: s[2].(string),
		// UNUSED s[3].(bool)
		Target:          s[4].(string),
		Services:        s[5].([]string),
		Ports:           interfaceSliceToPorts(s[6]),
		ICMPBlocks:      toStringSlice(s[7]),
		Masquerade:      s[8].(bool),
		ForwardPorts:    interfaceSliceToForwardPorts(s[9]),
		Interfaces:      s[10].([]string),
		SourceAddresses: s[11].([]string),
		RichRules:       s[12].([]string),
		Protocols:       s[13].([]string),
		SourcePorts:     interfaceSliceToPorts(s[14]),
	}
}

func (z *ZoneSettings) ToSlice() []interface{} {
	return []interface{}{
		z.Version,
		z.Name,
		z.Description,
		false, // UNUSED
		z.Target,
		z.Services,
		portsToInterfaceSlice(z.Ports),
		z.ICMPBlocks,
		z.Masquerade,
		forwardPortsToInterfaceSlice(z.ForwardPorts),
		z.Interfaces,
		z.SourceAddresses,
		z.RichRules,
		z.Protocols,
		portsToInterfaceSlice(z.SourcePorts),
		false, // needed but ???
	}
}

func ZoneSettingsFromMap2(s map[string]dbus.Variant) ZoneSettings {
	var zoneSettings ZoneSettings
	val, ok := s["version"]
	if ok {
		zoneSettings.Version = val.Value().(string)
	}
	val, ok = s["short"]
	if ok {
		zoneSettings.Name = val.Value().(string)
	}
	val, ok = s["description"]
	if ok {
		zoneSettings.Description = val.Value().(string)
	}
	val, ok = s["target"]
	if ok {
		zoneSettings.Target = val.Value().(string)
	}
	val, ok = s["services"]
	if ok {
		zoneSettings.Services = val.Value().([]string)
	}
	val, ok = s["ports"]
	if ok {
		zoneSettings.Ports = interfaceSliceToPorts(val.Value())
	}
	val, ok = s["icmp_blocks"]
	if ok {
		zoneSettings.ICMPBlocks = toStringSlice(val.Value())
	}
	val, ok = s["masquerade"]
	if ok {
		zoneSettings.Masquerade = val.Value().(bool)
	}
	val, ok = s["forward_ports"]
	if ok {
		zoneSettings.ForwardPorts = interfaceSliceToForwardPorts(val.Value())
	}
	val, ok = s["interfaces"]
	if ok {
		zoneSettings.Interfaces = val.Value().([]string)
	}
	val, ok = s["sources"]
	if ok {
		zoneSettings.SourceAddresses = val.Value().([]string)
	}
	val, ok = s["rules_str"]
	if ok {
		zoneSettings.RichRules = val.Value().([]string)
	}
	val, ok = s["protocols"]
	if ok {
		zoneSettings.Protocols = val.Value().([]string)
	}
	val, ok = s["source_ports"]
	if ok {
		zoneSettings.SourcePorts = interfaceSliceToPorts(val.Value())
	}
	val, ok = s["icmp_block_inversion"]
	if ok {
		zoneSettings.ICMPBlockInversion = val.Value().(bool)
	}
	val, ok = s["forward"]
	if ok {
		zoneSettings.Forwarded = val.Value().(bool)
	}
	val, ok = s["egress_priority"]
	if ok {
		zoneSettings.EgressPriority = val.Value().(int32)
	}
	val, ok = s["ingress_priority"]
	if ok {
		zoneSettings.IngressPriority = val.Value().(int32)
	}
	return zoneSettings
}

func (z *ZoneSettings) ToMap2() map[string]dbus.Variant {
	s := make(map[string]dbus.Variant)
	s["version"] = dbus.MakeVariant(z.Version)
	s["short"] = dbus.MakeVariant(z.Name)
	s["description"] = dbus.MakeVariant(z.Description)
	s["target"] = dbus.MakeVariant(z.Target)
	s["services"] = dbus.MakeVariant(z.Services)
	s["ports"] = dbus.MakeVariant(portsToInterfaceSlice(z.Ports))
	s["icmp_blocks"] = dbus.MakeVariant(z.ICMPBlocks)
	s["masquerade"] = dbus.MakeVariant(z.Masquerade)
	s["forward_ports"] = dbus.MakeVariant(forwardPortsToInterfaceSlice(z.ForwardPorts))
	s["interfaces"] = dbus.MakeVariant(z.Interfaces)
	s["sources"] = dbus.MakeVariant(z.SourceAddresses)
	s["rules_str"] = dbus.MakeVariant(z.RichRules)
	s["protocols"] = dbus.MakeVariant(z.Protocols)
	s["source_ports"] = dbus.MakeVariant(portsToInterfaceSlice(z.SourcePorts))
	s["icmp_block_inversion"] = dbus.MakeVariant(z.ICMPBlockInversion)
	s["forward"] = dbus.MakeVariant(z.Forwarded)
	s["egress_priority"] = dbus.MakeVariant(z.EgressPriority)
	s["ingress_priority"] = dbus.MakeVariant(z.IngressPriority)
	return s
}

type Port struct {
	Port     string
	Protocol string
}

func PortFromSlice(s []string) Port {
	return Port{
		Port:     s[0],
		Protocol: s[1],
	}
}

func (p *Port) ToSlice() []interface{} {
	return []interface{}{
		p.Port,
		p.Protocol,
	}
}

type ForwardPort struct {
	Port      string
	Protocol  string
	ToPort    string
	ToAddress string
}

func ForwardPortFromSlice(s []string) ForwardPort {
	return ForwardPort{
		Port:      s[0],
		Protocol:  s[1],
		ToPort:    s[2],
		ToAddress: s[3],
	}
}

func (p *ForwardPort) ToSlice() []interface{} {
	return []interface{}{
		p.Port,
		p.Protocol,
		p.ToPort,
		p.ToAddress,
	}
}

func interfaceSliceToPorts(s interface{}) []Port {
	var out []Port
	for _, p := range toStringSliceSlice(s) {
		out = append(out, PortFromSlice(p))
	}
	return out
}

func portsToInterfaceSlice(ports []Port) [][]interface{} {
	var out [][]interface{}
	for _, p := range ports {
		out = append(out, p.ToSlice())
	}
	return out
}

func interfaceSliceToForwardPorts(s interface{}) []ForwardPort {
	var out []ForwardPort
	for _, p := range toStringSliceSlice(s) {
		out = append(out, ForwardPort{
			Port:      p[0],
			Protocol:  p[1],
			ToPort:    p[2],
			ToAddress: p[3],
		})
	}
	return out
}

func forwardPortsToInterfaceSlice(ports []ForwardPort) [][]interface{} {
	var out [][]interface{}
	for _, p := range ports {
		out = append(out, p.ToSlice())
	}
	return out
}

func toStringSliceSlice(in interface{}) (out [][]string) {
	topSlice, ok := in.([][]interface{})
	if !ok {
		return nil
	}

	for _, slice := range topSlice {
		s := toStringSlice(slice)
		if len(s) == 0 {
			continue
		}
		out = append(out, s)
	}
	return
}

func toStringSlice(in interface{}) (out []string) {
	slice, ok := in.([]interface{})
	if !ok {
		return nil
	}

	for _, i := range slice {
		s, ok := i.(string)
		if !ok {
			continue
		}
		out = append(out, s)
	}
	return
}
