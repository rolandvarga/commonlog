package commonlog

import (
	"testing"
)

func TestParse(t *testing.T) {
	log := []string{
		"127.0.0.1 user-identifier frank [10/Oct/2000:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
	}

	entries := Parse(log)
	first := entries[0]

	wantIP := "127.0.0.1"
	if first.IP != wantIP {
		t.Errorf("Not identifying IP address.\nGOT '%s' || WANT '%s'", first.IP, wantIP)
	}

	wantDT := "[10/Oct/2000:13:55:36 -0700]"
	if first.DateTime != wantDT {
		t.Errorf("Not identifying DateTime.\nGOT '%s' || WANT '%s'", first.DateTime, wantDT)
	}

}

func TestDraw(t *testing.T) {
	log := []string{
		"127.0.0.1 - - [09/Feb/2019:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"103.75.54.153 - - [09/Feb/2019:13:55:35 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"103.75.54.153 - - [09/Feb/2019:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"103.75.54.153 - - [09/Feb/2019:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"9.234.152.164 - - [09/Feb/2019:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"9.234.152.164 - - [09/Feb/2019:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"67.233.108.118 - - [09/Feb/2019:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"185.219.65.88 - - [09/Feb/2019:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"89.76.143.107 - - [09/Feb/2019:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"89.76.143.107 - - [09/Feb/2019:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"224.250.71.200 - - [09/Feb/2019:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"254.196.136.219 - - [09/Feb/2019:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"70.225.104.238 - - [09/Feb/2019:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"187.55.126.186 - - [09/Feb/2019:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"79.24.73.164 - - [09/Feb/2019:13:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"103.75.54.153 - - [09/Feb/2019:14:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"103.75.54.153 - - [09/Feb/2019:14:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"103.75.54.153 - - [09/Feb/2019:14:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"9.234.152.164 - - [09/Feb/2019:14:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"9.234.152.164 - - [09/Feb/2019:14:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"67.233.108.118 - - [09/Feb/2019:14:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"185.219.65.88 - - [09/Feb/2019:14:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"89.76.143.107 - - [09/Feb/2019:14:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"89.76.143.107 - - [09/Feb/2019:14:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"224.250.71.200 - - [09/Feb/2019:14:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"254.196.146.219 - - [09/Feb/2019:14:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"70.225.104.238 - - [09/Feb/2019:14:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"187.55.126.186 - - [09/Feb/2019:14:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"79.24.73.164 - - [09/Feb/2019:14:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"103.75.54.153 - - [09/Feb/2019:15:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"103.75.54.153 - - [09/Feb/2019:15:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"103.75.54.153 - - [09/Feb/2019:15:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"9.234.152.164 - - [09/Feb/2019:15:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"9.234.152.164 - - [09/Feb/2019:15:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"67.233.108.118 - - [09/Feb/2019:15:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"185.219.65.88 - - [09/Feb/2019:15:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"89.76.153.107 - - [09/Feb/2019:15:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"89.76.153.107 - - [09/Feb/2019:15:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"224.250.71.200 - - [09/Feb/2019:15:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"254.196.156.219 - - [09/Feb/2019:15:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"70.225.104.238 - - [09/Feb/2019:15:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"187.55.126.186 - - [09/Feb/2019:15:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
		"79.24.73.164 - - [09/Feb/2019:15:55:36 -0700] \"GET /apache_pb.gif HTTP/1.0\" 200 2326",
	}

	entries := Parse(log)
	unique := UniqueViews(entries)
	Draw(unique)
}
