package info_test

import (
	"runtime/debug"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/openshift/backplane-cli/pkg/info"
	infoMock "github.com/openshift/backplane-cli/pkg/info/mocks"
)

var _ = Describe("Info", func() {
	var (
		mockCtrl             *gomock.Controller
		mockBuildInfoService *infoMock.MockBuildInfoService
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockBuildInfoService = infoMock.NewMockBuildInfoService(mockCtrl)
		info.DefaultBuildInfoService = mockBuildInfoService
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("When getting build version", func() {
		It("Should return the pre-set Version is available", func() {
			info.Version = "whatever"

			version := info.DefaultInfoService.GetVersion()
			Expect(version).To(Equal("whatever"))
		})
		It("Should return a version when go bulid info is available and there is no pre-set Version", func() {
			info.Version = ""
			mockBuildInfoService.EXPECT().GetBuildInfo().Return(&debug.BuildInfo{
				Main: debug.Module{
					Version: "v1.2.3",
				},
			}, true).Times(1)

			version := info.DefaultInfoService.GetVersion()
			Expect(version).To(Equal("1.2.3"))
		})
		It("Should return an unknown when no way to determine version", func() {
			info.Version = ""
			mockBuildInfoService.EXPECT().GetBuildInfo().Return(nil, false).Times(1)
			version := info.DefaultInfoService.GetVersion()
			Expect(version).To(Equal("unknown"))
		})
	})
})
