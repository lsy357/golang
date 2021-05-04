package pool

import (
	"context"
	"errors"
	"hello/utils"
	"net"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type Task interface {
	process() error
}

type ReadTask struct {
	conn net.Conn
	fd   int
}

func (task *ReadTask) process() error {
	var err error
	// process read task
	return err
}

type WriteTask struct {
	conn     net.Conn
	fd       int
	deviceID int64
	payload  []byte
}

func (task *WriteTask) process() error {
	var err error
	// process write task
	return err
}

type taskQueue struct {
	ctx       context.Context
	taskQueue chan Task
	handle    HandleFunc
	size      int
	wg        sync.WaitGroup
	status    int8 // 0: closed, 1: running
	done      context.CancelFunc
}

func (queue *taskQueue) Start() {
	queue.status = 1
	for i := queue.size; i > 0; i-- {
		queue.wg.Add(1)
		go func(ctx context.Context) {
			defer func() {
				utils.Recovery()
				queue.wg.Done()
			}()
			running := true
			for running {
				select {
				case <-ctx.Done():
					running = false
				case task := <-queue.taskQueue:
					err := task.process()
					if err != nil {
						logrus.WithError(err).Error("handle task fail")
					}
				}
			}
		}(queue.ctx)
	}
}

func (queue *taskQueue) Submit(task Task) error {
	if queue.status == 0 {
		return errors.New("queue is closed")
	}
	var err error
	select {
	case queue.taskQueue <- task:
		// handle
	case <-time.After(500 * time.Microsecond):
		err := errors.New("conn submit fail")
		logrus.WithError(err).Error("Connection channel is blocked!")
		// send to kafka
		return err
	}
	return err
}

func (queue *taskQueue) Close() {
	queue.status = 0
	queue.done()
	queue.wg.Wait()

out:
	for {
		select {
		case task := <-queue.taskQueue:
			err := task.process()
			if err != nil {
				logrus.WithError(err).Error("handle task fail")
			}
		case <-time.After(100 * time.Microsecond):
			break out
		}
	}

	close(queue.taskQueue)
}

type HandleFunc func(conn net.Conn)

func newWorkerPool(_size int) *taskQueue {
	_ctx, cancel := context.WithCancel(context.Background())
	return &taskQueue{
		ctx:       _ctx,
		taskQueue: make(chan Task, 3000),
		handle:    handleRequest,
		size:      _size,
		done:      cancel,
		wg:        sync.WaitGroup{},
	}
}

func handleRequest(conn net.Conn) {
	// handle client request
}
