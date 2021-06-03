// Copyright (c) 2018 Yandex LLC. All rights reserved.
// Use of this source code is governed by a MPL 2.0
// license that can be found in the LICENSE file.
// Author: Vladimir Skipor <skipor@yandex-team.ru>

package coreutil

import "a.yandex-team.ru/load/projects/pandora/core"

func ReturnSampleIfBorrowed(s core.Sample) {
	borrowed, ok := s.(core.BorrowedSample)
	if !ok {
		return
	}
	borrowed.Return()
}
