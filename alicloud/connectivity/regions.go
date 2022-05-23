package connectivity

// Region represents ECS region
type Region string

// Constants of region definition
const (
	Hangzhou    = Region("cn-hangzhou")
	Qingdao     = Region("cn-qingdao")
	Beijing     = Region("cn-beijing")
	Hongkong    = Region("cn-hongkong")
	Shenzhen    = Region("cn-shenzhen")
	Shanghai    = Region("cn-shanghai")
	Zhangjiakou = Region("cn-zhangjiakou")
	Huhehaote   = Region("cn-huhehaote")
	ChengDu     = Region("cn-chengdu")
	HeYuan      = Region("cn-heyuan")
	WuLanChaBu  = Region("cn-wulanchabu")
	GuangZhou   = Region("cn-guangzhou")

	APSouthEast1 = Region("ap-southeast-1")
	APNorthEast1 = Region("ap-northeast-1")
	APSouthEast2 = Region("ap-southeast-2")
	APSouthEast3 = Region("ap-southeast-3")
	APSouthEast5 = Region("ap-southeast-5")
	APSouthEast6 = Region("ap-southeast-6")
	APSouthEast7 = Region("ap-southeast-7")

	APSouth1 = Region("ap-south-1")

	USWest1 = Region("us-west-1")
	USEast1 = Region("us-east-1")

	MEEast1 = Region("me-east-1")

	EUCentral1 = Region("eu-central-1")
	EUWest1    = Region("eu-west-1")

	RusWest1 = Region("rus-west-1")

	ShenZhenFinance     = Region("cn-shenzhen-finance-1")
	ShanghaiFinance     = Region("cn-shanghai-finance-1")
	ShanghaiFinance1Pub = Region("cn-shanghai-finance-1-pub")
	CnNorth2Gov1        = Region("cn-north-2-gov-1")
)

var ValidRegions = []Region{
	Hangzhou, Qingdao, Beijing, Shenzhen, Hongkong, Shanghai, Zhangjiakou, Huhehaote, ChengDu, HeYuan, WuLanChaBu, GuangZhou,
	USWest1, USEast1,
	APNorthEast1, APSouthEast1, APSouthEast2, APSouthEast3, APSouthEast5, APSouthEast6, APSouthEast7,
	APSouth1,
	MEEast1,
	EUCentral1, EUWest1,
	ShenZhenFinance, ShanghaiFinance, CnNorth2Gov1, ShanghaiFinance1Pub,
}

var EcsClassicSupportedRegions = []Region{Hangzhou}
var EcsSpotNoSupportedRegions = []Region{Hangzhou}
var EcsSccSupportedRegions = []Region{Hangzhou}
var SlbGuaranteedSupportedRegions = []Region{Hangzhou}
var DrdsSupportedRegions = []Region{Hangzhou}
var DrdsClassicNoSupportedRegions = []Region{Hangzhou}
var GpdbSupportedRegions = []Region{Hangzhou}

// Some Ram resources only one can be owned by one account at the same time,
// skipped here to avoid multi regions concurrency conflict.
var RamNoSkipRegions = []Region{Hangzhou}
var CenNoSkipRegions = []Region{Hangzhou}
var KmsSkippedRegions = []Region{Hangzhou}

// Actiontrail only one can be owned by one account at the same time,
// skipped here to avoid multi regions concurrency conflict.
var ActiontrailNoSkipRegions = []Region{Hangzhou}
var FcNoSupportedRegions = []Region{Hangzhou}
var DatahubSupportedRegions = []Region{Hangzhou}
var RdsClassicNoSupportedRegions = []Region{Hangzhou}
var RdsMultiAzNoSupportedRegions = []Region{Hangzhou}
var RdsPPASNoSupportedRegions = []Region{Hangzhou}
var RouteTableNoSupportedRegions = []Region{Hangzhou}
var ApiGatewayNoSupportedRegions = []Region{Hangzhou}
var OtsHighPerformanceNoSupportedRegions = []Region{Hangzhou}
var OtsCapacityNoSupportedRegions = []Region{Hangzhou}
var PrivateIpNoSupportedRegions = []Region{Hangzhou}
var SwarmSupportedRegions = []Region{Hangzhou}
var ManagedKubernetesSupportedRegions = []Region{Hangzhou}
var ServerlessKubernetesSupportedRegions = []Region{Hangzhou}
var KubernetesSupportedRegions = []Region{Hangzhou}
var NasClassicSupportedRegions = []Region{Hangzhou}
var CasClassicSupportedRegions = []Region{Hangzhou}
var CRNoSupportedRegions = []Region{Hangzhou}
var MongoDBClassicNoSupportedRegions = []Region{Hangzhou}
var MongoDBMultiAzSupportedRegions = []Region{Hangzhou}
var DdoscooSupportedRegions = []Region{Hangzhou}
var DdosbgpSupportedRegions = []Region{Hangzhou}

//var NetworkAclSupportedRegions = []Region{Hangzhou}
var EssScalingConfigurationMultiSgSupportedRegions = []Region{Hangzhou}
var SlbClassicNoSupportedRegions = []Region{Hangzhou}
var NasNoSupportedRegions = []Region{Hangzhou}
var OssVersioningSupportedRegions = []Region{Hangzhou}
var OssSseSupportedRegions = []Region{Hangzhou}
var GpdbClassicNoSupportedRegions = []Region{Hangzhou}
var OnsNoSupportRegions = []Region{Hangzhou}
var AlikafkaSupportedRegions = []Region{Hangzhou}
var SmartagSupportedRegions = []Region{Hangzhou}
var YundunDbauditSupportedRegions = []Region{Hangzhou}
var HttpHttpsHealthCheckMehtodSupportedRegions = []Region{Hangzhou}
var HBaseClassicSupportedRegions = []Region{Hangzhou}
var EdasSupportedRegions = []Region{Hangzhou}
var CloudConfigSupportedRegions = []Region{Hangzhou}
var DBReadwriteSplittingConnectionSupportedRegions = []Region{Hangzhou}
var KVstoreClassicNetworkInstanceSupportRegions = []Region{Hangzhou}
var MaxComputeSupportRegions = []Region{Hangzhou}
var FnfSupportRegions = []Region{Hangzhou}
var PrivateLinkRegions = []Region{Hangzhou}
var BrainIndustrialRegions = []Region{Hangzhou}
var EciContainerGroupRegions = []Region{Hangzhou}
var TsdbInstanceSupportRegions = []Region{Hangzhou}
var VpcIpv6SupportRegions = []Region{Hangzhou}
var EssdSupportRegions = []Region{Hangzhou}
var AdbReserverUnSupportRegions = []Region{Hangzhou}
var KmsKeyHSMSupportRegions = []Region{Hangzhou}
var DmSupportRegions = []Region{Hangzhou}
var BssOpenApiSupportRegions = []Region{Hangzhou}
var EipAddressBGPProSupportRegions = []Region{Hangzhou}
var CenTransitRouterVpcAttachmentSupportRegions = []Region{Hangzhou} // Not all of APSouthEast1 and HangZhou zones support vpc attachment
var ARMSSupportRegions = []Region{Hangzhou}
var SaeSupportRegions = []Region{Hangzhou}
var HbrSupportRegions = []Region{Hangzhou}
var EcdSupportRegions = []Region{Hangzhou}
var EcpSupportRegions = []Region{Hangzhou}
var SddpSupportRegions = []Region{Hangzhou}
var DfsSupportRegions = []Region{Hangzhou}
var EventBridgeSupportRegions = []Region{Hangzhou}
var AlbSupportRegions = []Region{Hangzhou}
var IMMSupportRegions = []Region{Hangzhou}
var CenTRSupportRegions = []Region{Hangzhou}
var VbrSupportRegions = []Region{Hangzhou}
var ClickHouseSupportRegions = []Region{Hangzhou}
var ClickHouseBackupPolicySupportRegions = []Region{Hangzhou}
var DatabaseGatewaySupportRegions = []Region{Hangzhou}
var CloudSsoSupportRegions = []Region{Hangzhou}
var SWASSupportRegions = []Region{Hangzhou}
var SurveillanceSystemSupportRegions = []Region{Hangzhou}
var VodSupportRegions = []Region{Hangzhou}
var OpenSearchSupportRegions = []Region{Hangzhou}
var GraphDatabaseSupportRegions = []Region{Hangzhou}
var DBFSSystemSupportRegions = []Region{Hangzhou}
var EAISSystemSupportRegions = []Region{Hangzhou}
var CloudAuthSupportRegions = []Region{Hangzhou}
var MHUBSupportRegions = []Region{Hangzhou}
var ActiontrailSupportRegions = []Region{Hangzhou}
var VpcTrafficMirrorSupportRegions = []Region{Hangzhou}
var EcdUserSupportRegions = []Region{Hangzhou}
var VpcIpv6GatewaySupportRegions = []Region{Hangzhou}
var CmsDynamicTagGroupSupportRegions = []Region{Hangzhou}
var OOSApplicationSupportRegions = []Region{Hangzhou}
var DTSSupportRegions = []Region{Hangzhou}
var OOSSupportRegions = []Region{Hangzhou}
var MongoDBSupportRegions = []Region{Hangzhou}
var MongoDBServerlessSupportRegions = []Region{Hangzhou}
var FnFSupportRegions = []Region{Hangzhou}
var GaSupportRegions = []Region{Hangzhou}
var AlidnsSupportRegions = []Region{Hangzhou}
var VPCVbrHaSupportRegions = []Region{Hangzhou}
var ROSSupportRegions = []Region{Hangzhou}
var VPCBgpGroupSupportRegions = []Region{Hangzhou}
var NASSupportRegions = []Region{Hangzhou}
var HBRSupportRegions = []Region{Hangzhou}
var NASCPFSSupportRegions = []Region{Hangzhou}
var WAFSupportRegions = []Region{Hangzhou}
var MSCSupportRegions = []Region{Hangzhou}

// Other regions requires the custom should have icp
var FCCustomDomainSupportRegions = []Region{Hangzhou}
var RDCupportRegions = []Region{Hangzhou}
var MSEGatewaySupportRegions = []Region{Hangzhou}
var BrainIndustrialSupportRegions = []Region{Hangzhou}
var TestSalveRegions = []Region{Hangzhou}
var TestPvtzRegions = []Region{Hangzhou}
var ECPSupportRegions = []Region{Hangzhou}
var DCDNSupportRegions = []Region{Hangzhou}
var GpdbElasticInstanceSupportRegions = []Region{Hangzhou}
var PolarDBSupportRegions = []Region{Hangzhou}
var ESSSupportRegions = []Region{Hangzhou}
var SimpleApplicationServerNotSupportRegions = []Region{Hangzhou}
var CRSupportRegions = []Region{Hangzhou}
var MSESupportRegions = []Region{Hangzhou}
var LogResourceSupportRegions = []Region{Hangzhou}
var AliKafkaSupportRegions = []Region{Hangzhou}
var BastionhostSupportRegions = []Region{Hangzhou}
var ACKSystemDiskEncryptionSupportRegions = []Region{Hangzhou}
