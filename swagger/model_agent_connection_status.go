/*
 * syntropy-controller
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AgentConnectionStatus string

// List of AgentConnectionStatus
const (
	PENDING_AgentConnectionStatus AgentConnectionStatus = "PENDING"
	WARNING_AgentConnectionStatus AgentConnectionStatus = "WARNING"
	ERROR__AgentConnectionStatus AgentConnectionStatus = "ERROR"
	CONNECTED_AgentConnectionStatus AgentConnectionStatus = "CONNECTED"
	OFFLINE_AgentConnectionStatus AgentConnectionStatus = "OFFLINE"
)