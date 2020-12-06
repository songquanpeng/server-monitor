package main

type Server struct {
	Ip      string
	GPUInfo GPUInfo
}

type GPUInfo struct {
	Timestamp     string `xml:"timestamp"`
	DriverVersion string `xml:"driver_version"`
	AttachedGpus  int    `xml:"attached_gpus"`
	Gpu           []struct {
		ProductName string `xml:"product_name"`
		Utilization struct {
			GpuUtil     string `xml:"gpu_util"`
			MemoryUtil  string `xml:"memory_util"`
			EncoderUtil string `xml:"encoder_util"`
			DecoderUtil string `xml:"decoder_util"`
		} `xml:"utilization"`
	} `xml:"gpu"`
}
