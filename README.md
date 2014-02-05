# Apple iTunes to Atlassian HipChat Status

A simple application that takes the current track and artist from iTunes and updates your Atlassian HipChat status to reflect.

## Requirements

- While the application is written in Google Go, it will spawn a process that utilizes AppleScript to check iTunes for the current track. So any OSX should do but this is tested against the latest release, 10.9.
- Compilation requires Google Go 1.2
- Atlassian HipChat account

## Installation

The application uses only the standard Go libraries so installation is as simple as compiling the binary on your platform; assuming your GOROOT and GOPATH are setup correctly.

```
$ git clone https://github.com/christianvozar/itunes-hipchat-status.git
$ go build
```

## Usage

From the commandline you can utilize the application by passing in your HipChat username and [authentication token](https://www.hipchat.com/docs/apiv2/auth).
```
./itunes-hipchat-status -user=yourname@domain.com -token=XXXXXXXXXXXXXXXXXXXXXXXXXX
```

A more appropriate usage is to schedule the updater to run on a cron schedule of every 15 or 30 seconds.


## License

Copyright 2014, Rogue Ethic, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
