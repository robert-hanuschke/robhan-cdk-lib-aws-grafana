//go:build no_runtime_type_checking

package robhancdklibawsgrafana

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkspaceBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (w *jsiiProxy_WorkspaceBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (w *jsiiProxy_WorkspaceBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceBase) validateGetWorkspaceArnParameters(workspaceId *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceBase) validateGetWorkspaceIdParameters(workspaceArn *string) error {
	return nil
}

func validateWorkspaceBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWorkspaceBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateWorkspaceBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewWorkspaceBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

