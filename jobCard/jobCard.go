package jobCard

const (
    HTTP_MASH_UP = 1 
)
type jobCard struct {
   Jobid []byte `json:"jobid"`
   Jobtype int `json:"jobtype"`
   Ipaddress string `json:"ipaddress"`
   Payload []byte `json:"-"`
}

func init() {
  
}
