// Copyright (c) 2018 ChrisOboe
//
// This file is part of tvhcc
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"os/exec"
	"tvhcc/tvhapi"
)

func listChannels(api tvhapi.Tvhapi) {
	channels, error := api.GetChannels()
	if error != nil {
		fmt.Println(error.Error())
	}
	for _, channel := range channels.Entries {
		fmt.Println(channel.Val)
	}
}

func drawEpg(api tvhapi.Tvhapi) {
	epg, error := api.GetEpg()
	if error != nil {
		fmt.Println(error.Error())
	}

	simpleEpg := make(map[string][]string)
	longestChannelName := 0
	for _, entry := range epg.Entries {
		simpleEpg[entry.ChannelName] = append(simpleEpg[entry.ChannelName], entry.Title)
		if len(entry.ChannelName) > longestChannelName {
			longestChannelName = len(entry.ChannelName)
		}
	}

	for channel, titles := range simpleEpg {
		fmt.Print(channel)
		for i := len(channel); i <= longestChannelName; i++ {
			fmt.Print(" ")
		}
		currentPosition := longestChannelName
		currentTitle := 0
		terminalWidth, _, _ := terminal.GetSize(0)
		for len(titles) >= currentTitle+1 && len(titles[currentTitle])+currentPosition+3 <= terminalWidth {
			fmt.Print(" | " + titles[currentTitle])
			currentPosition += len(titles[currentTitle]) + 3
			currentTitle++
		}
		fmt.Println()
	}
}

func play(api tvhapi.Tvhapi, channel string) {
	channelId, error := api.GetId(channel)
	if error != nil {
		fmt.Println("This channel doesn't exist.")
	}
	cmd := exec.Command("mpv", api.GetStream(channelId))
	cmd.Start()
}

func printUsage() {
	fmt.Println("Usage: tvhcc <server> [epg|channels|play <channelName>]")
}

func main() {
	if len(os.Args) <= 2 {
		printUsage()
		os.Exit(0)
	}

	server := os.Args[1]
	tvhapi := tvhapi.Init(server)

	switch os.Args[2] {
	case "epg":
		drawEpg(tvhapi)
	case "channels":
		listChannels(tvhapi)
	case "play":
		if len(os.Args) <= 3 {
			printUsage()
			os.Exit(0)
		}
		play(tvhapi, os.Args[3])
	default:
		printUsage()
	}
}
