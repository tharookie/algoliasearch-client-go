// Code generated by go generate. DO NOT EDIT.

package opt

import (
    "context"
    "strings"

    "github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractPaginationLimitedTo(opts ...interface{}) *opt.PaginationLimitedToOption {
    for _, o := range opts {
        if v, ok := o.(opt.PaginationLimitedToOption); ok {
            return &v
        }
    }
    return nil
}