/*
 * Application Manager API
 *
 * Application Manager APIs to control Apache Flink jobs
 *
 * API version: 2.1.0
 * Contact: platform@ververica.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package appmanagerapi

type SystemInformationStatus struct {
	JvmVersion string `json:"jvmVersion,omitempty"`
	ResourceQuota *ResourceQuota `json:"resourceQuota,omitempty"`
	RevisionInformation *RevisionInformation `json:"revisionInformation,omitempty"`
}
