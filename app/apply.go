package app

import (
	"bytes"
	"context"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/restmapper"
	//"fmt"
)

var cli, _ = dynamic.NewForConfig(getConfig())

func Apply(f map[string]interface{}) {
	mapper := restmapper.NewDeferredDiscoveryRESTMapper(
			memory.NewMemCacheClient(getDiscover()),
		)
	for _,v := range f {
			d := yaml.NewYAMLOrJSONDecoder(bytes.NewReader(v.([]byte)), 768) //
					for {
			var r runtime.RawExtension; if err := d.Decode(&r); err != nil { break }; // exit
			o := &unstructured.Unstructured{}
			_, gvk, _ := unstructured.UnstructuredJSONScheme.Decode(r.Raw, nil, o)
			m, _ := mapper.RESTMapping(/**/ gvk.GroupKind(), gvk.Version)
			var dri dynamic.ResourceInterface
			dri = cli.Resource(m.Resource)
			if m.Scope.Name() == meta.RESTScopeNameNamespace {
					dri = cli.Resource(m.Resource).Namespace(o.GetNamespace())
			}
			/*_, err :=*/ dri.Apply(
			context.Background(), o.GetName(), o, v1.ApplyOptions{} /*otherRes*/,
			)
					}
			// ----------------------------------- ---
	}
	//
	//\ get deploy sta
	//
}

func init() {
	//
	///
	//\\\\\\\\\\\\\\\\\\\\\\\\\
	/////
	//\
}