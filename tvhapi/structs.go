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

package tvhapi

type Id string

type Channel struct {
	Key Id
	Val string
}

type ChannelList struct {
	Entries []Channel
}

type Epg struct {
	EventId       int
	EpisodeId     int
	ChannelName   string
	ChannelUuid   Id
	ChannelNumber string
	ChannelIcon   string
	Start         int
	Stop          int
	Title         string
	Subtitle      string
	Summary       string
	Description   string
	Widescreen    int
	Hd            int
	NextEventId   int
	Genre         []int
}

type EpgGrid struct {
	TotalCount int
	Entries    []Epg
}
