package cursor

import (
	"testing"

	"github.com/gofrs/uuid/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ValidUUID = uuid.Must(uuid.NewV7()).String()

func Test_newCursor(t *testing.T) {
	_ = assert.New(t)
	require := require.New(t)

	type args struct {
		lc string
		rc string
	}
	tests := []struct {
		name    string
		args    args
		want    Cursor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Successful",
			args:    args{lc: ValidUUID, rc: ValidUUID},
			want:    Cursor{Left: ValidUUID, Right: ValidUUID},
			wantErr: false,
		},
		{
			name:    "Missing both cursors",
			args:    args{lc: "", rc: ""},
			want:    Cursor{Left: "", Right: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newCursor(tt.args.lc, tt.args.rc)
			if tt.wantErr {
				require.Error(err)
				return
			}

			require.Equal(tt.want, got)
		})
	}
}

func TestNewLeftCursor(t *testing.T) {
	_ = assert.New(t)
	require := require.New(t)

	type args struct {
		lc string
	}
	tests := []struct {
		name    string
		args    args
		want    Cursor
		wantErr bool
	}{
		{
			name:    "Successful",
			args:    args{lc: ValidUUID},
			want:    Cursor{Left: ValidUUID, Right: ""},
			wantErr: false,
		},
		{
			name:    "Missing left cursor",
			args:    args{lc: ""},
			want:    Cursor{Left: "", Right: ""},
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLeft(tt.args.lc)
			if tt.wantErr {
				require.Error(err)
				return
			}

			require.Equal(tt.want, got)
		})
	}
}

func TestNewRightCursor(t *testing.T) {
	_ = assert.New(t)
	require := require.New(t)

	type args struct {
		rc string
	}
	tests := []struct {
		name    string
		args    args
		want    Cursor
		wantErr bool
	}{
		{
			name:    "Successful",
			args:    args{rc: ValidUUID},
			want:    Cursor{Left: "", Right: ValidUUID},
			wantErr: false,
		},
		{
			name:    "Missing right cursor",
			args:    args{rc: ""},
			want:    Cursor{Left: "", Right: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRight(tt.args.rc)
			if tt.wantErr {
				require.Error(err)
				return
			}

			require.Equal(tt.want, got)
		})
	}
}

func TestNewBoundedCursor(t *testing.T) {
	_ = assert.New(t)
	require := require.New(t)

	type args struct {
		lc string
		rc string
	}
	tests := []struct {
		name    string
		args    args
		want    Cursor
		wantErr bool
	}{
		{
			name:    "Successful",
			args:    args{lc: ValidUUID, rc: ValidUUID},
			want:    Cursor{Left: ValidUUID, Right: ValidUUID},
			wantErr: false,
		},
		{
			name:    "Missing left cursor",
			args:    args{lc: "", rc: ValidUUID},
			want:    Cursor{Left: "", Right: ValidUUID},
			wantErr: true,
		},
		{
			name:    "Missing right cursor",
			args:    args{lc: ValidUUID, rc: ""},
			want:    Cursor{Left: ValidUUID, Right: ""},
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBounded(tt.args.lc, tt.args.rc)
			if tt.wantErr {
				require.Error(err)
				return
			}

			require.Equal(tt.want, got)
		})
	}
}

func TestCursor_IsEmpty(t *testing.T) {
	_ = assert.New(t)
	require := require.New(t)

	type fields struct {
		Left  string
		Right string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Not empty",
			fields: fields{Left: ValidUUID, Right: ValidUUID},
			want:   false,
		},
		{
			name:   "Empty",
			fields: fields{Left: "", Right: ""},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cursor{
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			require.Equal(c.IsEmpty(), tt.want)
		})
	}
}

func TestCursor_HasLeft(t *testing.T) {
	_ = assert.New(t)
	require := require.New(t)

	type fields struct {
		Left  string
		Right string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Has left and doesn't have right",
			fields: fields{Left: ValidUUID, Right: ""},
			want:   true,
		},
		{
			name:   "Has left and has right",
			fields: fields{Left: ValidUUID, Right: ValidUUID},
			want:   true,
		},
		{
			name:   "Has neither left or right",
			fields: fields{Left: "", Right: ""},
			want:   false,
		},
		{
			name:   "Doesn't have left and has right",
			fields: fields{Left: "", Right: ValidUUID},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cursor{
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			require.Equal(c.HasLeft(), tt.want)
		})
	}
}

func TestCursor_HasRight(t *testing.T) {
	_ = assert.New(t)
	require := require.New(t)

	type fields struct {
		Left  string
		Right string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Has right and doesn't have left",
			fields: fields{Left: "", Right: ValidUUID},
			want:   true,
		},
		{
			name:   "Has right and has left",
			fields: fields{Left: ValidUUID, Right: ValidUUID},
			want:   true,
		},
		{
			name:   "Has neither left or right",
			fields: fields{Left: "", Right: ""},
			want:   false,
		},
		{
			name:   "Doesn't have right and has left",
			fields: fields{Left: ValidUUID, Right: ""},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cursor{
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			require.Equal(c.HasRight(), tt.want)
		})
	}
}

func TestCursor_IsLeft(t *testing.T) {
	_ = assert.New(t)
	require := require.New(t)

	type fields struct {
		Left  string
		Right string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Has left and doesn't have right",
			fields: fields{Left: ValidUUID, Right: ""},
			want:   true,
		},
		{
			name:   "Has left and has right",
			fields: fields{Left: ValidUUID, Right: ValidUUID},
			want:   false,
		},
		{
			name:   "Has neither left or right",
			fields: fields{Left: "", Right: ""},
			want:   false,
		},
		{
			name:   "Doesn't have left and has right",
			fields: fields{Left: "", Right: ValidUUID},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cursor{
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			require.Equal(c.IsLeft(), tt.want)
		})
	}
}

func TestCursor_IsRight(t *testing.T) {
	_ = assert.New(t)
	require := require.New(t)

	type fields struct {
		Left  string
		Right string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Has left and doesn't have right",
			fields: fields{Left: ValidUUID, Right: ""},
			want:   false,
		},
		{
			name:   "Has left and has right",
			fields: fields{Left: ValidUUID, Right: ValidUUID},
			want:   false,
		},
		{
			name:   "Has neither left or right",
			fields: fields{Left: "", Right: ""},
			want:   false,
		},
		{
			name:   "Doesn't have left and has right",
			fields: fields{Left: "", Right: ValidUUID},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cursor{
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			require.Equal(c.IsRight(), tt.want)
		})
	}
}

func TestCursor_IsBounded(t *testing.T) {
	_ = assert.New(t)
	require := require.New(t)

	type fields struct {
		Left  string
		Right string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Has left and doesn't have right",
			fields: fields{Left: ValidUUID, Right: ""},
			want:   false,
		},
		{
			name:   "Has both right and left",
			fields: fields{Left: ValidUUID, Right: ValidUUID},
			want:   true,
		},
		{
			name:   "Has neither left or right",
			fields: fields{Left: "", Right: ""},
			want:   false,
		},
		{
			name:   "Doesn't have left and has right",
			fields: fields{Left: "", Right: ValidUUID},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cursor{
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			require.Equal(c.IsBounded(), tt.want)
		})
	}
}
