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

/*
 * This header file documents the plugin interfaces that need to be implemented by the
 * asic specific driver to integrate with the northbound stack
 */

#ifndef CUSTOM_CHIP_H
#define CUSTOM_CHIP_H
#include "pluginCommon.h"

/* Plugin init method
 * bootMode   - Specifies desired driver boot mode BOOT_MODE_COLDBOOT or BOOT_MODE_WARMBOOT
 * ifMapCount - Specifies list of interfaces to map into linux netdev
 * ifMap      - Interface information to use when mapping ports to linux netdev
 * macAddr    - Global switch mac address
 */
int Init(int bootMode, int ifMapCount, ifMapInfo_t *ifMap, uint8_t *macAddr);

/* Plugin deinit method
 * cacheSwState - Specifies whether to sync SW state for subsequent Warmboot
 */
int Deinit(int cacheSwState);

/* Utility method to retrieve max ports in system */
int GetMaxSysPorts();

/* Method to update port attributes
 * flags - Mask of port attributes that need to be updated
 * info  - Data containing attribute update values
 */
int UpdatePortConfig(int flags, portConfig *info);

/* Method to create a vlan
 * vlanId - Vlan ID of vlan to be created
 */
int CreateVlan(int vlanId);

/* Method to delete a vlan
 * vlanId - Vlan ID of vlan to be created
 */
int DeleteVlan(int vlanId);

/* Method to add ports to a vlan
 * vlanId         - Vlan ID of vlan
 * portCount      - Count of tagged ports
 * untagPortCount - Count of untagged ports
 * portList       - List of tagged ports
 * untagPortList  - List of untagged ports
 */
int UpdateVlanAddPorts(int vlanId, int portCount, int untagPortCount, int *portList, int *untagPortList);

/* Method to delete ports from a vlan
 * vlanId    - Vlan ID of vlan
 * portCount - Count of ports to be deleted
 * portList  - List of ports to be deleted
 */
int UpdateVlanDeletePorts(int vlanId, int portCount, int *portList);

/* Method to create an ip interface
 * macAddr - Source mac address to use for ip interface
 * vlanId  - Vlan id to use for vlan router interface
 */
int CreateIPIntf(uint8_t *macAddr, int vlanId);

/* Method to create an ip interface
 * vlanId  - Vlan id of vlan router interface
 */
int DeleteIPIntf(int vlanId);

/* Method to create ip nexthop
 * nextHopFlags - Nexthop flags applicable to this nexthop
 * vlanId       - vlan id of router interface that nexthop is reachable on
 * routePhyIntf - port number that nexthop is reachable on
 * macAddr      - mac address of nexthop
 */
uint64_t CreateIPNextHop(uint32_t nextHopFlags, int vlanId, int routerPhyIntf, uint8_t *macAddr, int ip_type);

/* Method to delete ip nexthop
 * nextHopId - H/W id of nexthop to delete
 */
int DeleteIPNextHop(uint64_t nextHopId);

/* Method to update ip nexthop
 * nextHopFlags - Nexthop flags applicable to this nexthop
 * vlanId       - vlan id of router interface that nexthop is reachable on
 * routePhyIntf - port number that nexthop is reachable on
 * macAddr      - mac address of nexthop
 */
int UpdateIPNextHop(uint32_t nextHopFlags, uint64_t nextHopId, int vlanId, int routerPhyIntf, uint8_t *macAddr), int ip_type;

/* Method to create a next hop group
 * numOfNh - Count of nexthops in group
 * nhIdArr - list of nexthop id's
 */
uint64_t CreateIPNextHopGroup(int numOfNh, uint64_t *nhIdArr);

/* Method to delete next hop group
 * ecmpGrpId - H/W id of nexthop group to be deleted
 */
int DeleteIPNextHopGroup(uint64_t ecmpGrpId);

/* Method to update a nexthop group
 * numOfNh - Count of nexthops in group
 * nhIdArr - list of nexthop id's
 * ecmpGrpId - H/W id of nexthop group to be deleted
 */
int UpdateIPNextHopGroup(int numOfNh, uint64_t *nhIdArr, uint64_t ecmpGrpId);

/* Method to create a host/neighbor
 * ipAddr    - IP Address of host
 * nextHopId - H/W id of nexthop corresponding to host
 * ip_type   - IPv4/IPv6
 */
int CreateIPNeighbor(uint32_t *ipAddr, uint64_t nextHopId, int ip_type);

/* Method to create a host/neighbor
 * ipAddr    - IP Address of host
 * nextHopId - H/W id of nexthop corresponding to host
 * ip_type   - IPv4/IPv6
 */
int UpdateIPNeighbor(uint32_t *ipAddr, uint64_t nextHopId, int ip_type);

/* Method to create a host/neighbor
 * ipAddr    - IP Address of host
 * ip_type   - IPv4/IPv6
 */
int DeleteIPNeighbor(uint32_t *ipAddr, int ip_type);

/* Method to create ip route entry
 * ipPrefix   - Network prefix of route
 * ipMask     - Network mask of route
 * routeFlags - Route flags as applicable to this route
 * nextHopId  - H/W id of nexthop to use for this route
 */
int CreateIPRoute(uint8_t *ipPrefix, uint8_t *ipMask, uint32_t routeFlags, uint64_t nextHopId);

/* Method to create ip route entry
 * ipPrefix   - Network prefix of route
 * ipMask     - Network mask of route
 * routeFlags - Route flags as applicable to this route
 */
int DeleteIPRoute(uint8_t *ipPrefix, uint8_t *ipMask,uint32_t routeFlags);

/* Method to create a lag
 * hashType  - Hash algorithm to use for this lag
 * portCount - Count of port members in this lag
 * ports     - List of ports to be added as members of this lag
 */
int CreateLag(int hashType, int portCount, int *ports);

/* Method to delete a lag
 * lagID - H/W id of lag to be deleted
 */
int DeleteLag(int lagId);

/* Method to update a lag
 * hashType     - Hash algorithm to use for this lag
 * oldPortCount - Count of existing port members in this lag
 * oldPorts     - List of existing port members of this lag
 * portCount    - Count of new port members in this lag
 * ports        - List of new ports to be added as members of this lag
 */
int UpdateLag(int lagId, int hashType, int oldPortCount, int *oldPorts, int portCount, int *ports);

/* Method to create an stg
 * listlen  - Count of vlan members in this stg
 * vlanList - List of vlan id's that are members of this stg
 */
int CreateStg(int listLen, int *vlanList);

/* Method to delete an stg
 * stgId - Id of spanning tree group to be deleted
 */
int DeleteStg(int stgId);

/* Method to set stp state of a port
 * stgId - Id of spanning tree group
 * port  - Id of port memeber in the stg
 * stpState - Spanning tree state of member port to set
 */
int SetPortStpState(int stgId, int port, int stpState);

/* Method to get stp state of a port
 * stgId - Id of spanning tree group
 * port  - Id of port memeber in the stg
 */
int GetPortStpState(int stgId, int port);

/* Method to update an stg
 * stgId       - Id of stg to update
 * oldListLen  - Count of vlan members in this stg
 * oldVlanList - List of vlan id's that are members of this stg
 * newListLen  - Count of vlan members to add to this stg
 * newVlanList - List of vlan id's that are to be added as members of this stg
 */
int UpdateStgVlanList(int stgId, int oldListLen, int *oldVlanList, int newListLen, int *newVlanList); 

/* Method to flush FDB table per vlan
 * vlan - ID of vlan corresponding to which FDB table entries are to be flushed
 */
int FlushFdbByVlan(int vlan);
#endif //CUSTOM_CHIP_H
