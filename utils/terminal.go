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

package utils

import "github.com/fatih/color"

// Define global color variables.
var (
	Yellow       = color.New(color.FgHiYellow, color.Bold).SprintFunc()
	YellowItalic = color.New(color.FgHiYellow, color.Bold, color.Italic).SprintFunc()
	Green        = color.New(color.FgHiGreen, color.Bold).SprintFunc()
	Blue         = color.New(color.FgHiBlue, color.Bold).SprintFunc()
	Cyan         = color.New(color.FgCyan, color.Bold, color.Underline).SprintFunc()
	Red          = color.New(color.FgHiRed, color.Bold).SprintFunc()
	White        = color.New(color.FgWhite).SprintFunc()
	WhiteBold    = color.New(color.FgWhite, color.Bold).SprintFunc()
	forceDetail  = "yaml"
)

func Print() string {
	return Blue(`
╭╮╱╭╮╱╭╮╱╱╱╱╱╱╱╱╱╭╮╱╱╱╭╮╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╭╮
┃┃╭╯╰┳╯╰╮╱╱╱╱╱╱╱╭╯╰╮╱╭╯╰╮╱╱╱╱╱╱╱╱╱╱╱╱╱╱╭╯╰╮
┃╰┻╮╭┻╮╭╋━━╮╱╱╭━┻╮╭╋━┻╮╭╯╱╭━━┳╮╭┳━━┳━━┳┻╮╭╋━━┳━╮
┃╭╮┃┃╱┃┃┃╭╮┣━━┫━━┫┃┃╭╮┃┣━━┫┃━╋╋╋┫╭╮┃╭╮┃╭┫┃┃┃━┫╭╯
┃┃┃┃╰╮┃╰┫╰╯┣━━╋━━┃╰┫╭╮┃╰┳━┫┃━╋╋╋┫╰╯┃╰╯┃┃┃╰┫┃━┫┃
╰╯╰┻━╯╰━┫╭━╯╱╱╰━━┻━┻╯╰┻━╯╱╰━━┻╯╰┫╭━┻━━┻╯╰━┻━━┻╯
╱╱╱╱╱╱╱╱┃┃╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱┃┃
╱╱╱╱╱╱╱╱╰╯╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╰╯
`)
}
