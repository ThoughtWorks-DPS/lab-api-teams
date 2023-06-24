package domain

// gw-id: subdomain of gateway, e.g., dev.api.twdps.io
// gw-cert-managed: is this a supported gateway, cert-manager to manage?
// gw-custom-cert: where the customer supplies the cert?
// gw-ingress: ingress-gateway to attach

type Gateway struct {
	GatewayID          string `json:"gatewayID"`
	GatewayCertManaged bool   `json:"gatewayCertManaged"`
	GatewayCustomCert  string `json:"gatewayCustomCert"`
	GatewayIngress     string `json:"gatewayIngress"`
}

type GatewayRepository interface {
	GetGateways() ([]Gateway, error)
	GetGatewayByID(id string) (Gateway, error)
}
