// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package p

// Generic methods shouldn't trigger the monomorphization checker,
// because they cannot be implicitly instantiated when the type is.
type T[P any] struct{}

func (T[P]) m[Q any]() T[T[P]] { panic("unimplemented") }

// Non-generic methods still trigger the monomorphization checker.
type U[P /* ERROR "instantiation cycle" */ any] struct{}

func (U[P]) m() U[U[P]] { panic("unimplemented") }
