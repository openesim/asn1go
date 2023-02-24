package asn1go

import (
	"testing"
)

func TestUnmarshalAsn1(t *testing.T) {
	var asn1Blob = []byte(`value7 ProfileElement ::= genericFileManagement : {
  gfm-header {
    mandated NULL,
    identification 21
  },
  fileManagementCMD {
    {
      filePath : ''H,
      createFCP : {
        fileDescriptor '4221007C'H,
        fileID '2FFB'H,
        lcsi '05'H,
        securityAttributesReferenced '2F060E'H,
        efFileSize '04D8'H,
        shortEFID ''H,
        proprietaryEFInfo {
          specialFileInformation '40'H
        }
      },
      filePath : '7F10'H,
      createFCP : {
        fileDescriptor '4621001A'H,
        fileID '6F44'H,
        lcsi '05'H,
        securityAttributesReferenced '2F0607'H,
        efFileSize '82'H,
        shortEFID ''H,
        proprietaryEFInfo {
          specialFileInformation '00'H
        }
      },
      filePath : '7F105F3A'H,
      createFCP : {
        fileDescriptor '42210002'H,
        fileID '4F09'H,
        lcsi '05'H,
        securityAttributesReferenced '2F0607'H,
        efFileSize '14'H,
        shortEFID '08'H,
        proprietaryEFInfo {
          specialFileInformation '00'H,
          repeatPattern '00'H
        }
      },
      createFCP : {
        fileDescriptor '42210011'H,
        fileID '4F11'H,
        lcsi '05'H,
        securityAttributesReferenced '2F0607'H,
        efFileSize 'AA'H,
        shortEFID '10'H,
        proprietaryEFInfo {
          specialFileInformation '00'H
        }
      },
      createFCP : {
        fileDescriptor '4221000D'H,
        fileID '4F12'H,
        lcsi '05'H,
        securityAttributesReferenced '2F0607'H,
        efFileSize '82'H,
        shortEFID '18'H,
        proprietaryEFInfo {
          specialFileInformation '40'H,
          fillPattern '00FF'H
        }
      },
      createFCP : {
        fileDescriptor '42210011'H,
        fileID '4F13'H,
        lcsi '05'H,
        securityAttributesReferenced '2F0607'H,
        efFileSize 'AA'H,
        shortEFID '38'H,
        proprietaryEFInfo {
          specialFileInformation '40'H
        }
      },
      createFCP : {
        fileDescriptor '42210028'H,
        fileID '4F14'H,
        lcsi '05'H,
        securityAttributesReferenced '2F0607'H,
        efFileSize '0190'H,
        shortEFID '40'H,
        proprietaryEFInfo {
          specialFileInformation '00'H
        }
      },
      createFCP : {
        fileDescriptor '42210003'H,
        fileID '4F15'H,
        lcsi '05'H,
        securityAttributesReferenced '2F0607'H,
        efFileSize '1E'H,
        shortEFID '28'H,
        proprietaryEFInfo {
          specialFileInformation '40'H
        }
      },
      createFCP : {
        fileDescriptor '42210002'H,
        fileID '4F16'H,
        lcsi '05'H,
        securityAttributesReferenced '2F0607'H,
        efFileSize '14'H,
        shortEFID '30'H,
        proprietaryEFInfo {
          specialFileInformation '40'H
        }
      },
      fillFileContent : '0001'H,
      fillFileContent : '0002'H,
      fillFileContent : '0000'H,
      fillFileContent : '0000'H,
      fillFileContent : '0000'H,
      fillFileContent : '0000'H,
      fillFileContent : '0000'H,
      fillFileContent : '0000'H,
      fillFileContent : '0000'H,
      fillFileContent : '0000'H,
      createFCP : {
        fileDescriptor '42210014'H,
        fileID '4F19'H,
        lcsi '05'H,
        securityAttributesReferenced '2F0607'H,
        efFileSize 'C8'H,
        shortEFID '20'H,
        proprietaryEFInfo {
          specialFileInformation '40'H
        }
      },
      createFCP : {
        fileDescriptor '4221001C'H,
        fileID '4F3A'H,
        lcsi '05'H,
        securityAttributesReferenced '2F0607'H,
        efFileSize '0118'H,
        shortEFID '50'H,
        proprietaryEFInfo {
          specialFileInformation '00'H
        }
      },
      fillFileContent : '546573746E722E31FFFFFFFFFFFF069194982143F1FFFFFFFFFFFFFF546573746E722E32FFFFFFFFFFFF069194982143F2'H,
      createFCP : {
        fileDescriptor '4221000F'H,
        fileID '4F3D'H,
        lcsi '05'H,
        securityAttributesReferenced '2F0607'H,
        efFileSize '96'H,
        shortEFID '60'H,
        proprietaryEFInfo {
          specialFileInformation '40'H
        }
      },
      createFCP : {
        fileDescriptor '4221000A'H,
        fileID '4F4B'H,
        lcsi '05'H,
        securityAttributesReferenced '2F0607'H,
        efFileSize '64'H,
        shortEFID ''H,
        proprietaryEFInfo {
          specialFileInformation '40'H
        }
      },
      createFCP : {
        fileDescriptor '4221000A'H,
        fileID '4F4C'H,
        lcsi '05'H,
        securityAttributesReferenced '2F0607'H,
        efFileSize '64'H,
        shortEFID '58'H,
        proprietaryEFInfo {
          specialFileInformation '40'H,
          repeatPattern '00'H
        }
      },
      createFCP : {
        fileDescriptor '42210014'H,
        fileID '4F4D'H,
        lcsi '05'H,
        securityAttributesReferenced '2F0607'H,
        efFileSize 'C8'H,
        shortEFID ''H,
        proprietaryEFInfo {
          specialFileInformation '40'H
        }
      },
      createFCP : {
        fileDescriptor '42210028'H,
        fileID '4F51'H,
        lcsi '05'H,
        securityAttributesReferenced '2F0607'H,
        efFileSize '0190'H,
        shortEFID '48'H,
        proprietaryEFInfo {
          specialFileInformation '40'H
        }
      },
      filePath : '7F10'H,
      createFCP : {
        fileDescriptor '7821'H,
        fileID '5F3E'H,
        lcsi '05'H,
        securityAttributesReferenced '2F0601'H,
        pinStatusTemplateDO '81010A0B'H
      },
      filePath : '7F105F3E'H,
      createFCP : {
        fileDescriptor '4121'H,
        fileID '4F01'H,
        lcsi '05'H,
        securityAttributesReferenced '2F060A'H,
        efFileSize '02'H,
        proprietaryEFInfo {
          specialFileInformation '40'H
        }
      },
      fillFileContent : '0100'H,
      createFCP : {
        fileDescriptor '7921'H,
        fileID '4F02'H,
        lcsi '05'H,
        securityAttributesReferenced '2F060A'H,
        efFileSize '0400'H,
        proprietaryEFInfo {
          specialFileInformation '40'H
        }
      },
      createFCP : {
        fileDescriptor '4121'H,
        fileID '4F03'H,
        lcsi '05'H,
        securityAttributesReferenced '2F060A'H,
        efFileSize '64'H,
        shortEFID ''H,
        proprietaryEFInfo {
          specialFileInformation '40'H
        }
      },
      createFCP : {
        fileDescriptor '4121'H,
        fileID '4F04'H,
        lcsi '05'H,
        securityAttributesReferenced '2F060A'H,
        efFileSize '64'H,
        shortEFID ''H,
        proprietaryEFInfo {
          specialFileInformation '40'H
        }
      },
      filePath : ''H,
      createFCP : {
        fileDescriptor '7821'H,
        fileID '7F66'H,
        lcsi '05'H,
        securityAttributesReferenced '2F0601'H,
        pinStatusTemplateDO '010A0B'H
      },
      filePath : '7F66'H,
      createFCP : {
        fileDescriptor '7821'H,
        fileID '5F40'H,
        lcsi '05'H,
        securityAttributesReferenced '2F0601'H,
        pinStatusTemplateDO '010A0B'H
      },
      filePath : '7F665F40'H,
      createFCP : {
        fileDescriptor '4121'H,
        fileID '4F40'H,
        lcsi '05'H,
        securityAttributesReferenced '2F0602'H,
        efFileSize '01'H,
        shortEFID ''H,
        proprietaryEFInfo {
          specialFileInformation '40'H
        }
      },
      fillFileContent : '00'H,
      createFCP : {
        fileDescriptor '4121'H,
        fileID '4F41'H,
        lcsi '05'H,
        securityAttributesReferenced '2F0602'H,
        efFileSize '20'H,
        shortEFID ''H,
        proprietaryEFInfo {
          specialFileInformation '40'H
        }
      },
      fillFileContent : '06013C1E3C1E0000000000000000000000000000000000000000000000000000'H,
      createFCP : {
        fileDescriptor '4121'H,
        fileID '4F42'H,
        lcsi '05'H,
        securityAttributesReferenced '2F0602'H,
        efFileSize '06'H,
        shortEFID ''H,
        proprietaryEFInfo {
          specialFileInformation '40'H,
          repeatPattern '00'H
        }
      },
      createFCP : {
        fileDescriptor '4121'H,
        fileID '4F43'H,
        lcsi '05'H,
        securityAttributesReferenced '2F0604'H,
        efFileSize '20'H,
        shortEFID ''H,
        proprietaryEFInfo {
          specialFileInformation '40'H,
          repeatPattern '00'H
        }
      },
      createFCP : {
        fileDescriptor '4121'H,
        fileID '4F44'H,
        lcsi '05'H,
        securityAttributesReferenced '2F0604'H,
        efFileSize '01'H,
        shortEFID ''H,
        proprietaryEFInfo {
          specialFileInformation '40'H
        }
      },
      fillFileContent : '00'H
    }
  }
}
value8 ProfileElement ::= usim : {
  usim-header {
    mandated NULL,
    identification 8
  }
}`)

	type ProfileElement struct {
		Header struct {
			MajorVersion int
			MinorVersion int
			ProfileType  string
			Iccid        string
		}
	}
	var profileElement ProfileElement

	err := Unmarshal(asn1Blob, &profileElement)
	if err != nil {
		t.Error("error:", err)
	}
	t.Logf("%+v", profileElement)

}
