#!/bin/bash

for i in {1..27}
do
  directory="task${i}"
  mkdir $directory
done

for i in {1..27}
do
  directory="task${i}"
  touch $directory/main.go
done