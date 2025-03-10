package dremio_test

import (
	"context"
	"encoding/json"
	"github.com/mskcc/smile/internal/dremio"
	"github.com/mskcc/smile/internal/smile"
	"testing"
)

var newRequest = `
{
  "smileRequestId": "6cca6166-875a-11eb-ae9e-acde48001122",
  "igoRequestId": "22022_BZ",
  "genePanel": "GENESET101_BAITS",
  "projectManagerName": "bart simpson",
  "piEmail": "",
  "labHeadName": "beavis",
  "labHeadEmail": "beavis@mskcc.org",
  "investigatorName": "butthead",
  "investigatorEmail": "butthead@mskcc.org",
  "dataAnalystName": "",
  "dataAnalystEmail": "",
  "otherContactEmails": "metallica@mskcc.org",
  "dataAccessEmails": "",
  "qcAccessEmails": "",
  "isCmoRequest": true,
  "bicAnalysis": false,
  "samples": [
    {
      "smileSampleId": "afe74fba-8756-11eb-9b45-acde48001122",
      "smilePatientId": "6cc7394f-875a-11eb-91ec-acde48001122",
      "cmoSampleName": "C-TX6DNG-N001-d",
      "sampleName": "LMNO_4396_N",
      "sampleType": "Normal",
      "oncotreeCode": "TPLL",
      "collectionYear": "",
      "tubeId": "4157451784",
      "cfDNA2dBarcode": "8029250670",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "",
          "investigatorDecision": "Continue processing"
        }
      ],
      "libraries": [
        {
          "libraryIgoId": "22022_CC_3_1",
          "libraryConcentrationNgul": 34.2,
          "captureConcentrationNm": "1.461988304093567",
          "captureInputNg": "50.0",
          "captureName": "Pool-22022_BZ-22022_CC-Tube7_1",
          "runs": [
            {
              "runMode": "HiSeq High Output",
              "runId": "CRX_7395",
              "flowCellId": "HGJMLBBXY",
              "readLength": "101/8/8/101",
              "runDate": "2020-05-20",
              "flowCellLanes": [
                1,
                2,
                3,
                4,
                5,
                6,
                7
              ],
              "fastqs": [
                "/FASTQ/Project_22022_CC/Sample_LMNO_4396_N_IGO_22022_CC_3/LMNO_4396_N_IGO_22022_CC_3_S144_R1_001.fastq.gz",
                "/FASTQ/Project_22022_CC/Sample_LMNO_4396_N_IGO_22022_CC_3/LMNO_4396_N_IGO_22022_CC_3_S144_R2_001.fastq.gz"
              ]
            }
          ]
        }
      ],
      "cmoPatientId": "C-TX6DNG",
      "primaryId": "22022_CC_3",
      "investigatorSampleId": "LMNO_4396_N",
      "species": "Human",
      "sex": "F",
      "tumorOrNormal": "Normal",
      "preservation": "EDTA-Streck",
      "sampleClass": "Blood",
      "sampleOrigin": "Buffy Coat",
      "tissueLocation": "Blood",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "Buffy Coat",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-TX6DNG"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_CC_3"
        },
        {
          "namespace": "investigatorId",
          "value": "LMNO_4396_N"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
    },
    {
      "smileSampleId": "afe74d1c-8756-11eb-a95a-acde48001122",
      "smilePatientId": "6cca3573-875a-11eb-911e-acde48001122",
      "cmoSampleName": "C-DPCXX1-N001-d",
      "sampleName": "LMNO_3443_N",
      "sampleType": "Normal",
      "oncotreeCode": "MPNWP",
      "collectionYear": "2018",
      "tubeId": "",
      "cfDNA2dBarcode": "",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "",
          "investigatorDecision": "Continue processing"
        }
      ],
      "libraries": [
        {
          "barcodeId": "DUAL_TNG_LIB_XXX",
          "barcodeIndex": "TGCATGCT-GGCTATAA",
          "libraryIgoId": "22022_BZ_19_1_1_1",
          "libraryVolume": 35.0,
          "libraryConcentrationNgul": 26.1,
          "captureConcentrationNm": "1.9157088122605364",
          "captureInputNg": "50.0",
          "captureName": "Pool-22022_BZ-22022_CC-Tube7_1",
          "runs": [
            {
              "runMode": "HiSeq High Output",
              "runId": "CRX_7395",
              "flowCellId": "HXXXLBBYY",
              "readLength": "101/8/8/101",
              "runDate": "2020-05-20",
              "flowCellLanes": [
                1,
                2,
                3,
                4,
                5,
                6,
                7
              ],
              "fastqs": [
                "/FASTQ/Project_22022_BZ/Sample_LMNO_3443_N_IGO_22022_BZ_19/LMNO_3443_N_IGO_22022_BZ_19_S132_R1_001.fastq.gz",
                "/FASTQ/Project_22022_BZ/Sample_LMNO_3443_N_IGO_22022_BZ_19/LMNO_3443_N_IGO_22022_BZ_19_S132_R2_001.fastq.gz"
              ]
            }
          ]
        }
      ],
      "cmoPatientId": "C-DPCXX1",
      "primaryId": "22022_BZ_19",
      "investigatorSampleId": "LMNO_3443_N",
      "species": "Human",
      "sex": "M",
      "tumorOrNormal": "Normal",
      "preservation": "EDTA-Streck",
      "sampleClass": "other",
      "sampleOrigin": "Buffy Coat",
      "tissueLocation": "",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "Buffy Coat",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-DPCXX1"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_BZ_19"
        },
        {
          "namespace": "investigatorId",
          "value": "LMNO_3443_N"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
    },
    {
      "smileSampleId": "afe74c2e-8756-11eb-b556-acde48001122",
      "smilePatientId": "6cca38e3-875a-11eb-9e48-acde48001122",
      "cmoSampleName": "C-KXXL3J-P001-d",
      "sampleName": "LMNO_0034_T_DNA",
      "sampleType": "Primary",
      "oncotreeCode": "SARCNOS",
      "collectionYear": "2018",
      "tubeId": "TS-189768.2.T",
      "cfDNA2dBarcode": "8029250670",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "",
          "investigatorDecision": ""
        }
      ],
      "libraries": [
        {
          "barcodeId": "DUAL_IDT_LIB_661",
          "barcodeIndex": "GAATTACG-CAAGGAGA",
          "libraryIgoId": "22022_BZ_15_1_1_1",
          "libraryVolume": 35.0,
          "libraryConcentrationNgul": 27.0,
          "captureConcentrationNm": "3.7037037037037037",
          "captureInputNg": "100.0",
          "captureName": "Pool-22022_BZ-22022_CC-Tube7_1",
          "runs": [
            {
              "runMode": "HiSeq High Output",
              "runId": "CRX_7395",
              "flowCellId": "HXXXLBBYY",
              "readLength": "101/8/8/101",
              "runDate": "2020-05-20",
              "flowCellLanes": [
                1,
                2,
                3,
                4,
                5,
                6,
                7
              ],
              "fastqs": [
                "/FASTQ/Project_22022_BZ/Sample_LMNO_0034_T_DNA_IGO_22022_BZ_15/LMNO_0034_T_DNA_IGO_22022_BZ_15_S128_R2_001.fastq.gz",
                "/FASTQ/Project_22022_BZ/Sample_LMNO_0034_T_DNA_IGO_22022_BZ_15/LMNO_0034_T_DNA_IGO_22022_BZ_15_S128_R1_001.fastq.gz"
              ]
            }
          ]
        }
      ],
      "cmoPatientId": "C-KXXL3J",
      "primaryId": "22022_BZ_15",
      "investigatorSampleId": "LMNO_0034_T_DNA",
      "species": "Human",
      "sex": "M",
      "tumorOrNormal": "Tumor",
      "preservation": "OCT",
      "sampleClass": "Resection",
      "sampleOrigin": "Tissue",
      "tissueLocation": "Left Kidney",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "Tissue",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-KXXL3J"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_BZ_15"
        },
        {
          "namespace": "investigatorId",
          "value": "LMNO_0034_T_DNA"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
    },
    {
      "smileSampleId": "afe74ca6-8756-11eb-ae00-acde48001122",
      "smilePatientId": "6cca3bdc-875a-11eb-a412-acde48001122",
      "cmoSampleName": "C-XXX711-M001-d",
      "sampleName": "LMNO_9999_T",
      "sampleType": "Metastasis",
      "oncotreeCode": "PAAD",
      "collectionYear": "2017",
      "tubeId": "181446.13.T",
      "cfDNA2dBarcode": "8029250670",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "",
          "investigatorDecision": "Continue processing"
        }
      ],
      "libraries": [
        {
          "barcodeId": "DUAL_IDT_LIB_663",
          "barcodeIndex": "TACTAACG-CAAGTAAT",
          "libraryIgoId": "22022_BZ_17_1_1_1",
          "libraryVolume": 35.0,
          "libraryConcentrationNgul": 28.7,
          "captureConcentrationNm": "3.4843205574912894",
          "captureInputNg": "100.0",
          "captureName": "Pool-22022_BZ-22022_CC-Tube7_1",
          "runs": [
            {
              "runMode": "HiSeq High Output",
              "runId": "CRX_7395",
              "flowCellId": "HXXXLBBYY",
              "readLength": "101/8/8/101",
              "runDate": "2020-05-20",
              "flowCellLanes": [
                1,
                2,
                3,
                4,
                5,
                6,
                7
              ],
              "fastqs": [
                "/FASTQ/Project_22022_BZ/Sample_LMNO_9999_T_IGO_22022_BZ_17/LMNO_9999_T_IGO_22022_BZ_17_S130_R1_001.fastq.gz",
                "/FASTQ/Project_22022_BZ/Sample_LMNO_9999_T_IGO_22022_BZ_17/LMNO_9999_T_IGO_22022_BZ_17_S130_R2_001.fastq.gz"
              ]
            }
          ]
        }
      ],
      "cmoPatientId": "C-XXX711",
      "primaryId": "22022_BZ_17",
      "investigatorSampleId": "LMNO_9999_T",
      "species": "Human",
      "sex": "M",
      "tumorOrNormal": "",
      "preservation": "OCT",
      "sampleClass": "Resection",
      "sampleOrigin": "",
      "tissueLocation": "",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-XXX711"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_BZ_17"
        },
        {
          "namespace": "investigatorId",
          "value": "LMNO_9999_T"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
    },
    {
      "smileSampleId": "afe74e51-8756-11eb-9a0e-acde48001122",
      "smilePatientId": "6cca3d19-875a-11eb-92e8-acde48001122",
      "cmoSampleName": "C-PPPXX2-N001-d",
      "sampleName": "LMNO_8080_N",
      "sampleType": "Normal",
      "oncotreeCode": "COAD",
      "collectionYear": "",
      "tubeId": "8027185112",
      "cfDNA2dBarcode": "8029250670",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "Low quantity",
          "investigatorDecision": "Continue processing"
        }
      ],
      "libraries": [],
      "cmoPatientId": "C-PPPXX2",
      "primaryId": "22022_BZ_4",
      "investigatorSampleId": "LMNO_8080_N",
      "species": "Human",
      "sex": "F",
      "tumorOrNormal": "Normal",
      "preservation": "EDTA-Streck",
      "sampleClass": "Blood",
      "sampleOrigin": "Buffy Coat",
      "tissueLocation": "Blood",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "Buffy Coat",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-PPPXX2"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_BZ_4"
        },
        {
          "namespace": "investigatorId",
          "value": "LMNO_8080_N"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
    },
    {
      "smileSampleId": "afe74e17-8756-11eb-bb1d-acde48001122",
      "smilePatientId": "6cca38e3-875a-11eb-9e48-acde48001122",
      "cmoSampleName": "C-KXXL3J-N001-d",
      "sampleName": "LMNO_0034_N2",
      "sampleType": "Normal",
      "oncotreeCode": "SARCNOS",
      "collectionYear": "2019",
      "tubeId": "BLD-12345",
      "cfDNA2dBarcode": "8029250670",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "",
          "investigatorDecision": ""
        }
      ],
      "libraries": [
        {
          "barcodeId": "DUAL_TNG_LIB_6XX",
          "barcodeIndex": "TATAAACG-GGCTAGTT",
          "libraryIgoId": "22022_BZ_3_1_1_1",
          "libraryVolume": 35.0,
          "libraryConcentrationNgul": 23.4,
          "captureConcentrationNm": "2.1367521367521367",
          "captureInputNg": "50.0",
          "captureName": "Pool-22022_BZ-22022_CC-Tube7_1",
          "runs": [
            {
              "runMode": "HiSeq High Output",
              "runId": "CRX_7395",
              "flowCellId": "HXXXLBBYY",
              "readLength": "101/8/8/101",
              "runDate": "2020-05-20",
              "flowCellLanes": [
                1,
                2,
                3,
                4,
                5,
                6,
                7
              ],
              "fastqs": [
                "/FASTQ/Project_22022_BZ/Sample_LMNO_0034_N2_IGO_22022_BZ_3/LMNO_0034_N2_IGO_22022_BZ_3_S138_R1_001.fastq.gz",
                "/FASTQ/Project_22022_BZ/Sample_LMNO_0034_N2_IGO_22022_BZ_3/LMNO_0034_N2_IGO_22022_BZ_3_S138_R2_001.fastq.gz"
              ]
            }
          ]
        }
      ],
      "cmoPatientId": "C-KXXL3J",
      "primaryId": "22022_BZ_3",
      "investigatorSampleId": "LMNO_0034_N2",
      "species": "Human",
      "sex": "M",
      "tumorOrNormal": "Normal",
      "preservation": "EDTA-Streck",
      "sampleClass": "Blood",
      "sampleOrigin": "Buffy Coat",
      "tissueLocation": "",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "Buffy Coat",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-KXXL3J"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_BZ_3"
        },
        {
          "namespace": "investigatorId",
          "value": "LMNO_0034_N2"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
    },
    {
      "smileSampleId": "afe74d5c-8756-11eb-8a70-acde48001122",
      "smilePatientId": "6cca414a-875a-11eb-b133-acde48001122",
      "cmoSampleName": "C-X09281-N001-d",
      "sampleName": "LMNO_7787_N2",
      "sampleType": "Normal",
      "oncotreeCode": "MACR",
      "collectionYear": "2018",
      "tubeId": "",
      "cfDNA2dBarcode": "",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "",
          "investigatorDecision": "Continue processing"
        }
      ],
      "libraries": [
        {
          "barcodeId": "DUAL_IDT_LIB_XX8",
          "barcodeIndex": "GGTAGCCT-AAAGCTTA",
          "libraryIgoId": "22022_BZ_2_1_1_1",
          "libraryVolume": 35.0,
          "libraryConcentrationNgul": 24.7,
          "captureConcentrationNm": "2.0242914979757085",
          "captureInputNg": "50.0",
          "captureName": "Pool-22022_BZ-22022_CC-Tube7_1",
          "runs": [
            {
              "runMode": "HiSeq High Output",
              "runId": "CRX_7395",
              "flowCellId": "HXXXLBBYY",
              "readLength": "101/8/8/101",
              "runDate": "2020-05-20",
              "flowCellLanes": [
                1,
                2,
                3,
                4,
                5,
                6,
                7
              ],
              "fastqs": [
                "/FASTQ/Project_22022_BZ/Sample_LMNO_7787_N2_IGO_22022_BZ_2/LMNO_7787_N2_IGO_22022_BZ_2_S137_R1_001.fastq.gz",
                "/FASTQ/Project_22022_BZ/Sample_LMNO_7787_N2_IGO_22022_BZ_2/LMNO_7787_N2_IGO_22022_BZ_2_S137_R2_001.fastq.gz"
              ]
            }
          ]
        }
      ],
      "cmoPatientId": "C-X09281",
      "primaryId": "22022_BZ_2",
      "investigatorSampleId": "LMNO_7787_N2",
      "species": "Human",
      "sex": "M",
      "tumorOrNormal": "Normal",
      "preservation": "EDTA-Streck",
      "sampleClass": "Blood",
      "sampleOrigin": "Buffy Coat",
      "tissueLocation": "Blood",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "Buffy Coat",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-X09281"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_BZ_2"
        },
        {
          "namespace": "investigatorId",
          "value": "LMNO_7787_N2"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
    },
    {
      "smileSampleId": "afe74a42-8756-11eb-a483-acde48001122",
      "smilePatientId": "6cca438f-875a-11eb-8dda-acde48001122",
      "cmoSampleName": "C-XXA40X-P001-d",
      "sampleName": "LMNO_2020_DNA",
      "sampleType": "Primary",
      "oncotreeCode": "CHOS",
      "collectionYear": "2019",
      "tubeId": "TS-193394.2.T",
      "cfDNA2dBarcode": "8029250670",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "",
          "investigatorDecision": "Continue processing"
        }
      ],
      "libraries": [
        {
          "barcodeId": "DUAL_IDT_LIB_6X7",
          "barcodeIndex": "AGTGTACT-GCATACTT",
          "libraryIgoId": "22022_BZ_1_1_1_1",
          "libraryVolume": 35.0,
          "libraryConcentrationNgul": 23.3,
          "captureConcentrationNm": "5.150214592274678",
          "captureInputNg": "120.0",
          "captureName": "Pool-22022_BZ-22022_CC-Tube7_1",
          "runs": [
            {
              "runMode": "HiSeq High Output",
              "runId": "CRX_7395",
              "flowCellId": "HXXXLBBYY",
              "readLength": "101/8/8/101",
              "runDate": "2020-05-20",
              "flowCellLanes": [
                1,
                2,
                3,
                4,
                5,
                6,
                7
              ],
              "fastqs": [
                "/FASTQ/Project_22022_BZ/Sample_LMNO_2020_DNA_IGO_22022_BZ_1/LMNO_2020_DNA_IGO_22022_BZ_1_S133_R2_001.fastq.gz",
                "/FASTQ/Project_22022_BZ/Sample_LMNO_2020_DNA_IGO_22022_BZ_1/LMNO_2020_DNA_IGO_22022_BZ_1_S133_R1_001.fastq.gz"
              ]
            }
          ]
        }
      ],
      "cmoPatientId": "C-XXA40X",
      "primaryId": "22022_BZ_1",
      "investigatorSampleId": "LMNO_2020_DNA",
      "species": "Human",
      "sex": "M",
      "tumorOrNormal": "Tumor",
      "preservation": "OCT",
      "sampleClass": "Resection",
      "sampleOrigin": "Tissue",
      "tissueLocation": "Right Distal Femur",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "Tissue",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-XXA40X"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_BZ_1"
        },
        {
          "namespace": "investigatorId",
          "value": "LMNO_2020_DNA"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
    },
    {
      "smileSampleId": "afe74f80-8756-11eb-a0f5-acde48001122",
      "smilePatientId": "6cca45c7-875a-11eb-9cce-acde48001122",
      "cmoSampleName": "C-YXX89J-N001-d",
      "sampleName": "LMNO_3030_N",
      "sampleType": "Normal",
      "oncotreeCode": "TAML",
      "collectionYear": "2019",
      "tubeId": "BLD-33777",
      "cfDNA2dBarcode": "8029250670",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "",
          "investigatorDecision": ""
        }
      ],
      "libraries": [
        {
          "barcodeId": "DUAL_IDT_LIB_X51",
          "barcodeIndex": "TCACCGGC-GGCTTACT",
          "libraryIgoId": "22022_BZ_9_1_1_1",
          "libraryVolume": 35.0,
          "libraryConcentrationNgul": 20.7,
          "captureConcentrationNm": "2.4154589371980677",
          "captureInputNg": "50.0",
          "captureName": "Pool-22022_BZ-22022_CC-Tube7_1",
          "runs": [
            {
              "runMode": "HiSeq High Output",
              "runId": "CRX_7395",
              "flowCellId": "HXXXLBBYY",
              "readLength": "101/8/8/101",
              "runDate": "2020-05-20",
              "flowCellLanes": [
                1,
                2,
                3,
                4,
                5,
                6,
                7
              ],
              "fastqs": [
                "/FASTQ/Project_22022_BZ/Sample_LMNO_3030_N_IGO_22022_BZ_9/LMNO_3030_N_IGO_22022_BZ_9_S143_R1_001.fastq.gz",
                "/FASTQ/Project_22022_BZ/Sample_LMNO_3030_N_IGO_22022_BZ_9/LMNO_3030_N_IGO_22022_BZ_9_S143_R2_001.fastq.gz"
              ]
            }
          ]
        }
      ],
      "cmoPatientId": "C-YXX89J",
      "primaryId": "22022_BZ_9",
      "investigatorSampleId": "LMNO_3030_N",
      "species": "Human",
      "sex": "F",
      "tumorOrNormal": "Normal",
      "preservation": "EDTA-Streck",
      "sampleClass": "other",
      "sampleOrigin": "Buffy Coat",
      "tissueLocation": "",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "Buffy Coat",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-YXX89J"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_BZ_9"
        },
        {
          "namespace": "investigatorId",
          "value": "LMNO_3030_N"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
    },
    {
      "smileSampleId": "afe74f42-8756-11eb-8e66-acde48001122",
      "smilePatientId": "6cca3573-875a-11eb-911e-acde48001122",
      "cmoSampleName": "C-DPCXX1-M001-d",
      "sampleName": "LMNO_3443_T_DNA",
      "sampleType": "Metastasis",
      "oncotreeCode": "MPNWP",
      "collectionYear": "2019",
      "tubeId": "TS-019273.5.T",
      "cfDNA2dBarcode": "8029250670",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "",
          "investigatorDecision": ""
        }
      ],
      "libraries": [
        {
          "barcodeId": "DUAL_IDT_LIB_444",
          "barcodeIndex": "AGTTCATA-GGCCTTCA",
          "libraryIgoId": "22022_BZ_8_1_1_1",
          "libraryVolume": 35.0,
          "libraryConcentrationNgul": 24.2,
          "captureConcentrationNm": "4.958677685950414",
          "captureInputNg": "120.0",
          "captureName": "Pool-22022_BZ-22022_CC-Tube7_1",
          "runs": [
            {
              "runMode": "HiSeq High Output",
              "runId": "CRX_7395",
              "flowCellId": "HXXXLBBYY",
              "readLength": "101/8/8/101",
              "runDate": "2020-05-20",
              "flowCellLanes": [
                1,
                2,
                3,
                4,
                5,
                6,
                7
              ],
              "fastqs": [
                "/FASTQ/Project_22022_BZ/Sample_LMNO_3443_T_DNA_IGO_22022_BZ_8/LMNO_3443_T_DNA_IGO_22022_BZ_8_S142_R1_001.fastq.gz",
                "/FASTQ/Project_22022_BZ/Sample_LMNO_3443_T_DNA_IGO_22022_BZ_8/LMNO_3443_T_DNA_IGO_22022_BZ_8_S142_R2_001.fastq.gz"
              ]
            }
          ]
        }
      ],
      "cmoPatientId": "C-DPCXX1",
      "primaryId": "22022_BZ_8",
      "investigatorSampleId": "LMNO_3443_T_DNA",
      "species": "Human",
      "sex": "M",
      "tumorOrNormal": "Tumor",
      "preservation": "OCT",
      "sampleClass": "Resection",
      "sampleOrigin": "Tissue",
      "tissueLocation": "Lung/Middle Lobe",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "Tissue",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-DPCXX1"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_BZ_8"
        },
        {
          "namespace": "investigatorId",
          "value": "LMNO_3443_T_DNA"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
    },
    {
      "smileSampleId": "afe74c61-8756-11eb-8290-acde48001122",
      "smilePatientId": "6cca4a4f-875a-11eb-833e-acde48001122",
      "cmoSampleName": "C-999XX-M001-d",
      "sampleName": "LMNO_6565_T2_DNA",
      "sampleType": "Metastasis",
      "oncotreeCode": "MPNWP",
      "collectionYear": "2018",
      "tubeId": "TS-191305.1.T",
      "cfDNA2dBarcode": "8029250670",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "",
          "investigatorDecision": ""
        }
      ],
      "libraries": [
        {
          "barcodeId": "DUAL_IDT_LIB_XX2",
          "barcodeIndex": "AGTACTGC-GGTCTTAG",
          "libraryIgoId": "22022_BZ_16_1_1_1",
          "libraryVolume": 35.0,
          "libraryConcentrationNgul": 24.0,
          "captureConcentrationNm": "5.0",
          "captureInputNg": "120.0",
          "captureName": "Pool-22022_BZ-22022_CC-Tube7_1",
          "runs": [
            {
              "runMode": "HiSeq High Output",
              "runId": "CRX_7395",
              "flowCellId": "HXXXLBBYY",
              "readLength": "101/8/8/101",
              "runDate": "2020-05-20",
              "flowCellLanes": [
                1,
                2,
                3,
                4,
                5,
                6,
                7
              ],
              "fastqs": [
                "/FASTQ/Project_22022_BZ/Sample_LMNO_6565_T2_DNA_IGO_22022_BZ_16/LMNO_6565_T2_DNA_IGO_22022_BZ_16_S129_R1_001.fastq.gz",
                "/FASTQ/Project_22022_BZ/Sample_LMNO_6565_T2_DNA_IGO_22022_BZ_16/LMNO_6565_T2_DNA_IGO_22022_BZ_16_S129_R2_001.fastq.gz"
              ]
            }
          ]
        }
      ],
      "cmoPatientId": "C-999XX",
      "primaryId": "22022_BZ_16",
      "investigatorSampleId": "LMNO_6565_T2_DNA",
      "species": "Human",
      "sex": "F",
      "tumorOrNormal": "Tumor",
      "preservation": "OCT",
      "sampleClass": "Resection",
      "sampleOrigin": "Tissue",
      "tissueLocation": "Bladder Peritoneum",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "Tissue",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-999XX"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_BZ_16"
        },
        {
          "namespace": "investigatorId",
          "value": "LMNO_6565_T2_DNA"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
    },
    {
      "smileSampleId": "afe74bb5-8756-11eb-8365-acde48001122",
      "smilePatientId": "6cca4c9e-875a-11eb-b6d1-acde48001122",
      "cmoSampleName": "C-9XX8808-P001-d",
      "sampleName": "LMNO_0011_T",
      "sampleType": "Primary",
      "oncotreeCode": "GBC",
      "collectionYear": "2018",
      "tubeId": "TS-54321.9.T",
      "cfDNA2dBarcode": "8029250670",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "",
          "investigatorDecision": ""
        }
      ],
      "libraries": [
        {
          "barcodeId": "DUAL_IDT_LIB_YY9",
          "barcodeIndex": "TATACATC-AATCTGAA",
          "libraryIgoId": "22022_BZ_13_1_1_1",
          "libraryVolume": 35.0,
          "libraryConcentrationNgul": 25.3,
          "captureConcentrationNm": "4.743083003952569",
          "captureInputNg": "120.0",
          "captureName": "Pool-22022_BZ-22022_CC-Tube7_1",
          "runs": [
            {
              "runMode": "HiSeq High Output",
              "runId": "CRX_7395",
              "flowCellId": "HXXXLBBYY",
              "readLength": "101/8/8/101",
              "runDate": "2020-05-20",
              "flowCellLanes": [
                1,
                2,
                3,
                4,
                5,
                6,
                7
              ],
              "fastqs": [
                "/FASTQ/Project_22022_BZ/Sample_LMNO_0011_T_IGO_22022_BZ_13/LMNO_0011_T_IGO_22022_BZ_13_S126_R1_001.fastq.gz",
                "/FASTQ/Project_22022_BZ/Sample_LMNO_0011_T_IGO_22022_BZ_13/LMNO_0011_T_IGO_22022_BZ_13_S126_R2_001.fastq.gz"
              ]
            }
          ]
        }
      ],
      "cmoPatientId": "C-9XX8808",
      "primaryId": "22022_BZ_13",
      "investigatorSampleId": "LMNO_0011_T",
      "species": "Human",
      "sex": "F",
      "tumorOrNormal": "Tumor",
      "preservation": "OCT",
      "sampleClass": "Resection",
      "sampleOrigin": "Tissue",
      "tissueLocation": "Right Femur 4",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "Tissue",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-9XX8808"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_BZ_13"
        },
        {
          "namespace": "investigatorId",
          "value": "LMNO_0011_T"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
    },
    {
      "smileSampleId": "afe74ddc-8756-11eb-9c0d-acde48001122",
      "smilePatientId": "6cca4a4f-875a-11eb-833e-acde48001122",
      "cmoSampleName": "C-999XX-N002-d",
      "sampleName": "LMNO_6565_N2",
      "sampleType": "Normal",
      "oncotreeCode": "MPNWP",
      "collectionYear": "2019",
      "tubeId": "BLD-99111 BC",
      "cfDNA2dBarcode": "8029250670",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "",
          "investigatorDecision": ""
        }
      ],
      "libraries": [
        {
          "barcodeId": "DUAL_IDT_LIB_777",
          "barcodeIndex": "TGCGGGAC-TTAAGCCC",
          "libraryIgoId": "22022_BZ_21_1_1_1",
          "libraryVolume": 35.0,
          "libraryConcentrationNgul": 23.4,
          "captureConcentrationNm": "2.1367521367521367",
          "captureInputNg": "50.0",
          "captureName": "Pool-22022_BZ-22022_CC-Tube7_1",
          "runs": [
            {
              "runMode": "HiSeq High Output",
              "runId": "CRX_7395",
              "flowCellId": "HXXXLBBYY",
              "readLength": "101/8/8/101",
              "runDate": "2020-05-20",
              "flowCellLanes": [
                1,
                2,
                3,
                4,
                5,
                6,
                7
              ],
              "fastqs": [
                "/FASTQ/Project_22022_BZ/Sample_LMNO_6565_N2_IGO_22022_BZ_21/LMNO_6565_N2_IGO_22022_BZ_21_S135_R1_001.fastq.gz",
                "/FASTQ/Project_22022_BZ/Sample_LMNO_6565_N2_IGO_22022_BZ_21/LMNO_6565_N2_IGO_22022_BZ_21_S135_R2_001.fastq.gz"
              ]
            }
          ]
        }
      ],
      "cmoPatientId": "C-999XX",
      "primaryId": "22022_BZ_21",
      "investigatorSampleId": "LMNO_6565_N2",
      "species": "Human",
      "sex": "F",
      "tumorOrNormal": "Normal",
      "preservation": "EDTA-Streck",
      "sampleClass": "other",
      "sampleOrigin": "Buffy Coat",
      "tissueLocation": "",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "Buffy Coat",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-999XX"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_BZ_21"
        },
        {
          "namespace": "investigatorId",
          "value": "LMNO_6565_N2"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
    },
    {
      "smileSampleId": "afe74ae3-8756-11eb-a9f8-acde48001122",
      "smilePatientId": "6cca50e1-875a-11eb-bf5b-acde48001122",
      "cmoSampleName": "C-FFX222-N001-d",
      "sampleName": "LMNO_2299_N",
      "sampleType": "Normal",
      "oncotreeCode": "MACR",
      "collectionYear": "2017",
      "tubeId": "BLD-98765 BC",
      "cfDNA2dBarcode": "8029250670",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "",
          "investigatorDecision": ""
        }
      ],
      "libraries": [
        {
          "barcodeId": "DUAL_IDT_LIB_111",
          "barcodeIndex": "GGCCCTTA-GGCCTGAC",
          "libraryIgoId": "22022_BZ_10_1_1_1",
          "libraryVolume": 35.0,
          "libraryConcentrationNgul": 20.0,
          "captureConcentrationNm": "2.5",
          "captureInputNg": "50.0",
          "captureName": "Pool-22022_BZ-22022_CC-Tube7_1",
          "runs": [
            {
              "runMode": "HiSeq High Output",
              "runId": "CRX_7395",
              "flowCellId": "HXXXLBBYY",
              "readLength": "101/8/8/101",
              "runDate": "2020-05-20",
              "flowCellLanes": [
                1,
                2,
                3,
                4,
                5,
                6,
                7
              ],
              "fastqs": [
                "/FASTQ/Project_22022_BZ/Sample_LMNO_2299_N_IGO_22022_BZ_10/LMNO_2299_N_IGO_22022_BZ_10_S123_R1_001.fastq.gz",
                "/FASTQ/Project_22022_BZ/Sample_LMNO_2299_N_IGO_22022_BZ_10/LMNO_2299_N_IGO_22022_BZ_10_S123_R2_001.fastq.gz"
              ]
            }
          ]
        }
      ],
      "cmoPatientId": "C-FFX222",
      "primaryId": "22022_BZ_10",
      "investigatorSampleId": "LMNO_2299_N",
      "species": "Human",
      "sex": "F",
      "tumorOrNormal": "Normal",
      "preservation": "EDTA-Streck",
      "sampleClass": "other",
      "sampleOrigin": "Buffy Coat",
      "tissueLocation": "",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "Buffy Coat",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-FFX222"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_BZ_10"
        },
        {
          "namespace": "investigatorId",
          "value": "LMNO_2299_N"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
    },
    {
      "smileSampleId": "afe74e91-8756-11eb-8e29-acde48001122",
      "smilePatientId": "6cca3d19-875a-11eb-92e8-acde48001122",
      "cmoSampleName": "C-PPPXX2-P001-d",
      "sampleName": "LMNO_8080_T_DNA",
      "sampleType": "Primary",
      "oncotreeCode": "COAD",
      "collectionYear": "",
      "tubeId": "TS.928838.4.T",
      "cfDNA2dBarcode": "8029250670",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "",
          "investigatorDecision": ""
        }
      ],
      "libraries": [
        {
          "barcodeId": "DUAL_IDT_LIB_233",
          "barcodeIndex": "AGGCTGAG-GGTCCTCA",
          "libraryIgoId": "22022_BZ_5_1_1_1",
          "libraryVolume": 35.0,
          "libraryConcentrationNgul": 32.6,
          "captureConcentrationNm": "2.7607361963190185",
          "captureInputNg": "90.0",
          "captureName": "Pool-22022_BZ-22022_CC-Tube7_1",
          "runs": [
            {
              "runMode": "HiSeq High Output",
              "runId": "CRX_7395",
              "flowCellId": "HXXXLBBYY",
              "readLength": "101/8/8/101",
              "runDate": "2020-05-20",
              "flowCellLanes": [
                1,
                2,
                3,
                4,
                5,
                6,
                7
              ],
              "fastqs": [
                "/FASTQ/Project_22022_BZ/Sample_LMNO_8080_T_DNA_IGO_22022_BZ_5/LMNO_8080_T_DNA_IGO_22022_BZ_5_S139_R1_001.fastq.gz",
                "/FASTQ/Project_22022_BZ/Sample_LMNO_8080_T_DNA_IGO_22022_BZ_5/LMNO_8080_T_DNA_IGO_22022_BZ_5_S139_R2_001.fastq.gz"
              ]
            }
          ]
        }
      ],
      "cmoPatientId": "C-PPPXX2",
      "primaryId": "22022_BZ_5",
      "investigatorSampleId": "LMNO_8080_T_DNA",
      "species": "Human",
      "sex": "F",
      "tumorOrNormal": "Tumor",
      "preservation": "OCT",
      "sampleClass": "Resection",
      "sampleOrigin": "Tissue",
      "tissueLocation": "Liver/Right lobe",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "Tissue",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-PPPXX2"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_BZ_5"
        },
        {
          "namespace": "investigatorId",
          "value": "LMNO_8080_T_DNA"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
    },
    {
      "smileSampleId": "afe74b2b-8756-11eb-82f5-acde48001122",
      "smilePatientId": "6cca45c7-875a-11eb-9cce-acde48001122",
      "cmoSampleName": "C-YXX89J-P001-d",
      "sampleName": "LMNO_3030_T_DNA",
      "sampleType": "Primary",
      "oncotreeCode": "TAML",
      "collectionYear": "2019",
      "tubeId": "TS-223344.0.T",
      "cfDNA2dBarcode": "8029250670",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "",
          "investigatorDecision": ""
        }
      ],
      "libraries": [
        {
          "barcodeId": "DUAL_IDT_LIB_767",
          "barcodeIndex": "TTAGTGTG-GGTAGCTA",
          "libraryIgoId": "22022_BZ_11_1_1_1",
          "libraryVolume": 35.0,
          "libraryConcentrationNgul": 31.6,
          "captureConcentrationNm": "2.848101265822785",
          "captureInputNg": "90.0",
          "captureName": "Pool-22022_BZ-22022_CC-Tube7_1",
          "runs": [
            {
              "runMode": "HiSeq High Output",
              "runId": "CRX_7395",
              "flowCellId": "HXXXLBBYY",
              "readLength": "101/8/8/101",
              "runDate": "2020-05-20",
              "flowCellLanes": [
                1,
                2,
                3,
                4,
                5,
                6,
                7
              ],
              "fastqs": [
                "/FASTQ/Project_22022_BZ/Sample_LMNO_3030_T_DNA_IGO_22022_BZ_11/LMNO_3030_T_DNA_IGO_22022_BZ_11_S124_R1_001.fastq.gz",
                "/FASTQ/Project_22022_BZ/Sample_LMNO_3030_T_DNA_IGO_22022_BZ_11/LMNO_3030_T_DNA_IGO_22022_BZ_11_S124_R2_001.fastq.gz"
              ]
            }
          ]
        }
      ],
      "cmoPatientId": "C-YXX89J",
      "primaryId": "22022_BZ_11",
      "investigatorSampleId": "LMNO_3030_T_DNA",
      "species": "Human",
      "sex": "F",
      "tumorOrNormal": "Tumor",
      "preservation": "OCT",
      "sampleClass": "Resection",
      "sampleOrigin": "Tissue",
      "tissueLocation": "left retroperitoneum",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "Tissue",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-YXX89J"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_BZ_11"
        },
        {
          "namespace": "investigatorId",
          "value": "LMNO_3030_T_DNA"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
    },
    {
      "smileSampleId": "afe74b70-8756-11eb-8386-acde48001122",
      "smilePatientId": "6cca3bdc-875a-11eb-a412-acde48001122",
      "cmoSampleName": "C-XXX711-N001-d",
      "sampleName": "WT_PXX_bc",
      "sampleType": "Normal",
      "oncotreeCode": "PAAD",
      "collectionYear": "2017",
      "tubeId": "88354511236",
      "cfDNA2dBarcode": "8029250670",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "",
          "investigatorDecision": ""
        }
      ],
      "libraries": [
        {
          "barcodeId": "DUAL_IDT_LIB_658",
          "barcodeIndex": "AGTACGCG-GGCAATAA",
          "libraryIgoId": "22022_BZ_12_1_1_1",
          "libraryVolume": 35.0,
          "libraryConcentrationNgul": 21.3,
          "captureConcentrationNm": "2.347417840375587",
          "captureInputNg": "50.0",
          "captureName": "Pool-22022_BZ-22022_CC-Tube7_1",
          "runs": [
            {
              "runMode": "HiSeq High Output",
              "runId": "CRX_7395",
              "flowCellId": "HXXXLBBYY",
              "readLength": "101/8/8/101",
              "runDate": "2020-05-20",
              "flowCellLanes": [
                1,
                2,
                3,
                4,
                5,
                6,
                7
              ],
              "fastqs": [
                "/FASTQ/Project_22022_BZ/Sample_WT_PXX_bc_IGO_22022_BZ_12/WT_PXX_bc_IGO_22022_BZ_12_S125_R1_001.fastq.gz",
                "/FASTQ/Project_22022_BZ/Sample_WT_PXX_bc_IGO_22022_BZ_12/WT_PXX_bc_IGO_22022_BZ_12_S125_R2_001.fastq.gz"
              ]
            }
          ]
        }
      ],
      "cmoPatientId": "C-XXX711",
      "primaryId": "22022_BZ_12",
      "investigatorSampleId": "WT_PXX_bc",
      "species": "Human",
      "sex": "M",
      "tumorOrNormal": "",
      "preservation": "EDTA-Streck",
      "sampleClass": "Blood",
      "sampleOrigin": "Buffy Coat",
      "tissueLocation": "",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "Buffy Coat",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-XXX711"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_BZ_12"
        },
        {
          "namespace": "investigatorId",
          "value": "WT_PXX_bc"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
    },
    {
      "smileSampleId": "afe74ce1-8756-11eb-8145-acde48001122",
      "smilePatientId": "6cca4c9e-875a-11eb-b6d1-acde48001122",
      "cmoSampleName": "C-9XX8808-N001-d",
      "sampleName": "LMNO_0011_N",
      "sampleType": "Normal",
      "oncotreeCode": "GBC",
      "collectionYear": "2018",
      "tubeId": "BLD-98765",
      "cfDNA2dBarcode": "8029250670",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "",
          "investigatorDecision": "Continue processing"
        }
      ],
      "libraries": [
        {
          "barcodeId": "DUAL_IDT_LIB_332",
          "barcodeIndex": "TTGGGGCA-TTTTTATT",
          "libraryIgoId": "22022_BZ_18_1_1_1",
          "libraryVolume": 35.0,
          "libraryConcentrationNgul": 24.3,
          "captureConcentrationNm": "2.05761316872428",
          "captureInputNg": "50.0",
          "captureName": "Pool-22022_BZ-22022_CC-Tube7_1",
          "runs": [
            {
              "runMode": "HiSeq High Output",
              "runId": "CRX_7395",
              "flowCellId": "HXXXLBBYY",
              "readLength": "101/8/8/101",
              "runDate": "2020-05-20",
              "flowCellLanes": [
                1,
                2,
                3,
                4,
                5,
                6,
                7
              ],
              "fastqs": [
                "/FASTQ/Project_22022_BZ/Sample_LMNO_0011_N_IGO_22022_BZ_18/LMNO_0011_N_IGO_22022_BZ_18_S131_R2_001.fastq.gz",
                "/FASTQ/Project_22022_BZ/Sample_LMNO_0011_N_IGO_22022_BZ_18/LMNO_0011_N_IGO_22022_BZ_18_S131_R1_001.fastq.gz"
              ]
            }
          ]
        }
      ],
      "cmoPatientId": "C-9XX8808",
      "primaryId": "22022_BZ_18",
      "investigatorSampleId": "LMNO_0011_N",
      "species": "Human",
      "sex": "F",
      "tumorOrNormal": "Normal",
      "preservation": "EDTA-Streck",
      "sampleClass": "Blood",
      "sampleOrigin": "Buffy Coat",
      "tissueLocation": "",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "Buffy Coat",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-9XX8808"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_BZ_18"
        },
        {
          "namespace": "investigatorId",
          "value": "LMNO_0011_N"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
    },
    {
      "smileSampleId": "afe74ecc-8756-11eb-a6cc-acde48001122",
      "smilePatientId": "6cca5b30-875a-11eb-9186-acde48001122",
      "cmoSampleName": "C-XXC4XX-N001-d",
      "sampleName": "LMNO_8895_N",
      "sampleType": "Normal",
      "oncotreeCode": "GBC",
      "collectionYear": "",
      "tubeId": "",
      "cfDNA2dBarcode": "",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "",
          "investigatorDecision": ""
        }
      ],
      "libraries": [
        {
          "barcodeId": "DUAL_IDT_LIB_321",
          "barcodeIndex": "AAAGCTGA-AATTTCGA",
          "libraryIgoId": "22022_BZ_6_1_1_1",
          "libraryVolume": 35.0,
          "libraryConcentrationNgul": 25.0,
          "captureConcentrationNm": "2.0",
          "captureInputNg": "50.0",
          "captureName": "Pool-22022_BZ-22022_CC-Tube7_1",
          "runs": [
            {
              "runMode": "HiSeq High Output",
              "runId": "CRX_7395",
              "flowCellId": "HXXXLBBYY",
              "readLength": "101/8/8/101",
              "runDate": "2020-05-20",
              "flowCellLanes": [
                1,
                2,
                3,
                4,
                5,
                6,
                7
              ],
              "fastqs": [
                "/FASTQ/Project_22022_BZ/Sample_LMNO_8895_N_IGO_22022_BZ_6/LMNO_8895_N_IGO_22022_BZ_6_S140_R1_001.fastq.gz",
                "/FASTQ/Project_22022_BZ/Sample_LMNO_8895_N_IGO_22022_BZ_6/LMNO_8895_N_IGO_22022_BZ_6_S140_R2_001.fastq.gz"
              ]
            }
          ]
        }
      ],
      "cmoPatientId": "C-XXC4XX",
      "primaryId": "22022_BZ_6",
      "investigatorSampleId": "LMNO_8895_N",
      "species": "Human",
      "sex": "F",
      "tumorOrNormal": "Normal",
      "preservation": "EDTA-Streck",
      "sampleClass": "Blood",
      "sampleOrigin": "Buffy Coat",
      "tissueLocation": "",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "Buffy Coat",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-XXC4XX"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_BZ_6"
        },
        {
          "namespace": "investigatorId",
          "value": "LMNO_8895_N"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
    },
    {
      "smileSampleId": "afe74da1-8756-11eb-8903-acde48001122",
      "smilePatientId": "6cca438f-875a-11eb-8dda-acde48001122",
      "cmoSampleName": "C-XXA40X-N001-d",
      "sampleName": "LMNO_2020_N",
      "sampleType": "Normal",
      "oncotreeCode": "CHOS",
      "collectionYear": "2019",
      "tubeId": "BLD-98981 BC",
      "cfDNA2dBarcode": "8029250670",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "",
          "investigatorDecision": ""
        }
      ],
      "libraries": [
        {
          "barcodeId": "DUAL_IDT_LIB_866",
          "barcodeIndex": "GGGGCTGA-TTTATATA",
          "libraryIgoId": "22022_BZ_20_1_1_1",
          "libraryVolume": 35.0,
          "libraryConcentrationNgul": 23.4,
          "captureConcentrationNm": "2.1367521367521367",
          "captureInputNg": "50.0",
          "captureName": "Pool-22022_BZ-22022_CC-Tube7_1",
          "runs": [
            {
              "runMode": "HiSeq High Output",
              "runId": "CRX_7395",
              "flowCellId": "HXXXLBBYY",
              "readLength": "101/8/8/101",
              "runDate": "2020-05-20",
              "flowCellLanes": [
                1,
                2,
                3,
                4,
                5,
                6,
                7
              ],
              "fastqs": [
                "/FASTQ/Project_22022_BZ/Sample_LMNO_2020_N_IGO_22022_BZ_20/LMNO_2020_N_IGO_22022_BZ_20_S134_R1_001.fastq.gz",
                "/FASTQ/Project_22022_BZ/Sample_LMNO_2020_N_IGO_22022_BZ_20/LMNO_2020_N_IGO_22022_BZ_20_S134_R2_001.fastq.gz"
              ]
            }
          ]
        }
      ],
      "cmoPatientId": "C-XXA40X",
      "primaryId": "22022_BZ_20",
      "investigatorSampleId": "LMNO_2020_N",
      "species": "Human",
      "sex": "M",
      "tumorOrNormal": "Normal",
      "preservation": "EDTA-Streck",
      "sampleClass": "other",
      "sampleOrigin": "Buffy Coat",
      "tissueLocation": "Buffy Coat",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "Buffy Coat",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-XXA40X"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_BZ_20"
        },
        {
          "namespace": "investigatorId",
          "value": "LMNO_2020_N"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
    },
    {
      "smileSampleId": "afe74bf3-8756-11eb-9ed7-acde48001122",
      "smilePatientId": "6cca414a-875a-11eb-b133-acde48001122",
      "cmoSampleName": "C-X09281-M001-d",
      "sampleName": "LMNO_7787_T",
      "sampleType": "Metastasis",
      "oncotreeCode": "MACR",
      "collectionYear": "2018",
      "tubeId": "TS-828282",
      "cfDNA2dBarcode": "8029250670",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "",
          "investigatorDecision": ""
        }
      ],
      "libraries": [
        {
          "barcodeId": "DUAL_IDT_LIB_990",
          "barcodeIndex": "AAGGCTGA-GGCTCTAG",
          "libraryIgoId": "22022_BZ_14_1_1_1",
          "libraryVolume": 35.0,
          "libraryConcentrationNgul": 25.9,
          "captureConcentrationNm": "3.8610038610038613",
          "captureInputNg": "100.0",
          "captureName": "Pool-22022_BZ-22022_CC-Tube7_1",
          "runs": [
            {
              "runMode": "HiSeq High Output",
              "runId": "CRX_7395",
              "flowCellId": "HXXXLBBYY",
              "readLength": "101/8/8/101",
              "runDate": "2020-05-20",
              "flowCellLanes": [
                1,
                2,
                3,
                4,
                5,
                6,
                7
              ],
              "fastqs": [
                "/FASTQ/Project_22022_BZ/Sample_LMNO_7787_T_IGO_22022_BZ_14/LMNO_7787_T_IGO_22022_BZ_14_S127_R1_001.fastq.gz",
                "/FASTQ/Project_22022_BZ/Sample_LMNO_7787_T_IGO_22022_BZ_14/LMNO_7787_T_IGO_22022_BZ_14_S127_R2_001.fastq.gz"
              ]
            }
          ]
        }
      ],
      "cmoPatientId": "C-X09281",
      "primaryId": "22022_BZ_14",
      "investigatorSampleId": "LMNO_7787_T",
      "species": "Human",
      "sex": "M",
      "tumorOrNormal": "Tumor",
      "preservation": "OCT",
      "sampleClass": "Resection",
      "sampleOrigin": "Tissue",
      "tissueLocation": "Left Kidney Hilar Lymph Node",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "Tissue",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-X09281"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_BZ_14"
        },
        {
          "namespace": "investigatorId",
          "value": "LMNO_7787_T"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
    },
    {
      "smileSampleId": "afe74f11-8756-11eb-a0db-acde48001122",
      "smilePatientId": "6cca50e1-875a-11eb-bf5b-acde48001122",
      "cmoSampleName": "C-FFX222-P001-d",
      "sampleName": "LMNO_2299_T_DNA",
      "sampleType": "Primary",
      "oncotreeCode": "MACR",
      "collectionYear": "2019",
      "tubeId": "TS-888999.1.T",
      "cfDNA2dBarcode": "8029250670",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "",
          "investigatorDecision": ""
        }
      ],
      "libraries": [
        {
          "barcodeId": "DUAL_IDT_LIB_653",
          "barcodeIndex": "TCAATTAT-CCTGAAGG",
          "libraryIgoId": "22022_BZ_7_1_1_1",
          "libraryVolume": 35.0,
          "libraryConcentrationNgul": 26.7,
          "captureConcentrationNm": "3.7453183520599254",
          "captureInputNg": "100.0",
          "captureName": "Pool-22022_BZ-22022_CC-Tube7_1",
          "runs": [
            {
              "runMode": "HiSeq High Output",
              "runId": "CRX_7395",
              "flowCellId": "HXXXLBBYY",
              "readLength": "101/8/8/101",
              "runDate": "2020-05-20",
              "flowCellLanes": [
                1,
                2,
                3,
                4,
                5,
                6,
                7
              ],
              "fastqs": [
                "/FASTQ/Project_22022_BZ/Sample_LMNO_2299_T_DNA_IGO_22022_BZ_7/LMNO_2299_T_DNA_IGO_22022_BZ_7_S141_R1_001.fastq.gz",
                "/FASTQ/Project_22022_BZ/Sample_LMNO_2299_T_DNA_IGO_22022_BZ_7/LMNO_2299_T_DNA_IGO_22022_BZ_7_S141_R2_001.fastq.gz"
              ]
            }
          ]
        }
      ],
      "cmoPatientId": "C-FFX222",
      "primaryId": "22022_BZ_7",
      "investigatorSampleId": "LMNO_2299_T_DNA",
      "species": "Human",
      "sex": "F",
      "tumorOrNormal": "Tumor",
      "preservation": "OCT",
      "sampleClass": "Resection",
      "sampleOrigin": "Tissue",
      "tissueLocation": "Abdomen",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "Tissue",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-FFX222"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_BZ_7"
        },
        {
          "namespace": "investigatorId",
          "value": "LMNO_2299_T_DNA"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
    }
  ],
  "pooledNormals": [
    "/FASTQ/Project_POOLEDNORMALS/Sample_FFPEPOOLEDNORMAL_IGO_GENESET101_ACCGTCCT/FFPEPOOLEDNORMAL_IGO_GENESET101_ACCGTCCT_S118_R1_001.fastq.gz",
    "/FASTQ/Project_POOLEDNORMALS/Sample_FFPEPOOLEDNORMAL_IGO_GENESET101_ACCGTCCT/FFPEPOOLEDNORMAL_IGO_GENESET101_ACCGTCCT_S118_R2_001.fastq.gz",
    "/FASTQ/Project_POOLEDNORMALS/Sample_FROZENPOOLEDNORMAL_IGO_GENESET101_TGTCTAAC/FROZENPOOLEDNORMAL_IGO_GENESET101_TGTCTAAC_S22_R1_001.fastq.gz",
    "/FASTQ/Project_POOLEDNORMALS/Sample_FROZENPOOLEDNORMAL_IGO_GENESET101_TGTCTAAC/FROZENPOOLEDNORMAL_IGO_GENESET101_TGTCTAAC_S22_R2_001.fastq.gz"
  ],
  "igoProjectId": "22022"
}
`

var updatedRequest = `
[
{
  "smileRequestId": "6cc9f987-875a-11eb-ada3-acde48001122",
  "igoRequestId": "22022_BZ",
  "genePanel": "GENESET101_BAITS_NEWNAME",
  "projectManagerName": "Bar, Foo",
  "piEmail": "updated_pi_email@mskcc.org",
  "labHeadName": "Foo Bar",
  "labHeadEmail": "request1pi_updated@mskcc.org",
  "investigatorName": "John Smith",
  "investigatorEmail": "smithj@mskcc.org",
  "dataAnalystName": "Poin Dexter",
  "dataAnalystEmail": "dexterp@mskcc.org",
  "otherContactEmails": "dexterp@mskcc.org",
  "dataAccessEmails": "",
  "qcAccessEmails": "",
  "isCmoRequest": true,
  "bicAnalysis": false,
  "pooledNormals": [
    "/FASTQ/Project_POOLEDNORMALS/Sample_FFPEPOOLEDNORMAL_IGO_GENESET101_TAGCTTGA/FFPEPOOLEDNORMAL_IGO_GENESET101_TAGCTTGA_S23_R1_001.fastq.gz",
    "/FASTQ/Project_POOLEDNORMALS/Sample_FFPEPOOLEDNORMAL_IGO_GENESET101_TAGCTTGA/FFPEPOOLEDNORMAL_IGO_GENESET101_TAGCTTGA_S23_R2_001.fastq.gz",
    "/FASTQ/Project_POOLEDNORMALS/Sample_FROZENPOOLEDNORMAL_IGO_AMBIGUOUS_TTAGGCTG/FROZENPOOLEDNORMAL_IGO_AMBIGUOUS_TTAGGCTG_S86_R1_001.fastq.gz",
    "/FASTQ/Project_POOLEDNORMALS/Sample_FROZENPOOLEDNORMAL_IGO_AMBIGUOUS_TTAGGCTG/FROZENPOOLEDNORMAL_IGO_AMBIGUOUS_TTAGGCTG_S86_R2_001.fastq.gz",
    "/FASTQ/Project_POOLEDNORMALS/Sample_FROZENPOOLEDNORMAL_IGO_HEMESET_v1_TTAGGCTG/FROZENPOOLEDNORMAL_IGO_HEMESET_v1_TTAGGCTG_S147_R1_001.fastq.gz",
    "/FASTQ/Project_POOLEDNORMALS/Sample_FROZENPOOLEDNORMAL_IGO_HEMESET_v1_TTAGGCTG/FROZENPOOLEDNORMAL_IGO_HEMESET_v1_TTAGGCTG_S147_R2_001.fastq.gz",
    "/FASTQ/Project_POOLEDNORMALS/Sample_FROZENPOOLEDNORMAL_IGO_GENESET101_TTAGGCTG/FROZENPOOLEDNORMAL_IGO_GENESET101_TTAGGCTG_S196_R1_001.fastq.gz",
    "/FASTQ/Project_POOLEDNORMALS/Sample_FROZENPOOLEDNORMAL_IGO_GENESET101_TTAGGCTG/FROZENPOOLEDNORMAL_IGO_GENESET101_TTAGGCTG_S196_R2_001.fastq.gz",
    "/FASTQ/Project_POOLEDNORMALS/Sample_MOUSEPOOLEDNORMAL_IGO_AMBIGUOUS_GGCGTCAT/MOUSEPOOLEDNORMAL_IGO_AMBIGUOUS_GGCGTCAT_S61_R1_001.fastq.gz",
    "/FASTQ/Project_POOLEDNORMALS/Sample_MOUSEPOOLEDNORMAL_IGO_AMBIGUOUS_GGCGTCAT/MOUSEPOOLEDNORMAL_IGO_AMBIGUOUS_GGCGTCAT_S61_R2_001.fastq.gz"
  ],
  "igoProjectId": "22022"
},
{
  "smileRequestId": "6cca6166-875a-11eb-ae9e-acde48001122",
  "igoRequestId": "22022_BZ",
  "genePanel": "GENESET101_BAITS",
  "projectManagerName": "bart simpson",
  "piEmail": "pi_email@mskcc.org",
  "labHeadName": "beavis",
  "labHeadEmail": "beavis@mskcc.org",
  "investigatorName": "butthead",
  "investigatorEmail": "butthead@mskcc.org",
  "dataAnalystName": "",
  "dataAnalystEmail": "",
  "otherContactEmails": "metallica@mskcc.org",
  "dataAccessEmails": "",
  "qcAccessEmails": "",
  "isCmoRequest": true,
  "bicAnalysis": false,
  "pooledNormals": [
    "/FASTQ/Project_POOLEDNORMALS/Sample_FFPEPOOLEDNORMAL_IGO_GENESET101_ACCGTCCT/FFPEPOOLEDNORMAL_IGO_GENESET101_ACCGTCCT_S118_R1_001.fastq.gz",
    "/FASTQ/Project_POOLEDNORMALS/Sample_FFPEPOOLEDNORMAL_IGO_GENESET101_ACCGTCCT/FFPEPOOLEDNORMAL_IGO_GENESET101_ACCGTCCT_S118_R2_001.fastq.gz",
    "/FASTQ/Project_POOLEDNORMALS/Sample_FROZENPOOLEDNORMAL_IGO_GENESET101_TGTCTAAC/FROZENPOOLEDNORMAL_IGO_GENESET101_TGTCTAAC_S22_R1_001.fastq.gz",
    "/FASTQ/Project_POOLEDNORMALS/Sample_FROZENPOOLEDNORMAL_IGO_GENESET101_TGTCTAAC/FROZENPOOLEDNORMAL_IGO_GENESET101_TGTCTAAC_S22_R2_001.fastq.gz"
  ],
  "igoProjectId": "22022"
}
]
`
var updatedSample = `
[
{
      "cmoSampleName": "C-TX6DNG-N001-d-UPDATE",
      "sampleName": "LMNO_4396_N",
      "sampleType": "Normal",
      "oncotreeCode": "TPLL",
      "collectionYear": "",
      "tubeId": "4157451784",
      "cfDNA2dBarcode": "8029250670",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "",
          "investigatorDecision": "Continue processing"
        }
      ],
      "libraries": [
        {
          "libraryIgoId": "22022_CC_3_1",
          "libraryConcentrationNgul": 34.2,
          "captureConcentrationNm": "1.461988304093567",
          "captureInputNg": "50.0",
          "captureName": "Pool-22022_BZ-22022_CC-Tube7_1",
          "runs": [
            {
              "runMode": "HiSeq High Output",
              "runId": "CRX_7395",
              "flowCellId": "HGJMLBBXY",
              "readLength": "101/8/8/101",
              "runDate": "2020-05-20",
              "flowCellLanes": [
                1,
                2,
                3,
                4,
                5,
                6,
                7
              ],
              "fastqs": [
                "/FASTQ/Project_22022_CC/Sample_LMNO_4396_N_IGO_22022_CC_3/LMNO_4396_N_IGO_22022_CC_3_S144_R1_001.fastq.gz",
                "/FASTQ/Project_22022_CC/Sample_LMNO_4396_N_IGO_22022_CC_3/LMNO_4396_N_IGO_22022_CC_3_S144_R2_001.fastq.gz"
              ]
            }
          ]
        }
      ],
      "cmoPatientId": "C-TX6DNG",
      "primaryId": "22022_CC_3",
      "investigatorSampleId": "LMNO_4396_N",
      "species": "Human",
      "sex": "F",
      "tumorOrNormal": "Normal",
      "preservation": "EDTA-Streck",
      "sampleClass": "Blood",
      "sampleOrigin": "Buffy Coat",
      "tissueLocation": "Blood",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "Buffy Coat",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-TX6DNG"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_CC_3"
        },
        {
          "namespace": "investigatorId",
          "value": "LMNO_4396_N"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
},
{
      "cmoSampleName": "C-TX6DNG-N001-d",
      "sampleName": "LMNO_4396_N",
      "sampleType": "Normal",
      "oncotreeCode": "TPLL",
      "collectionYear": "",
      "tubeId": "4157451784",
      "cfDNA2dBarcode": "8029250670",
      "qcReports": [
        {
          "qcReportType": "LIBRARY",
          "comments": "",
          "investigatorDecision": "Continue processing"
        }
      ],
      "libraries": [
        {
          "libraryIgoId": "22022_CC_3_1",
          "libraryConcentrationNgul": 34.2,
          "captureConcentrationNm": "1.461988304093567",
          "captureInputNg": "50.0",
          "captureName": "Pool-22022_BZ-22022_CC-Tube7_1",
          "runs": [
            {
              "runMode": "HiSeq High Output",
              "runId": "CRX_7395",
              "flowCellId": "HGJMLBBXY",
              "readLength": "101/8/8/101",
              "runDate": "2020-05-20",
              "flowCellLanes": [
                1,
                2,
                3,
                4,
                5,
                6,
                7
              ],
              "fastqs": [
                "/FASTQ/Project_22022_CC/Sample_LMNO_4396_N_IGO_22022_CC_3/LMNO_4396_N_IGO_22022_CC_3_S144_R1_001.fastq.gz",
                "/FASTQ/Project_22022_CC/Sample_LMNO_4396_N_IGO_22022_CC_3/LMNO_4396_N_IGO_22022_CC_3_S144_R2_001.fastq.gz"
              ]
            }
          ]
        }
      ],
      "cmoPatientId": "C-TX6DNG",
      "primaryId": "22022_CC_3",
      "investigatorSampleId": "LMNO_4396_N",
      "species": "Human",
      "sex": "F",
      "tumorOrNormal": "Normal",
      "preservation": "EDTA-Streck",
      "sampleClass": "Blood",
      "sampleOrigin": "Buffy Coat",
      "tissueLocation": "Blood",
      "baitSet": "GENESET101_BAITS",
      "genePanel": "GENESET101_BAITS",
      "datasource": "igo",
      "igoComplete": true,
      "cmoSampleIdFields": {
        "naToExtract": "",
        "sampleType": "Buffy Coat",
        "normalizedPatientId": "MRN_REDACTED",
        "recipe": "GENESET101_BAITS"
      },
      "patientAliases": [
        {
          "namespace": "cmo",
          "value": "C-TX6DNG"
        }
      ],
      "sampleAliases": [
        {
          "namespace": "igoId",
          "value": "22022_CC_3"
        },
        {
          "namespace": "investigatorId",
          "value": "LMNO_4396_N"
        }
      ],
      "additionalProperties": {
        "isCmoSample": "true",
        "igoRequestId": "22022_BZ"
      }
}
]
`

var dArgs = dremio.DremioArgs{
	Host:         "localhost",
	Username:     "",
	Password:     "",
	ObjectStore:  "\"local-minio\".smile",
	RequestTable: "requests",
	SampleTable:  "samples",
}

func TestNewRequest(t *testing.T) {

	var r smile.Request
	err := json.Unmarshal([]byte(newRequest), &r)
	if err != nil {
		t.Error(err)
	}

	dr, err := dremio.NewDremioRepos(dArgs)
	if err != nil {
		t.Error(err)
	}
	err = dr.AddRequest(context.Background(), r)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateRequest(t *testing.T) {

	var r []smile.Request
	err := json.Unmarshal([]byte(updatedRequest), &r)
	if err != nil {
		t.Error(err)
	}

	dr, err := dremio.NewDremioRepos(dArgs)
	if err != nil {
		t.Error(err)
	}
	// tbd: query before and after
	err = dr.UpdateRequest(context.Background(), r)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateSample(t *testing.T) {

	var s []smile.Sample
	err := json.Unmarshal([]byte(updatedSample), &s)
	if err != nil {
		t.Error(err)
	}

	dr, err := dremio.NewDremioRepos(dArgs)
	if err != nil {
		t.Error(err)
	}
	// tbd: query before and after
	err = dr.UpdateSample(context.Background(), s)
	if err != nil {
		t.Error(err)
	}
}
