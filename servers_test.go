package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp_GetServers(t *testing.T) {
	type args struct {
		page   string
		sort   string
		by     string
		filter string
	}
	tests := []struct {
		name        string
		args        args
		wantServers []ServerCore
		wantErr     bool
	}{
		{
			"v no sort",
			args{"1", "", "", ""},
			[]ServerCore{
				{"s3.example.com", "test server 3", 948, 1000, "Grand Larceny", "English", false},
				{"s4.example.com", "test server 4", 50, 50, "rivershell", "Polish", true},
				{"ss.southcla.ws", "Scavenge and Survive Official", 4, 32, "Scavenge & Survive by Southclaws", "English", false},
				{"s2.example.com", "test server 2", 0, 100, "Grand Larceny", "English", false},
			},
			false,
		},
		{
			"v desc",
			args{"1", "asc", "", ""},
			[]ServerCore{
				{"s2.example.com", "test server 2", 0, 100, "Grand Larceny", "English", false},
				{"ss.southcla.ws", "Scavenge and Survive Official", 4, 32, "Scavenge & Survive by Southclaws", "English", false},
				{"s4.example.com", "test server 4", 50, 50, "rivershell", "Polish", true},
				{"s3.example.com", "test server 3", 948, 1000, "Grand Larceny", "English", false},
			},
			false,
		},
		{
			"v pass",
			args{"1", "", "", "password"},
			[]ServerCore{
				{"s3.example.com", "test server 3", 948, 1000, "Grand Larceny", "English", false},
				{"ss.southcla.ws", "Scavenge and Survive Official", 4, 32, "Scavenge & Survive by Southclaws", "English", false},
				{"s2.example.com", "test server 2", 0, 100, "Grand Larceny", "English", false},
			},
			false,
		},
		{
			"v empty",
			args{"1", "", "", "empty"},
			[]ServerCore{
				{"s3.example.com", "test server 3", 948, 1000, "Grand Larceny", "English", false},
				{"s4.example.com", "test server 4", 50, 50, "rivershell", "Polish", true},
				{"ss.southcla.ws", "Scavenge and Survive Official", 4, 32, "Scavenge & Survive by Southclaws", "English", false},
			},
			false,
		},
		{
			"v full",
			args{"1", "", "", "full"},
			[]ServerCore{
				{"s3.example.com", "test server 3", 948, 1000, "Grand Larceny", "English", false},
				{"ss.southcla.ws", "Scavenge and Survive Official", 4, 32, "Scavenge & Survive by Southclaws", "English", false},
				{"s2.example.com", "test server 2", 0, 100, "Grand Larceny", "English", false},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotServers, err := app.GetServers(tt.args.page, tt.args.sort, tt.args.by, tt.args.filter)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantServers, gotServers)
		})
	}
}
