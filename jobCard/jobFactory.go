package jobCard 

import (
   "fmt"
   "os/exec"
   "log"
   "bytes"
   "encoding/json"
)
type Export struct {

}
var activeJobs []jobCard

func init() {
 
}

func (c Export) Createjob(jobtype int, ipaddr string) jobCard {
   var tempjc jobCard
   uuid, err := exec.Command("uuidgen").Output()
   tempjc.Jobid = uuid
   fmt.Printf("%s", uuid)
   if err != nil {
      log.Fatal(err)
   }
   tempjc.Jobtype = HTTP_MASH_UP
   tempjc.Ipaddress = "10.10.10.10"
   tempjc.Payload = marshaltojson(tempjc)
   activeJobs = append(activeJobs, tempjc)
   return tempjc
}

func (c Export) Deletejob(jobid []byte) bool {
   //iterate through the job, find job id and then apply delete  
   index, error := findjob(jobid)
   if error != nil {
      //do not do anything
      return false
   } else {
      activeJobs = append(activeJobs[:index], activeJobs[index+1:]...)
      return true      
   } 
}

func (c Export) DumpAllJobs() {
  for il := 0;il <= len(activeJobs) - 1;il++ {
      fmt.Printf("%s", activeJobs[il].Jobid) 
      fmt.Printf("%s", activeJobs[il].Payload)
  }
}


//This is definitely an inefficient way to search, we may have to change it
func findjob(jobid []byte) (int, *string) {
   for il := 0;il <= len(activeJobs) - 1;il++ {
      if bytes.Compare(activeJobs[il].Jobid, jobid) == 0 {
           return il, nil
      }
   }
   var error = new(string)
   *error = "not found"
   return -1, error
}
 
func (c Export) Unmarshaljson(jsonjb []byte) jobCard {
    var tempjb jobCard

    json.Unmarshal(jsonjb, &tempjb)

    fmt.Println(tempjb.Jobid)

    return tempjb

} 

func marshaltojson(job jobCard) []byte {
  var tempjb jobCard
 
  tempjb.Jobid = job.Jobid
  tempjb.Jobtype = job.Jobtype
  tempjb.Ipaddress = job.Ipaddress
 
  jsonval, _ := json.Marshal(tempjb)
  fmt.Println(string(jsonval))
  return jsonval
}
