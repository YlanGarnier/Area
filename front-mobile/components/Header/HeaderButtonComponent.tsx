import {StyleSheet, TouchableOpacity, View} from "react-native";
import React from "react";

type Props = {
    onPressAction: () => void,
    logo: () => React.JSX.Element,
    backgroundColor: string
}
export const HeaderButtonComponent = ({onPressAction, logo, backgroundColor}: Props) => {
    return (
        <TouchableOpacity
            style={[styles.HeaderButtonComponent, styles.shadowProp, {backgroundColor: backgroundColor}]}
            onPress={onPressAction}>
            {logo()}
        </TouchableOpacity>
    )
}

const styles = StyleSheet.create({
    HeaderButtonComponent: {
        width: 40,
        height: 40,
        borderRadius: 50,
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
    },
    shadowProp: {
        shadowColor: "rgba(0,0,0,0.74)",
        borderRadius: 100,
        shadowOffset: {
            width: 0,
            height: 1,
        },
        shadowOpacity: 0.36,
        shadowRadius: 6.68,
        elevation: 11,
    },
})
