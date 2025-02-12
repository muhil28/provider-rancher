package bootstrap

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("rancher2_bootstrap", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "github"
        r.ShortGroup = "bootstrap"

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["token"].(string); ok {
				conn["token"] = []byte(a)
				conn["token_key"] = []byte(a)
			}
			if a, ok := attr["url"].(string); ok {
				conn["url"] = []byte(a)
				conn["api_url"] = []byte(a)
			}
			conn["insecure"] = []byte("true")
			return conn, nil
		}
    })
}
