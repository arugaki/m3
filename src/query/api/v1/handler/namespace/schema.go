// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package namespace

import (
	"io/ioutil"
	"net/http"
	"path"

	clusterclient "github.com/m3db/m3/src/cluster/client"
	"github.com/m3db/m3/src/cluster/kv"
	"github.com/m3db/m3/src/dbnode/namespace/kvadmin"
	"github.com/m3db/m3/src/query/api/v1/handler"
	"github.com/m3db/m3/src/query/generated/proto/admin"
	"github.com/m3db/m3/src/query/util/logging"
	xerrors "github.com/m3db/m3/src/x/errors"
	"github.com/m3db/m3/src/x/net/http"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

var (
	// M3DBSchemaURL is the url for the M3DB schema handler.
	M3DBSchemaURL = path.Join(handler.RoutePrefixV1, M3DBServiceSchemaPathName)

	// SchemaHTTPMethod is the HTTP method used with this resource.
	SchemaHTTPMethod = http.MethodPost
)

// SchemaHandler is the handler for namespace schema upserts.
type SchemaHandler Handler

// For unit test purpose.
var newAdminService = kvadmin.NewAdminService

// NewSchemaHandler returns a new instance of SchemaHandler.
func NewSchemaHandler(client clusterclient.Client) *SchemaHandler {
	return &SchemaHandler{client: client}
}

func (h *SchemaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := logging.WithContext(ctx)

	md, rErr := h.parseRequest(r)
	if rErr != nil {
		logger.Error("unable to parse request", zap.Error(rErr))
		xhttp.Error(w, rErr.Inner(), rErr.Code())
		return
	}

	opts := handler.NewServiceOptions("kv", r.Header, nil)
	resp, err := h.Add(md, opts)
	if err != nil {
		if err == kv.ErrNotFound || xerrors.InnerError(err) == kv.ErrNotFound {
			logger.Error("namespaces metadata key does not exist", zap.Error(err))
			xhttp.Error(w, err, http.StatusInternalServerError)
			return
		}
		if err == kvadmin.ErrNamespaceNotFound || xerrors.InnerError(err) == kvadmin.ErrNamespaceNotFound {
			logger.Error("namespace does not exist", zap.Error(err))
			xhttp.Error(w, err, http.StatusNotFound)
			return
		}

		logger.Error("unable to deploy schema to namespace", zap.Error(err))
		xhttp.Error(w, err, http.StatusBadRequest)
		return
	}

	xhttp.WriteProtoMsgJSONResponse(w, &resp, logger)
}

type SchemaAddRequest struct {
	Name      string            `yaml:"name"`
	MsgName   string            `yaml:"msgName"`
	ProtoName string            `yaml:"protoName"`
	ProtoMap  map[string]string `yaml:"protoMap,flow"`
}

func (h *SchemaHandler) parseRequest(r *http.Request) (*admin.NamespaceSchemaAddRequest, *xhttp.ParseError) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, xhttp.NewParseError(err, http.StatusBadRequest)
	}

	var schemaAddReq SchemaAddRequest
	if err := yaml.Unmarshal(body, &schemaAddReq); err != nil {
		return nil, xhttp.NewParseError(err, http.StatusBadRequest)
	}
	return &admin.NamespaceSchemaAddRequest{
		Name:      schemaAddReq.Name,
		MsgName:   schemaAddReq.MsgName,
		ProtoName: schemaAddReq.ProtoName,
		ProtoMap:  schemaAddReq.ProtoMap,
	}, nil
}

// Add adds schema to an existing namespace.
func (h *SchemaHandler) Add(addReq *admin.NamespaceSchemaAddRequest, opts handler.ServiceOptions) (admin.NamespaceSchemaAddResponse, error) {
	var emptyRep = admin.NamespaceSchemaAddResponse{}

	kvOpts := kv.NewOverrideOptions().
		SetEnvironment(opts.ServiceEnvironment).
		SetZone(opts.ServiceZone)

	store, err := h.client.Store(kvOpts)
	if err != nil {
		return emptyRep, err
	}

	schemaAdmin := newAdminService(store, M3DBNodeNamespacesKey, nil)
	deployID, err := schemaAdmin.DeploySchema(addReq.Name, addReq.ProtoName, addReq.MsgName, addReq.ProtoMap)
	if err != nil {
		return emptyRep, err
	}
	return admin.NamespaceSchemaAddResponse{DeployID: deployID}, nil
}
