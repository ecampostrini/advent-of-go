#!/usr/bin/env bash

day=$1

if [ -d $day ]; then
  echo "Directory '${day}' already exists. No action will be done"
  exit 0
fi

mkdir -p ${day} && cd ${day} && touch day${day}.go && cd .. && ls ${day}/*
