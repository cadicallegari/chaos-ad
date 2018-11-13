// +build unit

package storage_test

import (
	"testing"
	"time"

	"cadicallegari/chaos-ad/pkg/storage"
)

func TestShouldReturnProperlyStatusWhenEmpty(t *testing.T) {
	s, err := storage.New()
	if err != nil {
		t.Errorf("Error not expected: %s\n", err)
	}

	_, ok := s.Lookup("abc")
	if ok {
		t.Errorf("Should return false when key was not found")
	}
}

func TestShouldAddAndRecoveryProperly(t *testing.T) {
	s, err := storage.New()
	if err != nil {
		t.Errorf("Error not expected: %s\n", err)
	}

	key := "thekey"

	_, ok := s.Lookup(key)
	if ok {
		t.Error("Should return false when key was not found")
	}

	if err := s.Add(key, time.Now()); err != nil {
		t.Errorf("Not expected error inserting value for key '%s'", key)
	}

	_, ok = s.Lookup(key)
	if !ok {
		t.Errorf("The key '%s' should exists, but don't =(", key)
	}

	if err := s.Del(key); err != nil {
		t.Errorf("Not expected error removing the key '%s': %s", key, err)
	}

	_, ok = s.Lookup(key)
	if ok {
		t.Error("the key was removed, why it was found?")
	}

}
