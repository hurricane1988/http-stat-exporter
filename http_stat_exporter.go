/*
Copyright 2024 Hurricane Authors

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

package main

import (
	"fmt"
	"github.com/hurricane1988/http-stat-exporter/collector/version"
	"github.com/hurricane1988/http-stat-exporter/utils"
	ctrl "sigs.k8s.io/controller-runtime"
)

// TODO: https://github.com/davecheney/httpstat/blob/master/main.go

var (
	setupLog = ctrl.Log.WithName("http-stat-exporter")
)

func main() {
	// Registered the terminal information.
	fmt.Println(utils.Print())
	version.Print()
}
