#!/bin/bash

LIST_DOCKER_IMAGE_HASHES=$(docker images chain4energy --format "{{ title .Tag }}" | awk '!/Debug/ && !/V[0-9-]+/' | awk '{print tolower($0)}')
