# Motivation

Welcome there. You're on the right place if you are looking for some resources to build up your Kubernetes Admission
Webhook.

In essence, an admission webhook is just an HTTP server that takes requests from the kube-apiserver and responds. So
there are many ways to implement it.

In this repository, we aim to find the easiest way to

- Develop
- Test
- Deploy

the admission webhooks.

All the features talked about here come from the
great [controller-runtime](https://github.com/kubernetes-sigs/controller-runtime) library. There is hard work on it, and
it's a perfect library to build Kubernetes Native applications. Enough talk. Let's get started to explore.

# Table of Contents

<!-- toc -->

- [Validating](#validating)
    - [1. Like an extension method](#1-like-an-extension-method)
    - [2. With a handler](#2-with-a-handler)
- [Defaulting](#defaulting)
    - [1. Like an extension method](#1-like-an-extension-method-1)
    - [2. With a handler](#2-with-a-handler-1)
- [Resources](#resources)

<!-- /toc -->

## Validating

The validating webhooks are for validating objects. For example, you want all pods contains some cpu or memory limits.
Or you want all deployments at least has 2 replicas. In short, we validate the objects.

We have 2 examples to do that and the main scenario is a pod is trying to be created, and we validate that pod has an
annotation like this.

```yaml
annotations:
  requested-pod-annotation: foo
```

### 1. Like an extension method

In this example, we build a validator while implementing `admission.CustomValidator` interface of
the `controller-runtime`.

Also, the testing is very easy like you are testing a simple method.

See [custom_validator.go](examples/custom_validator/custom_validator.go)

### 2. With a handler

In this example, we have more control over the incoming `admission.Request`. For example, you can take any metadata from
the request like the Auth information, dry run options to validate more.

Testing is a bit more complex than the custom validator.

See [validator_handler.go](examples/validator_handler/validator_handler.go)

## Defaulting

The defaulting (aka Mutating) webhooks are for patching objects if needed. For example, you want to label all pods which
are trying to be created in a specific namespace. Or you want to create another object when an object trying to be
created. But you got the idea. We make a patch.

We have 2 examples to do that and the main scenario is a pod is trying to be created, and we want it to have an
annotation like this.

```yaml
annotations:
  requested-pod-annotation: foo
```

So we patch it whatever happens.

### 1. Like an extension method

In this example, we build a defaulter while implementing `admission.CustomDefaulter` interface of
the `controller-runtime`.

Also, the testing is very easy like you are testing a simple method.

See [custom_defaulter.go](examples/custom_defaulter/custom_defaulter.go)

### 2. With a handler

In this example, we have more control over the incoming `admission.Request`. For example, you can take any metadata from
the request like the Auth information, dry run options, etc.

Testing is a bit more complex than the custom defaulter.

See [defaulter_handler.go](examples/defaulter_handler/defaulter_handler.go)

# Resources

- [controller-runtime](https://github.com/kubernetes-sigs/controller-runtime)
- [skaffold](https://github.com/GoogleContainerTools/skaffold)
