# SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Gardener contributors.
#
# SPDX-License-Identifier: Apache-2.0

.PHONY: generate-license
generate-license:
	for f in $(shell find . -name "*.go" -o -name "*.sh"); do \
		reuse addheader -r \
			--copyright="SAP SE or an SAP affiliate company and Open Component Model contributors." \
			--license="Apache-2.0" \
			$$f \
			--skip-unrecognised; \
	done
