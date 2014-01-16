package test_test

import (
	. "static_proxy_server/lib/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var numberset = []struct {
	x      int
	y      int
	result int
}{
	{1, 2, 3},
	{1, 2, 3},
	{2, 4, 6},
}

var _ = Describe("Test", func() {
  Describe("#Add", func(){
    It("should eq expect", func(){
      for _, set := range numberset {
        Expect(Add(set.x, set.y)).To(Equal(set.result))
      }
    })
  })
})
