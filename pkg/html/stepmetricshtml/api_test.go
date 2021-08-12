package stepmetricshtml_test

import (
	"testing"

	"github.com/openshift/sippy/pkg/html/htmltesthelpers"
	"github.com/openshift/sippy/pkg/html/stepmetricshtml"
)

type apiTestCase struct {
	name    string
	request stepmetricshtml.Request
	// We don't care about the ordering of either MultistageDetails or
	// StepDetails, so we key by their respective names and iterate over the
	// result list when we run the test.
	expectedMultistageDetails map[string]stepmetricshtml.MultistageDetails
	expectedStepDetails       map[string]stepmetricshtml.StepDetails
}

func TestStepMetricsAPI(t *testing.T) {
	testCases := []apiTestCase{
		{
			name: "all multistage jobs",
			request: stepmetricshtml.Request{
				MultistageJobName: stepmetricshtml.All,
			},
			expectedMultistageDetails: map[string]stepmetricshtml.MultistageDetails{
				"e2e-aws": {
					Name: "e2e-aws",
					Trend: stepmetricshtml.Trend{
						Trajectory: stepmetricshtml.TrendTrajectoryFlat,
						Delta:      0,
					},
					StepDetails: map[string]stepmetricshtml.StepDetail{
						"aws-specific": stepmetricshtml.StepDetail{
							Name: "aws-specific",
							Trend: stepmetricshtml.Trend{
								Trajectory: stepmetricshtml.TrendTrajectoryFlat,
								Delta:      0,
							},
						},
						"ipi-install": stepmetricshtml.StepDetail{
							Name: "ipi-install",
							Trend: stepmetricshtml.Trend{
								Trajectory: stepmetricshtml.TrendTrajectoryFlat,
								Delta:      0,
							},
						},
						"openshift-e2e-test": stepmetricshtml.StepDetail{
							Name: "openshift-e2e-test",
							Trend: stepmetricshtml.Trend{
								Trajectory: stepmetricshtml.TrendTrajectoryFlat,
								Delta:      0,
							},
						},
					},
				},
				"e2e-gcp": {
					Name: "e2e-gcp",
					Trend: stepmetricshtml.Trend{
						Trajectory: stepmetricshtml.TrendTrajectoryFlat,
						Delta:      0,
					},
					StepDetails: map[string]stepmetricshtml.StepDetail{
						"gcp-specific": stepmetricshtml.StepDetail{
							Name: "gcp-specific",
							Trend: stepmetricshtml.Trend{
								Trajectory: stepmetricshtml.TrendTrajectoryFlat,
								Delta:      0,
							},
						},
						"ipi-install": stepmetricshtml.StepDetail{
							Name: "ipi-install",
							Trend: stepmetricshtml.Trend{
								Trajectory: stepmetricshtml.TrendTrajectoryFlat,
								Delta:      0,
							},
						},
						"openshift-e2e-test": stepmetricshtml.StepDetail{
							Name: "openshift-e2e-test",
							Trend: stepmetricshtml.Trend{
								Trajectory: stepmetricshtml.TrendTrajectoryFlat,
								Delta:      0,
							},
						},
					},
				},
			},
		},
		{
			name: "specific multistage job name",
			request: stepmetricshtml.Request{
				MultistageJobName: "e2e-aws",
			},
			expectedMultistageDetails: map[string]stepmetricshtml.MultistageDetails{
				"e2e-aws": {
					Name: "e2e-aws",
					Trend: stepmetricshtml.Trend{
						Trajectory: stepmetricshtml.TrendTrajectoryFlat,
						Delta:      0,
					},
					StepDetails: map[string]stepmetricshtml.StepDetail{
						"aws-specific": stepmetricshtml.StepDetail{
							Name: "aws-specific",
							Trend: stepmetricshtml.Trend{
								Trajectory: stepmetricshtml.TrendTrajectoryFlat,
								Delta:      0,
							},
						},
						"ipi-install": stepmetricshtml.StepDetail{
							Name: "ipi-install",
							Trend: stepmetricshtml.Trend{
								Trajectory: stepmetricshtml.TrendTrajectoryFlat,
								Delta:      0,
							},
						},
						"openshift-e2e-test": stepmetricshtml.StepDetail{
							Name: "openshift-e2e-test",
							Trend: stepmetricshtml.Trend{
								Trajectory: stepmetricshtml.TrendTrajectoryFlat,
								Delta:      0,
							},
						},
					},
				},
			},
		},
		{
			name: "all step names",
			request: stepmetricshtml.Request{
				StepName: stepmetricshtml.All,
			},
			expectedStepDetails: map[string]stepmetricshtml.StepDetails{
				"openshift-e2e-test": {
					Name: "openshift-e2e-test",
					Trend: stepmetricshtml.Trend{
						Trajectory: stepmetricshtml.TrendTrajectoryFlat,
						Delta:      0,
					},
					ByMultistage: map[string]stepmetricshtml.StepDetail{
						"e2e-aws": stepmetricshtml.StepDetail{
							Name: "openshift-e2e-test",
							Trend: stepmetricshtml.Trend{
								Trajectory: stepmetricshtml.TrendTrajectoryFlat,
								Delta:      0,
							},
						},
						"e2e-gcp": stepmetricshtml.StepDetail{
							Name: "openshift-e2e-test",
							Trend: stepmetricshtml.Trend{
								Trajectory: stepmetricshtml.TrendTrajectoryFlat,
								Delta:      0,
							},
						},
					},
				},
				"ipi-install": {
					Name: "ipi-install",
					Trend: stepmetricshtml.Trend{
						Trajectory: stepmetricshtml.TrendTrajectoryFlat,
						Delta:      0,
					},
					ByMultistage: map[string]stepmetricshtml.StepDetail{
						"e2e-aws": stepmetricshtml.StepDetail{
							Name: "ipi-install",
							Trend: stepmetricshtml.Trend{
								Trajectory: stepmetricshtml.TrendTrajectoryFlat,
								Delta:      0,
							},
						},
						"e2e-gcp": stepmetricshtml.StepDetail{
							Name: "ipi-install",
							Trend: stepmetricshtml.Trend{
								Trajectory: stepmetricshtml.TrendTrajectoryFlat,
								Delta:      0,
							},
						},
					},
				},
				"aws-specific": {
					Name: "aws-specific",
					Trend: stepmetricshtml.Trend{
						Trajectory: stepmetricshtml.TrendTrajectoryFlat,
						Delta:      0,
					},
					ByMultistage: map[string]stepmetricshtml.StepDetail{
						"e2e-aws": stepmetricshtml.StepDetail{
							Name: "aws-specific",
							Trend: stepmetricshtml.Trend{
								Trajectory: stepmetricshtml.TrendTrajectoryFlat,
								Delta:      0,
							},
						},
					},
				},
				"gcp-specific": {
					Name: "gcp-specific",
					Trend: stepmetricshtml.Trend{
						Trajectory: stepmetricshtml.TrendTrajectoryFlat,
						Delta:      0,
					},
					ByMultistage: map[string]stepmetricshtml.StepDetail{
						"e2e-gcp": stepmetricshtml.StepDetail{
							Name: "gcp-specific",
							Trend: stepmetricshtml.Trend{
								Trajectory: stepmetricshtml.TrendTrajectoryFlat,
								Delta:      0,
							},
						},
					},
				},
			},
		},
		{
			name: "specific step name",
			request: stepmetricshtml.Request{
				StepName: "openshift-e2e-test",
			},
			expectedStepDetails: map[string]stepmetricshtml.StepDetails{
				"openshift-e2e-test": {
					Name: "openshift-e2e-test",
					Trend: stepmetricshtml.Trend{
						Trajectory: stepmetricshtml.TrendTrajectoryFlat,
						Delta:      0,
					},
					ByMultistage: map[string]stepmetricshtml.StepDetail{
						"e2e-aws": stepmetricshtml.StepDetail{
							Name: "openshift-e2e-test",
							Trend: stepmetricshtml.Trend{
								Trajectory: stepmetricshtml.TrendTrajectoryFlat,
								Delta:      0,
							},
						},
						"e2e-gcp": stepmetricshtml.StepDetail{
							Name: "openshift-e2e-test",
							Trend: stepmetricshtml.Trend{
								Trajectory: stepmetricshtml.TrendTrajectoryFlat,
								Delta:      0,
							},
						},
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			a := stepmetricshtml.NewStepMetricsAPI(
				htmltesthelpers.GetTestReport("a-job-name", "test-name", "4.9"),
				htmltesthelpers.GetTestReport("a-job-name", "test-name", "4.9"),
			)

			resp := a.Fetch(testCase.request)

			if testCase.request.MultistageJobName != "" {
				assertAllMultistageDetails(t, resp.MultistageDetails, testCase.expectedMultistageDetails)
			}

			if testCase.request.StepName != "" {
				assertAllStepDetails(t, resp.StepDetails, testCase.expectedStepDetails)
			}
		})
	}
}

func assertAllMultistageDetails(t *testing.T, have []stepmetricshtml.MultistageDetails, want map[string]stepmetricshtml.MultistageDetails) {
	t.Helper()

	if len(have) != len(want) {
		t.Errorf("size mismatch, have: %d, want: %d", len(have), len(want))
	}

	for _, multistageDetails := range have {
		if _, ok := want[multistageDetails.Name]; !ok {
			t.Errorf("expected to find multistage details for: %s", multistageDetails.Name)
		}

		assertMultistageDetails(t, multistageDetails, want[multistageDetails.Name])
	}
}

func assertMultistageDetails(t *testing.T, have, want stepmetricshtml.MultistageDetails) {
	t.Helper()

	if have.Name != want.Name {
		t.Errorf("name mismatch, have: %s, want: %s", have.Name, want.Name)
	}

	assertTrend(t, have.Trend, want.Trend)

	for stageName := range want.StepDetails {
		if _, ok := have.StepDetails[stageName]; !ok {
			t.Errorf("missing step details for: %s", stageName)
		}

		assertStepDetail(t, have.StepDetails[stageName], want.StepDetails[stageName])
	}
}

func assertTrend(t *testing.T, have, want stepmetricshtml.Trend) {
	t.Helper()

	if have.Trajectory != want.Trajectory {
		t.Errorf("trajectory mismatch, have: %s, want: %s", have.Trajectory, want.Trajectory)
	}

	if have.Delta != want.Delta {
		t.Errorf("delta mismatch, have: %0.2f, want: %0.2f", have.Delta, want.Delta)
	}

	if have.Current.Name != have.Previous.Name {
		t.Errorf("trend name mismatch, current: %s, previous: %s", have.Current.Name, have.Previous.Name)
	}
}

func assertAllStepDetails(t *testing.T, have []stepmetricshtml.StepDetails, want map[string]stepmetricshtml.StepDetails) {
	t.Helper()

	if len(have) != len(want) {
		t.Errorf("size mismatch, have: %d, want: %d", len(have), len(want))
	}

	for _, stepDetails := range have {
		if _, ok := want[stepDetails.Name]; !ok {
			t.Errorf("expected to find step details for: %s", stepDetails.Name)
		}

		assertStepDetails(t, stepDetails, want[stepDetails.Name])
	}
}

func assertStepDetails(t *testing.T, have, want stepmetricshtml.StepDetails) {
	t.Helper()

	if have.Name != want.Name {
		t.Errorf("name mismatch, have: %s, want: %s", have.Name, want.Name)
	}

	assertTrend(t, have.Trend, want.Trend)

	for multistageName := range want.ByMultistage {
		if _, ok := have.ByMultistage[multistageName]; !ok {
			t.Errorf("missing step details for multistage name: %s", multistageName)
		}

		assertStepDetail(t, have.ByMultistage[multistageName], want.ByMultistage[multistageName])
	}
}

func assertStepDetail(t *testing.T, have, want stepmetricshtml.StepDetail) {
	if have.Name != want.Name {
		t.Errorf("name mismatch, have: %s, want: %s", have.Name, want.Name)
	}

	assertTrend(t, have.Trend, want.Trend)
}
