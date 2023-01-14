package io

const (
	NIIVersion1 = 1
	NIIVersion2 = 2
)

const (
	NII1HeaderSize = 348
	NII2HeaderSize = 540
)

// Nii1Header defines the structure of the NIFTI-1 header
type Nii1Header struct {
	SizeofHdr      int32
	DataTypeUnused [10]uint8
	DbName         [18]uint8
	Extents        int32
	SessionError   int16
	Regular        uint8
	DimInfo        uint8
	Dim            [8]int16
	IntentP1       float32
	IntentP2       float32
	IntentP3       float32
	IntentCode     int16
	Datatype       int16
	Bitpix         int16
	SliceStart     int16
	Pixdim         [8]float32
	VoxOffset      float32
	SclSlope       float32
	SclInter       float32
	SliceEnd       int16
	SliceCode      uint8
	XyztUnits      uint8
	CalMax         float32
	CalMin         float32
	SliceDuration  float32
	Toffset        float32
	Glmax          int32
	Glmin          int32
	Descrip        [80]uint8
	AuxFile        [24]uint8
	QformCode      int16
	SformCode      int16
	QuaternB       float32
	QuaternC       float32
	QuaternD       float32
	QoffsetX       float32
	QoffsetY       float32
	QoffsetZ       float32
	SrowX          [4]float32
	SrowY          [4]float32
	SrowZ          [4]float32
	IntentName     [16]uint8
	Magic          [4]uint8
}

// Nii2Header defines the structure of the NIFTI-2 header
type Nii2Header struct {
	SizeofHdr     int32
	Magic         [8]uint8
	Datatype      int16
	Bitpix        int16
	Dim           [8]int64
	IntentP1      float64
	IntentP2      float64
	IntentP3      float64
	Pixdim        [8]float64
	VoxOffset     int64
	SclSlope      float64
	SclInter      float64
	CalMax        float64
	CalMin        float64
	SliceDuration float64
	Toffset       float64
	SliceStart    int64
	SliceEnd      int64
	Descrip       [80]uint8
	AuxFile       [24]uint8
	QformCode     int32
	SformCode     int32
	QuaternB      float64
	QuaternC      float64
	QuaternD      float64
	QoffsetX      float64
	QoffsetY      float64
	QoffsetZ      float64
	SrowX         [4]float64
	SrowY         [4]float64
	SrowZ         [4]float64
	SliceCode     int32
	XyztUnits     int32
	IntentCode    int32
	IntentName    [16]uint8
	DimInfo       uint8
	UnusedStr     [15]uint8
}
