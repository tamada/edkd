package main

import "testing"

func TestParseArguments(t *testing.T) {
	testdata := []struct {
		giveArgs   []string
		wontHubeny bool
		wontSamex  bool
		wontSamey  bool
		wontRadian bool
		wontHelp   bool
		wontError  bool
		message    string
	}{
		{[]string{"edkd"}, false, false, false, false, false, false, "ヘルプを表示"},
		{[]string{"edkd", "./testdata/result_utf.csv"}, false, false, false, false, false, false, "成功"},
		{[]string{"edkd", "0", "0", "180", "90"}, false, false, false, false, false, false, "成功"},
		{[]string{"edkd", "-H"}, true, false, false, false, false, false, "ヒュベニの公式適応"},
		{[]string{"edkd", "-x"}, false, true, false, false, false, false, "同経度"},
		{[]string{"edkd", "-y"}, false, false, true, false, false, false, "同緯度"},
		{[]string{"edkd", "-r"}, false, false, false, true, false, false, "引数：弧度法"},
		{[]string{"edkd", "-h"}, false, false, false, false, true, false, "ヘルプを表示"},
		{[]string{"edkd", "-unknown-flag"}, false, false, false, false, false, true, "未知のフラグ"},
	}
	for _, td := range testdata {
		opts, err := parseArgs(td.giveArgs)
		if (err == nil) && td.wontError {
			t.Errorf("parseArgs(%v) wont error, but got no error: %s", td.giveArgs, td.message)
		}
		if err != nil && !td.wontError {
			t.Errorf("parseArgs(%v) wont no error, but got error: %s (%s)", td.giveArgs, err.Error(), td.message)
		}
		if err != nil {
			continue
		}
		if opts.Hubeny != td.wontHubeny {
			t.Errorf("parseArgs(%v) Hubeny did not match, wont %v, but got %v", td.giveArgs, td.wontHubeny, opts.Hubeny)
		}
		if opts.samex != td.wontSamex {
			t.Errorf("parseArgs(%v) samex did not match, wont %v, but got %v", td.giveArgs, td.wontSamex, opts.samex)
		}
		if opts.samey != td.wontSamey {
			t.Errorf("parseArgs(%v) samey did not match, wont %v, but got %v", td.giveArgs, td.wontSamey, opts.samey)
		}
		if opts.radian != td.wontRadian {
			t.Errorf("parseArgs(%v) radian did not match, wont %v, but got %v", td.giveArgs, td.wontRadian, opts.radian)
		}
		if opts.help != td.wontHelp {
			t.Errorf("parseArgs(%v) help did not match, wont %v, but got %v", td.giveArgs, td.wontHelp, opts.help)
		}
	}
}
