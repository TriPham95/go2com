from pydicom import dcmread

fpath = "./test_data/File 12948.dcm"
# fpath = "./test_data/File 1.dcm"
fpath = "./test_data/File 18001"
fpath = "./test_data/File 32000"
fpath = "./test_data/File 11636.dcm"
fpath = "./test_data/File 32000"
fpath = "./test_data/File 4000.dcm"
fpath = "./test_data/File 12000"
fpath = "./test_data/File 160.dcm"
fpath = "./test_data/File 8000"
# fpath = "./test_data/File 12000"
fpath = '/home/tripg/Documents/1.dcm'
fpath = '/home/tripg/Downloads/mammo_dicoms/1.3.12.2.1107.5.12.7.5054.30000019090500530000000000556.dicom'
fpath = '/home/tripg/Documents/dicom/test_data/400014'
fpath = '/home/tripg/Documents/dicom/mammo_dicoms/1.2.840.113619.2.255.10452022879169.3670200508103440.2701.dicom'
fpath = '/home/tripg/Documents/dicom/test_data/File 12986'
fpath = '/home/tripg/Documents/dicom/ptt1.dcm'
fpath = '/home/tripg/Documents/dicom/test_data/File 10000.dcm'
fpath = '/home/tripg/Documents/dicom/dicoms_mr_func/MR.1.3.46.670589.11.38317.5.0.4476.2014042516042547586'
# fpath = '/home/tripg/Documents/dicom/dicoms_struct/N2D_0001.dcm'
fpath = '/home/tripg/Documents/dicom/test_data/File 10051.dcm'
fpath = '/home/tripg/Documents/dicom/test_data/40009'
fpath = '/home/tripg/Documents/dicom/test_data/File 12943.dcm'
fpath = '/home/tripg/Documents/dicom/mammo_dicoms/1.3.12.2.1107.5.12.7.3367.30000018112001512650000000209.dicom'

ds = dcmread(fpath, force=True)
print(ds)