// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	helm "github.com/jenkins-x/jx/pkg/helm"
)

func AnyHelmHelmer() helm.Helmer {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(helm.Helmer))(nil)).Elem()))
	var nullValue helm.Helmer
	return nullValue
}

func EqHelmHelmer(value helm.Helmer) helm.Helmer {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue helm.Helmer
	return nullValue
}
