port=1992
pid=$(netstat -nlp | grep :$port | awk '{print $7}' | awk -F"/" '{ print $1 }');
if [ -n "$pid" ]; then 
    kill -15 $pid
fi
