#!/bin/bash
find . -name "*.sh" | cut -d "." -f2 | rev | cut -d "/" -f1 | rev
