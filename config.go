// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package logp

import "time"

// Config contains the configuration options for the logger. To create a Config
// from a common.Config use logp/config.Build.
type Config struct {
	JSON      bool     `mapstructure:"json"`      // Write logs as JSON.
	Level     Level    `mapstructure:"level"`     // Logging level (error, warning, info, debug).
	Selectors []string `mapstructure:"selectors"` // Selectors for debug level logging.

	toObserver  bool
	toIODiscard bool
	ToStderr    bool `mapstructure:"to_stderr"`
	ToSyslog    bool `mapstructure:"to_syslog"`
	ToFiles     bool `mapstructure:"to_files"`
	ToEventLog  bool `mapstructure:"to_eventlog"`

	Files FileConfig `mapstructure:"files"`

	addCaller   bool // Adds package and line number info to messages.
	development bool // Controls how DPanic behaves.
}

// FileConfig contains the configuration options for the file output.
type FileConfig struct {
	Path        string        `mapstructure:"path"`
	Name        string        `mapstructure:"name"`
	MaxSize     uint          `mapstructure:"rotateeverybytes" validate:"min=1"`
	MaxBackups  uint          `mapstructure:"keepfiles" validate:"max=1024"`
	Permissions uint32        `mapstructure:"permissions"`
	Interval    time.Duration `mapstructure:"interval"`
}

var defaultConfig = Config{
	Level:    InfoLevel,
	ToFiles:  false,
	ToStderr: true,
	Files: FileConfig{
		MaxSize:     10 * 1024 * 1024,
		MaxBackups:  7,
		Permissions: 0600,
		Interval:    0,
	},
	addCaller: true,
}

// DefaultConfig returns the default config options.
func DefaultConfig() Config {
	return defaultConfig
}