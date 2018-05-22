/**
 * @prettier
 * @flow
 */
import React from "react";
import { StyleSheet, View, TextInput, Text } from "react-native";

import Control from "./Control";
import colors from "./colors";
import FormButton from "./FormButton";

const text = {
  phone: {
    label: "Enter your phone number",
    button: "Submit",
    placeholder: "(234) 567-8910",
    keyboardType: "phone-pad",
  },
  code: {
    label:
      "We've sent you a text with a confirmation code. Enter the code to proceed.",
    button: "Submit Code",
    placeholder: "0000",
    keyboardType: "phone-pad",
  },
  username: {
    label: "Enter a username",
    button: "Save",
    placeholder: "",
    keyboardType: undefined,
  },
};
const stepKey = {
  "1": "phone",
  "2": "code",
  "3": "username",
};

type Props = {
  control: Control,
  loginInputValue: string,
  loginErrMessage: string,
  ...typeof Control.state,
};

module.exports = class Login extends React.Component<Props> {
  textInputRef: ?TextInput;
  stepNumber = () => {
    if (this.props.codeSubmitted) {
      return "3";
    }
    if (this.props.phoneSubmitted) {
      return "2";
    }
    return "1";
  };
  componentDidMount() {}
  componentDidUpdate(prevProps: Props) {
    // Refocus needed on keyboard type change when proceeding to
    // username step
    if (
      this.props.codeSubmitted &&
      !prevProps.codeSubmitted &&
      this.textInputRef
    ) {
      this.textInputRef.focus();
    }
  }
  render() {
    let step = stepKey[this.stepNumber()];
    return (
      <View style={styles.view}>
        <Text style={styles.loginHeader}>Sign In To C4</Text>
        <View style={styles.loginFormContainer}>
          <View style={styles.loginTextContainer}>
            <Text style={styles.loginLabel}>{text[step].label}</Text>
            {(() => {
              if (this.props.loginErrMessage) {
                return (
                  <Text style={styles.loginError}>
                    {this.props.loginErrMessage}
                  </Text>
                );
              }
            })()}
          </View>
          <TextInput
            testID={"LoginInput"}
            ref={ref => (this.textInputRef = ref)}
            style={styles.loginInput}
            placeholder={text[step].placeholder}
            key={text[step].keyboardType}
            keyboardType={text[step].keyboardType}
            onChangeText={loginInputValue => {
              this.props.control.setState({
                loginInputValue,
              });
            }}
            placeholderTextColor={"rgba(255,255,255,0.3)"}
            autoCapitalize={"none"}
            onSubmitEditing={this.props.control.loginSubmit}
            value={this.props.loginInputValue}
          />
          <FormButton
            testID={"LoginSubmit"}
            inverse={true}
            weight="bold"
            color={colors.white}
            text={text[step].button}
            onPress={this.props.control.loginSubmit}
          />
        </View>
      </View>
    );
  }
};

const styles = StyleSheet.create({
  view: {
    alignSelf: "stretch",
    flex: 1,
  },
  loginTextContainer: {},
  loginInput: {
    color: colors.white,
    borderWidth: 1,
    borderRadius: 3,
    borderColor: colors.white,
    margin: 0,
    marginBottom: 10,
    height: 44,
    paddingLeft: 10,
    paddingRight: 10,
    paddingTop: 12,
    paddingBottom: 12,
    fontSize: 13,
  },
  loginHeader: {
    fontSize: 20,
    color: colors.white,
    textAlign: "center",
    marginTop: 20,
    marginBottom: 100,
  },
  loginFormContainer: {
    paddingRight: 35,
    paddingLeft: 35,
  },
  loginError: {
    fontSize: 14,
    color: "red",
    marginBottom: 5,
  },
  loginLabel: {
    marginBottom: 10,
    fontSize: 20,
    color: "white",
  },
});
