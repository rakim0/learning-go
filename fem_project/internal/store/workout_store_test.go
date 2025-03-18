package store

import (
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("pgx", "host=localhost port=5433 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		t.Fatalf("opening test db: %v", err)
	}
	// run the migrations for test db
	err = Migrate(db, "../../migrations")
	if err != nil {
		t.Fatalf("migrating db error: %v", err)
	}
	_, err = db.Exec(`TRUNCATE workouts, workout_entries CASCADE`)
	if err != nil {
		t.Fatalf("truncating tables %v", err)
	}
	return db
}

func TestCreateWorkout(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()
	store := NewPostgresWorkoutStore(db)

	tests := []struct {
		name    string
		workout *Workout
		wantErr bool
	}{
		{
			name: "valid workout",
			workout: &Workout{
				Title:           "push day",
				Description:     "upper body day",
				DurationMinutes: 60,
				CaloriesBurned:  200,
				Entries: []WorkoutEntry{
					{
						ExerciseName: "Bench Presses",
						Sets:         3,
						Reps:         IntPtr(10),
						Weight:       FloatPtr(135.5),
						Notes:        "warm up properly",
						OrderIndex:   1,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "workout with invalid workout entries",
			workout: &Workout{
				Title:           "fully body",
				Description:     "complete workout",
				DurationMinutes: 90,
				CaloriesBurned:  500,
				Entries: []WorkoutEntry{
					{
						ExerciseName: "Plank",
						Sets:         3,
						Reps:         IntPtr(10),
						Weight:       FloatPtr(135.5),
						OrderIndex:   1,
					},
					{
						ExerciseName:    "Squats",
						Sets:            4,
						Reps:            IntPtr(10),
						Weight:          FloatPtr(135.5),
						DurationSeconds: IntPtr(20),
						OrderIndex:      2,
					},
				},
			},
			wantErr: true,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			createdWorkout, err := store.CreateWorkout(testcase.workout)
			if testcase.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, testcase.workout.Title, createdWorkout.Title)
			assert.Equal(t, testcase.workout.Description, createdWorkout.Description)
			assert.Equal(t, testcase.workout.DurationMinutes, createdWorkout.DurationMinutes)
			assert.Equal(t, testcase.workout.CaloriesBurned, createdWorkout.CaloriesBurned)
			assert.Equal(t, testcase.workout.Entries, createdWorkout.Entries)

			retrieved, err := store.GetWorkoutByID(int64(createdWorkout.ID))
			require.NoError(t, err)
			assert.Equal(t, createdWorkout.ID, retrieved.ID)
			assert.Equal(t, len(testcase.workout.Entries), len(retrieved.Entries))
			for i := range retrieved.Entries {
				assert.Equal(t, testcase.workout.Entries[i].ExerciseName, retrieved.Entries[i].ExerciseName)
				assert.Equal(t, testcase.workout.Entries[i].Sets, retrieved.Entries[i].Sets)
				assert.Equal(t, testcase.workout.Entries[i].OrderIndex, retrieved.Entries[i].OrderIndex)
			}
		})
	}
}

func IntPtr(i int) *int {
	return &i
}
func FloatPtr(f float64) *float64 {
	return &f
}
