import React, { useEffect } from "react";
import { StyleSheet, Text, View } from "react-native";
import { useFonts } from "expo-font";
import * as Font from "expo-font";

type Props = {
    title: string;
    button?: React.JSX.Element;
};

export const HeaderComponent = ({ title, button }: Props) => {
    return (
        <View style={styles.HeaderComponent}>
            <Text
                style={[
                    styles.HeaderComponent__title,
                    { color: "#8F5495" },
                ]}
            >
                {title}
            </Text>
            <View style={styles.HeaderComponent__button}>{button}</View>
        </View>
    );
};

const styles = StyleSheet.create({
    HeaderComponent: {
        width: "100%",
        height: 65,
        display: "flex",
        flexDirection: "row",
        justifyContent: "space-between",
        alignItems: "center",
        marginTop: "10%",
    },
    HeaderComponent__title: {
        fontSize: 50,
        marginLeft: 15,
    },
    HeaderComponent__button: {
        height: "100%",
        width: 100,
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
    },
});
