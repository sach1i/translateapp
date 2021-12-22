package logging_test

import (
	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap/zapcore"
	"testing"
	l "translateapp/internal/logging"
)

func TestNewLogger(t *testing.T) {
	t.Parallel()

	logger := l.NewLogger("", true)
	if logger == nil {
		t.Fatal("expected logger to never be nil")
	}
}

func TestLevelToZapLevls(t *testing.T) {
	t.Parallel()

	cases := []struct {
		input string
		want  zapcore.Level
	}{
		{input: l.LevelDebug, want: zapcore.DebugLevel},
		{input: l.LevelInfo, want: zapcore.InfoLevel},
		{input: l.LevelWarning, want: zapcore.WarnLevel},
		{input: l.LevelError, want: zapcore.ErrorLevel},
		{input: l.LevelCritical, want: zapcore.DPanicLevel},
		{input: l.LevelAlert, want: zapcore.PanicLevel},
		{input: l.LevelEmergency, want: zapcore.FatalLevel},
		{input: "unknown", want: zapcore.WarnLevel},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.input, func(t *testing.T) {
			t.Parallel()

			got := l.LevelToZapLevel(tc.input)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Fatalf("mismatch (-want, +got):\n%s", diff)
			}
		})
	}
}
