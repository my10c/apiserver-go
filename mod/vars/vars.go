//
// BSD 3-Clause License
//
// Copyright (c) 2023, © Badassops LLC / Luc Suryo
// All rights reserved.
//

package vars

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"
)

var (
	MyVersion   = "0.0.1"
	now         = time.Now()
	MyProgname  = path.Base(os.Args[0])
	myAuthor    = "Luc Suryo"
	myCopyright = "Copyright 2019 - " + strconv.Itoa(now.Year()) + " ©Badassops LLC"
	myLicense   = "License 3-Clause BSD, https://opensource.org/licenses/BSD-3-Clause ♥"
	myEmail     = "<luc@badassops.com>"
	MyInfo      = fmt.Sprintf("%s (version %s)\n%s\n%s\nWritten by %s %s\n",
		MyProgname, MyVersion, myCopyright, myLicense, myAuthor, myEmail)
	MyDescription = "Simple API server"

	// default values
	LogsDir       = "/var/log/aprserver-go"
	LogFile       = fmt.Sprintf("%s.log", MyProgname)
	LogMaxSize    = 128 // megabytes
	LogMaxBackups = 14  // 14 files
	LogMaxAge     = 14  // 14 days
)
