#!/bin/bash

PAGES=( \
	"https://item.taobao.com/item.htm?spm=a1z10.1-c.w4004-16494593951.2.5d72d0c1hfIsU2&id=554830049251" \
	"https://item.taobao.com/item.htm?spm=a1z10.1-c.w4004-16494593951.4.5d72d0c11kJe88&id=553334445686" \
	)

PAGE_OPTS=""
for page in ${PAGES[@]}; do
	PAGE_OPTS="-page ${page} ${PAGE_OPTS}"
done

`pwd`/spider ${PAGE_OPTS}
