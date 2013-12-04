// Copyright 2013 The Crashwatch Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"code.google.com/p/go.exp/fsnotify"
	"github.com/mirtchovski/gosxnotifier"
)

var dir = flag.String("dir", "/Library/Logs/DiagnosticReports/:~/Library/Logs/DiagnosticReports", "column-separated list of directories to monitor")

func main() {
	flag.Parse()

	home := os.Getenv("HOME")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range strings.Split(*dir, ":") {
		v := strings.Replace(v, "~", home, -1)
		fmt.Println("monitoring:", v)
		go func() {
			err = watcher.Watch(v)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	for {
		select {
		case ev := <-watcher.Event:
			if ev.IsCreate() && !strings.HasSuffix(ev.Name, ".plist") {
				note := gosxnotifier.NewNotification("program" + path.Ext(ev.Name))
				note.Title = path.Base(ev.Name)
				note.Subtitle = ev.Name
				note.Sound = gosxnotifier.Basso

				note.Link = "file://" + ev.Name
				err := note.Push()
				if err != nil {
					log.Println("notification push error:", err)
				}
			} else {
				log.Println("event:", ev)
			}
		case err := <-watcher.Error:
			log.Println("error:", err)
		}
	}
}
