//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the
//  License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing,
//  software distributed under the License is distributed on an "AS
//  IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
//  express or implied. See the License for the specific language
//  governing permissions and limitations under the License.

package main

import (
	"net/http"
)

type GetLogsHandler struct {
	mr *MsgRing
}

func NewGetLogsHandler(mr *MsgRing) *GetLogsHandler {
	return &GetLogsHandler{mr: mr}
}

func (h *GetLogsHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	messages := h.mr.Messages()
	stringMessages := make([]string, len(messages))
	for i, message := range messages {
		stringMessages[i] = string(message)
	}

	rv := struct {
		Messages []string `json:"messages"`
	}{
		Messages: stringMessages,
	}
	mustEncode(w, rv)
}
