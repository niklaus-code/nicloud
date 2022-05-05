yum install epel-release.noarch -y && yum install bridge-utils -y

ifconfig eth0 0.0.0.0 up

brctl addbr br85

brctl addif br85 eth0 &&

ifconfig br85 10.0.0.1/24 up &&

route add default gw 10.0.0.254
