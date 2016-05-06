package concourse_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("DeleteResourceVersion", func() {
	Context("when ATC request succeeds", func() {
		BeforeEach(func() {
			expectedURL := "/api/v1/pipelines/mypipeline/resources/myresource/versions/123/delete"
			atcServer.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("DELETE", expectedURL),
					ghttp.RespondWithJSONEncoded(http.StatusOK, ""),
				),
			)
		})

		It("sends delete resource version request to ATC", func() {
			found, err := client.DeleteResourceVersion("mypipeline", "myresource", 123)
			Expect(err).NotTo(HaveOccurred())
			Expect(found).To(BeTrue())

			Expect(atcServer.ReceivedRequests()).To(HaveLen(1))
		})
	})

	Context("when pipeline or resource does not exist", func() {
		BeforeEach(func() {
			expectedURL := "/api/v1/pipelines/mypipeline/resources/myresource/versions/123/delete"
			atcServer.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("DELETE", expectedURL),
					ghttp.RespondWithJSONEncoded(http.StatusNotFound, ""),
				),
			)
		})

		It("returns a ResourceNotFoundError", func() {
			found, err := client.DeleteResourceVersion("mypipeline", "myresource", 123)
			Expect(err).NotTo(HaveOccurred())
			Expect(found).To(BeFalse())
		})
	})

	Context("when ATC responds with an error", func() {
		BeforeEach(func() {
			expectedURL := "/api/v1/pipelines/mypipeline/resources/myresource/versions/123/delete"

			atcServer.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("DELETE", expectedURL),
					ghttp.RespondWithJSONEncoded(http.StatusBadRequest, ""),
				),
			)
		})

		It("returns an error", func() {
			found, err := client.DeleteResourceVersion("mypipeline", "myresource", 123)
			Expect(found).To(BeFalse())
			Expect(err).To(HaveOccurred())
		})
	})
})
