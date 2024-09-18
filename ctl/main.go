/*
 * Copyright The Kmesh Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at:
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
	"istio.io/istio/pkg/kube"

	logcmd "kmesh.net/kmesh/ctl/log"
)

func main() {
	rootCmd := &cobra.Command{
		Use:          "kmeshctl",
		Short:        "Kmesh command line tools to operate and debug Kmesh",
		SilenceUsage: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	kubeconfig := ""
	configContext := ""
	revision := ""

	rc, err := kube.DefaultRestConfig(kubeconfig, configContext)
	if err != nil {
		fmt.Printf("failed to get rest config: %v", err)
		os.Exit(1)
	}

	kubecli, err := kube.NewCLIClient(kube.NewClientConfigForRestConfig(rc), kube.WithRevision(revision))
	if err != nil {
		fmt.Printf("failed to create kubecli: %v", err)
		os.Exit(1)
	}

	podName := "kmesh-7jz9n"
	podNamespace := "kmesh-system"

	fw, err := kubecli.NewPortForwarder(podName, podNamespace, "", 0, 15200)
	if err != nil {
		fmt.Printf("failed to create port forwarder: %v\n", err)
		os.Exit(1)
	}
	if err := fw.Start(); err != nil {
		fmt.Printf("failed to start port forwarder: %v\n", err)
		os.Exit(1)
	}
	defer fw.Close()

	method := "GET"
	req, err := http.NewRequest(method, fmt.Sprintf("http://%s/%s", fw.Address(), "/debug/loggers?name=default"), nil)
	if err != nil {
		fmt.Printf("failed to create http request: %v\n", err)
		os.Exit(1)
	}
	resp, err := (&http.Client{
		Timeout: time.Second * 15,
	}).Do(req.WithContext(context.Background()))
	if err != nil {
		fmt.Printf("failed to do http request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("http status is %v which is not %v\n", resp.StatusCode, http.StatusOK)
		os.Exit(1)
	}

	out, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed to io.ReadAll: %v", err)
		os.Exit(1)
	}

	fmt.Printf("response body is %s", string(out))

	rootCmd.AddCommand(logcmd.NewCmd())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
