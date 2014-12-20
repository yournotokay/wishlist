package integration_test

import (
  "strconv"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/sclevine/agouti/core"
  . "github.com/sclevine/agouti/matchers"

  . "github.com/yournotokay/wishlist/server"
)

var _ = Describe("UserLogin", func() {
  var page Page
  port := 3133
  baseUri :=  "http://localhost:" + strconv.Itoa(port)

  BeforeEach(func() {

    var err error
    page, err = agoutiDriver.Page(Use().Browser("firefox"))
    Expect(err).NotTo(HaveOccurred())
  })

  AfterEach(func() {
    StopMyApp(port)
    page.Destroy()
  })

  It("should manage user authentication", func() {
    By("redirecting the user to the login form from the home page", func() {
      Expect(page.Navigate(baseUri)).To(Succeed())
      Expect(page).To(HaveURL(baseUri + "/login"))
    })
  })
})
