
# Introduction

gomp is a go permission manager library which following ABAC ( attribute based access control ) and inspired from `Minecraft Permission Plugin`. Currently only come with some basic operation add, remove, check permission, get permission childs.

# Examples

Usually we will only doing the adding and checking, if you need removing you can do it too. 

> `adding` and `removing` is not required to call `LoadPermissions`, only call when you need check permission & get permissions childs.

```go
package main

import (
	"github.com/Oskang09/gopm"
)

func main() {
	pm := gopm.New()

	// will add 3 lines permissions
	pm.AddPermissions([]string{
		"merchant.create",
		"merchant.update",
		"merchant.store.*",
		"merchant.view.file",
		"merchant.view.report",
	})

	// will add single line permission
	pm.AddPermission("merchant.item.10")

	pm.SavePermissions()                         // [merchant.item.10 merchant.create merchant.update merchant.store.* merchant.view.file merchant.view.report]

	// load permission should be called after complete
	// adding the permission to manager
	pm.LoadPermissions()

	pm.GetPermissionChilds("merchant")           // [update store item create]
	pm.GetPermissionChilds("merchant.view")      // [file report]
	pm.HasPermission("merchant")                 // true
	pm.HasPermission("merchant.create")          // true
	pm.HasPermission("merchant.delete")          // false
	pm.HasPermission("merchant.store.create")    // true
	pm.HasPermission("merchant.store.create.10") // true
	pm.HasPermission("merchant.item.11")         // false
}
```


# Benchmark

## Code

```go
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
		"other.permission.node",
		"check.for.int.10",
	})

	pm.LoadPermissions()

	pm.HasPermission("check.for.int.10")
	pm.HasPermission("check.for.int.11")
	pm.HasPermission("other.permission.node")
	pm.HasPermission("merchant.create")
	pm.HasPermission("merchant.create.30")
	pm.HasPermission("merchant.store.1600452598542728428")
	pm.HasPermission("merchant.store.1600452598542728428.view.card")
	pm.HasPermission("merchant.store.1600452598542728428.view.name")
	pm.HasPermission("other.permission.node")
}
```

## Result

```
goos: darwin
goarch: arm64
pkg: github.com/Oskang09/gopm
BenchmarkPermissionCheck-10     1000000000               0.0000244 ns/op
PASS
ok      github.com/Oskang09/gopm        0.101s
```