package network

type Route struct {
	Adv               bool        `json:"adv,omitempty"`
	Advbgp            bool        `json:"advbgp,omitempty"`
	Advertise         string      `json:"advertise,omitempty"`
	Advisis           bool        `json:"advisis,omitempty"`
	Advospf           bool        `json:"advospf,omitempty"`
	Advrip            bool        `json:"advrip,omitempty"`
	Bgp               bool        `json:"bgp,omitempty"`
	Cost              int         `json:"cost,omitempty"`
	Cost1             int         `json:"cost1,omitempty"`
	Data              bool        `json:"data,omitempty"`
	Data0             bool        `json:"data0,omitempty"`
	Detail            bool        `json:"detail,omitempty"`
	Dhcp              bool        `json:"dhcp,omitempty"`
	Direct            bool        `json:"direct,omitempty"`
	Distance          int         `json:"distance,omitempty"`
	Dynamic           bool        `json:"dynamic,omitempty"`
	Failedprobes      int         `json:"failedprobes,omitempty"`
	Flags             bool        `json:"flags,omitempty"`
	Gateway           string      `json:"gateway,omitempty"`
	Gatewayname       string      `json:"gatewayname,omitempty"`
	Isis              bool        `json:"isis,omitempty"`
	Lbroute           bool        `json:"lbroute,omitempty"`
	Monitor           string      `json:"monitor,omitempty"`
	Monstatcode       int         `json:"monstatcode,omitempty"`
	Monstatparam1     int         `json:"monstatparam1,omitempty"`
	Monstatparam2     int         `json:"monstatparam2,omitempty"`
	Monstatparam3     int         `json:"monstatparam3,omitempty"`
	Msr               string      `json:"msr,omitempty"`
	Nat               bool        `json:"nat,omitempty"`
	Netmask           string      `json:"netmask,omitempty"`
	Network           string      `json:"network,omitempty"`
	Ospf              bool        `json:"ospf,omitempty"`
	Ownergroup        string      `json:"ownergroup,omitempty"`
	Permanent         bool        `json:"permanent,omitempty"`
	Protocol          interface{} `json:"protocol,omitempty"`
	Retain            int         `json:"retain,omitempty"`
	Rip               bool        `json:"rip,omitempty"`
	Routeowners       interface{} `json:"routeowners,omitempty"`
	Routetype         string      `json:"routetype,omitempty"`
	State             int         `json:"state,omitempty"`
	Static            bool        `json:"Static,omitempty"`
	Td                int         `json:"td,omitempty"`
	Totalfailedprobes int         `json:"totalfailedprobes,omitempty"`
	Totalprobes       int         `json:"totalprobes,omitempty"`
	Tunnel            bool        `json:"tunnel,omitempty"`
	Type              bool        `json:"type,omitempty"`
	Weight            int         `json:"weight,omitempty"`
}
