package main

import (
	"encoding/json"
	"go/format"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestGeneratedBudgetValidationAndNoNetworkDryRun(t *testing.T) {
	data, err := os.ReadFile("../../unified-openapi/openapi3.json")
	if err != nil {
		t.Fatal(err)
	}
	var s spec
	if err := json.Unmarshal(data, &s); err != nil {
		t.Fatal(err)
	}
	ops, drops := collectOps(&s, "Budget")
	if len(drops) > 0 {
		t.Fatal(drops)
	}
	var selected []operation
	for _, op := range ops {
		if op.Action == "create" {
			selected = append(selected, op)
		}
	}
	src, err := render("Budget", "validationfixture", "budget", selected)
	if err != nil {
		t.Fatal(err)
	}
	text := string(src)
	if !(strings.Index(text, "PrepareRequestBody") < strings.Index(text, "ConfirmPlan") && strings.Index(text, "ConfirmPlan") < strings.Index(text, "deps.APIClient()")) {
		t.Fatal("validation/plan must precede client creation")
	}
	dir, err := os.MkdirTemp("../..", ".gencli-validation-test-")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(dir) })
	for name, content := range map[string][]byte{"cmd.go": src, "cmd_test.go": []byte(generatedValidationTests)} {
		formatted, err := format.Source(content)
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, name), formatted, 0600); err != nil {
			t.Fatal(err)
		}
	}
	cmd := exec.Command("go", "test", "-count=1", ".")
	cmd.Dir = dir
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("fixture failed %v\n%s", err, out)
	}
}

const generatedValidationTests = `package validationfixture
import (
 "bytes"
 "context"
 "encoding/json"
 "io"
 "net/http"
 "os"
 "strings"
 "testing"
 clipkg "github.com/flexera-public/flexera-cli/internal/cli"
 "github.com/flexera-public/flexera-cli/internal/catalog"
)
type noHTTP struct{t *testing.T}
func(d noHTTP)Do(*http.Request)(*http.Response,error){d.t.Fatal("unexpected HTTP/client auth");return nil,nil}
func TestValidationOrder(t *testing.T){
 t.Setenv("HOME",t.TempDir());t.Setenv("FLEXERA_CLI_CONFIG","")
 index,err:=catalog.Load();if err!=nil{t.Fatal(err)};entry,_:=index.Lookup("Budget_Budget_create")
 for _,tc:=range []struct{name,body string;skip bool;want int}{
  {"valid-preview",string(entry.RequestExample),false,0},
  {"invalid","{}",false,2},
  {"skip","{}",true,0},
  {"syntax","{} {}",true,2},
  {"lost-field",` + "`" + `{"unknown":"private-input"}` + "`" + `,true,2},
 }{t.Run(tc.name,func(t *testing.T){var out,stderr bytes.Buffer;root,deps:=clipkg.NewRootCmd(clipkg.RootOptions{BaseHTTP:noHTTP{t},Getenv:os.Getenv});root.SetOut(&out);root.SetErr(&stderr);root.AddCommand(NewCmd());args:=[]string{"budget","create","--org-id","123","--body",tc.body,"--dry-run"};if tc.skip{args=append(args,"--no-validate")};if code:=clipkg.Execute(context.Background(),root,deps,args);code!=tc.want{t.Fatalf("exit %d want %d: %s",code,tc.want,stderr.String())};if tc.want==0{var plan map[string]any;if err:=json.Unmarshal(out.Bytes(),&plan);err!=nil{t.Fatal(err)};status:=plan["validation"].(map[string]any)["status"];want:="ok";if tc.skip{want="skipped"};if status!=want{t.Fatalf("status=%v",status)}}else if out.Len()!=0{t.Fatal("invalid input produced plan")};if strings.Contains(stderr.String(),"private-input"){t.Fatal("secret leaked")}})}
}
type captureHTTP struct{body []byte}
func(d *captureHTTP)Do(r *http.Request)(*http.Response,error){body,err:=io.ReadAll(r.Body);if err!=nil{return nil,err};d.body=body;return &http.Response{StatusCode:201,Header:http.Header{"Content-Type":[]string{"application/json"}},Body:io.NopCloser(strings.NewReader("{}")),Request:r},nil}
func TestPreviewBodyMatchesWire(t *testing.T){
 t.Setenv("HOME",t.TempDir())
 index,err:=catalog.Load();if err!=nil{t.Fatal(err)};entry,_:=index.Lookup("Budget_Budget_create")
 var preview bytes.Buffer
 root,deps:=clipkg.NewRootCmd(clipkg.RootOptions{BaseHTTP:noHTTP{t}});root.SetOut(&preview);root.SetErr(io.Discard);root.AddCommand(NewCmd())
 args:=[]string{"budget","create","--org-id","123","--body",string(entry.RequestExample)}
 if code:=clipkg.Execute(context.Background(),root,deps,append(append([]string{},args...),"--dry-run"));code!=0{t.Fatalf("preview exit=%d",code)}
 var plan struct{Plan struct{Body json.RawMessage}}
 if err:=json.Unmarshal(preview.Bytes(),&plan);err!=nil{t.Fatal(err)}
 capture:=&captureHTTP{}
 root,deps=clipkg.NewRootCmd(clipkg.RootOptions{BaseHTTP:capture});root.SetOut(io.Discard);root.SetErr(io.Discard);root.AddCommand(NewCmd())
 if code:=clipkg.Execute(context.Background(),root,deps,append(args,"--access-token","fixture-token"));code!=0{t.Fatalf("apply exit=%d",code)}
 var expected,actual any
 if err:=json.Unmarshal(plan.Plan.Body,&expected);err!=nil{t.Fatal(err)}
 if err:=json.Unmarshal(capture.body,&actual);err!=nil{t.Fatal(err)}
 expectedJSON,_:=json.Marshal(expected);actualJSON,_:=json.Marshal(actual)
 if string(expectedJSON)!=string(actualJSON){t.Fatalf("preview/wire differ: %s / %s",expectedJSON,actualJSON)}
}
`
