package gopm

import (
	"testing"
)

func BenchmarkPermissionCheck(b *testing.B) {
	pm := New()
	pm.AddPermissions([]string{
		"merchant.edit",
		"merchant.create.5",
		"merchant.create.*",
		"merchant.store.1600452598542728428.view",
		"merchant.store.1600452598542728428.edit",
		"merchant.store.1600452598542728428.view.*",
		"merchant.store.1600452598542728429.view.*",
		"other.permission.node",
		"check.for.int.10",
	})

	pm.LoadPermissions()

	pm.GetPermissionChilds("merchant.store")

	pm.HasPermission("check.for.int.10")
	pm.HasPermission("check.for.int.11")
	pm.HasPermission("other.permission.node")
	pm.HasPermission("merchant.create")
	pm.HasPermission("merchant.create.30")
	pm.HasPermission("merchant.store.1600452598542728428")
	pm.HasPermission("merchant.store.1600452598542728428.view.card")
	pm.HasPermission("merchant.store.1600452598542728428.view.name ")
	pm.HasPermissions([]string{"org.fast-reply.view"})
}
