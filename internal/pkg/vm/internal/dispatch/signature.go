package dispatch

import (
	"bytes"
	"reflect"

	"github.com/filecoin-project/go-filecoin/internal/pkg/encoding"
	"github.com/filecoin-project/specs-actors/actors/runtime"
	"github.com/pkg/errors"
)

// MethodSignature wraps a specific method and allows you to encode/decodes input/output bytes into concrete types.
type MethodSignature interface {
	// ArgInterface returns the typed argument expected by the actor method.
	ArgInterface(argBytes []byte) (interface{}, error)
	// ReturnInterface returns the methods typed return.
	ReturnInterface(returnBytes []byte) (interface{}, error)
}

type methodSignature struct {
	method method
}

var _ MethodSignature = (*methodSignature)(nil)

// ArgInterface implement MethodSignature.
func (ms *methodSignature) ArgInterface(argBytes []byte) (interface{}, error) {
	// decode arg1 (this is the payload for the actor method)
	t := ms.method.Type().In(1)
	v := reflect.New(t)

	// Dragons: fix this on our encoding library
	obj := v.Elem().Interface()
	if unmarsh, ok := obj.(runtime.CBORUnmarshaler); ok {
		buf := bytes.NewBuffer(argBytes)
		auxv := reflect.New(t.Elem())
		obj = auxv.Interface()
		unmarsh = obj.(runtime.CBORUnmarshaler)
		if err := unmarsh.UnmarshalCBOR(buf); err != nil {
			return nil, err
		}
		return unmarsh, nil
	}

	if err := encoding.Decode(argBytes, v.Interface()); err != nil {
		return nil, errors.Wrap(err, "failed to decode bytes as method argument")
	}

	// dereference the extra pointer created by `reflect.New()`
	return v.Elem().Interface(), nil
}

// ReturnInterface implement MethodSignature.
func (ms *methodSignature) ReturnInterface(returnBytes []byte) (interface{}, error) {
	// decode arg1 (this is the payload for the actor method)
	t := ms.method.Type().Out(0)
	v := reflect.New(t)
	if err := encoding.Decode(returnBytes, v.Interface()); err != nil {
		return nil, errors.Wrap(err, "failed to decode return bytes for method")
	}

	return v.Interface(), nil
}
