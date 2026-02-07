set -euo pipefail

VERSION="${1:-}"
if [ -z "$VERSION" ]; then
  echo "Usage: bash release.sh <version>" >&2
  echo "Example: bash release.sh 1.0.0" >&2
  exit 1
fi

BRANCH="$(git rev-parse --abbrev-ref HEAD)"
if [ "$BRANCH" != "publish" ]; then
  echo "Error: you must be on branch 'publish' (current: '$BRANCH')" >&2
  exit 1
fi

TAG="$VERSION"
case "$TAG" in
  v*) ;;
  *) TAG="v${TAG}" ;;
esac

if git rev-parse -q --verify "refs/tags/${TAG}" >/dev/null; then
  echo "Error: tag '${TAG}' already exists" >&2
  exit 1
fi

git fetch
git pull
git merge master
git tag "${TAG}"
git push origin "${TAG}"
git checkout master

# finish
exit 0
