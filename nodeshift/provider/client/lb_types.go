package client

import (
	"encoding/json"
	"time"
)

type LoadBalancerConfig struct {
	Name            string           `json:"name"`
	Replicas        map[string]int   `json:"replicas"`
	CPUUUIDs        []string         `json:"cpuUuids"`
	ForwardingRules []ForwardingRule `json:"forwardingRules"`
	VPCUUID         string           `json:"vpcUuid"`
}

type ForwardingRule struct {
	In  RuleEndpoint `json:"in"`
	Out RuleEndpoint `json:"out"`
}

type RuleEndpoint struct {
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
}

type LoadBalancerConfigResponse struct {
	UUID   string `json:"uuid"`
	Status string `json:"status"`
	TaskID string `json:"taskId"`
}

type GetLBResponse struct {
	UUID            string              `json:"uuid"`
	Name            string              `json:"name"`
	TaskID          string              `json:"taskId"`
	Status          string              `json:"status"`
	ReplicasAmount  int                 `json:"replicasAmount"`
	CPUAmount       int                 `json:"cpuAmount"`
	PriceInUSD      string              `json:"priceInUsd"`
	CreatedAt       time.Time           `json:"createdAt"`
	LoadBalancerCfg LoadBalancerDetails `json:"loadBalancerConfig"`
	VPC             VPCInfo             `json:"vpc"`
	Deployments     []Deployment        `json:"deployments"`
	Tags            []Tag               `json:"tags"`
}

type LoadBalancerDetails struct {
	DNSRecord       string             `json:"dnsRecord"`
	Replicas        []ReplicaInfo      `json:"replicas"`
	ForwardingRules []ForwardingRuleV2 `json:"forwardingRules"`
	Backends        []Backend          `json:"backends"`
}

type ReplicaInfo struct {
	Name   string `json:"name"`
	IP     string `json:"ip"`
	Region string `json:"region"`
	Status int    `json:"status"`
}

type ForwardingRuleV2 struct {
	In                RuleEndpoint               `json:"in"`
	Out               RuleEndpoint               `json:"out"`
	CertificateParams map[string]json.RawMessage `json:"certificateParams"`
}

type Backend struct {
	Public  BackendInterface `json:"public"`
	Private PrivateInterface `json:"private"`
}

type BackendInterface struct {
	ID string `json:"id"`
	IP string `json:"ip"`
}

type PrivateInterface struct {
	ID                string `json:"id"`
	IP                string `json:"ip"`
	VPCName           string `json:"vpcName"`
	VPCBootstrapToken string `json:"vpcBootstrapToken"`
}

type VPCInfo struct {
	UUID               string        `json:"uuid"`
	Name               string        `json:"name"`
	Description        string        `json:"description"`
	AddressRangePrefix string        `json:"addressRangePrefix"`
	NetworkSize        int           `json:"networkSize"`
	Token              string        `json:"token"`
	Resources          []VPCResource `json:"resources"`
	Tags               []Tag         `json:"tags"`
	CreatedAt          time.Time     `json:"created_at"`
}

type VPCResource struct {
	UUID      string    `json:"uuid"`
	Hostname  string    `json:"hostname"`
	IP        string    `json:"ip"`
	Status    int       `json:"status"`
	CRU       int       `json:"cru"`
	MRU       int       `json:"mru"`
	SRU       int       `json:"sru"`
	TaskID    string    `json:"taskId"`
	CreatedAt time.Time `json:"created_at"`
}

type Deployment struct {
	UUID         string    `json:"uuid"`
	Status       int       `json:"status"`
	IP           string    `json:"ip"`
	CRU          int       `json:"cru"`
	MRU          int       `json:"mru"`
	SRU          int       `json:"sru"`
	HRU          int       `json:"hru"`
	Hostname     string    `json:"hostname"`
	IPv4         bool      `json:"ipv4"`
	SSHKey       string    `json:"sshKey"`
	ImageVersion string    `json:"imageVersion"`
	Price        string    `json:"price"`
	Region       string    `json:"region"`
	HDDType      int       `json:"hddType"`
	ChargeRate   float64   `json:"chargeRate"`
	Tags         []Tag     `json:"tags"`
	CreatedAt    time.Time `json:"createdAt"`
	AuthType     int       `json:"authType"`
	Password     string    `json:"password"`
}

type Tag struct {
	TagUUID string `json:"tagUuid"`
	Name    string `json:"name"`
}
