package integration_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/sclevine/agouti/core"

  "testing"
)

func TestWishlist(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Wishlist Suite")
}

var agoutiDriver WebDriver

var _ = BeforeSuite(func() {
  var err error

  // Choose a WebDriver

  agoutiDriver, err = PhantomJS()
  // agoutiDriver, err = Selenium():q
  // agoutiDriver, err = Chrome()

  Expect(err).NotTo(HaveOccurred())
  Expect(agoutiDriver.Start()).To(Succeed())
})

var _ = AfterSuite(func() {
  agoutiDriver.Stop()
})
