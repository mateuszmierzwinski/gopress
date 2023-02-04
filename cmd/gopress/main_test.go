package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBindAddress(t *testing.T) {
	testCases := []struct {
		testName       string
		execFunc       func(string) string
		resultExpected string
	}{
		{
			testName: "BIND_IP and PORT present",
			execFunc: func(s string) string {
				return map[string]string{
					"BIND_IP": "0.0.0.0",
					"PORT":    "8080",
				}[s]
			},
			resultExpected: "0.0.0.0:8080",
		}, {
			testName: "BIND_IP only present",
			execFunc: func(s string) string {
				return map[string]string{
					"BIND_IP": "1.2.3.4",
				}[s]
			},
			resultExpected: "0.0.0.0:8080",
		}, {
			testName: "PORT only present",
			execFunc: func(s string) string {
				return map[string]string{
					"PORT": "9090",
				}[s]
			},
			resultExpected: "0.0.0.0:8080",
		}, {
			testName: "none present present",
			execFunc: func(s string) string {
				return ""
			},
			resultExpected: "0.0.0.0:8080",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			result := getBindAddress(testCase.execFunc)
			assert.Equal(t,
				testCase.resultExpected,
				result,
				"Test %s failed. Expected result '%s' but got '%s'",
				testCase.testName,
				testCase.resultExpected,
				result)
		})
	}
}
