// Copyright 2015 David du Colombier. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pcap

const (
	HeaderSize       = 24         // Header size.
	RecordHeaderSize = 16         // Record header size.
	Magic            = 0xa1b2c3d4 // Magic number.
	MaxSnapLen       = 262144     // Maximum snapshot length.
)

// Link-layer header types.
const (
	LINKTYPE_NULL                       = DLT_NULL
	LINKTYPE_ETHERNET                   = DLT_EN10MB // also for 100Mb and up
	LINKTYPE_EXP_ETHERNET               = DLT_EN3MB  // 3Mb experimental Ethernet
	LINKTYPE_AX25                       = DLT_AX25
	LINKTYPE_PRONET                     = DLT_PRONET
	LINKTYPE_CHAOS                      = DLT_CHAOS
	LINKTYPE_IEEE802_5                  = DLT_IEEE802 // DLT_IEEE802 is used for 802.5 Token Ring
	LINKTYPE_ARCNET_BSD                 = DLT_ARCNET  // BSD-style headers
	LINKTYPE_SLIP                       = DLT_SLIP
	LINKTYPE_PPP                        = DLT_PPP
	LINKTYPE_FDDI                       = DLT_FDDI
	LINKTYPE_PPP_HDLC                   = 50  // PPP in HDLC-like framing
	LINKTYPE_PPP_ETHER                  = 51  // NetBSD PPP-over-Ethernet
	LINKTYPE_SYMANTEC_FIREWALL          = 99  // Symantec Enterprise Firewall
	LINKTYPE_ATM_RFC1483                = 100 // LLC/SNAP-encapsulated ATM
	LINKTYPE_RAW                        = 101 // raw IP
	LINKTYPE_SLIP_BSDOS                 = 102 // BSD/OS SLIP BPF header
	LINKTYPE_PPP_BSDOS                  = 103 // BSD/OS PPP BPF header
	LINKTYPE_MATCHING_MIN               = 104 // lowest value in the "matching" range
	LINKTYPE_C_HDLC                     = 104 // Cisco HDLC
	LINKTYPE_IEEE802_11                 = 105 // IEEE 802.11 (wireless)
	LINKTYPE_ATM_CLIP                   = 106 // Linux Classical IP over ATM
	LINKTYPE_FRELAY                     = 107 // Frame Relay
	LINKTYPE_LOOP                       = 108 // OpenBSD loopback
	LINKTYPE_ENC                        = 109 // OpenBSD IPSEC enc
	LINKTYPE_LANE8023                   = 110 // ATM LANE + 802.3
	LINKTYPE_HIPPI                      = 111 // NetBSD HIPPI
	LINKTYPE_HDLC                       = 112 // NetBSD HDLC framing
	LINKTYPE_LINUX_SLL                  = 113 // Linux cooked socket capture
	LINKTYPE_LTALK                      = 114 // Apple LocalTalk hardware
	LINKTYPE_ECONET                     = 115 // Acorn Econet
	LINKTYPE_IPFILTER                   = 116
	LINKTYPE_PFLOG                      = 117 // OpenBSD DLT_PFLOG
	LINKTYPE_CISCO_IOS                  = 118 // For Cisco-internal use
	LINKTYPE_IEEE802_11_PRISM           = 119 // 802.11 plus Prism II monitor mode radio metadata header
	LINKTYPE_IEEE802_11_AIRONET         = 120 // 802.11 plus FreeBSD Aironet driver radio metadata header
	LINKTYPE_HHDLC                      = 121
	LINKTYPE_IP_OVER_FC                 = 122 // RFC 2625 IP-over-Fibre Channel
	LINKTYPE_SUNATM                     = 123 // Solaris+SunATM
	LINKTYPE_RIO                        = 124 // RapidIO
	LINKTYPE_PCI_EXP                    = 125 // PCI Express
	LINKTYPE_AURORA                     = 126 // Xilinx Aurora link layer
	LINKTYPE_IEEE802_11_RADIOTAP        = 127 // 802.11 plus radiotap radio metadata header
	LINKTYPE_TZSP                       = 128 // Tazmen Sniffer Protocol
	LINKTYPE_ARCNET_LINUX               = 129 // Linux-style headers
	LINKTYPE_JUNIPER_MLPPP              = 130
	LINKTYPE_JUNIPER_MLFR               = 131
	LINKTYPE_JUNIPER_ES                 = 132
	LINKTYPE_JUNIPER_GGSN               = 133
	LINKTYPE_JUNIPER_MFR                = 134
	LINKTYPE_JUNIPER_ATM2               = 135
	LINKTYPE_JUNIPER_SERVICES           = 136
	LINKTYPE_JUNIPER_ATM1               = 137
	LINKTYPE_APPLE_IP_OVER_IEEE1394     = 138 // Apple IP-over-IEEE 1394 cooked header
	LINKTYPE_MTP2_WITH_PHDR             = 139
	LINKTYPE_MTP2                       = 140
	LINKTYPE_MTP3                       = 141
	LINKTYPE_SCCP                       = 142
	LINKTYPE_DOCSIS                     = 143 // DOCSIS MAC frames
	LINKTYPE_LINUX_IRDA                 = 144 // Linux-IrDA
	LINKTYPE_IBM_SP                     = 145
	LINKTYPE_IBM_SN                     = 146
	LINKTYPE_USER0                      = 147
	LINKTYPE_USER1                      = 148
	LINKTYPE_USER2                      = 149
	LINKTYPE_USER3                      = 150
	LINKTYPE_USER4                      = 151
	LINKTYPE_USER5                      = 152
	LINKTYPE_USER6                      = 153
	LINKTYPE_USER7                      = 154
	LINKTYPE_USER8                      = 155
	LINKTYPE_USER9                      = 156
	LINKTYPE_USER10                     = 157
	LINKTYPE_USER11                     = 158
	LINKTYPE_USER12                     = 159
	LINKTYPE_USER13                     = 160
	LINKTYPE_USER14                     = 161
	LINKTYPE_USER15                     = 162
	LINKTYPE_IEEE802_11_AVS             = 163 // 802.11 plus AVS radio metadata header
	LINKTYPE_JUNIPER_MONITOR            = 164
	LINKTYPE_BACNET_MS_TP               = 165
	LINKTYPE_PPP_PPPD                   = 166
	LINKTYPE_JUNIPER_PPPOE              = 167
	LINKTYPE_JUNIPER_PPPOE_ATM          = 168
	LINKTYPE_GPRS_LLC                   = 169 // GPRS LLC
	LINKTYPE_GPF_T                      = 170 // GPF-T (ITU-T G.7041/Y.1303)
	LINKTYPE_GPF_F                      = 171 // GPF-T (ITU-T G.7041/Y.1303)
	LINKTYPE_GCOM_T1E1                  = 172
	LINKTYPE_GCOM_SERIAL                = 173
	LINKTYPE_JUNIPER_PIC_PEER           = 174
	LINKTYPE_ERF_ETH                    = 175 // Ethernet
	LINKTYPE_ERF_POS                    = 176 // Packet-over-SONET
	LINKTYPE_LINUX_LAPD                 = 177
	LINKTYPE_JUNIPER_ETHER              = 178
	LINKTYPE_JUNIPER_PPP                = 179
	LINKTYPE_JUNIPER_FRELAY             = 180
	LINKTYPE_JUNIPER_CHDLC              = 181
	LINKTYPE_MFR                        = 182
	LINKTYPE_JUNIPER_VP                 = 183
	LINKTYPE_A429                       = 184
	LINKTYPE_A653_ICM                   = 185
	LINKTYPE_USB                        = 186
	LINKTYPE_BLUETOOTH_HCI_H4           = 187
	LINKTYPE_IEEE802_16_MAC_CPS         = 188
	LINKTYPE_USB_LINUX                  = 189
	LINKTYPE_CAN20B                     = 190
	LINKTYPE_IEEE802_15_4_LINUX         = 191
	LINKTYPE_PPI                        = 192
	LINKTYPE_IEEE802_16_MAC_CPS_RADIO   = 193
	LINKTYPE_JUNIPER_ISM                = 194
	LINKTYPE_IEEE802_15_4               = 195
	LINKTYPE_SITA                       = 196
	LINKTYPE_ERF                        = 197
	LINKTYPE_RAIF1                      = 198
	LINKTYPE_IPMB                       = 199
	LINKTYPE_JUNIPER_ST                 = 200
	LINKTYPE_BLUETOOTH_HCI_H4_WITH_PHDR = 201
	LINKTYPE_AX25_KISS                  = 202
	LINKTYPE_LAPD                       = 203
	LINKTYPE_PPP_WITH_DIR               = 204 // PPP
	LINKTYPE_C_HDLC_WITH_DIR            = 205 // Cisco HDLC
	LINKTYPE_FRELAY_WITH_DIR            = 206 // Frame Relay
	LINKTYPE_LAPB_WITH_DIR              = 207 // LAPB
	LINKTYPE_IPMB_LINUX                 = 209
	LINKTYPE_FLEXRAY                    = 210
	LINKTYPE_MOST                       = 211
	LINKTYPE_LIN                        = 212
	LINKTYPE_X2E_SERIAL                 = 213
	LINKTYPE_X2E_XORAYA                 = 214
	LINKTYPE_IEEE802_15_4_NONASK_PHY    = 215
	LINKTYPE_LINUX_EVDEV                = 216
	LINKTYPE_GSMTAP_UM                  = 217
	LINKTYPE_GSMTAP_ABIS                = 218
	LINKTYPE_MPLS                       = 219
	LINKTYPE_USB_LINUX_MMAPPED          = 220
	LINKTYPE_DECT                       = 221
	LINKTYPE_AOS                        = 222
	LINKTYPE_WIHART                     = 223
	LINKTYPE_FC_2                       = 224
	LINKTYPE_FC_2_WITH_FRAME_DELIMS     = 225
	LINKTYPE_IPNET                      = 226
	LINKTYPE_CAN_SOCKETCAN              = 227
	LINKTYPE_IPV4                       = 228
	LINKTYPE_IPV6                       = 229
	LINKTYPE_IEEE802_15_4_NOFCS         = 230
	LINKTYPE_DBUS                       = 231
	LINKTYPE_JUNIPER_VS                 = 232
	LINKTYPE_JUNIPER_SRX_E2E            = 233
	LINKTYPE_JUNIPER_FIBRECHANNEL       = 234
	LINKTYPE_DVB_CI                     = 235
	LINKTYPE_MUX27010                   = 236
	LINKTYPE_STANAG_5066_D_PDU          = 237
	LINKTYPE_JUNIPER_ATM_CEMIC          = 238
	LINKTYPE_NFLOG                      = 239
	LINKTYPE_NETANALYZER                = 240
	LINKTYPE_NETANALYZER_TRANSPARENT    = 241
	LINKTYPE_IPOIB                      = 242
	LINKTYPE_MPEG_2_TS                  = 243
	LINKTYPE_NG40                       = 244
	LINKTYPE_NFC_LLCP                   = 245
	LINKTYPE_PFSYNC                     = 246
	LINKTYPE_INFINIBAND                 = 247
	LINKTYPE_SCTP                       = 248
	LINKTYPE_USBPCAP                    = 249
	LINKTYPE_BLUETOOTH_LE_LL            = 251
	LINKTYPE_WIRESHARK_UPPER_PDU        = 252
	LINKTYPE_NETLINK                    = 253
	LINKTYPE_BLUETOOTH_LINUX_MONITOR    = 254
	LINKTYPE_BLUETOOTH_BREDR_BB         = 255
	LINKTYPE_BLUETOOTH_LE_LL_WITH_PHDR  = 256
	LINKTYPE_PROFIBUS_DL                = 257
	LINKTYPE_PKTAP                      = 258
	LINKTYPE_EPON                       = 259
	LINKTYPE_IPMI_HPM_2                 = 260
	LINKTYPE_MATCHING_MAX               = 260 // highest value in the "matching" range
)

const (
	DLT_NULL                       = 0  // BSD loopback encapsulation
	DLT_EN10MB                     = 1  // Ethernet (10Mb)
	DLT_EN3MB                      = 2  // Experimental Ethernet (3Mb)
	DLT_AX25                       = 3  // Amateur Radio AX.25
	DLT_PRONET                     = 4  // Proteon ProNET Token Ring
	DLT_CHAOS                      = 5  // Chaos
	DLT_IEEE802                    = 6  // 802.5 Token Ring
	DLT_ARCNET                     = 7  // ARCNET, with BSD-style header
	DLT_SLIP                       = 8  // Serial Line IP
	DLT_PPP                        = 9  // Point-to-point Protocol
	DLT_FDDI                       = 10 // FDDI
	DLT_ATM_RFC1483                = 11 // LLC-encapsulated ATM
	DLT_RAW1                       = 14 // raw IP
	DLT_RAW2                       = 12 // raw IP
	DLT_SLIP_BSDOS1                = 13 // BSD/OS Serial Line IP
	DLT_PPP_BSDOS1                 = 14 // BSD/OS Point-to-point Protocol
	DLT_SLIP_BSDOS2                = 15 // BSD/OS Serial Line IP
	DLT_PPP_BSDOS2                 = 16 // BSD/OS Point-to-point Protocol
	DLT_PFSYNC                     = 18
	DLT_ATM_CLIP                   = 19 // Linux Classical-IP over ATM
	DLT_REDBACK_SMARTEDGE          = 32
	DLT_PPP_SERIAL                 = 50 // PPP over serial with HDLC encapsulation
	DLT_PPP_ETHER                  = 51 // PPP over Ethernet
	DLT_SYMANTEC_FIREWALL          = 99
	DLT_MATCHING_MIN               = 104
	DLT_C_HDLC                     = 104 // Cisco HDLC
	DLT_CHDLC                      = DLT_C_HDLC
	DLT_IEEE802_11                 = 105 // IEEE 802.11 wireless
	DLT_FRELAY                     = 107
	DLT_LOOP1                      = 12
	DLT_LOOP2                      = 108
	DLT_ENC1                       = 13
	DLT_ENC2                       = 109
	DLT_LINUX_SLL                  = 113
	DLT_LTALK                      = 114
	DLT_ECONET                     = 115
	DLT_IPFILTER                   = 116
	DLT_PFLOG                      = 117
	DLT_CISCO_IOS                  = 118
	DLT_PRISM_HEADER               = 119
	DLT_AIRONET_HEADER             = 120
	DLT_PFSYNC1                    = 121
	DLT_HHDLC                      = 121
	DLT_IP_OVER_FC                 = 122
	DLT_SUNATM                     = 123 // Solaris+SunATM
	DLT_RIO                        = 124 // RapidIO
	DLT_PCI_EXP                    = 125 // PCI Express
	DLT_AURORA                     = 126 // Xilinx Aurora link layer
	DLT_IEEE802_11_RADIO           = 127 // 802.11 plus radiotap radio header
	DLT_TZSP                       = 128 // Tazmen Sniffer Protocol
	DLT_ARCNET_LINUX               = 129 // ARCNET
	DLT_JUNIPER_MLPPP              = 130
	DLT_JUNIPER_MLFR               = 131
	DLT_JUNIPER_ES                 = 132
	DLT_JUNIPER_GGSN               = 133
	DLT_JUNIPER_MFR                = 134
	DLT_JUNIPER_ATM2               = 135
	DLT_JUNIPER_SERVICES           = 136
	DLT_JUNIPER_ATM1               = 137
	DLT_APPLE_IP_OVER_IEEE1394     = 138
	DLT_MTP2_WITH_PHDR             = 139 // pseudo-header with various info, followed by MTP2
	DLT_MTP2                       = 140 // MTP2, without pseudo-header
	DLT_MTP3                       = 141 // MTP3, without pseudo-header or MTP2
	DLT_SCCP                       = 142 // SCCP, without pseudo-header or MTP2 or MTP3
	DLT_DOCSIS                     = 143
	DLT_LINUX_IRDA                 = 144
	DLT_IBM_SP                     = 145
	DLT_IBM_SN                     = 146
	DLT_USER0                      = 147
	DLT_USER1                      = 148
	DLT_USER2                      = 149
	DLT_USER3                      = 150
	DLT_USER4                      = 151
	DLT_USER5                      = 152
	DLT_USER6                      = 153
	DLT_USER7                      = 154
	DLT_USER8                      = 155
	DLT_USER9                      = 156
	DLT_USER10                     = 157
	DLT_USER11                     = 158
	DLT_USER12                     = 159
	DLT_USER13                     = 160
	DLT_USER14                     = 161
	DLT_USER15                     = 162
	DLT_IEEE802_11_RADIO_AVS       = 163 // 802.11 plus AVS radio header
	DLT_JUNIPER_MONITOR            = 164
	DLT_BACNET_MS_TP               = 165
	DLT_PPP_PPPD                   = 166
	DLT_JUNIPER_PPPOE              = 167
	DLT_JUNIPER_PPPOE_ATM          = 168
	DLT_GPRS_LLC                   = 169 // GPRS LLC
	DLT_GPF_T                      = 170 // GPF-T (ITU-T G.7041/Y.1303)
	DLT_GPF_F                      = 171 // GPF-F (ITU-T G.7041/Y.1303)
	DLT_GCOM_T1E1                  = 172
	DLT_GCOM_SERIAL                = 173
	DLT_JUNIPER_PIC_PEER           = 174
	DLT_ERF_ETH                    = 175 // Ethernet
	DLT_ERF_POS                    = 176 // Packet-over-SONET
	DLT_LINUX_LAPD                 = 177
	DLT_JUNIPER_ETHER              = 178
	DLT_JUNIPER_PPP                = 179
	DLT_JUNIPER_FRELAY             = 180
	DLT_JUNIPER_CHDLC              = 181
	DLT_MFR                        = 182
	DLT_JUNIPER_VP                 = 183
	DLT_A429                       = 184
	DLT_A653_ICM                   = 185
	DLT_USB                        = 186
	DLT_BLUETOOTH_HCI_H4           = 187
	DLT_IEEE802_16_MAC_CPS         = 188
	DLT_USB_LINUX                  = 189
	DLT_CAN20B                     = 190
	DLT_IEEE802_15_4_LINUX         = 191
	DLT_PPI                        = 192
	DLT_IEEE802_16_MAC_CPS_RADIO   = 193
	DLT_JUNIPER_ISM                = 194
	DLT_IEEE802_15_4               = 195
	DLT_SITA                       = 196
	DLT_ERF                        = 197
	DLT_RAIF1                      = 198
	DLT_IPMB                       = 199
	DLT_JUNIPER_ST                 = 200
	DLT_BLUETOOTH_HCI_H4_WITH_PHDR = 201
	DLT_AX25_KISS                  = 202
	DLT_LAPD                       = 203
	DLT_PPP_WITH_DIR               = 204 // PPP - don't confuse with DLT_PPP_WITH_DIRECTION
	DLT_C_HDLC_WITH_DIR            = 205 // Cisco HDLC
	DLT_FRELAY_WITH_DIR            = 206 // Frame Relay
	DLT_LAPB_WITH_DIR              = 207 // LAPB
	DLT_IPMB_LINUX                 = 209
	DLT_FLEXRAY                    = 210
	DLT_MOST                       = 211
	DLT_LIN                        = 212
	DLT_X2E_SERIAL                 = 213
	DLT_X2E_XORAYA                 = 214
	DLT_IEEE802_15_4_NONASK_PHY    = 215
	DLT_LINUX_EVDEV                = 216
	DLT_GSMTAP_UM                  = 217
	DLT_GSMTAP_ABIS                = 218
	DLT_MPLS                       = 219
	DLT_USB_LINUX_MMAPPED          = 220
	DLT_DECT                       = 221
	DLT_AOS                        = 222
	DLT_WIHART                     = 223
	DLT_FC_2                       = 224
	DLT_FC_2_WITH_FRAME_DELIMS     = 225
	DLT_IPNET                      = 226
	DLT_CAN_SOCKETCAN              = 227
	DLT_IPV4                       = 228
	DLT_IPV6                       = 229
	DLT_IEEE802_15_4_NOFCS         = 230
	DLT_DBUS                       = 231
	DLT_JUNIPER_VS                 = 232
	DLT_JUNIPER_SRX_E2E            = 233
	DLT_JUNIPER_FIBRECHANNEL       = 234
	DLT_DVB_CI                     = 235
	DLT_MUX27010                   = 236
	DLT_STANAG_5066_D_PDU          = 237
	DLT_JUNIPER_ATM_CEMIC          = 238
	DLT_NFLOG                      = 239
	DLT_NETANALYZER                = 240
	DLT_NETANALYZER_TRANSPARENT    = 241
	DLT_IPOIB                      = 242
	DLT_MPEG_2_TS                  = 243
	DLT_NG40                       = 244
	DLT_NFC_LLCP                   = 245
	DLT_PFSYNC2                    = 246
	DLT_INFINIBAND                 = 247
	DLT_SCTP                       = 248
	DLT_USBPCAP                    = 249
	DLT_RTAC_SERIAL                = 250
	DLT_BLUETOOTH_LE_LL            = 251
	DLT_WIRESHARK_UPPER_PDU        = 252
	DLT_NETLINK                    = 253
	DLT_BLUETOOTH_LINUX_MONITOR    = 254
	DLT_BLUETOOTH_BREDR_BB         = 255
	DLT_BLUETOOTH_LE_LL_WITH_PHDR  = 256
	DLT_PROFIBUS_DL                = 257
	DLT_PKTAP1                     = DLT_USER2
	DLT_PKTAP2                     = 258
	DLT_EPON                       = 259
	DLT_IPMI_HPM_2                 = 260
	DLT_MATCHING_MAX               = 260 // highest value in the "matching" range
)