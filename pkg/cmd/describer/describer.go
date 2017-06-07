package describer

import (
	"fmt"
	"reflect"

	"github.com/golang/glog"
	"github.com/k8sdb/apimachinery/client/clientset"
	"github.com/k8sdb/cli/pkg/cmd/decoder"
	"k8s.io/kubernetes/pkg/kubectl"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/runtime"
)

type Describer interface {
	Describe(object runtime.Object, describerSettings *kubectl.DescriberSettings) (output string, err error)
}

func NewDescriber(f cmdutil.Factory) Describer {
	return newHumanReadableDescriber(f)
}

type handlerEntry struct {
	describeFunc reflect.Value
	args         []reflect.Value
}

type humanReadableDescriber struct {
	cmdutil.Factory
	extensionsClient clientset.ExtensionInterface
	handlerMap       map[reflect.Type]*handlerEntry
}

func newHumanReadableDescriber(f cmdutil.Factory) *humanReadableDescriber {
	restClonfig, _ := f.ClientConfig()
	describer := &humanReadableDescriber{
		Factory:          f,
		extensionsClient: clientset.NewForConfigOrDie(restClonfig),
		handlerMap:       make(map[reflect.Type]*handlerEntry),
	}
	describer.addDefaultHandlers()
	return describer
}

func (h *humanReadableDescriber) addDefaultHandlers() {
	h.Handler(h.describeElastic)
	h.Handler(h.describePostgres)
	h.Handler(h.describeSnapshot)
	h.Handler(h.describeDormantDatabase)
}

func (h *humanReadableDescriber) Handler(describeFunc interface{}) error {
	describeFuncValue := reflect.ValueOf(describeFunc)
	if err := h.validateDescribeHandlerFunc(describeFuncValue); err != nil {
		glog.Errorf("Unable to add describe handler: %v", err)
		return err
	}

	objType := describeFuncValue.Type().In(0)

	h.handlerMap[objType] = &handlerEntry{
		describeFunc: describeFuncValue,
	}
	return nil
}

func (h *humanReadableDescriber) validateDescribeHandlerFunc(describeFunc reflect.Value) error {
	if describeFunc.Kind() != reflect.Func {
		return fmt.Errorf("invalid describe handler. %#v is not a function", describeFunc)
	}
	funcType := describeFunc.Type()
	if funcType.NumIn() != 2 || funcType.NumOut() != 2 {
		return fmt.Errorf("invalid describe handler." +
			"Must accept 2 parameters and return 2 value.")
	}

	if funcType.In(1) != reflect.TypeOf((*kubectl.DescriberSettings)(nil)) ||
		funcType.Out(0) != reflect.TypeOf((string)("")) ||
		funcType.Out(1) != reflect.TypeOf((*error)(nil)).Elem() {
		return fmt.Errorf("invalid describe handler. The expected signature is: "+
			"func handler(item %v, describerSettings *kubectl.DescriberSettings) (string, error)", funcType.In(0))
	}
	return nil
}

func (h *humanReadableDescriber) Describe(obj runtime.Object, describerSettings *kubectl.DescriberSettings) (string, error) {
	kind := obj.GetObjectKind().GroupVersionKind().Kind
	switch obj.(type) {
	case *runtime.UnstructuredList, *runtime.Unstructured, *runtime.Unknown:
		if objBytes, err := runtime.Encode(clientset.ExtendedCodec, obj); err == nil {

			if decodedObj, err := decoder.Decode(kind, objBytes); err == nil {
				obj = decodedObj
			}
		}
	}

	t := reflect.TypeOf(obj)
	if handler := h.handlerMap[t]; handler != nil {
		args := []reflect.Value{reflect.ValueOf(obj), reflect.ValueOf(describerSettings)}
		resultValue := handler.describeFunc.Call(args)
		if err := resultValue[1].Interface(); err != nil {
			return resultValue[0].Interface().(string), err.(error)
		}

		return resultValue[0].Interface().(string), nil
	}

	return "", fmt.Errorf(`kubedb doesn't support: "%v"`, kind)
}
