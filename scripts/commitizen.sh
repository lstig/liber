#!/bin/bash
exec < /dev/tty && yarn cz --hook || true