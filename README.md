# CEFCO

Cloud-Edge File Coordination Operator

## Run For Test

```bash
# clone repo
$ git clone https://github.com/inspursoft/cefco

# create crd
$ kubectl apply -f artifacts/crd.yaml
# run for test
$ make run

# You have to build the Reader and Writer docker image before creating.
# Get more info here:
#   http://git.inspur.com/middleware/idx-component/docker
# And then you need a stream and a consumer within the nats-jetstream
# create an example
$ kubectl apply -f artifacts/examples/example-mini.yaml
# delete the example
$ kubectl delete -f artifacts/examples/example-mini.yaml
```

## Build

```bash
# For executable
$ make filesync-controller VERSION=0.1.0 noRace=T
# run
$ ./filesync-controller --kubeconfig /root/.kube/config

# For docker image
$ make filesync-controller-docker filesyncVersion=0.1.0
# push it to your registry
# run
$ kubectl apply -f artifacts/crd.yaml
$ kubectl apply -f artifacts/rbac.yaml
$ kubectl apply -f artifacts/deployment.yaml

# create an example
$ kubectl apply -f artifacts/examples/example-mini.yaml
# delete the example
$ kubectl delete -f artifacts/examples/example-mini.yaml
```

## Tools

```bash
# Create a stream
$ nats --user s1 --password s1 str add example --subjects="example.>" --storage=memory --retention=workq --discard=new --max-msgs=-1 --max-bytes=-1 --max-age=-1 --max-msg-size=-1 --dupe-window=2s --replicas=1
# Create a consumer
nats --user s1 --password s1 con add example mini --deliver=all --replay=instant --filter="example.mini" --max-deliver=-1 --max-pending=1 --pull
```