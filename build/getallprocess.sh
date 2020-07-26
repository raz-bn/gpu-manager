#!/bin/bash
PID=$1
ps --forest -o pid= $(ps -e --no-header -o pid,ppid|awk -vp=$PID 'function r(s){print s;s=a[s];while(s){sub(",","",s);t=s;sub(",.*","",t);sub("[0-9]+","",s);r(t)}}{a[$2]=a[$2]","$1}END{r(p)}') 