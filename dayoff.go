package dayoff

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

const (
	ExitCodeOK    int = iota // 0
	ExitCodeError            // 1
)

var days []string

func prepare() error {
	var array []string

	usr, err := user.Current()
	if err != nil {
		fmt.Fprintln(os.Stderr, "[Error] Could not get current user:", err)
		return err
	}

	path := filepath.Join(usr.HomeDir, "/.config/dayoff/days.json")
	str, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[Error] Could not read config.json:", err)
		return err
	}

	if err := json.Unmarshal(str, &array); err != nil {
		fmt.Fprintln(os.Stderr, "[Error] JSON Unmarshal:", err)
		return err
	}

	days = array
	return nil
}

func IsDayOff(t time.Time) bool {
	if len(days) == 0 {
		if err := prepare(); err != nil {
			return false
		}
	}

	var JST = time.FixedZone("JST", 3600*9)
	jst := t.In(JST)
	str := fmt.Sprintf("%04d-%02d-%02d", jst.Year(), jst.Month(), jst.Day())

	for _, s := range days {
		if s == str {
			return true
		}
	}
	return false
}
