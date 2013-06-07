package m3u8_parser

import "testing"

func verify(t *testing.T, testname string, testcase, input, output, expected string) {
        if output != expected {
                t.Errorf("should %s: %s with input = %s: output %s != %s", testname, testcase, input, output, expected)
        }
}