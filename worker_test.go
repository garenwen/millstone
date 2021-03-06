package millstone_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/garenwen/millstone"
)

func TestRedactURL(t *testing.T) {
	t.Parallel()

	broker := "amqp://guest:guest@localhost:5672"
	redactedURL := millstone.RedactURL(broker)
	assert.Equal(t, "amqp://localhost:5672", redactedURL)
}

func TestPreConsumeHandler(t *testing.T) {
	t.Parallel()
	
	worker := &millstone.Worker{}

	worker.SetPreConsumeHandler(SamplePreConsumeHandler)
	assert.True(t, worker.PreConsumeHandler())
}

func SamplePreConsumeHandler(w *millstone.Worker) bool {
	return true
}
