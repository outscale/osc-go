# CreateLoadBalancerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**Listeners** | [**[]ListenerForCreation**](ListenerForCreation.md) | One or more listeners to create. | 
**LoadBalancerName** | **string** | The unique name of the load balancer (32 alphanumeric or hyphen characters maximum, but cannot start or end with a hyphen). | 
**LoadBalancerType** | **string** | The type of load balancer: &#x60;internet-facing&#x60; or &#x60;internal&#x60;. Use this parameter only for load balancers in a Net. | [optional] 
**SecurityGroups** | **[]string** | (Net only) One or more IDs of security groups you want to assign to the load balancer. If not specified, the default security group of the Net is assigned to the load balancer. | [optional] 
**Subnets** | **[]string** | One or more IDs of Subnets in your Net that you want to attach to the load balancer. | [optional] 
**SubregionNames** | **[]string** | One or more names of Subregions (currently, only one Subregion is supported). This parameter is not required if you create a load balancer in a Net. To create an internal load balancer, use the &#x60;LoadBalancerType&#x60; parameter. | [optional] 
**Tags** | [**[]ResourceTag**](ResourceTag.md) | One or more tags assigned to the load balancer. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


