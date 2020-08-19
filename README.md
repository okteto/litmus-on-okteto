# Chaos Engineering with Litmus and Okteto Cloud

This post will show you how we can use [Litmus](https://litmuschaos.io) and [Okteto Cloud](https://cloud.okteto.com) to show you how to start Chaos testing your Kubernetes applications.


## Why Do We Need Chaos Engineering?

Cloud Native applications are, by definition, highly distributed, elastic, resistent to failure and loosely coupled. That's easy to say, and even diagram. But how do we validate that our applications will perform as expected under different failure conditions?

Enter Chaos engineering. [Chaos engineering](https://en.wikipedia.org/wiki/Chaos_engineering) is the discipline of experimenting on a software system in production in order to build confidence in the system's capability to withstand turbulent and unexpected conditions. Chaos Engineering is a great tool to help us find weaknesses and misconfiguration in our services. It is particular important on Cloud Native applications, which, due to their distributed and elastic nature, need to be resilient by default.

[Litmus](https://litmus.io) is a framework for practicing Chaos Engineering in Cloud Native environments. Litmus provides a chaos-operator, a large set of chaos experiments in its [hub](https://hub.litmuschaos.io/), detailed documentation, quick Demo, and a friendly community.


## Prerequisites

For this example you'll need a free Okteto Cloud account, and `kubectl` installed on your local machine.

## Deploy your Development Environment

[![Develop on Okteto](https://okteto.com/develop-okteto.svg)](https://cloud.okteto.com/deploy?repository=https://github.com/okteto/litmus-on-okteto&branch=community-call)

This will automatically deploy the following resources on your Okteto Cloud account:
- The Litmus Chaos operator
- The Pod-Delete experiment
- The Hello-World application, configured with 2 replicas

## Download your Kubernetes Credentials

If this is the first time you use Okteto Cloud, you'll need to [download your Kubernetes credentials](https://okteto.com/docs/cloud/credentials) in order to run `kubectl` commands. 

## Get the Code

```console
git clone https://github.com/okteto/litmus-on-okteto
```

## Call the Application

Open a local terminal, and run the script below. This script will call the application once very second. You can get the application's endpoint for the Okteto Cloud URL.

```console
./call.sh <your application endpoint>
```

```console
Hello world!
Hello world!
Hello world!
...
```



## Start the chaos experiment

Open `chaos/engine.yaml` with your favorite text editor. This is the file that defines which experiment to run, and on which resources. 

The file has three main sections:
- `appinfo`: This tells the Litmus operator which application to target. You have to specify a label selector and the type of resource.
- `experiments`: A list of experiments to run. In this case, we are running the [Pod Delete experiment](https://docs.litmuschaos.io/docs/pod-delete/).
- `experiments.spec.components`: The experiment-specific value overrides. In this case, we are telling the experiment to kill 1 pod over 30 seconds. The allowed files are defined on the [ChaosExperiment manifest](chaos/experiment.yaml). 

Open a second terminal, and run the following command to start the chaos experiment:

```console
kubectl apply -f chaos/engine.yaml
```

```console
chaosengine.litmuschaos.io/pod-killer-chaos created
```

## Observe Chaos results

The experiment will kill one of our application's pods. If you run the command below once the experiment has started, you'll see how a random pod is killed and then automatically recreated:

```console
kubectl get pod -l=app=hello-world
```

```console
NAME                                 READY   STATUS              RESTARTS   AGE
hello-world-75947547d4-2fcbc         1/1     Running             0          57m
hello-world-75947547d4-c6wsv         0/1     ContainerCreating   0          10s
```

While the experiment is running, keep an eye on the `call` process we started. Notice how the calls were never interrupted? That's because our application is resilient to pod destruction 💪🏻!

### Reporting Information

When an experiment is created, a ChaosResult resource will created to hold the result of the experiment. The `status.verdict` is set to `Awaited` when the experiment is in progress, eventually changing to either `Pass` or `Fail`.

```console
kubectl describe chaosresult pod-killer-chaos-pod-delete
```

```console
Name:         pod-killer-chaos-pod-delete
Namespace:    rberrelleza
Labels:       name=pod-killer-chaos-pod-delete
Annotations:  <none>
API Version:  litmuschaos.io/v1alpha1
Kind:         ChaosResult
Metadata:
  Creation Timestamp:  2020-08-05T21:14:05Z
  Generation:          5
  Resource Version:    165298631
  Self Link:           /apis/litmuschaos.io/v1alpha1/namespaces/rberrelleza/chaosresults/pod-killer-chaos-pod-delete
  UID:                 a7f50d28-1f14-4a03-9013-94e72d69eb72
Spec:
  Engine:      pod-killer-chaos
  Experiment:  pod-delete
Status:
  Experimentstatus:
    Fail Step:  N/A
    Phase:      Running
    Verdict:    Awaited
Events:
  Type    Reason   Age   From                     Message
  ----    ------   ----  ----                     -------
  Normal  Summary  45m   experiment-l0k004-5x2fj  pod-delete experiment has been Passed
```

## Conclusion

In this post, we showed how you can deploy a replicable development environment that includes an application, the LitmusChaos operator, and your chaos experiment, all in one click. Then, we ran a chaos experiment, validating that our application is resilient to a pod failure. 

Interested in Kubernetes and Chaos Engineering? Join us in the [Okteto](https://kubernetes.slack.com/messages/CM1QMQGS0/) and [Litmus](https://kubernetes.slack.com/messages/CNXNB0ZTN/) slack communities and let's get the conversation started!