// Copyright 2016 fatedier, fatedier@gmail.com
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

package assets

import (
	"embed"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/xxl6097/go-http/server/inter"
	"github.com/xxl6097/go-http/server/route"
	http2 "github.com/xxl6097/go-service-framework/pkg/http"
	"io/fs"
	"net/http"
	"os"
)

//go:embed * static/*
var content embed.FS

func init() {
	subFs, err := fs.Sub(content, ".")
	fmt.Sprintf("FS:%v err:%v\n", subFs, err)
}

func Load() http.FileSystem {
	fsys := http.FS(content)
	fmt.Sprintf("fs:%v\n", fsys)
	return fsys
}

type assets struct {
}

func (this *assets) Setup(_router *mux.Router) {
	fsys := Load()
	router := _router.NewRoute().Subrouter()
	//newStaticFiles(router)
	//route.RouterUtil.AddNoAuthPrefix("/")
	route.RouterUtil.AddNoAuthPrefix("static")
	route.RouterUtil.AddNoAuthPrefix("favicon.ico")
	router.Handle("/favicon.ico", http.FileServer(fsys)).Methods("GET")
	router.PathPrefix("/").Handler(http2.MakeHTTPGzipHandler(http.StripPrefix("/", http.FileServer(fsys)))).Methods("GET")
}

func newStaticFiles(router *mux.Router) {
	route.RouterUtil.AddNoAuthPrefix("files")
	staticPrefix := "/files/"
	baseDir, _ := os.Getwd()
	router.PathPrefix(staticPrefix).Handler(http.StripPrefix(staticPrefix, http.FileServer(http.Dir(baseDir))))
}

func NewRoute() inter.IRoute {
	opt := &assets{}
	return opt
}
