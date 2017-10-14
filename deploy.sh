#!/bin/sh

./build.sh && rsync -av . root@quiz.do:/opt/text-adventure && ssh 'root@quiz.do' 'supervisorctl reload'
