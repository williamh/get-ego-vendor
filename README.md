# get-ego-vendor

The purpose of this utility is to generate the  contents of the
EGO_VENDOR variable which is used in Gentoo ebuilds for go packages that
use Go modules and do not vendor their dependencies.

The package uses go modules if the source tree includes files named
go.mod and go.sum. If it also includes a directory named vendor in its top
level source directory, it vendors its dependencies and this
utility is not needed.

If it doesn't vendor its dependencies, the output from running this
utility should be added to the Gentoo ebuild for the package.

Below is an example of how to run this utility on a package foo.

```
$cd /path/to/foo
$ go mod vendor
$ get-ego-vendor > ego-vendor.txt
```

Then ego-vendor.txt should be inserted into the ebuild.
