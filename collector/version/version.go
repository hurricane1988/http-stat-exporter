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

package version

import (
	"fmt"
	"runtime"
)

var (
	GitVersion   = "v1.0.0"
	GitCommit    = "unknown"
	GitTreeState = "unknown"
	BuildDate    = "unknown"
	GitMajor     = "unknown"
	GitMinor     = "unknown"
)

type Info struct {
	Company      string `json:"Company"`
	Author       string `json:"Author"`
	GitVersion   string `json:"gitVersion"`
	GitMajor     string `json:"gitMajor"`
	GitMinor     string `json:"gitMinor"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

func Get() Info {
	// These variables typically come from -ldflags settings and
	// in their absence fallback to the default settings
	return Info{
		GitVersion:   GitVersion,
		GitMajor:     GitMajor,
		GitMinor:     GitMinor,
		GitCommit:    GitCommit,
		GitTreeState: GitTreeState,
		BuildDate:    BuildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

// Print the version information.
func Print() {
	v := Get()
	fmt.Printf(`
----------------------------------------------
#   GitCommit: %s
#   BuildDate: %s
#   GoVersion: %s
#   Compiler: %s
#   Platform: %s
----------------------------------------------
`, v.GitCommit, v.BuildDate, v.GoVersion, v.Compiler, v.Platform)
}
