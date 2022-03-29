service="./nicloud"
pid_file=goftp.pid
 
 
function start() {
    ${service} &
    if [[ $? -eq 0 ]]; then
        echo $! > ${pid_file}
    else exit 1
    fi
}

start
