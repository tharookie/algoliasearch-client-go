// Code generated by go generate. DO NOT EDIT.

package opt

import (
    "context"
    "strings"

    "github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractAutoGenerateObjectidIfNotExist(opts ...interface{}) *opt.AutoGenerateObjectidIfNotExistOption {
    for _, o := range opts {
        if v, ok := o.(opt.AutoGenerateObjectidIfNotExistOption); ok {
            return &v
        }
    }
    return nil
}