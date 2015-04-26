package ecs

type DescribeZonesArgs struct {
	RegionId string
}

type AvailableResourceCreationType struct {
	ResourceTypes []string //TODO support enum for Instance, Disk, VSwitch
}

type AvailableDiskCategoriesType struct {
	DiskCategories []string //TODO support enum for cloud, ephemeral, ephemeral_ssd
}

type ZoneType struct {
	ZoneId                    string
	LocalName                 string
	AvailableResourceCreation AvailableResourceCreationType
	AvailableDiskCategories   AvailableDiskCategoriesType
}

type DescribeZonesRespones struct {
	CommonResponse
	Zones struct {
		Zone []ZoneType
	}
}

func (client *Client) DescribeZones(regionId string) (zones []ZoneType, err *ECSError) {
	args := DescribeZonesArgs{
		RegionId: regionId,
	}
	response := DescribeZonesRespones{}

	err = client.Invoke("DescribeZones", &args, &response)

	if err == nil {
		return response.Zones.Zone, nil
	}

	return []ZoneType{}, err
}