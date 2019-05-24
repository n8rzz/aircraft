package aircraft

import "testing"

func assertFloat(t *testing.T, result, expected float64) {
	t.Helper()
	if result != expected {
		t.Errorf("result `%.2f` expected `%.2f", result, expected)
	}
}

func assertInt(t *testing.T, result, expected int) {
	t.Helper()
	if result != expected {
		t.Errorf("result `%d` expected `%d", result, expected)
	}
}

func assertString(t *testing.T, result, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("result `%s` expected `%s", result, expected)
	}
}

func assertState(t *testing.T, result, expected State) {
	t.Helper()
	if result != expected {
		t.Errorf("result `%d` expected `%d", result, expected)
	}
}

func TestAircraftUpdateAltitude(t *testing.T) {
	t.Run(".UpdateAltitude()", func(t *testing.T) {
		t.Run("when `#altitudeState` is `neutral`", func(t *testing.T) {
			t.Run("`#altitude` should not change", func(t *testing.T) {
				deltaTime := 1.0
				aircraft := Aircraft{
					altitude:             1000,
					altitudeAcceleration: 100,
					altitudeState:        neutral,
					altitudeTarget:       1000,
				}

				aircraft.UpdateAltitude(deltaTime)

				result := aircraft.altitude
				expected := 1000

				assertInt(t, result, expected)
			})

			t.Run("`#targetAltitude should reset to `#altitude`", func(t *testing.T) {
				deltaTime := 1.0
				aircraft := Aircraft{
					altitude:             1000,
					altitudeAcceleration: 100,
					altitudeState:        neutral,
					altitudeTarget:       2000,
				}

				aircraft.UpdateAltitude(deltaTime)

				result := aircraft.altitudeTarget
				expected := 1000

				assertInt(t, result, expected)
			})
		})

		t.Run("when `#altitudeState` is decreasing", func(t *testing.T) {
			t.Run("`#altitude` should decrement", func(t *testing.T) {
				deltaTime := 1.0
				aircraft := Aircraft{
					altitude:             1000,
					altitudeAcceleration: 100,
					altitudeState:        decreasing,
					altitudeTarget:       500,
				}

				aircraft.UpdateAltitude(deltaTime)

				result := aircraft.altitude
				expected := 900

				assertInt(t, result, expected)
			})

			t.Run("when `#altitude` < `#altitudeTarget`", func(t *testing.T) {
				t.Run("`#altitude` should set to `#altitudeTarget`", func(t *testing.T) {
					deltaTime := 1.0
					aircraft := Aircraft{
						altitude:             505,
						altitudeAcceleration: 100,
						altitudeState:        decreasing,
						altitudeTarget:       500,
					}

					aircraft.UpdateAltitude(deltaTime)

					result := aircraft.altitude
					expected := 500

					assertInt(t, result, expected)
				})

				t.Run("`#altitudeState` should set to neutral", func(t *testing.T) {
					deltaTime := 1.0
					aircraft := Aircraft{
						altitude:             505,
						altitudeAcceleration: 100,
						altitudeState:        decreasing,
						altitudeTarget:       500,
					}

					aircraft.UpdateAltitude(deltaTime)

					result := aircraft.altitudeState
					expected := neutral

					assertState(t, result, expected)
				})
			})
		})

		t.Run("when `#altitudeState` is increasing", func(t *testing.T) {
			t.Run("`#altitude` should increment", func(t *testing.T) {
				deltaTime := 1.0
				aircraft := Aircraft{
					altitude:             500,
					altitudeAcceleration: 100,
					altitudeState:        increasing,
					altitudeTarget:       1000,
				}

				aircraft.UpdateAltitude(deltaTime)

				result := aircraft.altitude
				expected := 600

				assertInt(t, result, expected)
			})

			t.Run("when `#altitude` < `#altitudeTarget`", func(t *testing.T) {
				t.Run("`#altitude` should set to `#altitudeTarget`", func(t *testing.T) {
					deltaTime := 1.0
					aircraft := Aircraft{
						altitude:             495,
						altitudeAcceleration: 100,
						altitudeState:        increasing,
						altitudeTarget:       500,
					}

					aircraft.UpdateAltitude(deltaTime)

					result := aircraft.altitude
					expected := 500

					assertInt(t, result, expected)
				})

				t.Run("`#altitudeState` should set to neutral", func(t *testing.T) {
					deltaTime := 1.0
					aircraft := Aircraft{
						altitude:             495,
						altitudeAcceleration: 100,
						altitudeState:        increasing,
						altitudeTarget:       500,
					}

					aircraft.UpdateAltitude(deltaTime)

					result := aircraft.altitudeState
					expected := neutral

					assertState(t, result, expected)
				})
			})
		})
	})
}

func TestAircraftUpdateHeading(t *testing.T) {
	t.Run(".UpdateHeading()", func(t *testing.T) {
		t.Run("when `#headingState` is `neutral`", func(t *testing.T) {
			t.Run("`#heading` should not change", func(t *testing.T) {
				deltaTime := 1.0
				aircraft := Aircraft{
					heading:             90,
					headingAcceleration: 5,
					headingState:        neutral,
					headingTarget:       270,
				}

				aircraft.UpdateHeading(deltaTime)

				result := aircraft.heading
				expected := 90

				assertInt(t, result, expected)
			})

			t.Run("`#targetHeading should reset to `#heading`", func(t *testing.T) {
				deltaTime := 1.0
				aircraft := Aircraft{
					heading:             90,
					headingAcceleration: 5,
					headingState:        neutral,
					headingTarget:       270,
				}

				aircraft.UpdateHeading(deltaTime)

				result := aircraft.headingTarget
				expected := 90

				assertInt(t, result, expected)
			})
		})

		t.Run("when `#headingState` is decreasing", func(t *testing.T) {
			t.Run("`#heading` should decrement", func(t *testing.T) {
				deltaTime := 1.0
				aircraft := Aircraft{
					heading:             90,
					headingAcceleration: 5,
					headingState:        decreasing,
					headingTarget:       10,
				}

				aircraft.UpdateHeading(deltaTime)

				result := aircraft.heading
				expected := 85

				assertInt(t, result, expected)
			})

			t.Run("`#heading` recalculates from 359 when decreasing through 0", func(t *testing.T) {
				deltaTime := 1.0
				aircraft := Aircraft{
					heading:             5,
					headingAcceleration: 10,
					headingState:        decreasing,
					headingTarget:       330,
				}

				aircraft.UpdateHeading(deltaTime)

				result := aircraft.heading
				expected := 355

				assertInt(t, result, expected)
			})

			t.Run("when `#heading` < `#headingTarget`", func(t *testing.T) {
				t.Run("`#heading` should set to `#headingTarget`", func(t *testing.T) {
					deltaTime := 1.0
					aircraft := Aircraft{
						heading:             95,
						headingAcceleration: 5,
						headingState:        decreasing,
						headingTarget:       100,
					}

					aircraft.UpdateHeading(deltaTime)

					result := aircraft.heading
					expected := 100

					assertInt(t, result, expected)
				})

				t.Run("`#headingState` should set to neutral", func(t *testing.T) {
					deltaTime := 1.0
					aircraft := Aircraft{
						heading:             95,
						headingAcceleration: 5,
						headingState:        decreasing,
						headingTarget:       100,
					}

					aircraft.UpdateHeading(deltaTime)

					result := aircraft.headingState
					expected := neutral

					assertState(t, result, expected)
				})
			})
		})

		t.Run("when `#headingState` is increasing", func(t *testing.T) {
			t.Run("`#heading` should increment", func(t *testing.T) {
				deltaTime := 1.0
				aircraft := Aircraft{
					heading:             90,
					headingAcceleration: 5,
					headingState:        increasing,
					headingTarget:       270,
				}

				aircraft.UpdateHeading(deltaTime)

				result := aircraft.heading
				expected := 95

				assertInt(t, result, expected)
			})

			t.Run("when `#heading` > `#headingTarget`", func(t *testing.T) {
				t.Run("`#heading` should set to `#headingTarget`", func(t *testing.T) {
					deltaTime := 1.0
					aircraft := Aircraft{
						heading:             90,
						headingAcceleration: 5,
						headingState:        increasing,
						headingTarget:       85,
					}

					aircraft.UpdateHeading(deltaTime)

					result := aircraft.heading
					expected := 85

					assertInt(t, result, expected)
				})

				t.Run("`#headingState` should set to neutral", func(t *testing.T) {
					deltaTime := 1.0
					aircraft := Aircraft{
						heading:             90,
						headingAcceleration: 5,
						headingState:        neutral,
						headingTarget:       270,
					}

					aircraft.UpdateHeading(deltaTime)

					result := aircraft.headingState
					expected := neutral

					assertState(t, result, expected)
				})

				t.Run("`#heading` recalculates from 0 when increasing through 359", func(t *testing.T) {
					deltaTime := 1.0
					aircraft := Aircraft{
						heading:             355,
						headingAcceleration: 10,
						headingState:        increasing,
						headingTarget:       10,
					}

					aircraft.UpdateHeading(deltaTime)

					result := aircraft.heading
					expected := 5

					assertInt(t, result, expected)
				})
			})
		})
	})
}

func TestAircraftUpdateSpeed(t *testing.T) {
	t.Run("when `#speedState is `0`", func(t *testing.T) {
		t.Run(".Update() should return early", func(t *testing.T) {
			deltaTimeMock := 1.0
			aircraft := Aircraft{
				speed:             0.0,
				speedAcceleration: 0.5,
				speedState:        neutral,
				speedTarget:       5.0,
			}

			aircraft.UpdateSpeed(deltaTimeMock)

			result := aircraft.speed
			expected := 0.0

			assertFloat(t, result, expected)
		})

		t.Run(".Update() should not change `#speedState`", func(t *testing.T) {
			deltaTimeMock := 1.0
			aircraft := Aircraft{
				speed:             0.0,
				speedAcceleration: 0.5,
				speedState:        neutral,
				speedTarget:       5.0,
			}

			aircraft.UpdateSpeed(deltaTimeMock)

			result := aircraft.speedState
			expected := neutral

			assertState(t, result, expected)
		})
	})

	t.Run("when `#speedState is `1`", func(t *testing.T) {
		t.Run(".Update() returns early when `#speed` == `#speedTarget`", func(t *testing.T) {
			deltaTimeMock := 1.0
			aircraft := Aircraft{
				speed:             5.0,
				speedAcceleration: 0.5,
				speedState:        increasing,
				speedTarget:       5.0,
			}

			aircraft.UpdateSpeed(deltaTimeMock)

			result := aircraft.speed
			expected := 5.0

			assertFloat(t, result, expected)
		})

		t.Run(".Update() increments `#speed` when `#speed` is < `#speedTarget`", func(t *testing.T) {
			deltaTimeMock := 1.0
			aircraft := Aircraft{
				speed:             0.0,
				speedAcceleration: 0.5,
				speedState:        increasing,
				speedTarget:       5.0,
			}

			aircraft.UpdateSpeed(deltaTimeMock)

			result := aircraft.speed
			expected := 0.5

			assertFloat(t, result, expected)
		})

		t.Run("when next value is > `#speedTarget`", func(t *testing.T) {
			t.Run("sets `#speed` to `#speedTarget`", func(t *testing.T) {
				deltaTimeMock := 1.0
				aircraft := Aircraft{
					speed:             4.9,
					speedAcceleration: 0.5,
					speedState:        increasing,
					speedTarget:       5.0,
				}

				aircraft.UpdateSpeed(deltaTimeMock)

				result := aircraft.speed
				expected := 5.0

				assertFloat(t, result, expected)
			})

			t.Run(".Update() sets `#speedState` to `neutral`", func(t *testing.T) {
				deltaTimeMock := 1.0
				aircraft := Aircraft{
					speed:             4.9,
					speedAcceleration: 0.5,
					speedState:        increasing,
					speedTarget:       5.0,
				}

				aircraft.UpdateSpeed(deltaTimeMock)

				result := aircraft.speedState
				expected := neutral

				assertState(t, result, expected)
			})
		})
	})

	t.Run("when `#speedState` is `neutral`", func(t *testing.T) {
		t.Run("decrements `#speed`", func(t *testing.T) {
			deltaTimeMock := 1.0
			aircraft := Aircraft{
				speed:             5.0,
				speedAcceleration: 0.5,
				speedState:        decreasing,
				speedTarget:       0.0,
			}

			aircraft.UpdateSpeed(deltaTimeMock)

			result := aircraft.speed
			expected := 4.5

			assertFloat(t, result, expected)
		})

		t.Run("when `#speed` > `#speedTarget`", func(t *testing.T) {
			t.Run("resets `#speedTarget`", func(t *testing.T) {
				deltaTimeMock := 1.0
				aircraft := Aircraft{
					speed:             0.1,
					speedAcceleration: 0.5,
					speedState:        decreasing,
					speedTarget:       0.0,
				}

				aircraft.UpdateSpeed(deltaTimeMock)

				result := aircraft.speed
				expected := 0.0

				assertFloat(t, result, expected)
			})

			t.Run("sets `#speedState` to `neutral", func(t *testing.T) {
				deltaTimeMock := 1.0
				aircraft := Aircraft{
					speed:             0.1,
					speedAcceleration: 0.5,
					speedState:        decreasing,
					speedTarget:       0.0,
				}

				aircraft.UpdateSpeed(deltaTimeMock)

				result := aircraft.speedState
				expected := neutral

				assertState(t, result, expected)
			})
		})
	})
}
