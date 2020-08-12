package errors

import (
	"fmt"
	"strings"
	"testing"
)

const errorMessage = "Thrown from bottom of the stack"

func bottom() error {
	return New(errorMessage)
}

func middle() error {
	return bottom()
}

func top() error {
	return middle()
}

func TestNew(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		enableStackTrace bool
	}{
		{"Error with stack trace", args{"Testing errors with a stack"}, true, true},
		{"Error without stack trace", args{"Testing errors without a stack"}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.enableStackTrace {
				EnableStackTraces()
			} else {
				DisableStackTraces()
			}

			if err := New(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
			}

			err := top()
			if err == nil {
				t.Errorf("expected call to top() to return an error with a stack trace formatted in the message")
			}

			msg := err.Error()

			var searchFor string

			if tt.enableStackTrace {
				searchFor = "Stack Trace:"
				if !strings.Contains(msg, searchFor) {
					t.Errorf("Expected stack trace to contain %s, got: %s", searchFor, msg)
				}

				searchFor = errorMessage
				if !strings.Contains(msg, searchFor) {
					t.Errorf("Expected stack trace to contain %s, got: %s", searchFor, msg)
				}
			} else {
				if msg != errorMessage {
					t.Errorf("Expected Error() to return %s, got: %s", errorMessage, msg)
				}
			}

		})
	}
}

func errorIfMissing(msg string, searchFor string) string {
	if !strings.Contains(msg, searchFor) {
		return fmt.Sprintf("Expected stack trace to contain %s, got: %s", searchFor, msg)
	}
	return ""
}
