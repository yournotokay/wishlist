package integration_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/sclevine/agouti/core"
  //. "github.com/sclevine/agouti/matchers"

  . "github.com/yournotokay/wishlist/server"
)

var _ = Describe("UserLogin", func() {
  var page Page

  BeforeEach(func() {
    go StartMyApp(3123)

    var err error
    page, err = agoutiDriver.Page(Use().Browser("firefox"))
    Expect(err).NotTo(HaveOccurred())
  })

  AfterEach(func() {
    page.Destroy()
  })

  It("should manage user authentication", func() {
        By("redirecting the user to the login form from the home page", func() {
            Expect(page.Navigate("http://localhost:3123")).To(Succeed())
            })
    })
})
