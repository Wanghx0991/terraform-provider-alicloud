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
	NanJing     = Region("cn-nanjing")

	APSouthEast1 = Region("ap-southeast-1")
	APNorthEast1 = Region("ap-northeast-1")
	APNorthEast2 = Region("ap-northeast-2")
	APSouthEast2 = Region("ap-southeast-2")
	APSouthEast3 = Region("ap-southeast-3")
	APSouthEast5 = Region("ap-southeast-5")
	APSouthEast6 = Region("ap-southeast-6")
	APSouthEast7 = Region("ap-southeast-7")

	APSouth1 = Region("ap-south-1")

	USWest1 = Region("us-west-1")
	USEast1 = Region("us-east-1")

	MEEast1    = Region("me-east-1")
	MECentral1 = Region("me-central-1")

	EUCentral1 = Region("eu-central-1")
	EUWest1    = Region("eu-west-1")

	RusWest1 = Region("rus-west-1")

	HangzhouFinance     = Region("cn-hangzhou-finance")
	HangzhouFinanceOSS  = Region("cn-hzfinance")
	HangzhouFinanceOSS1 = Region("cn-hzjbp")

	BeijingFinance1   = Region("cn-beijing-finance-1")
	BeijingFinancePub = Region("cn-beijing-finance-1-pub")

	ShanghaiFinance     = Region("cn-shanghai-finance-1")
	ShanghaiFinance1Pub = Region("cn-shanghai-finance-1-pub")
	ShenZhenFinance1    = Region("cn-shenzhen-finance-1")
	ShenzhenFinance2    = Region("cn-szfinance")
	ShenzhenFinance     = Region("cn-shenzhen-finance")

	CnNorth2Gov1 = Region("cn-north-2-gov-1")
)

var ValidRegions = []Region{
	Hangzhou, Qingdao, Beijing, Shenzhen, Hongkong, Shanghai, Zhangjiakou, Huhehaote, ChengDu, HeYuan, WuLanChaBu, GuangZhou, NanJing,
	USWest1, USEast1,
	APNorthEast1, APNorthEast2, APSouthEast1, APSouthEast2, APSouthEast3, APSouthEast5, APSouthEast6, APSouthEast7,
	APSouth1,
	MEEast1, MECentral1,
	EUCentral1, EUWest1,
	CnNorth2Gov1,
	ShenZhenFinance1, ShenzhenFinance2, ShenzhenFinance,
	ShanghaiFinance1Pub, ShanghaiFinance,
	HangzhouFinance, HangzhouFinanceOSS, HangzhouFinanceOSS1,
	BeijingFinance1, BeijingFinancePub,
}

var EcsClassicSupportedRegions = []Region{Shenzhen, Shanghai, Beijing, Qingdao, Hangzhou, Hongkong, USWest1, APSouthEast1, ShanghaiFinance}
var EcsSpotNoSupportedRegions = []Region{APSouth1, ShanghaiFinance}
var EcsSccSupportedRegions = []Region{Shanghai, Beijing, Zhangjiakou, ShanghaiFinance}
var SlbGuaranteedSupportedRegions = []Region{Qingdao, Beijing, Hangzhou, Shanghai, Shenzhen, Zhangjiakou, Huhehaote, APSouthEast1, USEast1, ShanghaiFinance}
var DrdsSupportedRegions = []Region{Beijing, Shenzhen, Hangzhou, Qingdao, Hongkong, Shanghai, Huhehaote, Zhangjiakou, APSouthEast1, ShanghaiFinance}
var DrdsClassicNoSupportedRegions = []Region{Hongkong, ShanghaiFinance}
var GpdbSupportedRegions = []Region{Beijing, Shenzhen, Hangzhou, Shanghai, Hongkong, ShanghaiFinance}

// Some Ram resources only one can be owned by one account at the same time,
// skipped here to avoid multi regions concurrency conflict.
var RamNoSkipRegions = []Region{Hangzhou, EUCentral1, APSouth1, ShanghaiFinance}
var CenNoSkipRegions = []Region{Shanghai, EUCentral1, APSouth1, ShanghaiFinance}
var KmsSkippedRegions = []Region{Beijing, Shanghai}

// Actiontrail only one can be owned by one account at the same time,
// skipped here to avoid multi regions concurrency conflict.
var ActiontrailNoSkipRegions = []Region{Hangzhou, EUCentral1, APSouth1, ShanghaiFinance}
var FcNoSupportedRegions = []Region{MEEast1, ShanghaiFinance}
var DatahubSupportedRegions = []Region{Beijing, Hangzhou, Shanghai, Shenzhen, APSouthEast1, ShanghaiFinance}
var RdsClassicNoSupportedRegions = []Region{APSouth1, APSouthEast2, APSouthEast3, APNorthEast1, EUCentral1, EUWest1, MEEast1, ShanghaiFinance}
var RdsMultiAzNoSupportedRegions = []Region{Qingdao, APNorthEast1, APSouthEast5, MEEast1, ShanghaiFinance}
var RdsPPASNoSupportedRegions = []Region{Qingdao, USEast1, APNorthEast1, EUCentral1, MEEast1, APSouthEast2, APSouthEast3, APSouth1, APSouthEast5, ChengDu, EUWest1, ShanghaiFinance}
var RouteTableNoSupportedRegions = []Region{Beijing, Hangzhou, Shenzhen, ShanghaiFinance}
var ApiGatewayNoSupportedRegions = []Region{Zhangjiakou, Huhehaote, USEast1, USWest1, EUWest1, MEEast1, ShanghaiFinance}
var OtsHighPerformanceNoSupportedRegions = []Region{Qingdao, Zhangjiakou, Huhehaote, APSouthEast2, APSouthEast5, APNorthEast1, EUCentral1, MEEast1, ShanghaiFinance}
var OtsCapacityNoSupportedRegions = []Region{APSouthEast1, USWest1, USEast1, ShanghaiFinance}
var PrivateIpNoSupportedRegions = []Region{Beijing, Hangzhou, Shenzhen, ShanghaiFinance}
var SwarmSupportedRegions = []Region{Qingdao, Beijing, Zhangjiakou, Huhehaote, Hangzhou, Shanghai, Shenzhen, Hongkong, APNorthEast1, APSouthEast1, APSouthEast2,
	APSouthEast3, USWest1, USEast1, EUCentral1, ShanghaiFinance}
var ManagedKubernetesSupportedRegions = []Region{Beijing, Hangzhou, Shanghai, Shenzhen, ChengDu, Hongkong, APSouthEast1, APSouthEast2, EUCentral1, USWest1, ShanghaiFinance}
var ServerlessKubernetesSupportedRegions = []Region{Beijing, Hangzhou, Shanghai, APSouthEast1, APSouthEast3, APSouthEast5, APSouth1, Huhehaote, ShanghaiFinance}
var KubernetesSupportedRegions = []Region{Beijing, Zhangjiakou, Huhehaote, Hangzhou, Shanghai, Shenzhen, Hongkong, APNorthEast1, APSouthEast1,
	APSouthEast2, APSouthEast3, APSouthEast5, APSouth1, USEast1, USWest1, EUWest1, MEEast1, EUCentral1, ShanghaiFinance}
var NasClassicSupportedRegions = []Region{ShanghaiFinance, Hangzhou, Qingdao, Beijing, Hongkong, Shenzhen, Shanghai, Zhangjiakou, Huhehaote, ShenZhenFinance1}
var CasClassicSupportedRegions = []Region{Hangzhou, APSouth1, MEEast1, EUCentral1, APNorthEast1, APSouthEast2, ShanghaiFinance}
var CRNoSupportedRegions = []Region{Beijing, Hangzhou, Qingdao, Huhehaote, Zhangjiakou, ShanghaiFinance}
var MongoDBClassicNoSupportedRegions = []Region{Huhehaote, Zhangjiakou, APSouthEast2, APSouthEast3, ShanghaiFinance, APSouthEast5, APSouth1, USEast1, USWest1, APNorthEast1}
var MongoDBMultiAzSupportedRegions = []Region{Hangzhou, Beijing, Shenzhen, EUCentral1, ShanghaiFinance}
var DdoscooSupportedRegions = []Region{Hangzhou, APSouthEast1, ShanghaiFinance}
var DdosbgpSupportedRegions = []Region{Hangzhou, Beijing, Shenzhen, Qingdao, ShanghaiFinance, Shanghai, Zhangjiakou, Huhehaote}

//var NetworkAclSupportedRegions = []Region{Hangzhou, Beijing, Shanghai, Hongkong, APSouthEast5, APSouth1}
var EssScalingConfigurationMultiSgSupportedRegions = []Region{APSouthEast1, APSouth1, ShanghaiFinance}
var SlbClassicNoSupportedRegions = []Region{APNorthEast1, APSouthEast2, APSouthEast3, APSouthEast5, APSouth1, USEast1, MEEast1, EUCentral1, EUWest1, Huhehaote, Zhangjiakou, ShanghaiFinance}
var NasNoSupportedRegions = []Region{Qingdao, APSouth1, APSouthEast3, APSouthEast5, ShanghaiFinance}
var OssVersioningSupportedRegions = []Region{APSouth1, ShanghaiFinance}
var OssSseSupportedRegions = []Region{Qingdao, Hangzhou, Beijing, Shanghai, Shenzhen, Hongkong, APNorthEast1, APSouth1, USEast1, ShanghaiFinance}
var GpdbClassicNoSupportedRegions = []Region{APSouthEast2, APSouthEast3, APSouthEast5, APSouth1, USEast1, USWest1, APNorthEast1, EUCentral1, ShanghaiFinance}
var OnsNoSupportRegions = []Region{APSouthEast5, ShanghaiFinance}
var AlikafkaSupportedRegions = []Region{Hangzhou, Qingdao, Beijing, Hongkong, Shenzhen, Shanghai, Zhangjiakou, Huhehaote, ChengDu, HeYuan, APNorthEast1, APSouthEast1, APSouthEast3, EUCentral1, EUWest1, USEast1, USWest1, ShanghaiFinance}
var SmartagSupportedRegions = []Region{Shanghai, ShanghaiFinance, Hongkong, APSouthEast1, APSouthEast2, APSouthEast3, APSouthEast5, EUCentral1, APNorthEast1, ShanghaiFinance}
var YundunDbauditSupportedRegions = []Region{Hangzhou, Beijing, Shanghai, ShanghaiFinance}
var HttpHttpsHealthCheckMehtodSupportedRegions = []Region{Hangzhou, Beijing, Shanghai, EUWest1, ChengDu, Qingdao, Hongkong, Shenzhen, APSouthEast5, Zhangjiakou, Huhehaote, MEEast1, APSouth1, EUCentral1, USWest1, APSouthEast3, APSouthEast2, APSouthEast1, APNorthEast1, ShanghaiFinance}
var HBaseClassicSupportedRegions = []Region{Hangzhou, Beijing, Shanghai, Shenzhen, ShanghaiFinance}
var EdasSupportedRegions = []Region{Hangzhou, Beijing, Shanghai, Shenzhen, Zhangjiakou, Qingdao, Hongkong, ShanghaiFinance}
var CloudConfigSupportedRegions = []Region{Shanghai, ShanghaiFinance}
var DBReadwriteSplittingConnectionSupportedRegions = []Region{APSouthEast1, ShanghaiFinance}
var KVstoreClassicNetworkInstanceSupportRegions = []Region{Hangzhou, Beijing, Shanghai, APSouthEast1, USEast1, USWest1, ShanghaiFinance}
var MaxComputeSupportRegions = []Region{ShanghaiFinance}
var FnfSupportRegions = []Region{Hangzhou, Beijing, Shanghai, Shenzhen, USWest1, ShanghaiFinance}
var PrivateLinkRegions = []Region{EUCentral1, ShanghaiFinance}
var BrainIndustrialRegions = []Region{Hangzhou, ShanghaiFinance}
var EciContainerGroupRegions = []Region{Hangzhou, ShanghaiFinance}
var TsdbInstanceSupportRegions = []Region{Beijing, Hangzhou, Shenzhen, Shanghai, ShenZhenFinance1, Qingdao, Zhangjiakou, ShanghaiFinance, Hongkong, USWest1, APNorthEast1, EUWest1, APSouthEast1, APSouthEast2, APSouthEast3, EUCentral1, APSouthEast5, Zhangjiakou, CnNorth2Gov1, ShanghaiFinance}
var VpcIpv6SupportRegions = []Region{Hangzhou, Shanghai, Shenzhen, Beijing, Huhehaote, Hongkong, APSouthEast1, ShanghaiFinance, ShanghaiFinance}
var EssdSupportRegions = []Region{Zhangjiakou, Huhehaote, ShanghaiFinance}
var AdbReserverUnSupportRegions = []Region{EUCentral1, ShanghaiFinance}
var KmsKeyHSMSupportRegions = []Region{Beijing, Zhangjiakou, Hangzhou, Shanghai, Shenzhen, Hongkong, APSouthEast1, APSouthEast2, APSouthEast3, APSouthEast5, USEast1, ShanghaiFinance}
var DmSupportRegions = []Region{Hangzhou, ShanghaiFinance}
var BssOpenApiSupportRegions = []Region{Hangzhou, Shanghai, APSouthEast1, ShanghaiFinance}
var EipAddressBGPProSupportRegions = []Region{Hongkong, ShanghaiFinance}
var CenTransitRouterVpcAttachmentSupportRegions = []Region{EUCentral1, ShanghaiFinance} // Not all of APSouthEast1 and HangZhou zones support vpc attachment
var ARMSSupportRegions = []Region{Hangzhou, Shanghai, Beijing, APSouthEast1, ShanghaiFinance}
var SaeSupportRegions = []Region{Hangzhou, Shanghai, Beijing, Zhangjiakou, Shenzhen, USWest1, ShanghaiFinance}
var HbrSupportRegions = []Region{Hangzhou, ShanghaiFinance}
var EcdSupportRegions = []Region{Hangzhou, Shanghai, Beijing, Shenzhen, Hongkong, APSouthEast1, APSouthEast2, ShanghaiFinance}
var EcpSupportRegions = []Region{Hangzhou, Shanghai, Beijing, Shenzhen, ShanghaiFinance}
var SddpSupportRegions = []Region{Hangzhou, Zhangjiakou, APSouthEast1, ShanghaiFinance}
var DfsSupportRegions = []Region{Hangzhou, Zhangjiakou, Shanghai, Beijing, HeYuan, ChengDu, APSouthEast5, USEast1, RusWest1, ShanghaiFinance}
var EventBridgeSupportRegions = []Region{Hangzhou, Zhangjiakou, Shanghai, Shenzhen, Beijing, HeYuan, ChengDu, Huhehaote, Hongkong, EUCentral1, USWest1, USEast1, ShanghaiFinance}
var AlbSupportRegions = []Region{Hangzhou, Shanghai, Qingdao, Zhangjiakou, Beijing, WuLanChaBu, Shenzhen, ChengDu, Hongkong, APSouthEast1, APSouthEast2, APSouthEast3, APSouthEast5, APNorthEast1, EUCentral1, USEast1, APSouth1, ShanghaiFinance}
var IMMSupportRegions = []Region{Hangzhou, Zhangjiakou, APSouthEast1, Shenzhen, Beijing, Shanghai, ShanghaiFinance}
var CenTRSupportRegions = []Region{EUCentral1, APSouthEast1, Hangzhou, Shanghai, Beijing, Shenzhen, Hongkong, APSouthEast1, USEast1, APSouth1, ShanghaiFinance}
var VbrSupportRegions = []Region{Hangzhou, ShanghaiFinance}
var ClickHouseSupportRegions = []Region{Hangzhou, Qingdao, Beijing, Shenzhen, Hongkong, Shanghai, Zhangjiakou, Huhehaote, ChengDu, USWest1, USEast1, APSouthEast1, EUCentral1, EUWest1, APNorthEast1, APSouthEast1, APSouthEast5, ShanghaiFinance}
var ClickHouseBackupPolicySupportRegions = []Region{Shanghai, ShanghaiFinance}
var DatabaseGatewaySupportRegions = []Region{Hangzhou, Zhangjiakou, Shanghai, Beijing, Qingdao, Huhehaote, Shenzhen, ChengDu, Hongkong, APNorthEast1, APSouth1, APSouthEast1, APSouthEast2, APSouthEast3, EUWest1, EUCentral1, APSouthEast5, USWest1, USEast1, ShanghaiFinance}
var CloudSsoSupportRegions = []Region{Shanghai, USWest1, ShanghaiFinance}
var SWASSupportRegions = []Region{Qingdao, Hangzhou, Beijing, Shenzhen, Shanghai, GuangZhou, Huhehaote, ChengDu, Zhangjiakou, Hongkong, APSouthEast1, ShanghaiFinance}
var SurveillanceSystemSupportRegions = []Region{Beijing, Shenzhen, Qingdao, ShanghaiFinance}
var VodSupportRegions = []Region{Shanghai, ShanghaiFinance}
var OpenSearchSupportRegions = []Region{Beijing, Shenzhen, Hangzhou, Zhangjiakou, Qingdao, Shanghai, APSouthEast1, ShanghaiFinance}
var GraphDatabaseSupportRegions = []Region{Shenzhen, Beijing, Qingdao, Shanghai, Hongkong, Zhangjiakou, Hangzhou, APSouthEast1, APSouthEast5, USWest1, USEast1, APSouth1, ShanghaiFinance}
var DBFSSystemSupportRegions = []Region{Hangzhou, ShanghaiFinance}
var EAISSystemSupportRegions = []Region{Hangzhou, ShanghaiFinance}
var CloudAuthSupportRegions = []Region{Hangzhou, ShanghaiFinance}
var MHUBSupportRegions = []Region{Shanghai, ShanghaiFinance}
var ActiontrailSupportRegions = []Region{Hangzhou, Qingdao, Beijing, Shenzhen, Hongkong, Shanghai, Zhangjiakou, Huhehaote, ChengDu, HeYuan, WuLanChaBu, GuangZhou, APSouthEast1, APSouthEast2, APSouthEast3, APSouthEast5, APNorthEast1, USWest1, USEast1, EUCentral1, EUWest1, APSouth1, MEEast1, ShanghaiFinance}
var VpcTrafficMirrorSupportRegions = []Region{Hangzhou, Beijing, Zhangjiakou, Qingdao, Huhehaote, Shenzhen, Hongkong, APSouthEast2, ChengDu, USEast1, USWest1, EUWest1, ShanghaiFinance}
var EcdUserSupportRegions = []Region{Shanghai, ShanghaiFinance}
var VpcIpv6GatewaySupportRegions = []Region{Qingdao, Beijing, Zhangjiakou, Huhehaote, WuLanChaBu, Hangzhou, Shanghai, Shenzhen, GuangZhou, Hongkong, ChengDu, HeYuan, APSouthEast1, APSouthEast6, USEast1, EUCentral1, ShanghaiFinance}
var CmsDynamicTagGroupSupportRegions = []Region{Shanghai, ShanghaiFinance}
var OOSApplicationSupportRegions = []Region{Hangzhou, ShanghaiFinance}
var DTSSupportRegions = []Region{Hangzhou, APSouth1, ShenZhenFinance1, CnNorth2Gov1, Qingdao, ShanghaiFinance, USWest1, APNorthEast1, Beijing, Hongkong, APSouthEast1, APSouthEast3, EUCentral1, APSouthEast5, Shenzhen, APSouthEast2, Huhehaote, USEast1, Zhangjiakou, EUWest1, MEEast1, Shanghai, ShanghaiFinance}
var OOSSupportRegions = []Region{APSouthEast5, USWest1, EUWest1, Qingdao, ChengDu, Shanghai, Huhehaote, Shenzhen, APNorthEast1, APSouthEast1, EUCentral1, Hangzhou, Beijing, APSouth1, APSouthEast3, USEast1, Zhangjiakou, Hongkong, APSouthEast2, ShanghaiFinance}
var MongoDBSupportRegions = []Region{APSouth1, Shanghai, APSouthEast2, WuLanChaBu, CnNorth2Gov1, Hangzhou, Beijing, Qingdao, Zhangjiakou, USWest1, GuangZhou, APSouthEast6, EUWest1, ChengDu, APSouthEast1, APSouthEast3, APSouthEast5, ShanghaiFinance, Hongkong, HeYuan, Huhehaote, USEast1, EUCentral1, APNorthEast1, Shenzhen, ShenZhenFinance1, MEEast1, ShanghaiFinance}
var MongoDBServerlessSupportRegions = []Region{APSouthEast5, Shanghai, USEast1, Hongkong, HeYuan, Zhangjiakou, APSouthEast6, GuangZhou, Huhehaote, Beijing, Shenzhen, WuLanChaBu, ChengDu, Hangzhou, Qingdao, USWest1, APSouthEast1, ShanghaiFinance}
var FnFSupportRegions = []Region{Shenzhen, Beijing, Shanghai, APSouthEast1, USWest1, Hangzhou, ShanghaiFinance}
var GaSupportRegions = []Region{Hangzhou, ShanghaiFinance}
var AlidnsSupportRegions = []Region{Hangzhou, APSouthEast1, ShanghaiFinance}
var VPCVbrHaSupportRegions = []Region{Hangzhou, ShanghaiFinance}
var ROSSupportRegions = []Region{USWest1, HeYuan, Zhangjiakou, Hongkong, APSouthEast3, EUCentral1, Huhehaote, APSouthEast6, Shenzhen, APSouth1, Qingdao, GuangZhou, APSouthEast2, WuLanChaBu, EUWest1, MEEast1, ChengDu, Shanghai, APSouthEast1, APSouthEast5, USEast1, Beijing, APNorthEast1, Hangzhou, ShanghaiFinance}
var VPCBgpGroupSupportRegions = []Region{Hangzhou, ShanghaiFinance}
var NASSupportRegions = []Region{HeYuan, Huhehaote, APSouthEast5, WuLanChaBu, CnNorth2Gov1, Qingdao, ChengDu, Hangzhou, APSouth1, ShenZhenFinance1, EUCentral1, Shenzhen, APSouthEast2, Beijing, Shanghai, ShanghaiFinance, APSouthEast1, APSouthEast6, APNorthEast1, APSouthEast3, GuangZhou, USEast1, EUWest1, Hongkong, Zhangjiakou, USWest1, ShanghaiFinance}
var HBRSupportRegions = []Region{Beijing, ChengDu, Huhehaote, Qingdao, Shanghai, Shenzhen, Zhangjiakou, Hangzhou, ShanghaiFinance}
var NASCPFSSupportRegions = []Region{Hangzhou, Shenzhen, Beijing, Shanghai, HeYuan, Huhehaote, WuLanChaBu, Qingdao, ChengDu, ShanghaiFinance}
var WAFSupportRegions = []Region{Hangzhou, APSouth1, ShanghaiFinance}
var MSCSupportRegions = []Region{Hangzhou, ShanghaiFinance}

// Other regions requires the custom should have icp
var FCCustomDomainSupportRegions = []Region{EUCentral1, APSouthEast1, ShanghaiFinance}
var RDCupportRegions = []Region{Shanghai, ShanghaiFinance}
var MSEGatewaySupportRegions = []Region{Shenzhen, Hangzhou, Shanghai, Beijing, ShanghaiFinance}
var BrainIndustrialSupportRegions = []Region{Hangzhou, ShanghaiFinance}
var TestSalveRegions = []Region{Hangzhou, ShanghaiFinance}
var TestPvtzRegions = []Region{Hangzhou, ShanghaiFinance}
var ECPSupportRegions = []Region{Beijing, Hangzhou, ShanghaiFinance}
var DCDNSupportRegions = []Region{Hangzhou, APSouthEast1, APNorthEast1, ShanghaiFinance}
var GpdbElasticInstanceSupportRegions = []Region{EUCentral1, Beijing, Hangzhou, Shanghai, Shenzhen, APSouthEast1, APSouthEast5, Hongkong, ShanghaiFinance}
var PolarDBSupportRegions = []Region{Hangzhou, ShanghaiFinance}
var ESSSupportRegions = []Region{Beijing, ShanghaiFinance}
var SimpleApplicationServerNotSupportRegions = []Region{EUCentral1, ShanghaiFinance}
var CRSupportRegions = []Region{WuLanChaBu, APSouthEast2, Hangzhou, ShenZhenFinance1, MEEast1, APSouth1, ShanghaiFinance, APNorthEast1, APSouthEast5, CnNorth2Gov1, Hongkong, Huhehaote, Beijing, ChengDu, APSouthEast3, Shenzhen, USEast1, GuangZhou, Qingdao, Zhangjiakou, EUWest1, Shanghai, APSouthEast1, HeYuan, EUCentral1, USWest1, ShanghaiFinance}
var MSESupportRegions = []Region{Zhangjiakou, USWest1, Shenzhen, ChengDu, Qingdao, APSouthEast3, USEast1, Hangzhou, APNorthEast1, ShenZhenFinance1, APSouthEast1, APSouthEast2, APSouthEast5, Beijing, EUWest1, Shanghai, ShanghaiFinance, Huhehaote, APSouth1, CnNorth2Gov1, Hongkong, HeYuan, EUCentral1, ShanghaiFinance}
var LogResourceSupportRegions = []Region{HeYuan, ShanghaiFinance}
var AliKafkaSupportRegions = []Region{Beijing, CnNorth2Gov1, Qingdao, APSouthEast3, Huhehaote, APSouth1, EUWest1, ShenZhenFinance1, ChengDu, USEast1, USWest1, Hangzhou, Zhangjiakou, Shenzhen, Shanghai, Hongkong, HeYuan, APSouthEast5, APNorthEast1, ShanghaiFinance, APSouthEast1, EUCentral1, ShanghaiFinance}
var BastionhostSupportRegions = []Region{CnNorth2Gov1, Qingdao, ShanghaiFinance, EUCentral1, EUWest1, ChengDu, Shanghai, HeYuan, APNorthEast1, MEEast1, APSouth1, Hongkong, Zhangjiakou, USWest1, APSouthEast1, APSouthEast2, Huhehaote, APSouthEast5, Beijing, Hangzhou, ShenZhenFinance1, APSouthEast3, USEast1, Shenzhen, ShanghaiFinance}
var ACKSystemDiskEncryptionSupportRegions = []Region{Hongkong, ShanghaiFinance}
var DdosBasicSupportRegions = []Region{WuLanChaBu, APSouth1, HeYuan, Shenzhen, MEEast1, APSouthEast1, Huhehaote, CnNorth2Gov1, ChengDu, USEast1, Hangzhou, ShanghaiFinance, ShenZhenFinance1, GuangZhou, APSouthEast2, Beijing, EUCentral1, USWest1, APNorthEast1, Qingdao, APSouthEast3, APSouthEast5, APSouthEast6, Shanghai, Hongkong, Zhangjiakou, EUWest1, ShanghaiFinance}
var TagSupportRegions = []Region{Huhehaote, APSouthEast5, CnNorth2Gov1, HeYuan, APSouthEast2, Beijing, APSouthEast3, USWest1, WuLanChaBu, GuangZhou, MEEast1, ShenZhenFinance1, Shanghai, ShanghaiFinance, EUCentral1, APSouthEast1, USEast1, Hangzhou, Hongkong, Qingdao, Zhangjiakou, Shenzhen, EUWest1, APNorthEast1, APSouth1, ChengDu, ShanghaiFinance}
var GraphDatabaseDbInstanceSupportRegions = []Region{Hangzhou, ShanghaiFinance}
