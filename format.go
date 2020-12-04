package rslog

import (
	"bytes"
	"strconv"
	"strings"
	"time"
)

func itoa(i int, wid int) (buf []byte) {
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	b[bp] = byte('0' + i)
	buf = append(buf, b[bp:]...)
	return
}
func formatHeader(buf *bytes.Buffer, t time.Time, pcInfo PcInfo, projectName, levelStr string) {
	year, month, day := t.Date()
	buf.Write(itoa(year, 4))
	buf.WriteString("-")
	buf.Write(itoa(int(month), 2))
	buf.WriteString("-")
	buf.Write(itoa(day, 2))
	buf.WriteString(" ")

	hour, min, sec := t.Clock()
	buf.Write(itoa(hour, 2))
	buf.WriteString(":")
	buf.Write(itoa(min, 2))
	buf.WriteString(":")
	buf.Write(itoa(sec, 2))
	buf.WriteString(".")
	buf.Write(itoa(t.Nanosecond()/1e3, 6))
	buf.WriteString(" ")

	if strings.Contains(pcInfo.File, projectName) {
		buf.WriteString(".")
		n := strings.Index(pcInfo.File, projectName)
		buf.WriteString(strings.Join(strings.Split(pcInfo.File, "")[n+len(projectName):], ""))
	} else {
		buf.WriteString(pcInfo.File)
	}
	buf.WriteString(":")
	buf.WriteString(strconv.Itoa(pcInfo.Line))
	buf.WriteString(": ")

	buf.WriteString("[")
	buf.WriteString(GetFuncName(pcInfo, projectName))
	buf.WriteString("] ")

	buf.WriteString("[")
	buf.WriteString(levelStr)
	buf.WriteString("] ")
}
