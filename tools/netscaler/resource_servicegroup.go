package netscaler

import (
	"github.com/chiradeep/go-nitro/config/basic"
	"github.com/chiradeep/go-nitro/config/lb"
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceNetScalerServicegroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createServicegroupFunc,
		Read:          readServicegroupFunc,
		Update:        updateServicegroupFunc,
		Delete:        deleteServicegroupFunc,
		Schema: map[string]*schema.Schema{
			"appflowlog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"autoscale": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cacheable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cachetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cipheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cka": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clttimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cmp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"customserverid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"downstateflush": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dupweight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"graceful": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hashid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"healthmonitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"includemembers": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"maxbandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxclient": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxreq": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"memberport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"monitornamesvc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"netprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pathmonitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pathmonitorindv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rtspsessionidremap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"servername": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicegroupname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"svrtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcpb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcpprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"useproxyport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"lbvserver": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createServicegroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In createServicegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	var servicegroupName string
	if v, ok := d.GetOk("servicegroupname"); ok {
		servicegroupName = v.(string)
	} else {
		servicegroupName = resource.PrefixedUniqueId("tf-servicegroup-")
		d.Set("servicegroupname", servicegroupName)
	}
	servicegroup := basic.Servicegroup{
		Appflowlog:         d.Get("appflowlog").(string),
		Autoscale:          d.Get("autoscale").(string),
		Cacheable:          d.Get("cacheable").(string),
		Cachetype:          d.Get("cachetype").(string),
		Cip:                d.Get("cip").(string),
		Cipheader:          d.Get("cipheader").(string),
		Cka:                d.Get("cka").(string),
		Clttimeout:         d.Get("clttimeout").(int),
		Cmp:                d.Get("cmp").(string),
		Comment:            d.Get("comment").(string),
		Customserverid:     d.Get("customserverid").(string),
		Delay:              d.Get("delay").(int),
		Downstateflush:     d.Get("downstateflush").(string),
		Dupweight:          d.Get("dupweight").(int),
		Graceful:           d.Get("graceful").(string),
		Hashid:             d.Get("hashid").(int),
		Healthmonitor:      d.Get("healthmonitor").(string),
		Httpprofilename:    d.Get("httpprofilename").(string),
		Includemembers:     d.Get("includemembers").(bool),
		Maxbandwidth:       d.Get("maxbandwidth").(int),
		Maxclient:          d.Get("maxclient").(int),
		Maxreq:             d.Get("maxreq").(int),
		Memberport:         d.Get("memberport").(int),
		Monitornamesvc:     d.Get("monitornamesvc").(string),
		Monthreshold:       d.Get("monthreshold").(int),
		Netprofile:         d.Get("netprofile").(string),
		Newname:            d.Get("newname").(string),
		Pathmonitor:        d.Get("pathmonitor").(string),
		Pathmonitorindv:    d.Get("pathmonitorindv").(string),
		Port:               d.Get("port").(int),
		Rtspsessionidremap: d.Get("rtspsessionidremap").(string),
		Sc:                 d.Get("sc").(string),
		Serverid:           d.Get("serverid").(int),
		Servername:         d.Get("servername").(string),
		Servicegroupname:   d.Get("servicegroupname").(string),
		Servicetype:        d.Get("servicetype").(string),
		Sp:                 d.Get("sp").(string),
		State:              d.Get("state").(string),
		Svrtimeout:         d.Get("svrtimeout").(int),
		Tcpb:               d.Get("tcpb").(string),
		Tcpprofilename:     d.Get("tcpprofilename").(string),
		Td:                 d.Get("td").(int),
		Useproxyport:       d.Get("useproxyport").(string),
		Usip:               d.Get("usip").(string),
		Weight:             d.Get("weight").(int),
	}

	_, err := client.AddResource(netscaler.Servicegroup.Type(), servicegroupName, &servicegroup)
	if err != nil {
		return err
	}

	d.SetId(servicegroupName)
	lbvserver := d.Get("lbvserver").(string)

	binding := lb.Lbvserverservicegroupbinding{
		Name:             lbvserver,
		Servicegroupname: servicegroupName,
	}

	err = client.BindResource(netscaler.Lbvserver.Type(), lbvserver, netscaler.Servicegroup.Type(), servicegroupName, &binding)
	if err != nil {
		return err
	}
	err = readServicegroupFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this servicegroup but we can't read it ?? %s", servicegroupName)
		return nil
	}
	return nil
}

func readServicegroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In readServicegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	servicegroupName := d.Id()
	log.Printf("[DEBUG] netscaler-provider: Reading servicegroup state %s", servicegroupName)
	data, err := client.FindResource(netscaler.Servicegroup.Type(), servicegroupName)
	if err != nil {
		log.Printf("[WARN] netscaler-provider: Clearing servicegroup state %s", servicegroupName)
		d.SetId("")
		return nil
	}
	d.Set("servicegroupname", data["servicegroupname"])
	d.Set("appflowlog", data["appflowlog"])
	d.Set("autoscale", data["autoscale"])
	d.Set("cacheable", data["cacheable"])
	d.Set("cachetype", data["cachetype"])
	d.Set("cip", data["cip"])
	d.Set("cipheader", data["cipheader"])
	d.Set("cka", data["cka"])
	d.Set("clttimeout", data["clttimeout"])
	d.Set("cmp", data["cmp"])
	d.Set("comment", data["comment"])
	d.Set("customserverid", data["customserverid"])
	d.Set("delay", data["delay"])
	d.Set("downstateflush", data["downstateflush"])
	d.Set("dupweight", data["dupweight"])
	d.Set("graceful", data["graceful"])
	d.Set("hashid", data["hashid"])
	d.Set("healthmonitor", data["healthmonitor"])
	d.Set("httpprofilename", data["httpprofilename"])
	d.Set("includemembers", data["includemembers"])
	d.Set("maxbandwidth", data["maxbandwidth"])
	d.Set("maxclient", data["maxclient"])
	d.Set("maxreq", data["maxreq"])
	d.Set("memberport", data["memberport"])
	d.Set("monitornamesvc", data["monitornamesvc"])
	d.Set("monthreshold", data["monthreshold"])
	d.Set("netprofile", data["netprofile"])
	d.Set("newname", data["newname"])
	d.Set("pathmonitor", data["pathmonitor"])
	d.Set("pathmonitorindv", data["pathmonitorindv"])
	d.Set("port", data["port"])
	d.Set("rtspsessionidremap", data["rtspsessionidremap"])
	d.Set("sc", data["sc"])
	d.Set("serverid", data["serverid"])
	d.Set("servername", data["servername"])
	d.Set("servicegroupname", data["servicegroupname"])
	d.Set("servicetype", data["servicetype"])
	d.Set("sp", data["sp"])
	d.Set("state", data["state"])
	d.Set("svrtimeout", data["svrtimeout"])
	d.Set("tcpb", data["tcpb"])
	d.Set("tcpprofilename", data["tcpprofilename"])
	d.Set("td", data["td"])
	d.Set("useproxyport", data["useproxyport"])
	d.Set("usip", data["usip"])
	d.Set("weight", data["weight"])

	return nil

}

func updateServicegroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In updateServicegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	servicegroupName := d.Get("servicegroupname").(string)

	servicegroup := basic.Servicegroup{
		Servicegroupname: d.Get("servicegroupname").(string),
	}
	hasChange := false
	if d.HasChange("appflowlog") {
		log.Printf("[DEBUG]  netscaler-provider: Appflowlog has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Appflowlog = d.Get("appflowlog").(string)
		hasChange = true
	}
	if d.HasChange("autoscale") {
		log.Printf("[DEBUG]  netscaler-provider: Autoscale has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Autoscale = d.Get("autoscale").(string)
		hasChange = true
	}
	if d.HasChange("cacheable") {
		log.Printf("[DEBUG]  netscaler-provider: Cacheable has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Cacheable = d.Get("cacheable").(string)
		hasChange = true
	}
	if d.HasChange("cachetype") {
		log.Printf("[DEBUG]  netscaler-provider: Cachetype has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Cachetype = d.Get("cachetype").(string)
		hasChange = true
	}
	if d.HasChange("cip") {
		log.Printf("[DEBUG]  netscaler-provider: Cip has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Cip = d.Get("cip").(string)
		hasChange = true
	}
	if d.HasChange("cipheader") {
		log.Printf("[DEBUG]  netscaler-provider: Cipheader has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Cipheader = d.Get("cipheader").(string)
		hasChange = true
	}
	if d.HasChange("cka") {
		log.Printf("[DEBUG]  netscaler-provider: Cka has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Cka = d.Get("cka").(string)
		hasChange = true
	}
	if d.HasChange("clttimeout") {
		log.Printf("[DEBUG]  netscaler-provider: Clttimeout has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Clttimeout = d.Get("clttimeout").(int)
		hasChange = true
	}
	if d.HasChange("cmp") {
		log.Printf("[DEBUG]  netscaler-provider: Cmp has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Cmp = d.Get("cmp").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  netscaler-provider: Comment has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("customserverid") {
		log.Printf("[DEBUG]  netscaler-provider: Customserverid has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Customserverid = d.Get("customserverid").(string)
		hasChange = true
	}
	if d.HasChange("delay") {
		log.Printf("[DEBUG]  netscaler-provider: Delay has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Delay = d.Get("delay").(int)
		hasChange = true
	}
	if d.HasChange("downstateflush") {
		log.Printf("[DEBUG]  netscaler-provider: Downstateflush has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Downstateflush = d.Get("downstateflush").(string)
		hasChange = true
	}
	if d.HasChange("dupweight") {
		log.Printf("[DEBUG]  netscaler-provider: Dupweight has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Dupweight = d.Get("dupweight").(int)
		hasChange = true
	}
	if d.HasChange("graceful") {
		log.Printf("[DEBUG]  netscaler-provider: Graceful has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Graceful = d.Get("graceful").(string)
		hasChange = true
	}
	if d.HasChange("hashid") {
		log.Printf("[DEBUG]  netscaler-provider: Hashid has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Hashid = d.Get("hashid").(int)
		hasChange = true
	}
	if d.HasChange("healthmonitor") {
		log.Printf("[DEBUG]  netscaler-provider: Healthmonitor has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Healthmonitor = d.Get("healthmonitor").(string)
		hasChange = true
	}
	if d.HasChange("httpprofilename") {
		log.Printf("[DEBUG]  netscaler-provider: Httpprofilename has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Httpprofilename = d.Get("httpprofilename").(string)
		hasChange = true
	}
	if d.HasChange("includemembers") {
		log.Printf("[DEBUG]  netscaler-provider: Includemembers has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Includemembers = d.Get("includemembers").(bool)
		hasChange = true
	}
	if d.HasChange("maxbandwidth") {
		log.Printf("[DEBUG]  netscaler-provider: Maxbandwidth has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Maxbandwidth = d.Get("maxbandwidth").(int)
		hasChange = true
	}
	if d.HasChange("maxclient") {
		log.Printf("[DEBUG]  netscaler-provider: Maxclient has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Maxclient = d.Get("maxclient").(int)
		hasChange = true
	}
	if d.HasChange("maxreq") {
		log.Printf("[DEBUG]  netscaler-provider: Maxreq has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Maxreq = d.Get("maxreq").(int)
		hasChange = true
	}
	if d.HasChange("memberport") {
		log.Printf("[DEBUG]  netscaler-provider: Memberport has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Memberport = d.Get("memberport").(int)
		hasChange = true
	}
	if d.HasChange("monitornamesvc") {
		log.Printf("[DEBUG]  netscaler-provider: Monitornamesvc has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Monitornamesvc = d.Get("monitornamesvc").(string)
		hasChange = true
	}
	if d.HasChange("monthreshold") {
		log.Printf("[DEBUG]  netscaler-provider: Monthreshold has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Monthreshold = d.Get("monthreshold").(int)
		hasChange = true
	}
	if d.HasChange("netprofile") {
		log.Printf("[DEBUG]  netscaler-provider: Netprofile has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Netprofile = d.Get("netprofile").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  netscaler-provider: Newname has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("pathmonitor") {
		log.Printf("[DEBUG]  netscaler-provider: Pathmonitor has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Pathmonitor = d.Get("pathmonitor").(string)
		hasChange = true
	}
	if d.HasChange("pathmonitorindv") {
		log.Printf("[DEBUG]  netscaler-provider: Pathmonitorindv has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Pathmonitorindv = d.Get("pathmonitorindv").(string)
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG]  netscaler-provider: Port has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Port = d.Get("port").(int)
		hasChange = true
	}
	if d.HasChange("rtspsessionidremap") {
		log.Printf("[DEBUG]  netscaler-provider: Rtspsessionidremap has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Rtspsessionidremap = d.Get("rtspsessionidremap").(string)
		hasChange = true
	}
	if d.HasChange("sc") {
		log.Printf("[DEBUG]  netscaler-provider: Sc has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Sc = d.Get("sc").(string)
		hasChange = true
	}
	if d.HasChange("serverid") {
		log.Printf("[DEBUG]  netscaler-provider: Serverid has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Serverid = d.Get("serverid").(int)
		hasChange = true
	}
	if d.HasChange("servername") {
		log.Printf("[DEBUG]  netscaler-provider: Servername has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Servername = d.Get("servername").(string)
		hasChange = true
	}
	if d.HasChange("servicegroupname") {
		log.Printf("[DEBUG]  netscaler-provider: Servicegroupname has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Servicegroupname = d.Get("servicegroupname").(string)
		hasChange = true
	}
	if d.HasChange("servicetype") {
		log.Printf("[DEBUG]  netscaler-provider: Servicetype has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Servicetype = d.Get("servicetype").(string)
		hasChange = true
	}
	if d.HasChange("sp") {
		log.Printf("[DEBUG]  netscaler-provider: Sp has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Sp = d.Get("sp").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  netscaler-provider: State has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.State = d.Get("state").(string)
		hasChange = true
	}
	if d.HasChange("svrtimeout") {
		log.Printf("[DEBUG]  netscaler-provider: Svrtimeout has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Svrtimeout = d.Get("svrtimeout").(int)
		hasChange = true
	}
	if d.HasChange("tcpb") {
		log.Printf("[DEBUG]  netscaler-provider: Tcpb has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Tcpb = d.Get("tcpb").(string)
		hasChange = true
	}
	if d.HasChange("tcpprofilename") {
		log.Printf("[DEBUG]  netscaler-provider: Tcpprofilename has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Tcpprofilename = d.Get("tcpprofilename").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  netscaler-provider: Td has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Td = d.Get("td").(int)
		hasChange = true
	}
	if d.HasChange("useproxyport") {
		log.Printf("[DEBUG]  netscaler-provider: Useproxyport has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Useproxyport = d.Get("useproxyport").(string)
		hasChange = true
	}
	if d.HasChange("usip") {
		log.Printf("[DEBUG]  netscaler-provider: Usip has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Usip = d.Get("usip").(string)
		hasChange = true
	}
	if d.HasChange("weight") {
		log.Printf("[DEBUG]  netscaler-provider: Weight has changed for servicegroup %s, starting update", servicegroupName)
		servicegroup.Weight = d.Get("weight").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Servicegroup.Type(), servicegroupName, &servicegroup)
		if err != nil {
			return fmt.Errorf("Error updating servicegroup %s", servicegroupName)
		}
	}
	return readServicegroupFunc(d, meta)
}

func deleteServicegroupFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In deleteServicegroupFunc")
	client := meta.(*NetScalerNitroClient).client
	servicegroupName := d.Id()
	err := client.DeleteResource(netscaler.Servicegroup.Type(), servicegroupName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
