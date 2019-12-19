/*
 * AppManager API
 *
 * HTTP REST API to connect to the AppManager
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package appManagerApiClient

import core "k8s.io/api/core/v1"

type VolumeAndMount struct {
	Name        string            `json:"name,omitempty"`
	Volume      *core.Volume      `json:"volume,omitempty"`
	VolumeMount *core.VolumeMount `json:"volumeMount,omitempty"`
}