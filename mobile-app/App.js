/**
 * @prettier
 * @flow
 */
import React from "react";
import { StyleSheet, View, ActivityIndicator } from "react-native";
import SafeAreaView from "react-native-safe-area-view";
import Control from "./src/Control";
import LoginView from "./src/LoginView";
import GameView from "./src/GameView";
import IndexView from "./src/IndexView";
import colors from "./src/colors";

type Props = {};
type State = typeof Control.state;

export default class App extends React.Component<Props, State> {
  control: Control;
  constructor(props: Props) {
    super(props);
    this.control = new Control(this);
    this.state = this.control.state;
  }
  componentDidMount() {
    this.control.appDidMount();
  }
  _renderView = () => {
    if (this.state.loading) {
      return <ActivityIndicator size="large" color={colors.text_color} />;
    }
    if (this.state.loggedIn === false) {
      return React.createElement(LoginView, {
        ...this.state,
        control: this.control,
      });
    }
    if (this.state.game) {
      return React.createElement(GameView, {
        ...this.state,
        control: this.control,
      });
    }
    return React.createElement(IndexView, {
      ...this.state,
      control: this.control,
    });
  };
  render() {
    return (
      <SafeAreaView style={styles.safeContainer}>
        <View style={styles.container}>{this._renderView()}</View>
      </SafeAreaView>
    );
  }
}

let styles = StyleSheet.create({
  container: {
    flex: 1,
    alignSelf: "stretch",
    alignItems: "center",
    justifyContent: "center",
  },
  safeContainer: {
    backgroundColor: colors.background,
    flex: 1,
    alignItems: "center",
    justifyContent: "center",
  },
});
