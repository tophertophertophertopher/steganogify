package main_test

import (
	// . "github.com/topherbullock/steganogify"
	"github.com/onsi/gomega/gexec"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var payload string = "hello world"

var _ = Describe("Main", func() {

	Context("when given valid encoding args", func() {

		BeforeEach(func() {
			command := exec.Command(steganogifyPath, "encode", "./")
			_, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ShouldNot(HaveOccurred())

		})

		It("saves an encoded gif stego file to the output path", func() {

		})

	})

	Context("given valid decoding args", func() {

		Context("and no output path supplied", func() {
			It("prints the payload to stdout", func() {

			})
		})

		Context("and output path is supplied", func() {
			It("writes the payload to the file", func() {

			})
		})
	})

})
