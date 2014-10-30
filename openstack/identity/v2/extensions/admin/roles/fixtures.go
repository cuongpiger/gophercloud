package roles

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/rackspace/gophercloud/testhelper"
	fake "github.com/rackspace/gophercloud/testhelper/client"
)

func MockListRoleResponse(t *testing.T) {
	th.Mux.HandleFunc("/OS-KSADMN/roles", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "roles": [
        {
            "id": "123",
            "name": "compute:admin",
            "description": "Nova Administrator"
        }
    ]
}
  `)
	})
}
