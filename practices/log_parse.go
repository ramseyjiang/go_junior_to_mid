package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	const formatTime = "2006-01-02 15:04:05" // magic time in Go, 15 means 0h-23h, if change 15 to 03, mean 1h-12h

	// a line of log
	logsExample := `[2021-08-27T17:39:54.173Z] "GET /healthz HTTP/1.1" 200 - 0 61 225 - "111.114.195.106,10.0.0.11" "okhttp/3.12.1" "0557b0bd-4c1c-4c7a-ab7f-2120d67bee2f" "example.com" "172.16.0.1:8080"`

	// your defined log format, $_ means the content you can ignore.
	logsFormat := `\[$time_stamp\] \"$http_method $request_path $_\" $response_code - $_ $_ $_ - \"$ips\" \"$_\" \"$_\" \"$website_name\" \"$_\"`

	// transform all the defined variable into a regex-readable named format
	regexFormat := regexp.MustCompile(`\$([\w_]*)`).ReplaceAllString(logsFormat, `(?P<$1>.*)`)

	// compile the result
	result := regexp.MustCompile(regexFormat)

	// find all the matched data from the logsExample
	matches := result.FindStringSubmatch(logsExample)

	for i, k := range result.SubexpNames() {
		// ignore the first and the $_
		if i == 0 || k == "_" {
			continue
		} else {
			if k == "time_stamp" {
				dateTime, _ := time.Parse(time.RFC3339Nano, matches[i])
				matches[i] = dateTime.Format(formatTime)
			}

			// print the defined variable,
			// %-15 means all arrows will in the same column.
			// The first s mean k, the second s means matches[i]
			fmt.Printf("%-15s => %s\n", k, matches[i])
		}
	}
}
