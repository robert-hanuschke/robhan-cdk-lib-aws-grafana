//go:build no_runtime_type_checking

package robhancdklibawsgrafana

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_Workspace) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (w *jsiiProxy_Workspace) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (w *jsiiProxy_Workspace) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (w *jsiiProxy_Workspace) validateGetWorkspaceArnParameters(workspaceId *string) error {
	return nil
}

func (w *jsiiProxy_Workspace) validateGetWorkspaceIdParameters(workspaceArn *string) error {
	return nil
}

func validateWorkspace_FromWorkspaceAttributesParameters(scope constructs.Construct, id *string, attrs *WorkspaceAttributes) error {
	return nil
}

func validateWorkspace_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWorkspace_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateWorkspace_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateWorkspace_IsWorkspaceParameters(x interface{}) error {
	return nil
}

func validateNewWorkspaceParameters(scope constructs.Construct, id *string, props *WorkspaceProps) error {
	return nil
}

