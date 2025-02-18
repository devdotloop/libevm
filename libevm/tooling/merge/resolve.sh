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

OVERWRITING=(".github/workflows/go.yml" ".golangci.yml")
git checkout --ours -- "${OVERWRITING[@]}";
git add "${OVERWRITING[@]}";

git status | grep "deleted by them" | awk '{print $NF}' | xargs -r git rm || echo "No files deleted upstream";

cd ./libevm/tooling/merge;
ACCEPTING=$(git status | grep "both modified" | awk '{print $NF}' | grep -P "\.go$" | go run . resolve theirs);
if [[ -n "${ACCEPTING[*]// }" ]]; then # $x// removes whitespace
    echo "${ACCEPTING[@]}" | xargs git checkout --theirs -- ;
    echo "${ACCEPTING[@]}" | xargs git add;
fi
cd -;
