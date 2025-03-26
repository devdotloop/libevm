# Release

In the following, we create a release candidate tag `v1.13.14-0.2.0.rc.4`.

1. Create two branches, usually from the `main` branch:

    ```bash
    git fetch origin main:main
    git checkout main
    git checkout -b release/v1.13.14-0.2.0.rc.4
    git push -u origin release/v1.13.14-0.2.0.rc.4
    git checkout -b myusername/release/v1.13.14-0.2.0.rc.4
    git push -u origin myusername/release/v1.13.14-0.2.0.rc.4
    ```

    The `myusername/release/v1.13.14-0.2.0.rc.4` branch will be used to add "release modifications" and will target the branch `release/v1.13.14-0.2.0.rc.4` in a pull request.
1. Run script `./cherrypick.sh` which cherry picks all Geth commits listed in [cherrypicks](cherrypicks)
    - you may have to resolve conflicts
1. Modify [params/version.libevm.go](/params/version.libevm.go)
    - Change the `LibEVMVersion` to the correct final version name
    - Change the `LibEVMReleaseType` to the correct release type, for example `ReleaseCandidate`
    - If planning a release candidate: set `libEVMReleaseCandidate` to the correct number; in this case `4`
    - If needed: change the `LibEVMVersionMajor`, `LibEVMVersionMinor` and `LibEVMVersionPatch` numbers
1. Commit your modifications to [params/version.libevm.go](/params/version.libevm.go) with a commit title `chore: release v1.13.14-0.2.0.rc.4`.
1. Push your modified branch to the remote `git push`
1. Open a pull request from your modified branch `myusername/release/v1.13.14-0.2.0.rc.4` and targeting `release/v1.13.14-0.2.0.rc.4`, for example with `https://github.com/ava-labs/libevm/compare/release/v1.13.14-0.2.0.rc.4...myusername/release/v1.13.14-0.2.0.rc.4?expand=1`. Set the tile to "chore: release `v1.13.14-0.2.0.rc.4`"
1. Wait for all the checks to pass on the pull request
1. Fast forward merge your branch in the release branch:

    ```bash
    git checkout release/v1.13.14-0.2.0.rc.4
    git merge --ff-only myusername/release/v1.13.14-0.2.0.rc.4
    git push
    ```

    This will also close the ongoing pull request.
1. Finally create your tag and push it using the release branch:

    ```bash
    git tag v1.13.14-0.2.0.rc.4
    git push origin v1.13.14-0.2.0.rc.4
    ```

1. (optional) you can then cleanup with:

    ```bash
    git branch -D release/v1.13.14-0.2.0.rc.4 myusername/release/v1.13.14-0.2.0.rc.4
    git push -d origin myusername/release/v1.13.14-0.2.0.rc.4
    ```
