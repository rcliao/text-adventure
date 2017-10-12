#!/bin/sh

./build.sh && rsync -av main root@quiz.do:/opt/text-adventure && ssh 'root@quiz.do' 'supervisorctl reload'
