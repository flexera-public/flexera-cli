package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	cliconfig "github.com/flexera-public/flexera-cli/internal/config"
)

// EnvPrefix is the environment-variable namespace for all configuration.
// Combined with viper's key replacer ("-" -> "_") this preserves the legacy
// env var names exactly: flag "org-id" -> FLEXERA_CLI_ORG_ID, etc.
const EnvPrefix = "FLEXERA_CLI"

// Persistent flag / viper key names. These double as the suffixes (after
// the FLEXERA_CLI_ prefix and "-"->"_" replacement) of the supported env vars.
const (
	FlagConfig       = "config"
	FlagZone         = "zone"
	FlagAPIBaseURL   = "api-base-url"
	FlagLoginBaseURL = "login-base-url"
	FlagOutput       = "output"
	FlagJSONStyle    = "json-style"
	FlagOutJQ        = "out-jq"
	FlagOutFields    = "out-fields"
	FlagRawOutput    = "raw-output"
	FlagNoValidate   = "no-validate"
	FlagAccessToken  = "access-token"
	FlagClientID     = "client-id"
	FlagClientSecret = "client-secret"
	FlagRefreshToken = "refresh-token"
	FlagOrgID        = "org-id"
	FlagDebug        = "debug"
)

// persistentKeys are the config keys bound to viper (every persistent flag
// except --config and --debug, which are handled out of band).
var persistentKeys = []string{
	FlagZone, FlagAPIBaseURL, FlagLoginBaseURL, FlagOutput, FlagJSONStyle,
	FlagAccessToken, FlagClientID, FlagClientSecret, FlagRefreshToken, FlagOrgID,
}

// RegisterPersistentFlags declares the global flags on the root command's
// persistent flag set. Defaults are intentionally empty so that
// cliconfig.ResolveCommon remains the single owner of default values
// (keeping viper's flag-default at the lowest precedence rung).
func RegisterPersistentFlags(pf *pflag.FlagSet) {
	pf.String(FlagConfig, "", "config file (default $HOME/.flexera/config.yaml)")
	pf.String(FlagZone, "", "API zone (nam|eu|apac|test)")
	pf.String(FlagAPIBaseURL, "", "override API base URL")
	pf.String(FlagLoginBaseURL, "", "override login base URL")
	pf.StringP(FlagOutput, "o", "", "output format (json|table)")
	pf.String(FlagJSONStyle, "auto", "JSON whitespace style (auto|pretty|compact)")
	pf.String(FlagOutJQ, "", "shape JSON output with a jq expression")
	pf.String(FlagOutFields, "", "project JSON output fields (comma-separated paths)")
	pf.BoolP(FlagRawOutput, "r", false, "write jq string results without JSON quotes (requires --out-jq)")
	pf.Bool(FlagNoValidate, false, "skip API schema constraints (never JSON syntax or request data-loss checks)")
	pf.String(FlagAccessToken, "", "static bearer access token")
	pf.String(FlagClientID, "", "OAuth client ID")
	pf.String(FlagClientSecret, "", "OAuth client secret")
	pf.String(FlagRefreshToken, "", "OAuth refresh token")
	pf.Int(FlagOrgID, 0, "organization ID")
	pf.BoolP(FlagDebug, "d", false, "log HTTP requests/responses to stderr (Authorization redacted)")
}

// NewViper builds a viper instance wired for configuration precedence
// flag > env > config file > default. cfgFile, when non-empty, overrides the
// default ~/.flexera/config.yaml location. The persistent flag set is bound
// so changed flags win over env and file.
func NewViper(pf *pflag.FlagSet, cfgFile string) (*viper.Viper, error) {
	return newViper(pf, cfgFile, true)
}

// Discovery ignores implicit configuration while honoring explicit --config.
func newViper(pf *pflag.FlagSet, cfgFile string, implicit bool) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType("yaml")

	explicit := strings.TrimSpace(cfgFile) != ""
	if explicit {
		v.SetConfigFile(cfgFile)
	} else if implicit {
		if home, err := os.UserHomeDir(); err == nil {
			v.AddConfigPath(filepath.Join(home, ".flexera"))
			v.SetConfigName("config")
		}
	}
	if err := func() error {
		if explicit || implicit {
			return v.ReadInConfig()
		}
		return nil
	}(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if explicit {
			// An explicitly requested config file must exist and parse.
			return nil, fmt.Errorf("reading config file %q: %w", cfgFile, err)
		}
		if !asConfigNotFound(err, &notFound) {
			return nil, fmt.Errorf("reading config: %w", err)
		}
	}

	v.SetEnvPrefix(EnvPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	v.AutomaticEnv()

	for _, key := range persistentKeys {
		if f := pf.Lookup(key); f != nil {
			if err := v.BindPFlag(key, f); err != nil {
				return nil, fmt.Errorf("binding flag %q: %w", key, err)
			}
		}
	}
	return v, nil
}

func asConfigNotFound(err error, target *viper.ConfigFileNotFoundError) bool {
	if nf, ok := err.(viper.ConfigFileNotFoundError); ok {
		*target = nf
		return true
	}
	// os.PathError for a missing dir/file also counts as "no config".
	return os.IsNotExist(err)
}

// Resolve reads the merged configuration out of viper and applies the shared
// defaulting + validation in cliconfig.ResolveCommon. Because viper has
// already merged env and config-file values into each key, a no-op getenv is
// passed so ResolveCommon only supplies defaults.
func Resolve(v *viper.Viper) (cliconfig.CommonConfig, error) {
	if _, err := ResolveJSONStyle(v); err != nil {
		return cliconfig.CommonConfig{}, err
	}
	opts := cliconfig.CommonOptions{
		Zone:         v.GetString(FlagZone),
		APIBaseURL:   v.GetString(FlagAPIBaseURL),
		LoginBaseURL: v.GetString(FlagLoginBaseURL),
		Output:       v.GetString(FlagOutput),
		AccessToken:  v.GetString(FlagAccessToken),
		ClientID:     v.GetString(FlagClientID),
		ClientSecret: v.GetString(FlagClientSecret),
		RefreshToken: v.GetString(FlagRefreshToken),
		OrgID:        v.GetInt(FlagOrgID),
	}
	return cliconfig.ResolveCommon(opts, func(string) string { return "" })
}

// ResolveJSONStyle reads flag > env > file > default without changing the
// SDK's CommonConfig. Root initialization passes this style to Printer.
func ResolveJSONStyle(v *viper.Viper) (JSONStyle, error) {
	return ParseJSONStyle(v.GetString(FlagJSONStyle))
}
