// SPDX-License-Identifier: AGPL-3.0-only
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>

package cache

import (
	"context"
	"time"
)

func NewNoop() RedisCache {
	return noop{}
}

type noop struct{}

func (n noop) GetMany(ctx context.Context, keys []string) GetManyResult {
	return GetManyResult{}
}

func (n noop) SetMany(ctx context.Context, data map[string]any, ttl time.Duration) error {
	return nil
}

func (n noop) Get(context.Context, string, any) (bool, error) {
	return false, nil
}

func (n noop) Set(context.Context, string, any, time.Duration) error {
	return nil
}

func (n noop) Del(context.Context, ...string) error {
	return nil
}
