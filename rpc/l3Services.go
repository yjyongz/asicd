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
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

// This file defines all interfaces provided for the L3 service
package rpc

import (
	"asicdInt"
	"asicdServices"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const RADD = 0
const RDEL = 1

const Gbroker = "sjc-ads-6624:9092"
const Gtopic = "test"

type IPv4RouteConfigWrapper struct {
	Ipv4Route []asicdInt.IPv4Route
	Operation int32
}

type IPv6RouteConfigWrapper struct {
	Ipv6Route []asicdInt.IPv6Route
	Operation int32
}

func OneWayHandleIPv4Kafka(broker string, topic string, nexthopInfo IPv4RouteConfigWrapper) error {
	lbroker := broker
	ltopic := topic

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": lbroker})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		return err
	}

	fmt.Printf("Created Producer %v\n", p)

	// Optional delivery channel, if not specified the Producer object's
	// .Events channel is used.
	deliveryChan := make(chan kafka.Event)

	value, _ := json.Marshal(nexthopInfo)
	err = p.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &ltopic, Partition: kafka.PartitionAny}, Value: []byte(value)}, deliveryChan)

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)
	return nil
}

func OneWayHandleIPv6Kafka(broker string, topic string, nexthopInfo IPv6RouteConfigWrapper) error {
	lbroker := broker
	ltopic := topic

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": lbroker})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		return err
	}

	fmt.Printf("Created Producer %v\n", p)

	// Optional delivery channel, if not specified the Producer object's
	// .Events channel is used.
	deliveryChan := make(chan kafka.Event)

	value, _ := json.Marshal(nexthopInfo)
	err = p.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &ltopic, Partition: kafka.PartitionAny}, Value: []byte(value)}, deliveryChan)

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)
	return nil
}

func OnewayHandleIPv6Route(ipv6RouteList []*asicdInt.IPv6Route, operation int32) *IPv6RouteConfigWrapper {
	nexthopInfo := new(IPv6RouteConfigWrapper)
	nexthopInfo.Operation = operation
	nexthopInfo.Ipv6Route = make([]asicdInt.IPv6Route, len(ipv6RouteList))
	for idx, route := range ipv6RouteList {
		nexthopInfo.Ipv6Route[idx].DestinationNw = route.DestinationNw
		nexthopInfo.Ipv6Route[idx].NetworkMask = route.NetworkMask
		nexthopInfo.Ipv6Route[idx].Protocol = route.Protocol
		//nexthopInfo[idx].Cost = route.Cost
		//nexthopInfo[idx].NullRoute = route.NullRoute
		nexthopInfo.Ipv6Route[idx].NextHopList = make([]*asicdInt.IPv6NextHop, len(route.NextHopList))
		for jdx, nexthop := range route.NextHopList {
			nexthopInfo.Ipv6Route[idx].NextHopList[jdx] = new(asicdInt.IPv6NextHop)
			nexthopInfo.Ipv6Route[idx].NextHopList[jdx].NextHopIp = nexthop.NextHopIp
			//nexthopInfo.Ipv4Route[idx].NextHop[jdx].NextHopIntRef = nexthop.NextHopIfType
			nexthopInfo.Ipv6Route[idx].NextHopList[jdx].Weight = nexthop.Weight
		}
	}
	return nexthopInfo
}

func OnewayHandleIPv4Route(ipv4RouteList []*asicdInt.IPv4Route, operation int32) *IPv4RouteConfigWrapper {
	nexthopInfo := new(IPv4RouteConfigWrapper)
	nexthopInfo.Operation = operation
	nexthopInfo.Ipv4Route = make([]asicdInt.IPv4Route, len(ipv4RouteList))
	for idx, route := range ipv4RouteList {
		nexthopInfo.Ipv4Route[idx].DestinationNw = route.DestinationNw
		nexthopInfo.Ipv4Route[idx].NetworkMask = route.NetworkMask
		nexthopInfo.Ipv4Route[idx].Protocol = route.Protocol
		nexthopInfo.Ipv4Route[idx].NextHopList = make([]*asicdInt.IPv4NextHop, len(route.NextHopList))
		for jdx, nexthop := range route.NextHopList {
			nexthopInfo.Ipv4Route[idx].NextHopList[jdx] = new(asicdInt.IPv4NextHop)
			nexthopInfo.Ipv4Route[idx].NextHopList[jdx].NextHopIp = nexthop.NextHopIp
			nexthopInfo.Ipv4Route[idx].NextHopList[jdx].Weight = nexthop.Weight
		}
	}
	return nexthopInfo
}

//Utility method to retrieve list of ifindex to ifname mapping
func (svcHdlr AsicDaemonServiceHandler) GetBulkIntf(currMarker, count asicdInt.Int) (*asicdInt.IntfGetInfo, error) {
	bulkObj := asicdInt.NewIntfGetInfo()
	return bulkObj, nil
}

//Logical Intf related services
func (svcHdlr AsicDaemonServiceHandler) CreateLogicalIntf(confObj *asicdServices.LogicalIntf) (rv bool, err error) {
	return rv, err
}
func (svcHdlr AsicDaemonServiceHandler) UpdateLogicalIntf(oldConfIntfObj, newConfIntfObj *asicdServices.LogicalIntf, attrset []bool, op []*asicdServices.PatchOpInfo) (rv bool, err error) {
	return rv, err
}
func (svcHdlr AsicDaemonServiceHandler) DeleteLogicalIntf(confObj *asicdServices.LogicalIntf) (rv bool, err error) {
	return rv, err
}
func (svcHdlr AsicDaemonServiceHandler) GetBulkLogicalIntfState(currMarker, count asicdServices.Int) (*asicdServices.LogicalIntfStateGetInfo, error) {
	bulkObj := asicdServices.NewLogicalIntfStateGetInfo()
	return bulkObj, nil
}
func (svcHdlr AsicDaemonServiceHandler) GetLogicalIntfState(name string) (*asicdServices.LogicalIntfState, error) {
	return nil, nil
}

//IPv4 interface related services
func (svcHdlr AsicDaemonServiceHandler) CreateIPv4Intf(ipv4IntfObj *asicdServices.IPv4Intf) (rv bool, err error) {
	return rv, err
}
func (svcHdlr AsicDaemonServiceHandler) UpdateIPv4Intf(oldIPv4IntfObj, newIPv4IntfObj *asicdServices.IPv4Intf, attrset []bool, op []*asicdServices.PatchOpInfo) (rv bool, err error) {
	return rv, err
}
func (svcHdlr AsicDaemonServiceHandler) DeleteIPv4Intf(ipv4IntfObj *asicdServices.IPv4Intf) (rv bool, err error) {
	return rv, err
}
func (svcHdlr AsicDaemonServiceHandler) GetIPv4IntfState(intfRef string) (*asicdServices.IPv4IntfState, error) {
	return nil, nil
}
func (svcHdlr AsicDaemonServiceHandler) GetBulkIPv4IntfState(currMarker, count asicdServices.Int) (*asicdServices.IPv4IntfStateGetInfo, error) {
	bulkObj := asicdServices.NewIPv4IntfStateGetInfo()
	return bulkObj, nil
}

//IPv4 Neighbor related services
func (svcHdlr AsicDaemonServiceHandler) CreateIPv4Neighbor(ipAddr string, macAddr string, vlanId, ifIndex int32) (rval int32, err error) {
	return rval, err
}
func (svcHdlr AsicDaemonServiceHandler) UpdateIPv4Neighbor(ipAddr string, macAddr string, vlanId, ifIndex int32) (rval int32, err error) {
	return rval, err
}
func (svcHdlr AsicDaemonServiceHandler) DeleteIPv4Neighbor(ipAddr string, macAddr string, vlanId, ifIndex int32) (rval int32, err error) {
	return rval, err
}
func (svcHdlr AsicDaemonServiceHandler) GetBulkArpEntryHwState(currMarker, count asicdServices.Int) (*asicdServices.ArpEntryHwStateGetInfo, error) {
	bulkObj := asicdServices.NewArpEntryHwStateGetInfo()
	return bulkObj, nil
}
func (svcHdlr AsicDaemonServiceHandler) GetArpEntryHwState(ipAddr string) (*asicdServices.ArpEntryHwState, error) {
	return nil, nil
}

//IPv4 Route related services
func (svcHdlr AsicDaemonServiceHandler) OnewayCreateIPv4Route(ipv4RouteList []*asicdInt.IPv4Route) error {
	nexthopInfo := OnewayHandleIPv4Route(ipv4RouteList, RADD)
	err := OneWayHandleIPv4Kafka(Gbroker, Gtopic, *nexthopInfo)
	if err != nil {
		return err
	}
	return nil
}

func (svcHdlr AsicDaemonServiceHandler) OnewayDeleteIPv4Route(ipv4RouteList []*asicdInt.IPv4Route) error {
	nexthopInfo := OnewayHandleIPv4Route(ipv4RouteList, RDEL)
	err := OneWayHandleIPv4Kafka(Gbroker, Gtopic, *nexthopInfo)
	if err != nil {
		return err
	}
	return nil
}

//IPv6 Route related services
func (svcHdlr AsicDaemonServiceHandler) OnewayCreateIPv6Route(ipv6RouteList []*asicdInt.IPv6Route) error {
	nexthopInfo := OnewayHandleIPv6Route(ipv6RouteList, RADD)
	err := OneWayHandleIPv6Kafka(Gbroker, Gtopic, *nexthopInfo)
	if err != nil {
		return err
	}
	return nil
}

func (svcHdlr AsicDaemonServiceHandler) OnewayDeleteIPv6Route(ipv6RouteList []*asicdInt.IPv6Route) error {
	nexthopInfo := OnewayHandleIPv6Route(ipv6RouteList, RDEL)
	err := OneWayHandleIPv6Kafka(Gbroker, Gtopic, *nexthopInfo)
	if err != nil {
		return err
	}
	return nil
}

func (svcHdlr AsicDaemonServiceHandler) GetBulkIPv4RouteHwState(currMarker, count asicdServices.Int) (*asicdServices.IPv4RouteHwStateGetInfo, error) {
	bulkObj := asicdServices.NewIPv4RouteHwStateGetInfo()
	return bulkObj, nil
}
func (svcHdlr AsicDaemonServiceHandler) GetIPv4RouteHwState(destinationNw string) (*asicdServices.IPv4RouteHwState, error) {
	return nil, nil
}

func (svcHdlr AsicDaemonServiceHandler) CreateSubIPv4Intf(subipv4IntfObj *asicdServices.SubIPv4Intf) (rv bool, err error) {
	return rv, err
}

func (svcHdlr AsicDaemonServiceHandler) UpdateSubIPv4Intf(oldSubIPv4IntfObj,
	newSubIPv4IntfObj *asicdServices.SubIPv4Intf, attrset []bool, op []*asicdServices.PatchOpInfo) (rv bool, err error) {
	return rv, err
}

func (svcHdlr AsicDaemonServiceHandler) DeleteSubIPv4Intf(subipv4IntfObj *asicdServices.SubIPv4Intf) (rv bool, err error) {
	return rv, err
}

// IPv6 Interface Create/Update/Delete API's
func (svcHdlr AsicDaemonServiceHandler) CreateIPv6Intf(ipv6IntfObj *asicdServices.IPv6Intf) (rv bool, err error) {
	return rv, err
}

func (svcHdlr AsicDaemonServiceHandler) UpdateIPv6Intf(oldIPv6IntfObj, newIPv6IntfObj *asicdServices.IPv6Intf, attrset []bool, op []*asicdServices.PatchOpInfo) (bool, error) {
	return true, nil
}

func (svcHdlr AsicDaemonServiceHandler) DeleteIPv6Intf(ipv6IntfObj *asicdServices.IPv6Intf) (rv bool, err error) {
	return rv, err
}

func (svcHdlr AsicDaemonServiceHandler) GetIPv6IntfState(intfRef string) (ipv6IntfState *asicdServices.IPv6IntfState, err error) {
	return ipv6IntfState, err
}

func (svcHdlr AsicDaemonServiceHandler) GetBulkIPv6IntfState(currMarker, count asicdServices.Int) (*asicdServices.IPv6IntfStateGetInfo, error) {
	bulkObj := asicdServices.NewIPv6IntfStateGetInfo()
	return bulkObj, nil
}

//IPv6 Neighbor related services
func (svcHdlr AsicDaemonServiceHandler) CreateIPv6Neighbor(ipAddr string, macAddr string, vlanId, ifIndex int32) (rval int32, err error) {
	return rval, err
}
func (svcHdlr AsicDaemonServiceHandler) UpdateIPv6Neighbor(ipAddr string, macAddr string, vlanId, ifIndex int32) (rval int32, err error) {
	return rval, err
}
func (svcHdlr AsicDaemonServiceHandler) DeleteIPv6Neighbor(ipAddr string, macAddr string, vlanId, ifIndex int32) (rval int32, err error) {
	return rval, err
}

// Sub IPv6 Interface Create/Delete/Update API's
func (svcHdlr AsicDaemonServiceHandler) CreateSubIPv6Intf(subipv6IntfObj *asicdServices.SubIPv6Intf) (rv bool, err error) {
	return rv, err
}

func (svcHdlr AsicDaemonServiceHandler) UpdateSubIPv6Intf(oldSubIPv6IntfObj,
	newSubIPv6IntfObj *asicdServices.SubIPv6Intf, attrset []bool, op []*asicdServices.PatchOpInfo) (rv bool, err error) {
	return rv, err
}

func (svcHdlr AsicDaemonServiceHandler) DeleteSubIPv6Intf(subipv6IntfObj *asicdServices.SubIPv6Intf) (rv bool, err error) {
	return rv, err
}

func (svcHdlr AsicDaemonServiceHandler) GetBulkNdpEntryHwState(currMarker, count asicdServices.Int) (*asicdServices.NdpEntryHwStateGetInfo, error) {
	return nil, nil
}

func (svcHdlr AsicDaemonServiceHandler) GetNdpEntryHwState(ipAddr string) (*asicdServices.NdpEntryHwState, error) {
	return nil, nil
}

func (svcHdlr AsicDaemonServiceHandler) GetBulkLinkScopeIpState(currMarker, count asicdServices.Int) (*asicdServices.LinkScopeIpStateGetInfo, error) {
	return nil, nil
}

func (svcHdlr AsicDaemonServiceHandler) GetLinkScopeIpState(lsIpAddr string) (*asicdServices.LinkScopeIpState, error) {
	return nil, nil
}

func (svcHdlr AsicDaemonServiceHandler) GetIPv6RouteHwState(destinationNw string) (*asicdServices.IPv6RouteHwState, error) {
	return nil, nil
}

func (svcHdlr AsicDaemonServiceHandler) GetBulkIPv6RouteHwState(currMarker, count asicdServices.Int) (*asicdServices.IPv6RouteHwStateGetInfo, error) {
	return nil, nil
}
