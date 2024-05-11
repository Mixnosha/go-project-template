package jsonwrt

import (
	"os"
	"time"
)

type JsonWriter struct {
	f         *os.File
	oldTime   string
	pathToDir string
}

func NewJsonWriter(pathDir string) *JsonWriter {
	t := getHour()
	writer := &JsonWriter{
		oldTime:   t,
		pathToDir: pathDir,
	}
	writer.SetNewWriter()
	return writer
}

func getHour() string {
	return time.Now().Format("2006-01-02:15h")
}

func (w *JsonWriter) SetNewWriter() error {
	t := getHour()
	if w.f != nil {
		w.f.Close()
	}
	filename := w.pathToDir + t + ".json"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	w.f = file

	return nil
}

func (w *JsonWriter) Write(b []byte) (n int, err error) {

	if getHour() == w.oldTime {
		return w.f.Write(b)
	} else {
		w.SetNewWriter()
		return w.f.Write(b)
	}
}

