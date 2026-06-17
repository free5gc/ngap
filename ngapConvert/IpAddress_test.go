package ngapConvert

import "testing"

func TestIPAddressToString(t *testing.T) {
	tests := []struct {
		name     string
		ipv4     string
		ipv6     string
		wantIPv4 string
		wantIPv6 string
	}{
		{
			name:     "ipv4",
			ipv4:     "192.0.2.1",
			wantIPv4: "192.0.2.1",
		},
		{
			name:     "ipv6",
			ipv6:     "2001:db8::1",
			wantIPv6: "2001:db8::1",
		},
		{
			name:     "ipv4 and ipv6",
			ipv4:     "192.0.2.1",
			ipv6:     "2001:db8::1",
			wantIPv4: "192.0.2.1",
			wantIPv6: "2001:db8::1",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ipv4, ipv6 := IPAddressToString(IPAddressToNgap(test.ipv4, test.ipv6))
			if ipv4 != test.wantIPv4 {
				t.Fatalf("unexpected ipv4: got %q, want %q", ipv4, test.wantIPv4)
			}
			if ipv6 != test.wantIPv6 {
				t.Fatalf("unexpected ipv6: got %q, want %q", ipv6, test.wantIPv6)
			}
		})
	}
}
