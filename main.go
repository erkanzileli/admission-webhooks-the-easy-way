package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

var (
	certsDir = flag.String("certs-dir", "/app/certs", "certificates directory that contains the tls.crt and tls.key files")
)

func init() {
	flag.Parse()
}

func main() {
	// set up a Manager
	log.Info("setting up manager")
	mgr, err := manager.New(config.GetConfigOrDie(), manager.Options{})
	if err != nil {
		log.Error(err, "unable to set up overall controller manager")
		os.Exit(1)
	}

	// create a webhook server
	log.Info("setting up webhook server")
	hookServer := &webhook.Server{
		Port:    8443,
		CertDir: *certsDir,
	}
	if err = mgr.Add(hookServer); err != nil {
		panic(err)
	}

	// register webhooks
	log.Info("registering webhooks to the webhook server")

	hookServer.Register("/handle-validate-v1-pod", &webhook.Admission{Handler: NewPodValidatingWebhook()})
	hookServer.Register("/validate-v1-pod", &webhook.Admission{Handler: NewCustomPodValidator()})

	hookServer.Register("/handle-mutate-v1-pod", &webhook.Admission{Handler: NewPodMutatingWebhook()})
	hookServer.Register("/mutate-v1-pod", &webhook.Admission{Handler: NewCustomPodDefaulter()})

	log.Info("starting manager")
	if err = mgr.Start(signals.SetupSignalHandler()); err != nil {
		log.Error(err, "unable to run manager")
		os.Exit(1)
	}
}
