package http

import "net/http"

func NewRouter(h *Handlers) *http.ServeMux {
	mux := http.NewServeMux()

	apiPrefix := "/api/v1"

	routes := map[string]map[string]http.HandlerFunc{
		apiPrefix + "/storage": {
			http.MethodGet:  h.GetStorage,
			http.MethodPost: h.SetStorage,
		},
		apiPrefix + "/storage/sync": {
			http.MethodPost: h.SyncStorage,
		},
		apiPrefix + "/storage/consistency": {
			http.MethodGet: h.CheckConsistency,
		},
	}

	for path, methods := range routes {
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			if handler, ok := methods[r.Method]; ok {
				handler(w, r)
				return
			}

			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		})
	}

	return mux
}
