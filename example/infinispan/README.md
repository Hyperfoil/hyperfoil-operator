# Local Benchmark Demo

Quick-start guide for deploying and testing the Hyperfoil Operator alongside Infinispan on a local `kind` cluster.

## Setup Cluster

```bash
# Ensure you are in the root directory
cd /path/to/hyperfoil-operator

# Create local cluster
kind create cluster --name benchmark-demo
kubectl cluster-info --context kind-benchmark-demo

# Install OLM and Infinispan Operator
curl -sL https://github.com/operator-framework/operator-lifecycle-manager/releases/download/v0.45.0/install.sh | bash -s v0.45.0
kubectl create -f https://operatorhub.io/install/infinispan.yaml
```

## Build and Deploy Operator

```bash
make install
make docker-build IMG=hyperfoil-operator:dev
kind load docker-image hyperfoil-operator:dev --name benchmark-demo
make deploy IMG=hyperfoil-operator:dev
```

## Deploy Resources

```bash
kubectl apply -f example/infinispan/infinispan.yaml
kubectl apply -f example/infinispan/hyperfoil.yaml
```

## Configure Credentials and Port-Forwarding

1. Wait for the Hyperfoil controller pod to start:
```bash
kubectl get pod hyperfoil-controller -w
```

2. In other terminal, do the port-forward to the Hyperfoil service:
```bash
kubectl port-forward svc/hyperfoil 8090:8090
```

3. Retrieve Infinispan credentials:
```bash
kubectl get secret example-infinispan-generated-secret -o jsonpath="{.data.identities\.yaml}" | base64 --decode
```

4. **Action:** Open `example/infinispan/hotrod-benchmark-demo.yaml` and update the username and password fields with the decoded credentials.

## Run Benchmark

1. Upload the benchmark definition:
```bash
curl -X POST \
  -H "Content-Type: text/vnd.yaml" \
  --data-binary "@example/infinispan/hotrod-benchmark-demo.yaml" \
  http://localhost:8090/benchmark
```

2. Open [http://localhost:8090](http://localhost:8090) in your browser to access the Hyperfoil Web UI.

3. Trigger the benchmark execution:
```text
run hotrod-benchmark-demo
```

4. Download the generated report (replace `0000` with your run ID):
```bash
curl -X GET "http://localhost:8090/run/0000/report" -o 0000.html
```

## Clean Up

```bash
kind delete cluster --name benchmark-demo
```
