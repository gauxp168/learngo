package taillog

import (
	"context"
	"fmt"
	"github.com/hpcloud/tail"
	"simqo.com/mygospace/learngo/example/logagent/kafka"
)

type TailTask struct {
	path string
	topic string
	instance *tail.Tail
	ctx context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	tailObj = &TailTask{
		path:path,
		topic:topic,
	}
	return
}

func (t *TailTask) init()  {
	config := tail.Config{
		ReOpen:true,
		Follow:true,
		Location:&tail.SeekInfo{Offset:0, Whence:2},
		MustExist:false,
		Poll:true,
	}	
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
	}
	go t.run()
}

func (t *TailTask) run()  {
	for  {
		select {
		case <-t.ctx.Done():
			fmt.Printf("tail task:%s_%s 结束了...\n",t.path, t.topic)
			return
		case line := <- t.instance.Lines:
			kafka.SendToChan(t.topic, line.Text)

		}
	}
}
