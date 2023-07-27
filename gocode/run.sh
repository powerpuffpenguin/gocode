#!/bin/bash
set -e

cd `dirname "$BASH_SOURCE"`

./gocode dir /opt/google/go/src/ -r \
    -E /opt/google/go/src/cmd/compile/internal/types2/testdata/fixedbugs/issue50372.go \
    -E /opt/google/go/src/go/parser/testdata/issue42951/not_a_file.go/invalid.go \
    -R /opt/google/go/src/cmd/compile/internal/syntax/testdata/ \
    -R /opt/google/go/src/cmd/compile/internal/types2/testdata/check/ \
    -R /opt/google/go/src/cmd/compile/internal/types2/testdata/examples/ \
    -R /opt/google/go/src/cmd/compile/internal/types2/testdata/fixedbugs/ \
    -R /opt/google/go/src/go/types/testdata/check/ \
    -R /opt/google/go/src/go/types/testdata/examples/ \
    -R /opt/google/go/src/go/types/testdata/fixedbugs/
