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

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type Tvhapi struct {
	http   *http.Client
	server string
}

func Init(server string) Tvhapi {
	t := Tvhapi{&http.Client{Timeout: time.Second * 10}, server}
	return t
}

func (t Tvhapi) GetChannels() (ChannelList, error) {
	path := "/api/channel/list?start=0&limit=999999999&sort=number&dir=ASC&all=1"
	var cl ChannelList
	error := t.getJson(path, &cl)
	if error != nil {
		return cl, error
	}
	return cl, nil
}

func (t Tvhapi) GetEpg() (EpgGrid, error) {
	path := "/api/epg/events/grid?dir=ASC&sort=channelNumber&start=0&limit=20000&mode=now"

	var epg EpgGrid
	error := t.getJson(path, &epg)
	if error != nil {
		return epg, error
	}
	return epg, nil
}

func (t Tvhapi) GetStream(channelId Id) string {
	return t.server + "/stream/channel/" + string(channelId)
}

func (t Tvhapi) GetId(channelName string) (Id, error) {
	cl, error := t.GetChannels()
	if error != nil {
		return Id(""), error
	}

	for _, channel := range cl.Entries {
		if channel.Val == channelName {
			return channel.Key, nil
		}
	}

	return Id(""), errors.New("No channel with that name found.")
}

// helper functions
func (t Tvhapi) getJson(path string, target interface{}) error {
	response, error := t.http.Get(t.server + path)
	if error != nil {
		return error
	}
	defer response.Body.Close()

	error = json.NewDecoder(response.Body).Decode(target)
	if error != nil {
		return error
	}
	return nil
}
