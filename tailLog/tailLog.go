package tailLog

import (
	"fmt"
	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
)

func Init(path string) (err error) {
	tailObj, err = tail.TailFile(path, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		fmt.Println("init tail failed, err:", err)
		return
	}
	return
}

func ReadChan() <-chan *tail.Line{
	return tailObj.Lines
}