#!/bin/bash

if [ -f .env ]; then
  source .env
fi

HOME=${HOME:-$HOME}
WORK=${WORK:-$PWD}

echo -e "Workdir on host: $WORK\nStarting shell, Enjoy your work !"

docker build -f ./Dockerfile -t work . >/dev/null
docker run -ti --rm \
    --hostname "factory" \
    --volume "$WORK":"/home/worker/work" \
    --volume "$HOME/.gitconfig":"/home/worker/.gitconfig" \
    --volume "$HOME/.ssh:/home/worker/.ssh" \
    work