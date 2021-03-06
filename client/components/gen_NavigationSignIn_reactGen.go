// Code generated by reactGen. DO NOT EDIT.

package components

import "myitcv.io/react"

type NavigationSignInElem struct {
	react.Element
}

func buildNavigationSignIn(cd react.ComponentDef) react.Component {
	return NavigationSignInDef{ComponentDef: cd}
}

func buildNavigationSignInElem(props NavigationSignInProps, children ...react.Element) *NavigationSignInElem {
	return &NavigationSignInElem{
		Element: react.CreateElement(buildNavigationSignIn, props, children...),
	}
}

func (n NavigationSignInDef) RendersElement() react.Element {
	return n.Render()
}

// SetState is an auto-generated proxy proxy to update the state for the
// NavigationSignIn component.  SetState does not immediately mutate n.State()
// but creates a pending state transition.
func (n NavigationSignInDef) SetState(state NavigationSignInState) {
	n.ComponentDef.SetState(state)
}

// State is an auto-generated proxy to return the current state in use for the
// render of the NavigationSignIn component
func (n NavigationSignInDef) State() NavigationSignInState {
	return n.ComponentDef.State().(NavigationSignInState)
}

// IsState is an auto-generated definition so that NavigationSignInState implements
// the myitcv.io/react.State interface.
func (n NavigationSignInState) IsState() {}

var _ react.State = NavigationSignInState{}

// GetInitialStateIntf is an auto-generated proxy to GetInitialState
func (n NavigationSignInDef) GetInitialStateIntf() react.State {
	return NavigationSignInState{}
}

func (n NavigationSignInState) EqualsIntf(val react.State) bool {
	return n == val.(NavigationSignInState)
}

// IsProps is an auto-generated definition so that NavigationSignInProps implements
// the myitcv.io/react.Props interface.
func (n NavigationSignInProps) IsProps() {}

// Props is an auto-generated proxy to the current props of NavigationSignIn
func (n NavigationSignInDef) Props() NavigationSignInProps {
	uprops := n.ComponentDef.Props()
	return uprops.(NavigationSignInProps)
}

func (n NavigationSignInProps) EqualsIntf(val react.Props) bool {
	return n == val.(NavigationSignInProps)
}

var _ react.Props = NavigationSignInProps{}
