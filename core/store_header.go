/*
 * Copyright (C) 2019 IBM, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy ofthe License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specificlanguage governing permissions and
 * limitations under the License.
 *
 */

package core

import (
	"time"

	"github.com/spf13/viper"
)

type storeHeaderNone struct {
}

// Transform transforms a flow before being stored
func (s *storeHeaderNone) AddStoreHeader(flows []interface{}, startTime time.Time, endTime time.Time) interface{} {
	return flows
}

// NewTransformNone create a new transform
func NewStoreHeaderNone(cfg *viper.Viper) (interface{}, error) {
	return &storeHeaderNone{}, nil
}
