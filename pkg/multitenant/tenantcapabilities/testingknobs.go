// Copyright 2023 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package tenantcapabilities

import "github.com/cockroachdb/cockroach/pkg/base"

// TestingKnobs contain testing helpers which are used by various components
// that enable tenant capabilities.
type TestingKnobs struct {
	// WatcherTestingKnobs can be used to test the tenant capabilities Watcher.
	WatcherTestingKnobs base.ModuleTestingKnobs

	// AuthorizerSkipAdminSplitCapabilityChecks, if set, skips capability checks
	// for AdminSplit requests in the Authorizer for secondary tenants.
	AuthorizerSkipAdminSplitCapabilityChecks bool
}

// ModuleTestingKnobs is part of the base.ModuleTestingKnobs interface.
func (t *TestingKnobs) ModuleTestingKnobs() {}

var _ base.ModuleTestingKnobs = (*TestingKnobs)(nil)
