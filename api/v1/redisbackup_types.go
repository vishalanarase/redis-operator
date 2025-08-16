/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RedisBackupSpec defines the desired state of RedisBackup
type RedisBackupSpec struct {
	// Name of the Redis cluster to back up
	ClusterName string `json:"clusterName"`

	// Backup storage configuration
	Storage BackupStorage `json:"storage"`

	// Duration to retain backups
	RetentionPeriod string `json:"retentionPeriod,omitempty"`
}

type BackupStorage struct {
	// S3 compatible storage
	S3 *S3Storage `json:"s3,omitempty"`

	// GCS compatible storage
	//GCS *GCSStorage `json:"gcs,omitempty"`

	// Azure compatible storage
	//Azure *AzureStorage `json:"azure,omitempty"`
}

type S3Storage struct {
	// S3 bucket name
	Bucket string `json:"bucket"`

	// S3 endpoint
	Endpoint string `json:"endpoint"`

	// S3 secret reference
	SecretRef string `json:"secretRef"`
}

// RedisBackupStatus defines the observed state of RedisBackup
type RedisBackupStatus struct {
	// Backup state
	State bool `json:"state"`

	// Conditions represent the latest available observations of a backup's state
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// RedisBackup is the Schema for the redisbackups API
type RedisBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedisBackupSpec   `json:"spec,omitempty"`
	Status RedisBackupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedisBackupList contains a list of RedisBackup
type RedisBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RedisBackup{}, &RedisBackupList{})
}
