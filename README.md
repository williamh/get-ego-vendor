# get-ego-vendor

The purpose of this utility is to generate the  contents of the
EGO_SUM variable which is used in Gentoo ebuilds for go packages that
use Go modules.

The package uses go modules if the source tree includes a file named
go.mod and optionally a file named go.sum in its top-level directory.

If go.sum does not exist or vendor does exist, this utility is not
needed. Otherwise, the output from running this utility should be
inserted into the ebuild.

Below is an example of how to run this utility on a package foo.

```
$cd /path/to/foo
$ get-ego-vendor > ego-sum.txt
```

Then ego-sum.txt should be inserted into the ebuild.
