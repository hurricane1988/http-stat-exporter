/*
Copyright 2024 Hurricane Authors

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

package utils

import (
	"context"
	"crypto/tls"
	"encoding/pem"
	"errors"
	"github.com/fatih/color"
	"github.com/hurricane1988/http-stat-exporter/collector/constants"
	"io"
	"mime"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"strings"
)

func ReadClientCert(ctx context.Context, filename string) []tls.Certificate {
	Log := log.FromContext(ctx)
	// 定义证书和秘钥变量
	var (
		keyPem  []byte
		certPem []byte
	)
	if filename == "" {
		return nil
	}
	
	certFileBytes, err := os.ReadFile(filename)
	if err != nil {
		Log.Error(err, "failed to read http client certificate.",
			"filename", filename,
		)
	}
	
	for {
		block, rest := pem.Decode(certFileBytes)
		if block == nil {
			break
		}
		certFileBytes = rest
		
		if strings.HasPrefix(block.Type, constants.SuffixPrivateKey) {
			// pem.EncodeToMemory returns the PEM encoding of byte.
			keyPem = pem.EncodeToMemory(block)
		}
		if strings.HasPrefix(block.Type, constants.SuffixCertificate) {
			certPem = pem.EncodeToMemory(block)
		}
	}
	cert, err := tls.X509KeyPair(certPem, keyPem)
	if err != nil {
		Log.Error(err, "unable to load client cert and key pair")
	}
	return []tls.Certificate{cert}
}

func ParseURL(ctx context.Context, uri string) *url.URL {
	Log := log.FromContext(ctx)
	// Ensure the URI has a scheme or a double slash prefix
	if !strings.Contains(uri, "://") && !strings.HasPrefix(uri, "//") {
		uri = "//" + uri
	}
	
	// Parse the URI
	parsedURL, err := url.Parse(uri)
	if err != nil {
		Log.Error(err, "could not parse url.",
			"url", uri,
		)
	}
	
	// Set the scheme to http or https if it's empty
	if parsedURL.Scheme == constants.EmptyHttpScheme {
		parsedURL.Scheme = constants.DefaultHttpScheme
		if !strings.HasSuffix(parsedURL.Host, constants.HttpListenPort) {
			parsedURL.Scheme += "s"
		}
	}
	
	return parsedURL
}

// HeaderKeyValue parses the given HTTP header string into its key and value parts
func HeaderKeyValue(ctx context.Context, header string) (string, string) {
	Log := log.FromContext(ctx)
	// Locate the position of ':', used to split the key and value
	i := strings.Index(header, ":")
	if i == -1 {
		// Log an error as ": " is the expected delimiter
		Log.Error(errors.New("invalid header format"), "header has invalid format, missing ':'",
			"header", header,
		)
	}
	// Separate and clean the key and value
	key := strings.TrimRight(header[:i], " ")   // Removes trailing spaces from the key
	value := strings.TrimLeft(header[i:], " :") // Removes leading spaces and any possible colon from the value
	return key, value
}

func DialContext(network string) func(ctx context.Context, network, addr string) (net.Conn, error) {
	return func(ctx context.Context, _, addr string) (net.Conn, error) {
		return (&net.Dialer{
			Timeout:   constants.HttpTimeOut,
			KeepAlive: constants.HttpKeepAlive,
		}).DialContext(ctx, network, addr)
	}
}

func IsRedirect(resp *http.Response) bool {
	return resp.StatusCode > 299 && resp.StatusCode < 400
}

func CreateBody(ctx context.Context, body string) io.Reader {
	Log := log.FromContext(ctx)
	if strings.HasPrefix(body, "@") {
		filename := body[1:]
		f, err := os.Open(filename)
		if err != nil {
			Log.Error(err, "failed to open data file", "filename", filename)
		}
		return f
	}
	return strings.NewReader(body)
}

func GetFilenameFromHeaders(headers http.Header) string {
	// if the Content-Disposition header is set parse it
	if hdr := headers.Get("Content-Disposition"); hdr != "" {
		// pull the media type, and subsequent params, from
		// the body of the header field
		mt, params, err := mime.ParseMediaType(hdr)
		
		// if there was no error and the media type is attachment
		if err == nil && mt == "attachment" {
			if filename := params["filename"]; filename != "" {
				return filename
			}
		}
	}
	
	// return an empty string if we were unable to determine the filename
	return ""
}
func ReadResponseBody(ctx context.Context, saveOutput bool, outputFile string, req *http.Request, resp *http.Response) string {
	Log := log.FromContext(ctx)
	if IsRedirect(resp) || req.Method == http.MethodHead {
		return ""
	}
	
	w := io.Discard
	msg := color.CyanString("Body discarded")
	
	if saveOutput || outputFile != "" {
		filename := outputFile
		
		if saveOutput {
			// try to get the filename from the Content-Disposition header
			// otherwise fall back to the RequestURI
			if filename = GetFilenameFromHeaders(resp.Header); filename == "" {
				filename = path.Base(req.URL.RequestURI())
			}
			
			if filename == "/" {
				Log.Info("No remote filename; specify output filename with -o to save response body")
			}
		}
		
		f, err := os.Create(filename)
		if err != nil {
			Log.Error(err, "unable to create file", "filename", filename)
		}
		defer f.Close()
		w = f
		msg = color.CyanString("Body read")
	}
	
	if _, err := io.Copy(w, resp.Body); err != nil && w != io.Discard {
		Log.Error(err, "failed to read response body")
	}
	
	return msg
}

type headers []string

func (h *headers) String() string {
	var o []string
	for _, v := range *h {
		o = append(o, "-H "+v)
	}
	return strings.Join(o, " ")
}

func (h *headers) Set(v string) error {
	*h = append(*h, v)
	return nil
}

func (h *headers) Len() int {
	return len(*h)
}

func (h *headers) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *headers) Less(i, j int) bool {
	a, b := (*h)[i], (*h)[j]
	
	// server always sorts at the top
	if a == "Server" {
		return true
	}
	if b == "Server" {
		return false
	}
	
	endToEnd := func(n string) bool {
		// https://www.w3.org/Protocols/rfc2616/rfc2616-sec13.html#sec13.5.1
		switch n {
		case "Connection",
			"Keep-Alive",
			"Proxy-Authenticate",
			"Proxy-Authorization",
			"TE",
			"Trailers",
			"Transfer-Encoding",
			"Upgrade":
			return false
		default:
			return true
		}
	}
	
	x, y := endToEnd(a), endToEnd(b)
	if x == y {
		// both are of the same class
		return a < b
	}
	return x
}
