package singleton

import "fmt"

type Log interface {
	writeLog(message string) (err error)
}

type dbLog struct {
}
type localLog struct {
}
type cloudLog struct {
}

func (d *dbLog) writeLog(log string) (err error) {
	fmt.Printf("日志%s写入了数据库日志", log)
	return
}

func (l *localLog) writeLog(log string) (err error) {
	fmt.Printf("日志%s写入了本地日志", log)
	return
}
func (c *cloudLog) writeLog(log string) (err error) {
	fmt.Printf("日志%s写入了云端日志", log)
	return
}

func newLog(logType string) Log {
	switch logType {
	case "db":
		return &dbLog{}
	case "local":
		return &localLog{}
	case "cloud":
		return &cloudLog{}
	}
	return nil
}
