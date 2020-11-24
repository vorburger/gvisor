// Copyright 2020 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package header

import (
	"encoding/binary"
	"fmt"
	"time"

	"gvisor.dev/gvisor/pkg/tcpip"
)

// IGMP represents an IGMP header stored in a byte array.
type IGMP []byte

// IGMP implements `Transport`.
var _ Transport = (*IGMP)(nil)

const (
	// IGMPMinimumSize is the minimum size of a valid IGMP packet in bytes,
	// as per RFC 2236, Section 2, Page 2.
	IGMPMinimumSize = 8

	// IGMPQueryMinimumSize is the minimum size of a valid Membership Query
	// Message in bytes, as per RFC 2236, Section 2, Page 2.
	IGMPQueryMinimumSize = 8

	// IGMPReportMinimumSize is the minimum size of a valid Report Message in
	// bytes, as per RFC 2236, Section 2, Page 2.
	IGMPReportMinimumSize = 8

	// IGMPLeaveMessageMinimumSize is the minimum size of a valid Leave Message
	// in bytes, as per RFC 2236, Section 2, Page 2.
	IGMPLeaveMessageMinimumSize = 8

	// IGMPTTL is the TTL for all IGMP messages, as per RFC 2236, Section 3, Page
	// 3.
	IGMPTTL = 1

	// igmpTypeOffset defines the offset of the type field in an IGMP message.
	igmpTypeOffset = 0

	// igmpMaxRespTimeOffset defines the offset of the MaxRespTime field in an
	// IGMP message.
	igmpMaxRespTimeOffset = 1

	// igmpChecksumOffset defines the offset of the checksum field in an IGMP
	// message.
	igmpChecksumOffset = 2

	// igmpGroupAddressOffset defines the offset of the Group Address field in an
	// IGMP message.
	igmpGroupAddressOffset = 4

	// IGMPProtocolNumber is IGMP's transport protocol number.
	IGMPProtocolNumber tcpip.TransportProtocolNumber = 2
)

// IGMPType is the IGMP type field as per RFC 2236.
type IGMPType byte

// Values for the IGMP Type described in RFC 2236 Section 2.1, Page 2.
// Descriptions below come from there.
const (
	// IGMPMembershipQuery indicates that the message type is Membership Query.
	// "There are two sub-types of Membership Query messages:
	// - General Query, used to learn which groups have members on an
	//   attached network.
	// - Group-Specific Query, used to learn if a particular group
	//   has any members on an attached network.
	// These two messages are differentiated by the Group Address, as
	// described in section 1.4 ."
	IGMPMembershipQuery IGMPType = 0x11
	// IGMPv1MembershipReport indicates that the message is a Membership Report
	// generated by a host using the IGMPv1 protocol: "an additional type of
	// message, for backwards-compatibility with IGMPv1"
	IGMPv1MembershipReport IGMPType = 0x12
	// IGMPv2MembershipReport indicates that the Message type is a Membership
	// Report generated by a host using the IGMPv2 protocol.
	IGMPv2MembershipReport IGMPType = 0x16
	// IGMPLeaveGroup indicates that the message type is a Leave Group
	// notification message.
	IGMPLeaveGroup IGMPType = 0x17
)

// Type is the IGMP type field.
func (b IGMP) Type() IGMPType { return IGMPType(b[igmpTypeOffset]) }

// SetType sets the IGMP type field.
func (b IGMP) SetType(t IGMPType) { b[igmpTypeOffset] = byte(t) }

// MaxRespTime gets the MaxRespTimeField. This is meaningful only in Membership
// Query messages, in other cases it is set to 0 by the sender and ignored by
// the receiver.
func (b IGMP) MaxRespTime() time.Duration {
	// As per RFC 2236 section 2.2,
	//
	//  The Max Response Time field is meaningful only in Membership Query
	//  messages, and specifies the maximum allowed time before sending a
	//  responding report in units of 1/10 second.  In all other messages, it
	//  is set to zero by the sender and ignored by receivers.
	return DecisecondToDuration(b[igmpMaxRespTimeOffset])
}

// SetMaxRespTime sets the MaxRespTimeField.
func (b IGMP) SetMaxRespTime(m byte) { b[igmpMaxRespTimeOffset] = m }

// Checksum is the IGMP checksum field.
func (b IGMP) Checksum() uint16 {
	return binary.BigEndian.Uint16(b[igmpChecksumOffset:])
}

// SetChecksum sets the IGMP checksum field.
func (b IGMP) SetChecksum(checksum uint16) {
	binary.BigEndian.PutUint16(b[igmpChecksumOffset:], checksum)
}

// GroupAddress gets the Group Address field.
func (b IGMP) GroupAddress() tcpip.Address {
	return tcpip.Address(b[igmpGroupAddressOffset:][:IPv4AddressSize])
}

// SetGroupAddress sets the Group Address field.
func (b IGMP) SetGroupAddress(address tcpip.Address) {
	if n := copy(b[igmpGroupAddressOffset:], address); n != IPv4AddressSize {
		panic(fmt.Sprintf("copied %d bytes, expected %d", n, IPv4AddressSize))
	}
}

// SourcePort implements Transport.SourcePort.
func (IGMP) SourcePort() uint16 {
	return 0
}

// DestinationPort implements Transport.DestinationPort.
func (IGMP) DestinationPort() uint16 {
	return 0
}

// SetSourcePort implements Transport.SetSourcePort.
func (IGMP) SetSourcePort(uint16) {
}

// SetDestinationPort implements Transport.SetDestinationPort.
func (IGMP) SetDestinationPort(uint16) {
}

// Payload implements Transport.Payload.
func (IGMP) Payload() []byte {
	return nil
}

// IGMPCalculateChecksum calculates the IGMP checksum over the provided IGMP
// header.
func IGMPCalculateChecksum(h IGMP) uint16 {
	// The header contains a checksum itself, set it aside to avoid checksumming
	// the checksum and replace it afterwards.
	existingXsum := h.Checksum()
	h.SetChecksum(0)
	xsum := ^Checksum(h, 0)
	h.SetChecksum(existingXsum)
	return xsum
}

// DecisecondToDuration converts a value representing deci-seconds to a
// time.Duration.
func DecisecondToDuration(ds uint8) time.Duration {
	return time.Duration(ds) * time.Second / 10
}
