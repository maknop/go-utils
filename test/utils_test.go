package test

import (
	utils "github.com/mknop/go-utils/pkg"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testing functionality of comparing two data structures", func() {
	Describe("Ensure basic functionality of comparing two data structures", func() {
		Context("By removing one object from one map and adding it to another other", func() {
			It("Should check if objects exist in the GroupMap", func() {
				hListAlpha := utils.GroupMap{
					Key: "alpha",
					Value: map[string]bool{
						"obj1": true,
						"obj2": true,
					},
				}

				Expect(hListAlpha.Exists("obj1")).To(BeTrue())
				Expect(hListAlpha.Exists("obj3")).To(BeFalse())
			})

			It("Should add new entries as needed", func() {
				hListBeta := utils.GroupMap{
					Key: "beta",
					Value: map[string]bool{
						"obj3": true,
						"obj4": true,
						"obj5": true,
					},
				}

				Expect(hListBeta.Exists("obj6")).To(BeFalse())
				hListBeta.Add(hListBeta, "obj6")
				Expect(hListBeta.Exists("obj6")).To(BeTrue())
			})
		})
	})
})
