#Overview

Deploy a node server on a cluster.

#Parts
- Vagrant
- coreos
- docker
- etcd
- confd
- fleet

##Services
A service can be defined as a process that you want to run on a machine. systemd is a service manager that you can use to run, stop, init services. And fleet is a manager of that, which is used to manage services in a cluster.

A service is composed of these parts.

- Unit
- Service
- Install
- X-Fleet

```
// Example service file
// app@.service
[Unit]
Description=App on port %i

Requires=etcd.service
Requires=docker.service
Requires=app-discovery@%i.service

After=etcd.service
After=docker.service
Before=app-discovery@%i.service

[Service]
TimeoutStartSec=0
KillMode=none

EnvironmentFile=/etc/environment

ExecStartPre=-/usr/bin/docker kill app%i
ExecStartPre=-/usr/bin/docker rm app%i
ExecStartPre=/usr/bin/docker pull tutum/hello-world

ExecStart=/usr/bin/docker run --rm --name app%i -p ${COREOS_PRIVATE_IPV4}:%i:80 tutum/hello-world

ExecStop=/usr/bin/docker stop app%i

[Install]
WantedBy=multi-user.target

[X-Fleet]
X-Conflicts=app@*.service





// app-discovery@.service
[Unit]
Description=Announce app%i service

Requires=etcd.service
Requires=app@%i.service

After=etcd.service
After=app@%i.service
BindsTo=app@%i.service

[Service]
EnvironmentFile=/etc/environment

ExecStart=/bin/sh -c '\
    while true; do \
        curl -f ${COREOS_PRIVATE_IPV4}:%i; \
        if [ $? -eq 0 ]; then \
            etcdctl set /services/app%i \'{"host": "%H", "ipv4_private_addr": ${COREOS_PRIVATE_IPV4}, "ipv4_public_addr": ${COREOS_PUBLIC_IPV4}, "port": %i}\' --ttl 60; \
        else \
            etcdctl rm /services/app%i; \
        fi; \
        sleep 45; \
    done'

ExecStop=/usr/bin/etcdctl rm /services/app%i

[X-Fleet]
X-ConditionMachineOf=app@%i.service


```

Filename<br>
name@port.service<br>
%i - references string AFTER '@'

Requires and After directives indicate the defined service can only start after those other services have started. ```=-``` means allow failure.

TimeoutStartSec allows docker pull to run without systemd timing out.
KillMode, allows us to stop docker without systemd thinking it failed.

BindsTo directive binds this service to another service.

A app service should be associated with a sidekick discovery service whose responsibility is to monitor the life of the main app and report to etcd so other services can discovery services through etcd.

####Templates

	unit@.service
	
The @ indicates that this file is a tempalte file. When this template is instantiated the id will be put after the @.

fleetctl and systemd both accept symbolic links.

	ln -s app@.service app@8080.service
	
you can pass a directory into fleetctl to start everything at once

	fleetctl start instances/*


##Fleet
Fleet is a manager of services in a cluster. You can access the command line tool through one of the machines in the cluster. It uses etcd to communicate across machines.

ssh requires ssh-agent run 

	// if ssh is required
	eval $(ssh-agent)
	ssh-add
	
	// 
	ssh-A core@coreos_node_public_IP
	
	// inside one of the machines in the cluster
	// lists machines
	fleetctl list-machines
	// ssh into the machine running that service
	fleetctl ssh <service>
	
####Starting a service
	
	// register service to fleetctl
	fleetctl submit foo.service
	
	// list registered unit files
	fleetctl list-unit-files
	
	// to view unit file
	fleetctl cat foo.service
	
	// schedule a service in cluster
	fleetctl load foo.service
	
	// list units and their schedules/status
	fleetctl list-units
	
	// actually start service
	fleetctl start foo.service
	
	// stop service
	fleetctl stop foo.service
	
	// unload from systems it was loaded in
	fleetctl unload foo.service
	
	// remove from fleet completely
	fleetctl destroy foo.service
	
	// detailed status of service
	fleetctl status foo.service
	
	// see journal entry of service
	fleetctl journal [--lines 20] [-f (continue to pass logs)] foo.service

	
Summary

- submit
- load
- start
- stop
- unload
- destroy

sample flow

	fleetctl submit app@.service
	fleetctl submit app-discovery@.service
	fleetctl load app@80.service
	fleetctl load app-discovery@80.service
	fleetctl start app@80.service
	
fleetctl can also just take a whole folder

	ln -s app@.service instances/app@9000.service
	ln -s app@.service instances/app@9001.service
	ln -s app-discovery@.service instances/app-discovery@9000.service
	ln -s app-discovery@.service instances/app-discovery@9001.service
	// start all the symlinks in instances dir
	fleetctl start instances/*
	


##etcd
etcd is a high availablity key value store. It is similar to redis except it was built as a cluster with consistency and failure in mind. It's write and read speeds are not as fast as redis but it has 100% uptime and best used for cluster coordination.

	https://discovery.etcd.io/new
	// returns
	https://discovery.etcd.io/8d27bd52a298d9eedcbf4c37dde4346c
	// visiting this link will give you a json of connected nodes in the cluster
	
in the cloud-config file (user-data)
provide etcd with the discovery token for every new cluster (even reloads)

You can retrieve the url by looking at the etcd.service that a given machine is using.

	cat /run/systemd/system/etcd.service.d/20-cloudinit.conf
	
1. first machine connects, hits url, finds no other nodes, designates itself as cluster leader.
2. subsequent machines connect, and connects to machines found
3. data found is stored in etcd /_etcd/machines

```
etcdctl ls /_etcd/machines --recursive
```

####etcdctl
Commandline

	// lists keys
	etcdctl ls <key> [--recursive]
	
	// get key value
	etcdctl get <key>
	
	// create new dir
	etcdctl mkdir /<dirname>
	
	// make a key
	etcdctl mk <key> <value>
	
	// update a key
	etcdctl update <key> <value>
	
	// ttl time to live (seconds)
	etcdctl mkdir /shortlived --ttl 10
	// update ttl
	etcdctl updatedir /shortlived --ttl 50
	
	// change value or key/create
	etcdctl set <key> <value>
	
	// same effect for dir
	etcdctl setdir <key>
	
	// remove key
	etcdctl rm <key> [--recursive]
	
	// only remove empty dir or key
	etcdctl rmdir <key>
	
	// watch key for changes
	etcdctl watch [--recursive] <key>
	
	etcdctl exec-watch --recursive <key> -- echo 'hi'
	
HTTP/JSON api<br>
Port 4001 for key queries<br>
Port 7001 for etcd settings

	curl -L http://127.0.0.1:4001/v2/keys/
	// inside docker use
	http://172.17.42.1:4001
	
	// take in flags as query params example
	curl -L http://127.0.0.1:4001/v2/keys/?recursive=true
	
	// get version	
	curl -L http://127.0.0.1:4001/version
	// view cluster leader relationship
	curl -L http://127.0.0.1:4001/v2/stats/leader
	// view self stats
	curl -L http://127.0.0.1:4001/v2/stats/self
	// stats about operations
	curl -L http://127.0.0.1:4001/v2/stats/store
	
	// settings use port 7001
	// set etcd settings
	curl -L http://127.0.0.1:7001/v2/admin/config
	// list machines
	curl -L http://127.0.0.1:7001/v2/admin/machines

##Confd
	
	# download confd move to bin
	wget -O confd https://github.com/kelseyhightower/confd/releases/download/v0.6.3/confd-0.6.3-linux-amd64
	cp confd /opt/bin/.
	chmod +x /opt/bin/confd
	mkdir -p /etc/confd/{conf.d,templates}
	
	cp /home/core/share/confd/nginx.toml /etc/confd/conf.d/.
	cp /home/core/share/confd/nginx.tmpl /etc/confd/templates/.
	
the first section downloads confd and moves it to bin, and creates the folder structure confd expects to work with.

confd/conf.d/ will store the confd configs<br>
confd/templates will store the templates to work with

confd watches for key changes in etcd and on change will replace a file with a specified template file.

	// example app.toml
	[template]
	src = "nginx.tmpl"
	dest = "/etc/nginx/sites-enabled/app.conf"
	keys = [
    	"/services"
	]
	owner = "root"
	mode = "0644"
	check_cmd = "/usr/sbin/nginx -t"
	reload_cmd = "/usr/sbin/service nginx reload"
	
src - confd looks in the confd/templates folder for this file<br>
dest - where to move the src file on key change<br>
keys - keys to watch for in etcd<br>
owner - owner of the moved file<br>
mode - mode of the moved file<br>
check_cmd - command to use to check syntax of rendered template file<br>
reload_cmd - command to run when reloading<br>


	
##Operations

	fleetctl submit <service>
	fleetctl load <service>@port
	fleetctl start <service>@port
	
	etcdctl get /announce/services/<service>port
