// Code generated by go generate. DO NOT EDIT.

package opt

import (
    "context"
    "strings"

    "github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractAroundLatLng(opts ...interface{}) *opt.AroundLatLngOption {
    for _, o := range opts {
        if v, ok := o.(opt.AroundLatLngOption); ok {
            return &v
        }
    }
    return nil
}