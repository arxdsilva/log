package log

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func Test_Logger_withfield_should_add_to_stdout(t *testing.T) {
	r, w, err := os.Pipe()
	require.Nil(t, err)

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	l := New("service-name", WithLevel("INFO"), WithOutput(w))
	l.WithFields(
		zap.String("somefield", "somevalue"),
		zap.String("somefield2", "somevalue2")).Info("some log")

	w.Close()
	out := <-outC

	require.Contains(t, out, "somefield")
	require.Contains(t, out, "somevalue")
	require.Contains(t, out, "some log")
}

func Test_Logger_DEBUG_should_not_add_to_stdout(t *testing.T) {
	r, w, err := os.Pipe()
	require.Nil(t, err)

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	l := New("service-name", WithLevel("ERROR"), WithOutput(w))
	l.WithField(zap.String("somefield", "somevalue")).Debug("some log")

	w.Close()
	out := <-outC

	require.Nil(t, err)
	require.NotContains(t, out, "somefield")
	require.NotContains(t, out, "somevalue")
	require.NotContains(t, out, "some log")
}

func Test_Logger_WARN_should_not_add_to_stdout(t *testing.T) {
	r, w, err := os.Pipe()
	require.Nil(t, err)

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	l := New("service-name", WithLevel("ERROR"), WithOutput(w))
	l.WithField(zap.String("somefield", "somevalue")).Warn("some log")

	w.Close()
	out := <-outC

	require.NotContains(t, out, "somefield")
	require.NotContains(t, out, "somevalue")
	require.NotContains(t, out, "some log")
}

func Test_Logger_ERROR_should_add_to_stdout(t *testing.T) {
	r, w, err := os.Pipe()
	require.Nil(t, err)

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	l := New("service-name", WithLevel("ERROR"), WithOutput(w))
	l.WithField(zap.String("somefield", "somevalue")).Error("some error")

	w.Close()
	out := <-outC

	require.Nil(t, err)
	require.Contains(t, out, "somefield")
	require.Contains(t, out, "somevalue")
	require.Contains(t, out, "some error")
}

func Test_Logger_InfoWithFields_should_add_to_stdout(t *testing.T) {
	r, w, err := os.Pipe()
	require.Nil(t, err)

	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	l := New("service-name", WithLevel("INFO"), WithOutput(w))
	l.Info("some log", zap.String("somefield", "somevalue"))

	w.Close()
	out := <-outC

	require.Contains(t, out, "somefield")
	require.Contains(t, out, "somevalue")
	require.Contains(t, out, "some log")
}

func Test_Logger_withfields_should_add_to_stdout(t *testing.T) {
	r, w, err := os.Pipe()
	require.Nil(t, err)

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	l := New("service-name", WithLevel("INFO"), WithOutput(w))
	l.WithFields(zap.String("vessel", "italy"), zap.String("vessel2", "france")).Info("some log")

	w.Close()
	out := <-outC

	require.Contains(t, out, "vessel")
	require.Contains(t, out, "italy")
	require.Contains(t, out, "vessel2")
	require.Contains(t, out, "france")
	require.Contains(t, out, "some log")
}
