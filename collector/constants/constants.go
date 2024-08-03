/*
Copyright 2024 Faw Authors

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

package constants

import "time"

// Default http setup value.

const (
	MaxRedirects = 10
)

// Default version value.

const (
	VersionDev = "v1.0.0" // for -v flag, updated during the release process with -ldflags=-X=collector.version.version=... =
)

// Default certificate suffix value.
const (
	SuffixPrivateKey  = "PRIVATE KEY"
	SuffixCertificate = "CERTIFICATE"
)

// Default url value.
const (
	HttpListenPort  = ":80"
	HttpDoubleSlash = "//"
)

// Default http request method.
const (
	PostHttpMethod    = "POST"    // 用于请求数据，从服务器获取资源
	GetHttpMethod     = "GET"     // 用于提交数据给服务器，通常用于提交表单或上传文件
	PutHttpMethod     = "PUT"     // 用于更新服务器上的资源，传输数据的整体
	DeleteHttpMethod  = "DELETE"  // 用于删除服务器上的资源
	HeadHttpMethod    = "HEAD"    // 与 GET 类似，但只请求资源的头部信息，不返回实际数据
	OptionsHttpMethod = "OPTIONS" // 用于获取服务器支持的请求方法
	PatchHttpMethod   = "PATCH"   // 用于局部更新服务器上的资源
)

// Default http scheme value.
const (
	EmptyHttpScheme    = ""
	SecurityHttpScheme = "https"
	DefaultHttpScheme  = "http"
)

// Http request value.
const (
	HttpTimeOut   = 30 * time.Second
	HttpKeepAlive = 30 * time.Second
)
