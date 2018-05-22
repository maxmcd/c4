/**
 * @prettier
 * @flow
 */
import React from "react";
import { StyleSheet, View, TouchableOpacity, Text } from "react-native";
import type { ViewStyleProp } from "react-native/Libraries/StyleSheet/StyleSheet";
import colors from "./colors";

type Props = {
    color: string,
    testID?: string,
    onPress: () => Promise,
    style?: ViewStyleProp,
    textStyle?: {},
    text: string,
    active: boolean,
    weight: "regular" | "bold",
    inverse?: boolean,
};
export default class FormButton extends React.Component<Props> {
    render() {
        let textColor, backgroundColor;
        let color = this.props.inverse ? colors.white : this.props.color;
        let white = this.props.inverse ? this.props.color : colors.white;
        if (this.props.active) {
            textColor = white;
            backgroundColor = color;
        } else {
            textColor = color;
            backgroundColor = white;
        }

        return (
            <TouchableOpacity
                accessible={false}
                testID={this.props.testID}
                activeOpacity={0.7}
                onPress={this.props.onPress}
            >
                <View
                    style={[
                        styles.view,
                        {
                            backgroundColor: backgroundColor,
                        },
                        this.props.style,
                    ]}
                >
                    <Text numberOfLines={1} style={[styles.text]}>
                        {this.props.text}
                    </Text>
                </View>
            </TouchableOpacity>
        );
    }
}

let styles = StyleSheet.create({
    view: {
        justifyContent: "center",
        alignItems: "center",
        height: 44,
        borderRadius: 5,
        borderWidth: 1,
        margin: 0,
        marginBottom: 10,
        padding: 12,
    },
    text: {
        color: "black",
        fontSize: 16,
    },
});
