// Copyright (c) Red Hat, Inc.
// SPDX-License-Identifier: MPL-2.0

package iso_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIso(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ISO Builder Suite")
}
