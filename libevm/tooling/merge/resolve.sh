#!/usr/bin/env bash

#  Copyright 2025 the libevm authors.
#
#  The libevm additions to go-ethereum are free software: you can redistribute
#  them and/or modify them under the terms of the GNU Lesser General Public License
#  as published by the Free Software Foundation, either version 3 of the License,
#  or (at your option) any later version.
#
#  The libevm additions are distributed in the hope that they will be useful,
#  but WITHOUT ANY WARRANTY; without even the implied warranty of
#  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser
#  General Public License for more details.
#
#  You should have received a copy of the GNU Lesser General Public License
#  along with the go-ethereum library. If not, see
#  <http://www.gnu.org/licenses/>.

set -o pipefail;
set -eux;

cd "$(git rev-parse --show-toplevel)";

# Files for which we are taking full control have any incoming changes
# discarded.
OVERWRITING=(".github/workflows/go.yml" ".golangci.yml")
git checkout --ours -- "${OVERWRITING[@]}";
git add "${OVERWRITING[@]}";

# Files deleted by geth are deleted locally too because they're guaranteed to
# not be our code due to the libevm naming convention.
git status | grep "deleted by them" | awk '{print $NF}' | xargs -r git rm || echo "No files deleted upstream";

# There are a number of conflicts due to the module renaming, which can be
# resolved mechanically by accepting the incoming changes. This is a blunt
# first-pass approach, accepting entire files. It will start an interactive
# terminal UI with instructions.
cd ./libevm/tooling/merge;
ACCEPTING=$(git status | grep "both modified" | awk '{print $NF}' | grep -P "\.go$" | go run . resolve theirs);
if [[ -n "${ACCEPTING[*]// }" ]]; then # $x// removes whitespace
    echo "${ACCEPTING[@]}" | xargs git checkout --theirs -- ;
    echo "${ACCEPTING[@]}" | xargs git add;
fi
cd -;
