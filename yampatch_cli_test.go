package main_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("the yampatch cli", func() {
	var (
		session *gexec.Session
		err     error
	)

	Context("When you run the command without arguments", func() {
		BeforeEach(func() {
			command := exec.Command(binPath)
			session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
		})

		It("Runs, and exits 0", func() {
			Eventually(session, 5).Should(gexec.Exit(0))
			Eventually(session.Out, 5).Should(gbytes.Say(`YAML`))
		})

		// It("Prints an error/usage message", func() {
		// 	Eventually(session, 5).Should(gexec.Exit(0))
		// 	Eventually(session.Out, 5).Should(gbytes.Say(`some error`))
		// })
	})

	Context("When you run the command with just the input file argument", func() {
		BeforeEach(func() {
			command := exec.Command(binPath, "fixtures/simple_target.yml")
			session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
		})

		It("prints the contents of that file, unchanged, to stdout", func() {
			Eventually(session, 5).Should(gexec.Exit(0))
			Eventually(session.Out, 5).Should(gbytes.Say(`---
my-key1: my-key1-starting-value
`))
		})

	})
})
