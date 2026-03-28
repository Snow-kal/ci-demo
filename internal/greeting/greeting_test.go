package greeting

import "testing"

func TestBuild(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     string
		wantErr   bool
		wantValue string
	}{
		{
			name:      "valid name",
			input:     "Go Dev",
			wantErr:   false,
			wantValue: "Hello Go Dev, your CI pipeline can build and test this app!",
		},
		{
			name:    "blank name",
			input:   "   ",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := Build(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if got != tc.wantValue {
				t.Fatalf("unexpected greeting, got %q, want %q", got, tc.wantValue)
			}
		})
	}
}
