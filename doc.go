// Copyright 2013 The Crashwatch Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
The crashwatch command monitors osx directories for crash dumps
created by the ReportCrash program. When a new crash dump is
created a notification event is sent to the default OSX notifier.

Without a list of paths, crashwatch monitors the defauls
/Library/Logs/DiagnosticReports and ~/Library/Logs/DiagnosticReports.

Usage:
        crashwatch [-dir PATH1:PATH2]

The flags are:
        -dir
                a column-separated list of directories to monitor
*/
package main
