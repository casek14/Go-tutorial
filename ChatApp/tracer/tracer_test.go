package tracer

import (
	"testing"
	"bytes"
)

func TestNew(t *testing.T){
	var buf bytes.Buffer
	tracer := New(&buf)

	if tracer == nil{
		t.Error("Return from nil should not be nil !")
	}else{
		tracer.Trace("Hello tracer pkg.")
		if buf.String() != "Hello tracer pkg."{
			t.Errorf("Trace should not write  '%s'.",buf.String())
		}
	}
	}

	func TestOff(t testing.T){
		var silentTracer Tracer = Off()
		silentTracer.Trace("something")
	}