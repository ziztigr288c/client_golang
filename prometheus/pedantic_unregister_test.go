package prometheus

import (
	"testing"
)

func TestPedanticRegistryUnregister(t *testing.T) {
	r := NewPedanticRegistry()

	g := NewGauge(GaugeOpts{
		Name: "test_gauge",
		Help: "test help",
	})

	if err := r.Register(g); err != nil {
		t.Fatalf("unexpected error registering gauge: %s", err)
	}

	// Registering again should fail.
	if err := r.Register(g); err == nil {
		t.Fatal("expected error registering gauge twice, got nil")
	}

	// Unregister the gauge.
	if !r.Unregister(g) {
		t.Fatal("expected Unregister to return true")
	}

	// Registering again should succeed now.
	if err := r.Register(g); err != nil {
		t.Fatalf("unexpected error registering gauge after unregister: %s", err)
	}

	// Gather should succeed without errors.
	mfs, err := r.(Gatherer).Gather()
	if err != nil {
		t.Fatalf("unexpected error gathering: %s", err)
	}

	if len(mfs) != 1 {
		t.Fatalf("expected 1 metric family, got %d", len(mfs))
	}
}
