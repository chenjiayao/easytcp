package session

import (
	"github.com/DarthPestilane/easytcp/session/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestSessions(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			Sessions()
			wg.Done()
		}()
	}
	wg.Wait()
	assert.NotNil(t, manager)
	assert.Equal(t, manager, Sessions())
}

func TestManager_Add(t *testing.T) {
	mg := &Manager{}
	assert.NotPanics(t, func() { mg.Add(nil) })

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sess := mock.NewMockSession(ctrl)
	sess.EXPECT().ID().MinTimes(1).Return("sess id")

	mg.Add(sess)

	v, ok := mg.Sessions.Load(sess.ID())
	assert.True(t, ok)
	assert.Equal(t, v, sess)
}

func TestManager_Get(t *testing.T) {
	mg := &Manager{}
	assert.Nil(t, mg.Get("not found"))

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sess := mock.NewMockSession(ctrl)
	sess.EXPECT().ID().MinTimes(1).Return("sess id")

	mg.Sessions.Store(sess.ID(), sess)
	s := mg.Get(sess.ID())
	assert.NotNil(t, s)
	assert.Equal(t, s, sess)
}

func TestManager_Range(t *testing.T) {

}

func TestManager_Remove(t *testing.T) {

}
