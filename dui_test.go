package documentoUnicoDeIdentidad

import "testing"

func TestAValidDocument(t *testing.T) {
	validDocument := "00016297-5"
	err := Validate(validDocument)
	if err == nil {
		t.Logf("It's a valid document.")
	} else {
		t.Fatalf("It's an invalid document.")		
	}
}

func TestAnotherValidDocument(t *testing.T) {
	validDocument := "02392210-8"
	err := Validate(validDocument)
	if err == nil {
		t.Logf("It's a valid document.")
	} else {
		t.Fatalf("It's an invalid document.")		
	}
}

func TestAnInvalidDocument(t *testing.T) {
	invalidDocument := "123456789-1"
	err := Validate(invalidDocument)
	if err != nil {
		t.Logf("It's an invalid document.")
	}
}


func TestAnInvalidDocumentWithInvalidFormat(t *testing.T) {
	invalidDocument := "1234567891"
	err := Validate(invalidDocument)
	if err != nil{
		t.Logf("It's an invalid document with invalid format.")
	} 
}