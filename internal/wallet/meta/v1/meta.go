// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"time"
)

type ObjectMetaAccessor interface {
	GetObjectMeta() Object
}

// Object lets you work with object metadata from any of the versioned or
// internal API objects. Attempting to set or retrieve a field on an object that does
// not support that field (Name, UID, Namespace on lists) will be a no-op and return
// a default value.
type Object interface {
	GetID() uint64
	SetID(id uint64)
	GetCreatedAt() time.Time
	SetCreatedAt(createdAt time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updatedAt time.Time)
}

// ListInterface lets you work with list metadata from any of the versioned or
// internal API objects. Attempting to set or retrieve a field on an object that does
// not support that field will be a no-op and return a default value.
type ListInterface interface {
	GetTotalCount() int64
	SetTotalCount(count int64)
	GetCurrentPage() int64
	SetCurrentPage(offset int64)
	GetTotalPage() int64
	SetTotalPage(limit int64)
}

var _ ListInterface = &ListMeta{}

func (meta *ListMeta) GetTotalCount() int64      { return meta.TotalCount }
func (meta *ListMeta) SetTotalCount(count int64) { meta.TotalCount = count }

func (meta *ListMeta) GetListMeta() ListInterface { return meta }

func (obj *ObjectMeta) GetObjectMeta() Object { return obj }

var _ Object = &ObjectMeta{}

func (meta *ObjectMeta) GetID() uint64                    { return meta.ID }
func (meta *ObjectMeta) SetID(id uint64)                  { meta.ID = id }
func (meta *ObjectMeta) GetCreatedAt() time.Time          { return meta.CreatedAt }
func (meta *ObjectMeta) SetCreatedAt(createdAt time.Time) { meta.CreatedAt = createdAt }
func (meta *ObjectMeta) GetUpdatedAt() time.Time          { return meta.UpdatedAt }
func (meta *ObjectMeta) SetUpdatedAt(updatedAt time.Time) { meta.UpdatedAt = updatedAt }
