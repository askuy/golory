// Copyright 2018 golory Authors @1pb.club. All Rights Reserved.
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

// golory log component
package log

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

// Common logger configurations
type CommonCfg struct {
	Debug bool
	Level string
	Path  string
}

// TODO: for dynamic config & enrich log api
type Logger struct {
	*zap.Logger
}

// Initiate logger from config
func Boot(cfg CommonCfg) *Logger {
	var js string
	if cfg.Debug {
		js = fmt.Sprintf(`{
      "level": "%s",
      "encoding": "json",
      "outputPaths": ["stdout"],
      "errorOutputPaths": ["stdout"]
      }`, cfg.Level)
	} else {
		js = fmt.Sprintf(`{
      "level": "%s",
      "encoding": "json",
      "outputPaths": ["%s"],
      "errorOutputPaths": ["%s"]
      }`, cfg.Level, cfg.Path, cfg.Path)
	}

	var zcfg zap.Config
	if err := json.Unmarshal([]byte(js), &zcfg); err != nil {
		panic(err)
	}
	zcfg.EncoderConfig = zap.NewProductionEncoderConfig()
	zcfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	var err error
	l, err := zcfg.Build()
	if err != nil {
		log.Fatal("init logger error: ", err)
	}
	return &Logger{l}

}
