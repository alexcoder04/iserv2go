#!/bin/sh

# we dont want to proceed if something fails
set -e

# check if version number is passed
if [ -z "$1" ]; then
  echo "Please pass the version number as first argument"
  exit 1
fi

VERSION="$1"
TAG_NAME="v$VERSION"

# check if version number is in right format
echo "Checking version for right format..."
if [ echo "$VERSION" | grep -qE '^[0-9]+\.[0-9]+\.[0-9]+$' ]; then
  echo "Your version number is not in the semver format"
  exit 1
fi

# check if this version already exists
echo "Checking if $TAG_NAME already exists..."
if git tag | grep -qE "^$TAG_NAME\$"; then
  echo "The tag $TAG_NAME already exists, please increase the version"
  exit 1
fi

# check whether the version number is newer
echo "Checking if the new version is newer than latest existing..."
LATEST_EXISTING="$(git tag --sort=-version:refname | head -n 1 | tr -d v)"
if [ "$(echo "$LATEST_EXISTING\n$VERSION" | sort -r | head -n 1)" != "$VERSION" ]; then
  echo "Your specified version is older than the latest already existsing version"
  exit 1
fi

# tidy the deps
echo "Tidying the project dependencies..."
go mod tidy

# check for unstaged changes
echo "Checking for unstaged changes..."
if [ ! -z "$(git status -s)" ]; then
  echo "You have unstaged changes, please commit them first"
  exit 1
fi

# create new tag
echo "Tagging..."
git tag "$TAG_NAME"

# push everything, pushing the tags will trigger release generation on github
echo "Pushing..."
git push
git push --tags
