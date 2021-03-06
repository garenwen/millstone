package millstone_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	
	"github.com/garenwen/millstone"
	"github.com/garenwen/millstone/config"
)

func TestRegisterTasks(t *testing.T) {
	t.Parallel()

	server := getTestServer(t)
	err := server.RegisterTasks(map[string]interface{}{
		"test_task": func() error { return nil },
	})
	assert.NoError(t, err)

	_, err = server.GetRegisteredTask("test_task")
	assert.NoError(t, err, "test_task is not registered but it should be")
}

func TestRegisterTask(t *testing.T) {
	t.Parallel()

	server := getTestServer(t)
	err := server.RegisterTask("test_task", func() error { return nil })
	assert.NoError(t, err)

	_, err = server.GetRegisteredTask("test_task")
	assert.NoError(t, err, "test_task is not registered but it should be")
}

func TestRegisterTaskInRaceCondition(t *testing.T) {
	t.Parallel()

	server := getTestServer(t)
	for i:=0; i<10; i++ {
		go func() {
			err := server.RegisterTask("test_task", func() error { return nil })
			assert.NoError(t, err)
			_, err = server.GetRegisteredTask("test_task")
			assert.NoError(t, err, "test_task is not registered but it should be")
		}()
	}
}

func TestGetRegisteredTask(t *testing.T) {
	t.Parallel()

	server := getTestServer(t)
	_, err := server.GetRegisteredTask("test_task")
	assert.Error(t, err, "test_task is registered but it should not be")
}

func TestGetRegisteredTaskNames(t *testing.T) {
	t.Parallel()

	server := getTestServer(t)

	taskName := "test_task"
	err := server.RegisterTask(taskName, func() error { return nil })
	assert.NoError(t, err)

	taskNames := server.GetRegisteredTaskNames()
	assert.Equal(t, 1, len(taskNames))
	assert.Equal(t, taskName, taskNames[0])
}

func TestNewWorker(t *testing.T) {
	t.Parallel()

	server := getTestServer(t)

	server.NewWorker("test_worker", 1)
	assert.NoError(t, nil)
}

func TestNewCustomQueueWorker(t *testing.T) {
	t.Parallel()

	server := getTestServer(t)

	server.NewCustomQueueWorker("test_customqueueworker", 1, "test_queue")
	assert.NoError(t, nil)
}

func getTestServer(t *testing.T) *millstone.Server {
	server, err := millstone.NewServer(&config.Config{
		Broker:        "amqp://guest:guest@localhost:5672/",
		DefaultQueue:  "millstone_tasks",
		ResultBackend: "redis://127.0.0.1:6379",
		Lock:          "redis://127.0.0.1:6379",
		AMQP: &config.AMQPConfig{
			Exchange:      "millstone_exchange",
			ExchangeType:  "direct",
			BindingKey:    "millstone_task",
			PrefetchCount: 1,
		},
	})
	if err != nil {
		t.Error(err)
	}
	return server
}
