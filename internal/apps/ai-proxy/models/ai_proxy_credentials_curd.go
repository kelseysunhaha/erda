// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by erda-cli. DO NOT EDIT.
// Source: .erda/ai-proxy/migrations/ai-proxy/20230706-credentials.sql

package models

import (
	"gorm.io/gorm"
)

// FieldID returns the Field interface{} for the field ai_proxy_credentials.id
func (this *AIProxyCredentials) FieldID() Field { return field{name: "id"} }

// FieldCreatedAt returns the Field interface{} for the field ai_proxy_credentials.created_at
func (this *AIProxyCredentials) FieldCreatedAt() Field { return field{name: "created_at"} }

// FieldUpdatedAt returns the Field interface{} for the field ai_proxy_credentials.updated_at
func (this *AIProxyCredentials) FieldUpdatedAt() Field { return field{name: "updated_at"} }

// FieldDeletedAt returns the Field interface{} for the field ai_proxy_credentials.deleted_at
func (this *AIProxyCredentials) FieldDeletedAt() Field { return field{name: "deleted_at"} }

// FieldAccessKeyID returns the Field interface{} for the field ai_proxy_credentials.access_key_id
func (this *AIProxyCredentials) FieldAccessKeyID() Field { return field{name: "access_key_id"} }

// FieldSecretKeyID returns the Field interface{} for the field ai_proxy_credentials.secret_key_id
func (this *AIProxyCredentials) FieldSecretKeyID() Field { return field{name: "secret_key_id"} }

// FieldName returns the Field interface{} for the field ai_proxy_credentials.name
func (this *AIProxyCredentials) FieldName() Field { return field{name: "name"} }

// FieldPlatform returns the Field interface{} for the field ai_proxy_credentials.platform
func (this *AIProxyCredentials) FieldPlatform() Field { return field{name: "platform"} }

// FieldDescription returns the Field interface{} for the field ai_proxy_credentials.description
func (this *AIProxyCredentials) FieldDescription() Field { return field{name: "description"} }

// FieldEnabled returns the Field interface{} for the field ai_proxy_credentials.enabled
func (this *AIProxyCredentials) FieldEnabled() Field { return field{name: "enabled"} }

// FieldExpiredAt returns the Field interface{} for the field ai_proxy_credentials.expired_at
func (this *AIProxyCredentials) FieldExpiredAt() Field { return field{name: "expired_at"} }

// FieldProviderName returns the Field interface{} for the field ai_proxy_credentials.provider_name
func (this *AIProxyCredentials) FieldProviderName() Field { return field{name: "provider_name"} }

// FieldProviderInstanceID returns the Field interface{} for the field ai_proxy_credentials.provider_instance_id
func (this *AIProxyCredentials) FieldProviderInstanceID() Field {
	return field{name: "provider_instance_id"}
}

func (this *AIProxyCredentials) Creator(db *gorm.DB) Creator {
	return &creator{db: db, model: this}
}

func (this *AIProxyCredentials) Deleter(db *gorm.DB) Deleter {
	return &deleter{
		db:    db.Model(this),
		model: this,
		where: make([]Where, 0),
	}
}

func (this *AIProxyCredentials) Updater(db *gorm.DB) Updater {
	return &updater{
		db:      db.Model(this),
		model:   this,
		where:   nil,
		updates: make(map[string]any),
	}
}

func (this *AIProxyCredentials) Getter(db *gorm.DB) Getter {
	return &getter{
		db:    db.Model(this),
		model: this,
		where: make([]Where, 0),
	}
}

// FieldID returns the Field interface{} for the field ai_proxy_credentials.id
func (list AIProxyCredentialsList) FieldID() Field { return field{name: "id"} }

// FieldCreatedAt returns the Field interface{} for the field ai_proxy_credentials.created_at
func (list AIProxyCredentialsList) FieldCreatedAt() Field { return field{name: "created_at"} }

// FieldUpdatedAt returns the Field interface{} for the field ai_proxy_credentials.updated_at
func (list AIProxyCredentialsList) FieldUpdatedAt() Field { return field{name: "updated_at"} }

// FieldDeletedAt returns the Field interface{} for the field ai_proxy_credentials.deleted_at
func (list AIProxyCredentialsList) FieldDeletedAt() Field { return field{name: "deleted_at"} }

// FieldAccessKeyID returns the Field interface{} for the field ai_proxy_credentials.access_key_id
func (list AIProxyCredentialsList) FieldAccessKeyID() Field { return field{name: "access_key_id"} }

// FieldSecretKeyID returns the Field interface{} for the field ai_proxy_credentials.secret_key_id
func (list AIProxyCredentialsList) FieldSecretKeyID() Field { return field{name: "secret_key_id"} }

// FieldName returns the Field interface{} for the field ai_proxy_credentials.name
func (list AIProxyCredentialsList) FieldName() Field { return field{name: "name"} }

// FieldPlatform returns the Field interface{} for the field ai_proxy_credentials.platform
func (list AIProxyCredentialsList) FieldPlatform() Field { return field{name: "platform"} }

// FieldDescription returns the Field interface{} for the field ai_proxy_credentials.description
func (list AIProxyCredentialsList) FieldDescription() Field { return field{name: "description"} }

// FieldEnabled returns the Field interface{} for the field ai_proxy_credentials.enabled
func (list AIProxyCredentialsList) FieldEnabled() Field { return field{name: "enabled"} }

// FieldExpiredAt returns the Field interface{} for the field ai_proxy_credentials.expired_at
func (list AIProxyCredentialsList) FieldExpiredAt() Field { return field{name: "expired_at"} }

// FieldProviderName returns the Field interface{} for the field ai_proxy_credentials.provider_name
func (list AIProxyCredentialsList) FieldProviderName() Field { return field{name: "provider_name"} }

// FieldProviderInstanceID returns the Field interface{} for the field ai_proxy_credentials.provider_instance_id
func (list AIProxyCredentialsList) FieldProviderInstanceID() Field {
	return field{name: "provider_instance_id"}
}

func (list *AIProxyCredentialsList) Pager(db *gorm.DB) Pager {
	return &pager{
		db:    db,
		list:  list,
		where: nil,
	}
}