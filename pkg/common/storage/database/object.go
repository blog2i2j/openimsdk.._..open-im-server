// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package database

import (
	"context"
	"time"

	"github.com/openimsdk/open-im-server/v3/pkg/common/storage/model"
)

type ObjectInfo interface {
	SetObject(ctx context.Context, obj *model.Object) error
	Take(ctx context.Context, engine string, name string) (*model.Object, error)
	Delete(ctx context.Context, engine string, name []string) error
	FindExpirationObject(ctx context.Context, engine string, expiration time.Time, needDelType []string, count int64) ([]*model.Object, error)
	GetKeyCount(ctx context.Context, engine string, key string) (int64, error)

	GetEngineCount(ctx context.Context, engine string) (int64, error)
	GetEngineInfo(ctx context.Context, engine string, limit int, skip int) ([]*model.Object, error)
	UpdateEngine(ctx context.Context, oldEngine, oldName string, newEngine string) error
}
