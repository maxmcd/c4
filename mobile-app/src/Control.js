/**
 * @prettier
 * @flow
 */
import React from "react";
import { AppState } from "react-native";
import User from "./User";
// import Game from "./Game";

const state = {
    loggedIn: false,
    loading: true,
    loginType: null,
    sendingLogin: false,
    phoneSubmitted: false,
    codeSubmitted: false,
    messages: [],
    game: null,
    loginInputValue: "",
};

export default class Control {
    setState: ({}) => void;
    component: React.Component<{}, typeof state>;
    state: {
        loggedIn: boolean,
        loading: boolean,
        loginType?: ?string,
        sendingLogin: boolean,
        phoneSubmitted: boolean,
        codeSubmitted: boolean,
        messages: Array<{}>,
        loginInputValue: string,
    };
    User: User;
    state = state;
    static state = state;
    constructor(component: React.Component<{}, typeof state>) {
        this.component = component;
        this.setState = (state: {}) => {
            this.state = { ...this.state, ...state };
            component.setState(state);
        };
        this.User = new User(this);
        AppState.addEventListener("change", this.handleAppStateChange);
    }
    handleAppStateChange = (nextAppState: string) => {
        console.log(nextAppState);
        // if (nextAppState.match(/background/)) {
        //   this.exitVideoStream(nextAppState);
        // }
    };
    loginSubmit = async () => {
        if (this.state.sendingLogin) {
            return;
        }
        this.setState({
            sendingLogin: true,
            loginErrMessage: "",
        });
        let errMessage;
        if (this.state.phoneSubmitted && this.state.codeSubmitted) {
            errMessage = await this.User.updateUser(this.state.loginInputValue);
        } else if (this.state.phoneSubmitted) {
            errMessage = await this.User.checkCode(this.state.loginInputValue);
        } else {
            errMessage = await this.User.authenticate(
                this.state.loginInputValue,
            );
        }
        this.setState({ sendingLogin: false });
        if (errMessage) {
            return this.setState({ loginErrMessage: errMessage.message });
        }
        if (this.state.phoneSubmitted) {
            return this.setState({
                loggedIn: !!this.User.username,
                loginInputValue: "",
                codeSubmitted: true,
            });
        }
        this.setState({
            phoneSubmitted: true,
            loginInputValue: "",
        });
    };
    appDidMount = async () => {
        const userSave = await this.User.loadUser();
        this.setState({
            loggedIn: !!userSave,
            loading: false,
        });
    };
    logout = () => {
        this.User.logout();
        this.setState({
            loggedIn: false,
        });
    };
}
