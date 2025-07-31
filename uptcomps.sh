#! /usr/bin/env bash

for comp in ./components/*; do
    comp=${comp##*/}
    templui -f add $comp
done

