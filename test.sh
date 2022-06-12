#!/bin/bash

set -e

if command -v go &> /dev/null; then
  go install github.com/vgchl/terminal-ui
	terminal_ui=$(go env GOPATH)/bin/terminal-ui
else
  terminal_ui="docker run -t --rm terminal-ui"
fi

$terminal_ui header --title "Workspace setup"
$terminal_ui header "Installing dependencies"
$terminal_ui alert-info "You'll need to do something."
$terminal_ui alert-done "Installed dependencies"
$terminal_ui header "Code Generation"
$terminal_ui header --style light "Code Generation"
$terminal_ui header --style flat "Code Generation"
$terminal_ui alert-done "Code generated"
$terminal_ui alert-warn "Oh oh!"
$terminal_ui alert-error "Oops.. that failed"
