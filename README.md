# `gh collab-montage`

[![CI](https://github.com/spenserblack/gh-collab-montage/actions/workflows/ci.yml/badge.svg)](https://github.com/spenserblack/gh-collab-montage/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/spenserblack/gh-collab-montage/graph/badge.svg?token=jhBJJl4oQ2)](https://codecov.io/gh/spenserblack/gh-collab-montage)

Combine avatars of all contributors into one image. Just call `gh collab-montage` and this extension will query all contributors, fetch their avatars, and
combine them all into one PNG. This extension also provides a public library to provide utilities for making your own montage.

## Installation

```shell
gh extension install spenserblack/gh-collab-montage
```

## Features

- Specify if you want the avatars to be circles or squares
- Specify the margin between avatars

## Possible Features

These features will be implemented if they are requested. Raise an issue if you're interested!

- Generating montages from other sources (collaborators, organization members, a manually-defined file, etc.)
- ???

## Example

The following example was generated with `gh collab-montage -R o2sh/onefetch --size 100 --margin 5`.

![montage](https://github.com/spenserblack/gh-collab-montage/assets/8546709/4c2b673a-2ae0-4958-a2a6-69a56185ed23)


## Credits

This extension was inspired by [avatar-montage](https://github.com/benbalter/avatar-montage).
