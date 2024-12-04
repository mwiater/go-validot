package plugins

import (
	"fmt"
	"net"
	"strings"
)

// IPAddressValidationPlugin validates that the value of a specific environment variable
// key is a valid IP address. It can enforce constraints on IP versions (e.g., IPv4 or IPv6)
// and ensure that the IP address is private.
type IPAddressValidationPlugin struct {
	Key               string   // The key of the environment variable to validate.
	AllowedIPVersions []string // A list of allowed IP versions, e.g., "IPv4", "IPv6".
	MustBePrivate     bool     // If true, enforces that the IP address must be private.
}

// Validate checks if the value associated with the given key is a valid IP address
// and meets the specified criteria, such as allowed versions and private IP enforcement.
//
// Parameters:
//   - key: The key of the environment variable being validated.
//   - value: The value of the environment variable to validate.
//
// Returns:
//   - bool: Indicates whether this plugin handled the validation.
//   - error: An error if the value is invalid or nil if it passes validation.
func (p *IPAddressValidationPlugin) Validate(key, value string) (bool, error) {
	if key != p.Key {
		return false, nil // Plugin does not handle this key.
	}

	ip := net.ParseIP(strings.TrimSpace(value))
	if ip == nil {
		return true, fmt.Errorf("value for key %q must be a valid IP address", key)
	}

	if len(p.AllowedIPVersions) > 0 {
		validVersion := false
		for _, version := range p.AllowedIPVersions {
			switch strings.ToLower(version) {
			case "ipv4":
				if ip.To4() != nil {
					validVersion = true
				}
			case "ipv6":
				if ip.To16() != nil && ip.To4() == nil {
					validVersion = true
				}
			}
			if validVersion {
				break
			}
		}
		if !validVersion {
			return true, fmt.Errorf("value for key %q must be one of the following IP versions: %v", key, p.AllowedIPVersions)
		}
	}

	if p.MustBePrivate {
		if !isPrivateIP(ip) {
			return true, fmt.Errorf("value for key %q must be a private IP address", key)
		}
	}

	return true, nil
}

// isPrivateIP determines if the given IP address belongs to a private IP range.
//
// Parameters:
//   - ip: The IP address to check.
//
// Returns:
//   - bool: True if the IP address is private, otherwise false.
func isPrivateIP(ip net.IP) bool {
	privateIPBlocks := []*net.IPNet{}

	// Define private IP ranges.
	for _, cidr := range []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"::1/128",
		"fc00::/7",
	} {
		_, block, _ := net.ParseCIDR(cidr)
		privateIPBlocks = append(privateIPBlocks, block)
	}

	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}

// Name returns the name of the plugin.
//
// Returns:
//   - string: The name of the plugin.
func (p *IPAddressValidationPlugin) Name() string {
	return "IPAddressValidationPlugin"
}
