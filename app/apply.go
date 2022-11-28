package app

import (
	"bytes"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/restmapper"
	"context"
)

var cli, _ = dynamic.NewForConfig(getConfig())

func Apply(f map[string]interface{}) {
	d := yaml.NewYAMLOrJSONDecoder(bytes.NewReader(f["k8sic.yml"].([]byte)), 768) //
	for {
			resources, _ := restmapper.GetAPIGroupResources(getDiscover())
			mapper := restmapper.NewDiscoveryRESTMapper(resources)
			var rawObjc runtime.RawExtension; if err := d.Decode(&rawObjc); err != nil { break }; // exit
			///"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
			///obj, gvk, err := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode(rawObjc.Raw, nil, nil)
			obj, gvk, err := unstructured.UnstructuredJSONScheme.Decode(rawObjc.Raw, nil, nil)
			// fmt.Println(obj); fmt.Println(gvk); fmt.Println(err)
			mping, err := mapper.RESTMapping(/**/ gvk.GroupKind(), gvk.Version)
			// cli.Resource(mping.Resource).Apply(context.Background(), name, obj)
	}
	///
}

func init() {
	//
	///
	//\\\\\
	/////
	//\
}