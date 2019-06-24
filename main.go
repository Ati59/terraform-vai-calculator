// Update the variables below from the output of your terraform apply, assuming the attachment name is "attachment"

// output "attachment_name" {
//  value = "${aws_volume_attachment.attachment.device_name}"
// }

// output "attachment_instance_id" {
//   value = "${aws_volume_attachment.attachment.instance_id}"
// }

// output "attachment_vol_id" {
//   value = "${aws_volume_attachment.attachment.volume_id}"
// }

package main

import (
	"os"
	"bytes"
	"fmt"
	"github.com/hashicorp/terraform/helper/hashcode"
)


func volumeAttachmentID(name, volumeID, instanceID string) string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s-", name))
	buf.WriteString(fmt.Sprintf("%s-", instanceID))
	buf.WriteString(fmt.Sprintf("%s-", volumeID))

	return fmt.Sprintf("Computed Volume Attachment ID : vai-%d", hashcode.String(buf.String()))
}

func main() {
	if os.Args != nil && len(os.Args) == 4 {
	  name := os.Args[1]
	  vID := os.Args[2]
	  iID := os.Args[3]

	  fmt.Println("Name :", name)
	  fmt.Println("VolumeId :", vID)
	  fmt.Println("InstanceId :", iID)
	  fmt.Println(volumeAttachmentID(name, vID, iID))
	} else {
	  fmt.Println("*** ERROR *** Usage : ./terraform-vai-calculator [name] [VolumeId] [InstanceId]")
	}
}

