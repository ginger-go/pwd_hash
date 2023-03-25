package pwd_hash_test

import (
	"github.com/ginger-go/pwd_hash"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("test: pwd_hash", func() {

	var examplePwd1 = "123456"
	var examplePwd2 = "654321"

	It("should hash the password", func() {
		pwd1Hash := pwd_hash.Hash(examplePwd1)
		Expect(pwd_hash.Verify(examplePwd1, pwd1Hash)).To(BeTrue())

		pwd2Hash := pwd_hash.Hash(examplePwd2)
		Expect(pwd_hash.Verify(examplePwd2, pwd2Hash)).To(BeTrue())
	})

	It("should not verify the wrong password", func() {
		pwd1Hash := pwd_hash.Hash(examplePwd1)
		Expect(pwd_hash.Verify(examplePwd2, pwd1Hash)).To(BeFalse())

		pwd2Hash := pwd_hash.Hash(examplePwd2)
		Expect(pwd_hash.Verify(examplePwd1, pwd2Hash)).To(BeFalse())
	})
})
