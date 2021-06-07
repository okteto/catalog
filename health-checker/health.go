package health

// SimpleHealthClient represents a client that obtains health service
// in real time and has no historical health data available to it.
type SimpleHealthClient struct{}

func (s *SimpleHealthClient) Get() []ServiceHealth {
	return []ServiceHealth{
		{
			ServiceID: "3d8201b4-d152-4d84-b8bc-eb87817a617e",
			HealthResults: []HealthResult{
				{
					Healthy:   true,
					Timestamp: "1621778498",
				},
			},
		},
		{
			ServiceID: "a41a7cb2-6ab1-489d-9dab-ff008e4ce34e",
			HealthResults: []HealthResult{
				{
					Healthy:   false,
					Timestamp: "1621778553",
				},
			},
		},
		{
			ServiceID: "028d4690-f9b8-4cb2-9edb-a44d831dbed3",
			HealthResults: []HealthResult{
				{
					Healthy:   true,
					Timestamp: "1621778592",
				},
			},
		},
		{
			ServiceID: "d2647c9b-2e15-4978-a97f-2bcad9126f57",
			HealthResults: []HealthResult{
				{
					Healthy:   true,
					Timestamp: "1621778628",
				},
			},
		},
	}
}

// AdvancedHealthClient represents a client that obtains historical service
// health from a persistent store.
type AdvancedHealthClient struct{}

func (a *AdvancedHealthClient) Get() []ServiceHealth {
	return []ServiceHealth{
		{
			ServiceID: "3d8201b4-d152-4d84-b8bc-eb87817a617e",
			HealthResults: []HealthResult{
				{
					Healthy:   true,
					Timestamp: "1621778498",
				},
				{
					Healthy:   true,
					Timestamp: "1621778488",
				},
				{
					Healthy:   false,
					Timestamp: "1621778478",
				},
				{
					Healthy:   true,
					Timestamp: "1621778468",
				},
			},
		},
		{
			ServiceID: "a41a7cb2-6ab1-489d-9dab-ff008e4ce34e",
			HealthResults: []HealthResult{
				{
					Healthy:   false,
					Timestamp: "1621778553",
				},
				{
					Healthy:   false,
					Timestamp: "1621778543",
				},
				{
					Healthy:   false,
					Timestamp: "1621778533",
				},
				{
					Healthy:   false,
					Timestamp: "1621778523",
				},
			},
		},
		{
			ServiceID: "028d4690-f9b8-4cb2-9edb-a44d831dbed3",
			HealthResults: []HealthResult{
				{
					Healthy:   true,
					Timestamp: "1621778592",
				},
				{
					Healthy:   true,
					Timestamp: "1621778582",
				},
				{
					Healthy:   true,
					Timestamp: "1621778572",
				},
				{
					Healthy:   true,
					Timestamp: "1621778562",
				},
			},
		},
		{
			ServiceID: "d2647c9b-2e15-4978-a97f-2bcad9126f57",
			HealthResults: []HealthResult{
				{
					Healthy:   false,
					Timestamp: "1621778628",
				},
				{
					Healthy:   false,
					Timestamp: "1621778618",
				},
				{
					Healthy:   true,
					Timestamp: "1621778608",
				},
				{
					Healthy:   true,
					Timestamp: "1621778598",
				},
			},
		},
	}
}
