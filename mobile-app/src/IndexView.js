/**
 * @prettier
 * @flow
 */
import React from "react";
import { StyleSheet, Text, View } from "react-native";
import Control from "./Control";

type Props = {
  control: Control,
  ...typeof Control.state,
};
type State = {};

export default class IndexView extends React.Component<Props, State> {
  componentDidMount() {
    if (!this.props.control.User.username) {
      this.props.control.logout();
    }
  }
  render() {
    return (
      <View style={styles.container}>
        <Text>Open up App.js to start working on your app!</Text>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: "center",
    justifyContent: "center",
  },
});
