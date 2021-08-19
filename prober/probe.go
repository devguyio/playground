/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"k8s.io/apimachinery/pkg/util/sets"

	"knative.dev/networking/pkg/apis/networking/v1alpha1"
	"knative.dev/networking/pkg/status"
	"knative.dev/pkg/logging"
)

type TargetLister struct {
}

func (t TargetLister) ListProbeTargets(ctx context.Context, ingress interface{}) ([]status.ProbeTarget, error) {
	u, _ := url.Parse("http://www.google.com")
	uls := []*url.URL{u}

	return []status.ProbeTarget{
		{
			PodIPs:  sets.NewString("127.0.0.1"),
			PodPort: "8080", Port: "8080", URLs: uls,
		},
	}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("K-Network-Hash", "44136fa355b3678a1146ad16f7e8649e94fb4fc21fe77e8310c060f61caaff8a")
	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/view", handler)
	http.HandleFunc("/view/health", handler)
	go http.ListenAndServe(":8080", nil)

	ctx := context.Background()
	logger := logging.FromContext(ctx)

	statusProber := status.NewProber(
		logger.Named("status-manager"),
		TargetLister{},
		func(ing interface{}) {
			logger.Info("Ready callback triggered.")
		})
	in := v1alpha1.Ingress{}
	d := make(chan struct{})
	statusProber.Start(d)
	time.Sleep(3 * time.Second)
	for i := 0; i < 3; i++ {
		r, _ := statusProber.IsReady(ctx, &in)
		logger.Infof("Ready: %t", r)
		if r {
			break
		}
		time.Sleep(3 * time.Second)
	}

}
