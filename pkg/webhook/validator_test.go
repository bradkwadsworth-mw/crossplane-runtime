/*
Copyright 2022 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package webhook

import (
	"context"
	"testing"

	"github.com/crossplane/crossplane-runtime/pkg/test"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/crossplane/crossplane-runtime/pkg/errors"
)

var errBoom = errors.New("boom")

func TestValidateCreate(t *testing.T) {
	type args struct {
		obj runtime.Object
		fns []ValidateCreateFn
	}
	type want struct {
		err error
	}
	cases := map[string]struct {
		reason string
		args
		want
	}{
		"Success": {
			reason: "Functions without errors should be executed successfully",
			args: args{
				fns: []ValidateCreateFn{
					func(_ context.Context, _ runtime.Object) error {
						return nil
					},
				},
			},
		},
		"Failure": {
			reason: "Functions with errors should return with error",
			args: args{
				fns: []ValidateCreateFn{
					func(_ context.Context, _ runtime.Object) error {
						return errBoom
					},
				},
			},
			want: want{
				err: errBoom,
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			v := NewValidator(WithValidateCreationFns(tc.fns...))
			err := v.ValidateCreate(context.TODO(), tc.args.obj)
			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nValidateCreate(...): -want, +got\n%s\n", tc.reason, diff)
			}
		})
	}
}

func TestValidateUpdate(t *testing.T) {
	type args struct {
		oldObj runtime.Object
		newObj runtime.Object
		fns    []ValidateUpdateFn
	}
	type want struct {
		err error
	}
	cases := map[string]struct {
		reason string
		args
		want
	}{
		"Success": {
			reason: "Functions without errors should be executed successfully",
			args: args{
				fns: []ValidateUpdateFn{
					func(_ context.Context, _, _ runtime.Object) error {
						return nil
					},
				},
			},
		},
		"Failure": {
			reason: "Functions with errors should return with error",
			args: args{
				fns: []ValidateUpdateFn{
					func(_ context.Context, _, _ runtime.Object) error {
						return errBoom
					},
				},
			},
			want: want{
				err: errBoom,
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			v := NewValidator(WithValidateUpdateFns(tc.fns...))
			err := v.ValidateUpdate(context.TODO(), tc.args.oldObj, tc.args.newObj)
			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nValidateUpdate(...): -want, +got\n%s\n", tc.reason, diff)
			}
		})
	}
}

func TestValidateDelete(t *testing.T) {
	type args struct {
		obj runtime.Object
		fns []ValidateDeleteFn
	}
	type want struct {
		err error
	}
	cases := map[string]struct {
		reason string
		args
		want
	}{
		"Success": {
			reason: "Functions without errors should be executed successfully",
			args: args{
				fns: []ValidateDeleteFn{
					func(_ context.Context, _ runtime.Object) error {
						return nil
					},
				},
			},
		},
		"Failure": {
			reason: "Functions with errors should return with error",
			args: args{
				fns: []ValidateDeleteFn{
					func(_ context.Context, _ runtime.Object) error {
						return errBoom
					},
				},
			},
			want: want{
				err: errBoom,
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			v := NewValidator(WithValidateDeletionFns(tc.fns...))
			err := v.ValidateDelete(context.TODO(), tc.args.obj)
			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nValidateDelete(...): -want, +got\n%s\n", tc.reason, diff)
			}
		})
	}
}
