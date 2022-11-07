package go2com

import (
	"fmt"
	_ "image/jpeg"
	"testing"

	"github.com/okieraised/go2com/pkg/nifti/reader"
	"github.com/stretchr/testify/assert"
)

func TestNii1(t *testing.T) {
	assert := assert.New(t)

	filePath := "/home/tripg/Documents/nifti/Arnow^Corie^Shelvey^OM_segmented.nii"
	filePath = "/home/tripg/Documents/nifti/RGB16_4D.nii.gz"
	filePath = "/home/tripg/Documents/nifti/someones_anatomy.nii.gz"
	filePath = "/home/tripg/Documents/nifti/someones_epi.nii.gz"
	filePath = "/home/tripg/Documents/nifti/RGB8_4D.nii.gz"
	filePath = "/home/tripg/Documents/nifti/jaw.nii.gz"
	filePath = "/home/tripg/Documents/nifti/Arnow^Corie^Shelvey^OM_segmented.nii"
	filePath = "/home/tripg/Documents/nifti/knee.nii.gz"
	//filePath = "/home/tripg/Documents/nifti/ExBox11/fmri.nii.gz"
	filePath = "/home/tripg/Documents/nifti/ExBox11/structural.nii.gz"
	//filePath = "/home/tripg/Documents/nifti/ExBox11/structural_brain.nii.gz"
	//filePath = "/home/tripg/Documents/nifti/JHU_MNI_SS_T1.nii.gz"
	//filePath = "/home/tripg/Documents/nifti/avg152T1_LR_nifti2.nii.gz"
	filePath = "/Users/TriPham/Downloads/avg152T1_RL_nifti.nii"

	niiReader, err := reader.NewNiiReader(filePath)
	assert.NoError(err)
	err = niiReader.Parse()
	assert.NoError(err)

	shape := niiReader.GetImgShape()
	fmt.Println("shape", shape)
	fmt.Println(niiReader.GetUnitsOfMeasurements())
	fmt.Println("dtype", niiReader.GetDatatype())
	fmt.Println("nbyper", niiReader.GetNiiData().Data.NByPer)
	fmt.Println("GetSFromCode()", niiReader.GetSFormCode())
	fmt.Println("GetQFromCode()", niiReader.GetQFormCode())
	fmt.Println(niiReader.GetOrientation())

	res, err := niiReader.GetSlice(1, 0)
	assert.NoError(err)

	fmt.Println("len res", len(res))
	fmt.Println("len res0", len(res[0]))

	//fmt.Println(res)
	res2, err := niiReader.GetVolume(0)
	assert.NoError(err)

	fmt.Println("len res2", len(res2))

	//for _, elem := range res {
	//	fmt.Println(elem)
	//
	//}

	//fmt.Println(niiReader.GetTimeSeries(26, 30, 16))

	//for x := 0; x < 57; x++ {
	//	for y := 0; y < 67; y++ {
	//		for z := 0; z < 56; z++ {
	//			fmt.Println(niiReader.GetTimeSeries(int64(x), int64(y), int64(z)))
	//		}
	//	}
	//}

	shape = niiReader.GetImgShape()
	fmt.Println(shape)
}
