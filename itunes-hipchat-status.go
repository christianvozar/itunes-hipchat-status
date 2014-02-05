// Copyright 2014, Rogue Ethic, Inc.
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
//
// Contributor(s):
//   Christian Vozar (christian@rogueethic.com)

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const (
	HipChatApiUrl     = "https://api.hipchat.com"
	HipChatApiVersion = "v2"
)

type HipChatUser struct {
	Name         string       `json:"name,omitempty"`
	Title        string       `json:"title,omitempty"`
	Presence     presenceInfo `json:"presence,omitempty"`
	MentionName  string       `json:"mention_name,omitempty"`
	Timezone     string       `json:"timezone,omitempty"`
	Email        string       `json:"email,omitempty"`
	IsGroupAdmin bool         `json:"is_group_admin"`
}

type presenceInfo struct {
	Status string `json:"status, show"`
	Show   string `json:"show, show"`
}

func main() {
	flagVersion := flag.Bool("version", false, "Display application version.")
	flagUser := flag.String("user", "", "Atlassian HipChat ID or Email of user to update.")
	flagAuthToken := flag.String("token", "", "Atlassian HipChat API v2 authentication token.")
	flag.Parse()

	// Handle no command-line paramters
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Output version
	if *flagVersion {
		fmt.Println("iTunes to Atlassian HipChat Status", Version)
		os.Exit(0)
	}

	userInformation := viewHipChatUser(*flagUser, *flagAuthToken)
	userInformation.Presence.Status = getItunesInformation()
	if userInformation.Presence.Show == "" {
		userInformation.Presence.Show = "chat"
	}
	updateHipChatUser(userInformation, *flagUser, *flagAuthToken)
}

func updateHipChatUser(u HipChatUser, e, a string) {
	messageUri := fmt.Sprintf("%s/%s/user/%s?auth_token=%s", HipChatApiUrl, HipChatApiVersion, e, a)

	messagePayload, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	body := strings.NewReader(string(messagePayload))

	httpClient := &http.Client{}
	req, err := http.NewRequest("PUT", messageUri, body)
	req.Header.Add("content-type", "application/json")

	_, err = httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}

func viewHipChatUser(e, a string) HipChatUser {
	var hipChatData HipChatUser

	messageUri := fmt.Sprintf("%s/%s/user/%s?auth_token=%s", HipChatApiUrl, HipChatApiVersion, e, a)
	res, err := http.Get(messageUri)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &hipChatData)

	return hipChatData
}

func getItunesInformation() string {
	appleScriptRuntime := "osascript"
	arg0 := "-e"
	cmd := exec.Command(appleScriptRuntime, arg0, `tell application "iTunes"
if it is running then
set trackname to name of current track
set artistname to artist of current track
set albumname to album of current track

if artistname is null then
set artistshow to ""
else if artistname is "" then
set artistshow to ""
else
set artistshow to " | " & artistname & ""
end if

set output to trackname & artistshow
end if
end tell`)

	out, err := cmd.Output()
	if err != nil {
		return ""
	}

	iTunesInformation := strings.TrimSpace(string(out))

	// HipChat status cannot exceed 50 characters.
	if len(iTunesInformation) > 50 {
		return iTunesInformation[0:46] + "..."
	}

	return iTunesInformation
}
