package commonlog

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type LogEntry struct {
	IP       string
	DateTime string
}

func Parse(log []string) (o []LogEntry) {
	for _, l := range log {
		ipReg, _ := regexp.Compile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
		ip := ipReg.FindString(l)

		dtReg, _ := regexp.Compile(`\[.+\]`)
		dt := dtReg.FindString(l)

		entry := LogEntry{
			IP:       ip,
			DateTime: dt,
		}
		o = append(o, entry)
	}
	return o
}

func UniqueViews(entries []LogEntry) map[string]map[string]int {
	var format = "[02/Jan/2006:15:04:05 -0700]"
	o := make(map[string]map[string]int)

	for _, entry := range entries {
		t, err := time.Parse(format, entry.DateTime)
		if err != nil {
			fmt.Println(err.Error())
		}

		rounded := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())

		in, ok := o[rounded.String()]
		if !ok {
			in = make(map[string]int)
			o[rounded.String()] = in
		}
		in[entry.IP]++
	}
	return o
}

func Draw(entries map[string]map[string]int) {
	for k, v := range entries {
		fmt.Printf("\n%s", k)
		fmt.Printf("\n%s\t\t%s\t%s\n", "IP", "Count", "Histogram")
		for ik, iv := range v {
			hist := strings.Repeat("∎", iv)
			fmt.Printf("%s\t%d\t%s\n", ik, iv, hist)
		}
	}
}
