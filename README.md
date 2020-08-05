# Chaos Engineering with Litmus and Okteto

This post will show you how you can use Litmus and Okteto to start chaos testing your 

# Instal Litmus, the experiments and the target application

[![Develop on Okteto](https://okteto.com/develop-okteto.svg)](https://cloud.okteto.com/deploy?repository=https://github.com/okteto/litmus-on-okteto)


# Call the Application

Open a local terminal, and run the following script. This will simply call the application once very second. Replace the with the URL of your application.

```console
./call.sh https://hello-world-rberrelleza.cloud.okteto.net
```

# Start the chaos experiment

Apply the ChaosEngine manifest to trigger the experiment.

To start the experiment run the following command on a second terminal.
```
kubectl apply -f chaos/engine.yaml
```

## Observe Chaos results

When an experiment is created, a ChaosResult resource will created to hold the result of the experiment. The `status.verdict` is set to `Awaited` when the experiment is in progress, eventually changing to either `Pass` or `Fail`.

```console
kubectl describe chaosresult pod-killer-chaos-experiment -n nginx
```

NOTE: ChaosResult CR name will be <chaos-engine-name>-<chaos-experiment-name>