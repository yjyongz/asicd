//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//	 Unless required by applicable law or agreed to in writing, software
//	 distributed under the License is distributed on an "AS IS" BASIS,
//	 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	 See the License for the specific language governing permissions and
//	 limitations under the License.
//
//   This is a auto-generated file, please do not edit!
// _______   __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __  
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  | 
// |  |__   |  |     |  |__   \  V  /     |   (----  \   \/    \/   /  |  |  ---|  |---- |  ,---- |  |__|  | 
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   | 
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  | 
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__| 
//                                                                                                           
		
include "asicdInt.thrift"
namespace go asicdServices
typedef i32 int
typedef i16 uint16
struct IPv6Intf {
	1 : string IntfRef
	2 : string IpAddr
}
struct Vlan {
	1 : i32 VlanId
	2 : list<string> IntfList
	3 : list<string> UntagIntfList
}
struct IPv4IntfState {
	1 : string IntfRef
	2 : i32 IfIndex
	3 : string IpAddr
	4 : string OperState
	5 : i32 NumUpEvents
	6 : string LastUpEventTime
	7 : i32 NumDownEvents
	8 : string LastDownEventTime
	9 : string L2IntfType
	10 : i32 L2IntfId
}
struct IPv4IntfStateGetInfo {
	1: int StartIdx
	2: int EndIdx
	3: int Count
	4: bool More
	5: list<IPv4IntfState> IPv4IntfStateList
}
struct PortState {
	1 : string IntfRef
	2 : i32 IfIndex
	3 : string Name
	4 : string OperState
	5 : i32 NumUpEvents
	6 : string LastUpEventTime
	7 : i32 NumDownEvents
	8 : string LastDownEventTime
	9 : i32 Pvid
	10 : i64 IfInOctets
	11 : i64 IfInUcastPkts
	12 : i64 IfInDiscards
	13 : i64 IfInErrors
	14 : i64 IfInUnknownProtos
	15 : i64 IfOutOctets
	16 : i64 IfOutUcastPkts
	17 : i64 IfOutDiscards
	18 : i64 IfOutErrors
	19 : string ErrDisableReason
	20 : string PresentInHW
}
struct PortStateGetInfo {
	1: int StartIdx
	2: int EndIdx
	3: int Count
	4: bool More
	5: list<PortState> PortStateList
}
struct IPv6IntfState {
	1 : string IntfRef
	2 : i32 IfIndex
	3 : string IpAddr
	4 : string OperState
	5 : i32 NumUpEvents
	6 : string LastUpEventTime
	7 : i32 NumDownEvents
	8 : string LastDownEventTime
	9 : string L2IntfType
	10 : i32 L2IntfId
}
struct IPv6IntfStateGetInfo {
	1: int StartIdx
	2: int EndIdx
	3: int Count
	4: bool More
	5: list<IPv6IntfState> IPv6IntfStateList
}
struct SubIPv4Intf {
	1 : string IpAddr
	2 : string IntfRef
	3 : string Type
	4 : string MacAddr
	5 : bool Enable
}
struct IPv4Intf {
	1 : string IntfRef
	2 : string IpAddr
}
struct VlanState {
	1 : i32 VlanId
	2 : string VlanName
	3 : string OperState
	4 : i32 IfIndex
}
struct VlanStateGetInfo {
	1: int StartIdx
	2: int EndIdx
	3: int Count
	4: bool More
	5: list<VlanState> VlanStateList
}
struct ArpEntryHwState {
	1 : string IpAddr
	2 : string MacAddr
	3 : string Vlan
	4 : string Port
}
struct ArpEntryHwStateGetInfo {
	1: int StartIdx
	2: int EndIdx
	3: int Count
	4: bool More
	5: list<ArpEntryHwState> ArpEntryHwStateList
}
struct LogicalIntf {
	1 : string Name
	2 : string Type
}
struct SubIPv6Intf {
	1 : string IpAddr
	2 : string IntfRef
	3 : string Type
	4 : string MacAddr
	5 : bool Enable
}
struct IPv4RouteHwState {
	1 : string DestinationNw
	2 : string NextHopIps
	3 : string RouteCreatedTime
	4 : string RouteUpdatedTime
}
struct IPv4RouteHwStateGetInfo {
	1: int StartIdx
	2: int EndIdx
	3: int Count
	4: bool More
	5: list<IPv4RouteHwState> IPv4RouteHwStateList
}
struct Port {
	1 : string IntfRef
	2 : i32 IfIndex
	3 : string Description
	4 : string PhyIntfType
	5 : string AdminState
	6 : string MacAddr
	7 : i32 Speed
	8 : string Duplex
	9 : string Autoneg
	10 : string MediaType
	11 : i32 Mtu
	12 : string BreakOutMode
}
struct PortGetInfo {
	1: int StartIdx
	2: int EndIdx
	3: int Count
	4: bool More
	5: list<Port> PortList
}
struct MacTableEntryState {
	1 : string MacAddr
	2 : i32 VlanId
	3 : i32 Port
}
struct MacTableEntryStateGetInfo {
	1: int StartIdx
	2: int EndIdx
	3: int Count
	4: bool More
	5: list<MacTableEntryState> MacTableEntryStateList
}
struct LogicalIntfState {
	1 : string Name
	2 : i32 IfIndex
	3 : string SrcMac
	4 : string OperState
	5 : i64 IfInOctets
	6 : i64 IfInUcastPkts
	7 : i64 IfInDiscards
	8 : i64 IfInErrors
	9 : i64 IfInUnknownProtos
	10 : i64 IfOutOctets
	11 : i64 IfOutUcastPkts
	12 : i64 IfOutDiscards
	13 : i64 IfOutErrors
}
struct LogicalIntfStateGetInfo {
	1: int StartIdx
	2: int EndIdx
	3: int Count
	4: bool More
	5: list<LogicalIntfState> LogicalIntfStateList
}

struct PatchOpInfo {
    1 : string Op
    2 : string Path
    3 : string Value
}
			        
service ASICDServices extends asicdInt.ASICDINTServices {
	bool CreateIPv6Intf(1: IPv6Intf config);
	bool UpdateIPv6Intf(1: IPv6Intf origconfig, 2: IPv6Intf newconfig, 3: list<bool> attrset, 4: list<PatchOpInfo> op);
	bool DeleteIPv6Intf(1: IPv6Intf config);

	bool CreateVlan(1: Vlan config);
	bool UpdateVlan(1: Vlan origconfig, 2: Vlan newconfig, 3: list<bool> attrset, 4: list<PatchOpInfo> op);
	bool DeleteVlan(1: Vlan config);

	IPv4IntfStateGetInfo GetBulkIPv4IntfState(1: int fromIndex, 2: int count);
	IPv4IntfState GetIPv4IntfState(1: string IntfRef);
	PortStateGetInfo GetBulkPortState(1: int fromIndex, 2: int count);
	PortState GetPortState(1: string IntfRef);
	IPv6IntfStateGetInfo GetBulkIPv6IntfState(1: int fromIndex, 2: int count);
	IPv6IntfState GetIPv6IntfState(1: string IntfRef);
	bool CreateSubIPv4Intf(1: SubIPv4Intf config);
	bool UpdateSubIPv4Intf(1: SubIPv4Intf origconfig, 2: SubIPv4Intf newconfig, 3: list<bool> attrset, 4: list<PatchOpInfo> op);
	bool DeleteSubIPv4Intf(1: SubIPv4Intf config);

	bool CreateIPv4Intf(1: IPv4Intf config);
	bool UpdateIPv4Intf(1: IPv4Intf origconfig, 2: IPv4Intf newconfig, 3: list<bool> attrset, 4: list<PatchOpInfo> op);
	bool DeleteIPv4Intf(1: IPv4Intf config);

	VlanStateGetInfo GetBulkVlanState(1: int fromIndex, 2: int count);
	VlanState GetVlanState(1: i32 VlanId);
	ArpEntryHwStateGetInfo GetBulkArpEntryHwState(1: int fromIndex, 2: int count);
	ArpEntryHwState GetArpEntryHwState(1: string IpAddr);
	bool CreateLogicalIntf(1: LogicalIntf config);
	bool UpdateLogicalIntf(1: LogicalIntf origconfig, 2: LogicalIntf newconfig, 3: list<bool> attrset, 4: list<PatchOpInfo> op);
	bool DeleteLogicalIntf(1: LogicalIntf config);

	bool CreateSubIPv6Intf(1: SubIPv6Intf config);
	bool UpdateSubIPv6Intf(1: SubIPv6Intf origconfig, 2: SubIPv6Intf newconfig, 3: list<bool> attrset, 4: list<PatchOpInfo> op);
	bool DeleteSubIPv6Intf(1: SubIPv6Intf config);

	IPv4RouteHwStateGetInfo GetBulkIPv4RouteHwState(1: int fromIndex, 2: int count);
	IPv4RouteHwState GetIPv4RouteHwState(1: string DestinationNw);
	bool CreatePort(1: Port config);
	bool UpdatePort(1: Port origconfig, 2: Port newconfig, 3: list<bool> attrset, 4: list<PatchOpInfo> op);
	bool DeletePort(1: Port config);

	PortGetInfo GetBulkPort(1: int fromIndex, 2: int count);
	Port GetPort(1: string IntfRef);
	MacTableEntryStateGetInfo GetBulkMacTableEntryState(1: int fromIndex, 2: int count);
	MacTableEntryState GetMacTableEntryState(1: string MacAddr);
	LogicalIntfStateGetInfo GetBulkLogicalIntfState(1: int fromIndex, 2: int count);
	LogicalIntfState GetLogicalIntfState(1: string Name);
}