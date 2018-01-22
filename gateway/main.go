// Gateway for TCP and HTTP services.
//
// Copyright 2018 github.com/xiaoenai. All Rights Reserved.
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
//
package main

import (
	"github.com/xiaoenai/ants/gateway/client"
	"github.com/xiaoenai/ants/gateway/http"
	"github.com/xiaoenai/ants/gateway/tcp"
)

func main() {
	client.Init(cfg.InnerClient)

	if cfg.EnableOuterHttp {
		go http.Serve(cfg.OuterHttpServer)
	}

	if cfg.EnableOuterTcp {
		go tcp.Serve(cfg.OuterTcpServer)
	}

	select {}
}