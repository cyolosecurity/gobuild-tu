package analyze

import "testing"

// a sample of a list with a single entry in the expected JSON format
var sample = "[\n\t{\n\t\t\"ID\": 0,\n\t\t\"Mode\": \"link-install\",\n\t\t\"Package\": \"command-line-arguments\",\n\t\t\"Deps\": [\n\t\t\t1\n\t\t],\n\t\t\"Objdir\": \"/var/folders/s_/_blhkrnx47793cqf6rbzhx2h0000gn/T/go-build498529669/b001/\",\n\t\t\"Target\": \"main\",\n\t\t\"Priority\": 678,\n\t\t\"Built\": \"main\",\n\t\t\"BuildID\": \"S5-nwDHJ69frbrI9UIvZ/6mba6NiX_hEvMvgo9ojQ/47TBGsuRSZ_go12wX3xn/ea8rIMilgyd15Q9rW_GV\",\n\t\t\"TimeReady\": \"2022-02-20T16:48:52.768918+02:00\",\n\t\t\"TimeStart\": \"2022-02-20T16:48:52.768934+02:00\",\n\t\t\"TimeDone\": \"2022-02-20T16:48:52.769853+02:00\",\n\t\t\"Cmd\": null\n\t}]"

func TestAnalyzeJSON(t *testing.T) {
	JSON([]byte(sample))
}
