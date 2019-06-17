package sconfig_test

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/psiemens/sconfig"
	"github.com/spf13/cobra"
)

func TestEnvironment(t *testing.T) {
	type Specification struct {
		String        string
		Bool          bool
		Int           int
		Int16         int16
		Int32         int32
		Int64         int64
		Uint          uint
		Uint8         uint8
		Uint16        uint16
		Uint32        uint32
		Uint64        uint64
		Float32       float32
		Float64       float64
		Duration      time.Duration
		StringSlice   []string
		BoolSlice     []bool
		IntSlice      []int
		DurationSlice []time.Duration
	}

	var s Specification

	os.Clearenv()

	os.Setenv("ENV_STRING", "apple")
	os.Setenv("ENV_BOOL", "true")
	os.Setenv("ENV_INT", "123")
	os.Setenv("ENV_INT16", "123")
	os.Setenv("ENV_INT32", "123")
	os.Setenv("ENV_INT64", "123")
	os.Setenv("ENV_UINT", "123")
	os.Setenv("ENV_UINT8", "123")
	os.Setenv("ENV_UINT16", "123")
	os.Setenv("ENV_UINT32", "123")
	os.Setenv("ENV_UINT64", "123")
	os.Setenv("ENV_FLOAT32", "123.123")
	os.Setenv("ENV_FLOAT64", "123.123")
	os.Setenv("ENV_DURATION", "5s")
	os.Setenv("ENV_STRINGSLICE", "apple,banana,orange")
	os.Setenv("ENV_BOOLSLICE", "true,false,true")
	os.Setenv("ENV_INTSLICE", "1,2,3")
	os.Setenv("ENV_DURATIONSLICE", "5s,10s,20m")

	defer os.Clearenv()

	err := sconfig.New(&s).
		FromEnvironment("ENV").
		Parse()
	if err != nil {
		t.Fail()
	}

	if s.String != "apple" {
		t.Errorf("expected %s, got %s", "apple", s.String)
	}

	if s.Bool != true {
		t.Errorf("expected %v, got %v", true, s.Bool)
	}

	if s.Int != 123 {
		t.Errorf("expected %d, got %d", 123, s.Int)
	}

	if s.Int16 != 123 {
		t.Errorf("expected %d, got %d", 123, s.Int16)
	}

	if s.Int32 != 123 {
		t.Errorf("expected %d, got %d", 123, s.Int32)
	}

	if s.Int64 != 123 {
		t.Errorf("expected %d, got %d", 123, s.Int64)
	}

	if s.Uint != 123 {
		t.Errorf("expected %d, got %d", 123, s.Uint)
	}

	if s.Uint8 != 123 {
		t.Errorf("expected %d, got %d", 123, s.Uint8)
	}

	if s.Uint16 != 123 {
		t.Errorf("expected %d, got %d", 123, s.Uint16)
	}

	if s.Uint32 != 123 {
		t.Errorf("expected %d, got %d", 123, s.Uint32)
	}

	if s.Uint64 != 123 {
		t.Errorf("expected %d, got %d", 123, s.Uint64)
	}

	if s.Float32 != 123.123 {
		t.Errorf("expected %f, got %f", 123.123, s.Float32)
	}

	if s.Float64 != 123.123 {
		t.Errorf("expected %f, got %f", 123.123, s.Float64)
	}

	if s.Duration != 5*time.Second {
		t.Errorf("expected %s, got %s", 5*time.Second, s.Duration)
	}

	if len(s.StringSlice) != 3 ||
		s.StringSlice[0] != "apple" ||
		s.StringSlice[1] != "banana" ||
		s.StringSlice[2] != "orange" {
		t.Errorf(
			"expected %#v, got %#v",
			[]string{"apple", "banana", "orange"},
			s.StringSlice,
		)
	}

	if len(s.BoolSlice) != 3 ||
		s.BoolSlice[0] != true ||
		s.BoolSlice[1] != false ||
		s.BoolSlice[2] != true {
		t.Errorf("expected %#v, got %#v", []bool{true, false, true}, s.BoolSlice)
	}

	if len(s.IntSlice) != 3 ||
		s.IntSlice[0] != 1 ||
		s.IntSlice[1] != 2 ||
		s.IntSlice[2] != 3 {
		t.Errorf("expected %#v, got %#v", []int{1, 2, 3}, s.IntSlice)
	}

	if len(s.DurationSlice) != 3 ||
		s.DurationSlice[0] != 5*time.Second ||
		s.DurationSlice[1] != 10*time.Second ||
		s.DurationSlice[2] != 20*time.Minute {
		t.Errorf(
			"expected %#v, got %#v",
			[]time.Duration{5 * time.Second, 10 * time.Second, 20 * time.Minute},
			s.DurationSlice,
		)
	}
}

func TestCommandLineFlags(t *testing.T) {
	type Specification struct {
		String        string          `flag:"string"`
		Bool          bool            `flag:"bool"`
		Int           int             `flag:"int"`
		Int16         int16           `flag:"int16"`
		Int32         int32           `flag:"int32"`
		Int64         int64           `flag:"int64"`
		Uint          uint            `flag:"uint"`
		Uint8         uint8           `flag:"uint8"`
		Uint16        uint16          `flag:"uint16"`
		Uint32        uint32          `flag:"uint32"`
		Uint64        uint64          `flag:"uint64"`
		Float32       float32         `flag:"float32"`
		Float64       float64         `flag:"float64"`
		Duration      time.Duration   `flag:"dur"`
		StringSlice   []string        `flag:"strings"`
		BoolSlice     []bool          `flag:"bools"`
		IntSlice      []int           `flag:"ints"`
		DurationSlice []time.Duration `flag:"durs"`
	}

	var s Specification

	c := &cobra.Command{
		Use:  "c",
		Args: cobra.ArbitraryArgs,
		Run:  func(_ *cobra.Command, _ []string) {},
	}

	err := sconfig.New(&s).
		BindFlags(c.PersistentFlags()).
		Parse()
	if err != nil {
		t.Fail()
	}

	err = executeCommand(
		c,
		"--string=apple",
		"--bool",
		"--int=123",
		"--int16=123",
		"--int32=123",
		"--int64=123",
		"--uint=123",
		"--uint8=123",
		"--uint16=123",
		"--uint32=123",
		"--uint64=123",
		"--float32=123.123",
		"--float64=123.123",
		"--dur=5s",
		"--strings=apple,banana,orange",
		"--bools=true,false,true",
		"--ints=1,2,3",
		"--durs=5s,10s,20m",
	)

	if s.String != "apple" {
		t.Errorf("expected %s, got %s", "apple", s.String)
	}

	if s.Bool != true {
		t.Errorf("expected %v, got %v", true, s.Bool)
	}

	if s.Int != 123 {
		t.Errorf("expected %d, got %d", 123, s.Int)
	}

	if s.Int16 != 123 {
		t.Errorf("expected %d, got %d", 123, s.Int16)
	}

	if s.Int32 != 123 {
		t.Errorf("expected %d, got %d", 123, s.Int32)
	}

	if s.Int64 != 123 {
		t.Errorf("expected %d, got %d", 123, s.Int64)
	}

	if s.Uint != 123 {
		t.Errorf("expected %d, got %d", 123, s.Uint)
	}

	if s.Uint8 != 123 {
		t.Errorf("expected %d, got %d", 123, s.Uint8)
	}

	if s.Uint16 != 123 {
		t.Errorf("expected %d, got %d", 123, s.Uint16)
	}

	if s.Uint32 != 123 {
		t.Errorf("expected %d, got %d", 123, s.Uint32)
	}

	if s.Uint64 != 123 {
		t.Errorf("expected %d, got %d", 123, s.Uint64)
	}

	if s.Float32 != 123.123 {
		t.Errorf("expected %f, got %f", 123.123, s.Float32)
	}

	if s.Float64 != 123.123 {
		t.Errorf("expected %f, got %f", 123.123, s.Float64)
	}

	if s.Duration != 5*time.Second {
		t.Errorf("expected %s, got %s", 5*time.Second, s.Duration)
	}

	if len(s.StringSlice) != 3 ||
		s.StringSlice[0] != "apple" ||
		s.StringSlice[1] != "banana" ||
		s.StringSlice[2] != "orange" {
		t.Errorf(
			"expected %#v, got %#v",
			[]string{"apple", "banana", "orange"},
			s.StringSlice,
		)
	}

	if len(s.BoolSlice) != 3 ||
		s.BoolSlice[0] != true ||
		s.BoolSlice[1] != false ||
		s.BoolSlice[2] != true {
		t.Errorf("expected %#v, got %#v", []bool{true, false, true}, s.BoolSlice)
	}

	if len(s.IntSlice) != 3 ||
		s.IntSlice[0] != 1 ||
		s.IntSlice[1] != 2 ||
		s.IntSlice[2] != 3 {
		t.Errorf("expected %#v, got %#v", []int{1, 2, 3}, s.IntSlice)
	}

	if len(s.DurationSlice) != 3 ||
		s.DurationSlice[0] != 5*time.Second ||
		s.DurationSlice[1] != 10*time.Second ||
		s.DurationSlice[2] != 20*time.Minute {
		t.Errorf(
			"expected %#v, got %#v",
			[]time.Duration{5 * time.Second, 10 * time.Second, 20 * time.Minute},
			s.DurationSlice,
		)
	}
}

func TestCommandLineShortFlags(t *testing.T) {
	type Specification struct {
		Environment string `flag:"env,e"`
		Host        string `flag:"host"`
		Port        int    `flag:"port,p"`
	}

	var s1 Specification

	c1 := &cobra.Command{
		Use:  "c",
		Args: cobra.ArbitraryArgs,
		Run:  func(_ *cobra.Command, _ []string) {},
	}

	err := sconfig.New(&s1).
		BindFlags(c1.PersistentFlags()).
		Parse()
	if err != nil {
		t.Fail()
	}

	// -h flag is invalid
	err = executeCommand(
		c1,
		"-e=PROD",
		"-p=80",
		"-h=127.0.0.1",
	)
	if err == nil {
		t.Errorf("expected error due to invalid -h flag")
	}

	var s2 Specification

	c2 := &cobra.Command{
		Use:  "c",
		Args: cobra.ArbitraryArgs,
		Run:  func(_ *cobra.Command, _ []string) {},
	}

	err = sconfig.New(&s2).
		BindFlags(c2.PersistentFlags()).
		Parse()
	if err != nil {
		t.Fail()
	}

	err = executeCommand(
		c2,
		"-e=PROD",
		"-p=80",
		"--host=127.0.0.1",
	)

	if s2.Environment != "PROD" {
		t.Errorf("expected %s, got %s", "PROD", s2.Environment)
	}

	if s2.Host != "127.0.0.1" {
		t.Errorf("expected %s, got %s", "127.0.0.1", s2.Host)
	}

	if s2.Port != 80 {
		t.Errorf("expected %d, got %d", 80, s2.Port)
	}
}

func TestEnvironmentAndCommandLineFlags(t *testing.T) {
	type Specification struct {
		Environment string `flag:"env,e"`
		Host        string `flag:"host"`
		Port        int    `flag:"port,p"`
	}

	var s Specification

	os.Clearenv()

	os.Setenv("ENV_ENVIRONMENT", "TEST")
	os.Setenv("ENV_HOST", "127.0.0.1")
	os.Setenv("ENV_PORT", "80")

	defer os.Clearenv()

	c := &cobra.Command{
		Use:  "c",
		Args: cobra.ArbitraryArgs,
		Run:  func(_ *cobra.Command, _ []string) {},
	}

	err := sconfig.New(&s).
		FromEnvironment("ENV").
		BindFlags(c.PersistentFlags()).
		Parse()
	if err != nil {
		t.Fail()
	}

	// override ENV_PORT env var with --port flag
	err = executeCommand(c, "--port=8080")

	if s.Environment != "TEST" {
		t.Errorf("expected %s, got %s", "TEST", s.Environment)
	}

	if s.Host != "127.0.0.1" {
		t.Errorf("expected %s, got %s", "127.0.0.1", s.Host)
	}

	if s.Port != 8080 {
		t.Errorf("expected %d, got %d", 80, s.Port)
	}
}

func TestDefaults(t *testing.T) {
	type Specification struct {
		Environment string `flag:"env,e"`
		Host        string `flag:"host" default:"127.0.0.1"`
		Port        int    `flag:"port,p"`
	}

	var s Specification

	os.Clearenv()

	os.Setenv("ENV_ENVIRONMENT", "TEST")
	os.Setenv("ENV_PORT", "80")

	defer os.Clearenv()

	c := &cobra.Command{
		Use:  "c",
		Args: cobra.ArbitraryArgs,
		Run:  func(_ *cobra.Command, _ []string) {},
	}

	err := sconfig.New(&s).
		FromEnvironment("ENV").
		BindFlags(c.PersistentFlags()).
		Parse()
	if err != nil {
		t.Fail()
	}

	// override env var with --port flag
	err = executeCommand(c, "--port=8080")

	if s.Environment != "TEST" {
		t.Errorf("expected %s, got %s", "TEST", s.Environment)
	}

	// host should fallback to default value
	if s.Host != "127.0.0.1" {
		t.Errorf("expected %s, got %s", "127.0.0.1", s.Host)
	}

	if s.Port != 8080 {
		t.Errorf("expected %d, got %d", 80, s.Port)
	}
}

func executeCommand(root *cobra.Command, args ...string) error {
	buf := new(bytes.Buffer)
	root.SetOutput(buf)
	root.SetArgs(args)

	_, err := root.ExecuteC()

	return err
}
