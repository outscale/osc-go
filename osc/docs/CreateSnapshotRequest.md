# CreateSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A description for the snapshot. | [optional] 
**DryRun** | **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 
**FileLocation** | **string** | The pre-signed URL of the snapshot you want to import from the OSU bucket. | [optional] 
**SnapshotSize** | **int64** | The size of the snapshot created in your account, in bytes. This size must be exactly the same as the source snapshot one. | [optional] 
**SourceRegionName** | **string** | The name of the source Region, which must be the same as the Region of your account. | [optional] 
**SourceSnapshotId** | **string** | The ID of the snapshot you want to copy. | [optional] 
**VolumeId** | **string** | The ID of the volume you want to create a snapshot of. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


